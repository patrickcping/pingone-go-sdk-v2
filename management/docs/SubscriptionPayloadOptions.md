# SubscriptionPayloadOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaximumPayloadLimit** | Pointer to [**SubscriptionPayloadOptionsMaximumPayloadLimit**](SubscriptionPayloadOptionsMaximumPayloadLimit.md) |  | [optional] 
**PayloadFormat** | Pointer to [**SubscriptionPayloadOptionsPayloadFormat**](SubscriptionPayloadOptionsPayloadFormat.md) |  | [optional] 

## Methods

### NewSubscriptionPayloadOptions

`func NewSubscriptionPayloadOptions() *SubscriptionPayloadOptions`

NewSubscriptionPayloadOptions instantiates a new SubscriptionPayloadOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionPayloadOptionsWithDefaults

`func NewSubscriptionPayloadOptionsWithDefaults() *SubscriptionPayloadOptions`

NewSubscriptionPayloadOptionsWithDefaults instantiates a new SubscriptionPayloadOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaximumPayloadLimit

`func (o *SubscriptionPayloadOptions) GetMaximumPayloadLimit() SubscriptionPayloadOptionsMaximumPayloadLimit`

GetMaximumPayloadLimit returns the MaximumPayloadLimit field if non-nil, zero value otherwise.

### GetMaximumPayloadLimitOk

`func (o *SubscriptionPayloadOptions) GetMaximumPayloadLimitOk() (*SubscriptionPayloadOptionsMaximumPayloadLimit, bool)`

GetMaximumPayloadLimitOk returns a tuple with the MaximumPayloadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumPayloadLimit

`func (o *SubscriptionPayloadOptions) SetMaximumPayloadLimit(v SubscriptionPayloadOptionsMaximumPayloadLimit)`

SetMaximumPayloadLimit sets MaximumPayloadLimit field to given value.

### HasMaximumPayloadLimit

`func (o *SubscriptionPayloadOptions) HasMaximumPayloadLimit() bool`

HasMaximumPayloadLimit returns a boolean if a field has been set.

### GetPayloadFormat

`func (o *SubscriptionPayloadOptions) GetPayloadFormat() SubscriptionPayloadOptionsPayloadFormat`

GetPayloadFormat returns the PayloadFormat field if non-nil, zero value otherwise.

### GetPayloadFormatOk

`func (o *SubscriptionPayloadOptions) GetPayloadFormatOk() (*SubscriptionPayloadOptionsPayloadFormat, bool)`

GetPayloadFormatOk returns a tuple with the PayloadFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayloadFormat

`func (o *SubscriptionPayloadOptions) SetPayloadFormat(v SubscriptionPayloadOptionsPayloadFormat)`

SetPayloadFormat sets PayloadFormat field to given value.

### HasPayloadFormat

`func (o *SubscriptionPayloadOptions) HasPayloadFormat() bool`

HasPayloadFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


