package v1

import (
	"net/http"

	openapi "github.com/jrxFive/analytics-hw/pkg/bitly"
)

type PaginatedGroupBitLinksResponse struct {
	Bitlink  openapi.BitlinkBody
	Response *http.Response
	Err      error
}

type GroupBitLinksRequestExecutorFn func() (openapi.Bitlinks, *http.Response, error)
type PaginatedGroupBitLinksRequest struct {
	req openapi.ApiGetBitlinksByGroupRequest
}

// Aggregate all paginated responses from upstream API, as a slice.
func (p PaginatedGroupBitLinksRequest) SliceResponse(fn GroupBitLinksRequestExecutorFn) ([]PaginatedGroupBitLinksResponse, error) {
	allGroupBitlinks := make([]PaginatedGroupBitLinksResponse, 0, 0)

	for {
		bitLinks, response, err := fn()
		if err != nil {
			allGroupBitlinks = append(allGroupBitlinks, PaginatedGroupBitLinksResponse{
				Response: response,
				Err:      err,
			})
			return allGroupBitlinks, err
		}

		for _, link := range *bitLinks.Links {
			allGroupBitlinks = append(allGroupBitlinks, PaginatedGroupBitLinksResponse{
				Bitlink:  link,
				Response: nil,
				Err:      nil,
			})
		}

		if page, ok := bitLinks.GetPaginationOk(); ok {
			if len(*page.Next) == 0 {
				return allGroupBitlinks, nil
			}

			p.req = p.req.Page(page.GetPage() + 1)
		}
	}

}

//For each paginated response send results through given channel.
func (p PaginatedGroupBitLinksRequest) ChannelResponse(ch chan PaginatedGroupBitLinksResponse, fn GroupBitLinksRequestExecutorFn) {
	for {
		bitLinks, response, err := fn()
		if err != nil {
			ch <- PaginatedGroupBitLinksResponse{
				Response: response,
				Err:      err,
			}
			close(ch)
			return
		}

		for _, link := range *bitLinks.Links {
			ch <- PaginatedGroupBitLinksResponse{
				Bitlink:  link,
				Response: response,
				Err:      nil,
			}
		}

		if page, ok := bitLinks.GetPaginationOk(); ok {
			if len(*page.Next) == 0 {
				close(ch)
				return
			}

			p.req = p.req.Page(page.GetPage() + 1)
		}
	}
}
