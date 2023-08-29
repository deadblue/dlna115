package ssdp

import (
	"net"

	"github.com/deadblue/dlna115/internal/upnp"
	"golang.org/x/net/ipv4"
)

// Server is a simple SSDP server implementation.
type Server struct {
	// Joined net interfaces
	jnis []net.Interface
	// Packet connection
	pc *ipv4.PacketConn

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
