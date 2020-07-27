NAME    := postmaster
PACKAGE := github.com/biosugar0/$(NAME)
GIT     := $(shell git rev-parse --short HEAD)
SOURCE_DATE_EPOCH ?= $(shell date +%s)
DATE    := $(shell date -u -d @${SOURCE_DATE_EPOCH} +%FT%T%Z)
VERSION  ?= v0.1.0
IMG_NAME := biosugar0/postmaster
IMAGE    := ${IMG_NAME}:${VERSION}

build:  ## Builds the CLI
	@go build \
	-ldflags "-w -s -X ${PACKAGE}/cmd.version=${VERSION} -X ${PACKAGE}/cmd.commit=${GIT} -X ${PACKAGE}/cmd.date=${DATE}" \
	-a -tags netgo -o execs/${NAME} main.go
