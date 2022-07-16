# PasswordPolicyHistory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | Specifies the number of prior passwords to keep for prevention of password re-use. The value must be a positive, non-zero integer. | [optional] 
**RetentionDays** | Pointer to **int32** | The length of time to keep recent passwords for prevention of password re-use. The value must be a positive, non-zero integer. | [optional] 

## Methods

### NewPasswordPolicyHistory

`func NewPasswordPolicyHistory() *PasswordPolicyHistory`

NewPasswordPolicyHistory instantiates a new PasswordPolicyHistory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyHistoryWithDefaults

`func NewPasswordPolicyHistoryWithDefaults() *PasswordPolicyHistory`

NewPasswordPolicyHistoryWithDefaults instantiates a new PasswordPolicyHistory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *PasswordPolicyHistory) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PasswordPolicyHistory) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PasswordPolicyHistory) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PasswordPolicyHistory) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetRetentionDays

`func (o *PasswordPolicyHistory) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *PasswordPolicyHistory) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *PasswordPolicyHistory) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *PasswordPolicyHistory) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


