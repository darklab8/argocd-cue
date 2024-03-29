#!/bin/sh

if uname -a | grep -q x86_64; then
    echo "x86_64"
    wget https://get.helm.sh/helm-v3.10.3-linux-amd64.tar.gz -O - | tar xz && mv linux-amd64/helm /usr/local/bin/helm && chmod a+x /usr/local/bin/helm
    wget https://github.com/cue-lang/cue/releases/download/v0.7.1/cue_v0.7.1_linux_amd64.tar.gz -O - | tar xz && mv cue /usr/local/bin/cue && chmod +x /usr/local/bin/cue
elif uname -a | grep -q aarch64; then
    echo "aarch64"
    wget https://get.helm.sh/helm-v3.10.3-linux-arm64.tar.gz -O - | tar xz && mv linux-arm64/helm /usr/local/bin/helm && chmod a+x /usr/local/bin/helm
    wget https://github.com/cue-lang/cue/releases/download/v0.7.1/cue_v0.7.1_linux_arm64.tar.gz -O - | tar xz && mv cue /usr/local/bin/cue && chmod +x /usr/local/bin/cue
else # neither amd64 or arm64
    echo "Unsupported CPU vendor: $(uname -a)" >&2
fi
