package upnp

const (
	DeviceTypeMediaServer1   = "urn:schemas-upnp-org:device:MediaServer:1"
	DeviceTypeMediaRenderer1 = "urn:schemas-upnp-org:device:MediaRenderer:1"
)

type Device interface {
	// DeviceType return device type
	DeviceType() string
	// USN return unique service name of device.
	USN() string

	// GetDescURL return device desc URL on ip.
	GetDescURL(ip string) string
}
