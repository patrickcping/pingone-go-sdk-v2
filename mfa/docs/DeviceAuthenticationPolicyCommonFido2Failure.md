# DeviceAuthenticationPolicyCommonFido2Failure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int32** | The maximum number of times that authentication can fail before user is blocked for the specified period. Minimum is &#x60;1&#x60;, maximum is &#x60;7&#x60;. | [optional] 
**CoolDown** | Pointer to [**DeviceAuthenticationPolicyCommonFido2FailureCoolDown**](DeviceAuthenticationPolicyCommonFido2FailureCoolDown.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyCommonFido2Failure

`func NewDeviceAuthenticationPolicyCommonFido2Failure() *DeviceAuthenticationPolicyCommonFido2Failure`

NewDeviceAuthenticationPolicyCommonFido2Failure instantiates a new DeviceAuthenticationPolicyCommonFido2Failure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonFido2FailureWithDefaults

`func NewDeviceAuthenticationPolicyCommonFido2FailureWithDefaults() *DeviceAuthenticationPolicyCommonFido2Failure`

NewDeviceAuthenticationPolicyCommonFido2FailureWithDefaults instantiates a new DeviceAuthenticationPolicyCommonFido2Failure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DeviceAuthenticationPolicyCommonFido2Failure) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeviceAuthenticationPolicyCommonFido2Failure) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeviceAuthenticationPolicyCommonFido2Failure) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DeviceAuthenticationPolicyCommonFido2Failure) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCoolDown

`func (o *DeviceAuthenticationPolicyCommonFido2Failure) GetCoolDown() DeviceAuthenticationPolicyCommonFido2FailureCoolDown`

GetCoolDown returns the CoolDown field if non-nil, zero value otherwise.

### GetCoolDownOk

`func (o *DeviceAuthenticationPolicyCommonFido2Failure) GetCoolDownOk() (*DeviceAuthenticationPolicyCommonFido2FailureCoolDown, bool)`

GetCoolDownOk returns a tuple with the CoolDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolDown

`func (o *DeviceAuthenticationPolicyCommonFido2Failure) SetCoolDown(v DeviceAuthenticationPolicyCommonFido2FailureCoolDown)`

SetCoolDown sets CoolDown field to given value.

### HasCoolDown

`func (o *DeviceAuthenticationPolicyCommonFido2Failure) HasCoolDown() bool`

HasCoolDown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


