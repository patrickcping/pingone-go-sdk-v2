# DeviceAuthenticationPolicyCommonMobileApplicationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The application&#39;s ID. | 
**Type** | Pointer to [**EnumPingIDApplicationType**](EnumPingIDApplicationType.md) |  | [optional] 
**Push** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerPush**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerPush.md) |  | [optional] 
**PushTimeout** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushTimeout**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushTimeout.md) |  | [optional] 
**PairingKeyLifetime** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime.md) |  | [optional] 
**PushLimit** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit.md) |  | [optional] 
**Otp** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerOtp**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerOtp.md) |  | [optional] 
**DeviceAuthorization** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerDeviceAuthorization**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerDeviceAuthorization.md) |  | [optional] 
**AutoEnrollment** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerAutoEnrollment**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerAutoEnrollment.md) |  | [optional] 
**BiometricsEnabled** | Pointer to **bool** | You can set &#x60;biometricsEnabled&#x60; to true to allow users to use biometric authentication methods (such as fingerprint or facial recognition) for MFA. If set to false, users will not be able to use biometric methods for authentication. | [optional] 
**PairingDisabled** | Pointer to **bool** | You can set &#x60;pairingDisabled&#x60; to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices. | [optional] 
**IntegrityDetection** | Pointer to [**EnumMFADevicePolicyMobileIntegrityDetection**](EnumMFADevicePolicyMobileIntegrityDetection.md) |  | [optional] 
**IpPairingConfiguration** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerIpPairingConfiguration**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerIpPairingConfiguration.md) |  | [optional] 
**NewRequestDurationConfiguration** | Pointer to [**DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfiguration**](DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfiguration.md) |  | [optional] 

## Methods

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInner

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInner(id string, ) *DeviceAuthenticationPolicyCommonMobileApplicationsInner`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInner instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerWithDefaults

`func NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerWithDefaults() *DeviceAuthenticationPolicyCommonMobileApplicationsInner`

NewDeviceAuthenticationPolicyCommonMobileApplicationsInnerWithDefaults instantiates a new DeviceAuthenticationPolicyCommonMobileApplicationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetType() EnumPingIDApplicationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetTypeOk() (*EnumPingIDApplicationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetType(v EnumPingIDApplicationType)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPush

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPush() DeviceAuthenticationPolicyCommonMobileApplicationsInnerPush`

GetPush returns the Push field if non-nil, zero value otherwise.

### GetPushOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPushOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerPush, bool)`

GetPushOk returns a tuple with the Push field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPush

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetPush(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerPush)`

SetPush sets Push field to given value.

### HasPush

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasPush() bool`

HasPush returns a boolean if a field has been set.

### GetPushTimeout

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPushTimeout() DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushTimeout`

GetPushTimeout returns the PushTimeout field if non-nil, zero value otherwise.

### GetPushTimeoutOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPushTimeoutOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushTimeout, bool)`

GetPushTimeoutOk returns a tuple with the PushTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushTimeout

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetPushTimeout(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushTimeout)`

SetPushTimeout sets PushTimeout field to given value.

### HasPushTimeout

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasPushTimeout() bool`

HasPushTimeout returns a boolean if a field has been set.

### GetPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPairingKeyLifetime() DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime`

GetPairingKeyLifetime returns the PairingKeyLifetime field if non-nil, zero value otherwise.

### GetPairingKeyLifetimeOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPairingKeyLifetimeOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime, bool)`

GetPairingKeyLifetimeOk returns a tuple with the PairingKeyLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetPairingKeyLifetime(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerPairingKeyLifetime)`

SetPairingKeyLifetime sets PairingKeyLifetime field to given value.

### HasPairingKeyLifetime

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasPairingKeyLifetime() bool`

HasPairingKeyLifetime returns a boolean if a field has been set.

### GetPushLimit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPushLimit() DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit`

GetPushLimit returns the PushLimit field if non-nil, zero value otherwise.

### GetPushLimitOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPushLimitOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit, bool)`

GetPushLimitOk returns a tuple with the PushLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPushLimit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetPushLimit(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerPushLimit)`

SetPushLimit sets PushLimit field to given value.

### HasPushLimit

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasPushLimit() bool`

HasPushLimit returns a boolean if a field has been set.

### GetOtp

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetOtp() DeviceAuthenticationPolicyCommonMobileApplicationsInnerOtp`

GetOtp returns the Otp field if non-nil, zero value otherwise.

### GetOtpOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetOtpOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerOtp, bool)`

GetOtpOk returns a tuple with the Otp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtp

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetOtp(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerOtp)`

SetOtp sets Otp field to given value.

### HasOtp

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasOtp() bool`

HasOtp returns a boolean if a field has been set.

