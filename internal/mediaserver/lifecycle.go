package mediaserver

import (
	"context"
	"errors"
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
	l, err := net.Listen("tcp", ":5000")
	if err != nil {
		return err
	}
	log.Printf(
		"Starting server at: %s://%s",
		l.Addr().Network(),
		l.Addr().String(),
	)
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
