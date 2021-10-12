package middleware

import (
	"fmt"
	"strings"

	"github.com/cespare/xxhash"
	cacheMiddleware "github.com/jrxFive/analytics-hw/third_party/echo-http-cache"
	"github.com/jrxFive/analytics-hw/types"
	"github.com/labstack/echo/v4"
)

var (
	// CacheHasher uses echo.Context to create a hash key to check if response is in cache.
	// Uses XXHash as hashing algorithm
	CacheHasher = func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(echo.HeaderAuthorization)
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
	}

	// KeyAuthContext verifies that auth key has been supplied and adds to context
	KeyAuthContext = func(key string, c echo.Context) (bool, error) {
		if key == "" {
			return false, types.ErrMissingKeyAuth
		}

		c.Set(types.ContextBitlyAPIKey, key)
		return true, nil
	}

	// RequestIDSkipper skip adding RequestID if path is to healthz
	RequestIDSkipper = func(c echo.Context) bool {
		return strings.Contains(c.Path(), "healthz")
	}
)
