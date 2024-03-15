FSPATH := $(PWD)
OS := $(shell uname -s)
ARCH := $(shell uname -m)
USER_HOME := $(shell echo ${HOME})
GOPATH := $(shell go env GOPATH)
BUILD_PATH := .build
RELEASE_PATH := .releases
OS_ARCH := linux/amd64 linux/arm64 windows/amd64
CLI_PATH := $(FSPATH)/cmd/cli
CLI_EXECUTABLE_NAME := ya
CLI_DAEMON_EMBED_PATH := $(CLI_PATH)/embed
DAEMON_GO_PATH := $(FSPATH)/cmd/daemon-go
DAEMON_EXECUTABLE_NAME := ya-daemon
VERSION_FILE := VERSION
VERSION := $(shell cat ./$(VERSION_FILE) || git describe --tags --abbrev=0)
LICENSE_FILE := $(FSPATH)/LICENSE
CLI_LD_FLAGS := -ldflags="-s -w -X '$(CLI_PATH)/main.version=$(VERSION)'"
DAEMON_LD_FLAGS := -ldflags="-s -w"