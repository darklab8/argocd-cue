package identifier

import (
	"os"
	"strings"

	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Deployment string

const (
	Helm      Deployment = "helm_tool.cue"
	Manifests Deployment = "manifests_tool.cue"
)

func containsAnyFile(workdir utils_types.FilePath, filename string) bool {
	files, err := os.ReadDir(workdir.ToString())
	logus.LogStdout.CheckFatal(err, "contains any file")

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			return true
		}
	}
	return false
}

func IdentifyDeployment(workdir utils_types.FilePath) Deployment {
	if containsAnyFile(workdir, string(Helm)) {
		return Helm
	}

	if containsAnyFile(workdir, string(Manifests)) {
		return Manifests
	}

	logus.LogStdout.Panic(
		"not recognized package type, expected package file",
		typelog.String("manifests_package_identifier", string(Manifests)),
		typelog.String("Helm_package_identifier", string(Helm)),
	)
	panic("Not recognized kubernetes package")
}
