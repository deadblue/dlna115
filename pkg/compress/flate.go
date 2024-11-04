package compress

import (
	"bytes"
	"compress/flate"
	"encoding/base64"
	"io"
)

func Encode(source string) string {
	buf := &bytes.Buffer{}
	fw, _ := flate.NewWriter(buf, flate.BestSpeed)
	fw.Write([]byte(source))
	fw.Close()
	return base64.RawURLEncoding.EncodeToString(buf.Bytes())
}

func Decode(source string) (result string, err error) {
	raw, err := base64.RawURLEncoding.DecodeString(source)
	if err != nil {
		return
	}
	fr := flate.NewReader(bytes.NewReader(raw))
	raw, err = io.ReadAll(fr)
	if err == nil {
		result = string(raw)
	}
	return
}
