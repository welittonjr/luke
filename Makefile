BINARY_NAME=luke
BIN_DIR=bin

.PHONY: build test lint clean run-example

build:
	go build -o $(BIN_DIR)/$(BINARY_NAME) ./cmd/...

test:
	go test ./...

lint:
	go vet ./...

clean:
	rm -rf $(BIN_DIR)

run-example:
	./$(BIN_DIR)/$(BINARY_NAME) run "diga olá"
