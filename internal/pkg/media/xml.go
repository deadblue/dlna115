package media

import (
	"encoding/xml"
	"net/http"
	"strconv"
)

var (
	header = []byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n")

	headerSize = len(header)

	xmlMimeType = "text/xml"
)

func renderXml(rw http.ResponseWriter, doc any, withXmlHeader bool) (err error) {
	body, err := xml.Marshal(doc)
	if err != nil {
		return
	}
	bodySize := len(body)
	if withXmlHeader {
		bodySize += headerSize
	}
	rw.Header().Set("Content-Type", xmlMimeType)
	rw.Header().Set("Content-Length", strconv.Itoa(bodySize))
	rw.WriteHeader(200)
	if withXmlHeader {
		if _, err = rw.Write(header); err != nil {
			return
		}
	}
	_, err = rw.Write(body)
	return
}

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
