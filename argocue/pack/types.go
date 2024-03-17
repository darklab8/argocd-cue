package pack

import (
	"log"
	"os"
	"strings"

	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Package int64

const (
	Helm Package = iota
	Manifests
)

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

func IdentifyPackage(workdir utils_types.FilePath) Package {
	if containsAnyFile(workdir, "helm_tool.cue") {
		return Helm
	}

	if containsAnyFile(workdir, "manifests_tool.cue") {
		return Manifests
	}

	panic("Not recognized kubernetes package")
}
