# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GORUN=$(GOCMD) run

BUILD_FILES = $(shell go list -f '{{range .GoFiles}}{{$$.Dir}}/{{.}} {{end}}' ./...)

LDFLAGS = -s -w

bin/rename: $(BUILD_FILES)
	$(GOBUILD) -trimpath -o "$@" -ldflags='$(LDFLAGS)' ./main.go 
build: bin/rename
.PHONY: build

test: $(BUILD_FILES)
	$(GOTEST) ./...  -coverprofile c.out
.PHONY: test

clean:
	rm -rf bin dist c.out
.PHONY: clean