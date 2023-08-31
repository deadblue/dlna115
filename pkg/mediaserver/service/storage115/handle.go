package storage115

import (
	"net/http"
	"strings"

	"github.com/deadblue/elevengo"
)

const (
	SetupUrl = "/storage/setup"
	VideoURL = "/storage/video/"

	videoUrlLen = len(VideoURL)
)

func (s *Service) RegisterTo(mux *http.ServeMux) {
	mux.HandleFunc(VideoURL, s.HandleVideo)
}

func (s *Service) HandleVideo(rw http.ResponseWriter, req *http.Request) {
	fileName := req.URL.Path[videoUrlLen:]
	if len(fileName) == 0 {
		rw.WriteHeader(http.StatusForbidden)
		return
	}
	// Get pickcode & extension name
	dotIndex := strings.IndexRune(fileName, '.')
	if dotIndex < 0 {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	pickCode, extName := fileName[:dotIndex], fileName[dotIndex+1:]
	if extName != "m3u8" {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	ticket := &elevengo.VideoTicket{}
	if err := s.ea.VideoCreateTicket(pickCode, ticket); err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	http.Redirect(rw, req, ticket.Url, http.StatusTemporaryRedirect)
}
