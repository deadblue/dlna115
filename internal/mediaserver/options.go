package mediaserver

type Options struct {
	// Listening port
	Port uint
	// Unique ID
	UUID string
	// Friendly name
	Name string

	// Play video use HLS or mp4, make sure your player supports this.
	UseHLS bool
}
