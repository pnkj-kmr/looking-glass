package gateways

import "fmt"

func ciscoPingCmd(proto, dst string) (cmd string) {
	cmd = fmt.Sprintf("ping %s", dst)
	return cmd
}

func ciscoTracerouteCmd(proto, dst string) (cmd string) {
	cmd = fmt.Sprintf("traceroute %s", dst)
	return cmd
}

func ciscoBGPCmd(proto, dst string) (cmd string) {
	if proto == IPv4 {
		cmd = fmt.Sprintf("sh bgp ipv4 unicast")
	} else {
		cmd = fmt.Sprintf("sh bgp ipv6 unicast")
	}
	return cmd
}
