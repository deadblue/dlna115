package mediaserver

import (
	"net/http"

	"github.com/deadblue/dlna115/internal/mediaserver/service/connectionmanager"
	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory"
	"github.com/deadblue/dlna115/internal/mediaserver/service/storage115"
)

type Server struct {
	// Closed flag
	cf int32
	// Error channel
	ec chan error
	// Core HTTP server
	hs *http.Server
	// Services
	ss  *storage115.Service
	cds *contentdirectory.Service
	cms *connectionmanager.Service

	descXml []byte
}

func New(uuid string) *Server {
	// Create storage service
	ss := storage115.New()
	// Make server
	s := &Server{
		cf: 0,
		ec: make(chan error, 1),
		hs: &http.Server{},
		// Services
		ss:  ss,
		cds: contentdirectory.New(ss),
		cms: connectionmanager.New(),
		// Make description xml
		descXml: makeDeviceDesc(uuid),
	}
	// Register handle functions
	mux := http.NewServeMux()
	// Register storage service URLs
	ss.RegisterTo(mux)
	// Device description URL
	mux.HandleFunc(deviceDescUrl, s.handleDescDeviceXml)
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
