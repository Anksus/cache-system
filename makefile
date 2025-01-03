# Makefile for building and running a Go application

# Variables

APP_NAME = cache
SRC_DIR = .
BUILD_DIR = bin
GO_FILES = $(SRC_DIR)/cmd/main.go

.PHONY: build
build:
	@echo "Building the application..."
	go build -o $(BUILD_DIR)/$(APP_NAME) $(GO_FILES)


.PHONY: run
run: build
	@echo "Running the application..."
	@$(BUILD_DIR)/$(APP_NAME) start
