package connectionmanager

import "net/http"

const (
	descUrl    = "/ConnectionManager/scpd.xml"
	controlUrl = "/ConnectionManager/control"
	eventUrl   = "/ConnectionManager/event"
)

func (s *Service) RegisterTo(mux *http.ServeMux) {
	mux.HandleFunc(descUrl, s.HandleDescXml)
	mux.HandleFunc(controlUrl, s.HandleControl)
	mux.HandleFunc(eventUrl, s.HandleEvent)
}
