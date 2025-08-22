# DeviceAuthenticationPolicyPingOneMFAAllOfMobile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A boolean that specifies whether the method is enabled or disabled in the policy. | 
**Otp** | [**DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp**](DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp.md) |  | 
**Applications** | Pointer to [**[]DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications**](DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications.md) |  | [optional] 
**PromptForNicknameOnPairing** | Pointer to **bool** | Set to &#x60;true&#x60; if you want to allow users to provide nicknames for devices during pairing. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobile

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobile(enabled bool, otp DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp, ) *DeviceAuthenticationPolicyPingOneMFAAllOfMobile`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobile instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileWithDefaults

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileWithDefaults() *DeviceAuthenticationPolicyPingOneMFAAllOfMobile`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileWithDefaults instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetOtp

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) GetOtp() DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) GetOtpOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) SetOtp(v DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp)`

SetOtp sets Otp field to given value.


### GetApplications

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) GetApplications() []DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) GetApplicationsOk() (*[]DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) SetApplications(v []DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) GetPromptForNicknameOnPairing() bool`

GetPromptForNicknameOnPairing returns the PromptForNicknameOnPairing field if non-nil, zero value otherwise.

### GetPromptForNicknameOnPairingOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) GetPromptForNicknameOnPairingOk() (*bool, bool)`

GetPromptForNicknameOnPairingOk returns a tuple with the PromptForNicknameOnPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) SetPromptForNicknameOnPairing(v bool)`

SetPromptForNicknameOnPairing sets PromptForNicknameOnPairing field to given value.

### HasPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobile) HasPromptForNicknameOnPairing() bool`

HasPromptForNicknameOnPairing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


