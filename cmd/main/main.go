package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/deadblue/dlna115/internal/mediaserver"
	"github.com/google/uuid"
)

func main() {
	// Handle OS signal
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	defer func() {
		signal.Stop(sigChan)
		close(sigChan)
	}()

	// Create media service
	ms := mediaserver.New(uuid.NewString())
	// Start media service
	var err error
	if err = ms.Startup(); err != nil {
		log.Fatal(err)
	}

	// Loop
	for running := true; running; {
		select {
		case <-sigChan:
			log.Printf("Shutdown DLNA media server ...")
			ms.Shutdown()
		case err = <-ms.ErrChan():
			if err != nil {
				log.Printf("Media server closed with error: %s", err)
			} else {
				log.Println("Media server closed normally!")
			}
			running = false
		}
	}

}
