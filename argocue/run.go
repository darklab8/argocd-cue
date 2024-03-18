package argocue

import (
	"os/exec"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/argocd-cue/argocue/pack"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func Run(workdir utils_types.FilePath, command Command) {

	package_type := pack.IdentifyPackage(workdir)

	switch package_type {
	case pack.Manifests:

		switch command {
		case Commands.Generate:
			{
				RenderManifest(workdir)
			}
		case Commands.GetParameters:
			{
				GetManifestsParameters(workdir)
			}
		default:
			logus.LogStdout.Fatal("not chosen command", typelog.Any("commands", Commands))
		}

	case pack.Helm:
		switch command {
		case Commands.Generate:
			{
				RenderHelm(workdir)
			}
		case Commands.GetParameters:
			{
				GetHelmParameters(workdir)
			}
		default:
			logus.LogStdout.Fatal("not chosen command", typelog.Any("commands", Commands))
		}

	default:
		logus.LogStdout.Fatal("not recognized package type")
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
