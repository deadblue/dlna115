package connectionmanager

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) HandleControl(rw http.ResponseWriter, req *http.Request) {
	// Dummy impl
	rw.WriteHeader(http.StatusOK)
}

func (s *Service) HandleEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "SUBSCRIBE" {
		subId := fmt.Sprintf("uuid:%s", uuid.NewString())
		req.Header.Set("SID", subId)
	}
	rw.WriteHeader(http.StatusOK)
}
