# SignOnPolicyActionMFAAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticator** | Pointer to [**SignOnPolicyActionMFAAllOfAuthenticator**](SignOnPolicyActionMFAAllOfAuthenticator.md) |  | [optional] 
**BoundBiometrics** | Pointer to [**SignOnPolicyActionMFAAllOfBoundBiometrics**](SignOnPolicyActionMFAAllOfBoundBiometrics.md) |  | [optional] 
**Email** | Pointer to [**SignOnPolicyActionMFAAllOfEmail**](SignOnPolicyActionMFAAllOfEmail.md) |  | [optional] 
**SecurityKey** | Pointer to [**SignOnPolicyActionMFAAllOfSecurityKey**](SignOnPolicyActionMFAAllOfSecurityKey.md) |  | [optional] 
**Sms** | Pointer to [**SignOnPolicyActionMFAAllOfSms**](SignOnPolicyActionMFAAllOfSms.md) |  | [optional] 
**Voice** | Pointer to [**SignOnPolicyActionMFAAllOfVoice**](SignOnPolicyActionMFAAllOfVoice.md) |  | [optional] 
**Applications** | Pointer to [**[]SignOnPolicyActionMFAAllOfApplications**](SignOnPolicyActionMFAAllOfApplications.md) | The applications collection specifies all the native native applications that are allowed in the sign-on policy action.  If the applications collection is empty, a push notification is not allowed for the action. | [optional] 
**DeviceAuthenticationPolicy** | Pointer to [**SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy**](SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy.md) |  | [optional] 
**NoDeviceMode** | Pointer to [**EnumSignOnPolicyNoDeviceMode**](EnumSignOnPolicyNoDeviceMode.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionMFAAllOf

`func NewSignOnPolicyActionMFAAllOf() *SignOnPolicyActionMFAAllOf`

NewSignOnPolicyActionMFAAllOf instantiates a new SignOnPolicyActionMFAAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAAllOfWithDefaults

`func NewSignOnPolicyActionMFAAllOfWithDefaults() *SignOnPolicyActionMFAAllOf`

NewSignOnPolicyActionMFAAllOfWithDefaults instantiates a new SignOnPolicyActionMFAAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticator

`func (o *SignOnPolicyActionMFAAllOf) GetAuthenticator() SignOnPolicyActionMFAAllOfAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *SignOnPolicyActionMFAAllOf) GetAuthenticatorOk() (*SignOnPolicyActionMFAAllOfAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *SignOnPolicyActionMFAAllOf) SetAuthenticator(v SignOnPolicyActionMFAAllOfAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *SignOnPolicyActionMFAAllOf) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetBoundBiometrics

`func (o *SignOnPolicyActionMFAAllOf) GetBoundBiometrics() SignOnPolicyActionMFAAllOfBoundBiometrics`

GetBoundBiometrics returns the BoundBiometrics field if non-nil, zero value otherwise.

### GetBoundBiometricsOk

`func (o *SignOnPolicyActionMFAAllOf) GetBoundBiometricsOk() (*SignOnPolicyActionMFAAllOfBoundBiometrics, bool)`

GetBoundBiometricsOk returns a tuple with the BoundBiometrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundBiometrics

`func (o *SignOnPolicyActionMFAAllOf) SetBoundBiometrics(v SignOnPolicyActionMFAAllOfBoundBiometrics)`

SetBoundBiometrics sets BoundBiometrics field to given value.

### HasBoundBiometrics

`func (o *SignOnPolicyActionMFAAllOf) HasBoundBiometrics() bool`

HasBoundBiometrics returns a boolean if a field has been set.

### GetEmail

`func (o *SignOnPolicyActionMFAAllOf) GetEmail() SignOnPolicyActionMFAAllOfEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignOnPolicyActionMFAAllOf) GetEmailOk() (*SignOnPolicyActionMFAAllOfEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignOnPolicyActionMFAAllOf) SetEmail(v SignOnPolicyActionMFAAllOfEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignOnPolicyActionMFAAllOf) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSecurityKey

`func (o *SignOnPolicyActionMFAAllOf) GetSecurityKey() SignOnPolicyActionMFAAllOfSecurityKey`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *SignOnPolicyActionMFAAllOf) GetSecurityKeyOk() (*SignOnPolicyActionMFAAllOfSecurityKey, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *SignOnPolicyActionMFAAllOf) SetSecurityKey(v SignOnPolicyActionMFAAllOfSecurityKey)`

SetSecurityKey sets SecurityKey field to given value.

### HasSecurityKey

`func (o *SignOnPolicyActionMFAAllOf) HasSecurityKey() bool`

HasSecurityKey returns a boolean if a field has been set.

### GetSms

`func (o *SignOnPolicyActionMFAAllOf) GetSms() SignOnPolicyActionMFAAllOfSms`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *SignOnPolicyActionMFAAllOf) GetSmsOk() (*SignOnPolicyActionMFAAllOfSms, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *SignOnPolicyActionMFAAllOf) SetSms(v SignOnPolicyActionMFAAllOfSms)`

SetSms sets Sms field to given value.

### HasSms

`func (o *SignOnPolicyActionMFAAllOf) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetVoice

`func (o *SignOnPolicyActionMFAAllOf) GetVoice() SignOnPolicyActionMFAAllOfVoice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *SignOnPolicyActionMFAAllOf) GetVoiceOk() (*SignOnPolicyActionMFAAllOfVoice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *SignOnPolicyActionMFAAllOf) SetVoice(v SignOnPolicyActionMFAAllOfVoice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *SignOnPolicyActionMFAAllOf) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetApplications

`func (o *SignOnPolicyActionMFAAllOf) GetApplications() []SignOnPolicyActionMFAAllOfApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *SignOnPolicyActionMFAAllOf) GetApplicationsOk() (*[]SignOnPolicyActionMFAAllOfApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *SignOnPolicyActionMFAAllOf) SetApplications(v []SignOnPolicyActionMFAAllOfApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *SignOnPolicyActionMFAAllOf) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFAAllOf) GetDeviceAuthenticationPolicy() SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy`

GetDeviceAuthenticationPolicy returns the DeviceAuthenticationPolicy field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPolicyOk

`func (o *SignOnPolicyActionMFAAllOf) GetDeviceAuthenticationPolicyOk() (*SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy, bool)`

GetDeviceAuthenticationPolicyOk returns a tuple with the DeviceAuthenticationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFAAllOf) SetDeviceAuthenticationPolicy(v SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy)`

SetDeviceAuthenticationPolicy sets DeviceAuthenticationPolicy field to given value.

### HasDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFAAllOf) HasDeviceAuthenticationPolicy() bool`

HasDeviceAuthenticationPolicy returns a boolean if a field has been set.

### GetNoDeviceMode

`func (o *SignOnPolicyActionMFAAllOf) GetNoDeviceMode() EnumSignOnPolicyNoDeviceMode`

GetNoDeviceMode returns the NoDeviceMode field if non-nil, zero value otherwise.

### GetNoDeviceModeOk

`func (o *SignOnPolicyActionMFAAllOf) GetNoDeviceModeOk() (*EnumSignOnPolicyNoDeviceMode, bool)`

GetNoDeviceModeOk returns a tuple with the NoDeviceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDeviceMode

`func (o *SignOnPolicyActionMFAAllOf) SetNoDeviceMode(v EnumSignOnPolicyNoDeviceMode)`

SetNoDeviceMode sets NoDeviceMode field to given value.

### HasNoDeviceMode

`func (o *SignOnPolicyActionMFAAllOf) HasNoDeviceMode() bool`

HasNoDeviceMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


