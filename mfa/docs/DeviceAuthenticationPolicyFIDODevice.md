# DeviceAuthenticationPolicyFIDODevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Enabled or disabled in the policy. | 
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


