# DeviceAuthenticationPolicyTotp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enabled or disabled in the policy. | 
**Otp** | [**DeviceAuthenticationPolicyTotpOtp**](DeviceAuthenticationPolicyTotpOtp.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyTotp

`func NewDeviceAuthenticationPolicyTotp(enabled bool, otp DeviceAuthenticationPolicyTotpOtp, ) *DeviceAuthenticationPolicyTotp`

NewDeviceAuthenticationPolicyTotp instantiates a new DeviceAuthenticationPolicyTotp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyTotpWithDefaults

`func NewDeviceAuthenticationPolicyTotpWithDefaults() *DeviceAuthenticationPolicyTotp`

NewDeviceAuthenticationPolicyTotpWithDefaults instantiates a new DeviceAuthenticationPolicyTotp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyTotp) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyTotp) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyTotp) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOtp

`func (o *DeviceAuthenticationPolicyTotp) GetOtp() DeviceAuthenticationPolicyTotpOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyTotp) GetOtpOk() (*DeviceAuthenticationPolicyTotpOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyTotp) SetOtp(v DeviceAuthenticationPolicyTotpOtp)`

SetOtp sets Otp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


