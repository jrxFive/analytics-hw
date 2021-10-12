package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/matryer/is"
)

func TestHealthz_HeadHealthz(t *testing.T) {
	var tests = []struct {
		name string
		req  *http.Request
		want func(rec *httptest.ResponseRecorder) bool
	}{
		{
			name: "200 OK",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodHead, fmt.Sprintf("/"), nil)

				return req
			}(),
			want: func(rec *httptest.ResponseRecorder) bool {
				is := is.New(t)
				is.Equal(http.StatusNoContent, rec.Code)

				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			e := echo.New()
			h := Healthz{}

			e.HEAD("/", h.HeadHealthz)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, tt.req)
			defer e.Close()

			if !tt.want(rec) {
				is.Fail()
			}
		})
	}
}

func TestHealthz_GetHealthz(t *testing.T) {
	var tests = []struct {
		name string
		req  *http.Request
		want func(rec *httptest.ResponseRecorder) bool
	}{
		{
			name: "200 OK",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/"), nil)

				return req
			}(),
			want: func(rec *httptest.ResponseRecorder) bool {
				is := is.New(t)
				is.Equal(http.StatusOK, rec.Code)

				return true
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			is := is.New(t)

			e := echo.New()
			h := Healthz{}

			e.GET("/", h.GetHealthz)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, tt.req)
			defer e.Close()

			if !tt.want(rec) {
				is.Fail()
			}
		})
	}
}
