package ssdp

import (
	"errors"
	"log"
	"net"
)

func (s *Server) Startup() (err error) {
	s.conn, err = net.ListenMulticastUDP("udp4", nil, serverAddr)
	if err == nil {
		go s.loop()
	}
	return
}

func (s *Server) Shutdown() {
	if s.conn != nil {
		_ = s.conn.Close()
	}
}

func (s *Server) Done() <-chan struct{} {
	return s.done
}

func (s *Server) loop() {
	buf := make([]byte, 1500)
	for {
		if size, raddr, err := s.conn.ReadFromUDP(buf); err != nil {
			if !errors.Is(err, net.ErrClosed) {
				log.Printf("Read error: %s", err.Error())
			}
			break
		} else {
			req := &Request{}
			if err = req.Unmarshal(buf[:size]); err == nil {
				go s.handleRequest(raddr, req)
			} else {
				log.Printf("Parse SSDP message failed: %s\nRaw: %s", err, buf[:size])
			}
		}
	}
	close(s.done)
}
