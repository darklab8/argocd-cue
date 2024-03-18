package helm

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/argocd-cue/argocue/settings"
	"github.com/darklab8/argocd-cue/argocue/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
	"gopkg.in/yaml.v3"
)

type Helm struct {
	parameters *settings.AppParameters
}

func NewHelm(parameters *settings.AppParameters) Helm { return Helm{parameters: parameters} }

func helmLoadDeps(workdir utils_types.FilePath) {
	build := exec.Command("helm", "dep", "update", "--kube-insecure-skip-tls-verify")
	build.Dir = workdir.ToString()
	build_out, err := build.Output()
	utils.HandleCmdError(build_out, err, "failed to run helm dep update")
}

func buildHelm(workdir utils_types.FilePath) {
	build := exec.Command("cue", "cmd", "build")
	build.Dir = workdir.ToString()
	build_out, err := build.Output()
	utils.HandleCmdError(build_out, err, "failed to cue cmd build")
}

func (h Helm) Generate(workdir utils_types.FilePath) {
	buildHelm(workdir)
	// HelmLoadDeps(workdir) // Not working correctly yet. TODO fix.

	command_exec := "helm"
	templating_commands := []string{"template"}

	templating_commands = append(templating_commands, h.parameters.HelmTemplateArgs...)

	templating_commands = append(templating_commands, ".")

	cmd := exec.Command(command_exec, templating_commands...)
	cmd.Dir = workdir.ToString()
	out, err := cmd.Output()
	utils.HandleCmdError(out, err, "failed to helm template")
	fmt.Println(string(out))
}

func newHelmParams(Map map[string]string) []settings.GetParameters {
	return []settings.GetParameters{
		{
			Name:           "helm-parameters",
			Title:          "Helm Parameters",
			CollectionType: "map",
			Map:            Map,
		},
	}
}

func (h Helm) GetParameters(workdir utils_types.FilePath) {
	buildHelm(workdir)
	data := make(map[string]string)

	yfile, err := os.ReadFile(utils_filepath.Join(workdir, "values.yaml").ToString())

	logus.LogStdout.CheckFatal(err, "Failed to read values")

	err2 := yaml.Unmarshal(yfile, &data)
	logus.LogStdout.CheckFatal(err2, "failed to unmarshal yaml")

	jsoned, err := json.Marshal(newHelmParams(data))
	logus.LogStdout.CheckWarn(err, "not able to marshal params")
	fmt.Println(string(jsoned))
}
