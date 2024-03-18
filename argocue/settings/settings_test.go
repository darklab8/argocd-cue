package settings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadSettings(t *testing.T) {

	data := []byte("[{\"array\":[\"--name-template=monitoring\",\"--namespace=production-monitoring\"],\"name\":\"helm_template_args\"}]")

	app := NewParameters()
	app.loadData(data)

	assert.Greater(t, app.HelmTemplateArgs, 0)
}

func TestLoadHelmParams(t *testing.T) {
	//json.dumps([{"name":"map_parameters","map":{"release_name":"customstuff"}}])
	data := []byte(`[{"name": "helm_parameters", "map": {"release_name": "customstuff"}}]`)
	app := NewParameters()
	app.loadData(data)

	assert.Equal(t, *app.HelmParameters.HelmReleaseName, "customstuff")
}

func TestLoadCommonParams(t *testing.T) {
	//json.dumps([{"name":"map_parameters","map":{"release_name":"customstuff"}}])
	data := []byte(`[{"name": "common_parameters", "map": {"cue_version": "customstuff"}}]`)
	app := NewParameters()
	app.loadData(data)

	assert.Equal(t, *app.CommonParameters.CueVersion, "customstuff")
}
