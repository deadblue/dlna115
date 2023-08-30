package mediaserver

type Options struct {
	// Listening port
	Port uint `yaml:"port,omitempty"`
	// Unique ID
	UUID string `yaml:"uuid,omitempty"`
	// Friendly name
	Name string `yaml:"name,omitempty"`
}
