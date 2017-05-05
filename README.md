## git-switch

If you need a way to easily and quickly modify your global git user config, try this out.

*This is still under active development*

## Installation

The easiest way to install right now would be:

`go get github.com/pulsedemon/git-switch`


## Commands

### Usage

```
$ git switch -h
Usage:
	git-switch ls
	git-switch add <name> <email>
	git-switch reset
	git-switch help
	git-switch -h | --help
	git-switch -v | --version
	git-switch <name_or_email>

Options:
   -v --version
   -h, --help

Commands:
  ls      list available users
	add 		add a new user to the config
  reset   switch back to the default user
  help    display help
```
