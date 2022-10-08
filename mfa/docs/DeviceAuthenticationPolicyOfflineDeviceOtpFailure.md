# DeviceAuthenticationPolicyOfflineDeviceOtpFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | The maximum number of times that the OTP entry can fail for a user, before they are blocked. | 
**CoolDown** | [**DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown**](DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown.md) |  | 

## Methods

### NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure(count int32, coolDown DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown, ) *DeviceAuthenticationPolicyOfflineDeviceOtpFailure`

NewDeviceAuthenticationPolicyOfflineDeviceOtpFailure instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureWithDefaults

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureWithDefaults() *DeviceAuthenticationPolicyOfflineDeviceOtpFailure`

NewDeviceAuthenticationPolicyOfflineDeviceOtpFailureWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtpFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) SetCount(v int32)`

SetCount sets Count field to given value.


### GetCoolDown

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) GetCoolDown() DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown`

GetCoolDown returns the CoolDown field if non-nil, zero value otherwise.

### GetCoolDownOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) GetCoolDownOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown, bool)`

GetCoolDownOk returns a tuple with the CoolDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoolDown

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtpFailure) SetCoolDown(v DeviceAuthenticationPolicyOfflineDeviceOtpFailureCoolDown)`

SetCoolDown sets CoolDown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


