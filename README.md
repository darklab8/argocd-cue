Argo cd plugin to support cue language.

- all commands with descriptions in Taskfile.yaml
- plugin in kustomize format in folder `plugin`
- sample cue application for deploymetn in folder `sample`
- declarative application to try sample, in folder `application`

# Features

- it is in customize format with pinned all dependencies
- using cue 0.7.1,  argocd v2.8.11, kubectl v1.29.1 at the moment of this writing
- and added script to use amd64 or arm64 cue binary depending on host cpu arch
- deploy with `kubectl apply -k ./plugin`
  - see all commands in [Taskfile](https://github.com/darklab8/infra/blob/master/k8s/modules/argo_cue/Taskfile.yaml)
- added [sample application](https://github.com/darklab8/infra/tree/master/k8s/modules/argo_cue/application) to try it on, with command in Taskfile described too
- tested on
  - local [Kind cluster](https://kind.sigs.k8s.io/) for amd64 cpu arch, kind v0.22.0, kube v1.29.2
  - microk8s with arm64 cpu architecture too, microk8s 1.28.7
- written very close according to [official guide](https://argo-cd.readthedocs.io/en/stable/operator-manual/config-management-plugins/) and their [helm example plugin](https://github.com/argoproj/argo-cd/tree/master/examples/plugins/helm)
