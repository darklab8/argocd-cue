package settings

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
)

type GetParameters struct {
	Name           string                 `json:"name"`
	Title          string                 `json:"title"`
	CollectionType string                 `json:"collectionType"`
	Map            map[string]interface{} `json:"map"`
}

type ApplicationParameters struct {
	// Commented stuff for to be implemented a bit in a future
	Common struct {
		// CueVersion *string `json:"cue_version"`
	} `json:"common"`
	Helm struct {
		TemplateArgs []string `json:"template_args"`
		// HelmVersion  *string  `json:"helm_version"`
		// ReleaseName  *string  `json:"release_name"`
		// Namespace    *string  `json:"namespace"`
	} `json:"helm"`
}

type AppOption func(app *ApplicationParameters)

func NewApplication(opts ...AppOption) *ApplicationParameters {
	app := &ApplicationParameters{}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

const (
	ARGOCUE_PARAM_KEY = "argocd-cue-parameters"
)

func (a *ApplicationParameters) Load() {
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

	app_parameters, found_app_param := os.LookupEnv("ARGOCD_APP_PARAMETERS")

	if !found_app_param && app_parameters == "" {
		return
	}
	logus.LogFile.Info("found ARGOCD_APP_PARAMETERS", typelog.String("ARGOCD_APP_PARAMETERS", app_parameters))

	var parameter_groups []GetParameters = make([]GetParameters, 0)
	err := json.Unmarshal([]byte(app_parameters), &parameter_groups)
	if logus.LogFile.CheckWarn(err, "failed to unmarshal paramaeters") {
		return
	}
	logus.LogFile.Info("succesfully unmarhslaed",
		typelog.Int("len(parameter_groups)", len(parameter_groups)),
		typelog.String("data", app_parameters),
	)

	for _, parameteter_group := range parameter_groups {

		logus.LogFile.Info("found parameter group", typelog.String("name", parameteter_group.Name))
		if parameteter_group.Name == ARGOCUE_PARAM_KEY {

			marshaled_map, err := json.Marshal(parameteter_group.Map)
			if logus.LogFile.CheckError(err, "failed to marshal parameteter_group.Map") {
				continue
			}

			err = json.Unmarshal(marshaled_map, a)
			if logus.LogFile.CheckError(err, "failed to unarmshal parameteter_group.Map") {
				continue
			}
		}
	}

}
