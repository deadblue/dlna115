package connectionmanager

import "net/http"

func (s *Service) HandleControl(rw http.ResponseWriter, req *http.Request) {
	// Dummy impl
	rw.WriteHeader(http.StatusOK)
}
