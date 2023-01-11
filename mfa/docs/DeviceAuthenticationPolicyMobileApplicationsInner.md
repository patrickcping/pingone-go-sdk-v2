# DeviceAuthenticationPolicyMobileApplicationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The application&#39;s ID. | 
**Push** | Pointer to [**DeviceAuthenticationPolicyMobileApplicationsInnerPush**](DeviceAuthenticationPolicyMobileApplicationsInnerPush.md) |  | [optional] 
**PushTimeout** | Pointer to [**DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout**](DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout.md) |  | [optional] 
**Otp** | Pointer to [**DeviceAuthenticationPolicyMobileApplicationsInnerOtp**](DeviceAuthenticationPolicyMobileApplicationsInnerOtp.md) |  | [optional] 
**DeviceAuthorization** | Pointer to [**DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization**](DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization.md) |  | [optional] 
**AutoEnrollment** | Pointer to [**DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment**](DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment.md) |  | [optional] 
**IntegrityDetection** | Pointer to [**EnumMFADevicePolicyMobileIntegrityDetection**](EnumMFADevicePolicyMobileIntegrityDetection.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyMobileApplicationsInner

`func NewDeviceAuthenticationPolicyMobileApplicationsInner(id string, ) *DeviceAuthenticationPolicyMobileApplicationsInner`

NewDeviceAuthenticationPolicyMobileApplicationsInner instantiates a new DeviceAuthenticationPolicyMobileApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyMobileApplicationsInnerWithDefaults

`func NewDeviceAuthenticationPolicyMobileApplicationsInnerWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInner`

NewDeviceAuthenticationPolicyMobileApplicationsInnerWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetId(v string)`

SetId sets Id field to given value.


### GetPush

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPush() DeviceAuthenticationPolicyMobileApplicationsInnerPush`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPushOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerPush, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetPush(v DeviceAuthenticationPolicyMobileApplicationsInnerPush)`

SetPush sets Push field to given value.

### HasPush

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasPush() bool`

HasPush returns a boolean if a field has been set.

### GetPushTimeout

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPushTimeout() DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout`

GetPushTimeout returns the PushTimeout field if non-nil, zero value otherwise.

### GetPushTimeoutOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPushTimeoutOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout, bool)`

GetPushTimeoutOk returns a tuple with the PushTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTimeout

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetPushTimeout(v DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout)`

SetPushTimeout sets PushTimeout field to given value.

### HasPushTimeout

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasPushTimeout() bool`

HasPushTimeout returns a boolean if a field has been set.

### GetOtp

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetOtp() DeviceAuthenticationPolicyMobileApplicationsInnerOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetOtpOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetOtp(v DeviceAuthenticationPolicyMobileApplicationsInnerOtp)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetDeviceAuthorization

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetDeviceAuthorization() DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization`

GetDeviceAuthorization returns the DeviceAuthorization field if non-nil, zero value otherwise.

### GetDeviceAuthorizationOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetDeviceAuthorizationOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization, bool)`

GetDeviceAuthorizationOk returns a tuple with the DeviceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorization

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetDeviceAuthorization(v DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization)`

SetDeviceAuthorization sets DeviceAuthorization field to given value.

### HasDeviceAuthorization

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasDeviceAuthorization() bool`

HasDeviceAuthorization returns a boolean if a field has been set.

### GetAutoEnrollment

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetAutoEnrollment() DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment`

GetAutoEnrollment returns the AutoEnrollment field if non-nil, zero value otherwise.

### GetAutoEnrollmentOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetAutoEnrollmentOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment, bool)`

GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollment

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetAutoEnrollment(v DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment)`

SetAutoEnrollment sets AutoEnrollment field to given value.

### HasAutoEnrollment

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasAutoEnrollment() bool`

HasAutoEnrollment returns a boolean if a field has been set.

### GetIntegrityDetection

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetIntegrityDetection() EnumMFADevicePolicyMobileIntegrityDetection`

GetIntegrityDetection returns the IntegrityDetection field if non-nil, zero value otherwise.

### GetIntegrityDetectionOk

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetIntegrityDetectionOk() (*EnumMFADevicePolicyMobileIntegrityDetection, bool)`

GetIntegrityDetectionOk returns a tuple with the IntegrityDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityDetection

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetIntegrityDetection(v EnumMFADevicePolicyMobileIntegrityDetection)`

SetIntegrityDetection sets IntegrityDetection field to given value.

### HasIntegrityDetection

`func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasIntegrityDetection() bool`

HasIntegrityDetection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


