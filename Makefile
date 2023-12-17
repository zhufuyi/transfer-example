SHELL := /bin/bash

PROJECT_NAME := "transfer"
PKG := "$(PROJECT_NAME)"
PKG_LIST := $(shell go list ${PKG}/... | grep -v /vendor/ | grep -v /api/)



.PHONY: mod
# add missing and remove unused modules
mod:
	go mod tidy


.PHONY: fmt
# go format *.go files
fmt:
	gofmt -s -w .


.PHONY: test
# go test *_test.go files, the parameter -count=1 means that caching is disabled
test:
	go test -count=1 -short ${PKG_LIST}


.PHONY: cover
# generate test coverage
cover:
	go test -short -coverprofile=cover.out -covermode=atomic ${PKG_LIST}
	go tool cover -html=cover.out


.PHONY: proto
# generate *.go and template code by proto files, if you do not refer to the proto file, the default is all the proto files in the api directory. you can specify the proto file, multiple files are separated by commas, e.g. make proto FILES=api/user/v1/user.proto. only for ⓶ Microservices created based on sql, ⓷ Web services created based on protobuf, ⓸ Microservices created based on protobuf, ⓹ grpc gateway service created based on protobuf
proto: mod fmt
	@bash scripts/protoc.sh $(FILES)


.PHONY: build
# build transfer for linux amd64 binary
build:
	@echo "building 'transfer', linux binary file will output to 'cmd/transfer'"
	@cd cmd/transfer && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build


.PHONY: run
# run service
run:
	@bash scripts/run.sh


.PHONY: run-nohup
# run server with nohup to local, if you want to stop the server, pass the parameter stop, e.g. make run-nohup CMD=stop
run-nohup:
	@bash scripts/run-nohup.sh $(CMD)


.PHONY: update-config
# update internal/config code base on yaml file
update-config:
	@sponge config --server-dir=.


.PHONY: clean
# clean binary file, cover.out, template file
clean:
	@rm -vrf cmd/transfer/transfer
	@rm -vrf cover.out
	@rm -vrf main.go transfer.gv
	@rm -vrf internal/ecode/*.go.gen*
	@rm -vrf internal/routers/*.go.gen*
	@rm -vrf internal/handler/*.go.gen*
	@rm -vrf internal/service/*.go.gen*
	@rm -rf transfer-binary.tar.gz
	@echo "clean finished"


# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m  %-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := all
