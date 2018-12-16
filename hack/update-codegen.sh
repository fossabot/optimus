#!/bin/bash

set -o errexit
set -o nounset
set -o pipefail

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd ${SCRIPT_ROOT}; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ${GOPATH}/src/k8s.io/code-generator)}

echo "You must have the k8s code-generator in $GOPATH/src/k8s.io/code-generator, othwerise this will fail"
../../../k8s.io/code-generator/generate-groups.sh all \
  github.com/cloudflavor/optimus/pkg/client  github.com/cloudflavor/optimus/pkg/apis \
  pipelines:v1 \
  --go-header-file ${SCRIPT_ROOT}/assets/license-header.txt
