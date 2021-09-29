package v1

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/matryer/is"

	openapi "github.com/jrxFive/analytics-hw/pkg/bitly"
)

func TestPaginatedGroupBitLinksRequest_SliceResponse(t *testing.T) {
	groupId := "Bl9iekhPWD3"

	tests := []struct {
		name    string
		server  *httptest.Server
		want    func(r []PaginatedGroupBitLinksResponse) bool
		wantErr bool
	}{
		{
			name: "one link no pagination in response",
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)

				m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("one-link-no-pagination.json")))
				})

				return s
			}(),
			want: func(r []PaginatedGroupBitLinksResponse) bool {
				is := is.New(t)

				is.Equal(1, len(r))
				is.Equal("bit.ly/3Aqhgo0", *r[0].Bitlink.Id)

				return true
			},

			wantErr: false,
		},
		{
			name: "two link pagination in response",
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
			want: func(r []PaginatedGroupBitLinksResponse) bool {
				is := is.New(t)
				is.Equal(3, len(r))

				is.Equal("bit.ly/3nPhiCD", *r[0].Bitlink.Id)
				is.Equal("bit.ly/2XxugJT", *r[1].Bitlink.Id)
				is.Equal("bit.ly/3Aqhgo0", *r[2].Bitlink.Id)

				return true
			},

			wantErr: false,
		},
		{
			name: "error forbidden",
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
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()

			bitlyClient := openapi.NewAPIClient(&openapi.Configuration{
				Servers: openapi.ServerConfigurations{
					{
						URL: tt.server.URL,
					},
				},
			})

			req := bitlyClient.BitlinksApi.GetBitlinksByGroup(context.TODO(), groupId)
			fn := func() (openapi.Bitlinks, *http.Response, error) {
				links, resp, err := bitlyClient.BitlinksApi.GetBitlinksByGroupExecute(req)
				return links, resp, err
			}

			p := PaginatedGroupBitLinksRequest{
				req,
			}

			got, err := p.SliceResponse(fn)
			if (err != nil) != tt.wantErr {
				t.Errorf("SliceResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if (err != nil) && tt.wantErr {
				return
			}

			tt.want(got)
		})
	}
}

func TestPaginatedGroupBitLinksRequest_ChannelResponse(t *testing.T) {
	groupId := "Bl9iekhPWD3"

	tests := []struct {
		name    string
		ch      chan PaginatedGroupBitLinksResponse
		server  *httptest.Server
		want    func(ch chan PaginatedGroupBitLinksResponse, wantErr bool) bool
		wantErr bool
	}{
		{
			name: "one link no pagination in response",
			ch:   make(chan PaginatedGroupBitLinksResponse, 10),
			server: func() *httptest.Server {
				m := http.NewServeMux()
				s := httptest.NewServer(m)

				m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

					w.Header().Add("Content-Type", "application/json")
					w.WriteHeader(http.StatusOK)
					w.Write([]byte(fixture("one-link-no-pagination.json")))
				})

				return s
			}(),
			want: func(ch chan PaginatedGroupBitLinksResponse, wantErr bool) bool {
				is := is.New(t)

				is.Equal(1, len(ch))

				r := <-ch
				is.Equal("bit.ly/3Aqhgo0", *r.Bitlink.Id)

				return true
			},
			wantErr: false,
		},
		{
			name: "three link pagination in response",
			ch:   make(chan PaginatedGroupBitLinksResponse, 10),
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
		{
			name: "error forbidden",
			ch:   make(chan PaginatedGroupBitLinksResponse, 10),
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
			want: func(ch chan PaginatedGroupBitLinksResponse, wantErr bool) bool {
				is := is.New(t)
				is.Equal(1, len(ch))

				r := <-ch
				is.True(r.Err != nil && wantErr)

				return true
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer tt.server.Close()

			bitlyClient := openapi.NewAPIClient(&openapi.Configuration{
				Servers: openapi.ServerConfigurations{
					{
						URL: tt.server.URL,
					},
				},
			})

			req := bitlyClient.BitlinksApi.GetBitlinksByGroup(context.TODO(), groupId)
			fn := func() (openapi.Bitlinks, *http.Response, error) {
				links, resp, err := bitlyClient.BitlinksApi.GetBitlinksByGroupExecute(req)
				return links, resp, err
			}

			p := PaginatedGroupBitLinksRequest{
				req,
			}

			p.ChannelResponse(tt.ch, fn)
			tt.want(tt.ch, tt.wantErr)
		})
	}
}
