package mediaserver

import (
	"fmt"
	"net/http"

	"github.com/deadblue/dlna115/pkg/mediaserver/service/connectionmanager"
	"github.com/deadblue/dlna115/pkg/mediaserver/service/contentdirectory"
	"github.com/deadblue/dlna115/pkg/mediaserver/service/storage115"
	"github.com/deadblue/dlna115/pkg/upnp"
	"github.com/deadblue/dlna115/pkg/util"
	"github.com/google/uuid"
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

func New(opts *Options, sopts *storage115.Options) *Server {
	// Instantiate services
	ss := storage115.New(sopts)
	cds := contentdirectory.New(ss)
	cms := connectionmanager.New()

	// Make server
	s := &Server{
		cf:  0,
		ec:  make(chan error, 1),
		sp:  util.DefaultNumber(opts.Port, 8115),
		hs:  &http.Server{},
		uss: []upnp.Service{cds, cms},
	}
	s.udn = fmt.Sprintf(
		"uuid:%s",
		util.DefaultStringFunc(opts.UUID, uuid.NewString),
	)
	s.initDesc(util.DefaultString(opts.Name, "115"))

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
