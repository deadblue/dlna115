package contentdirectory

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory/proto"
	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory/proto/didl"
	"github.com/deadblue/dlna115/internal/mediaserver/service/storageservice"
	"github.com/deadblue/dlna115/internal/soap"
	"github.com/deadblue/dlna115/internal/util"
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
		util.RenderXML(rw, envelope)
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
	items := s.ss.Browse(req.ObjectID)
	for _, item := range items {
		switch item := item.(type) {
		case *storageservice.Dir:
			// Make container object
			obj := (&didl.StorageFolderContainer{}).Init()
			obj.ID = item.ID
			obj.ParentID = req.ObjectID
			obj.StorageUsed = -1
			obj.Title = item.Name
			result.AppendContainer(obj)
			resp.TotalMatches += 1
		case *storageservice.VideoFile:
			// Make videoItem object
			obj := (&didl.VideoItem{}).Init()
			obj.ID = item.ID
			obj.ParentID = req.ObjectID
			obj.Title = item.Name
			obj.Res.ProtocolInfo = "http-get:*:video/mp4:*"
			obj.Res.Size = item.Size
			obj.Res.NrAudioChannels = item.AudioChannels
			obj.Res.SampleFrequency = item.AudioSampleRate
			obj.Res.Resolution = item.VideoResolution
			obj.Res.URL = item.PlayURL
			// Calculate bitrate
			obj.Res.Bitrate = int(float64(item.Size) / item.Duration)
			// Format Duration
			obj.Res.Duration = formatDuration(item.Duration)
			resp.TotalMatches += 1
		}
	}
	resp.NumberReturned = resp.TotalMatches
	resp.SetResult(result)
	ret = resp
	return
}

func formatDuration(seconds float64) string {
	m := int(seconds / 60)
	s := seconds - float64(m*60)
	h := int(m / 60)
	m = m % 60
	return fmt.Sprintf("%d:%02d:%06.3f", h, m, s)
}
