package manifests

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/argocd-cue/argocue/settings"
	"github.com/darklab8/argocd-cue/argocue/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Manifests struct {
	parameters *settings.AppParameters
}

func NewManifests(parameters *settings.AppParameters) Manifests {
	return Manifests{parameters: parameters}
}

func (m Manifests) Generate(workdir utils_types.FilePath) {
	cmd := exec.Command("cue", "cmd", "dump")
	cmd.Dir = workdir.ToString()
	out, err := cmd.Output()

	utils.HandleCmdError(out, err, "failed to run cue cmd dump")
	fmt.Println(string(out))
}

func mewManifestsParams(Map map[string]string) []settings.GetParameters {
	return []settings.GetParameters{
		{
			Name:           "manifests-parameters",
			Title:          "Manifests Parameters",
			CollectionType: "map",
			Map:            Map,
		},
	}
}

func (m Manifests) GetParameters(workdir utils_types.FilePath) {
	jsoned, err := json.Marshal(mewManifestsParams(map[string]string{}))
	logus.LogStdout.CheckWarn(err, "not able to marshal params")
	fmt.Println(string(jsoned))
}
