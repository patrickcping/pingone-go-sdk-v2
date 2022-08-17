# SignOnPolicyActionMFAAllOfEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | A boolean that specifies the enabled/disabled state of the MFA through email action. The default is disbled when creating a new policy. When enabled, it allows users to receive the one-time password and authenticate through email. | [optional] [default to false]

## Methods

### NewSignOnPolicyActionMFAAllOfEmail

`func NewSignOnPolicyActionMFAAllOfEmail() *SignOnPolicyActionMFAAllOfEmail`

NewSignOnPolicyActionMFAAllOfEmail instantiates a new SignOnPolicyActionMFAAllOfEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAAllOfEmailWithDefaults

`func NewSignOnPolicyActionMFAAllOfEmailWithDefaults() *SignOnPolicyActionMFAAllOfEmail`

NewSignOnPolicyActionMFAAllOfEmailWithDefaults instantiates a new SignOnPolicyActionMFAAllOfEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionMFAAllOfEmail) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionMFAAllOfEmail) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionMFAAllOfEmail) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SignOnPolicyActionMFAAllOfEmail) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


