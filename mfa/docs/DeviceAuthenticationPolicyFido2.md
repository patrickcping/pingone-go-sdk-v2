# DeviceAuthenticationPolicyFido2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | A boolean that specifies whether the method is enabled or disabled in the policy. | 
**PairingDisabled** | Pointer to **bool** | You can set &#x60;pairingDisabled&#x60; to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices. | [optional] 
**Fido2PolicyId** | Pointer to **string** | Specifies the UUID that represents the FIDO2 policy in PingOne. This property can be null. When null, the environment&#39;s default FIDO2 Policy is used. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyFido2

`func NewDeviceAuthenticationPolicyFido2(enabled bool, ) *DeviceAuthenticationPolicyFido2`

NewDeviceAuthenticationPolicyFido2 instantiates a new DeviceAuthenticationPolicyFido2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyFido2WithDefaults

`func NewDeviceAuthenticationPolicyFido2WithDefaults() *DeviceAuthenticationPolicyFido2`

NewDeviceAuthenticationPolicyFido2WithDefaults instantiates a new DeviceAuthenticationPolicyFido2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyFido2) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyFido2) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyFido2) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPairingDisabled

`func (o *DeviceAuthenticationPolicyFido2) GetPairingDisabled() bool`

GetPairingDisabled returns the PairingDisabled field if non-nil, zero value otherwise.

### GetPairingDisabledOk

`func (o *DeviceAuthenticationPolicyFido2) GetPairingDisabledOk() (*bool, bool)`

GetPairingDisabledOk returns a tuple with the PairingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingDisabled

`func (o *DeviceAuthenticationPolicyFido2) SetPairingDisabled(v bool)`

SetPairingDisabled sets PairingDisabled field to given value.

### HasPairingDisabled

`func (o *DeviceAuthenticationPolicyFido2) HasPairingDisabled() bool`

HasPairingDisabled returns a boolean if a field has been set.

### GetFido2PolicyId

`func (o *DeviceAuthenticationPolicyFido2) GetFido2PolicyId() string`

GetFido2PolicyId returns the Fido2PolicyId field if non-nil, zero value otherwise.

### GetFido2PolicyIdOk

`func (o *DeviceAuthenticationPolicyFido2) GetFido2PolicyIdOk() (*string, bool)`

GetFido2PolicyIdOk returns a tuple with the Fido2PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2PolicyId

`func (o *DeviceAuthenticationPolicyFido2) SetFido2PolicyId(v string)`

SetFido2PolicyId sets Fido2PolicyId field to given value.

### HasFido2PolicyId

`func (o *DeviceAuthenticationPolicyFido2) HasFido2PolicyId() bool`

HasFido2PolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


