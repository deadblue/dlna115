package soap

import (
	"bytes"
	"encoding/xml"
)

func Extract(data []byte) (name string, begin, end int64) {
	d := xml.NewDecoder(bytes.NewReader(data))
	var tmpOffset int64
	for {
		t, err := d.RawToken()
		if err != nil || t == nil {
			break
		}
		switch t := t.(type) {
		case xml.StartElement:
			if t.Name.Space == "s" && t.Name.Local == "Body" {
				begin = d.InputOffset()
			} else {
				if begin > 0 && name == "" {
					name = t.Name.Local
				}
			}
		case xml.EndElement:
			if t.Name.Space == "s" && t.Name.Local == "Body" {
				end = tmpOffset
			}
		}
		tmpOffset = d.InputOffset()
	}
	return
}
