package soap

import (
	"bytes"
	"encoding/xml"
)

func Extract(data []byte) (beginOffset, endOffset int64, name string) {
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
				beginOffset = d.InputOffset()
			} else {
				if beginOffset > 0 && name == "" {
					name = t.Name.Local
				}
			}
		case xml.EndElement:
			if t.Name.Space == "s" && t.Name.Local == "Body" {
				endOffset = tmpOffset
			}
		}
		tmpOffset = d.InputOffset()
	}
	return
}
