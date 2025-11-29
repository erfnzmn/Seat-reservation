APP_NAME ?= seat-reservation
CMD_PATH := ./cmd/server
BIN_DIR ?= bin
BIN := $(BIN_DIR)/$(APP_NAME)

.PHONY: build run test tidy fmt clean

build:
@echo "Building $(BIN)..."
@mkdir -p $(BIN_DIR)
go build -o $(BIN) $(CMD_PATH)

run:
go run $(CMD_PATH)

test:
go test ./...

tidy:
go mod tidy

fmt:
gofmt -w $(shell find . -name '*.go' -not -path './vendor/*')

clean:
rm -rf $(BIN_DIR)