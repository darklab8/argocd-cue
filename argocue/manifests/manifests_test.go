package manifests

import (
	"testing"

	"github.com/darklab8/argocd-cue/argocue/settings"
	"github.com/darklab8/argocd-cue/argocue/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func TestManifests(t *testing.T) {
	manifests_folder := utils_filepath.Join(utils.ProjectRoot, "examples", "manifests")
	NewManifests(settings.NewParameters()).Generate(utils_types.FilePath(manifests_folder))
}

func TestManifestsParams(t *testing.T) {
	manifests_folder := utils_filepath.Join(utils.ProjectRoot, "examples", "manifests")
	NewManifests(settings.NewParameters()).GetParameters(utils_types.FilePath(manifests_folder))
}
