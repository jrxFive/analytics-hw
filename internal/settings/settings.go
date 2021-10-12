package settings

import (
	"context"
	"sync"
	"time"

	"github.com/jrxFive/analytics-hw/types"
	"github.com/pkg/errors"

	"github.com/sethvargo/go-envconfig"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	envPrefixBitly  = "BITLY_"
	envPrefixCache  = "CACHE_"
	envPrefixServer = "SERVER_"
)

var (
	once   sync.Once
	logger *zap.SugaredLogger
)

type Settings struct {
	EnvSettings
	Logger *zap.SugaredLogger
}

type EnvSettings struct {
	BitlyClientSettings
	CacheSettings
	ServerSettings
}

type CacheSettings struct {
	TTL time.Duration `env:"TTL,default=60s"`
}

type BitlyClientSettings struct {
	BitlyServerURL     string        `env:"SERVER_URL,default=https://api-ssl.bitly.com/v4"`
	BitlyClientTimeout time.Duration `env:"CLIENT_TIMEOUT,default=5s"`
}

type ServerSettings struct {
	ServerIdleTimeout  time.Duration `env:"IDLE_TIMEOUT,default=30s"`
	ServerPort         int           `env:"PORT,default=8080"`
	ServerReadTimeout  time.Duration `env:"READ_TIMEOUT,default=1s"`
	ServerWriteTimeout time.Duration `env:"WRITE_TIMEOUT,default=30s"`
}

func NewSettings() (Settings, error) {
	var (
		bitlySettings  BitlyClientSettings
		cacheSettings  CacheSettings
		serverSettings ServerSettings
	)

	serverLooker := envconfig.PrefixLookuper(envPrefixServer, envconfig.OsLookuper())
	if err := envconfig.ProcessWith(context.Background(), &serverSettings, serverLooker); err != nil {
		return Settings{}, errors.Wrap(err, types.ErrRequiredSetting.Error())
	}

	bitlyLooker := envconfig.PrefixLookuper(envPrefixBitly, envconfig.OsLookuper())
	if err := envconfig.ProcessWith(context.Background(), &bitlySettings, bitlyLooker); err != nil {
		return Settings{}, errors.Wrap(err, types.ErrRequiredSetting.Error())
	}

	cacheLooker := envconfig.PrefixLookuper(envPrefixCache, envconfig.OsLookuper())
	if err := envconfig.ProcessWith(context.Background(), &cacheSettings, cacheLooker); err != nil {
		return Settings{}, errors.Wrap(err, types.ErrRequiredSetting.Error())
	}

	return Settings{
		EnvSettings: EnvSettings{
			BitlyClientSettings: bitlySettings,
			CacheSettings:       cacheSettings,
			ServerSettings:      serverSettings,
		},
		Logger: getLogger(),
	}, nil
}

func getLogger() *zap.SugaredLogger {
	once.Do(func() {
		logConfig := zap.NewDevelopmentConfig()
		logConfig.OutputPaths = []string{"stdout"}
		logConfig.ErrorOutputPaths = []string{"stdout"}
		logConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

		logBuilder, _ := logConfig.Build()
		logger = logBuilder.Sugar()
	})

	return logger
}
