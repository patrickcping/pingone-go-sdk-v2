# SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | A boolean that specifies the enabled/disabled state of automatic MFA for native devices paired with the user for the specified application. | [optional] [default to false]
**ExtraVerification** | Pointer to [**EnumSignOnPolicyExtraVerification**](EnumSignOnPolicyExtraVerification.md) |  | [optional] [default to DISABLED]

## Methods

### NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization

`func NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization() *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization`

NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorization instantiates a new SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorizationWithDefaults

`func NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorizationWithDefaults() *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization`

NewSignOnPolicyActionMFAApplicationsInnerDeviceAuthorizationWithDefaults instantiates a new SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraVerification

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetExtraVerification() EnumSignOnPolicyExtraVerification`

GetExtraVerification returns the ExtraVerification field if non-nil, zero value otherwise.

### GetExtraVerificationOk

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetExtraVerificationOk() (*EnumSignOnPolicyExtraVerification, bool)`

GetExtraVerificationOk returns a tuple with the ExtraVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVerification

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) SetExtraVerification(v EnumSignOnPolicyExtraVerification)`

SetExtraVerification sets ExtraVerification field to given value.

### HasExtraVerification

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) HasExtraVerification() bool`

HasExtraVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


