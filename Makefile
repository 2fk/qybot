APP?=bot
PROJECT?=github.com/2fk/qybot

VERSION?=$(shell git describe --tags --always)
COMMIT_HASH?=$(shell git rev-parse --short HEAD 2>/dev/null)
NOW?=$(shell date -u '+%Y-%m-%d %I:%M:%S %Z')

LDFLAGS += -X "${PROJECT}/bot.BuildAppName=${APP}"
LDFLAGS += -X "${PROJECT}/bot.BuildVersion=${VERSION}"
LDFLAGS += -X "${PROJECT}/bot.BuildTime=${NOW}"
LDFLAGS += -X "${PROJECT}/bot.BuildCommitHash=${COMMIT_HASH}"
LDFLAGS += -X "${PROJECT}/bot.BuildGoVersion=${shell go version}"
BUILD_FLAGS = "-v"
BUILD_TAGS = ""

default: build

.PHONY: build
build: govet
	CGO_ENABLED=0 GOOS= GOARCH= go build ${BUILD_FLAGS} -ldflags '-s -w ${LDFLAGS}' -tags '${BUILD_TAGS}' -o bin/${APP} cmd/*.go

.PHONY: govet
govet:
	@ go vet ./... && go fmt ./...

#build: clean
#	go build -o bin/bot cmd/*.go

clean:
	rm -rf bin/