package ssdp

import (
	"log"
	"net"
	"net/http"

	"github.com/deadblue/dlna115/internal/upnp"
)

func (s *Server) handle(req *Request, raddr *net.UDPAddr) {
	switch req.Method {
	case methodSearch:
		s.handleSearch(req, raddr)
	}
}

func (s *Server) handleSearch(req *Request, raddr *net.UDPAddr) {
	target := req.GetHeader(headerSearchTarget)
	if target != searchAll && target != upnp.DeviceTypeMediaServer1 {
		return
	}

	// Make response
	resp := &Response{
		StatusCode: http.StatusOK,
	}
	resp.SetHeader(headerServer, upnp.ServerName)
	resp.SetHeader(headerCacheControl, "max-age=3600")
	resp.SetHeader(headerExtension, "")
	resp.SetHeader(headerSearchTarget, upnp.DeviceTypeMediaServer1)

	// Prepare UDP
	conn, err := net.DialUDP(raddr.Network(), nil, raddr)
	if err != nil {
		log.Printf("Can not send packet to %s, error: %s", raddr.String(), err)
		return
	}
	resp.WriteTo(conn)
}
