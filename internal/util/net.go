package util

import (
	"io"
	"net"
)

func InterfaceHasIPv4(nif *net.Interface) bool {
	if addrs, err := nif.Addrs(); err == nil {
		for _, addr := range addrs {
			if addr.Network() != "ip+net" {
				continue
			}
			ipnet := addr.(*net.IPNet)
			if ipnet.IP.To4() != nil {
				return true
			}
		}
	}
	return false
}

// ForAllIPs calls fn with all local IPs
func ForAllIPs(skipLoopback bool, fn func(net.IP)) {
	nis, err := net.Interfaces()
	if err != nil {
		return
	}
	// Traverse netifs
	for _, ni := range nis {
		// Skip inactive netif
		if ni.Flags&(net.FlagUp|net.FlagRunning) == 0 {
			continue
		}
		// Skip loopback
		if skipLoopback && (ni.Flags&net.FlagLoopback != 0) {
			continue
		}
		// Get address from netif
		addrs, err := ni.Addrs()
		if err != nil {
			continue
		}
		// Traverse addresses
		for _, addr := range addrs {
			ipnet, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			ipv4 := ipnet.IP.To4()
			if ipv4 == nil {
				continue
			}
			fn(ipv4)
		}
	}
}

func Broadcast(msg io.WriterTo, localIp net.IP, remoteAddr *net.UDPAddr) (err error) {
	localAddr := &net.UDPAddr{
		// Set local IP to force system select correct network
		IP: localIp,
		// Let system choose an available port
		Port: 0,
	}
	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	if err != nil {
		return
	}
	_, err = msg.WriteTo(conn)
	return
}
