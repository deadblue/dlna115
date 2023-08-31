package login

import (
	"flag"
	"fmt"

	"github.com/deadblue/elevengo/option"
)

const (
	defaultPlatform = "linux"
)

func (c *Command) Init(args []string) (err error) {
	// Parse args
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Usage = func() {}
	var platform string
	fs.StringVar(&platform, "platform", defaultPlatform, "")
	fs.StringVar(&platform, "p", defaultPlatform, "")
	if err = fs.Parse(args); err != nil {
		return
	}

	// Save args
	switch platform {
	case "mac":
		c.Platform = option.QrcodeLoginMac
	case "windows":
		c.Platform = option.QrcodeLoginWindows
	default:
		c.Platform = option.QrcodeLoginLinux
	}
	if fs.NArg() > 0 {
		c.SaveFile = fs.Arg(0)
	}
	return
}

const usageTemplate = `
Usage: %s [-p paltform] [save-file]

Description:
    %s

Arguments:
    -p, -platform <platform>
        Simulte login on given platform.
        Supported platform: linux/mac/windows
    save-file
        File to save login cookie.

Example: 
    %s -p linux cookie.txt

`

func (c *Command) PrintUsage(prog string) {
	cmd := fmt.Sprintf("%s %s", prog, name)
	fmt.Printf(usageTemplate, cmd, desc, cmd)
}
