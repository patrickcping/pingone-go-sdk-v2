# DeviceAuthenticationPolicyMigrateData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceAuthenticationPolicyId** | **string** | The ID of the device authentication policy. | 
**Fido2PolicyId** | Pointer to **string** | The ID of the FIDO2 policy. | [optional] 

## Methods

### NewDeviceAuthenticationPolicyMigrateData

`func NewDeviceAuthenticationPolicyMigrateData(deviceAuthenticationPolicyId string, ) *DeviceAuthenticationPolicyMigrateData`

NewDeviceAuthenticationPolicyMigrateData instantiates a new DeviceAuthenticationPolicyMigrateData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyMigrateDataWithDefaults

`func NewDeviceAuthenticationPolicyMigrateDataWithDefaults() *DeviceAuthenticationPolicyMigrateData`

NewDeviceAuthenticationPolicyMigrateDataWithDefaults instantiates a new DeviceAuthenticationPolicyMigrateData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceAuthenticationPolicyId

`func (o *DeviceAuthenticationPolicyMigrateData) GetDeviceAuthenticationPolicyId() string`

GetDeviceAuthenticationPolicyId returns the DeviceAuthenticationPolicyId field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPolicyIdOk

`func (o *DeviceAuthenticationPolicyMigrateData) GetDeviceAuthenticationPolicyIdOk() (*string, bool)`

GetDeviceAuthenticationPolicyIdOk returns a tuple with the DeviceAuthenticationPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicyId

`func (o *DeviceAuthenticationPolicyMigrateData) SetDeviceAuthenticationPolicyId(v string)`

SetDeviceAuthenticationPolicyId sets DeviceAuthenticationPolicyId field to given value.


### GetFido2PolicyId

`func (o *DeviceAuthenticationPolicyMigrateData) GetFido2PolicyId() string`

GetFido2PolicyId returns the Fido2PolicyId field if non-nil, zero value otherwise.

### GetFido2PolicyIdOk

`func (o *DeviceAuthenticationPolicyMigrateData) GetFido2PolicyIdOk() (*string, bool)`

GetFido2PolicyIdOk returns a tuple with the Fido2PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFido2PolicyId

`func (o *DeviceAuthenticationPolicyMigrateData) SetFido2PolicyId(v string)`

SetFido2PolicyId sets Fido2PolicyId field to given value.

### HasFido2PolicyId

`func (o *DeviceAuthenticationPolicyMigrateData) HasFido2PolicyId() bool`

HasFido2PolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


