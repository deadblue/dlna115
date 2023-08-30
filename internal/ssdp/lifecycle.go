package ssdp

import (
	"errors"
	"log"
	"net"

	"github.com/deadblue/dlna115/internal/util"
	"golang.org/x/net/ipv4"
)

func (s *Server) Startup() (err error) {
	// Create packet connection
	npc, err := net.ListenPacket("udp4", ":1900")
	if err != nil {
		return
	}
	s.pc = ipv4.NewPacketConn(npc)
	// Receive multicast on all available interfaces
	if nifs, nerr := net.Interfaces(); nerr == nil {
		for _, nif := range nifs {
			// Skip loopback
			if nif.Flags&net.FlagLoopback != 0 {
				continue
			}
			if nif.Flags&net.FlagUp == 0 {
				continue
			}
			if nif.Flags&net.FlagRunning == 0 {
				continue
			}
			if !util.InterfaceHasIPv4(&nif) {
				continue
			}
			if jerr := s.pc.JoinGroup(&nif, multicastAddr); jerr == nil {
				// Save joined group, we need leave from them when shutdown.
				s.jnis = append(s.jnis, nif)
			}
		}
	}
	go s.loop()
	return
}

func (s *Server) Shutdown() {
	if s.pc != nil {
		// Leave joined groups
		for _, jni := range s.jnis {
			s.pc.LeaveGroup(&jni, multicastAddr)
		}
		s.pc.Close()
	}
}

func (s *Server) Done() <-chan struct{} {
	return s.done
}

func (s *Server) loop() {
	buf := make([]byte, 1500)
	for {
		size, _, src, err := s.pc.ReadFrom(buf)
		if err != nil {
			if !errors.Is(err, net.ErrClosed) {
				log.Printf("Read error: %s", err.Error())
			}
			break
		}
		raddr := src.(*net.UDPAddr)
		req := &Request{}
		if err = req.Unmarshal(buf[:size]); err == nil {
			go s.handleRequest(raddr, req)
		} else {
			log.Printf("Parse SSDP message failed: %s\nRaw: %s", err, buf[:size])
		}
	}
	close(s.done)
	s.pc = nil
}
