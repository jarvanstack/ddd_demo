make_dir:=$(shell pwd)
app_name:=$(shell basename $(make_dir))

# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: gen build run

## init: init project, create dadabase import sql
.PHONY: init
init:
	cd job/init_mysql/ && go run . && cd $(make_dir)

## gen: Gen protobuf files.
.PHONY: gen
gen:
	cd internal/interfaces/rpc/proto_file/ && \
	protoc -I ./ --go_out=./ --go-grpc_out=./ ./in/* && \
	rm -rf ../protos && cp -r protos/ ../ && rm -rf protos && \
	cd $(make_dir)

## tidy: Tidy go mod.	
.PHONY: tidy
tidy:
	go mod tidy

## build: Build app
.PHONY: build
build:
	go build -o ./bin/$(app_name) -gcflags "-N -l" -race cmd/main.go

## run: Run app
.PHONY: run
run:
	./bin/$(app_name) --config ./config.yaml

## help: Show this help info.
help: Makefile
.PHONY: help
	@printf "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
