package argocue

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
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

	command_exec := "helm"
	templating_commands := []string{"template"}

	if app_parameters, ok := os.LookupEnv("ARGOCD_APP_PARAMETERS"); ok && app_parameters != "" {
		logus.LogFile.Info("found ARGOCD_APP_PARAMETERS", typelog.String("ARGOCD_APP_PARAMETERS", app_parameters))

		parameters := make(map[string]interface{})
		err := json.Unmarshal([]byte(app_parameters), &parameters)
		if !logus.LogFile.CheckWarn(err, "failed to unmarshal args") {
			if value, ok := parameters["helm-template-args"]; ok {
				parameters_map := value.(map[string]string)
				typeloged_envs := []typelog.LogType{}
				for key, value := range parameters_map {
					typeloged_envs = append(typeloged_envs, typelog.String(key, value))
				}
				logus.LogFile.Info("all parameters", typeloged_envs...)
			}
		}
	}

	typeloged_envs := []typelog.LogType{}
	for _, env := range os.Environ() {
		values := strings.Split(env, "=")
		key := values[0]
		value := values[1]
		if len(value) == 2 && strings.Contains(key, "ARGOCD") {
			typeloged_envs = append(typeloged_envs, typelog.String(key, value))
		}
	}
	logus.LogFile.Info("all envs", typeloged_envs...)

	templating_commands = append(templating_commands, ".")

	cmd := exec.Command(command_exec, templating_commands...)
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

	logus.LogStdout.CheckFatal(err, "Failed to read values")

	err2 := yaml.Unmarshal(yfile, &data)
	logus.LogStdout.CheckFatal(err2, "failed to unmarshal yaml")

	jsoned, err := json.Marshal(NewHelmParams(data))
	logus.LogStdout.CheckWarn(err, "not able to marshal params")
	fmt.Println(string(jsoned))
}
