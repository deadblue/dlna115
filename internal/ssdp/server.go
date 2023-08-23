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
	// Done channel
	done chan struct{}
}

func (s *Server) Startup() (err error) {
	if s.conn, err = net.ListenMulticastUDP("udp4", nil, serverAddr); err != nil {
		return
	}
	go s.loop()
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
	log.Println("Start SSDP service...")
	buf := make([]byte, 1500)
	for {
		size, addr, err := s.conn.ReadFromUDP(buf)
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				log.Printf("Read error: %s", err.Error())
			}
			break
		}
		log.Printf("Receive %d bytes from peer [%s]:\n%s",
			size, addr.String(), string(buf[:size]),
		)
	}
	log.Printf("Server shutdown!")
	close(s.done)
}

func NewServer() *Server {
	server := &Server{
		done: make(chan struct{}, 1),
	}
	return server
}
