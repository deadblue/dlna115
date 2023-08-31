package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deadblue/dlna115/app"
	"github.com/deadblue/dlna115/app/daemon"
	"github.com/deadblue/dlna115/app/login"
)

const (
	usagePrefix = `
Usage: %s <command> [command-arguments]

Commands:

`

	usageSuffix = `
Run "%s help <command>" to check usage for command.

`
)

type CommandRegistry map[string]app.Command

func (cr CommandRegistry) Register(cmds ...app.Command) {
	for _, cmd := range cmds {
		cr[cmd.Name()] = cmd
	}

}

func (cr CommandRegistry) PrintUsage(prog string) {
	prefix := fmt.Sprintf(usagePrefix, prog)
	suffix := fmt.Sprintf(usageSuffix, prog)
	sb := &strings.Builder{}
	sb.WriteString(prefix)
	for _, cmd := range cr {
		cmdInfo := fmt.Sprintf("    %s\t%s\n", cmd.Name(), cmd.Desc())
		sb.WriteString(cmdInfo)
	}
	sb.WriteString(suffix)
	print(sb.String())
}

func (cr CommandRegistry) PrintHelp(prog string, args []string) {
	cmdName := ""
	if len(args) > 0 {
		cmdName = args[0]
	}
	if cmd, ok := cr[cmdName]; ok {
		cmd.PrintUsage(prog)
	} else {
		cr.PrintUsage(prog)
	}
}

func main() {
	// Load commands
	cr := make(CommandRegistry)
	cr.Register(
		&login.Command{}, &daemon.Command{},
	)

	// Extract arguments
	progName, cmdName := os.Args[0], ""
	var cmdArgs []string
	if len(os.Args) > 1 {
		cmdName, cmdArgs = os.Args[1], os.Args[2:]
	}

	if cmd, ok := cr[cmdName]; !ok {
		if cmdName == "help" {
			cr.PrintHelp(progName, cmdArgs)
		} else {
			cr.PrintUsage(progName)
		}
		return
	} else {
		var err error
		if err = cmd.Init(cmdArgs); err != nil {
			cmd.PrintUsage(progName)
			return
		}
		if err = cmd.Run(); err != nil {
			log.Fatal(err)
		}
	}
}
