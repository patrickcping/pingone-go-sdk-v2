# ListStatements200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to [**ListStatements200ResponseEmbedded**](ListStatements200ResponseEmbedded.md) |  | [optional] 
**Count** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 

## Methods

### NewListStatements200Response

`func NewListStatements200Response() *ListStatements200Response`

NewListStatements200Response instantiates a new ListStatements200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStatements200ResponseWithDefaults

`func NewListStatements200ResponseWithDefaults() *ListStatements200Response`

NewListStatements200ResponseWithDefaults instantiates a new ListStatements200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ListStatements200Response) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ListStatements200Response) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ListStatements200Response) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ListStatements200Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *ListStatements200Response) GetEmbedded() ListStatements200ResponseEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *ListStatements200Response) GetEmbeddedOk() (*ListStatements200ResponseEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *ListStatements200Response) SetEmbedded(v ListStatements200ResponseEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *ListStatements200Response) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetCount

`func (o *ListStatements200Response) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ListStatements200Response) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ListStatements200Response) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *ListStatements200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetSize

`func (o *ListStatements200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListStatements200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListStatements200Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListStatements200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


