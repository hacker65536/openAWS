SHELL := /bin/bash
GO := GO111MODULE=on go
REV := $(shell git rev-parse --short HEAD 2> /dev/null || echo 'unknown')
GO_VERSION := $(shell $(GO) version | sed -e 's/^[^0-9.]*\([0-9.]*\).*/\1/')
GO_DEPENDENCIES := $(call rwildcard,pkg/,*.go) $(call rwildcard,cmd/,*.go)

BRANCH     := $(shell git rev-parse --abbrev-ref HEAD 2> /dev/null  || echo 'unknown')
BUILD_DATE := $(shell date +%Y%m%d-%H:%M:%S)
CGO_ENABLED = 0


GOTEST := $(GO) test
VERSION ?= $(shell echo "$$(git for-each-ref refs/tags/ --count=1 --sort=-version:refname --format='%(refname:short)' 2>/dev/null)-dev+$(REV)" | sed 's/^v//')
DIRNAME := $(shell basename $(shell pwd))
BUILD_TIME_CONFIG_FLAGS ?= ""

BUILDFLAGS :=  -ldflags \
			   "-X github.com/hacker65536/${DIRNAME}/cmd.GitCommit=$(REV) \
			    -X github.com/hacker65536/${DIRNAME}/cmd.Version=$(VERSION) \
				$(BUILD_TIME_CONFIG_FLAGS)"

				

version:
	@echo ${VERSION}
install: $(GO_DEPENDENCIES) ## Install the binary
	GOBIN=${GOPATH}/bin $(GO) install $(BUILDFLAGS) 


# https://sourcegraph.com/github.com/jenkins-x/jx/-/blob/Makefile?L17
