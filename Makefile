make_dir:=$(shell pwd)
app_name:=$(shell basename $(make_dir))

# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: gen build run

## gen: Gen protobuf files.
gen:
	cd internal/interfaces/rpc/proto_file/ && \
	protoc -I ./ --go_out=./ --go-grpc_out=./ ./in/* && \
	rm -rf ../protos && cp -r protos/ ../ && rm -rf protos && \
	cd $(make_dir)

## tidy: Tidy go mod.	
tidy:
	go mod tidy
## build: Build app
build:
	go build -o ./bin/$(app_name) -gcflags "-N -l" -race cmd/main.go

## run: Run app
run:
	./bin/$(app_name) --config ./config.yaml

## help: Show this help info.
help: Makefile
.PHONY: help
	@printf "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:\n"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"
