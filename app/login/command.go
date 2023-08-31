package login

import (
	"github.com/deadblue/elevengo/option"
)

const (
	name = "login"
	desc = "Simulate 115 client login and export cookies."
)

type Command struct {
	Platform option.QrcodeLoginOption
	SaveFile string
}

func (c *Command) Name() string {
	return name
}

func (c *Command) Desc() string {
	return desc
}
