package upnp

import "encoding/xml"

const (
	NamespaceDevice = "urn:schemas-upnp-org:device-1-0"

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

type DeviceIcon struct {
	MimeType string `xml:"mimetype"`
	Width    int    `xml:"width"`
	Height   int    `xml:"height"`
	Depth    int    `xml:"depth"`
	URL      string `xml:"url"`
}

type DeviceService struct {
	ServiceType string `xml:"serviceType"`
	ServiceId   string `xml:"serviceId"`
	ScpdURL     string `xml:"SCPDURL"`
	ControlURL  string `xml:"controlURL"`
	EventSubURL string `xml:"eventSubURL"`
}

// Description is the whole device description document
type DeviceDesc struct {
	XMLName     xml.Name `xml:"root"`
	Xmlns       string   `xml:"xmlns,attr"`
	SpecVersion struct {
		Major int `xml:"major"`
		Minor int `xml:"minor"`
	} `xml:"specVersion"`
	Device struct {
		DeviceType string `xml:"deviceType"`

		// Basic information
		UDN string `xml:"UDN"`
		UPC string `xml:"UPC,omitempty"`

		FriendlyName string `xml:"friendlyName"`
		SerialNumber string `xml:"serialNumber,omitempty"`

		// Manufacturer information
		Manufacturer    string `xml:"manufacturer"`
		ManufacturerURL string `xml:"manufacturerURL,omitempty"`

		// Model information
		ModelName        string `xml:"modelName"`
		ModelDescription string `xml:"modelDescription,omitempty"`
		ModelNumber      string `xml:"modelNumber,omitempty"`
		ModelURL         string `xml:"modelURL,omitempty"`

		// Icon list
		IconList struct {
			Icons []DeviceIcon `xml:"icon"`
		} `xml:"iconList"`

		// Service list
		ServiceList struct {
			Services []DeviceService `xml:"service"`
		} `xml:"serviceList"`

		PresentationURL string `xml:"presentationURL,omitempty"`
	} `xml:"device"`
}

// Init fills fixed fields into Description document
func (dd *DeviceDesc) Init(deviceType string) *DeviceDesc {
	dd.Xmlns = NamespaceDevice
	dd.SpecVersion.Major = 1
	dd.SpecVersion.Minor = 0
	dd.Device.DeviceType = deviceType
	return dd
}
