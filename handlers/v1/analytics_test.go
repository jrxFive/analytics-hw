package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"sync"
	"testing"

	"github.com/labstack/echo/v4"
	is "github.com/matryer/is"

	"github.com/jrxFive/analytics-hw/internal/settings"
	openapi "github.com/jrxFive/analytics-hw/pkg/bitly"
	schemas "github.com/jrxFive/analytics-hw/schemas/v1"
)

func TestAnalytics_getQsParamOrDefault(t *testing.T) {
	type fields struct {
		client   *openapi.APIClient
		settings settings.Settings
	}
	type args struct {
		param string
		c     echo.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "group value set",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				param: QsGroup,
				c: func() echo.Context {
					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					rec := httptest.NewRecorder()

					c := e.NewContext(req, rec)
					c.SetParamNames(QsGroup)
					c.SetParamValues(ExpectedDefaultGroupGUID)
					return c
				}(),
			},
			want: ExpectedDefaultGroupGUID,
		},
		{
			name: "group value default",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				param: QsGroup,
				c: func() echo.Context {
					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					rec := httptest.NewRecorder()

					c := e.NewContext(req, rec)
					return c
				}(),
			},
			want: DefaultQsParams[QsGroup],
		},
		{
			name: "unit value set",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				param: QsUnit,
				c: func() echo.Context {
					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					rec := httptest.NewRecorder()

					c := e.NewContext(req, rec)
					c.SetParamNames(QsUnit)
					c.SetParamValues("hour")
					return c
				}(),
			},
			want: "hour",
		},
		{
			name: "unit value default",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				param: QsUnit,
				c: func() echo.Context {
					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					rec := httptest.NewRecorder()

					c := e.NewContext(req, rec)
					return c
				}(),
			},
			want: DefaultQsParams[QsUnit],
		},
		{
			name: "units value set",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				param: QsUnits,
				c: func() echo.Context {
					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					rec := httptest.NewRecorder()

					c := e.NewContext(req, rec)
					c.SetParamNames(QsUnits)
					c.SetParamValues("5")
					return c
				}(),
			},
			want: "5",
		},
		{
			name: "units value default",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				param: QsUnits,
				c: func() echo.Context {
					e := echo.New()
					req := httptest.NewRequest(http.MethodGet, "/", nil)
					rec := httptest.NewRecorder()

					c := e.NewContext(req, rec)
					return c
				}(),
			},
			want: DefaultQsParams[QsUnits],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Analytics{
				client:   tt.fields.client,
				settings: tt.fields.settings,
			}
			if got := a.getQsParamOrDefault(tt.args.param, tt.args.c); got != tt.want {
				t.Errorf("getQsParamOrDefault() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAnalytics_fetchUser(t *testing.T) {
	type fields struct {
		client   *openapi.APIClient
		settings settings.Settings
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		server  *httptest.Server
		want    func(u openapi.User, err error, wantErr bool) bool
		wantErr bool
	}{
		{
			name: "found user",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{ctx: context.TODO()},
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)

				m.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("user.json")))
				})

				return s
			}(),
			want: func(u openapi.User, err error, wantErr bool) bool {
				is := is.New(t)

				is.True(err == nil)
				is.Equal(ExpectedUsername, u.Name)
				is.Equal(ExpectedDefaultGroupGUID, *u.DefaultGroupGuid)

				return true
			},
			wantErr: false,
		},
		{
			name: "forbidden on fetch user",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{ctx: context.TODO()},
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)

				m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusForbidden)
					w.Write([]byte(fixture("forbidden.json")))
				})

				return s
			}(),
			want: func(u openapi.User, err error, wantErr bool) bool {
				is := is.New(t)

				is.True(err != nil)
				is.True(wantErr)
				return true
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()

			tt.fields.client = openapi.NewAPIClient(&openapi.Configuration{
				Servers: openapi.ServerConfigurations{
					{
						URL: tt.server.URL,
					},
				},
			})

			a := Analytics{
				client:   tt.fields.client,
				settings: tt.fields.settings,
			}

			got, err := a.fetchUser(tt.args.ctx)
			tt.want(got, err, tt.wantErr)
		})
	}
}

