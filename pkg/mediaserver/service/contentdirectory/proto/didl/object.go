package didl

import "encoding/xml"

type Object struct {
	// Attributes
	ID         string `xml:"id,attr"`
	ParentID   string `xml:"parentID,attr"`
	Restricted string `xml:"restricted,attr"`
	// Children
	Title   string `xml:"dc:title"`
	Creator string `xml:"creator,omitempty"`
	Class   string `xml:"upnp:class"`
}

type Res struct {
	XMLName xml.Name `xml:"res"`

	ProtocolInfo    string `xml:"protocolInfo,attr"`
	Resolution      string `xml:"resolution,attr"`
	Size            int64  `xml:"size,attr"`
	Bitrate         int    `xml:"bitrate,attr"`
	Duration        string `xml:"duration,attr"`
	NrAudioChannels int    `xml:"nrAudioChannels,attr"`
	SampleFrequency int    `xml:"sampleFrequency,attr"`

	URL string `xml:",chardata"`
}
