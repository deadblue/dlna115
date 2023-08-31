package mediaserver

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"sync/atomic"
)

// die marks server as closed, and notifies observers thru error channel.
func (s *Server) die(err error) {
	// Closed flag will be set to 1 when first error arrive here.
	if atomic.CompareAndSwapInt32(&s.cf, 0, 1) {
		if err != nil {
			s.ec <- err
		}
		close(s.ec)
	}
}

func (s *Server) serve(l net.Listener) {
	err := s.hs.Serve(l)
	if !errors.Is(err, http.ErrServerClosed) {
		// When error is not ErrServerClosed, send to error channel
		s.die(err)
	}
}

// Startup starts HTTP server in goroutine.
func (s *Server) Startup() (err error) {
	addr := fmt.Sprintf(":%d", s.sp)
	l, err := net.Listen("tcp4", addr)
	// TODO: Select a random port when error
	if err != nil {
		return err
	}
	log.Printf("Running DLNA media server at http://%s", l.Addr().String())
	go s.serve(l)
	return nil
}

// Shutdown tries to shutdown HTTP server.
func (s *Server) Shutdown() {
	go func() {
		if atomic.LoadInt32(&s.cf) == 0 {
			// Call shutdown and wait for its result
			err := s.hs.Shutdown(context.Background())
			s.die(err)
		}
	}()
}

// ErrChan returns a readonly error channel.
func (s *Server) ErrChan() <-chan error {
	return s.ec
}