func TestAnalytics_getGroupBitLinksWithChannel(t *testing.T) {
	type fields struct {
		client   *openapi.APIClient
		settings settings.Settings
	}
	type args struct {
		ctx     context.Context
		groupId string
		ch      chan PaginatedGroupBitLinksResponse
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		server  *httptest.Server
		want    func(ch chan PaginatedGroupBitLinksResponse, wantErr bool) bool
		wantErr bool
	}{
		{
			name: "paginated results",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				ctx:     context.TODO(),
				groupId: ExpectedDefaultGroupGUID,
				ch:      make(chan PaginatedGroupBitLinksResponse, 10),
			},
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)
				mu := sync.RWMutex{}

				counter := 0

				m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)

					mu.RLock()
					if counter == 0 {
						w.Write([]byte(fixture("three-link-first-pagination.json")))
					} else if counter == 1 {
						w.Write([]byte(fixture("three-link-second-pagination.json")))
					} else if counter == 2 {
						w.Write([]byte(fixture("three-link-third-pagination.json")))
					} else {
						t.Log("counter should not reach above two call")
						t.Fail()
					}
					mu.RUnlock()

					mu.Lock()
					counter += 1
					mu.Unlock()
				})

				return s
			}(),
			want: func(ch chan PaginatedGroupBitLinksResponse, wantErr bool) bool {
				is := is.New(t)
				is.Equal(3, len(ch))

				r := <-ch
				is.Equal("bit.ly/3nPhiCD", *r.Bitlink.Id)

				r = <-ch
				is.Equal("bit.ly/2XxugJT", *r.Bitlink.Id)

				r = <-ch
				is.Equal("bit.ly/3Aqhgo0", *r.Bitlink.Id)

				return true
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()

			a := Analytics{
				settings: tt.fields.settings,
			}

			a.client = openapi.NewAPIClient(&openapi.Configuration{
				Servers: openapi.ServerConfigurations{
					{
						URL: tt.server.URL,
					},
				},
			})

			a.getGroupBitLinksWithChannel(tt.args.ctx, tt.args.groupId, tt.args.ch)
			tt.want(tt.args.ch, tt.wantErr)
		})
	}
}

