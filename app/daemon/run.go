package daemon

import (
	"log"
	"os"
	"os/signal"

	"github.com/deadblue/dlna115/pkg/mediaserver"
	"github.com/deadblue/dlna115/pkg/ssdp"
	"github.com/deadblue/dlna115/pkg/storage/impl"
	"github.com/deadblue/dlna115/pkg/version"
)

func (c *Command) Run() (err error) {
	// Print full version at beginning
	log.Println(version.Full())

	// Load and parse config file
	options := &Options{}
	if err = options.Load(c.ConfigFile); err != nil {
		log.Fatalf("Load config file failed: %s", err)
	}

	// Initialize storage service
	storage := impl.New(&options.Storage)
	if err = storage.ApplyOptions(); err != nil {
		log.Fatalf("Initialize storage service failed: %s", err)
	}

	// Handle OS signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	defer func() {
		signal.Stop(sigChan)
		close(sigChan)
	}()

	// Create & start media server
	ms := mediaserver.New(&options.Media, storage)
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
	return
}
