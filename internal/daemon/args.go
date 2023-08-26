package daemon

import (
	"flag"
	"os"

	"github.com/deadblue/dlna115/internal/mediaserver"
	"github.com/google/uuid"
)

type Args struct {
	// Media options
	Media mediaserver.Options
	// Start an SSDP server for M-SEARCH
	WithSSDP bool
}

func (a *Args) Init() (err error) {
	// Get config from commandline
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	// Read
	fs.UintVar(&a.Media.Port, "media-port", 8115, "Listening port of media server.")
	fs.StringVar(&a.Media.UUID, "media-uuid", "", "The UUID of media server.")
	fs.BoolVar(&a.Media.UseHLS, "media-hls", true, "Streaming video via HLS.")
	fs.BoolVar(&a.WithSSDP, "ssdp", false, "Start an SSDP server.")
	var configFile string
	fs.StringVar(&configFile, "config", "", "Config file")
	if err = fs.Parse(os.Args[1:]); err != nil {
		return
	}
	if a.Media.UUID == "" {
		a.Media.UUID = uuid.NewString()
	}
	return
}
