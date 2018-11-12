image=quay.io/cloudflavor/pipelines
vers=v0.1.0-alpha
gobuild=go build -o _output/bin/pipelines
commit=$(shell git rev-parse --short HEAD)
tag=$(shell git describe --abbrev=0 --tags)

build: verify
	rm -rf _output || true
	mkdir -p _output/bin/
	$(gobuild) -v -ldflags "-X main.commit=$(commit) -X main.version=$(vers)" ./cmd/pipelines/main.go

generate:
	echo "test"

publish:
	docker build -t $(image):$(commit) .
	docker push $(image):$(commit)
	docker tag $(image):$(commit) $(image):latest
	docker push $(image):latest

push: tag
	docker build -t $(image):$(commit) .
	docker push $(image):$(commit)
	docker push $(image):$(vers)
	docker push $(image):latest
	git push --tags

tag:
	docker tag $(image):$(commit) $(image):latest
	docker tag $(image):$(commit) $(image):$(vers)
	git tag $(vers)

# TODO: implement coverage testing.
test:
	go test -race -coverprofile=coverage.txt -covermode=atomic -v ./pkg/plugins

gen:
	hack/update-codegen.sh

verify:
	hack/verify-codegen.sh
