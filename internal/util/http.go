package util

import (
	"encoding/xml"
	"net/http"
	"strconv"
)

const (
	mimeType = "text/xml"
)

func RenderXML(rw http.ResponseWriter, doc any) {
	if body, err := xml.Marshal(doc); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(err.Error()))
	} else {
		SendXML(rw, body)
	}
}

func SendXML(rw http.ResponseWriter, body []byte) {
	rw.Header().Set("Content-Type", mimeType)
	rw.Header().Set("Content-Length", strconv.Itoa(len(body)))
	rw.WriteHeader(http.StatusOK)
	rw.Write(body)
}
