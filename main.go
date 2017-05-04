package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"

	"github.com/docopt/docopt-go"
	"github.com/BurntSushi/toml"
)

var user_info, _ = user.Current()
var user_directory = user_info.HomeDir
var config_file = fmt.Sprintf("%s/.git_switch_config", user_directory)

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

	commanderr := runCommand(cmd, cmdArgs)
	if commanderr != nil {
		fmt.Println(commanderr)
		os.Exit(1)
	}

}

func goRun(scriptName string, args []string) (err error) {

	cmdArgs := make([]string, 2)
	cmdArgs[0] = "run"
	cmdArgs[1] = scriptName
	cmdArgs = append(cmdArgs, args...)
	osCmd := exec.Command("go", cmdArgs...)
	var out []byte
	out, err = osCmd.Output()
	fmt.Println(string(out))
	if err != nil {
		return
	}
	return

}

func runCommand(cmd string, args []string) (err error) {

	argv := make([]string, 1)
	argv[0] = cmd

	switch cmd {
	case "ls":
		fmt.Println("ran ls")
		return
	case "setup":
		fmt.Println((config_file))
		return
	case "help", "":
		return goRun("main.go", []string{"--help"})
	}

	return fmt.Errorf("%s is not a git-switch command. See 'git-switch help'", cmd)

}
