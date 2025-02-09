# Makefile for Go project

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=short-url

# Directories
SRC_DIR=./src
BIN_DIR=./bin

# Build the project
build:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)/...

# Run tests
test:
	$(GOTEST) ./...

# Clean build files
clean:
	$(GOCLEAN)
	rm -f $(BIN_DIR)/$(BINARY_NAME)

# Install dependencies
deps:
	$(GOGET) -v ./...

# Run the application
run:
	$(BIN_DIR)/$(BINARY_NAME)

.PHONY: build test clean deps run