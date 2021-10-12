CONTAINER_NAME=analytics-hw
BINARY_NAME=analytics-server
GOOS=darwin
GOARCH=amd64

.PHONY: build
build: fmt vet pkg/bitly
	go build -o $(BINARY_NAME) cmd/api/main.go

.PHONY: build.docker
build.docker:
	docker build -t $(BINARY_NAME) .

.PHONY: fmt
fmt:
	go $@ ./...

.PHONY: vet
vet:
	go $@ ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: fmt vet pkg/bitly
	go test -cover -race ./...

analytics-server: build

.PHONY: clean
clean:
	-rm $(BINARY_NAME)
	-rm -rf pkg/bitly

pkg/bitly:
	openapi-generator generate -i internal/specs/bitly/v4.json -g go  -o $@ --additional-properties=generateInterfaces=true
	rm $@/go.mod
	go mod tidy
