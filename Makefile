image=quay.io/cloudflavor/optimus
vers=v0.1.0-alpha
gobuild=go build -o _output/bin/optimus
commit=$(shell git rev-parse --short HEAD)
tag=$(shell git describe --abbrev=0 --tags)

build: verify
	rm -rf _output || true
	mkdir -p _output/bin/
	GOOS=linux GOARCH=amd64 $(gobuild) -v -ldflags "-X main.commit=$(commit) -X main.version=$(vers)" ./cmd/optimus/main.go

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
	echo "no tests available"

gen:
	hack/update-codegen.sh

verify:
	hack/verify-codegen.sh
