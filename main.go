package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/docopt/docopt-go"
)

var osUserInfo, _ = user.Current()
var userDirectory = osUserInfo.HomeDir
var configFile = fmt.Sprintf("%s/.git_switch_config", userDirectory)

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

func decodeConfig() tomlConfig {
	var config tomlConfig
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return config
}

func getCurrentGitUser() (string, string) {
	userName, _ := exec.Command("git", "config", "--global", "user.name").Output()
	userEmail, _ := exec.Command("git", "config", "--global", "user.email").Output()

	return fmt.Sprintf("%s", userName), fmt.Sprintf("%s", userEmail)
}

func switchGitUser(userName string, userEmail string) {
	exec.Command("git", "config", "--global", "user.name", userName).Run()
	exec.Command("git", "config", "--global", "user.email", userEmail).Run()

	fmt.Printf("  \033[1m%s\033[m\n", "Current user:")
	fmt.Printf("  %s\n", userName)
	fmt.Printf("  %s\n", userEmail)
}

func printAvailableUsers() {
	var config = decodeConfig()

	println()
	fmt.Printf("  \033[1m%s\033[m\n", "Available users:")
	for _, author := range config.Authors {
		fmt.Printf("  %s\n", author.Name)
		fmt.Printf("  %s\n", author.Email)
		println()
	}

	return
}

func searchForUser(query string) {
	println()
	matches := make([]int, 0)
	fmt.Printf("  \033[1m%s\033[m %s\n", "Searching for user:", query)

	var config = decodeConfig()

	println()

	for index, author := range config.Authors {

		if strings.Contains(strings.ToUpper(author.Name), strings.ToUpper(query)) {
			matches = append(matches, index)
		} else if strings.Contains(strings.ToUpper(author.Email), strings.ToUpper(query)) {
			matches = append(matches, index)
		}
	}

	if len(matches) > 1 {
		fmt.Printf("  \033[1;33m%s\033[m %s\n", "More than one match found for query:", query)
		fmt.Printf("  %s\n", "Try your search again with a more specific query.")
		printAvailableUsers()
	} else if len(matches) == 1 {
		fmt.Printf("  \033[1m%s\033[m\n", "Match found.  Switching...\n")
		switchGitUser(config.Authors[matches[0]].Name, config.Authors[matches[0]].Email)
	} else {
		fmt.Printf("  \033[1;31m%s\033[m %s\n", "No matches for:", query)
	}
	return
}

func runCommand(cmd string, args []string) (err error) {
	argv := make([]string, 1)
	argv[0] = cmd

	switch cmd {
	case "ls":
		currentUserName, currentUserEmail := getCurrentGitUser()
		println()
		fmt.Printf("  \033[1m%s\033[m\n", "Current user:")
		fmt.Printf("  %s", currentUserName)
		fmt.Printf("  %s", currentUserEmail)

		printAvailableUsers()
		return

	case "setup":
		// TODO: finish this
		fmt.Println((configFile))
		return

	case "help", "":
		return goRun("main.go", []string{"--help"})
	}

	if cmd != "" && len(args) == 0 {
		searchForUser(cmd)
		return
	}

	return fmt.Errorf("%s is not a git-switch command. See 'git-switch help'", cmd)
}
