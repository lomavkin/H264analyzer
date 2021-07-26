SRCS := $(shell find . -type f -name '*.go')

.PHONY: all
all: build

.PHONY: build
build: bin/H264analyzer

.PHONY: clean
clean:
	@rm -fr bin/

bin/H264analyzer: $(SRCS) go.mod
	go build -o bin/H264analyzer
