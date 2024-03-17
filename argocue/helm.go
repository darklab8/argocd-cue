package argocue

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
	"gopkg.in/yaml.v3"
)

func HelmLoadDeps(workdir utils_types.FilePath) {
	build := exec.Command("helm", "dep", "update", "--kube-insecure-skip-tls-verify")
	build.Dir = workdir.ToString()
	build_out, err := build.Output()
	HandleCmdError(build_out, err, "failed to run helm dep update")
}

func BuildHelm(workdir utils_types.FilePath) {
	build := exec.Command("cue", "cmd", "build")
	build.Dir = workdir.ToString()
	build_out, err := build.Output()
	HandleCmdError(build_out, err, "failed to cue cmd build")
}

func RenderHelm(workdir utils_types.FilePath) {
	BuildHelm(workdir)
	// HelmLoadDeps(workdir) // Not working correctly yet. TODO fix.
	cmd := exec.Command("helm", "template", ".")
	cmd.Dir = workdir.ToString()
	out, err := cmd.Output()
	HandleCmdError(out, err, "failed to helm template")
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
