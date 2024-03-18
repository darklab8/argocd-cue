package settings

import (
	"encoding/json"
	"os"
	"strings"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
)

type GetParameters struct {
	Name           string            `json:"name"`
	Title          string            `json:"title"`
	CollectionType string            `json:"collectionType"`
	Map            map[string]string `json:"map"`
}

type AppParameterGroup struct {
	Name  string            `json:"name"`
	Map   map[string]string `json:"map,omitempty"`
	Array []string          `json:"array,omitempty"`
}

const (
	APP_PARAMETER_HELM_TEMPLATE_ARGS_KEY = "helm_template_args"
	APP_PARAMETER_COMMON_PARAMETERS_KEY  = "common_parameters"
	APP_PARAMETER_HELM_PARAMETERS_KEY    = "helm_parameters"
)

type AppParameters struct {
	HelmTemplateArgs []string
	CommonParameters struct {
		// TODO Not yet having effect, but it will be
		CueVersion *string `json:"cue_version"`
	}
	HelmParameters struct {
		// TODO Not yet having effect, but it will be
		HelmVersion *string `json:"helm_version"`

		HelmNameTemplate *string `json:"name_template"`
		HelmNamespace    *string `json:"namespace"`
	}
}

type AppOption func(app *AppParameters)

func NewParameters(opts ...AppOption) *AppParameters {
	app := &AppParameters{}

	for _, opt := range opts {
		opt(app)
	}

	return app
}

const (
	HELM_TEMPLATE_ARGS_KEY = "helm_template_args"
	APP_PARAM_KEY          = "map_parameters"
)

func (app *AppParameters) loadData(data []byte) {
	var argo_params []AppParameterGroup = make([]AppParameterGroup, 0)
	err := json.Unmarshal(data, &argo_params)
	logus.LogStdout.CheckFatal(err, "failed to unmarshal")

	for _, argo_param_group := range argo_params {

		if argo_param_group.Name == APP_PARAMETER_HELM_TEMPLATE_ARGS_KEY {
			app.HelmTemplateArgs = argo_param_group.Array

		} else if argo_param_group.Name == APP_PARAMETER_COMMON_PARAMETERS_KEY {
			marshaled, err := json.Marshal(argo_param_group.Map)
			logus.LogStdout.CheckPanic(err, "failed to marshal values for map parameter")
			err = json.Unmarshal(marshaled, &(app.CommonParameters))
			logus.LogStdout.CheckPanic(err, "failed to unmarshal values for map parameter")

		} else if argo_param_group.Name == APP_PARAMETER_HELM_PARAMETERS_KEY {
			marshaled, err := json.Marshal(argo_param_group.Map)
			logus.LogStdout.CheckPanic(err, "failed to marshal values for map parameter")
			err = json.Unmarshal(marshaled, &(app.HelmParameters))
			logus.LogStdout.CheckPanic(err, "failed to unmarshal values for map parameter")
		}
	}
}

func (a *AppParameters) Load() {
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

	a.loadData([]byte(app_parameters))
}
