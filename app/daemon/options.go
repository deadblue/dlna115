package daemon

import (
	"os"

	"github.com/deadblue/dlna115/pkg/mediaserver"
	"github.com/deadblue/dlna115/pkg/storage/impl"
	"gopkg.in/yaml.v3"
)

type Options struct {
	// Storage options
	Storage impl.Options `yaml:"storage"`
	// Media options
	Media mediaserver.Options `yaml:"media"`
}

func (opts *Options) Load(configFile string) (err error) {
	// Parse config file
	file, err := os.Open(configFile)
	if err != nil {
		return
	}
	defer file.Close()
	decoder := yaml.NewDecoder(file)
	return decoder.Decode(opts)
}
