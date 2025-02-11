package gateways

import "fmt"

func huaweiPingCmd(proto, dst string) string {
	var cmd string
	if proto == IPv4 {
		cmd = fmt.Sprintf("ping %s", dst)
	} else {
		cmd = fmt.Sprintf("ping ipv6 %s", dst)
	}
	return cmd
}

func huaweiTracerouteCmd(proto, dst string) string {
	var cmd string
	if proto == IPv4 {
		cmd = fmt.Sprintf("tracert -m 6 %s", dst)
	} else {
		cmd = fmt.Sprintf("tracert ipv6 -v %s", dst)
	}
	return cmd
}

func huaweiBGPCmd(proto, dst string) string {
	var cmd string
	if proto == IPv4 {
		cmd = fmt.Sprintf("display bgp routing-table %s | no-more", dst)
	} else {
		cmd = fmt.Sprintf("display bgp ipv6 routing-table %s | no-more", dst)
	}
	return cmd
}
