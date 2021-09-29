CONTAINER_NAME=analytics-hw
BINARY_NAME=analytics-server
GOOS=darwin
GOARCH=amd64

.PHONY: build
.PHONY: fmt
.PHONY: vet
.PHONY: build.docker
.PHONY: test
.PHONY: clean

build: fmt vet pkg/bitly
	go build -o $(BINARY_NAME) cmd/api/main.go

build.docker:
	docker build -t $(BINARY_NAME) .

fmt:
	go $@ ./...

vet:
	go $@ ./...


test: fmt vet pkg/bitly
	go test -cover -race ./...

analytics-server: build

clean:
	-rm $(BINARY_NAME)
	-rm -rf pkg/bitly

pkg/bitly:
	openapi-generator generate -i internal/specs/bitly/v4.json -g go  -o $@ --additional-properties=generateInterfaces=true
	rm $@/go.mod
	go mod tidy
