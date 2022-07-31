# SignOnPolicyActionMFA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authenticator** | Pointer to [**SignOnPolicyActionMFAAuthenticator**](SignOnPolicyActionMFAAuthenticator.md) |  | [optional] 
**BoundBiometrics** | Pointer to [**SignOnPolicyActionMFABoundBiometrics**](SignOnPolicyActionMFABoundBiometrics.md) |  | [optional] 
**Email** | Pointer to [**SignOnPolicyActionMFAEmail**](SignOnPolicyActionMFAEmail.md) |  | [optional] 
**SecurityKey** | Pointer to [**SignOnPolicyActionMFASecurityKey**](SignOnPolicyActionMFASecurityKey.md) |  | [optional] 
**Sms** | Pointer to [**SignOnPolicyActionMFASms**](SignOnPolicyActionMFASms.md) |  | [optional] 
**Voice** | Pointer to [**SignOnPolicyActionMFAVoice**](SignOnPolicyActionMFAVoice.md) |  | [optional] 
**Applications** | Pointer to [**[]SignOnPolicyActionMFAApplicationsInner**](SignOnPolicyActionMFAApplicationsInner.md) | The applications collection specifies all the native native applications that are allowed in the sign-on policy action.  If the applications collection is empty, a push notification is not allowed for the action. | [optional] 
**DeviceAuthenticationPolicy** | Pointer to [**SignOnPolicyActionMFADeviceAuthenticationPolicy**](SignOnPolicyActionMFADeviceAuthenticationPolicy.md) |  | [optional] 
**NoDeviceMode** | Pointer to [**EnumSignOnPolicyNoDeviceMode**](EnumSignOnPolicyNoDeviceMode.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionMFA

`func NewSignOnPolicyActionMFA() *SignOnPolicyActionMFA`

NewSignOnPolicyActionMFA instantiates a new SignOnPolicyActionMFA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAWithDefaults

`func NewSignOnPolicyActionMFAWithDefaults() *SignOnPolicyActionMFA`

NewSignOnPolicyActionMFAWithDefaults instantiates a new SignOnPolicyActionMFA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthenticator

`func (o *SignOnPolicyActionMFA) GetAuthenticator() SignOnPolicyActionMFAAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *SignOnPolicyActionMFA) GetAuthenticatorOk() (*SignOnPolicyActionMFAAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *SignOnPolicyActionMFA) SetAuthenticator(v SignOnPolicyActionMFAAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *SignOnPolicyActionMFA) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetBoundBiometrics

`func (o *SignOnPolicyActionMFA) GetBoundBiometrics() SignOnPolicyActionMFABoundBiometrics`

GetBoundBiometrics returns the BoundBiometrics field if non-nil, zero value otherwise.

### GetBoundBiometricsOk

`func (o *SignOnPolicyActionMFA) GetBoundBiometricsOk() (*SignOnPolicyActionMFABoundBiometrics, bool)`

GetBoundBiometricsOk returns a tuple with the BoundBiometrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundBiometrics

`func (o *SignOnPolicyActionMFA) SetBoundBiometrics(v SignOnPolicyActionMFABoundBiometrics)`

SetBoundBiometrics sets BoundBiometrics field to given value.

### HasBoundBiometrics

`func (o *SignOnPolicyActionMFA) HasBoundBiometrics() bool`

HasBoundBiometrics returns a boolean if a field has been set.

### GetEmail

`func (o *SignOnPolicyActionMFA) GetEmail() SignOnPolicyActionMFAEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignOnPolicyActionMFA) GetEmailOk() (*SignOnPolicyActionMFAEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignOnPolicyActionMFA) SetEmail(v SignOnPolicyActionMFAEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignOnPolicyActionMFA) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSecurityKey

`func (o *SignOnPolicyActionMFA) GetSecurityKey() SignOnPolicyActionMFASecurityKey`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *SignOnPolicyActionMFA) GetSecurityKeyOk() (*SignOnPolicyActionMFASecurityKey, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *SignOnPolicyActionMFA) SetSecurityKey(v SignOnPolicyActionMFASecurityKey)`

SetSecurityKey sets SecurityKey field to given value.

### HasSecurityKey

`func (o *SignOnPolicyActionMFA) HasSecurityKey() bool`

HasSecurityKey returns a boolean if a field has been set.

### GetSms

`func (o *SignOnPolicyActionMFA) GetSms() SignOnPolicyActionMFASms`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *SignOnPolicyActionMFA) GetSmsOk() (*SignOnPolicyActionMFASms, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *SignOnPolicyActionMFA) SetSms(v SignOnPolicyActionMFASms)`

SetSms sets Sms field to given value.

### HasSms

`func (o *SignOnPolicyActionMFA) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetVoice

`func (o *SignOnPolicyActionMFA) GetVoice() SignOnPolicyActionMFAVoice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *SignOnPolicyActionMFA) GetVoiceOk() (*SignOnPolicyActionMFAVoice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *SignOnPolicyActionMFA) SetVoice(v SignOnPolicyActionMFAVoice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *SignOnPolicyActionMFA) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetApplications

`func (o *SignOnPolicyActionMFA) GetApplications() []SignOnPolicyActionMFAApplicationsInner`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *SignOnPolicyActionMFA) GetApplicationsOk() (*[]SignOnPolicyActionMFAApplicationsInner, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *SignOnPolicyActionMFA) SetApplications(v []SignOnPolicyActionMFAApplicationsInner)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *SignOnPolicyActionMFA) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFA) GetDeviceAuthenticationPolicy() SignOnPolicyActionMFADeviceAuthenticationPolicy`

GetDeviceAuthenticationPolicy returns the DeviceAuthenticationPolicy field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPolicyOk

`func (o *SignOnPolicyActionMFA) GetDeviceAuthenticationPolicyOk() (*SignOnPolicyActionMFADeviceAuthenticationPolicy, bool)`

GetDeviceAuthenticationPolicyOk returns a tuple with the DeviceAuthenticationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFA) SetDeviceAuthenticationPolicy(v SignOnPolicyActionMFADeviceAuthenticationPolicy)`

SetDeviceAuthenticationPolicy sets DeviceAuthenticationPolicy field to given value.

### HasDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFA) HasDeviceAuthenticationPolicy() bool`

HasDeviceAuthenticationPolicy returns a boolean if a field has been set.

### GetNoDeviceMode

`func (o *SignOnPolicyActionMFA) GetNoDeviceMode() EnumSignOnPolicyNoDeviceMode`

GetNoDeviceMode returns the NoDeviceMode field if non-nil, zero value otherwise.

### GetNoDeviceModeOk

`func (o *SignOnPolicyActionMFA) GetNoDeviceModeOk() (*EnumSignOnPolicyNoDeviceMode, bool)`

GetNoDeviceModeOk returns a tuple with the NoDeviceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDeviceMode

`func (o *SignOnPolicyActionMFA) SetNoDeviceMode(v EnumSignOnPolicyNoDeviceMode)`

SetNoDeviceMode sets NoDeviceMode field to given value.

### HasNoDeviceMode

`func (o *SignOnPolicyActionMFA) HasNoDeviceMode() bool`

HasNoDeviceMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


