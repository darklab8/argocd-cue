package manifests

import (
	"testing"

	"github.com/darklab8/argocd-cue/argocue/types"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func TestManifests(t *testing.T) {
	manifests_folder := utils_filepath.Join(types.ProjectRoot, "examples", "manifests")
	NewManifests().Generate(utils_types.FilePath(manifests_folder))
}

func TestManifestsParams(t *testing.T) {
	manifests_folder := utils_filepath.Join(types.ProjectRoot, "examples", "manifests")
	NewManifests().GetParameters(utils_types.FilePath(manifests_folder))
}
