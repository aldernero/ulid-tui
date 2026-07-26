BINARY := ulid-tui

.PHONY: build
build:
	go build -o $(BINARY) .

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	rm -f $(BINARY)
