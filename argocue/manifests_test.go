package argocue

import (
	"testing"

	"github.com/darklab8/go-utils/goutils/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

var ProjectRoot = utils_filepath.Dir(utils.GetCurrentFolder())

func TestManifests(t *testing.T) {
	manifests_folder := utils_filepath.Join(ProjectRoot, "examples", "manifests")
	NewManifests().Generate(utils_types.FilePath(manifests_folder))
}

func TestManifestsParams(t *testing.T) {
	manifests_folder := utils_filepath.Join(ProjectRoot, "examples", "manifests")
	NewManifests().GetParameters(utils_types.FilePath(manifests_folder))
}
