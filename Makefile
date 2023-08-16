#!/usr/bin/make

include .env
export

BUILD := ./build/webobs
SITES := `cat ./config/sites.yml`
SLACK := `cat ./config/slack.yml`
SOURCE := ./src
LDFLAGS := "-X 'main.YmlSiteConfig=$(SITES)' -X 'main.YmlSlackConfig=$(SLACK)'"

.PHONY : help init format build run test
.DEFAULT_GOAL : help

# HELP =================================================================================================================
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

init: ### Init project
	go mod download

format: ### Format go files
	gofmt -w $(SOURCE)

build: ### Build binary
	go build -ldflags $(LDFLAGS) -o $(BUILD) $(SOURCE)

run: ### Run project locally
	go run -ldflags $(LDFLAGS) $(SOURCE)

test: ### Run tests
	go test -v ./...