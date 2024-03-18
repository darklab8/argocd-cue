package argocue

import (
	"os/exec"

	"github.com/darklab8/argocd-cue/argocue/identifier"
	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Deployment interface {
	Generate(utils_types.FilePath)
	GetParameters(utils_types.FilePath)
}

func Run(workdir utils_types.FilePath, command Command) {

	package_type := identifier.IdentifyDeployment(workdir)
	var deployment Deployment

	switch package_type {
	case identifier.Manifests:
		deployment = NewManifests()
	case identifier.Helm:
		deployment = NewHelm()
	default:
		logus.LogStdout.Fatal("not recognized package type")
	}

	switch command {
	case Commands.Generate:
		deployment.Generate(workdir)
	case Commands.GetParameters:
		deployment.GetParameters(workdir)
	default:
		logus.LogStdout.Fatal("not chosen command", typelog.Any("commands", Commands))
	}
}

func HandleCmdError(out []byte, err error, msg string) {
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			logus.LogStdout.CheckPanic(err,
				msg,
				typelog.String("stdout", string(out)),
				typelog.String("stderr", string(err.Stderr)),
			)
		}
		if err, ok := err.(*exec.Error); ok {
			logus.LogStdout.CheckPanic(err,
				msg,
				typelog.String("stdout", string(out)),
				typelog.String("stderr", string(err.Error())),
			)
		}

		logus.LogStdout.CheckPanic(err,
			msg,
			typelog.String("stdout", string(out)),
			typelog.String("stderr", string(err.Error())),
		)
	}
}
