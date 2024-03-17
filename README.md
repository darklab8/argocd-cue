Argo cd plugin to support cue language.

![example](docs/example.png)

# Examples

[See folder examples](examples)

# Features

- you can use this plugin and cue to implement
    - kubernetes manfiests
    - helm charts

- it is in customize format with pinned all dependencies
- using cue 0.7.1,  argocd v2.8.11, kubectl v1.29.1 at the moment of this writing
    - dependencies are easily adjustable in install_deps.sh
- tested on
  - local [Kind cluster](https://kind.sigs.k8s.io/) for amd64 cpu arch, kind v0.22.0, kube v1.29.2
  - microk8s with arm64 cpu architecture too, microk8s 1.28.7
- written very close according to [official guide](https://argo-cd.readthedocs.io/en/stable/operator-manual/config-management-plugins/) and their [helm example plugin](https://github.com/argoproj/argo-cd/tree/master/examples/plugins/helm)

# Getting started with kind cluster

- install
    - [Taskfile](https://taskfile.dev/installation/)
    - [Kind](https://kind.sigs.k8s.io/docs/user/quick-start/)

- run `task cluster:create`
- run `task cluster:scan`
    - get observed ipaddress+port
    - replace cluster address+port in ~/.kube/config for kind cluster with this value

- deploy argo with `task argo:deploy` ( kubectl apply -k ./plugin )
- deploy sample apps `task argo:apply`
