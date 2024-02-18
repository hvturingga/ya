# -*- coding: utf-8 -*-
FSPATH := $(PWD)

OS := $(shell uname -s)
ARCH := $(shell uname -m)
USER_HOME := $(shell echo ${HOME})
GOPATH := $(shell go env GOPATH)

# build
BUILD_PATH := .build
RELEASE_PATH := .releases
OS_ARCH := linux/amd64 linux/arm64 windows/amd64

# cli
CLI_PATH := $(FSPATH)/cmd/cli
CLI_EXECUTABLE_NAME := ya
CLI_DAEMON_EMBED_PATH := $(CLI_PATH)/embed

# daemon
DAEMON_GO_PATH := $(FSPATH)/cmd/daemon-go
DAEMON_EXECUTABLE_NAME := ya-daemon

# version
VERSION_FILE := VERSION
VERSION := $(shell cat ./$(VERSION_FILE))

ifeq ($(VERSION),)
        VERSION := $(shell git describe --tags --abbrev=0)
endif

LICENSE_FILE := $(FSPATH)/LICENSE

# ld flags
CLI_LD_FLAGS := -ldflags="-s -w -X '$(CLI_PATH)/main.version=$(VERSION)'"
DAEMON_LD_FLAGS := -ldflags="-s -w"

# mk
MAKE_FSPATH := $(FSPATH)/hack/mk
MK_FILES := $(wildcard $(MAKE_FSPATH)/*.mk)

.PHONY: print build gen release pre-release clean release-all

print:
	@echo
	@echo "FSPATH: $(FSPATH)"
	@echo "VERSION: $(VERSION)"
	@echo ""
	@echo "OS: $(OS)"
	@echo "ARCH: $(ARCH)"
	@echo "GOPATH: $(GOPATH)"
	@echo "USER_HOME: $(USER_HOME)"
	@echo ""

include $(MK_FILES)

clean:
	@rm -rf $(BUILD_PATH)
	@rm -rf $(RELEASE_PATH)

pre-release:
	@make gen

release: clean pre-release release-all package-cli

release-all: build-daemon release-cli

CLI_TARGETS := $(foreach osarch,$(OS_ARCH),release-cli-$(osarch))
DAEMON_TARGETS := $(foreach osarch,$(OS_ARCH),build-daemon-$(osarch))

.PHONY: release-cli build-daemon $(CLI_TARGETS) $(DAEMON_TARGETS)

release-cli: $(CLI_TARGETS)

build-daemon: $(DAEMON_TARGETS)

$(CLI_TARGETS): release-cli-%:
	$(eval OS=$(word 1,$(subst /, ,$*)))
	$(eval ARCH=$(word 2,$(subst /, ,$*)))
	$(eval EXT=$(if $(filter windows,$(OS)),.exe,))
	$(eval CC=$(if $(filter arm64,$(ARCH)),aarch64-linux-gnu-gcc,$(if $(filter windows,$(OS)),x86_64-w64-mingw32-gcc,)))
	CGO_ENABLED=1 GOOS=$(OS) GOARCH=$(ARCH) CC=$(CC) go build $(CLI_LD_FLAGS) -o $(BUILD_PATH)/ya-$(OS)-$(ARCH)$(EXT) $(CLI_PATH) && upx -9 $(BUILD_PATH)/ya-$(OS)-$(ARCH)$(EXT)

$(DAEMON_TARGETS): build-daemon-%:
	$(eval OS=$(word 1,$(subst /, ,$*)))
	$(eval ARCH=$(word 2,$(subst /, ,$*)))
	$(eval EXT=$(if $(filter windows,$(OS)),.exe,))
	$(eval CC=$(if $(filter arm64,$(ARCH)),aarch64-linux-gnu-gcc,$(if $(filter windows,$(OS)),x86_64-w64-mingw32-gcc,)))
	@mkdir -p $(CLI_DAEMON_EMBED_PATH)
	CGO_ENABLED=1 GOOS=$(OS) GOARCH=$(ARCH) CC=$(CC) go build $(DAEMON_LD_FLAGS) -o $(CLI_DAEMON_EMBED_PATH)/ya-daemon-$(OS)-$(ARCH)$(EXT) $(DAEMON_GO_PATH) && upx -9 $(CLI_DAEMON_EMBED_PATH)/ya-daemon-$(OS)-$(ARCH)$(EXT)

package-cli:
	@mkdir -p $(RELEASE_PATH)
	$(foreach osarch,$(OS_ARCH),$(call package,$(osarch));)

define package
	$(eval OS=$(word 1,$(subst /, ,$1)))
	$(eval ARCH=$(word 2,$(subst /, ,$1)))
	$(eval EXT=$(if $(filter windows,$(OS)),.exe,))
	$(eval TARGET=ya-$(OS)-$(ARCH)$(EXT))
	$(eval PACKAGE_NAME=ya-$(VERSION)-$(OS)-$(ARCH))
	$(eval PACKAGE_PATH=$(RELEASE_PATH)/$(PACKAGE_NAME))
	@if [ "$(OS)" = "windows" ]; then \
		mkdir -p $(PACKAGE_PATH) && \
		cp $(BUILD_PATH)/$(TARGET) $(PACKAGE_PATH) && \
		cp $(LICENSE_FILE) $(PACKAGE_PATH) && \
		cd $(RELEASE_PATH) && \
		zip -r $(PACKAGE_NAME).zip $(PACKAGE_NAME) && \
		cd $(FSPATH); \
	else \
		mkdir -p $(PACKAGE_PATH) && \
		cp $(BUILD_PATH)/$(TARGET) $(PACKAGE_PATH) && \
		cp $(LICENSE_FILE) $(PACKAGE_PATH) && \
		cd $(RELEASE_PATH) && \
		tar -czvf $(PACKAGE_NAME).tar.gz $(PACKAGE_NAME) && \
		cd $(FSPATH); \
	fi
endef