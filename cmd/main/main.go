package main

import (
	"log"

	"github.com/deadblue/dlna115/internal/daemon"
)

func main() {
	if err := daemon.Run(); err != nil {
		log.Fatalf("DLNA115 exited with error: %s", err)
	}

}
