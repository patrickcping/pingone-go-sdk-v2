# ListVersions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to [**ListVersions200ResponseEmbedded**](ListVersions200ResponseEmbedded.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewListVersions200Response

`func NewListVersions200Response() *ListVersions200Response`

NewListVersions200Response instantiates a new ListVersions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVersions200ResponseWithDefaults

`func NewListVersions200ResponseWithDefaults() *ListVersions200Response`

NewListVersions200ResponseWithDefaults instantiates a new ListVersions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListVersions200Response) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListVersions200Response) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListVersions200Response) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListVersions200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *ListVersions200Response) GetEmbedded() ListVersions200ResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListVersions200Response) GetEmbeddedOk() (*ListVersions200ResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListVersions200Response) SetEmbedded(v ListVersions200ResponseEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ListVersions200Response) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetCount

`func (o *ListVersions200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListVersions200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListVersions200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListVersions200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSize

`func (o *ListVersions200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListVersions200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListVersions200Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListVersions200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


