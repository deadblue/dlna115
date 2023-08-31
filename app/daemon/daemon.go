package daemon

import (
	"log"
	"os"
	"os/signal"

	"github.com/deadblue/dlna115/pkg/mediaserver"
	"github.com/deadblue/dlna115/pkg/ssdp"
)

// Run starts daemon process.
func Run() (err error) {
	opts := &Options{}
	if err = opts.Init(); err != nil {
		log.Fatal(err)
	}

	// Handle OS signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	defer func() {
		signal.Stop(sigChan)
		close(sigChan)
	}()

	// Create & start media server
	ms := mediaserver.New(&opts.Media, &opts.Storage)
	// Start media service
	if err = ms.Startup(); err != nil {
		log.Fatal(err)
	}
	ssdp.NotifyDeviceAvailable(ms)

	// Create & start SSDP server
	ss := ssdp.NewServer(ms)
	_ = ss.Startup()

	// Wait OS signal
	<-sigChan
	log.Printf("Shutdown server ...")

	// Shutdown SSDP server
	ss.Shutdown()
	// Shutdown media server
	ssdp.NotifyDeviceUnavailable(ms)
	ms.Shutdown()

	// Wait SSDP server shutdown
	<-ss.Done()
	// Wait media server shutdown
	<-ms.ErrChan()

	log.Printf("Byebye")

	// Loop
	// for running := true; running; {
	// 	select {
	// 	case <-sigChan:
	// 		log.Printf("Shutdown DLNA media server ...")
	// 		ssdp.NotifyDeviceUnavailable(ms)
	// 		ms.Shutdown()
	// 	case err = <-ms.ErrChan():
	// 		if err != nil {
	// 			log.Printf("Media server closed with error: %s", err)
	// 		} else {
	// 			log.Println("Media server closed normally!")
	// 		}
	// 		// Shutdown SSDP server
	// 		ss.Shutdown()
	// 	case <-ss.Done():
	// 		running = false
	// 	}
	// }
	return
}
