package ssdp

import (
	"errors"
	"log"
	"net"
)

// Server is a simple SSDP server implementation.
type Server struct {
	// UDP connection
	conn *net.UDPConn
	// Error channel
	errCh chan error
}

func (s *Server) Startup() (err error) {
	// Resolve never fails!
	addr, _ := net.ResolveUDPAddr("udp4", "239.255.255.250:1900")
	if s.conn, err = net.ListenMulticastUDP("udp", nil, addr); err == nil {
		go s.loop()
	}
	return
}

func (s *Server) Shutdown() {
	if s.conn != nil {
		_ = s.conn.Close()
	}
	close(s.errCh)
}

func (s *Server) Error() <-chan error {
	return s.errCh
}

func (s *Server) loop() {
	log.Println("Start SSDP service...")
	buf := make([]byte, 1500)
	for {
		size, addr, err := s.conn.ReadFrom(buf)
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				log.Printf("Read error: %s", err.Error())
			}
			break
		}
		log.Printf("Receive %d bytes from peer [%s]: %s",
			size, addr.String(), buf[:size],
		)
	}
	log.Printf("Server shutdown!")
}

func NewServer() *Server {
	server := &Server{
		errCh: make(chan error, 1),
	}
	return server
}
