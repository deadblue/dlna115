package ssdp

import (
	"net"

	"github.com/deadblue/dlna115/internal/upnp"
)

// Server is a simple SSDP server implementation.
type Server struct {
	// UDP connection
	conn *net.UDPConn
	// Done channel
	done chan struct{}
	// UPnP device
	device upnp.Device
}

func NewServer(device upnp.Device) *Server {
	server := &Server{
		done:   make(chan struct{}, 1),
		device: device,
	}
	return server
}
