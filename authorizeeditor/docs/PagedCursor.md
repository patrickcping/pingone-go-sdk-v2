# PagedCursor

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityArray** | Pointer to [**EntityArray**](EntityArray.md) |  | [optional]
**HTTPResponse** | Pointer to **http.Response** |  | [optional]

## Methods

### GetPaginationSelfLink

`func (o *PagedCursor) GetPaginationSelfLink() LinksHATEOASValue`

GetPaginationSelfLink returns the `self` Links field if non-nil, zero value otherwise.

### GetPaginationSelfLinkOk

`func (o *PagedCursor) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool)`

GetPaginationSelfLinkOk returns a tuple with the `self` Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPaginationSelf

`func (o *PagedCursor) HasPaginationSelf() bool`

HasLinks returns a boolean if the `self` HAL field has been set.

### GetPaginationNextLink

`func (o *PagedCursor) GetPaginationNextLink() LinksHATEOASValue`

GetPaginationNextLink returns the `next` Links field if non-nil, zero value otherwise.

### GetPaginationNextLinkOk

`func (o *PagedCursor) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool)`

GetPaginationNextLinkOk returns a tuple with the `next` Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPaginationNext

`func (o *PagedCursor) HasPaginationNext() bool`

HasLinks returns a boolean if the `next` HAL field has been set.

### GetPaginationPreviousLink

`func (o *PagedCursor) GetPaginationPreviousLink() LinksHATEOASValue`

GetPaginationPreviousLink returns the `prev` Links field if non-nil, zero value otherwise.

### GetPaginationPreviousLinkOk

`func (o *PagedCursor) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool)`

GetPaginationPreviousLinkOk returns a tuple with the `prev` Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### HasPaginationPrevious

`func (o *PagedCursor) HasPaginationPrevious() bool`

HasLinks returns a boolean if the `prev` HAL field has been set.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
