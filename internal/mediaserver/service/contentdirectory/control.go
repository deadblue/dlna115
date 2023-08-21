package contentdirectory

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory/proto"
	"github.com/deadblue/dlna115/internal/soap"
	"github.com/deadblue/dlna115/internal/xmlhttp"
	"github.com/deadblue/elevengo"
)

func renderError(rw http.ResponseWriter, status int, err error) {
	if status == 0 {
		status = http.StatusInternalServerError
	}
	rw.WriteHeader(status)
	if err != nil {
		rw.Write([]byte(err.Error()))
	}
}

func (s *Service) HandleControl(rw http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		renderError(rw, http.StatusMethodNotAllowed, nil)
		return
	}
	// Read request payload
	payload, err := io.ReadAll(req.Body)
	if err != nil {
		renderError(rw, http.StatusBadRequest, err)
		return
	}
	// Parse request
	name, begin, end := soap.Extract(payload)
	if name == "" {
		renderError(rw, http.StatusBadRequest, nil)
		return
	}
	// Dispatch request
	var resp any
	switch name {
	case ActionBrowse:
		resp, err = s.handleActionBrowse(payload[begin:end])
	default:
		// TODO
	}
	// Render response
	if err != nil {
		renderError(rw, http.StatusInternalServerError, err)
	} else {
		envelope := (&soap.Envelope{}).Init(resp)
		xmlhttp.RenderXML(rw, envelope)
	}
}

func (s *Service) handleActionBrowse(payload []byte) (ret any, err error) {
	req := &proto.BrowseReq{}
	if err = xml.Unmarshal(payload, req); err != nil {
		return
	}
	resp := (&proto.BrowseResp{}).Init()
	// Get file list
	it, err := s.ea.FileIterate(req.ObjectID)
	if err != nil {
		return
	}
	file := &elevengo.File{}
	for ; err == nil; err = it.Next() {
		it.Get(file)
	}

	ret = resp
	return
}
