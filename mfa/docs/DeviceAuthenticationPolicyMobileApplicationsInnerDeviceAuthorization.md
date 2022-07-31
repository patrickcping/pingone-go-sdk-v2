# DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies the enabled or disabled state of automatic MFA for native devices paired with the user, for the specified application. | [optional] 
**ExtraVerification** | Pointer to **string** | Specifies the level of further verification when deviceAuthorization is enabled. The PingOne platform performs an extra verification check by sending a “silent” push notification to the customer native application, and receives a confirmation in return. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization

`func NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization() *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization`

NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorizationWithDefaults

`func NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorizationWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization`

NewDeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorizationWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExtraVerification

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) GetExtraVerification() string`

GetExtraVerification returns the ExtraVerification field if non-nil, zero value otherwise.

### GetExtraVerificationOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) GetExtraVerificationOk() (*string, bool)`

GetExtraVerificationOk returns a tuple with the ExtraVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVerification

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) SetExtraVerification(v string)`

SetExtraVerification sets ExtraVerification field to given value.

### HasExtraVerification

`func (o *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) HasExtraVerification() bool`

HasExtraVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


