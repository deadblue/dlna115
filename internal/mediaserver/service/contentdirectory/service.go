package contentdirectory

import (
	"fmt"
	"net/http"

	"github.com/deadblue/elevengo"
	"github.com/google/uuid"
)

type FileForwarder interface {
	GetAccessURL(accessCode string) string
}

type Service struct {
	ea *elevengo.Agent
	ff FileForwarder
}

func New(ea *elevengo.Agent, ff FileForwarder) *Service {
	return &Service{
		ea: ea,
		ff: ff,
	}
}

func (s *Service) HandleEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "SUBSCRIBE" {
		subId := fmt.Sprintf("uuid:%s", uuid.NewString())
		req.Header.Set("SID", subId)
	}
	rw.WriteHeader(http.StatusOK)
}
