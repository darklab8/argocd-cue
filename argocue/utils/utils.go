package utils

import (
	"os/exec"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
)

var ProjectRoot = utils_filepath.Dir(utils_filepath.Dir(utils.GetCurrentFolder()))

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
