package argocue

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/darklab8/argocue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
	"gopkg.in/yaml.v3"
)

func BuildHelm(workdir utils_types.FilePath) {
	build := exec.Command("cue", "cmd", "build")
	build.Dir = workdir.ToString()
	build_out, err := build.Output()
	logus.Log.CheckFatal(err, "failed to execute command build", typelog.String("stdout", string(build_out)))
}

func RenderHelm(workdir utils_types.FilePath) {
	BuildHelm(workdir)

	cmd := exec.Command("helm", "template", ".")
	cmd.Dir = workdir.ToString()
	out, err := cmd.Output()
	logus.Log.CheckFatal(err, "failed to execute command template", typelog.String("stdout", string(out)), typelog.OptError(err))
	fmt.Println(string(out))
}

func NewHelmParams(Map map[string]interface{}) []ApplicationParams {
	return []ApplicationParams{
		{
			Name:           "helm-parameters",
			Title:          "Helm Parameters",
			CollectionType: "map",
			Map:            Map,
		},
	}
}

func GetHelmParameters(workdir utils_types.FilePath) {
	BuildHelm(workdir)
	data := make(map[string]interface{})

	yfile, err := os.ReadFile(utils_filepath.Join(workdir, "values.yaml").ToString())

	logus.Log.CheckFatal(err, "Failed to read values")

	err2 := yaml.Unmarshal(yfile, &data)
	logus.Log.CheckFatal(err2, "failed to unmarshal yaml")

	jsoned, err := json.Marshal(NewHelmParams(data))
	logus.Log.CheckWarn(err, "not able to marshal params")
	fmt.Println(string(jsoned))
}
