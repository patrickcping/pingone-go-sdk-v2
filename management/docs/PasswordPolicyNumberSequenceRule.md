# PasswordPolicyNumberSequenceRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLength** | Pointer to **int32** | Maximum number of allowed sequential numbers in the password. Must be a value of 2 or 3. | [optional] 

## Methods

### NewPasswordPolicyNumberSequenceRule

`func NewPasswordPolicyNumberSequenceRule() *PasswordPolicyNumberSequenceRule`

NewPasswordPolicyNumberSequenceRule instantiates a new PasswordPolicyNumberSequenceRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyNumberSequenceRuleWithDefaults

`func NewPasswordPolicyNumberSequenceRuleWithDefaults() *PasswordPolicyNumberSequenceRule`

NewPasswordPolicyNumberSequenceRuleWithDefaults instantiates a new PasswordPolicyNumberSequenceRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLength

`func (o *PasswordPolicyNumberSequenceRule) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *PasswordPolicyNumberSequenceRule) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *PasswordPolicyNumberSequenceRule) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *PasswordPolicyNumberSequenceRule) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


