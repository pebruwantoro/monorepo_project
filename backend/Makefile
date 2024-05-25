.PHONY: build build-alpine clean test help default


BIN_NAME=go-feeder-boilerplate

VERSION    := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
IMAGE_NAME := "go-feeder-boilerplate"
GOTEST     := go test -v

GIT_COMMIT =$(shell git rev-parse HEAD)
GIT_DIRTY  =$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE =$(shell date '+%Y-%m-%d-%H:%M:%S')

default: test

help:
	@echo 'Management commands for {{values.component_id}}:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	@echo '    make build-alpine    Compile optimized for alpine linux.'
	@echo '    make package         Build final docker image with just the go binary inside'
	@echo '    make tag             Tag image created by package with latest, git commit and version'
	@echo '    make test            Run tests on a compiled project.'
	@echo '    make push            Push tagged images to registry'
	@echo '    make clean           Clean the directory tree.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X bitbucket.org/efishery/go-feeder-boilerplate/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X bitbucket.org/efishery/go-feeder-boilerplate/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

get-deps:
	dep ensure

build-alpine:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags '-w -linkmode external -extldflags "-static" -X bitbucket.org/efishery/go-feeder-boilerplate/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X bitbucket.org/efishery/go-feeder-boilerplate/version.BuildDate=${BUILD_DATE}' -o bin/${BIN_NAME}

package:
	@echo "building image ${BIN_NAME} ${VERSION} $(GIT_COMMIT)"
	docker build --build-arg VERSION=${VERSION} --build-arg GIT_COMMIT=$(GIT_COMMIT) -t $(IMAGE_NAME):local .

tag: 
	@echo "Tagging: latest ${VERSION} $(GIT_COMMIT)"
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):$(GIT_COMMIT)
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):${VERSION}
	docker tag $(IMAGE_NAME):local $(IMAGE_NAME):latest

push: tag
	@echo "Pushing docker image to registry: latest ${VERSION} $(GIT_COMMIT)"
	docker push $(IMAGE_NAME):$(GIT_COMMIT)
	docker push $(IMAGE_NAME):${VERSION}
	docker push $(IMAGE_NAME):latest

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	@echo "=================================================================================="
	@echo "Coverage Test"
	@echo "=================================================================================="
	go fmt ./... && $(GOTEST) -coverprofile coverage.cov -cover ./...
	@echo "\n"
	@echo "=================================================================================="
	@echo "All Package Coverage"
	@echo "=================================================================================="
	go tool cover -func coverage.cov

test-db:
	go clean -testcache && go test -cover --tags=integration ./...

generate:
	go generate ./...

lint:
	golangci-lint run ./...

docker-build: build-alpine
	docker build -f Dockerfile -t $(IMAGE_NAME):$(TAG) .
	docker images

docker-run:
	docker run --rm -p $(PORT):$(PORT) --env-file ${ROOT_DIR}/.env --mount type=bind,source=${ROOT_DIR}/.env,target="/.env" --network gourami_default --name=$(IMAGE_NAME) $(IMAGE_NAME)

rest: build
	@echo "  >  Starting Program..."
	./bin/${BIN_NAME} rest
	@echo "Process took $$(($$(date +%s)-$(STIME))) seconds"

mock-pkg:
	mockery --dir=internal/pkg/logger --output=internal/pkg/logger/mocks/ --all
	mockery --dir=internal/pkg/rest --output=internal/pkg/rest/mocks/ --all
	mockery --dir=internal/pkg/redis --output=internal/pkg/redis/mocks/ --all

mock-wrapper:
	mockery --dir=internal/app/wrapper/beeceptor --output=internal/app/wrapper/beeceptor/mocks/ --all

mock-repository:
	mockery --dir=internal/app/repository --output=internal/app/repository/mocks/ --all

mock-usecase:
	mockery --dir=internal/app/usecase/organization --output=internal/app/usecase/organization/mocks/ --all

mock: mock-pkg mock-wrapper mock-repository mock-usecase