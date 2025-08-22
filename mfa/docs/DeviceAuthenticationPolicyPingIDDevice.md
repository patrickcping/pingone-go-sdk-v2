# DeviceAuthenticationPolicyPingIDDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enabled or disabled in the policy. | 
**PairingDisabled** | Pointer to **bool** | You can set &#x60;pairingDisabled&#x60; to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices. | [optional] 
**Otp** | [**DeviceAuthenticationPolicyPingIDDeviceOtp**](DeviceAuthenticationPolicyPingIDDeviceOtp.md) |  | 
**PromptForNicknameOnPairing** | Pointer to **bool** | Set to &#x60;true&#x60; if you want to allow users to provide nicknames for devices during pairing. | [optional] 
**PairingKeyLifetime** | Pointer to [**DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime**](DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingIDDevice

`func NewDeviceAuthenticationPolicyPingIDDevice(enabled bool, otp DeviceAuthenticationPolicyPingIDDeviceOtp, ) *DeviceAuthenticationPolicyPingIDDevice`

NewDeviceAuthenticationPolicyPingIDDevice instantiates a new DeviceAuthenticationPolicyPingIDDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingIDDeviceWithDefaults

`func NewDeviceAuthenticationPolicyPingIDDeviceWithDefaults() *DeviceAuthenticationPolicyPingIDDevice`

NewDeviceAuthenticationPolicyPingIDDeviceWithDefaults instantiates a new DeviceAuthenticationPolicyPingIDDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyPingIDDevice) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPairingDisabled

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetPairingDisabled() bool`

GetPairingDisabled returns the PairingDisabled field if non-nil, zero value otherwise.

### GetPairingDisabledOk

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetPairingDisabledOk() (*bool, bool)`

GetPairingDisabledOk returns a tuple with the PairingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingDisabled

`func (o *DeviceAuthenticationPolicyPingIDDevice) SetPairingDisabled(v bool)`

SetPairingDisabled sets PairingDisabled field to given value.

### HasPairingDisabled

`func (o *DeviceAuthenticationPolicyPingIDDevice) HasPairingDisabled() bool`

HasPairingDisabled returns a boolean if a field has been set.

### GetOtp

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetOtp() DeviceAuthenticationPolicyPingIDDeviceOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetOtpOk() (*DeviceAuthenticationPolicyPingIDDeviceOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyPingIDDevice) SetOtp(v DeviceAuthenticationPolicyPingIDDeviceOtp)`

SetOtp sets Otp field to given value.


### GetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetPromptForNicknameOnPairing() bool`

GetPromptForNicknameOnPairing returns the PromptForNicknameOnPairing field if non-nil, zero value otherwise.

### GetPromptForNicknameOnPairingOk

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetPromptForNicknameOnPairingOk() (*bool, bool)`

GetPromptForNicknameOnPairingOk returns a tuple with the PromptForNicknameOnPairing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingIDDevice) SetPromptForNicknameOnPairing(v bool)`

SetPromptForNicknameOnPairing sets PromptForNicknameOnPairing field to given value.

### HasPromptForNicknameOnPairing

`func (o *DeviceAuthenticationPolicyPingIDDevice) HasPromptForNicknameOnPairing() bool`

HasPromptForNicknameOnPairing returns a boolean if a field has been set.

### GetPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetPairingKeyLifetime() DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime`

GetPairingKeyLifetime returns the PairingKeyLifetime field if non-nil, zero value otherwise.

### GetPairingKeyLifetimeOk

`func (o *DeviceAuthenticationPolicyPingIDDevice) GetPairingKeyLifetimeOk() (*DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime, bool)`

GetPairingKeyLifetimeOk returns a tuple with the PairingKeyLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingIDDevice) SetPairingKeyLifetime(v DeviceAuthenticationPolicyPingIDDevicePairingKeyLifetime)`

SetPairingKeyLifetime sets PairingKeyLifetime field to given value.

### HasPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingIDDevice) HasPairingKeyLifetime() bool`

HasPairingKeyLifetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


