#!/bin/bash
echo "utility needed: go get -u github.com/golang/dep/cmd/dep"

docker run -ti -v $GOPATH:/go -w /go/src/github.com/cloudflavor/optimus golang:1.10-stretch /bin/bash
