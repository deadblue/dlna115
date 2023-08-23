package conf

import (
	"flag"
	"os"
)

type Config struct {
	// UUID of the device
	UUID string
	// Start an SSDP server for M-SEARCH
	SSDP bool
	// Return HLS link of video
	VideoHLS bool
	// Media Server Port
	MediaServerPort uint
}

func (c *Config) Init() (err error) {
	// Get config from commandline
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	fs.StringVar(&c.UUID, "uuid", "", "")
	fs.BoolVar(&c.SSDP, "ssdp", false, "")
	fs.BoolVar(&c.VideoHLS, "hls", true, "")
	fs.UintVar(&c.MediaServerPort, "media-port", 5115, "")
	if err = fs.Parse(os.Args[1:]); err != nil {
		return
	}
	return
}
