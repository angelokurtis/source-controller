#!/usr/bin/env bash

set -e

make generate manifests
make install

export IMG=kurtis/source-controller
export TAG=latest
export BUILD_ARGS=--push
export BUILD_PLATFORMS=amd64
make docker-build

export REPO_DIGEST=$(docker inspect --format='{{index .RepoDigests 0}}' $IMG)
kubectl set image deployment/source-controller manager="$REPO_DIGEST" -n fluxcd
