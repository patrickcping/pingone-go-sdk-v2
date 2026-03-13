# DeviceAuthenticationPolicyCommonTotp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enabled or disabled in the policy. | 
**PairingDisabled** | Pointer to **bool** | You can set &#x60;pairingDisabled&#x60; to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices. | [optional] 
**Otp** | [**DeviceAuthenticationPolicyPingIDDeviceOtp**](DeviceAuthenticationPolicyPingIDDeviceOtp.md) |  | 
**PromptForNicknameOnPairing** | Pointer to **bool** | Set to &#x60;true&#x60; if you want to allow users to provide nicknames for devices during pairing. | [optional] 
**UriParameters** | Pointer to **map[string]string** | Object that you can use to provide key:value pairs for &#x60;otpauth&#x60; URI parameters. For example, if you provide a value for the &#x60;issuer&#x60; parameter, then authenticators that support that parameter will display the text you specify together with the OTP (in addition to the username). This can help users recognize which application the OTP is for. If you intend on using the same MFA policy for multiple applications, choose a name that reflects the group of applications. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyCommonTotp

`func NewDeviceAuthenticationPolicyCommonTotp(enabled bool, otp DeviceAuthenticationPolicyPingIDDeviceOtp, ) *DeviceAuthenticationPolicyCommonTotp`

NewDeviceAuthenticationPolicyCommonTotp instantiates a new DeviceAuthenticationPolicyCommonTotp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonTotpWithDefaults

`func NewDeviceAuthenticationPolicyCommonTotpWithDefaults() *DeviceAuthenticationPolicyCommonTotp`

NewDeviceAuthenticationPolicyCommonTotpWithDefaults instantiates a new DeviceAuthenticationPolicyCommonTotp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyCommonTotp) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyCommonTotp) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyCommonTotp) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPairingDisabled

`func (o *DeviceAuthenticationPolicyCommonTotp) GetPairingDisabled() bool`

GetPairingDisabled returns the PairingDisabled field if non-nil, zero value otherwise.

### GetPairingDisabledOk

`func (o *DeviceAuthenticationPolicyCommonTotp) GetPairingDisabledOk() (*bool, bool)`

GetPairingDisabledOk returns a tuple with the PairingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingDisabled

`func (o *DeviceAuthenticationPolicyCommonTotp) SetPairingDisabled(v bool)`

SetPairingDisabled sets PairingDisabled field to given value.

### HasPairingDisabled

`func (o *DeviceAuthenticationPolicyCommonTotp) HasPairingDisabled() bool`

HasPairingDisabled returns a boolean if a field has been set.

### GetOtp

`func (o *DeviceAuthenticationPolicyCommonTotp) GetOtp() DeviceAuthenticationPolicyPingIDDeviceOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyCommonTotp) GetOtpOk() (*DeviceAuthenticationPolicyPingIDDeviceOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyCommonTotp) SetOtp(v DeviceAuthenticationPolicyPingIDDeviceOtp)`

SetOtp sets Otp field to given value.


### GetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyCommonTotp) GetPromptForNicknameOnPairing() bool`

GetPromptForNicknameOnPairing returns the PromptForNicknameOnPairing field if non-nil, zero value otherwise.

### GetPromptForNicknameOnPairingOk

`func (o *DeviceAuthenticationPolicyCommonTotp) GetPromptForNicknameOnPairingOk() (*bool, bool)`

GetPromptForNicknameOnPairingOk returns a tuple with the PromptForNicknameOnPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyCommonTotp) SetPromptForNicknameOnPairing(v bool)`

SetPromptForNicknameOnPairing sets PromptForNicknameOnPairing field to given value.

### HasPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyCommonTotp) HasPromptForNicknameOnPairing() bool`

HasPromptForNicknameOnPairing returns a boolean if a field has been set.

### GetUriParameters

`func (o *DeviceAuthenticationPolicyCommonTotp) GetUriParameters() map[string]string`

GetUriParameters returns the UriParameters field if non-nil, zero value otherwise.

### GetUriParametersOk

`func (o *DeviceAuthenticationPolicyCommonTotp) GetUriParametersOk() (*map[string]string, bool)`

GetUriParametersOk returns a tuple with the UriParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUriParameters

`func (o *DeviceAuthenticationPolicyCommonTotp) SetUriParameters(v map[string]string)`

SetUriParameters sets UriParameters field to given value.

### HasUriParameters

`func (o *DeviceAuthenticationPolicyCommonTotp) HasUriParameters() bool`

HasUriParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


