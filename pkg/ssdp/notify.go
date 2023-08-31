package ssdp

import (
	"net"

	"github.com/deadblue/dlna115/pkg/upnp"
	"github.com/deadblue/dlna115/pkg/util"
)

// NotifyDeviceAvailable broadcasts a device available advertisement on all
// networks.
func NotifyDeviceAvailable(device upnp.Device) (err error) {
	// Prepare request
	req := &Request{
		Method: methodNotify,
	}
	req.SetHeader(headerHost, serverHost)
	req.SetHeader(headerCacheControl, "max-age=3600")
	req.SetHeader(headerNotificationType, device.DeviceType())
	req.SetHeader(headerNotificationSubType, notifyAlive)
	req.SetHeader(headerServer, upnp.ServerName)
	req.SetHeader(headerUniqueServiceName, device.DeviceUSN())

	util.ForAllIPs(true, func(ip net.IP) {
		req.SetHeader(headerLocation, device.GetDeviceDescURL(ip.String()))
		util.Broadcast(req, ip, multicastAddr)
	})
	return
}

// NotifyDeviceAvailable broadcasts a device unavailable advertisement on all
// networks.
func NotifyDeviceUnavailable(device upnp.Device) (err error) {
	// Prepare request
	req := &Request{
		Method: methodNotify,
	}
	req.SetHeader(headerHost, serverHost)
	req.SetHeader(headerNotificationType, device.DeviceType())
	req.SetHeader(headerNotificationSubType, notifyByebye)
	req.SetHeader(headerUniqueServiceName, device.DeviceUSN())

	util.ForAllIPs(true, func(ip net.IP) {
		util.Broadcast(req, ip, multicastAddr)
	})
	return
}
