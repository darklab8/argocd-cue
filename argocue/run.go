package argocue

import (
	"github.com/darklab8/argocd-cue/argocue/helm"
	"github.com/darklab8/argocd-cue/argocue/identifier"
	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/argocd-cue/argocue/manifests"
	"github.com/darklab8/argocd-cue/argocue/types"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

type Deployment interface {
	Generate(utils_types.FilePath)
	GetParameters(utils_types.FilePath)
}

func Run(workdir utils_types.FilePath, command types.Command) {

	package_type := identifier.IdentifyDeployment(workdir)
	var deployment Deployment

	switch package_type {
	case identifier.Manifests:
		deployment = manifests.NewManifests()
	case identifier.Helm:
		deployment = helm.NewHelm()
	default:
		logus.LogStdout.Fatal("not recognized package type")
	}

	switch command {
	case types.Commands.Generate:
		deployment.Generate(workdir)
	case types.Commands.GetParameters:
		deployment.GetParameters(workdir)
	default:
		logus.LogStdout.Fatal("not chosen command", typelog.Any("commands", types.Commands))
	}
}
