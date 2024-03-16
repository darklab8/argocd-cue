Argo cd plugin to support cue language.

- all commands with descriptions in Taskfile.yaml
- plugin in kustomize format in folder `plugin`
- sample cue application for deploymetn in folder `sample`
- declarative application to try sample, in folder `application`

# Changelog

it is in customize format with pinned all dependencies
- using cue 0.7.1, argocd v2.8.11, kubectl v1.29.1 at the moment of this writing
- and added script to use amd64 or arm64 cue binary depending on host cpu arch
- deploy with kubectl apply -k ./plugin
  - see all commands in Taskfile
- added sample application to try it on, with command in Taskfile described too
- tested on
  - local Kind cluster for amd64 cpu arch, kind v0.22.0, kube v1.29.2
  - microk8s with arm64 cpu architecture too, microk8s 1.28.7
- written very close according to official guide and their helm example plugin
