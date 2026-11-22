#!/usr/bin/env bash
set -euo pipefail

# renovate: datasource=github-releases depName=pulumi/pulumi
PULUMI_VERSION=3.277.0
# renovate: datasource=github-releases depName=kubernetes-sigs/kind
KIND_VERSION=v0.33.0

curl -fsSL https://get.pulumi.com | sh -s -- --version "${PULUMI_VERSION}"
go install "sigs.k8s.io/kind@${KIND_VERSION}"
python3 -m pip install --user build
