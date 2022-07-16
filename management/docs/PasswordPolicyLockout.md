# PasswordPolicyLockout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DurationSeconds** | Pointer to **int32** | The length of time before a password is automatically moved out of the lock out state. The value must be a positive, non-zero integer. | [optional] 
**FailureCount** | Pointer to **int32** | The number of tries before a password is placed in the lockout state. The value must be a positive, non-zero integer. | [optional] 

## Methods

### NewPasswordPolicyLockout

`func NewPasswordPolicyLockout() *PasswordPolicyLockout`

NewPasswordPolicyLockout instantiates a new PasswordPolicyLockout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyLockoutWithDefaults

`func NewPasswordPolicyLockoutWithDefaults() *PasswordPolicyLockout`

NewPasswordPolicyLockoutWithDefaults instantiates a new PasswordPolicyLockout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDurationSeconds

`func (o *PasswordPolicyLockout) GetDurationSeconds() int32`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *PasswordPolicyLockout) GetDurationSecondsOk() (*int32, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *PasswordPolicyLockout) SetDurationSeconds(v int32)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *PasswordPolicyLockout) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.

### GetFailureCount

`func (o *PasswordPolicyLockout) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *PasswordPolicyLockout) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *PasswordPolicyLockout) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *PasswordPolicyLockout) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


