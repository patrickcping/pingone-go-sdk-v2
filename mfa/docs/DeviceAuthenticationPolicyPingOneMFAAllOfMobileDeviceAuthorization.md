# DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Specifies the enabled or disabled state of automatic MFA for native devices paired with the user, for the specified application. | 
**ExtraVerification** | Pointer to [**EnumMFADevicePolicyMobileExtraVerification**](EnumMFADevicePolicyMobileExtraVerification.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization(enabled bool, ) *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorizationWithDefaults

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorizationWithDefaults() *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorizationWithDefaults instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetExtraVerification

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization) GetExtraVerification() EnumMFADevicePolicyMobileExtraVerification`

GetExtraVerification returns the ExtraVerification field if non-nil, zero value otherwise.

### GetExtraVerificationOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization) GetExtraVerificationOk() (*EnumMFADevicePolicyMobileExtraVerification, bool)`

GetExtraVerificationOk returns a tuple with the ExtraVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtraVerification

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization) SetExtraVerification(v EnumMFADevicePolicyMobileExtraVerification)`

SetExtraVerification sets ExtraVerification field to given value.

### HasExtraVerification

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization) HasExtraVerification() bool`

HasExtraVerification returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


