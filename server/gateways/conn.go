package gateways

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/pnkj-kmr/looking-glass/utils"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/knownhosts"
)

const (
	// DefaultTimeout connection timeout
	DefaultTimeout time.Duration = 10 * time.Second
	// IPv4 string
	IPv4 string = "ipv4"
	// IPv6 string
	IPv6 string = "ipv6"
	// Cisco string
	Cisco string = "Cisco"
	// Huawei string
	Huawei string = "Huawei"
)

type conn struct {
	src      string
	dial     string
	username string
	password string
	vendor   string
	timeout  time.Duration
}

// NewConn -  new conn refernce object
// which contains all conn related data
func NewConn(src, dial, usr, pwd, vendor string) Packet {
	return &conn{
		src:      src,
		dial:     dial,
		username: usr,
		password: pwd,
		vendor:   vendor,
		timeout:  DefaultTimeout,
	}
}

// Ping helps to do ping to target from conn
func (c *conn) Ping(proto, dst string, out chan<- []byte) error {
	var cmd string

	switch c.vendor {
	case Cisco:
		cmd = ciscoPingCmd(proto, dst)
	case Huawei:
		cmd = huaweiPingCmd(proto, dst)
	default:
		cmd = fmt.Sprintf("ping -c 4 %s", dst)
	}

	err := c.execute(cmd, out)
	if err != nil {
		return err
	}
	return nil
}

// Traceroute helps to do traceroute to target from conn
func (c *conn) Traceroute(proto, dst string, out chan<- []byte) error {
	var cmd string

	switch c.vendor {
	case Cisco:
		cmd = ciscoTracerouteCmd(proto, dst)
	case Huawei:
		cmd = huaweiTracerouteCmd(proto, dst)
	default:
		cmd = fmt.Sprintf("traceroute %s", dst)
	}

	err := c.execute(cmd, out)
	if err != nil {
		return err
	}
	return nil
}

// BGP helps to do BGP to target from conn
func (c *conn) BGP(proto, dst string, out chan<- []byte) error {
	var cmd string

	switch c.vendor {
	case Cisco:
		cmd = ciscoBGPCmd(proto, dst)
	case Huawei:
		cmd = huaweiBGPCmd(proto, dst)
	default:
		cmd = fmt.Sprintf("sh bgp")
	}

	err := c.execute(cmd, out)
	if err != nil {
		return err
	}
	return nil
}

// main method which does the ssh connection with src device and
// run the command on src device and return the output or error if any
func (c *conn) execute(cmd string, out chan<- []byte) (err error) {
	routineStarted := false
	defer func() {
		if !routineStarted {
			utils.L.Error("Exiting Error", zap.String("err", err.Error()))
			close(out)
		}
	}()

	client, err := c.dialConnection()
	if err != nil {
		utils.L.Error("Connection Error", zap.String("err", err.Error()), zap.String("conn", c.dial))
		return err
	}
	defer client.Close()

	session, err := client.NewSession()
	if err != nil {
		utils.L.Error("Session Error", zap.String("err", err.Error()), zap.String("src", c.src))
		return err
	}
	defer session.Close()

	stdout, err := session.StdoutPipe()
	if err != nil {
		utils.L.Error("Session outpipe", zap.String("err", err.Error()))
		return err
	}

	stderr, err := session.StderrPipe()
	if err != nil {
		utils.L.Error("Session errpipe", zap.String("err", err.Error()))
		return err
	}

	// calling listeners
	outChan := make(chan []byte)
	errChan := make(chan []byte)
	outputListener(stdout, stderr, outChan, errChan, out)
	routineStarted = true

	if err := session.Run(cmd); err != nil {
		utils.L.Error("Command Error", zap.String("err", err.Error()), zap.String("src", c.src), zap.String("cmd", cmd))
		return err
	}

	return nil
}

func (c *conn) dialConnection() (client *ssh.Client, err error) {
	homeDir := os.Getenv("HOME")
	if homeDir == "" {
		homeDir = "/root"
	}

	key, err := ioutil.ReadFile(filepath.Join(homeDir, ".ssh", "id_rsa"))
	if err != nil {
		utils.L.Error("unable to read private key", zap.String("err", err.Error()))
		return
	}

	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		utils.L.Error("unable to parse private key", zap.String("err", err.Error()))
		return
	}

	hostKeyCallback, err := knownhosts.New(filepath.Join(homeDir, ".ssh", "known_hosts"))
	if err != nil {
		utils.L.Error("could not create hostkeycallback function", zap.String("err", err.Error()))
		return
	}

	// var keyErr *knownhosts.KeyError

	config := &ssh.ClientConfig{
		User:    c.username,
		Timeout: c.timeout,
		Config: ssh.Config{
			KeyExchanges: preferredKexAlgos,
			Ciphers:      preferredCiphers,
			MACs:         supportedMACs,
		},
		Auth: []ssh.AuthMethod{
			ssh.Password(c.password),
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: hostKeyCallback,
		// HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error { return nil }),
		// HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// HostKeyCallback: ssh.HostKeyCallback(func(host string, remote net.Addr, pubKey ssh.PublicKey) error {
		// 	kh := checkKnownHosts()
		// 	hErr := kh(host, remote, pubKey)
		// 	if errors.As(hErr, &keyErr) && len(keyErr.Want) > 0 {
		// 		utils.L.Error("WARNING: %v is not a key of %s, either a MiTM attack or %s has reconfigured the host pub key.", string(pubKey.Marshal()), host, host)
		// 		return keyErr
		// 	} else if errors.As(hErr, &keyErr) && len(keyErr.Want) == 0 {
		// 		utils.L.Error("WARNING: %s is not trusted, adding this key: %q to known_hosts file.", host, string(pubKey.Marshal()))
		// 		return addHostKey(host, remote, pubKey)
		// 	}
		// 	utils.L.Error("Pub key exists for %s.", host)
		// 	return nil
		// }),
	}

	client, err = ssh.Dial("tcp", c.dial, config)
	return
}

// func checkKnownHosts() ssh.HostKeyCallback {
// 	createKnownHosts()
// 	kh, e := knownhosts.New(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
// 	utils.L.Debug(e)
// 	return kh
// }

// func createKnownHosts() {
// 	f, fErr := os.OpenFile(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"), os.O_CREATE, 0600)
// 	if fErr != nil {
// 		utils.L.Fatal(fErr)
// 	}
// 	f.Close()
// }

// func addHostKey(host string, remote net.Addr, pubKey ssh.PublicKey) error {
// 	// add host key if host is not found in known_hosts, error object is return, if nil then connection proceeds,
// 	// if not nil then connection stops.
// 	khFilePath := filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts")

// 	f, fErr := os.OpenFile(khFilePath, os.O_APPEND|os.O_WRONLY, 0600)
// 	if fErr != nil {
// 		return fErr
// 	}
// 	defer f.Close()

// 	knownHosts := knownhosts.Normalize(remote.String())
// 	_, fileErr := f.WriteString(knownhosts.Line([]string{knownHosts}, pubKey))
// 	return fileErr
// }
