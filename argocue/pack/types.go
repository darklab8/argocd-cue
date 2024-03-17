package pack

import (
	"os"
	"strings"

	"github.com/darklab8/argocue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type PackageType string

const (
	Helm      PackageType = "helm_tool.cue"
	Manifests PackageType = "manifests_tool.cue"
)

func containsAnyFile(workdir utils_types.FilePath, filename string) bool {
	files, err := os.ReadDir(workdir.ToString())
	logus.Log.CheckFatal(err, "contains any file")

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			return true
		}
	}
	return false
}

func IdentifyPackage(workdir utils_types.FilePath) PackageType {
	if containsAnyFile(workdir, string(Helm)) {
		return Helm
	}

	if containsAnyFile(workdir, string(Manifests)) {
		return Manifests
	}

	logus.Log.Panic(
		"not recognized package type, expected package file",
		typelog.String("manifests_package_identifier", string(Manifests)),
		typelog.String("Helm_package_identifier", string(Helm)),
	)
	panic("Not recognized kubernetes package")
}