### GetDeviceAuthorization

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetDeviceAuthorization() DeviceAuthenticationPolicyCommonMobileApplicationsInnerDeviceAuthorization`

GetDeviceAuthorization returns the DeviceAuthorization field if non-nil, zero value otherwise.

### GetDeviceAuthorizationOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetDeviceAuthorizationOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerDeviceAuthorization, bool)`

GetDeviceAuthorizationOk returns a tuple with the DeviceAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthorization

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetDeviceAuthorization(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerDeviceAuthorization)`

SetDeviceAuthorization sets DeviceAuthorization field to given value.

### HasDeviceAuthorization

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasDeviceAuthorization() bool`

HasDeviceAuthorization returns a boolean if a field has been set.

### GetAutoEnrollment

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetAutoEnrollment() DeviceAuthenticationPolicyCommonMobileApplicationsInnerAutoEnrollment`

GetAutoEnrollment returns the AutoEnrollment field if non-nil, zero value otherwise.

### GetAutoEnrollmentOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetAutoEnrollmentOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerAutoEnrollment, bool)`

GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoEnrollment

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetAutoEnrollment(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerAutoEnrollment)`

SetAutoEnrollment sets AutoEnrollment field to given value.

### HasAutoEnrollment

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasAutoEnrollment() bool`

HasAutoEnrollment returns a boolean if a field has been set.

### GetBiometricsEnabled

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetBiometricsEnabled() bool`

GetBiometricsEnabled returns the BiometricsEnabled field if non-nil, zero value otherwise.

### GetBiometricsEnabledOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetBiometricsEnabledOk() (*bool, bool)`

GetBiometricsEnabledOk returns a tuple with the BiometricsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiometricsEnabled

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetBiometricsEnabled(v bool)`

SetBiometricsEnabled sets BiometricsEnabled field to given value.

### HasBiometricsEnabled

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasBiometricsEnabled() bool`

HasBiometricsEnabled returns a boolean if a field has been set.

### GetPairingDisabled

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPairingDisabled() bool`

GetPairingDisabled returns the PairingDisabled field if non-nil, zero value otherwise.

### GetPairingDisabledOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetPairingDisabledOk() (*bool, bool)`

GetPairingDisabledOk returns a tuple with the PairingDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingDisabled

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetPairingDisabled(v bool)`

SetPairingDisabled sets PairingDisabled field to given value.

### HasPairingDisabled

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasPairingDisabled() bool`

HasPairingDisabled returns a boolean if a field has been set.

### GetIntegrityDetection

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetIntegrityDetection() EnumMFADevicePolicyMobileIntegrityDetection`

GetIntegrityDetection returns the IntegrityDetection field if non-nil, zero value otherwise.

### GetIntegrityDetectionOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetIntegrityDetectionOk() (*EnumMFADevicePolicyMobileIntegrityDetection, bool)`

GetIntegrityDetectionOk returns a tuple with the IntegrityDetection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrityDetection

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetIntegrityDetection(v EnumMFADevicePolicyMobileIntegrityDetection)`

SetIntegrityDetection sets IntegrityDetection field to given value.

### HasIntegrityDetection

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasIntegrityDetection() bool`

HasIntegrityDetection returns a boolean if a field has been set.

### GetIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetIpPairingConfiguration() DeviceAuthenticationPolicyCommonMobileApplicationsInnerIpPairingConfiguration`

GetIpPairingConfiguration returns the IpPairingConfiguration field if non-nil, zero value otherwise.

### GetIpPairingConfigurationOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetIpPairingConfigurationOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerIpPairingConfiguration, bool)`

GetIpPairingConfigurationOk returns a tuple with the IpPairingConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetIpPairingConfiguration(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerIpPairingConfiguration)`

SetIpPairingConfiguration sets IpPairingConfiguration field to given value.

### HasIpPairingConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasIpPairingConfiguration() bool`

HasIpPairingConfiguration returns a boolean if a field has been set.

### GetNewRequestDurationConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetNewRequestDurationConfiguration() DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfiguration`

GetNewRequestDurationConfiguration returns the NewRequestDurationConfiguration field if non-nil, zero value otherwise.

### GetNewRequestDurationConfigurationOk

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) GetNewRequestDurationConfigurationOk() (*DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfiguration, bool)`

GetNewRequestDurationConfigurationOk returns a tuple with the NewRequestDurationConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewRequestDurationConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) SetNewRequestDurationConfiguration(v DeviceAuthenticationPolicyCommonMobileApplicationsInnerNewRequestDurationConfiguration)`

SetNewRequestDurationConfiguration sets NewRequestDurationConfiguration field to given value.

### HasNewRequestDurationConfiguration

`func (o *DeviceAuthenticationPolicyCommonMobileApplicationsInner) HasNewRequestDurationConfiguration() bool`

HasNewRequestDurationConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


