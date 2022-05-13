#!/usr/bin/env bash

set -e

POD=$(kubectl get pods -l "app=source-controller" -o name)
kubectl exec "$POD" -- printenv > .tmp/.env
