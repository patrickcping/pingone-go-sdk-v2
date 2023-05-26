# EmailPhoneConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verify** | [**EnumVerify**](EnumVerify.md) |  | 
**CreateMfaDevice** | Pointer to **bool** | When enabled, PingOne Verify registers the email address or phone number with PingOne MFA as a verified MFA device. | [optional] 
**Otp** | Pointer to [**EmailPhoneConfigurationOtp**](EmailPhoneConfigurationOtp.md) |  | [optional] 

## Methods

### NewEmailPhoneConfiguration

`func NewEmailPhoneConfiguration(verify EnumVerify, ) *EmailPhoneConfiguration`

NewEmailPhoneConfiguration instantiates a new EmailPhoneConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailPhoneConfigurationWithDefaults

`func NewEmailPhoneConfigurationWithDefaults() *EmailPhoneConfiguration`

NewEmailPhoneConfigurationWithDefaults instantiates a new EmailPhoneConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerify

`func (o *EmailPhoneConfiguration) GetVerify() EnumVerify`

GetVerify returns the Verify field if non-nil, zero value otherwise.

### GetVerifyOk

`func (o *EmailPhoneConfiguration) GetVerifyOk() (*EnumVerify, bool)`

GetVerifyOk returns a tuple with the Verify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerify

`func (o *EmailPhoneConfiguration) SetVerify(v EnumVerify)`

SetVerify sets Verify field to given value.


### GetCreateMfaDevice

`func (o *EmailPhoneConfiguration) GetCreateMfaDevice() bool`

GetCreateMfaDevice returns the CreateMfaDevice field if non-nil, zero value otherwise.

### GetCreateMfaDeviceOk

`func (o *EmailPhoneConfiguration) GetCreateMfaDeviceOk() (*bool, bool)`

GetCreateMfaDeviceOk returns a tuple with the CreateMfaDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMfaDevice

`func (o *EmailPhoneConfiguration) SetCreateMfaDevice(v bool)`

SetCreateMfaDevice sets CreateMfaDevice field to given value.

### HasCreateMfaDevice

`func (o *EmailPhoneConfiguration) HasCreateMfaDevice() bool`

HasCreateMfaDevice returns a boolean if a field has been set.

### GetOtp

`func (o *EmailPhoneConfiguration) GetOtp() EmailPhoneConfigurationOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *EmailPhoneConfiguration) GetOtpOk() (*EmailPhoneConfigurationOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *EmailPhoneConfiguration) SetOtp(v EmailPhoneConfigurationOtp)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *EmailPhoneConfiguration) HasOtp() bool`

HasOtp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


