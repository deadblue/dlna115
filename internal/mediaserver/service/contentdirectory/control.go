package contentdirectory

import (
	"encoding/xml"
	"io"
	"net/http"

	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory/proto"
	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory/proto/didl"
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
	result := (&didl.Document{}).Init()
	// Get file list
	it, err := s.ea.FileIterate(req.ObjectID)
	if err != nil {
		return
	}
	for ; err == nil; err = it.Next() {
		file := &elevengo.File{}
		it.Get(file)
		if file.IsDirectory {
			cont := (&didl.StorageFolderContainer{}).Init()
			cont.ID = file.FileId
			cont.ParentID = req.ObjectID
			cont.StorageUsed = -1
			result.AppendContainer(cont)
			resp.TotalMatches += 1
		} else if file.IsVideo {
			item := (&didl.VideoItem{}).Init()
			item.ID = file.FileId
			item.ParentID = req.ObjectID
			item.Title = file.Name
			item.Res.ProtocolInfo = "http-get:*:video/mp4:*"
			item.Res.Size = file.Size
			item.Res.URL = s.ff.GetAccessURL(file.PickCode)
			resp.TotalMatches += 1
		} else {
			continue
		}
	}
	resp.NumberReturned = resp.TotalMatches
	resp.SetResult(result)
	ret = resp
	return
}
