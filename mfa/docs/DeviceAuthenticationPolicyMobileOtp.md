# DeviceAuthenticationPolicyMobileOtp

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Failure** | [**DeviceAuthenticationPolicyOfflineDeviceOtpFailure**](DeviceAuthenticationPolicyOfflineDeviceOtpFailure.md) |  | 
**Window** | Pointer to [**DeviceAuthenticationPolicyMobileOtpWindow**](DeviceAuthenticationPolicyMobileOtpWindow.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyMobileOtp

`func NewDeviceAuthenticationPolicyMobileOtp(failure DeviceAuthenticationPolicyOfflineDeviceOtpFailure, ) *DeviceAuthenticationPolicyMobileOtp`

NewDeviceAuthenticationPolicyMobileOtp instantiates a new DeviceAuthenticationPolicyMobileOtp object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyMobileOtpWithDefaults

`func NewDeviceAuthenticationPolicyMobileOtpWithDefaults() *DeviceAuthenticationPolicyMobileOtp`

NewDeviceAuthenticationPolicyMobileOtpWithDefaults instantiates a new DeviceAuthenticationPolicyMobileOtp object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailure

`func (o *DeviceAuthenticationPolicyMobileOtp) GetFailure() DeviceAuthenticationPolicyOfflineDeviceOtpFailure`

GetFailure returns the Failure field if non-nil, zero value otherwise.

### GetFailureOk

`func (o *DeviceAuthenticationPolicyMobileOtp) GetFailureOk() (*DeviceAuthenticationPolicyOfflineDeviceOtpFailure, bool)`

GetFailureOk returns a tuple with the Failure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailure

`func (o *DeviceAuthenticationPolicyMobileOtp) SetFailure(v DeviceAuthenticationPolicyOfflineDeviceOtpFailure)`

SetFailure sets Failure field to given value.


### GetWindow

`func (o *DeviceAuthenticationPolicyMobileOtp) GetWindow() DeviceAuthenticationPolicyMobileOtpWindow`

GetWindow returns the Window field if non-nil, zero value otherwise.

### GetWindowOk

`func (o *DeviceAuthenticationPolicyMobileOtp) GetWindowOk() (*DeviceAuthenticationPolicyMobileOtpWindow, bool)`

GetWindowOk returns a tuple with the Window field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindow

`func (o *DeviceAuthenticationPolicyMobileOtp) SetWindow(v DeviceAuthenticationPolicyMobileOtpWindow)`

SetWindow sets Window field to given value.

### HasWindow

`func (o *DeviceAuthenticationPolicyMobileOtp) HasWindow() bool`

HasWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


