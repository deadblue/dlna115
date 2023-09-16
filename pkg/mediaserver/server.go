package mediaserver

import (
	"fmt"
	"net/http"

	"github.com/deadblue/dlna115/pkg/mediaserver/service/connectionmanager"
	"github.com/deadblue/dlna115/pkg/mediaserver/service/contentdirectory"
	"github.com/deadblue/dlna115/pkg/storage"
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

func New(opts *Options, ss storage.StorageService) *Server {
	// Instantiate services
	cds := contentdirectory.New(ss)
	cms := connectionmanager.New()

	// Create HTTP handler
	mux := http.NewServeMux()
	// Register service handlers
	cds.MountTo(mux)
	cms.MountTo(mux)

	// Make server
	s := &Server{
		cf:  0,
		ec:  make(chan error, 1),
		sp:  util.DefaultNumber(opts.Port, 8115),
		hs:  &http.Server{Handler: mux},
		uss: []upnp.Service{cds, cms},
	}
	s.udn = fmt.Sprintf(
		"uuid:%s",
		util.DefaultStringFunc(opts.UUID, uuid.NewString),
	)
	// Device description URL
	s.initDesc(util.DefaultString(opts.Name, "115"))
	mux.HandleFunc(descUrl, s.handleDescXml)
	return s
}
