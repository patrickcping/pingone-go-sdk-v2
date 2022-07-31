# DeviceAuthenticationPolicyMobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled or disabled in the policy. | [optional] 
**Otp** | Pointer to [**DeviceAuthenticationPolicyMobileOtp**](DeviceAuthenticationPolicyMobileOtp.md) |  | [optional] 
**Applications** | Pointer to [**[]DeviceAuthenticationPolicyMobileApplicationsInner**](DeviceAuthenticationPolicyMobileApplicationsInner.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyMobile

`func NewDeviceAuthenticationPolicyMobile() *DeviceAuthenticationPolicyMobile`

NewDeviceAuthenticationPolicyMobile instantiates a new DeviceAuthenticationPolicyMobile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyMobileWithDefaults

`func NewDeviceAuthenticationPolicyMobileWithDefaults() *DeviceAuthenticationPolicyMobile`

NewDeviceAuthenticationPolicyMobileWithDefaults instantiates a new DeviceAuthenticationPolicyMobile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyMobile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyMobile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyMobile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeviceAuthenticationPolicyMobile) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOtp

`func (o *DeviceAuthenticationPolicyMobile) GetOtp() DeviceAuthenticationPolicyMobileOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyMobile) GetOtpOk() (*DeviceAuthenticationPolicyMobileOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyMobile) SetOtp(v DeviceAuthenticationPolicyMobileOtp)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *DeviceAuthenticationPolicyMobile) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetApplications

`func (o *DeviceAuthenticationPolicyMobile) GetApplications() []DeviceAuthenticationPolicyMobileApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeviceAuthenticationPolicyMobile) GetApplicationsOk() (*[]DeviceAuthenticationPolicyMobileApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeviceAuthenticationPolicyMobile) SetApplications(v []DeviceAuthenticationPolicyMobileApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeviceAuthenticationPolicyMobile) HasApplications() bool`

HasApplications returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


