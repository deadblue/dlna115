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
	fs.StringVar(&c.secret, "secret", "", "")
	fs.StringVar(&c.secret, "s", "", "")
	if err = fs.Parse(args); err != nil {
		return
	}

	// Save args
	switch platform {
	case "mac":
		c.platform = option.QrcodeLoginMac
	case "windows":
		c.platform = option.QrcodeLoginWindows
	default:
		c.platform = option.QrcodeLoginLinux
	}
	if fs.NArg() > 0 {
		c.saveFile = fs.Arg(0)
	}
	return
}

const usageTemplate = `
Usage: %s [-p paltform] [-s secret-key] [save-file]

Description:
    %s

Arguments:
    -p, -platform <platform>
        Simulte login on given platform.
        Supported platform: linux/mac/windows, default is linux.
    -s, -secret <secret-key>
        Secret key to encrypt credential, keep it secret!
    save-file
        File to save credential.

Example: 
    %s -p linux -s sesame credential.txt

`

func (c *Command) PrintUsage(prog string) {
	cmd := fmt.Sprintf("%s %s", prog, name)
	fmt.Printf(usageTemplate, cmd, desc, cmd)
}
