package contentdirectory

import "net/http"

const (
	descUrl    = "/ContentDirectory/scpd.xml"
	controlUrl = "/ContentDirectory/control"
	eventUrl   = "/ContentDirectory/event"
)

func (s *Service) RegisterTo(mux *http.ServeMux) {
	mux.HandleFunc(descUrl, s.HandleDescXml)
	mux.HandleFunc(controlUrl, s.HandleControl)
	mux.HandleFunc(eventUrl, s.HandleEvent)
}
