# OTPDeviceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verify** | [**EnumVerify**](EnumVerify.md) |  | 
**CreateMfaDevice** | Pointer to **bool** | When enabled, PingOne Verify registers the email address or phone number with PingOne MFA as a verified MFA device. | [optional] 
**Otp** | Pointer to [**OTPDeviceConfigurationOtp**](OTPDeviceConfigurationOtp.md) |  | [optional] 

## Methods

### NewOTPDeviceConfiguration

`func NewOTPDeviceConfiguration(verify EnumVerify, ) *OTPDeviceConfiguration`

NewOTPDeviceConfiguration instantiates a new OTPDeviceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOTPDeviceConfigurationWithDefaults

`func NewOTPDeviceConfigurationWithDefaults() *OTPDeviceConfiguration`

NewOTPDeviceConfigurationWithDefaults instantiates a new OTPDeviceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerify

`func (o *OTPDeviceConfiguration) GetVerify() EnumVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *OTPDeviceConfiguration) GetVerifyOk() (*EnumVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *OTPDeviceConfiguration) SetVerify(v EnumVerify)`

SetVerify sets Verify field to given value.


### GetCreateMfaDevice

`func (o *OTPDeviceConfiguration) GetCreateMfaDevice() bool`

GetCreateMfaDevice returns the CreateMfaDevice field if non-nil, zero value otherwise.

### GetCreateMfaDeviceOk

`func (o *OTPDeviceConfiguration) GetCreateMfaDeviceOk() (*bool, bool)`

GetCreateMfaDeviceOk returns a tuple with the CreateMfaDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMfaDevice

`func (o *OTPDeviceConfiguration) SetCreateMfaDevice(v bool)`

SetCreateMfaDevice sets CreateMfaDevice field to given value.

### HasCreateMfaDevice

`func (o *OTPDeviceConfiguration) HasCreateMfaDevice() bool`

HasCreateMfaDevice returns a boolean if a field has been set.

### GetOtp

`func (o *OTPDeviceConfiguration) GetOtp() OTPDeviceConfigurationOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *OTPDeviceConfiguration) GetOtpOk() (*OTPDeviceConfigurationOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *OTPDeviceConfiguration) SetOtp(v OTPDeviceConfigurationOtp)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *OTPDeviceConfiguration) HasOtp() bool`

HasOtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


