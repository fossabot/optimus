#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ${GOPATH}/src/k8s.io/code-generator)}

${SCRIPT_ROOT}vendor/k8s.io/code-generator/generate-groups.sh all \
  github.com/PI-Victor/pipelines/pkg/client  github.com/PI-Victor/pipelines/pkg/apis \
  cloudflavor.io:v1 \
  --go-header-file ${SCRIPT_ROOT}/assets/license-header.txt
