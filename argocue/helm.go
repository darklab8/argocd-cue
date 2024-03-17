package argocue

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/argocd-cue/argocue/pack"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_filepath"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
	"gopkg.in/yaml.v3"
)

func BuildHelm(workdir utils_types.FilePath) {
	build := exec.Command("cue", "cmd", "build")
	build.Dir = workdir.ToString()
	build_out, err := build.Output()
	logus.Log.CheckFatal(err, "failed to execute command build", typelog.String("stdout", string(build_out)))
}

func RenderHelm(workdir utils_types.FilePath) {
	BuildHelm(workdir)

	cmd := exec.Command("helm", "template", ".")
	cmd.Dir = workdir.ToString()
	out, err := cmd.Output()
	logus.Log.CheckFatal(err, "failed to execute command template", typelog.String("stdout", string(out)), typelog.OptError(err))
	fmt.Println(string(out))
}

type HelmParams struct {
	Name           string `json:"name"`
	Title          string `json:"title"`
	CollectionType string `json:"collectionType"`
	Map            map[string]interface{}
}

func NewHelmParams(Map map[string]interface{}) []HelmParams {
	return []HelmParams{
		{
			Name:           "helm-parameters",
			Title:          "Helm Parameters",
			CollectionType: "map",
			Map:            Map,
		},
	}
}

func GetHelmParameters(workdir utils_types.FilePath) {
	BuildHelm(workdir)
	data := make(map[string]interface{})

	yfile, err := os.ReadFile(utils_filepath.Join(workdir, "values.yaml").ToString())

	logus.Log.CheckFatal(err, "Failed to read values")

	err2 := yaml.Unmarshal(yfile, &data)
	logus.Log.CheckFatal(err2, "failed to unmarshal yaml")

	jsoned, err := json.Marshal(NewHelmParams(data))
	logus.Log.CheckWarn(err, "not able to marshal params")
	fmt.Println(string(jsoned))
}

func containsAnyFile(workdir utils_types.FilePath, filename string) bool {
	files, err := os.ReadDir(workdir.ToString())
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			return true
		}
	}
	return false
}

func IdentifyPackage(workdir utils_types.FilePath) pack.Package {
	if containsAnyFile(workdir, "helm_tool.cue") {
		return pack.Helm
	}

	if containsAnyFile(workdir, "manifests_tool.cue") {
		return pack.Manifests
	}

	panic("Not recognized kubernetes package")
}
