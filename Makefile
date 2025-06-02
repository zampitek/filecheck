# ----------------------------
# Configuration
# ----------------------------

# Program name
BINARY := filecheck

# Output directory
OUTPUT_DIR := bin

# Go package for version metadata (optional)
PKG := github.com/zampitek/filecheck/version

# Build metadata
DATE := $(shell date -u +%Y-%m-%dT%H.%M.%SZ)
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

# Installation
PREFIX ?= /usr/local
BINDIR := $(PREFIX)/bin

# ----------------------------
# Cross-compile targets
# ----------------------------

PLATFORMS := \
	linux_amd64 \
	linux_arm64 \
	darwin_amd64 \
	darwin_arm64 \
	windows_amd64 \
	windows_arm64

# ----------------------------
# Source files
# ----------------------------

GO_SOURCES := $(shell find . -name '*.go')

# ----------------------------
# Linker flags
# ----------------------------

LDFLAGS := -ldflags "-X '$(PKG).Commit=$(COMMIT)' -X '$(PKG).BuildDate=$(DATE)'"

# ----------------------------
# Phony targets
# ----------------------------

.PHONY: all build $(PLATFORMS) clean install uninstall help

# ----------------------------
# Default target
# ----------------------------

all: build

# ----------------------------
# Build targets
# ----------------------------

build: $(PLATFORMS)

$(PLATFORMS): %: $(GO_SOURCES)
	@echo "Building $(BINARY) for $*..."
	@mkdir -p $(OUTPUT_DIR)
	@GOOS=$(word 1,$(subst _, ,$*)) GOARCH=$(word 2,$(subst _, ,$*)) \
		go build $(LDFLAGS) -o $(OUTPUT_DIR)/$(BINARY)-$*

# ----------------------------
# Install target
# ----------------------------

install: $(OUTPUT_DIR)/$(BINARY)-linux_amd64
	@echo "Installing $(BINARY) to $(BINDIR)..."
	@mkdir -p $(BINDIR)
	@cp $(OUTPUT_DIR)/$(BINARY)-linux_amd64 $(BINDIR)/$(BINARY)
	@chmod 755 $(BINDIR)/$(BINARY)
	@echo "Installation complete."

# ----------------------------
# Uninstall target
# ----------------------------

uninstall:
	@echo "Uninstalling $(BINARY) from $(BINDIR)..."
	@rm -f $(BINDIR)/$(BINARY)
	@echo "Uninstallation complete."

# ----------------------------
# Clean target
# ----------------------------

clean:
	@echo "Cleaning up build artifacts..."
	@rm -rf $(OUTPUT_DIR)

# ----------------------------
# Help target
# ----------------------------

help:
	@echo "Available targets:"
	@echo "  all         Build all platforms"
	@echo "  build       Same as 'all'"
	@echo "  [platform]  Build a specific platform, e.g., make linux_amd64"
	@echo "              Available: $(PLATFORMS)"
	@echo "  install     Install the Linux amd64 binary to $(PREFIX)/bin"
	@echo "  uninstall   Remove the installed binary from $(PREFIX)/bin"
	@echo "  clean       Remove build artifacts"
	@echo "  help        Show this help message"
