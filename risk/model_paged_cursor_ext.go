package risk

import (
	"net/http"
)

const (
	PAGINATION_HAL_LINK_INDEX_SELF = "self"
	PAGINATION_HAL_LINK_INDEX_NEXT = "next"
	PAGINATION_HAL_LINK_INDEX_PREV = "prev"
)


type PagedCursor[T any] struct {
	Data  *T
	HTTPResponse *http.Response
}
