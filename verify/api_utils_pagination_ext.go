package verify

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/http"
)

type PagedResponseModel interface {
	IsPaginated() bool
	HasPaginationSelf() bool
	GetPaginationSelfLink() LinksHATEOASValue
	GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool)
	HasPaginationNext() bool
	GetPaginationNextLink() LinksHATEOASValue
	GetPaginationNextLinkOk() (*LinksHATEOASValue, bool)
	HasPaginationPrevious() bool
	GetPaginationPreviousLink() LinksHATEOASValue
	GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool)
}

type PagedIterator[T PagedResponseModel] iter.Seq2[PagedCursor[T], error]

func paginationIterator[T PagedResponseModel](ctx context.Context, apiClient *APIClient, initialPageAPIFunc func() (*T, *http.Response, error)) PagedIterator[T] {
	var err error
	return func(yield func(PagedCursor[T], error) bool) {

		initialCursor := PagedCursor[T]{}
		initialCursor.Data, initialCursor.HTTPResponse, err = initialPageAPIFunc()

		if !yield(initialCursor, err) {
			return
		}

		latestCursorData := *initialCursor.Data

		for latestCursorData.HasPaginationNext() {
			loopCursor := PagedCursor[T]{}

			var halResponse interface{}
			halResponse, loopCursor.HTTPResponse, err = apiClient.HALApi.ReadHALLink(ctx, latestCursorData.GetPaginationNextLink()).Execute()
			if err != nil {
				if !yield(loopCursor, err) {
					return
				}
				break
			}

			bytes, err := json.Marshal(halResponse)
			if err != nil {
				if !yield(loopCursor, err) {
					return
				}
				break
			}

			err = json.Unmarshal(bytes, &loopCursor.Data)
			if err != nil {
				if !yield(loopCursor, err) {
					return
				}
				break
			}

			if loopCursor.Data == nil {
				if !yield(loopCursor, fmt.Errorf("Paged results unexpectedly nil")) {
					return
				}
				break
			}

			latestCursorData = *loopCursor.Data

			if !yield(loopCursor, nil) {
				return
			}
		}
	}
}
