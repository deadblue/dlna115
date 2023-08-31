package contentdirectory

import (
	"fmt"
	"net/http"

	"github.com/deadblue/dlna115/pkg/upnp"
	"github.com/google/uuid"
)

func (s *Service) HandleEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method == upnp.MethodSubscribe {
		subId := fmt.Sprintf("uuid:%s", uuid.NewString())
		rw.Header().Set("SID", subId)
	}
	rw.WriteHeader(http.StatusOK)
}