func TestAnalytics_getGroupBitLinkCountryMetricsFromChannel(t *testing.T) {
	type fields struct {
		client   *openapi.APIClient
		settings settings.Settings
	}
	type args struct {
		ctx          context.Context
		timeUnit     openapi.TimeUnit
		units        string
		linksChannel chan PaginatedGroupBitLinksResponse
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		server      *httptest.Server
		want        map[string]schemas.AnalyticMetrics
		wantErrChan bool
	}{
		{
			name: "no clicks bitlink",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				ctx:      context.TODO(),
				timeUnit: openapi.DAY,
				units:    "30",
				linksChannel: func() chan PaginatedGroupBitLinksResponse {
					ch := make(chan PaginatedGroupBitLinksResponse, 10)

					ch <- PaginatedGroupBitLinksResponse{
						Bitlink: openapi.BitlinkBody{
							Id: referenceString("bit.ly/3nPhiCD"),
						},
						Response: nil,
						Err:      nil,
					}

					close(ch)
					return ch
				}(),
			},

			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)

				m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("bitlink3-click-metrics-no-clicks.json")))
				})

				return s
			}(),
			want:        map[string]schemas.AnalyticMetrics{},
			wantErrChan: false,
		},
		{
			name: "one country clicks bitlink",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				ctx:      context.TODO(),
				timeUnit: openapi.DAY,
				units:    "30",
				linksChannel: func() chan PaginatedGroupBitLinksResponse {
					ch := make(chan PaginatedGroupBitLinksResponse, 10)

					ch <- PaginatedGroupBitLinksResponse{
						Bitlink: openapi.BitlinkBody{
							Id: referenceString("bit.ly/3nPhiCD"),
						},
						Response: nil,
						Err:      nil,
					}

					close(ch)
					return ch
				}(),
			},
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)

				m.HandleFunc("/bitlinks/bit.ly/3nPhiCD/countries", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("bitlink1-click-metrics.json")))

				})

				return s
			}(),
			want: map[string]schemas.AnalyticMetrics{
				"bit.ly/3nPhiCD": {
					Metrics: []schemas.AnalyticMetric{
						{
							Value:              "US",
							Clicks:             2,
							AverageClicksByDay: 0.06666666666666667,
						},
					},
				},
			},
			wantErrChan: false,
		},
		{
			name: "multiple country clicks bitlink",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				ctx:      context.TODO(),
				timeUnit: openapi.DAY,
				units:    "30",
				linksChannel: func() chan PaginatedGroupBitLinksResponse {
					ch := make(chan PaginatedGroupBitLinksResponse, 10)

					ch <- PaginatedGroupBitLinksResponse{
						Bitlink: openapi.BitlinkBody{
							Id: referenceString("bit.ly/3Aqhgo0"),
						},
						Response: nil,
						Err:      nil,
					}

					close(ch)
					return ch
				}(),
			},
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)

				m.HandleFunc("/bitlinks/bit.ly/3Aqhgo0/countries", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("bitlink2-click-metrics-multiple-countries.json")))
				})

				return s
			}(),
			want: map[string]schemas.AnalyticMetrics{
				"bit.ly/3Aqhgo0": {
					Metrics: []schemas.AnalyticMetric{
						{
							Value:              "US",
							Clicks:             2,
							AverageClicksByDay: 0.06666666666666667,
						},
						{
							Value:              "UK",
							Clicks:             3,
							AverageClicksByDay: 0.1,
						},
					},
				},
			},
			wantErrChan: false,
		},
		{
			name: "multiple country clicks multiple bitlinks",
			fields: fields{
				settings: NoopLogger,
			},
			args: args{
				ctx:      context.TODO(),
				timeUnit: openapi.DAY,
				units:    "30",
				linksChannel: func() chan PaginatedGroupBitLinksResponse {
					ch := make(chan PaginatedGroupBitLinksResponse, 10)

					ch <- PaginatedGroupBitLinksResponse{
						Bitlink: openapi.BitlinkBody{
							Id: referenceString("bit.ly/3Aqhgo0"),
						},
						Response: nil,
						Err:      nil,
					}

					ch <- PaginatedGroupBitLinksResponse{
						Bitlink: openapi.BitlinkBody{
							Id: referenceString("bit.ly/3nPhiCD"),
						},
						Response: nil,
						Err:      nil,
					}

					close(ch)
					return ch
				}(),
			},
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)

				m.HandleFunc("/bitlinks/bit.ly/3Aqhgo0/countries", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("bitlink2-click-metrics-multiple-countries.json")))
				})

				m.HandleFunc("/bitlinks/bit.ly/3nPhiCD/countries", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("bitlink1-click-metrics.json")))
				})

				return s
			}(),
			want: map[string]schemas.AnalyticMetrics{
				"bit.ly/3Aqhgo0": {
					Metrics: []schemas.AnalyticMetric{
						{
							Value:              "US",
							Clicks:             2,
							AverageClicksByDay: 0.06666666666666667,
						},
						{
							Value:              "UK",
							Clicks:             3,
							AverageClicksByDay: 0.1,
						},
					},
				},
				"bit.ly/3nPhiCD": {
					Metrics: []schemas.AnalyticMetric{
						{
							Value:              "US",
							Clicks:             2,
							AverageClicksByDay: 0.06666666666666667,
						},
					},
				},
			},
			wantErrChan: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()

			a := Analytics{
				settings: tt.fields.settings,
			}

			a.client = openapi.NewAPIClient(&openapi.Configuration{
				Servers: openapi.ServerConfigurations{
					{
						URL: tt.server.URL,
					},
				},
			})

			got, got1 := a.getGroupBitLinkCountryMetricsFromChannel(tt.args.ctx, tt.args.timeUnit, tt.args.units, tt.args.linksChannel)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getGroupBitLinkCountryMetricsFromChannel() got = %v, want %v", got, tt.want)
			}

			if !tt.wantErrChan && len(got1) > 0 {
				t.Errorf("getGroupBitLinkCountryMetricsFromChannel() err = %v", <-got1)
			}
		})
	}
}

