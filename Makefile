BINARY_NAME = filecheck
OUTPUT_DIR = bin

all: build

build:
	@echo "Building $(BINARY_NAME)..."
	@mkdir -p $(OUTPUT_DIR)
	@go build -o $(OUTPUT_DIR)/$(BINARY_NAME) .

run: build
	@echo "Running $(BINARY_NAME)..."
	@./$(OUTPUT_DIR)/$(BINARY_NAME)

clean:
	@echo "Cleaning up..."
	@rm -rf $(OUTPUT_DIR)

fmt:
	@echo "Formatting code..."
	@go fmt ./...

deps:
	@echo "Installing dependencies..."
	@go mod tidy
