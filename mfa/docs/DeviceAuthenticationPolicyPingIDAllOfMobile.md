# DeviceAuthenticationPolicyPingIDAllOfMobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A boolean that specifies whether the method is enabled or disabled in the policy. | 
**Otp** | [**DeviceAuthenticationPolicyPingIDAllOfMobileOtp**](DeviceAuthenticationPolicyPingIDAllOfMobileOtp.md) |  | 
**Applications** | Pointer to [**[]DeviceAuthenticationPolicyPingIDAllOfMobileApplications**](DeviceAuthenticationPolicyPingIDAllOfMobileApplications.md) |  | [optional] 
**IpPairingConfiguration** | Pointer to [**DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration**](DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration.md) |  | [optional] 
**PromptForNicknameOnPairing** | Pointer to **bool** | Set to &#x60;true&#x60; if you want to allow users to provide nicknames for devices during pairing. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingIDAllOfMobile

`func NewDeviceAuthenticationPolicyPingIDAllOfMobile(enabled bool, otp DeviceAuthenticationPolicyPingIDAllOfMobileOtp, ) *DeviceAuthenticationPolicyPingIDAllOfMobile`

NewDeviceAuthenticationPolicyPingIDAllOfMobile instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingIDAllOfMobileWithDefaults

`func NewDeviceAuthenticationPolicyPingIDAllOfMobileWithDefaults() *DeviceAuthenticationPolicyPingIDAllOfMobile`

NewDeviceAuthenticationPolicyPingIDAllOfMobileWithDefaults instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOtp

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetOtp() DeviceAuthenticationPolicyPingIDAllOfMobileOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetOtpOk() (*DeviceAuthenticationPolicyPingIDAllOfMobileOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) SetOtp(v DeviceAuthenticationPolicyPingIDAllOfMobileOtp)`

SetOtp sets Otp field to given value.


### GetApplications

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetApplications() []DeviceAuthenticationPolicyPingIDAllOfMobileApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetApplicationsOk() (*[]DeviceAuthenticationPolicyPingIDAllOfMobileApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) SetApplications(v []DeviceAuthenticationPolicyPingIDAllOfMobileApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetIpPairingConfiguration() DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration`

GetIpPairingConfiguration returns the IpPairingConfiguration field if non-nil, zero value otherwise.

### GetIpPairingConfigurationOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetIpPairingConfigurationOk() (*DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration, bool)`

GetIpPairingConfigurationOk returns a tuple with the IpPairingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) SetIpPairingConfiguration(v DeviceAuthenticationPolicyPingIDAllOfMobileIpPairingConfiguration)`

SetIpPairingConfiguration sets IpPairingConfiguration field to given value.

### HasIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) HasIpPairingConfiguration() bool`

HasIpPairingConfiguration returns a boolean if a field has been set.

### GetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetPromptForNicknameOnPairing() bool`

GetPromptForNicknameOnPairing returns the PromptForNicknameOnPairing field if non-nil, zero value otherwise.

### GetPromptForNicknameOnPairingOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) GetPromptForNicknameOnPairingOk() (*bool, bool)`

GetPromptForNicknameOnPairingOk returns a tuple with the PromptForNicknameOnPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) SetPromptForNicknameOnPairing(v bool)`

SetPromptForNicknameOnPairing sets PromptForNicknameOnPairing field to given value.

### HasPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobile) HasPromptForNicknameOnPairing() bool`

HasPromptForNicknameOnPairing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


