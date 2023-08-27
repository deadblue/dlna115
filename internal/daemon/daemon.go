package daemon

import (
	"log"
	"os"
	"os/signal"

	"github.com/deadblue/dlna115/internal/mediaserver"
	"github.com/deadblue/dlna115/internal/ssdp"
)

// Run starts daemon process.
func Run() (err error) {
	args := (&Args{})
	if err = args.Init(); err != nil {
		log.Fatal(err)
	}

	// Handle OS signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	defer func() {
		signal.Stop(sigChan)
		close(sigChan)
	}()

	// Create media service
	ms := mediaserver.New(&args.Media)
	// Start media service
	if err = ms.Startup(); err != nil {
		log.Fatal(err)
	}
	ssdp.NotifyDeviceAvailable(ms)

	ss := ssdp.NewServer(ms)
	_ = ss.Startup()

	// Loop
	for running := true; running; {
		select {
		case <-sigChan:
			log.Printf("Shutdown DLNA media server ...")
			ssdp.NotifyDeviceUnavailable(ms)
			ms.Shutdown()
		case err = <-ms.ErrChan():
			if err != nil {
				log.Printf("Media server closed with error: %s", err)
			} else {
				log.Println("Media server closed normally!")
			}
			// Shutdown SSDP server
			ss.Shutdown()
		case <-ss.Done():
			running = false
		}
	}
	return
}
