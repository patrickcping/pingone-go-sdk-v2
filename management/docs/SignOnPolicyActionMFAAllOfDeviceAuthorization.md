# SignOnPolicyActionMFAAllOfDeviceAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | A boolean that specifies the enabled/disabled state of automatic MFA for native devices paired with the user for the specified application. | [optional] [default to false]
**ExtraVerification** | Pointer to [**EnumSignOnPolicyExtraVerification**](EnumSignOnPolicyExtraVerification.md) |  | [optional] [default to ENUMSIGNONPOLICYEXTRAVERIFICATION_DISABLED]

## Methods

### NewSignOnPolicyActionMFAAllOfDeviceAuthorization

`func NewSignOnPolicyActionMFAAllOfDeviceAuthorization() *SignOnPolicyActionMFAAllOfDeviceAuthorization`

NewSignOnPolicyActionMFAAllOfDeviceAuthorization instantiates a new SignOnPolicyActionMFAAllOfDeviceAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAAllOfDeviceAuthorizationWithDefaults

`func NewSignOnPolicyActionMFAAllOfDeviceAuthorizationWithDefaults() *SignOnPolicyActionMFAAllOfDeviceAuthorization`

NewSignOnPolicyActionMFAAllOfDeviceAuthorizationWithDefaults instantiates a new SignOnPolicyActionMFAAllOfDeviceAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionMFAAllOfDeviceAuthorization) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionMFAAllOfDeviceAuthorization) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionMFAAllOfDeviceAuthorization) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SignOnPolicyActionMFAAllOfDeviceAuthorization) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraVerification

`func (o *SignOnPolicyActionMFAAllOfDeviceAuthorization) GetExtraVerification() EnumSignOnPolicyExtraVerification`

GetExtraVerification returns the ExtraVerification field if non-nil, zero value otherwise.

### GetExtraVerificationOk

`func (o *SignOnPolicyActionMFAAllOfDeviceAuthorization) GetExtraVerificationOk() (*EnumSignOnPolicyExtraVerification, bool)`

GetExtraVerificationOk returns a tuple with the ExtraVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVerification

`func (o *SignOnPolicyActionMFAAllOfDeviceAuthorization) SetExtraVerification(v EnumSignOnPolicyExtraVerification)`

SetExtraVerification sets ExtraVerification field to given value.

### HasExtraVerification

`func (o *SignOnPolicyActionMFAAllOfDeviceAuthorization) HasExtraVerification() bool`

HasExtraVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


