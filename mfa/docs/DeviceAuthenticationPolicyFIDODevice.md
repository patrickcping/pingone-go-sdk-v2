# DeviceAuthenticationPolicyFIDODevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enabled or disabled in the policy. | 
**PairingDisabled** | Pointer to **bool** | You can set &#x60;pairingDisabled&#x60; to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices. | [optional] 
**FidoPolicyId** | Pointer to **string** | Specifies the FIDO policy UUID. This property can be null. When null, the environment&#39;s default FIDO Policy is used. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyFIDODevice

`func NewDeviceAuthenticationPolicyFIDODevice(enabled bool, ) *DeviceAuthenticationPolicyFIDODevice`

NewDeviceAuthenticationPolicyFIDODevice instantiates a new DeviceAuthenticationPolicyFIDODevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyFIDODeviceWithDefaults

`func NewDeviceAuthenticationPolicyFIDODeviceWithDefaults() *DeviceAuthenticationPolicyFIDODevice`

NewDeviceAuthenticationPolicyFIDODeviceWithDefaults instantiates a new DeviceAuthenticationPolicyFIDODevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *DeviceAuthenticationPolicyFIDODevice) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceAuthenticationPolicyFIDODevice) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceAuthenticationPolicyFIDODevice) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetPairingDisabled

`func (o *DeviceAuthenticationPolicyFIDODevice) GetPairingDisabled() bool`

GetPairingDisabled returns the PairingDisabled field if non-nil, zero value otherwise.

### GetPairingDisabledOk

`func (o *DeviceAuthenticationPolicyFIDODevice) GetPairingDisabledOk() (*bool, bool)`

GetPairingDisabledOk returns a tuple with the PairingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingDisabled

`func (o *DeviceAuthenticationPolicyFIDODevice) SetPairingDisabled(v bool)`

SetPairingDisabled sets PairingDisabled field to given value.

### HasPairingDisabled

`func (o *DeviceAuthenticationPolicyFIDODevice) HasPairingDisabled() bool`

HasPairingDisabled returns a boolean if a field has been set.

### GetFidoPolicyId

`func (o *DeviceAuthenticationPolicyFIDODevice) GetFidoPolicyId() string`

GetFidoPolicyId returns the FidoPolicyId field if non-nil, zero value otherwise.

### GetFidoPolicyIdOk

`func (o *DeviceAuthenticationPolicyFIDODevice) GetFidoPolicyIdOk() (*string, bool)`

GetFidoPolicyIdOk returns a tuple with the FidoPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFidoPolicyId

`func (o *DeviceAuthenticationPolicyFIDODevice) SetFidoPolicyId(v string)`

SetFidoPolicyId sets FidoPolicyId field to given value.

### HasFidoPolicyId

`func (o *DeviceAuthenticationPolicyFIDODevice) HasFidoPolicyId() bool`

HasFidoPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


