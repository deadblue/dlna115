package conf

import (
	"flag"
	"os"

	"github.com/google/uuid"
)

type Config struct {
	// Port for meida server
	MediaPort uint
	// UUID of media server
	MediaUUID string

	// Return HLS link of video
	VideoHLS bool
	// Start an SSDP server for M-SEARCH
	SSDP bool
}

func (c *Config) Init() (err error) {
	// Get config from commandline
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.UintVar(&c.MediaPort, "port", 8115, "Listening port of media server.")
	fs.StringVar(&c.MediaUUID, "uuid", "", "The UUID of media server.")
	fs.BoolVar(&c.VideoHLS, "video-hls", true, "Streaming video via HLS.")
	fs.BoolVar(&c.SSDP, "with-ssdp", false, "Start an SSDP server.")
	if err = fs.Parse(os.Args[1:]); err != nil {
		return
	}
	if c.MediaUUID == "" {
		c.MediaUUID = uuid.NewString()
	}
	return
}
