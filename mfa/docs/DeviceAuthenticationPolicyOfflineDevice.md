# DeviceAuthenticationPolicyOfflineDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Enabled or disabled in the policy. | [optional] 
**Otp** | Pointer to [**DeviceAuthenticationPolicyOfflineDeviceOtp**](DeviceAuthenticationPolicyOfflineDeviceOtp.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyOfflineDevice

`func NewDeviceAuthenticationPolicyOfflineDevice() *DeviceAuthenticationPolicyOfflineDevice`

NewDeviceAuthenticationPolicyOfflineDevice instantiates a new DeviceAuthenticationPolicyOfflineDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyOfflineDeviceWithDefaults

`func NewDeviceAuthenticationPolicyOfflineDeviceWithDefaults() *DeviceAuthenticationPolicyOfflineDevice`

NewDeviceAuthenticationPolicyOfflineDeviceWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyOfflineDevice) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyOfflineDevice) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyOfflineDevice) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeviceAuthenticationPolicyOfflineDevice) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetOtp

`func (o *DeviceAuthenticationPolicyOfflineDevice) GetOtp() DeviceAuthenticationPolicyOfflineDeviceOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyOfflineDevice) GetOtpOk() (*DeviceAuthenticationPolicyOfflineDeviceOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyOfflineDevice) SetOtp(v DeviceAuthenticationPolicyOfflineDeviceOtp)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *DeviceAuthenticationPolicyOfflineDevice) HasOtp() bool`

HasOtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


