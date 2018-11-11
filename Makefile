docker_tag ?= latest
docker_image ?= moorara/chia
docker_dir := /tmp/workspace


clean:
	@ rm -rf *.log coverage

dep:
	@ dep ensure -update

run:
	@ go run cmd/main.go

build:
	@ ./scripts/build.sh ./cmd/main.go ./chia

test:
	@ go test -v -race ./...

coverage:
	@ ./scripts/test-unit-cover.sh

docker:
	@ docker build -t $(docker_image):$(docker_tag) .

push:
	@ docker push $(docker_image):$(docker_tag)


.PHONY: clean
.PHONY: dep run build
.PHONY: test coverage
.PHONY: docker push
