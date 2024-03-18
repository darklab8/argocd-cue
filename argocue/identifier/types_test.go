package identifier

import (
	"testing"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/stretchr/testify/assert"
)

var ProjectRoot = utils_filepath.Dir(utils_filepath.Dir(utils.GetCurrentFolder()))

func TestType(t *testing.T) {
	helm_folder := utils_filepath.Join(ProjectRoot, "examples", "helm")
	assert.True(t, containsAnyFile(helm_folder, "helm_tool.cue"))
	assert.False(t, containsAnyFile(helm_folder, "helm_tool2.cue"))

	manifests_folder := utils_filepath.Join(ProjectRoot, "examples", "manifests")
	assert.True(t, containsAnyFile(manifests_folder, "manifests_tool.cue"))
	assert.False(t, containsAnyFile(manifests_folder, "manifests_tool2.cue"))

	assert.Equal(t, IdentifyDeployment(helm_folder), Helm)
	assert.Equal(t, IdentifyDeployment(manifests_folder), Manifests)
}
