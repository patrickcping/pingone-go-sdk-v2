# DeviceAuthenticationPolicyOfflineDeviceOtp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LifeTime** | [**DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime**](DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime.md) |  | 
**Failure** | [**DeviceAuthenticationPolicyOfflineDeviceOtpFailure**](DeviceAuthenticationPolicyOfflineDeviceOtpFailure.md) |  | 
**OtpLength** | Pointer to **int32** | Used to specify the length of the OTP that is shown to users. Minimum length is &#x60;6&#x60; digits and maximum is &#x60;10&#x60; digits. If the parameter is not provided, the default is &#x60;6&#x60; digits. | [optional] [default to 6]

## Methods

### NewDeviceAuthenticationPolicyOfflineDeviceOtp

`func NewDeviceAuthenticationPolicyOfflineDeviceOtp(lifeTime DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime, failure DeviceAuthenticationPolicyOfflineDeviceOtpFailure, ) *DeviceAuthenticationPolicyOfflineDeviceOtp`

NewDeviceAuthenticationPolicyOfflineDeviceOtp instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyOfflineDeviceOtpWithDefaults

`func NewDeviceAuthenticationPolicyOfflineDeviceOtpWithDefaults() *DeviceAuthenticationPolicyOfflineDeviceOtp`

NewDeviceAuthenticationPolicyOfflineDeviceOtpWithDefaults instantiates a new DeviceAuthenticationPolicyOfflineDeviceOtp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLifeTime

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetLifeTime() DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime`

GetLifeTime returns the LifeTime field if non-nil, zero value otherwise.

### GetLifeTimeOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetLifeTimeOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime, bool)`

GetLifeTimeOk returns a tuple with the LifeTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLifeTime

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) SetLifeTime(v DeviceAuthenticationPolicyOfflineDeviceOtpLifeTime)`

SetLifeTime sets LifeTime field to given value.


### GetFailure

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetFailure() DeviceAuthenticationPolicyOfflineDeviceOtpFailure`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetFailureOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpFailure, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) SetFailure(v DeviceAuthenticationPolicyOfflineDeviceOtpFailure)`

SetFailure sets Failure field to given value.


### GetOtpLength

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetOtpLength() int32`

GetOtpLength returns the OtpLength field if non-nil, zero value otherwise.

### GetOtpLengthOk

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) GetOtpLengthOk() (*int32, bool)`

GetOtpLengthOk returns a tuple with the OtpLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtpLength

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) SetOtpLength(v int32)`

SetOtpLength sets OtpLength field to given value.

### HasOtpLength

`func (o *DeviceAuthenticationPolicyOfflineDeviceOtp) HasOtpLength() bool`

HasOtpLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


