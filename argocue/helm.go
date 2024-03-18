package argocue

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
	"gopkg.in/yaml.v3"
)

type Helm struct {
}

func NewHelm() Helm { return Helm{} }

func helmLoadDeps(workdir utils_types.FilePath) {
	build := exec.Command("helm", "dep", "update", "--kube-insecure-skip-tls-verify")
	build.Dir = workdir.ToString()
	build_out, err := build.Output()
	HandleCmdError(build_out, err, "failed to run helm dep update")
}

func buildHelm(workdir utils_types.FilePath) {
	build := exec.Command("cue", "cmd", "build")
	build.Dir = workdir.ToString()
	build_out, err := build.Output()
	HandleCmdError(build_out, err, "failed to cue cmd build")
}

type ParameterGroup struct {
	Name  string   `json:"name"`
	Array []string `json:"array"`
}

func (h Helm) Generate(workdir utils_types.FilePath) {
	buildHelm(workdir)
	// HelmLoadDeps(workdir) // Not working correctly yet. TODO fix.

	command_exec := "helm"
	templating_commands := []string{"template"}

	if app_parameters, ok := os.LookupEnv("ARGOCD_APP_PARAMETERS"); ok && app_parameters != "" {
		logus.LogFile.Info("found ARGOCD_APP_PARAMETERS", typelog.String("ARGOCD_APP_PARAMETERS", app_parameters))

		var parameter_groups []ParameterGroup = make([]ParameterGroup, 0)
		err := json.Unmarshal([]byte(app_parameters), &parameter_groups)
		if !logus.LogFile.CheckWarn(err, "failed to unmarshal paramaeters") {
			logus.LogFile.Info("succesfully unmarhslaed", typelog.Int("len(parameter_groups)", len(parameter_groups)))

			for _, parameteter_group := range parameter_groups {
				logus.LogFile.Info("found parameter group", typelog.String("name", parameteter_group.Name))
				if parameteter_group.Name == "helm-template-args" {
					logus.LogFile.Info("found helm-template-args", typelog.Items("array", parameteter_group.Array))
					typeloged_envs := []typelog.LogType{}
					for i, value := range parameteter_group.Array {
						typeloged_envs = append(typeloged_envs, typelog.String(strconv.Itoa(i), value))
					}
					logus.LogFile.Info("all parameters", typeloged_envs...)

					templating_commands = append(templating_commands, parameteter_group.Array...)
				}
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

func newHelmParams(Map map[string]interface{}) []ApplicationParams {
	return []ApplicationParams{
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
	data := make(map[string]interface{})

	yfile, err := os.ReadFile(utils_filepath.Join(workdir, "values.yaml").ToString())

	logus.LogStdout.CheckFatal(err, "Failed to read values")

	err2 := yaml.Unmarshal(yfile, &data)
	logus.LogStdout.CheckFatal(err2, "failed to unmarshal yaml")

	jsoned, err := json.Marshal(newHelmParams(data))
	logus.LogStdout.CheckWarn(err, "not able to marshal params")
	fmt.Println(string(jsoned))
}
