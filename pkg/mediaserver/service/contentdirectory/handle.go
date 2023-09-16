package contentdirectory

import "net/http"

const (
	_DescUrl    = "/ContentDirectory/scpd.xml"
	_ControlUrl = "/ContentDirectory/control"
	_EventUrl   = "/ContentDirectory/event"

	// View URL for viewing file content
	_ViewUrl = "/ContentDirectory/view/"

	_ViewUrlLen = len(_ViewUrl)
)

func (s *Service) MountTo(mux *http.ServeMux) {
	mux.HandleFunc(_DescUrl, s.HandleDescXml)
	mux.HandleFunc(_ControlUrl, s.HandleControl)
	mux.HandleFunc(_EventUrl, s.HandleEvent)
	mux.HandleFunc(_ViewUrl, s.HandleView)
}
