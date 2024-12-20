package login

import (
	"github.com/deadblue/elevengo/option"
)

const (
	name = "login"
	desc = "Simulate 115 client login and export credential."
)

type Command struct {
	// Login options
	opts *option.QrcodeOptions
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
