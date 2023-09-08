package impl

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/deadblue/elevengo"
)

const (
	VideoURL = "/storage/video/"

	PlayTypeStream = "stream"
	PlayTypeFile   = "file"

	videoUrlLen = len(VideoURL)
)

func (s *Service) MountTo(mux *http.ServeMux) {
	mux.HandleFunc(VideoURL, s.HandleVideo)
}

func (s *Service) HandleVideo(rw http.ResponseWriter, req *http.Request) {
	// Parse request path
	relPath := req.URL.Path[videoUrlLen:]
	if len(relPath) == 0 {
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	parts := strings.SplitN(relPath, "/", 2)
	if len(parts) < 2 {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	playType, pickCode := parts[0], parts[1]

	// check pickcode
	if strings.IndexRune(pickCode, '.') > 0 {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	var err error
	switch playType {
	case PlayTypeStream:
		ticket := &elevengo.VideoTicket{}
		if err = s.ea.VideoCreateTicket(pickCode, ticket); err == nil {
			s.sendVideoStream(rw, req, ticket)
		}
	case PlayTypeFile:
		ticket, ok := s.dtc.Get(pickCode)
		if !ok {
			ticket = &elevengo.DownloadTicket{}
			if err = s.ea.DownloadCreateTicket(pickCode, ticket); err == nil {
				s.dtc.Put(pickCode, ticket)
			}
		}
		if ticket != nil {
			s.sendVideoFile(rw, req, ticket)
		}
	default:
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (s *Service) sendVideoStream(rw http.ResponseWriter, req *http.Request, ticket *elevengo.VideoTicket) {
	body, _ := s.ea.Fetch(ticket.Url)
	defer body.Close()

	rw.Header().Set("Content-Type", "application/x-mpegURL")
	rw.WriteHeader(http.StatusOK)
	io.Copy(rw, body)
}

func (s *Service) sendVideoFile(rw http.ResponseWriter, req *http.Request, ticket *elevengo.DownloadTicket) {
	var start int64 = 0
	if reqRange := req.Header.Get("Range"); reqRange != "" {
		index := strings.IndexRune(reqRange, '-')
		start, _ = strconv.ParseInt(reqRange[6:index], 10, 64)
	}

	var body io.ReadCloser
	var err error
	if start == 0 {
		body, err = s.ea.Fetch(ticket.Url)
	} else {
		body, err = s.ea.FetchRange(ticket.Url, elevengo.RangeMiddle(start, -1))
	}
	if err != nil {
		log.Printf("Fetch video from remote failed: %s", err)
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	defer body.Close()

	ext := filepath.Ext(ticket.FileName)
	rw.Header().Set("Content-Type", mime.TypeByExtension(ext))
	if start == 0 {
		rw.Header().Set("Content-Length", strconv.FormatInt(ticket.FileSize, 10))
		rw.WriteHeader(http.StatusOK)
	} else {
		contentSize := ticket.FileSize - start
		rw.Header().Set("Content-Length", strconv.FormatInt(contentSize, 10))
		contentRange := fmt.Sprintf(
			"bytes %d-%d/%d",
			start, ticket.FileSize-1, ticket.FileSize,
		)
		rw.Header().Set("Content-Range", contentRange)
		rw.WriteHeader(http.StatusPartialContent)
	}
	// Send video data
	io.Copy(rw, body)
}
