package main

import (
	"flag"

	"github.com/darklab8/argocd-cue/argocue"
	"github.com/darklab8/go-utils/goutils/utils"
)

func main() {
	var command string
	flag.StringVar(&command, "command", "undefined", "command to run")

	flag.Parse()
	argocue.Run(utils.GetCurrentFolder(), argocue.Command(command))
}
