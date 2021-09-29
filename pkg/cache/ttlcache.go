package cache

import (
	"time"

	"github.com/pkg/errors"

	ttlCache "github.com/ReneKroon/ttlcache/v2"
)

var (
	ErrFetchedKeyValueNotRequiredType = errors.New("value is not required type")
)

type ErrorLoggerFn func(key string, err error)

var NoopErrorLoggerFn = func(key string, err error) {
	return
}

type InMemoryTTLCache struct {
	cache         ttlCache.SimpleCache
	errorLoggerFn func(key string, err error)
}

func NewDefault() InMemoryTTLCache {
	return InMemoryTTLCache{
		cache:         ttlCache.NewCache(),
		errorLoggerFn: NoopErrorLoggerFn,
	}
}

func NewInMemoryTTLCache(cache ttlCache.SimpleCache, errLoggerFn ErrorLoggerFn) InMemoryTTLCache {
	return InMemoryTTLCache{
		cache:         cache,
		errorLoggerFn: errLoggerFn,
	}
}

func (i InMemoryTTLCache) Get(key string) ([]byte, bool) {
	value, err := i.cache.Get(key)
	if err != nil {
		i.errorLoggerFn(key, err)
		return nil, false
	}

	checkedValue, ok := value.([]byte)
	if !ok {
		i.errorLoggerFn(key, errors.Wrapf(err, ErrFetchedKeyValueNotRequiredType.Error()))
		return nil, false
	}

	return checkedValue, true
}

func (i InMemoryTTLCache) Set(key string, response []byte, expiration time.Time) {
	duration := expiration.Sub(time.Now())

	if err := i.cache.SetWithTTL(key, response, duration); err != nil {
		i.errorLoggerFn(key, err)
	}
}

func (i InMemoryTTLCache) Release(key string) {
	if err := i.cache.Remove(key); err != nil {
		i.errorLoggerFn(key, err)
	}
}
