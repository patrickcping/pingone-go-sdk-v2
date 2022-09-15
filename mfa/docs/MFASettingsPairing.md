# MFASettingsPairing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxAllowedDevices** | **int32** | An integer that defines the maximum number of MFA devices each user can have. This can be any number up to 15. The default value is 5. | [default to 5]
**PairingKeyFormat** | [**EnumMFASettingsPairingKeyFormat**](EnumMFASettingsPairingKeyFormat.md) |  | 

## Methods

### NewMFASettingsPairing

`func NewMFASettingsPairing(maxAllowedDevices int32, pairingKeyFormat EnumMFASettingsPairingKeyFormat, ) *MFASettingsPairing`

NewMFASettingsPairing instantiates a new MFASettingsPairing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFASettingsPairingWithDefaults

`func NewMFASettingsPairingWithDefaults() *MFASettingsPairing`

NewMFASettingsPairingWithDefaults instantiates a new MFASettingsPairing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxAllowedDevices

`func (o *MFASettingsPairing) GetMaxAllowedDevices() int32`

GetMaxAllowedDevices returns the MaxAllowedDevices field if non-nil, zero value otherwise.

### GetMaxAllowedDevicesOk

`func (o *MFASettingsPairing) GetMaxAllowedDevicesOk() (*int32, bool)`

GetMaxAllowedDevicesOk returns a tuple with the MaxAllowedDevices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAllowedDevices

`func (o *MFASettingsPairing) SetMaxAllowedDevices(v int32)`

SetMaxAllowedDevices sets MaxAllowedDevices field to given value.


### GetPairingKeyFormat

`func (o *MFASettingsPairing) GetPairingKeyFormat() EnumMFASettingsPairingKeyFormat`

GetPairingKeyFormat returns the PairingKeyFormat field if non-nil, zero value otherwise.

### GetPairingKeyFormatOk

`func (o *MFASettingsPairing) GetPairingKeyFormatOk() (*EnumMFASettingsPairingKeyFormat, bool)`

GetPairingKeyFormatOk returns a tuple with the PairingKeyFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingKeyFormat

`func (o *MFASettingsPairing) SetPairingKeyFormat(v EnumMFASettingsPairingKeyFormat)`

SetPairingKeyFormat sets PairingKeyFormat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


