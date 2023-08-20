package soap

import "encoding/xml"

type Envelope struct {
	XMLName       xml.Name `xml:"s:Envelope"`
	XmlNs         string   `xml:"xmlns:s,attr"`
	EncodingStyle string   `xml:"s:encodingStyle,attr"`
	Body          struct {
		Data any
	} `xml:"s:Body"`
}

func (e *Envelope) Init() *Envelope {
	e.XmlNs = "http://schemas.xmlsoap.org/soap/envelope/"
	e.EncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
	return e
}

func (e *Envelope) SetData(data any) {
	e.Body.Data = data
}
