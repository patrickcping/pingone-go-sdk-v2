# DeviceAuthenticationPolicyCommonMobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A boolean that specifies whether the method is enabled or disabled in the policy. | 
**Otp** | [**DeviceAuthenticationPolicyCommonMobileOtp**](DeviceAuthenticationPolicyCommonMobileOtp.md) |  | 
**Applications** | Pointer to [**[]DeviceAuthenticationPolicyCommonMobileApplicationsInner**](DeviceAuthenticationPolicyCommonMobileApplicationsInner.md) |  | [optional] 
**IpPairingConfiguration** | Pointer to [**DeviceAuthenticationPolicyCommonMobileIpPairingConfiguration**](DeviceAuthenticationPolicyCommonMobileIpPairingConfiguration.md) |  | [optional] 
**PromptForNicknameOnPairing** | Pointer to **bool** | Set to &#x60;true&#x60; if you want to allow users to provide nicknames for devices during pairing. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyCommonMobile

`func NewDeviceAuthenticationPolicyCommonMobile(enabled bool, otp DeviceAuthenticationPolicyCommonMobileOtp, ) *DeviceAuthenticationPolicyCommonMobile`

NewDeviceAuthenticationPolicyCommonMobile instantiates a new DeviceAuthenticationPolicyCommonMobile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonMobileWithDefaults

`func NewDeviceAuthenticationPolicyCommonMobileWithDefaults() *DeviceAuthenticationPolicyCommonMobile`

NewDeviceAuthenticationPolicyCommonMobileWithDefaults instantiates a new DeviceAuthenticationPolicyCommonMobile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyCommonMobile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyCommonMobile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyCommonMobile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOtp

`func (o *DeviceAuthenticationPolicyCommonMobile) GetOtp() DeviceAuthenticationPolicyCommonMobileOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyCommonMobile) GetOtpOk() (*DeviceAuthenticationPolicyCommonMobileOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyCommonMobile) SetOtp(v DeviceAuthenticationPolicyCommonMobileOtp)`

SetOtp sets Otp field to given value.


### GetApplications

`func (o *DeviceAuthenticationPolicyCommonMobile) GetApplications() []DeviceAuthenticationPolicyCommonMobileApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeviceAuthenticationPolicyCommonMobile) GetApplicationsOk() (*[]DeviceAuthenticationPolicyCommonMobileApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeviceAuthenticationPolicyCommonMobile) SetApplications(v []DeviceAuthenticationPolicyCommonMobileApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeviceAuthenticationPolicyCommonMobile) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobile) GetIpPairingConfiguration() DeviceAuthenticationPolicyCommonMobileIpPairingConfiguration`

GetIpPairingConfiguration returns the IpPairingConfiguration field if non-nil, zero value otherwise.

### GetIpPairingConfigurationOk

`func (o *DeviceAuthenticationPolicyCommonMobile) GetIpPairingConfigurationOk() (*DeviceAuthenticationPolicyCommonMobileIpPairingConfiguration, bool)`

GetIpPairingConfigurationOk returns a tuple with the IpPairingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobile) SetIpPairingConfiguration(v DeviceAuthenticationPolicyCommonMobileIpPairingConfiguration)`

SetIpPairingConfiguration sets IpPairingConfiguration field to given value.

### HasIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobile) HasIpPairingConfiguration() bool`

HasIpPairingConfiguration returns a boolean if a field has been set.

### GetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyCommonMobile) GetPromptForNicknameOnPairing() bool`

GetPromptForNicknameOnPairing returns the PromptForNicknameOnPairing field if non-nil, zero value otherwise.

### GetPromptForNicknameOnPairingOk

`func (o *DeviceAuthenticationPolicyCommonMobile) GetPromptForNicknameOnPairingOk() (*bool, bool)`

GetPromptForNicknameOnPairingOk returns a tuple with the PromptForNicknameOnPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyCommonMobile) SetPromptForNicknameOnPairing(v bool)`

SetPromptForNicknameOnPairing sets PromptForNicknameOnPairing field to given value.

### HasPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyCommonMobile) HasPromptForNicknameOnPairing() bool`

HasPromptForNicknameOnPairing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


