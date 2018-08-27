#!/bin/bash
echo "utility needed: go get -u github.com/golang/dep/cmd/dep"

docker run -ti -v $(pwd):/go/src/github.com/pi-victor/pipelines -w /go/src/github.com/pi-victor/pipelines golang:1.10-stretch /bin/bash
