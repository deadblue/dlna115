package login

import (
	"flag"
	"fmt"

	"github.com/deadblue/elevengo/option"
)

func (c *Command) Init(args []string) (err error) {
	// Parse args
	fs := flag.NewFlagSet("", flag.ContinueOnError)
	fs.Usage = func() {}
	fs.StringVar(&c.secret, "secret", "", "")
	fs.StringVar(&c.secret, "s", "", "")
	var platform string
	fs.StringVar(&platform, "platform", "", "")
	fs.StringVar(&platform, "p", "", "")
	if err = fs.Parse(args); err != nil {
		return
	}
	// Save args
	c.opts = option.Qrcode()
	switch platform {
	case "android":
		c.opts.LoginAndroid()
	case "ios":
		c.opts.LoginIos()
	case "tv":
		c.opts.LoginTv()
	case "wechat":
		c.opts.LoginWechatMiniApp()
	case "alipay":
		c.opts.LoginAlipayMiniApp()
	case "qandroid":
		c.opts.LoginQandroid()
	}
	if fs.NArg() > 0 {
		c.saveFile = fs.Arg(0)
	}
	return
}

const usageTemplate = `
Usage: %s [-s secret-key] [-p platform] [save-file]

Description:
    %s

Arguments:
    -s, -secret <secret-key>
        Secret key to encrypt credential, keep it secret!
    -p, -platform <login-platform>
        Simulate login on specific platform.
        Supported: web/android/ios/tv/wechat/alipay/qandroid
        Default: web
    save-file
        File to save credential.

Example: 
    %s -s sesame credential.txt

`

func (c *Command) PrintUsage(prog string) {
	cmd := fmt.Sprintf("%s %s", prog, name)
	fmt.Printf(usageTemplate, cmd, desc, cmd)
}
