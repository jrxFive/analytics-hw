FROM golang:1.15.4-alpine3.12 as go-mod-builder

WORKDIR /go/src/github.com/jrxFive/analytics-hw
COPY go.mod .
COPY go.sum .

RUN go mod download
COPY . .

FROM go-mod-builder as api-builder

WORKDIR /go/src/github.com/jrxFive/analytics-hw

RUN apk update \
    && apk add --no-cache --virtual build-base gcc libtool musl-dev \
    && go test -parallel 8 -race -cover ./... \
    && GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o analytics-hw cmd/api/main.go

FROM alpine:3.12

RUN apk update \
    && apk add --no-cache --virtual .build-deps ca-certificates \
    && rm -rf /var/cache/apk/*

WORKDIR /src/
COPY --from=api-builder /go/src/github.com/jrxFive/analytics-hw .
CMD ["./analytics-hw"]