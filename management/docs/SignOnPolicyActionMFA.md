# SignOnPolicyActionMFA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Condition** | Pointer to [**SignOnPolicyActionCommonConditionOrOrInner**](SignOnPolicyActionCommonConditionOrOrInner.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the sign-on policy assignment resource’s unique identifier. | [optional] [readonly] 
**Priority** | **int32** | An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property. | 
**SignOnPolicy** | Pointer to [**SignOnPolicyActionCommonSignOnPolicy**](SignOnPolicyActionCommonSignOnPolicy.md) |  | [optional] 
**Type** | [**EnumSignOnPolicyType**](EnumSignOnPolicyType.md) |  | 
**Authenticator** | Pointer to [**SignOnPolicyActionMFAAllOfAuthenticator**](SignOnPolicyActionMFAAllOfAuthenticator.md) |  | [optional] 
**BoundBiometrics** | Pointer to [**SignOnPolicyActionMFAAllOfBoundBiometrics**](SignOnPolicyActionMFAAllOfBoundBiometrics.md) |  | [optional] 
**Email** | Pointer to [**SignOnPolicyActionMFAAllOfEmail**](SignOnPolicyActionMFAAllOfEmail.md) |  | [optional] 
**SecurityKey** | Pointer to [**SignOnPolicyActionMFAAllOfSecurityKey**](SignOnPolicyActionMFAAllOfSecurityKey.md) |  | [optional] 
**Sms** | Pointer to [**SignOnPolicyActionMFAAllOfSms**](SignOnPolicyActionMFAAllOfSms.md) |  | [optional] 
**Voice** | Pointer to [**SignOnPolicyActionMFAAllOfVoice**](SignOnPolicyActionMFAAllOfVoice.md) |  | [optional] 
**Applications** | Pointer to [**[]SignOnPolicyActionMFAAllOfApplications**](SignOnPolicyActionMFAAllOfApplications.md) | The applications collection specifies all the native native applications that are allowed in the sign-on policy action.  If the applications collection is empty, a push notification is not allowed for the action. | [optional] 
**DeviceAuthenticationPolicy** | Pointer to [**SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy**](SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy.md) |  | [optional] 
**NoDevicesMode** | Pointer to [**EnumSignOnPolicyNoDeviceMode**](EnumSignOnPolicyNoDeviceMode.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionMFA

`func NewSignOnPolicyActionMFA(priority int32, type_ EnumSignOnPolicyType, ) *SignOnPolicyActionMFA`

NewSignOnPolicyActionMFA instantiates a new SignOnPolicyActionMFA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionMFAWithDefaults

`func NewSignOnPolicyActionMFAWithDefaults() *SignOnPolicyActionMFA`

NewSignOnPolicyActionMFAWithDefaults instantiates a new SignOnPolicyActionMFA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SignOnPolicyActionMFA) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SignOnPolicyActionMFA) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SignOnPolicyActionMFA) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SignOnPolicyActionMFA) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCondition

`func (o *SignOnPolicyActionMFA) GetCondition() SignOnPolicyActionCommonConditionOrOrInner`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SignOnPolicyActionMFA) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SignOnPolicyActionMFA) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SignOnPolicyActionMFA) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEnvironment

`func (o *SignOnPolicyActionMFA) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SignOnPolicyActionMFA) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SignOnPolicyActionMFA) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SignOnPolicyActionMFA) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SignOnPolicyActionMFA) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyActionMFA) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyActionMFA) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignOnPolicyActionMFA) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *SignOnPolicyActionMFA) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SignOnPolicyActionMFA) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SignOnPolicyActionMFA) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSignOnPolicy

`func (o *SignOnPolicyActionMFA) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy`

GetSignOnPolicy returns the SignOnPolicy field if non-nil, zero value otherwise.

### GetSignOnPolicyOk

`func (o *SignOnPolicyActionMFA) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool)`

GetSignOnPolicyOk returns a tuple with the SignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicy

`func (o *SignOnPolicyActionMFA) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy)`

SetSignOnPolicy sets SignOnPolicy field to given value.

### HasSignOnPolicy

`func (o *SignOnPolicyActionMFA) HasSignOnPolicy() bool`

HasSignOnPolicy returns a boolean if a field has been set.

### GetType

`func (o *SignOnPolicyActionMFA) GetType() EnumSignOnPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignOnPolicyActionMFA) GetTypeOk() (*EnumSignOnPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignOnPolicyActionMFA) SetType(v EnumSignOnPolicyType)`

SetType sets Type field to given value.


### GetAuthenticator

`func (o *SignOnPolicyActionMFA) GetAuthenticator() SignOnPolicyActionMFAAllOfAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *SignOnPolicyActionMFA) GetAuthenticatorOk() (*SignOnPolicyActionMFAAllOfAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *SignOnPolicyActionMFA) SetAuthenticator(v SignOnPolicyActionMFAAllOfAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *SignOnPolicyActionMFA) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetBoundBiometrics

