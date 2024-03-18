package manifests

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/argocd-cue/argocue/types"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Manifests struct {
}

func NewManifests() Manifests { return Manifests{} }

func (m Manifests) Generate(workdir utils_types.FilePath) {
	cmd := exec.Command("cue", "cmd", "dump")
	cmd.Dir = workdir.ToString()
	out, err := cmd.Output()

	types.HandleCmdError(out, err, "failed to run cue cmd dump")
	fmt.Println(string(out))
}

func mewManifestsParams(Map map[string]interface{}) []types.ApplicationParams {
	return []types.ApplicationParams{
		{
			Name:           "manifests-parameters",
			Title:          "Manifests Parameters",
			CollectionType: "map",
			Map:            Map,
		},
	}
}

func (m Manifests) GetParameters(workdir utils_types.FilePath) {
	jsoned, err := json.Marshal(mewManifestsParams(map[string]interface{}{}))
	logus.LogStdout.CheckWarn(err, "not able to marshal params")
	fmt.Println(string(jsoned))
}
