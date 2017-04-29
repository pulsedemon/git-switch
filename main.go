package main

import (
	"fmt"
	"os"

	"github.com/docopt/docopt-go"
)

func main() {

	usage := `usage: git  [--version]
    <command> [<args>...]

    commands:
      ls      list available users
      help    display help

    options:
       -v --version
       -h, --help
`

	args, _ := docopt.Parse(usage, nil, true, "0.0.1", false)
	cmd := args["<command>"].(string)
	cmdArgs := args["<args>"].([]string)

	command_err := runCommand(cmd, cmdArgs)
	if command_err != nil {
		fmt.Println(command_err)
		os.Exit(1)
	}

}

func runCommand(cmd string, args []string) (err error) {
	argv := make([]string, 1)
	argv[0] = cmd

	switch cmd {
	case "ls":
		fmt.Println("ran ls")
		return
	case "help", "":
		fmt.Println("asked for help")
		return
	}

	return fmt.Errorf("%s is not a git-switch command. See 'git-switch help'", cmd)
}
