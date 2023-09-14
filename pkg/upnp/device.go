package upnp

const (
	DeviceTypeMediaServer1   = "urn:schemas-upnp-org:device:MediaServer:1"
	DeviceTypeMediaRenderer1 = "urn:schemas-upnp-org:device:MediaRenderer:1"
)

// Device interface should be impled by an UPnP device.
type Device interface {
	// DeviceType returns UPnP device type that device impled.
	DeviceType() string
	// USN returns unique service name of this device.
	DeviceUSN() string
	// GetDeviceDescURL returns device desc URL on given ip.
	GetDeviceDescURL(ip string) string
	// DeviceIsTarget indicates whether device matches the searching target.
	DeviceIsTarget(targetType string) bool
}
