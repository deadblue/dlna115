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
			Icons []Icon `xml:"icon"`
		} `xml:"iconList"`

		// Service list
		ServiceList struct {
			Services []Service `xml:"service"`
		} `xml:"serviceList"`

		PresentationURL string `xml:"presentationURL,omitempty"`
	} `xml:"device"`
}

// Init fills fixed fields into Description document
func (d *Description) Init(deviceType string) *Description {
	d.Xmlns = "urn:schemas-upnp-org:device-1-0"
	d.SpecVersion.Major = 1
	d.SpecVersion.Minor = 0
	d.Device.DeviceType = deviceType
	return d
}
