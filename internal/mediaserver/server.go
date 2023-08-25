package mediaserver

import (
	"fmt"
	"net/http"

	"github.com/deadblue/dlna115/internal/mediaserver/service/connectionmanager"
	"github.com/deadblue/dlna115/internal/mediaserver/service/contentdirectory"
	"github.com/deadblue/dlna115/internal/mediaserver/service/storage115"
	"github.com/deadblue/dlna115/internal/upnp"
)

type Server struct {
	// Closed flag
	cf int32
	// Error channel
	ec chan error

	// Server port
	sp uint
	// HTTP server
	hs *http.Server

	// Unique device name
	udn string
	// UPnP services
	uss []upnp.Service

	// Description XML content
	desc []byte
}

func New(opts *Options) *Server {
	// Instantiate services
	ss := storage115.New()
	cds := contentdirectory.New(ss)
	cms := connectionmanager.New()

	// Make server
	s := &Server{
		cf:  0,
		ec:  make(chan error, 1),
		sp:  opts.Port,
		hs:  &http.Server{},
		udn: fmt.Sprintf("uuid:%s", opts.UUID),
		uss: []upnp.Service{cds, cms},
	}
	s.initDesc(opts.Name)

	// Create HTTP handler
	mux := http.NewServeMux()
	// Device description URL
	mux.HandleFunc(descUrl, s.handleDescXml)
	// Register service URLs
	ss.RegisterTo(mux)
	cds.RegisterTo(mux)
	cms.RegisterTo(mux)
	// Set handler to HTTP server
	s.hs.Handler = mux
	return s
}
