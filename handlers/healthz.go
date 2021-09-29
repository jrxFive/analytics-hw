package handlers

import (
	"net/http"

	"github.com/jrxFive/analytics-hw/schemas"
	"github.com/labstack/echo/v4"
)

type Healthz struct{}

func NewHealthz() Healthz {
	return Healthz{}
}

func (h Healthz) HeadHealthz(c echo.Context) error {
	return c.NoContent(http.StatusNoContent)
}

func (h Healthz) GetHealthz(c echo.Context) error {
	response := schemas.HealthzResponse{Status: "OK"}
	return c.JSONPretty(http.StatusOK, response, " ")
}
