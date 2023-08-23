package ssdp

import (
	"fmt"
	"net"

	"github.com/deadblue/dlna115/internal/upnp"
)

func NotifyAvailable(
	uuid string,
	deviceType string,
	serverPort int,
	descUrlPath string,
) (err error) {
	// Get net interfaces
	nifs, err := net.Interfaces()
	if err != nil {
		return
	}

	// Prepare request
	req := &Request{
		Method: methodNotify,
	}
	req.SetHeader("HOST", serverHost)
	req.SetHeader("CACHE-CONTROL", "max-age=3600")
	req.SetHeader("NT", deviceType)
	req.SetHeader("NTS", notifyAlive)
	req.SetHeader("SERVER", upnp.ServerTag)
	usn := fmt.Sprintf("uuid:%s:%s", uuid, deviceType)
	req.SetHeader("USN", usn)

	for _, nif := range nifs {
		// Skip inactive net interface
		if (nif.Flags&net.FlagUp == 0) || (nif.Flags&net.FlagRunning == 0) {
			continue
		}
		addrs, err := nif.Addrs()
		if err != nil {
			continue
		}
		// Find IPv4 Address
		for _, addr := range addrs {
			netip, ok := addr.(*net.IPNet)
			if !ok {
				continue
			}
			ipv4 := netip.IP.To4()
			if ipv4 == nil {
				continue
			}
			location := fmt.Sprintf(
				"http://%s:%d/%s",
				ipv4.String(), serverPort, descUrlPath,
			)
			req.SetHeader("LOCATION", location)
			// Send broadcast
			conn, err := net.DialUDP(serverAddr.Network(), nil, serverAddr)
			if err != nil {
				continue
			}
			req.WriteTo(conn)
		}
	}
	return
}
