package mediaserver

import (
	"encoding/xml"
)

var (
	header = []byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n")

	headerSize = len(header)
)

func marshalXml(v any) (data []byte, err error) {
	body, err := xml.Marshal(v)
	if err != nil {
		return
	}
	data = make([]byte, headerSize+len(body))
	copy(data, header)
	copy(data[headerSize:], body)
	return
}
