package management

import (
	"net/http"
)

const (
	PAGINATION_HAL_LINK_INDEX_SELF = "self"
	PAGINATION_HAL_LINK_INDEX_NEXT = "next"
	PAGINATION_HAL_LINK_INDEX_PREV = "prev"
)


type PagedCursor struct {
	EntityArray  *EntityArray
	HTTPResponse *http.Response
}

func (o *PagedCursor) hasHalLink(linkIndex string) bool {

	if o.EntityArray == nil {
		return false
	}

	if l, ok := o.EntityArray.GetLinksOk(); ok && l != nil {
		links := *l
		if v, ok := links[linkIndex]; ok {
			if h, ok := v.GetHrefOk(); ok && h != nil && *h != "" {
				return true
			}
		}
	}
	return false
}

func (o *PagedCursor) getHalLink(linkIndex string) LinksHATEOASValue {

	var ret LinksHATEOASValue
	if o.EntityArray == nil {
		return ret
	}

	if l, ok := o.EntityArray.GetLinksOk(); ok && l != nil {
		links := *l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	return ret
}

func (o *PagedCursor) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {

	if o.EntityArray == nil {
		return nil, false
	}

	if l, ok := o.EntityArray.GetLinksOk(); ok && l != nil {
		links := *l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o *PagedCursor) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o *PagedCursor) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o *PagedCursor) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o *PagedCursor) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o *PagedCursor) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o *PagedCursor) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o *PagedCursor) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o *PagedCursor) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o *PagedCursor) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o *PagedCursor) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}
