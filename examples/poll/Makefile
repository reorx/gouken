all: build

# Dev
run:
	go run cmd/rpcserver/main.go

.PHONY: proto
proto:
	protoc ./proto/*.proto --go_out=plugins=grpc:.

# Test
test: test-candidate

TESTBIN := go test -v
export PROJECTPATH := $(shell pwd)

test-candidate:
	@echo
	$(TESTBIN) ./test/candidate

# Build
MAIN := ./cmd/rpcserver
BUILDNAME := poll-app
build: clean
	@mkdir -p build/osx build/linux
	@echo build osx
	go build -o build/osx/$(BUILDNAME) $(MAIN)
	@echo build linux
	GOOS=linux GOARCH=amd64 go build -o build/linux/$(BUILDNAME) $(MAIN)

clean:
	@echo clean
	@rm -rf build
