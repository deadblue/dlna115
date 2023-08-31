package daemon

import (
	"flag"
	"os"

	"github.com/deadblue/dlna115/pkg/mediaserver"
	"github.com/deadblue/dlna115/pkg/mediaserver/service/storage115"
	"gopkg.in/yaml.v3"
)

type Options struct {
	// Storage options
	Storage storage115.Options `yaml:"storage"`
	// Media options
	Media mediaserver.Options `yaml:"media"`
}

func (opts *Options) Init() (err error) {
	// Get config file from command line
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	var configFile string
	fs.StringVar(&configFile, "config", "", "Config file")
	if err = fs.Parse(os.Args[1:]); err != nil {
		return
	}
	// Parse config file
	file, err := os.Open(configFile)
	if err != nil {
		return
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	return decoder.Decode(opts)
}
