package media

import (
	"net/http"

	"github.com/deadblue/elevengo"
)

type Server struct {
	// Core HTTP server
	hs *http.Server
	// 115 agent
	ea *elevengo.Agent

	descXml []byte
}

func (s *Server) Run() error {
	return s.hs.ListenAndServe()
}

func New(uuid string) *Server {
	s := &Server{
		// Initialize HTTP server
		hs: &http.Server{
			Addr: ":5000",
		},
		// Init 115 agent
		ea: elevengo.Default(),
		// Make description xml
		descXml: makeDeviceDesc(uuid),
	}
	mux := http.NewServeMux()
	mux.HandleFunc(deviceDescUrl, s.handleDescDeviceXml)
	// mux.HandleFunc(connectionManagerDescUrl, s.HandleDescDeviceXml)
	mux.HandleFunc(connectionManagerControlUrl, s.handleConnectionManagerControl)
	mux.HandleFunc(connectionManagerEventUrl, s.handleEvent)
	// mux.HandleFunc(contentDirectoryDescUrl, s.HandleDescDeviceXml)
	mux.HandleFunc(contentDirectoryControlUrl, s.handleContentDirectoryControl)
	mux.HandleFunc(contentDirectoryEventUrl, s.handleEvent)
	s.hs.Handler = mux
	return s
}