func TestAnalytics_GetByCountries(t *testing.T) {
	type fields struct {
		client   *openapi.APIClient
		settings settings.Settings
	}
	tests := []struct {
		name    string
		fields  fields
		req     *http.Request
		server  *httptest.Server
		want    func(rec *httptest.ResponseRecorder) bool
		wantErr bool
	}{
		{
			name: "defaults",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/"), nil)
				return req
			}(),
			fields: fields{
				settings: DebugLogger,
			},
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)
				mu := sync.RWMutex{}
				counter := 0

				m.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("user.json")))
				})

				m.HandleFunc("/groups/Bl9iekhPWD3/bitlinks", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)

					mu.RLock()
					if counter == 0 {
						w.Write([]byte(fixture("three-link-first-pagination.json")))
					} else if counter == 1 {
						w.Write([]byte(fixture("three-link-second-pagination.json")))
					} else if counter == 2 {
						w.Write([]byte(fixture("three-link-third-pagination.json")))
					} else {
						t.Log("counter should not reach above two call")
						t.Fail()
					}
					mu.RUnlock()

					mu.Lock()
					counter += 1
					mu.Unlock()
				})

				m.HandleFunc("/bitlinks/bit.ly/3nPhiCD/countries", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("bitlink1-click-metrics.json")))
				})

				m.HandleFunc("/bitlinks/bit.ly/3Aqhgo0/countries", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("bitlink2-click-metrics-multiple-countries.json")))
				})

				m.HandleFunc("/bitlinks/bit.ly/2XxugJT/countries", func(w http.ResponseWriter, r *http.Request) {
					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("bitlink3-click-metrics-no-clicks.json")))
				})

				return s
			}(),
			want: func(rec *httptest.ResponseRecorder) bool {
				is := is.New(t)
				is.Equal(http.StatusOK, rec.Code)

				ar := new(schemas.AnalyticsResponse)
				err := json.Unmarshal(rec.Body.Bytes(), ar)
				is.True(err == nil)

				is.Equal(schemas.AnalyticsResponse{
					Data: map[string]schemas.AnalyticMetrics{
						"bit.ly/3Aqhgo0": {
							Metrics: []schemas.AnalyticMetric{
								{
									Value:              "US",
									Clicks:             2,
									AverageClicksByDay: 0.06666666666666667,
								},
								{
									Value:              "UK",
									Clicks:             3,
									AverageClicksByDay: 0.1,
								},
							},
						},
						"bit.ly/3nPhiCD": {
							Metrics: []schemas.AnalyticMetric{
								{
									Value:              "US",
									Clicks:             2,
									AverageClicksByDay: 0.06666666666666667,
								},
							},
						},
					},
					Group: ExpectedDefaultGroupGUID,
					Unit:  "day",
					Units: "30",
					Facet: "country",
				}, *ar)

				return true
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()
			e := echo.New()
			is := is.New(t)

			a := Analytics{
				settings: tt.fields.settings,
			}

			a.client = openapi.NewAPIClient(&openapi.Configuration{
				Servers: openapi.ServerConfigurations{
					{
						URL: tt.server.URL,
					},
				},
			})

			e.GET("/", a.GetByCountries, FakeKeyAuthMiddleware)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, tt.req)
			defer e.Close()

			if !tt.want(rec) {
				is.Fail()
			}
		})
	}
}
