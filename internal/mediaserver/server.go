package mediaserver

import (
	"net/http"

	"github.com/deadblue/dlna115/internal/mediaserver/service/connectionmanager"
	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory"
	"github.com/deadblue/dlna115/internal/mediaserver/service/forward"
	"github.com/deadblue/elevengo"
)

type Server struct {
	// Core HTTP server
	hs *http.Server
	// Services
	fs  *forward.Service
	cds *contentdirectory.Service
	cms *connectionmanager.Service

	descXml []byte
}

func (s *Server) Run() error {
	return s.hs.ListenAndServe()
}

func New(uuid string) *Server {
	ea := elevengo.Default()
	fs := forward.New(ea)
	s := &Server{
		// Initialize HTTP server
		hs: &http.Server{
			Addr: ":5000",
		},
		// Services
		fs:  fs,
		cds: contentdirectory.New(ea, fs),
		cms: connectionmanager.New(),
		// Make description xml
		descXml: makeDeviceDesc(uuid),
	}
	// Register handle functions
	mux := http.NewServeMux()
	// Device description URL
	mux.HandleFunc(deviceDescUrl, s.handleDescDeviceXml)
	// Forwarder URLs
	mux.HandleFunc(forward.HandleURL, s.fs.HandleVideo)
	// ConnectionManager service URLs
	mux.HandleFunc(connectionmanager.DescUrl, s.cms.HandleDescXml)
	mux.HandleFunc(connectionmanager.ControlUrl, s.cms.HandleControl)
	mux.HandleFunc(connectionmanager.EventUrl, s.cms.HandleEvent)
	// ContentDirectory service URLs
	mux.HandleFunc(contentdirectory.DescUrl, s.cds.HandleDescXml)
	mux.HandleFunc(contentdirectory.ControlUrl, s.cds.HandleControl)
	mux.HandleFunc(contentdirectory.EventUrl, s.cds.HandleEvent)
	s.hs.Handler = mux
	return s
}
