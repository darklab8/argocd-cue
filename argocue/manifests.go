package argocue

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/darklab8/argocue/argocue/logus"
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

func NewManifestsParams(Map map[string]interface{}) []ApplicationParams {
	return []ApplicationParams{
		{
			Name:           "manifests-parameters",
			Title:          "Manifests Parameters",
			CollectionType: "map",
			Map:            Map,
		},
	}
}

func GetManifestsParameters(workdir utils_types.FilePath) {
	jsoned, err := json.Marshal(NewManifestsParams(map[string]interface{}{}))
	logus.Log.CheckWarn(err, "not able to marshal params")
	fmt.Println(string(jsoned))
}
