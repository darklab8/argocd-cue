package argocue

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/darklab8/argocd-cue/argocue/pack"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
	"github.com/stretchr/testify/assert"
)

func CleanFromYaml(manifests_folder utils_types.FilePath) {
	filepath.Walk(manifests_folder.ToString(),
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.Contains(info.Name(), "yaml") {
				os.Remove(path)
			}
			return nil
		})
}

func TestHelm(t *testing.T) {
	manifests_folder := utils_filepath.Join(ProjectRoot, "examples", "helm")
	CleanFromYaml(manifests_folder)
	RenderHelm(utils_types.FilePath(manifests_folder))
}

func TestHelmParams(t *testing.T) {
	manifests_folder := utils_filepath.Join(ProjectRoot, "examples", "helm")
	CleanFromYaml(manifests_folder)
	GetHelmParameters(utils_types.FilePath(manifests_folder))
}

func TestType(t *testing.T) {
	helm_folder := utils_filepath.Join(ProjectRoot, "examples", "helm")
	assert.True(t, containsAnyFile(helm_folder, "helm_tool.cue"))
	assert.False(t, containsAnyFile(helm_folder, "helm_tool2.cue"))

	manifests_folder := utils_filepath.Join(ProjectRoot, "examples", "manifests")
	assert.True(t, containsAnyFile(manifests_folder, "manifests_tool.cue"))
	assert.False(t, containsAnyFile(manifests_folder, "manifests_tool2.cue"))

	assert.Equal(t, IdentifyPackage(helm_folder), pack.Helm)
	assert.Equal(t, IdentifyPackage(manifests_folder), pack.Manifests)
}
