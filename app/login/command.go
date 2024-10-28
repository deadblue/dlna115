package login

const (
	name = "login"
	desc = "Simulate 115 client login and export credential."
)

type Command struct {
	// Secret to encrypt cookie
	secret string
	// File to save cookie
	saveFile string
}

func (c *Command) Name() string {
	return name
}

func (c *Command) Desc() string {
	return desc
}
