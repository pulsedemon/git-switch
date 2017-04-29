package main

import (
  "log"
  "github.com/docopt/docopt-go"
)

var Version = "0.0.1"

const Usage = `
  Usage:
  Options:
    -v, --version    output version
`

func main() {

  args, err := docopt.Parse(Usage, nil, true, Version, false)

  if err != nil {
		log.Fatalf("error: %s", err)
	}

  log.Fatalf("error: %s", args)


}
