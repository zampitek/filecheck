OUTPUT_DIR := bin
BINARY_NAME := filecheck
PKG := github.com/zampitek/filecheck/version
DATE := $(shell date -u +%Y-%m-%dT%H.%M.%SZ)
COMMIT := $(shell git rev-parse --short HEAD)

build: linux mac windows

linux:
	@echo "Building for Linux..."
	@mkdir -p $(OUTPUT_DIR)
	@GOOS=linux GOARCH=amd64 go build -ldflags "-X $(PKG).Commit=$(COMMIT) -X '$(PKG).BuildDate=$(DATE)'" -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux_amd64 .
	@GOOS=linux GOARCH=arm64 go build -ldflags "-X $(PKG).Commit=$(COMMIT) -X '$(PKG).BuildDate=$(DATE)'" -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux_arm64 .

mac:
	@echo "Building for macOS..."
	@mkdir -p $(OUTPUT_DIR)
	@GOOS=darwin GOARCH=amd64 go build -ldflags "-X $(PKG).Commit=$(COMMIT) -X '$(PKG).BuildDate=$(DATE)'" -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux_amd64 .
	@GOOS=darwin GOARCH=arm64 go build -ldflags "-X $(PKG).Commit=$(COMMIT) -X '$(PKG).BuildDate=$(DATE)'" -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux_arm64 .

windows:
	@echo "Building for Windows..."
	@mkdir -p $(OUTPUT_DIR)
	@GOOS=windows GOARCH=amd64 go build -ldflags "-X $(PKG).Commit=$(COMMIT) -X '$(PKG).BuildDate=$(DATE)'" -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux_amd64 .
	@GOOS=windows GOARCH=arm64 go build -ldflags "-X $(PKG).Commit=$(COMMIT) -X '$(PKG).BuildDate=$(DATE)'" -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux_arm64 .

clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)/*

fmt:
	@echo "Formatting code..."
	@go fmt ./...

deps:
	@echo "Installing dependencies..."
	@go mod tidy
