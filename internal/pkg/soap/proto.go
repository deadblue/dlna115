package soap

import "encoding/xml"

type Envelope[B any] struct {
	XMLName       xml.Name `xml:"s:Envelope"`
	XmlNs         string   `xml:"xmlns:s,attr"`
	EncodingStyle string   `xml:"s:encodingStyle,attr"`
	Body          struct {
		Data B
	} `xml:"s:Body"`
}

func (r *Envelope[B]) Init() *Envelope[B] {
	r.XmlNs = "http://schemas.xmlsoap.org/soap/envelope/"
	r.EncodingStyle = "http://schemas.xmlsoap.org/soap/encoding/"
	return r
}
