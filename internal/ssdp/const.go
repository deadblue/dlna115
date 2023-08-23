package ssdp

import "net"

const (
	serverHost = "239.255.255.250:1900"

	methodNotify = "NOTIFY"
	methodSearch = "M-SEARCH"

	notifyAlive  = "ssdp:alive"
	notifyByebye = "ssdp:byebye"

	// manDiscover = "\"ssdp:discover\""

	// searchAll        = "ssdp:all"
	// searchRootDevice = "upnp:rootdevice"
)

var (
	serverAddr, _ = net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
)
