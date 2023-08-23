package storage115

import (
	"log"
	"net/http"

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
	var err error
	fileCode := req.URL.Path[videoUrlLen:]
	ticket := &elevengo.VideoTicket{}
	if err = s.ea.VideoCreateTicket(fileCode, ticket); err != nil {
		// TODO: Send error to client
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	http.Redirect(rw, req, ticket.Url, http.StatusTemporaryRedirect)
}
