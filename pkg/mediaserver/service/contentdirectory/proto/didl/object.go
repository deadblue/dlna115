package didl

import "encoding/xml"

type Object struct {
	// Attributes
	ID         string `xml:"id,attr"`
	ParentID   string `xml:"parentID,attr"`
	Restricted string `xml:"restricted,attr"`
	// Children
	Title   string `xml:"dc:title"`
	Class   string `xml:"upnp:class"`
	Creator string `xml:"creator,omitempty"`
}

type Res struct {
	XMLName xml.Name `xml:"res"`

	ProtocolInfo string `xml:"protocolInfo,attr"`
	// File size
	Size int64 `xml:"size,attr"`
	// Media duration
	Duration string `xml:"duration,attr,omitempty"`
	// Media bitrate
	Bitrate int `xml:"bitrate,attr,omitempty"`
	// Audio channels
	NrAudioChannels int `xml:"nrAudioChannels,attr,omitempty"`
	// Audio samplerate
	SampleFrequency int `xml:"sampleFrequency,attr,omitempty"`
	// Video resolution
	Resolution string `xml:"resolution,attr,omitempty"`

	URL string `xml:",chardata"`
}
