package storage115

import (
	"log"
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
	mux.HandleFunc(SetupUrl, s.HandleSetup)
	mux.HandleFunc(VideoURL, s.HandleVideo)
}

// HandleSetup
func (s *Service) HandleSetup(rw http.ResponseWriter, req *http.Request) {
	cred := &elevengo.Credential{
		UID:  req.URL.Query().Get("uid"),
		CID:  req.URL.Query().Get("cid"),
		SEID: req.URL.Query().Get("seid"),
	}
	if err := s.ea.CredentialImport(cred); err != nil {
		// TODO: Send error to client
		log.Printf("Import credentail failed: %s", err)
		rw.WriteHeader(http.StatusServiceUnavailable)
	} else {
		rw.WriteHeader(http.StatusAccepted)
	}
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
