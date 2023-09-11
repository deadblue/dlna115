package impl

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/deadblue/elevengo"
)

const (
	PlayURL = "/storage/video/"

	PlayTypeHls  = "hls"
	PlayTypeFile = "file"

	playUrlLen = len(PlayURL)
)

var (
	rePlayPath = regexp.MustCompile(`^(\w+)-(\w+)/(\w+)\.(\w+)$`)

	errInvalidPath = errors.New("invalid request path")

	errUnsupportedType = errors.New("unsupported play type")
)

type PlayRequest struct {
	// Indicate whether it is a HEAD request
	IsHead bool

	// Play type
	Type string
	// Video extension
	VideoExt string

	// Video pick code
	PickCode string
	// Request extension
	RequestExt string

	// Range start for "file" type
	RangeStart int64
}

func (r *PlayRequest) Parse(req *http.Request) (err error) {
	// Check request method
	r.IsHead = req.Method == http.MethodHead
	// Extract parameters from path
	relPath := req.URL.Path[playUrlLen:]
	match := rePlayPath.FindAllStringSubmatch(relPath, 1)
	if len(match) == 0 || len(match[0]) != 5 {
		return errInvalidPath
	}
	r.Type = match[0][1]
	r.VideoExt = match[0][2]
	r.PickCode = match[0][3]
	r.RequestExt = match[0][4]
	// Parse range start
	if r.Type == PlayTypeFile {
		if reqRange := req.Header.Get("Range"); reqRange != "" {
			index := strings.IndexRune(reqRange, '-')
			r.RangeStart, _ = strconv.ParseInt(reqRange[6:index], 10, 64)
		}
	}
	return
}

func generatePlayUrl(file *elevengo.File, disableHLS bool) string {
	// Determine play type
	playType := PlayTypeHls
	if disableHLS {
		playType = PlayTypeFile
	}
	extName := filepath.Ext(file.Name)
	// Build play url
	return fmt.Sprintf(
		"%s%s-%s/%s%s",
		PlayURL,
		playType, extName[1:],
		file.PickCode, extName,
	)
}

func (s *Service) MountTo(mux *http.ServeMux) {
	mux.HandleFunc(PlayURL, s.HandlePlay)
}

func (s *Service) HandlePlay(rw http.ResponseWriter, req *http.Request) {
	var err error
	pr := &PlayRequest{}
	if err = pr.Parse(req); err != nil {
		log.Printf("Parse play request failed: %s", err)
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	// Check request ext with video ext
	if pr.RequestExt != pr.VideoExt {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	switch pr.Type {
	case PlayTypeHls:
		err = s.handlePlayHls(rw, pr)
	case PlayTypeFile:
		err = s.handlePlayFile(rw, pr)
	default:
		err = errUnsupportedType
	}
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
	}
}

func (s *Service) handlePlayHls(rw http.ResponseWriter, req *PlayRequest) (err error) {
	ticket, ok := s.vtc.Get(req.PickCode)
	if !ok {
		ticket = &elevengo.VideoTicket{}
		if err = s.ea.VideoCreateTicket(req.PickCode, ticket); err == nil {
			s.vtc.Put(req.PickCode, ticket)
		}
	}
	if err != nil {
		return
	}

	// Fetch HLS content from upstream
	body, _ := s.ea.Fetch(ticket.Url)
	defer body.Close()

	// Send to client
	rw.Header().Set("Content-Type", "application/x-mpegURL")
	rw.WriteHeader(http.StatusOK)
	io.Copy(rw, body)

	return
}

func (s *Service) handlePlayFile(rw http.ResponseWriter, req *PlayRequest) (err error) {
	ticket, ok := s.dtc.Get(req.PickCode)
	if !ok {
		ticket = &elevengo.DownloadTicket{}
		if err = s.ea.DownloadCreateTicket(req.PickCode, ticket); err == nil {
			s.dtc.Put(req.PickCode, ticket)
		}
	}
	if err != nil {
		return
	}

	var body io.ReadCloser
	if req.RangeStart == 0 {
		body, err = s.ea.Fetch(ticket.Url)
	} else {
		body, err = s.ea.FetchRange(ticket.Url, elevengo.RangeMiddle(req.RangeStart, -1))
	}
	if err != nil {
		log.Printf("Fetch video from remote failed: %s", err)
		rw.WriteHeader(http.StatusBadGateway)
		return
	}
	defer body.Close()

	rw.Header().Set("Content-Type", mime.TypeByExtension(req.VideoExt))
	if req.RangeStart == 0 {
		rw.Header().Set("Content-Length", strconv.FormatInt(ticket.FileSize, 10))
		rw.WriteHeader(http.StatusOK)
	} else {
		contentSize := ticket.FileSize - req.RangeStart
		rw.Header().Set("Content-Length", strconv.FormatInt(contentSize, 10))
		contentRange := fmt.Sprintf(
			"bytes %d-%d/%d",
			req.RangeStart, ticket.FileSize-1, ticket.FileSize,
		)
		rw.Header().Set("Content-Range", contentRange)
		rw.WriteHeader(http.StatusPartialContent)
	}
	// Send video data
	io.Copy(rw, body)

	return
}
