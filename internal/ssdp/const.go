package ssdp

import (
	"net"
)

const (
	serverHost = "239.255.255.250:1900"

	methodNotify = "NOTIFY"
	methodSearch = "M-SEARCH"

	headerHost                = "HOST"
	headerServer              = "SERVER"
	headerLocation            = "LOCATION"
	headerCacheControl        = "CACHE-CONTROL"
	headerExtension           = "EXT"
	headerSearchTarget        = "ST"
	headerUniqueServiceName   = "USN"
	headerNotificationType    = "NT"
	headerNotificationSubType = "NTS"

	notifyAlive  = "ssdp:alive"
	notifyByebye = "ssdp:byebye"

	searchAll = "ssdp:all"
)

var (
	multicastAddr, _ = net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
)
