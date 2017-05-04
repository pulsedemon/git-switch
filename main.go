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


type tomlConfig struct {
	Title   string
	Default userInfo
	Authors []userInfo
}

type userInfo struct {
	Name  string
	Email string
}

func main() {

	usage := `Usage: git-switch [--version]
    <command> [<args>...]

    commands:
      ls      list available users
      setup   generate config file if it doesn't exist
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


func getCurrentGitUser() (string, string) {
	user_name, _ := exec.Command("git", "config", "--global", "user.name").Output();
	user_email, _ := exec.Command("git", "config", "--global", "user.email").Output();

	return fmt.Sprintf("%s", user_name), fmt.Sprintf("%s", user_email)
}


func printAvailableUsers() {
	var config tomlConfig
	if _, err := toml.DecodeFile(config_file, &config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	println()
	fmt.Printf("  \033[1m%s\033[m\n", "Available users:")
	for _, author := range config.Authors {
		fmt.Printf("  %s\n", author.Name)
		fmt.Printf("  %s\n", author.Email)
		println()
	}
}


func runCommand(cmd string, args []string) (err error) {

	argv := make([]string, 1)
	argv[0] = cmd

	switch cmd {
	case "ls":
		current_user_name, current_user_email := getCurrentGitUser()
		println()
		fmt.Printf("  \033[1m%s\033[m\n", "Current user:")
		fmt.Printf("  %s", current_user_name)
		fmt.Printf("  %s", current_user_email)

		printAvailableUsers()
		return

	case "setup":
		// TODO: finish this
		fmt.Println((config_file))
		return

	case "help", "":
		return goRun("main.go", []string{"--help"})
	}

	return fmt.Errorf("%s is not a git-switch command. See 'git-switch help'", cmd)

}
