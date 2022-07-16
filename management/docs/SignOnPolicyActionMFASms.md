# SignOnPolicyActionMFASms

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | A boolean that specifies the enabled/disabled state of the MFA through SMS action. The default is disabled when creating a new policy. When enabled, it allows users to receive the one-time password and authenticate through SMS text message. | [optional] [default to false]

## Methods

### NewSignOnPolicyActionMFASms

`func NewSignOnPolicyActionMFASms() *SignOnPolicyActionMFASms`

NewSignOnPolicyActionMFASms instantiates a new SignOnPolicyActionMFASms object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFASmsWithDefaults

`func NewSignOnPolicyActionMFASmsWithDefaults() *SignOnPolicyActionMFASms`

NewSignOnPolicyActionMFASmsWithDefaults instantiates a new SignOnPolicyActionMFASms object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionMFASms) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionMFASms) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionMFASms) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SignOnPolicyActionMFASms) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


