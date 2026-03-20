# SubscriptionPayloadOptionsMaximumPayloadLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**EnumSubscriptionPayloadOptionsMaximumPayloadLimitType**](EnumSubscriptionPayloadOptionsMaximumPayloadLimitType.md) |  | [optional] 
**Size** | Pointer to **int32** | The maximum size of the payload based on &#x60;type&#x60;. For &#x60;EVENTS_PER_PAYLOAD&#x60; this can be from &#x60;1&#x60; to &#x60;500&#x60; events (defaults to &#x60;500&#x60;). For &#x60;KB_PER_PAYLOAD&#x60; this can be from &#x60;1&#x60; to &#x60;4096&#x60; kilobytes. | [optional] 

## Methods

### NewSubscriptionPayloadOptionsMaximumPayloadLimit

`func NewSubscriptionPayloadOptionsMaximumPayloadLimit() *SubscriptionPayloadOptionsMaximumPayloadLimit`

NewSubscriptionPayloadOptionsMaximumPayloadLimit instantiates a new SubscriptionPayloadOptionsMaximumPayloadLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPayloadOptionsMaximumPayloadLimitWithDefaults

`func NewSubscriptionPayloadOptionsMaximumPayloadLimitWithDefaults() *SubscriptionPayloadOptionsMaximumPayloadLimit`

NewSubscriptionPayloadOptionsMaximumPayloadLimitWithDefaults instantiates a new SubscriptionPayloadOptionsMaximumPayloadLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionPayloadOptionsMaximumPayloadLimit) GetType() EnumSubscriptionPayloadOptionsMaximumPayloadLimitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionPayloadOptionsMaximumPayloadLimit) GetTypeOk() (*EnumSubscriptionPayloadOptionsMaximumPayloadLimitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionPayloadOptionsMaximumPayloadLimit) SetType(v EnumSubscriptionPayloadOptionsMaximumPayloadLimitType)`

SetType sets Type field to given value.

### HasType

`func (o *SubscriptionPayloadOptionsMaximumPayloadLimit) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSize

`func (o *SubscriptionPayloadOptionsMaximumPayloadLimit) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *SubscriptionPayloadOptionsMaximumPayloadLimit) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *SubscriptionPayloadOptionsMaximumPayloadLimit) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *SubscriptionPayloadOptionsMaximumPayloadLimit) HasSize() bool`

HasSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


