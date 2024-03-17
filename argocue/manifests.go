package argocue

import (
	"fmt"
	"os/exec"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func RenderManifest(workdir utils_types.FilePath) {
	cmd := exec.Command("cue", "cmd", "dump")
	cmd.Dir = workdir.ToString()
	out, err := cmd.Output()

	logus.Log.CheckFatal(err, "failed to execute command", typelog.String("stdout", string(out)))
	fmt.Println(string(out))
}
