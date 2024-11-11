package verify

const (
	PAGINATION_HAL_LINK_INDEX_SELF = "self"
	PAGINATION_HAL_LINK_INDEX_NEXT = "next"
	PAGINATION_HAL_LINK_INDEX_PREV = "prev"
)

func (o *EntityArray) hasHalLink(linkIndex string) bool {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := *l
		if v, ok := links[linkIndex]; ok {
			if h, ok := v.GetHrefOk(); ok && h != nil && *h != "" {
				return true
			}
		}
	}
	return false
}

func (o *EntityArray) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o *EntityArray) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o *EntityArray) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o *EntityArray) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}
