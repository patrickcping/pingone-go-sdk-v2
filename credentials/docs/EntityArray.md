# EntityArray

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Embedded** | Pointer to [**EntityArrayEmbedded**](EntityArrayEmbedded.md) |  | [optional] 
**Size** | Pointer to **float32** |  | [optional] 

## Methods

### NewEntityArray

`func NewEntityArray() *EntityArray`

NewEntityArray instantiates a new EntityArray object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayWithDefaults

`func NewEntityArrayWithDefaults() *EntityArray`

NewEntityArrayWithDefaults instantiates a new EntityArray object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmbedded

`func (o *EntityArray) GetEmbedded() EntityArrayEmbedded`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *EntityArray) GetEmbeddedOk() (*EntityArrayEmbedded, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *EntityArray) SetEmbedded(v EntityArrayEmbedded)`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *EntityArray) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetSize

`func (o *EntityArray) GetSize() float32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *EntityArray) GetSizeOk() (*float32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *EntityArray) SetSize(v float32)`

SetSize sets Size field to given value.

### HasSize

`func (o *EntityArray) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


