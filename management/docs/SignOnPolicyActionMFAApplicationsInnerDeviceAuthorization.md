# SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | A boolean that specifies the enabled/disabled state of automatic MFA for native devices paired with the user for the specified application. | [optional] [default to false]
**ExtraVerification** | Pointer to **string** | Specifies the level of further verification when deviceAuthorization is enabled. The PingOne platform performs an extra verification check by sending a silent push notification to the customer native application, and receives a confirmation in return.  &#x60;extraVerification&#x60; can be one of the following levels: &#x60;disabled&#x60; (default): The PingOne platform does not perform the extra verification check. &#x60;permissive&#x60;: The PingOne platform performs the extra verification check. Upon timeout or failure to get a response from the native app, the MFA step is treated as successfully completed. &#x60;restrictive&#x60;: The PingOne platform performs the extra verification check.The PingOne platform performs the extra verification check. Upon timeout or failure to get a response from the native app, the MFA step is treated as failed.  | [optional] [default to "disabled"]

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

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetExtraVerification() string`

GetExtraVerification returns the ExtraVerification field if non-nil, zero value otherwise.

### GetExtraVerificationOk

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) GetExtraVerificationOk() (*string, bool)`

GetExtraVerificationOk returns a tuple with the ExtraVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVerification

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) SetExtraVerification(v string)`

SetExtraVerification sets ExtraVerification field to given value.

### HasExtraVerification

`func (o *SignOnPolicyActionMFAApplicationsInnerDeviceAuthorization) HasExtraVerification() bool`

HasExtraVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


