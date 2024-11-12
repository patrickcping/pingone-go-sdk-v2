package authorize

import (
	"context"
	"encoding/json"
	"fmt"
	"iter"
	"net/http"
)

type EntityArrayPagedIterator iter.Seq2[PagedCursor, error]

func (a *APIClient) paginationIterator(ctx context.Context, initialPageAPIFunc func() (*EntityArray, *http.Response, error)) EntityArrayPagedIterator {
	var err error
	return func(yield func(PagedCursor, error) bool) {

		initialCursor := PagedCursor{}
		initialCursor.EntityArray, initialCursor.HTTPResponse, err = initialPageAPIFunc()

		if !yield(initialCursor, err) {
			return
		}

		latestCursorData := *initialCursor.EntityArray

		for latestCursorData.HasPaginationNext() {
			loopCursor := PagedCursor{}

			var halResponse interface{}
			halResponse, loopCursor.HTTPResponse, err = a.HALApi.ReadHALLink(ctx, latestCursorData.GetPaginationNextLink()).Execute()
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

			err = json.Unmarshal(bytes, &loopCursor.EntityArray)
			if err != nil {
				if !yield(loopCursor, err) {
					return
				}
				break
			}

			if loopCursor.EntityArray == nil {
				if !yield(loopCursor, fmt.Errorf("Paged results unexpectedly nil")) {
					return
				}
				break
			}

			latestCursorData = *loopCursor.EntityArray

			if !yield(loopCursor, nil) {
				return
			}
		}
	}
}
