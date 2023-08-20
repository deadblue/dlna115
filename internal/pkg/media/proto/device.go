package proto

import (
	"encoding/xml"
)

type Icon struct {
	MimeType string `xml:"mimetype"`
	Width    int    `xml:"width"`
	Height   int    `xml:"height"`
	Depth    int    `xml:"depth"`
	URL      string `xml:"url"`
}

type Service struct {
	ServiceType string `xml:"serviceType"`
	ServiceId   string `xml:"serviceId"`
	ScpdURL     string `xml:"SCPDURL"`
	ControlURL  string `xml:"controlURL"`
	EventSubURL string `xml:"eventSubURL"`
}

// Description is the whole device description document
type Description struct {
	XMLName     xml.Name `xml:"root"`
	Xmlns       string   `xml:"xmlns,attr"`
	SpecVersion struct {
		Major int `xml:"major"`
		Minor int `xml:"minor"`
	} `xml:"specVersion"`
	Device struct {
		// Should be "urn:schemas-upnp-org:device:MediaServer:1"
		DeviceType string `xml:"deviceType"`

		// Basic information
		UDN             string `xml:"UDN"`
		FriendlyName    string `xml:"friendlyName"`
		PresentationURL string `xml:"presentationURL"`
		SerialNumber    string `xml:"serialNumber"`

		// Model information
		ModelName        string `xml:"modelName"`
		ModelNumber      string `xml:"modelNumber"`
		ModelDescription string `xml:"modelDescription"`
		ModelURL         string `xml:"modelURL"`
		Manufacturer     string `xml:"manufacturer"`
		ManufacturerURL  string `xml:"manufacturerURL"`

		// Icon list
		IconList struct {
			Icons []Icon `xml:"icon"`
		} `xml:"iconList"`

		// Service list
		ServiceList struct {
			Services []Service `xml:"service"`
		} `xml:"serviceList"`
	} `xml:"device"`
}

// Init fills fixed fields into Description document
func (d *Description) Init() *Description {
	d.Xmlns = "urn:schemas-upnp-org:device-1-0"
	d.SpecVersion.Major = 1
	d.SpecVersion.Minor = 0
	d.Device.DeviceType = "urn:schemas-upnp-org:device:MediaServer:1"
	return d
}
