package ssdp

import (
	"log"
	"net"
	"net/http"

	"github.com/deadblue/dlna115/internal/upnp"
)

func (s *Server) handleRequest(raddr *net.UDPAddr, req *Request) {
	switch req.Method {
	case methodSearch:
		s.handleSearch(raddr, req)
	}
}

func (s *Server) handleSearch(raddr *net.UDPAddr, req *Request) {
	// Directly return if there is no device to respond
	if s.device == nil {
		return
	}
	// Check search target
	target := req.GetHeader(headerSearchTarget)
	if target != searchAll && target != s.device.DeviceType() {
		return
	}

	// Create UDP connection to remote
	conn, err := net.DialUDP(raddr.Network(), nil, raddr)
	if err != nil {
		log.Printf("Can not send packet to %s, error: %s", raddr.String(), err)
		return
	}
	defer conn.Close()
	// Get local IP
	laddr := conn.LocalAddr().(*net.UDPAddr)
	localIP := laddr.IP.String()

	// Make response
	resp := &Response{
		StatusCode: http.StatusOK,
	}
	resp.SetHeader(headerCacheControl, "max-age=3600")
	resp.SetHeader(headerServer, upnp.ServerName)
	resp.SetHeader(headerExtension, "")
	resp.SetHeader(headerSearchTarget, s.device.DeviceType())
	resp.SetHeader(headerUniqueServiceName, s.device.DeviceUSN())
	resp.SetHeader(headerLocation, s.device.GetDeviceDescURL(localIP))
	// Send response
	resp.WriteTo(conn)
}
