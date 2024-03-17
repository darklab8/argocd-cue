package argocue

import (
	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func Run(workdir utils_types.FilePath, command Command) {
	switch command {
	case Commands.Generate:
		{

		}
	default:
		logus.Log.Fatal("not chosen command", typelog.Any("commands", Commands))
	}
}
