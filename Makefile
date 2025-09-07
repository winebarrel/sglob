.PHONY: all
all: vet test

.PHONY: vet
vet:
	go vet ./...

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: bench
bench:
	go test -bench . -benchmem
