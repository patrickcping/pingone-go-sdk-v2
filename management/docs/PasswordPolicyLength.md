# PasswordPolicyLength

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Max** | Pointer to **int32** | The maximum number of characters allowed for the password. Defaults to 255. This property is not enforced when not present. | [optional] [default to 255]
**Min** | Pointer to **int32** | The minimum number of characters required for the password. This can be from 8 to 32 (inclusive). Defaults to 8 characters. This property is not enforced when not present. | [optional] [default to 8]

## Methods

### NewPasswordPolicyLength

`func NewPasswordPolicyLength() *PasswordPolicyLength`

NewPasswordPolicyLength instantiates a new PasswordPolicyLength object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyLengthWithDefaults

`func NewPasswordPolicyLengthWithDefaults() *PasswordPolicyLength`

NewPasswordPolicyLengthWithDefaults instantiates a new PasswordPolicyLength object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMax

`func (o *PasswordPolicyLength) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *PasswordPolicyLength) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *PasswordPolicyLength) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *PasswordPolicyLength) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *PasswordPolicyLength) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *PasswordPolicyLength) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *PasswordPolicyLength) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *PasswordPolicyLength) HasMin() bool`

HasMin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


