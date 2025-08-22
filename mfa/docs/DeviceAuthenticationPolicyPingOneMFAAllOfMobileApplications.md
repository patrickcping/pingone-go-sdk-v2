# DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The application&#39;s ID. | 
**Push** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobilePush**](DeviceAuthenticationPolicyPingOneMFAAllOfMobilePush.md) |  | [optional] 
**PushTimeout** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout**](DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout.md) |  | [optional] 
**PairingKeyLifetime** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime**](DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime.md) |  | [optional] 
**PushLimit** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit**](DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit.md) |  | [optional] 
**Otp** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1**](DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1.md) |  | [optional] 
**DeviceAuthorization** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization**](DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization.md) |  | [optional] 
**AutoEnrollment** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment**](DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment.md) |  | [optional] 
**PairingDisabled** | Pointer to **bool** | You can set &#x60;pairingDisabled&#x60; to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices. | [optional] 
**IntegrityDetection** | Pointer to [**EnumMFADevicePolicyMobileIntegrityDetection**](EnumMFADevicePolicyMobileIntegrityDetection.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications(id string, ) *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileApplicationsWithDefaults

`func NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileApplicationsWithDefaults() *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications`

NewDeviceAuthenticationPolicyPingOneMFAAllOfMobileApplicationsWithDefaults instantiates a new DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetId(v string)`

SetId sets Id field to given value.


### GetPush

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPush() DeviceAuthenticationPolicyPingOneMFAAllOfMobilePush`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPushOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobilePush, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetPush(v DeviceAuthenticationPolicyPingOneMFAAllOfMobilePush)`

SetPush sets Push field to given value.

### HasPush

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasPush() bool`

HasPush returns a boolean if a field has been set.

### GetPushTimeout

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPushTimeout() DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout`

GetPushTimeout returns the PushTimeout field if non-nil, zero value otherwise.

### GetPushTimeoutOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPushTimeoutOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout, bool)`

GetPushTimeoutOk returns a tuple with the PushTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTimeout

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetPushTimeout(v DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushTimeout)`

SetPushTimeout sets PushTimeout field to given value.

### HasPushTimeout

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasPushTimeout() bool`

HasPushTimeout returns a boolean if a field has been set.

### GetPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPairingKeyLifetime() DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime`

GetPairingKeyLifetime returns the PairingKeyLifetime field if non-nil, zero value otherwise.

### GetPairingKeyLifetimeOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPairingKeyLifetimeOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime, bool)`

GetPairingKeyLifetimeOk returns a tuple with the PairingKeyLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetPairingKeyLifetime(v DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime)`

SetPairingKeyLifetime sets PairingKeyLifetime field to given value.

### HasPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasPairingKeyLifetime() bool`

HasPairingKeyLifetime returns a boolean if a field has been set.

### GetPushLimit

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPushLimit() DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit`

GetPushLimit returns the PushLimit field if non-nil, zero value otherwise.

### GetPushLimitOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPushLimitOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit, bool)`

GetPushLimitOk returns a tuple with the PushLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushLimit

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetPushLimit(v DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit)`

SetPushLimit sets PushLimit field to given value.

### HasPushLimit

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasPushLimit() bool`

HasPushLimit returns a boolean if a field has been set.

### GetOtp

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetOtp() DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetOtpOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetOtp(v DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetDeviceAuthorization

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetDeviceAuthorization() DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization`

GetDeviceAuthorization returns the DeviceAuthorization field if non-nil, zero value otherwise.

### GetDeviceAuthorizationOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetDeviceAuthorizationOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization, bool)`

GetDeviceAuthorizationOk returns a tuple with the DeviceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorization

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetDeviceAuthorization(v DeviceAuthenticationPolicyPingOneMFAAllOfMobileDeviceAuthorization)`

SetDeviceAuthorization sets DeviceAuthorization field to given value.

### HasDeviceAuthorization

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasDeviceAuthorization() bool`

HasDeviceAuthorization returns a boolean if a field has been set.

### GetAutoEnrollment

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetAutoEnrollment() DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment`

GetAutoEnrollment returns the AutoEnrollment field if non-nil, zero value otherwise.

### GetAutoEnrollmentOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetAutoEnrollmentOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment, bool)`

GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollment

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetAutoEnrollment(v DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment)`

SetAutoEnrollment sets AutoEnrollment field to given value.

### HasAutoEnrollment

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasAutoEnrollment() bool`

HasAutoEnrollment returns a boolean if a field has been set.

### GetPairingDisabled

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPairingDisabled() bool`

GetPairingDisabled returns the PairingDisabled field if non-nil, zero value otherwise.

### GetPairingDisabledOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetPairingDisabledOk() (*bool, bool)`

GetPairingDisabledOk returns a tuple with the PairingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingDisabled

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetPairingDisabled(v bool)`

SetPairingDisabled sets PairingDisabled field to given value.

### HasPairingDisabled

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasPairingDisabled() bool`

HasPairingDisabled returns a boolean if a field has been set.

### GetIntegrityDetection

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetIntegrityDetection() EnumMFADevicePolicyMobileIntegrityDetection`

GetIntegrityDetection returns the IntegrityDetection field if non-nil, zero value otherwise.

### GetIntegrityDetectionOk

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) GetIntegrityDetectionOk() (*EnumMFADevicePolicyMobileIntegrityDetection, bool)`

GetIntegrityDetectionOk returns a tuple with the IntegrityDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityDetection

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) SetIntegrityDetection(v EnumMFADevicePolicyMobileIntegrityDetection)`

SetIntegrityDetection sets IntegrityDetection field to given value.

### HasIntegrityDetection

`func (o *DeviceAuthenticationPolicyPingOneMFAAllOfMobileApplications) HasIntegrityDetection() bool`

HasIntegrityDetection returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


