package main

import (
	"os"

	"github.com/darklab8/argocd-cue/argocue"
	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/argocd-cue/argocue/settings"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		logus.LogStdout.Fatal("expected one of valid commands", typelog.Any("commands", settings.Commands))
	}

	workdir, err := os.Getwd()
	logus.LogStdout.CheckFatal(err, "failed to get workdir")
	argocue.Run(utils_types.FilePath(workdir), settings.Command(argsWithoutProg[0]))
}
