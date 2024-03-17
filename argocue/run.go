package argocue

import (
	"github.com/darklab8/argocd-cue/argocue/logus"
	"github.com/darklab8/argocd-cue/argocue/pack"
	"github.com/darklab8/go-typelog/typelog"
	"github.com/darklab8/go-utils/goutils/utils/utils_types"
)

func Run(workdir utils_types.FilePath, command Command) {

	package_type := pack.IdentifyPackage(workdir)

	switch package_type {
	case pack.Manifests:

		switch command {
		case Commands.Generate:
			{
				RenderManifest(workdir)
			}
		case Commands.GetParameters:
			{
				GetManifestsParameters(workdir)
			}
		default:
			logus.Log.Fatal("not chosen command", typelog.Any("commands", Commands))
		}

	case pack.Helm:
		switch command {
		case Commands.Generate:
			{
				RenderHelm(workdir)
			}
		case Commands.GetParameters:
			{
				GetHelmParameters(workdir)
			}
		default:
			logus.Log.Fatal("not chosen command", typelog.Any("commands", Commands))
		}

	default:
		logus.Log.Fatal("not recognized package type")
	}

}
