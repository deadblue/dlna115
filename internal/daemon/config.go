package daemon

import (
	"os"

	"gopkg.in/yaml.v3"
)

type CredentialSourceConfig struct {
	Type   string `yaml:"type"`
	Source string `yaml:"source"`
}

type Config struct {
	// Media config
	Media struct {
		Port uint   `yaml:"port,omitempty"`
		UUID string `yaml:"uuid,omitempty"`
		Name string `yaml:"name,omitempty"`
	} `yaml:"media"`

	// Storage config
	Storage struct {
		// Source to load 115 credential
		CredentialSource CredentialSourceConfig `yaml:"credential-source"`
		// Folders that displayed under root
		TopFolders []string `yaml:"top-folders"`
	} `yaml:"storage"`

	// SSDP config
	SSDP struct {
		Server bool `yaml:"server,omitempty"`
	} `yaml:"ssdp,omitempty"`
}

func (c *Config) Load(fileName string) (err error) {
	// Open config file
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()
	// Read and parse it
	decoder := yaml.NewDecoder(file)
	return decoder.Decode(c)
}
