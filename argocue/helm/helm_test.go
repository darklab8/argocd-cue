package helm

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/darklab8/argocd-cue/argocue/settings"
	"github.com/darklab8/argocd-cue/argocue/utils"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
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
	manifests_folder := utils_filepath.Join(utils.ProjectRoot, "examples", "helm")
	CleanFromYaml(manifests_folder)

	NewHelm(settings.NewApplication()).Generate(utils_types.FilePath(manifests_folder))
}

func TestHelmParams(t *testing.T) {
	manifests_folder := utils_filepath.Join(utils.ProjectRoot, "examples", "helm")
	CleanFromYaml(manifests_folder)
	NewHelm(settings.NewApplication()).GetParameters(utils_types.FilePath(manifests_folder))
}
