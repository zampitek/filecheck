BINARY_NAME = filecheck
OUTPUT_DIR = bin

all: linux

linux:
	@echo "Building for Linux..."
	@mkdir -p $(OUTPUT_DIR)
	@GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux_amd64 .
	@GOOS=linux GOARCH=arm64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-linux_arm64 .

mac:
	@echo "Building for macOS..."
	@mkdir -p $(OUTPUT_DIR)
	@GOOS=darwin GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-macOs_amd64 .
	@GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-macOs_arm64 .

windows:
	@echo "Building for Windows..."
	@mkdir -p $(OUTPUT_DIR)
	@GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-windows_amd64.exe .
	@GOOS=windows GOARCH=arm64 go build -o $(OUTPUT_DIR)/$(BINARY_NAME)-windows_arm64.exe .

clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)/*

fmt:
	@echo "Formatting code..."
	@go fmt ./...

deps:
	@echo "Installing dependencies..."
	@go mod tidy
