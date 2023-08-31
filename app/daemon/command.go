package daemon

const (
	name = "daemon"
	desc = "Run DLNA115 daemon."
)

type Command struct {
	ConfigFile string
}

func (c *Command) Name() string {
	return name
}

func (c *Command) Desc() string {
	return desc
}
