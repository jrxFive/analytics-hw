package v1

import (
	"io/ioutil"

	"github.com/jrxFive/analytics-hw/types"

	"github.com/labstack/echo/v4"

	"github.com/jrxFive/analytics-hw/internal/settings"
	"go.uber.org/zap"
)

const (
	ExpectedUsername         = "jrxfive"
	ExpectedDefaultGroupGUID = "Bl9iekhPWD3"
)

func FakeKeyAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(types.ContextBitlyAPIKey, "")
		return next(c)
	}
}

var (
	NoopLogger  = settings.Settings{Logger: zap.NewNop().Sugar()}
	DebugLogger = settings.Settings{Logger: zap.NewExample().Sugar()}
)

func referenceString(s string) *string {
	return &s
}

func fixture(path string) string {
	b, err := ioutil.ReadFile("testdata/" + path)
	if err != nil {
		panic(err)
	}
	return string(b)
}
