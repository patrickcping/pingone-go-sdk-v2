# SignOnPolicyActionMFAAllOfSms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | A boolean that specifies the enabled/disabled state of the MFA through SMS action. The default is disabled when creating a new policy. When enabled, it allows users to receive the one-time password and authenticate through SMS text message. | [optional] [default to false]

## Methods

### NewSignOnPolicyActionMFAAllOfSms

`func NewSignOnPolicyActionMFAAllOfSms() *SignOnPolicyActionMFAAllOfSms`

NewSignOnPolicyActionMFAAllOfSms instantiates a new SignOnPolicyActionMFAAllOfSms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAAllOfSmsWithDefaults

`func NewSignOnPolicyActionMFAAllOfSmsWithDefaults() *SignOnPolicyActionMFAAllOfSms`

NewSignOnPolicyActionMFAAllOfSmsWithDefaults instantiates a new SignOnPolicyActionMFAAllOfSms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionMFAAllOfSms) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionMFAAllOfSms) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionMFAAllOfSms) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SignOnPolicyActionMFAAllOfSms) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


