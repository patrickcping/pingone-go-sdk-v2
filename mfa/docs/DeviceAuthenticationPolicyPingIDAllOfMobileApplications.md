# DeviceAuthenticationPolicyPingIDAllOfMobileApplications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The application&#39;s ID. | 
**Type** | [**EnumPingIDApplicationType**](EnumPingIDApplicationType.md) |  | 
**Push** | Pointer to [**DeviceAuthenticationPolicyPingIDAllOfMobilePush**](DeviceAuthenticationPolicyPingIDAllOfMobilePush.md) |  | [optional] 
**PairingKeyLifetime** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime**](DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime.md) |  | [optional] 
**PushLimit** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit**](DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit.md) |  | [optional] 
**Otp** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1**](DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1.md) |  | [optional] 
**AutoEnrollment** | Pointer to [**DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment**](DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment.md) |  | [optional] 
**BiometricsEnabled** | Pointer to **bool** | You can set &#x60;biometricsEnabled&#x60; to true to allow users to use biometric authentication methods (such as fingerprint or facial recognition) for MFA. If set to false, users will not be able to use biometric methods for authentication. | [optional] 
**NewRequestDurationConfiguration** | Pointer to [**DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfiguration**](DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfiguration.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyPingIDAllOfMobileApplications

`func NewDeviceAuthenticationPolicyPingIDAllOfMobileApplications(id string, type_ EnumPingIDApplicationType, ) *DeviceAuthenticationPolicyPingIDAllOfMobileApplications`

NewDeviceAuthenticationPolicyPingIDAllOfMobileApplications instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobileApplications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyPingIDAllOfMobileApplicationsWithDefaults

`func NewDeviceAuthenticationPolicyPingIDAllOfMobileApplicationsWithDefaults() *DeviceAuthenticationPolicyPingIDAllOfMobileApplications`

NewDeviceAuthenticationPolicyPingIDAllOfMobileApplicationsWithDefaults instantiates a new DeviceAuthenticationPolicyPingIDAllOfMobileApplications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetType() EnumPingIDApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetTypeOk() (*EnumPingIDApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetType(v EnumPingIDApplicationType)`

SetType sets Type field to given value.


### GetPush

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetPush() DeviceAuthenticationPolicyPingIDAllOfMobilePush`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetPushOk() (*DeviceAuthenticationPolicyPingIDAllOfMobilePush, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetPush(v DeviceAuthenticationPolicyPingIDAllOfMobilePush)`

SetPush sets Push field to given value.

### HasPush

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) HasPush() bool`

HasPush returns a boolean if a field has been set.

### GetPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetPairingKeyLifetime() DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime`

GetPairingKeyLifetime returns the PairingKeyLifetime field if non-nil, zero value otherwise.

### GetPairingKeyLifetimeOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetPairingKeyLifetimeOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime, bool)`

GetPairingKeyLifetimeOk returns a tuple with the PairingKeyLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetPairingKeyLifetime(v DeviceAuthenticationPolicyPingOneMFAAllOfMobilePairingKeyLifetime)`

SetPairingKeyLifetime sets PairingKeyLifetime field to given value.

### HasPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) HasPairingKeyLifetime() bool`

HasPairingKeyLifetime returns a boolean if a field has been set.

### GetPushLimit

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetPushLimit() DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit`

GetPushLimit returns the PushLimit field if non-nil, zero value otherwise.

### GetPushLimitOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetPushLimitOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit, bool)`

GetPushLimitOk returns a tuple with the PushLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushLimit

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetPushLimit(v DeviceAuthenticationPolicyPingOneMFAAllOfMobilePushLimit)`

SetPushLimit sets PushLimit field to given value.

### HasPushLimit

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) HasPushLimit() bool`

HasPushLimit returns a boolean if a field has been set.

### GetOtp

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetOtp() DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetOtpOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetOtp(v DeviceAuthenticationPolicyPingOneMFAAllOfMobileOtp1)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetAutoEnrollment

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetAutoEnrollment() DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment`

GetAutoEnrollment returns the AutoEnrollment field if non-nil, zero value otherwise.

### GetAutoEnrollmentOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetAutoEnrollmentOk() (*DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment, bool)`

GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollment

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetAutoEnrollment(v DeviceAuthenticationPolicyPingOneMFAAllOfMobileAutoEnrollment)`

SetAutoEnrollment sets AutoEnrollment field to given value.

### HasAutoEnrollment

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) HasAutoEnrollment() bool`

HasAutoEnrollment returns a boolean if a field has been set.

### GetBiometricsEnabled

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetBiometricsEnabled() bool`

GetBiometricsEnabled returns the BiometricsEnabled field if non-nil, zero value otherwise.

### GetBiometricsEnabledOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetBiometricsEnabledOk() (*bool, bool)`

GetBiometricsEnabledOk returns a tuple with the BiometricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricsEnabled

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetBiometricsEnabled(v bool)`

SetBiometricsEnabled sets BiometricsEnabled field to given value.

### HasBiometricsEnabled

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) HasBiometricsEnabled() bool`

HasBiometricsEnabled returns a boolean if a field has been set.

### GetNewRequestDurationConfiguration

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetNewRequestDurationConfiguration() DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfiguration`

GetNewRequestDurationConfiguration returns the NewRequestDurationConfiguration field if non-nil, zero value otherwise.

### GetNewRequestDurationConfigurationOk

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) GetNewRequestDurationConfigurationOk() (*DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfiguration, bool)`

GetNewRequestDurationConfigurationOk returns a tuple with the NewRequestDurationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRequestDurationConfiguration

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) SetNewRequestDurationConfiguration(v DeviceAuthenticationPolicyPingIDAllOfMobileNewRequestDurationConfiguration)`

SetNewRequestDurationConfiguration sets NewRequestDurationConfiguration field to given value.

### HasNewRequestDurationConfiguration

`func (o *DeviceAuthenticationPolicyPingIDAllOfMobileApplications) HasNewRequestDurationConfiguration() bool`

HasNewRequestDurationConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


