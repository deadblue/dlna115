package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/deadblue/dlna115/app"
	"github.com/deadblue/dlna115/app/daemon"
	"github.com/deadblue/dlna115/app/login"
	"github.com/deadblue/dlna115/pkg/version"
)

const (
	usagePrefix = `
%s 

Usage: %s <command> [command-arguments]

Commands:

`

	usageSuffix = `
Run "%s help <command>" to check usage for command.

`
)

type CommandRegistry struct {
	names []string
	cmds  map[string]app.Command
}

func (cr *CommandRegistry) Register(cmds ...app.Command) {
	if cr.cmds == nil {
		cr.cmds = make(map[string]app.Command)
	}
	for _, cmd := range cmds {
		cr.names = append(cr.names, cmd.Name())
		cr.cmds[cmd.Name()] = cmd
	}
}

func (cr *CommandRegistry) PrintUsage(prog string) {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, usagePrefix, version.Full(), prog)
	for _, name := range cr.names {
		cmd := cr.cmds[name]
		fmt.Fprintf(sb, "    %s\t%s\n", name, cmd.Desc())
	}
	fmt.Fprintf(sb, usageSuffix, prog)
	print(sb.String())
}

func (cr *CommandRegistry) PrintHelp(prog string, args []string) {
	cmdName := ""
	if len(args) > 0 {
		cmdName = args[0]
	}
	if cmd, ok := cr.cmds[cmdName]; ok {
		cmd.PrintUsage(prog)
	} else {
		cr.PrintUsage(prog)
	}
}

func (cr *CommandRegistry) Get(name string) (cmd app.Command, ok bool) {
	cmd, ok = cr.cmds[name]
	return
}

func main() {
	// Load commands
	cr := &CommandRegistry{}
	cr.Register(
		&login.Command{}, &daemon.Command{},
	)

	// Extract arguments
	progName, cmdName := os.Args[0], ""
	var cmdArgs []string
	if len(os.Args) > 1 {
		cmdName, cmdArgs = os.Args[1], os.Args[2:]
	}

	if cmd, ok := cr.Get(cmdName); !ok {
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
