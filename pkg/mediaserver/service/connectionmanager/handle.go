package connectionmanager

import "net/http"

const (
	_DescUrl    = "/ConnectionManager/scpd.xml"
	_ControlUrl = "/ConnectionManager/control"
	_EventUrl   = "/ConnectionManager/event"
)

func (s *Service) MountTo(mux *http.ServeMux) {
	mux.HandleFunc(_DescUrl, s.HandleDescXml)
	mux.HandleFunc(_ControlUrl, s.HandleControl)
	mux.HandleFunc(_EventUrl, s.HandleEvent)
}
