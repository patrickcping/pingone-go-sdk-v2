# SignOnPolicyActionMFAAuthenticator

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | A boolean that specifies the enabled/disabled state of the MFA through a Time-based One-time Password (TOTP) action. The default is disabled when creating a new policy. When enabled, it allows users to authenticate using a TOTP authenticator application. | [optional] [default to false]

## Methods

### NewSignOnPolicyActionMFAAuthenticator

`func NewSignOnPolicyActionMFAAuthenticator() *SignOnPolicyActionMFAAuthenticator`

NewSignOnPolicyActionMFAAuthenticator instantiates a new SignOnPolicyActionMFAAuthenticator object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAAuthenticatorWithDefaults

`func NewSignOnPolicyActionMFAAuthenticatorWithDefaults() *SignOnPolicyActionMFAAuthenticator`

NewSignOnPolicyActionMFAAuthenticatorWithDefaults instantiates a new SignOnPolicyActionMFAAuthenticator object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionMFAAuthenticator) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionMFAAuthenticator) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionMFAAuthenticator) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SignOnPolicyActionMFAAuthenticator) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


