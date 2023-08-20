package media

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func (s *Server) handleDescDeviceXml(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/xml")
	rw.Header().Set("Content-Length", strconv.Itoa(len(s.descXml)))
	rw.Header().Set("Server", serverTag)
	rw.WriteHeader(200)
	rw.Write(s.descXml)
}

func (s *Server) handleEvent(rw http.ResponseWriter, req *http.Request) {
	if req.Method == "SUBSCRIBE" {
		subId := fmt.Sprintf("uuid:%s", uuid.NewString())
		req.Header.Set("SID", subId)
	}
	rw.Header().Set("Server", serverTag)
	rw.WriteHeader(200)
}

func (s *Server) handleConnectionManagerControl(rw http.ResponseWriter, req *http.Request) {
	// Dummy response
	rw.Header().Set("Server", serverTag)
	rw.WriteHeader(200)
}

func (s *Server) handleContentDirectoryControl(rw http.ResponseWriter, req *http.Request) {

}
