package daemon

import (
	"errors"
	"flag"
	"fmt"
)

var (
	ErrMissConfigFile = errors.New("config file required")
)

func (c *Command) Init(args []string) (err error) {
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Usage = func() {}

	fs.StringVar(&c.ConfigFile, "config", "", "")
	fs.StringVar(&c.ConfigFile, "c", "", "")
	if err = fs.Parse(args); err != nil {
		return
	}
	if c.ConfigFile == "" {
		err = ErrMissConfigFile
	}
	return
}

const usageTemplate = `
Usage: %s -c <config-file>

Description:
    %s

Arguments:
    -c, -config <config-file>
        Run daemon with config file.

Example: 
    %s -c config.yaml

`

func (c *Command) PrintUsage(prog string) {
	cmd := fmt.Sprintf("%s %s", prog, name)
	fmt.Printf(usageTemplate, cmd, desc, cmd)
}
