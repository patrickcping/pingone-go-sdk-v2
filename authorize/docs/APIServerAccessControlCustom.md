# APIServerAccessControlCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | If &#x60;TRUE&#x60;, custom policy will be used for the endpoint. Defaults to &#x60;FALSE&#x60;. | [optional] 

## Methods

### NewAPIServerAccessControlCustom

`func NewAPIServerAccessControlCustom() *APIServerAccessControlCustom`

NewAPIServerAccessControlCustom instantiates a new APIServerAccessControlCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerAccessControlCustomWithDefaults

`func NewAPIServerAccessControlCustomWithDefaults() *APIServerAccessControlCustom`

NewAPIServerAccessControlCustomWithDefaults instantiates a new APIServerAccessControlCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *APIServerAccessControlCustom) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *APIServerAccessControlCustom) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *APIServerAccessControlCustom) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *APIServerAccessControlCustom) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


