# combined dev build / test and package for deploy Makefile

DOCKER_IMAGES_BASE_NAME=evosalepos

# Image name to build the docker container from
IMAGE=$(DOCKER_IMAGES_BASE_NAME)/gomagik

# This Makefile is in APP_ROOT and called with APP_ROOT as pwd so:
APP_ROOT=$(shell pwd)

# App is always named after the directory it is in
NAME=$(notdir $(shell pwd))

# App are always in ROOT/src/APP_ROOT so:
ROOT=$(APP_ROOT)/../..

# Image name to generate from compiled src
IMAGE_BASE_NAME=$(DOCKER_IMAGES_BASE_NAME)/$(NAME)

# location ROOT is mapped to inside the docker container
LOCAL_ROOT=/go

# start a docker container with all the dev tools and this mono repo mapped to LOCAL_ROOT
DOCKER_CMD_TO_RUN=docker run --rm=true
DOCKER_CMD_TO_RUN+= -v $(ROOT):$(LOCAL_ROOT)
DOCKER_CMD_TO_RUN+= $(IMAGE)

# grab GIT info for use in version tags etc.
COMMIT_HASH=$(shell git rev-parse HEAD)
COMMIT_TIME=$(shell git show -s --format="%cI")

shell: gofmt
	docker run --rm=true -it -v $(ROOT):$(LOCAL_ROOT) $(IMAGE)

compile: gofmt build

package: compile
	mkdir -p ./bin
	cp ./Dockerfile ./bin
	cp ./tmp/$(NAME) ./bin
	-cp -r ./assets/* ./bin/
	docker build --rm=true -t $(DEPLOY_TAG) ./bin

build:
	$(DOCKER_CMD_TO_RUN) sh -c "cd /go && rm -rf /go/src/$(NAME)/bin && GOPATH=/go/_vendor /usr/local/go/bin/go build -ldflags '-X main.BUILD=$(COMMIT_TIME)__$(COMMIT_HASH)'  -o  /go/src/$(NAME)/tmp/$(NAME) /go/src/$(NAME)/*.go"

gofmt:
	$(DOCKER_CMD_TO_RUN) sh -c "cd /go && rm -rf /go/$(NAME)/bin && /usr/local/go/bin/gofmt -w /go/src/$(NAME)/*.go"
