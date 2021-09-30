# analytics-hw

## Design
Attempt to provide aggergate metrics of given `<TOKEN>`. By default, obtain `<TOKEN>` user's default group and fetch
each bitlink in the group and metrics per country. Allow for parameters
such as unit, units, group to be overridden via query string parameters. 

### Client Generation
Upstream API provides an openapi.spec, allowing for generation of the client. This will automatically occur when using
Make, if it doesn't exist, as long the necessary dependencies are installed.

### Caching
Utilize an opinionated [caching middleware](third_party/echo-http-cache) which was forked to provide per user caching instead of global
caching. To also show understanding of interfaces, removed the included adapter implementations and provided a simple [TTLCache](pkg/cache) that can
 be extended to use more traditonal caching (memcached, redis). To avoid re-processing the same request as the upstream API does not support granularity below `unit=minute`. At
most the cache will return a max-age/Expiration of 60 seconds(which is configurable via environment variable).

### Pagination Channels/WaitGroup
Since the response for getting a group's bitlinks can potentially be paginated. A custom type [PaginatedGroupBitLinksRequest](handlers/request.go)
wraps around the request to add pagination support. Two implementation are provided Slice and Channel. For the Handler/Controller
we use the channel since it does not block to get all results; on each request all results are passed through a channel for the
WaitGroup to perform aggergation.

### What's Missing
Due to the current day of the given `units` these values can change. While the server has an opinionated
caching layer to the minute (dimension/granularity) it does not currently have pagination support.

This would be possible to include later on if adding additional workers to fetch results after the day has passed and stored. Allowing
this API to only need to query for the current day to perform its calculations.

Validation of request's, since the current aggregation does not take a request Body, this was not implemented. Generally
using a library such as [validator](https://github.com/go-playground/validator) to verify correctness via struct tags.

CircuitBreakers/Retry Decorators for upstream API requests

OpenAPI specification to generate a client.

### Server Routes
The server currently implements two routes, one for health checking
```
xh http://<IP>:<PORT>/api/v1/healthz
```

The second being an aggregation of a user's default group per link metrics by country. The token specified is the same
API token from the upstream dashboard/api. This should currently only be run locally as for the server is not currently
implemented TLS/configured for. If a bitlink has no clicks, it will not be included in the response.

The default parameters if none are given:
```
group=default #will use default from user associated with <TOKEN>
units=30
unit=day
```

```
xh http://<IP>:<PORT>/api/v1/analytics/countries Authoization:"Bearer <TOKEN>"

HTTP/1.1 200 OK
Cache-Control: private, max-age=55
Content-Type: application/json; charset=UTF-8
Date: Mon, 27 Sep 2021 15:49:07 GMT
Expires: Mon, 27 Sep 2021 11:49:55 EDT:
Set-Cookie: _csrf=IzNJtODtI3qEo6FG06BUKMTDGFp3lhRk; Expires=Tue, 28 Sep 2021 15:48:54 GMT
Vary: Cookie,Accept-Encoding

{
    "data": {
        "bit.ly/2XxugJT": {
            "metrics": [
                {
                    "value": "US",
                    "clicks": 1,
                    "averageClicksByDay": 0.03333333333333333
                }
            ]
        },
        "bit.ly/3Aqhgo0": {
            "metrics": [
                {
                    "value": "US",
                    "clicks": 2,
                    "averageClicksByDay": 0.06666666666666667
                }
            ]
        }
    },
    "group": "Bl9iekhPWD3",
    "unit": "day",
    "units": "30",
    "facet": "country"
}
``` 

### Settings
The following settings are configurable via environment variables. None are specifically required and will default
to using the production/internet facing upstream API if not specified.

```golang
//Prefix CACHE_
type CacheSettings struct {
	TTL time.Duration `env:"TTL,default=60s"`
}

//Prefix BITLY_
type BitlyClientSettings struct {
	BitlyServerURL              string        `env:"SERVER_URL,default=https://api-ssl.bitly.com/v4"`
	BitlyClientTimeout          time.Duration `env:"CLIENT_TIMEOUT,default=5s"`
}

//Prefix SERVER_
type ServerSettings struct {
	ServerIdleTimeout  time.Duration `env:"IDLE_TIMEOUT,default=30s"`
	ServerPort         int           `env:"PORT,default=8080"`
	ServerReadTimeout  time.Duration `env:"READ_TIMEOUT,default=1s"`
	ServerWriteTimeout time.Duration `env:"WRITE_TIMEOUT,default=30s"`
}
```

## Dependencies
```
go 1.15
golangci-lint (not required, used for make lint)
openapi-generator (brew install openapi-generator if on osx)
docker 20.10.17
```

## Server Dependencies
- Echo framework for routing and general ease of use for an HTTP API.
- TTLCache, for a caching implementation used as part of a caching middleware
- XXHash for cache lookups
- Go-Envconfig, to configure application through environment variables
- Zap, for logging
- Cast, to convert values
- Oauth2, is required from the openapi client

- Is, part of testing assertions

```
module github.com/jrxFive/analytics-hw

go 1.15

require (
	github.com/ReneKroon/ttlcache/v2 v2.8.1
	github.com/cespare/xxhash v1.1.0
	github.com/labstack/echo/v4 v4.5.0
	github.com/matryer/is v1.4.0
	github.com/pkg/errors v0.8.1
	github.com/sethvargo/go-envconfig v0.3.5
	github.com/spf13/cast v1.4.1
	go.uber.org/zap v1.19.1
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
)

```

## Build
To create a binary and run, this will also generate the client library based on the openapi [spec](internal/specs/bitly/v4.json).

```bash
make
```

## Test
If not running within a docker container.

```bash
make test
```

The [Dockerfile](Dockerfile) will automatically run tests as part of building the container image.
```bash
docker build -t analytics-hw
```

## Run
Binary:
```bash
make && ./analytics-hw
```

Container:
```
docker run analytics-hw
```