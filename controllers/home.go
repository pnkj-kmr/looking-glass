package controllers

import (
	"fmt"

	"github.com/pnkj-kmr/looking-glass/gateways"
	"github.com/pnkj-kmr/looking-glass/repos"
	"github.com/pnkj-kmr/looking-glass/repos/jsondb"
	"github.com/pnkj-kmr/looking-glass/utils"
	"go.uber.org/zap"
)

// Execute func helps to execute cmd on src ip server
func Execute(src, dst string, proto int, out chan<- []byte) (err error) {
	utils.L.Debug("Entering into execute function", zap.Int("proto", proto), zap.String("src", src), zap.String("dst", dst))
	collection, err := jsondb.New(ipDetailTable)
	if err != nil {
		return
	}
	d, err := collection.Get(src)
	if err != nil {
		return
	}
	var ip repos.IPInfo
	err = ip.FromJSON(d)
	if err != nil {
		return
	}

	// creating connection object for query
	_dail := fmt.Sprintf("%s:%d", ip.IP, ip.Port)
	conn := gateways.NewConn(ip.IP, _dail, ip.Usr, ip.Pwd, ip.Vendor)
	utils.L.Debug("Created new connection", zap.String("conn", _dail))
	switch proto {
	case 1, 2:
		err := conn.Ping(getProto(proto), dst, out)
		if err != nil {
			return err
		}
	case 3, 4:
		err := conn.Traceroute(getProto(proto), dst, out)
		if err != nil {
			return err
		}
	case 5, 6:
		err := conn.BGP(getProto(proto), dst, out)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid arguments")
	}
	return nil
}

func getProto(proto int) string {
	utils.L.Debug("Protocal checking...", zap.Int("proto", proto))
	switch proto {
	case 1, 3, 5:
		return gateways.IPv4
	default:
		return gateways.IPv6
	}
}