`func (o *SignOnPolicyActionMFA) GetBoundBiometrics() SignOnPolicyActionMFAAllOfBoundBiometrics`

GetBoundBiometrics returns the BoundBiometrics field if non-nil, zero value otherwise.

### GetBoundBiometricsOk

`func (o *SignOnPolicyActionMFA) GetBoundBiometricsOk() (*SignOnPolicyActionMFAAllOfBoundBiometrics, bool)`

GetBoundBiometricsOk returns a tuple with the BoundBiometrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundBiometrics

`func (o *SignOnPolicyActionMFA) SetBoundBiometrics(v SignOnPolicyActionMFAAllOfBoundBiometrics)`

SetBoundBiometrics sets BoundBiometrics field to given value.

### HasBoundBiometrics

`func (o *SignOnPolicyActionMFA) HasBoundBiometrics() bool`

HasBoundBiometrics returns a boolean if a field has been set.

### GetEmail

`func (o *SignOnPolicyActionMFA) GetEmail() SignOnPolicyActionMFAAllOfEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignOnPolicyActionMFA) GetEmailOk() (*SignOnPolicyActionMFAAllOfEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignOnPolicyActionMFA) SetEmail(v SignOnPolicyActionMFAAllOfEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignOnPolicyActionMFA) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSecurityKey

`func (o *SignOnPolicyActionMFA) GetSecurityKey() SignOnPolicyActionMFAAllOfSecurityKey`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *SignOnPolicyActionMFA) GetSecurityKeyOk() (*SignOnPolicyActionMFAAllOfSecurityKey, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *SignOnPolicyActionMFA) SetSecurityKey(v SignOnPolicyActionMFAAllOfSecurityKey)`

SetSecurityKey sets SecurityKey field to given value.

### HasSecurityKey

`func (o *SignOnPolicyActionMFA) HasSecurityKey() bool`

HasSecurityKey returns a boolean if a field has been set.

### GetSms

`func (o *SignOnPolicyActionMFA) GetSms() SignOnPolicyActionMFAAllOfSms`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *SignOnPolicyActionMFA) GetSmsOk() (*SignOnPolicyActionMFAAllOfSms, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *SignOnPolicyActionMFA) SetSms(v SignOnPolicyActionMFAAllOfSms)`

SetSms sets Sms field to given value.

### HasSms

`func (o *SignOnPolicyActionMFA) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetVoice

`func (o *SignOnPolicyActionMFA) GetVoice() SignOnPolicyActionMFAAllOfVoice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *SignOnPolicyActionMFA) GetVoiceOk() (*SignOnPolicyActionMFAAllOfVoice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *SignOnPolicyActionMFA) SetVoice(v SignOnPolicyActionMFAAllOfVoice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *SignOnPolicyActionMFA) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetApplications

`func (o *SignOnPolicyActionMFA) GetApplications() []SignOnPolicyActionMFAAllOfApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *SignOnPolicyActionMFA) GetApplicationsOk() (*[]SignOnPolicyActionMFAAllOfApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *SignOnPolicyActionMFA) SetApplications(v []SignOnPolicyActionMFAAllOfApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *SignOnPolicyActionMFA) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFA) GetDeviceAuthenticationPolicy() SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy`

GetDeviceAuthenticationPolicy returns the DeviceAuthenticationPolicy field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPolicyOk

`func (o *SignOnPolicyActionMFA) GetDeviceAuthenticationPolicyOk() (*SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy, bool)`

GetDeviceAuthenticationPolicyOk returns a tuple with the DeviceAuthenticationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFA) SetDeviceAuthenticationPolicy(v SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy)`

SetDeviceAuthenticationPolicy sets DeviceAuthenticationPolicy field to given value.

### HasDeviceAuthenticationPolicy

`func (o *SignOnPolicyActionMFA) HasDeviceAuthenticationPolicy() bool`

HasDeviceAuthenticationPolicy returns a boolean if a field has been set.

### GetNoDevicesMode

`func (o *SignOnPolicyActionMFA) GetNoDevicesMode() EnumSignOnPolicyNoDeviceMode`

GetNoDevicesMode returns the NoDevicesMode field if non-nil, zero value otherwise.

### GetNoDevicesModeOk

`func (o *SignOnPolicyActionMFA) GetNoDevicesModeOk() (*EnumSignOnPolicyNoDeviceMode, bool)`

GetNoDevicesModeOk returns a tuple with the NoDevicesMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDevicesMode

`func (o *SignOnPolicyActionMFA) SetNoDevicesMode(v EnumSignOnPolicyNoDeviceMode)`

SetNoDevicesMode sets NoDevicesMode field to given value.

### HasNoDevicesMode

`func (o *SignOnPolicyActionMFA) HasNoDevicesMode() bool`

HasNoDevicesMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


