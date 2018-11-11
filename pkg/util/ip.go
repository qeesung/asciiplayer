package util

import (
	"net"
)

// GetIPList get the machine ip address list
func GetIPList() (IPList []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	IPList = make([]string, 0)
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok {
			if ipnet.IP.To4() != nil {
				IPList = append(IPList, ipnet.IP.String())
			}
		}
	}
	return
}
