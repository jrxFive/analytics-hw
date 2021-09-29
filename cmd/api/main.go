package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cespare/xxhash"

	"github.com/ReneKroon/ttlcache/v2"

	"github.com/jrxFive/analytics-hw/pkg/cache"

	cacheMiddleware "github.com/jrxFive/analytics-hw/third_party/echo-http-cache"

	"github.com/jrxFive/analytics-hw/types"

	openapi "github.com/jrxFive/analytics-hw/pkg/bitly"

	"github.com/jrxFive/analytics-hw/internal/settings"

	"github.com/jrxFive/analytics-hw/handlers"
	v1 "github.com/jrxFive/analytics-hw/handlers/v1"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	prefixV1 = "/api/v1"
)

func main() {
	e := echo.New()

	//settings
	s, err := settings.NewSettings()
	if err != nil {
		e.Logger.Fatal(fmt.Sprintf("err:%s, failed to instantiate settings, exiting", err.Error()))
		return
	}

	//middleware
	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		middleware.BodyLimit("256K"),
		middleware.CSRF(),
		middleware.RequestIDWithConfig(middleware.RequestIDConfig{
			Skipper: func(c echo.Context) bool {
				if strings.Contains(c.Path(), "healthz") {
					return false
				}

				return true
			},
			Generator: nil,
		}),
		middleware.Gzip(),
	)

	//bitly client
	bitlyClient := openapi.NewAPIClient(&openapi.Configuration{
		Servers: openapi.ServerConfigurations{
			{
				URL: s.BitlyServerURL,
			},
		},
		HTTPClient: &http.Client{
			Timeout: s.BitlyClientTimeout,
		},
	})

	//cache Middleware
	ttlCache := ttlcache.NewCache()
	if err := ttlCache.SetTTL(s.CacheSettings.TTL); err != nil {
		s.Logger.Fatal(err)
	}
	cacheImplForMiddleware := cache.NewInMemoryTTLCache(ttlCache, cache.NoopErrorLoggerFn)
	cacheMiddlewareImpl, err := cacheMiddleware.NewClient(
		cacheMiddleware.ClientWithAdapter(cacheImplForMiddleware),
		cacheMiddleware.ClientWithTTL(s.CacheSettings.TTL),
		cacheMiddleware.ClientWithHasher(func(c echo.Context) (string, error) {
			auth := c.Request().Header.Get(types.HeaderAuthorization)
			if auth == "" {
				return "", types.ErrMissingKeyAuth
			}

			cacheMiddleware.SortURLParams(c.Request().URL)

			xx := xxhash.New()
			for _, v := range []string{auth, c.Request().URL.String()} {
				if _, err := xx.Write([]byte(v)); err != nil {
					return "", err
				}
			}

			return fmt.Sprintf("%x", xx.Sum64()), nil
		}),
	)
	if err != nil {
		s.Logger.Fatal(err)
	}

	//healthz route
	healthz := handlers.NewHealthz()
	e.HEAD("/healthz", healthz.HeadHealthz)
	e.GET("/healthz", healthz.GetHealthz)

	//v1 routes
	analytics := v1.NewAnalytics(bitlyClient, s)
	analytics.Initialize(e.Group(fmt.Sprintf("%s/analytics", prefixV1),
		middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
			if len(key) == 0 {
				return false, types.ErrMissingKeyAuth
			}

			c.Set(types.ContextBitlyAPIKey, key)
			return true, nil
		}),
		cacheMiddlewareImpl.Middleware(),
	))

	for _, route := range e.Routes() {
		s.Logger.Info(route.Path, ",", route.Name, ",", route.Method)
	}

	//Server
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", s.ServerPort),
		ReadTimeout:       s.ServerReadTimeout,
		ReadHeaderTimeout: s.ServerReadTimeout,
		WriteTimeout:      s.ServerWriteTimeout,
		IdleTimeout:       s.ServerWriteTimeout,
	}

	e.Logger.Fatal(e.StartServer(server))
}
