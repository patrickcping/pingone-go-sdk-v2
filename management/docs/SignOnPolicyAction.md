# SignOnPolicyAction

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
**ConfirmIdentityProviderAttributes** | Pointer to **bool** | A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user&#39;s profile during account creation. This is an optional property. If omitted, the default value is set to false. | [optional] [default to false]
**EnforceLockoutForIdentityProviders** | Pointer to **bool** | A boolean that if set to true and if the user&#39;s account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented. | [optional] 
**Recovery** | Pointer to [**SignOnPolicyActionLoginAllOfRecovery**](SignOnPolicyActionLoginAllOfRecovery.md) |  | [optional] 
**Registration** | Pointer to [**SignOnPolicyActionIDPAllOfRegistration**](SignOnPolicyActionIDPAllOfRegistration.md) |  | [optional] 
**SocialProviders** | Pointer to [**[]SignOnPolicyActionLoginAllOfSocialProviders**](SignOnPolicyActionLoginAllOfSocialProviders.md) | An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow. | [optional] 
**Authenticator** | Pointer to [**SignOnPolicyActionMFAAllOfAuthenticator**](SignOnPolicyActionMFAAllOfAuthenticator.md) |  | [optional] 
**BoundBiometrics** | Pointer to [**SignOnPolicyActionMFAAllOfBoundBiometrics**](SignOnPolicyActionMFAAllOfBoundBiometrics.md) |  | [optional] 
**Email** | Pointer to [**SignOnPolicyActionMFAAllOfEmail**](SignOnPolicyActionMFAAllOfEmail.md) |  | [optional] 
**SecurityKey** | Pointer to [**SignOnPolicyActionMFAAllOfSecurityKey**](SignOnPolicyActionMFAAllOfSecurityKey.md) |  | [optional] 
**Sms** | Pointer to [**SignOnPolicyActionMFAAllOfSms**](SignOnPolicyActionMFAAllOfSms.md) |  | [optional] 
**Voice** | Pointer to [**SignOnPolicyActionMFAAllOfVoice**](SignOnPolicyActionMFAAllOfVoice.md) |  | [optional] 
**Applications** | Pointer to [**[]SignOnPolicyActionMFAAllOfApplications**](SignOnPolicyActionMFAAllOfApplications.md) | The applications collection specifies all the native native applications that are allowed in the sign-on policy action.  If the applications collection is empty, a push notification is not allowed for the action. | [optional] 
**DeviceAuthenticationPolicy** | Pointer to [**SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy**](SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy.md) |  | [optional] 
**NoDevicesMode** | Pointer to [**EnumSignOnPolicyNoDeviceMode**](EnumSignOnPolicyNoDeviceMode.md) |  | [optional] 
**DiscoveryRules** | Pointer to [**[]SignOnPolicyActionIDFirstAllOfDiscoveryRules**](SignOnPolicyActionIDFirstAllOfDiscoveryRules.md) | The list of IDP discovery rules that are evaluated in order when no user is associated with the user identifier. The maximum number of rules is 100. The condition on which this identity provider is used to authenticate the user is expressed using the PingOne policy condition language | [optional] 
**AcrValues** | Pointer to **string** | A string that designates the sign-on policies included in the authorization flow request. Options can include the PingOne predefined sign-on policies, Single_Factor and Multi_Factor, or any custom defined sign-on policy names. Sign-on policy names should be listed in order of preference, and they must be assigned to the application. This property can be configured on the identity provider action and is passed to the identity provider if the identity provider is of type &#x60;SAML&#x60; or &#x60;OPENID_CONNECT&#x60;. | [optional] 
**IdentityProvider** | [**SignOnPolicyActionIDPAllOfIdentityProvider**](SignOnPolicyActionIDPAllOfIdentityProvider.md) |  | 
**PassUserContext** | Pointer to **bool** | A boolean that specifies whether to pass in a login hint to the identity provider on the authentication request. Based on user context, the login hint is set if (1) the user is set on the flow, and (2) the user already has an account link for the identity provider. If both of these conditions are true, then the user is sent to the identity provider with a login hint equal to their externalId for the identity provider (saved on the account link). If these conditions are not true, then the API checks see if there is an OIDC login hint on the flow. If so, that login hint is used. If none of these conditions are true, the login hint parameter is not included on the authorization request to the identity provider. | [optional] 
**Agreement** | [**SignOnPolicyActionAgreementAllOfAgreement**](SignOnPolicyActionAgreementAllOfAgreement.md) |  | 
**DisableDeclineOption** | Pointer to **bool** | When enabled, the &#x60;Do Not Accept&#x60; button will terminate the Flow and display an error message to the user. | [optional] 
**Attributes** | [**[]SignOnPolicyActionProgressiveProfilingAllOfAttributes**](SignOnPolicyActionProgressiveProfilingAllOfAttributes.md) |  | 
**PreventMultiplePromptsPerFlow** | **bool** | A boolean that specifies whether the progressive profiling action will not be executed if another progressive profiling action has already been executed during the flow. This property is required. | 
**PromptIntervalSeconds** | **int32** | An integer that specifies how often to prompt the user to provide profile data for the configured attributes for which they do not have values. This property is required. | 
**PromptText** | **string** | A string that specifies text to display to the user when prompting for attribute values. This property is required. | 

## Methods

### NewSignOnPolicyAction

`func NewSignOnPolicyAction(priority int32, type_ EnumSignOnPolicyType, identityProvider SignOnPolicyActionIDPAllOfIdentityProvider, agreement SignOnPolicyActionAgreementAllOfAgreement, attributes []SignOnPolicyActionProgressiveProfilingAllOfAttributes, preventMultiplePromptsPerFlow bool, promptIntervalSeconds int32, promptText string, ) *SignOnPolicyAction`

NewSignOnPolicyAction instantiates a new SignOnPolicyAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionWithDefaults

`func NewSignOnPolicyActionWithDefaults() *SignOnPolicyAction`

NewSignOnPolicyActionWithDefaults instantiates a new SignOnPolicyAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SignOnPolicyAction) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SignOnPolicyAction) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SignOnPolicyAction) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SignOnPolicyAction) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCondition

`func (o *SignOnPolicyAction) GetCondition() SignOnPolicyActionCommonConditionOrOrInner`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SignOnPolicyAction) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SignOnPolicyAction) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SignOnPolicyAction) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEnvironment

`func (o *SignOnPolicyAction) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SignOnPolicyAction) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SignOnPolicyAction) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SignOnPolicyAction) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SignOnPolicyAction) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyAction) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyAction) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignOnPolicyAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *SignOnPolicyAction) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SignOnPolicyAction) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SignOnPolicyAction) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSignOnPolicy

`func (o *SignOnPolicyAction) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy`

GetSignOnPolicy returns the SignOnPolicy field if non-nil, zero value otherwise.

### GetSignOnPolicyOk

`func (o *SignOnPolicyAction) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool)`

GetSignOnPolicyOk returns a tuple with the SignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicy

`func (o *SignOnPolicyAction) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy)`

SetSignOnPolicy sets SignOnPolicy field to given value.

### HasSignOnPolicy

`func (o *SignOnPolicyAction) HasSignOnPolicy() bool`

HasSignOnPolicy returns a boolean if a field has been set.

### GetType

`func (o *SignOnPolicyAction) GetType() EnumSignOnPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignOnPolicyAction) GetTypeOk() (*EnumSignOnPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignOnPolicyAction) SetType(v EnumSignOnPolicyType)`

SetType sets Type field to given value.


### GetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyAction) GetConfirmIdentityProviderAttributes() bool`

GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field if non-nil, zero value otherwise.

### GetConfirmIdentityProviderAttributesOk

`func (o *SignOnPolicyAction) GetConfirmIdentityProviderAttributesOk() (*bool, bool)`

GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyAction) SetConfirmIdentityProviderAttributes(v bool)`

SetConfirmIdentityProviderAttributes sets ConfirmIdentityProviderAttributes field to given value.

### HasConfirmIdentityProviderAttributes

`func (o *SignOnPolicyAction) HasConfirmIdentityProviderAttributes() bool`

HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.

### GetEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyAction) GetEnforceLockoutForIdentityProviders() bool`

GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field if non-nil, zero value otherwise.

### GetEnforceLockoutForIdentityProvidersOk

`func (o *SignOnPolicyAction) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool)`

GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyAction) SetEnforceLockoutForIdentityProviders(v bool)`

SetEnforceLockoutForIdentityProviders sets EnforceLockoutForIdentityProviders field to given value.

### HasEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyAction) HasEnforceLockoutForIdentityProviders() bool`

HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.

### GetRecovery

`func (o *SignOnPolicyAction) GetRecovery() SignOnPolicyActionLoginAllOfRecovery`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *SignOnPolicyAction) GetRecoveryOk() (*SignOnPolicyActionLoginAllOfRecovery, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *SignOnPolicyAction) SetRecovery(v SignOnPolicyActionLoginAllOfRecovery)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *SignOnPolicyAction) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.

### GetRegistration

`func (o *SignOnPolicyAction) GetRegistration() SignOnPolicyActionIDPAllOfRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SignOnPolicyAction) GetRegistrationOk() (*SignOnPolicyActionIDPAllOfRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SignOnPolicyAction) SetRegistration(v SignOnPolicyActionIDPAllOfRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SignOnPolicyAction) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetSocialProviders

`func (o *SignOnPolicyAction) GetSocialProviders() []SignOnPolicyActionLoginAllOfSocialProviders`

GetSocialProviders returns the SocialProviders field if non-nil, zero value otherwise.

### GetSocialProvidersOk

`func (o *SignOnPolicyAction) GetSocialProvidersOk() (*[]SignOnPolicyActionLoginAllOfSocialProviders, bool)`

GetSocialProvidersOk returns a tuple with the SocialProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProviders

`func (o *SignOnPolicyAction) SetSocialProviders(v []SignOnPolicyActionLoginAllOfSocialProviders)`

SetSocialProviders sets SocialProviders field to given value.

### HasSocialProviders

`func (o *SignOnPolicyAction) HasSocialProviders() bool`

HasSocialProviders returns a boolean if a field has been set.

### GetAuthenticator

`func (o *SignOnPolicyAction) GetAuthenticator() SignOnPolicyActionMFAAllOfAuthenticator`

GetAuthenticator returns the Authenticator field if non-nil, zero value otherwise.

### GetAuthenticatorOk

`func (o *SignOnPolicyAction) GetAuthenticatorOk() (*SignOnPolicyActionMFAAllOfAuthenticator, bool)`

GetAuthenticatorOk returns a tuple with the Authenticator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticator

`func (o *SignOnPolicyAction) SetAuthenticator(v SignOnPolicyActionMFAAllOfAuthenticator)`

SetAuthenticator sets Authenticator field to given value.

### HasAuthenticator

`func (o *SignOnPolicyAction) HasAuthenticator() bool`

HasAuthenticator returns a boolean if a field has been set.

### GetBoundBiometrics

`func (o *SignOnPolicyAction) GetBoundBiometrics() SignOnPolicyActionMFAAllOfBoundBiometrics`

GetBoundBiometrics returns the BoundBiometrics field if non-nil, zero value otherwise.

### GetBoundBiometricsOk

`func (o *SignOnPolicyAction) GetBoundBiometricsOk() (*SignOnPolicyActionMFAAllOfBoundBiometrics, bool)`

GetBoundBiometricsOk returns a tuple with the BoundBiometrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundBiometrics

`func (o *SignOnPolicyAction) SetBoundBiometrics(v SignOnPolicyActionMFAAllOfBoundBiometrics)`

SetBoundBiometrics sets BoundBiometrics field to given value.

### HasBoundBiometrics

`func (o *SignOnPolicyAction) HasBoundBiometrics() bool`

HasBoundBiometrics returns a boolean if a field has been set.

### GetEmail

`func (o *SignOnPolicyAction) GetEmail() SignOnPolicyActionMFAAllOfEmail`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SignOnPolicyAction) GetEmailOk() (*SignOnPolicyActionMFAAllOfEmail, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SignOnPolicyAction) SetEmail(v SignOnPolicyActionMFAAllOfEmail)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SignOnPolicyAction) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetSecurityKey

`func (o *SignOnPolicyAction) GetSecurityKey() SignOnPolicyActionMFAAllOfSecurityKey`

GetSecurityKey returns the SecurityKey field if non-nil, zero value otherwise.

### GetSecurityKeyOk

`func (o *SignOnPolicyAction) GetSecurityKeyOk() (*SignOnPolicyActionMFAAllOfSecurityKey, bool)`

GetSecurityKeyOk returns a tuple with the SecurityKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityKey

`func (o *SignOnPolicyAction) SetSecurityKey(v SignOnPolicyActionMFAAllOfSecurityKey)`

SetSecurityKey sets SecurityKey field to given value.

### HasSecurityKey

`func (o *SignOnPolicyAction) HasSecurityKey() bool`

HasSecurityKey returns a boolean if a field has been set.

### GetSms

`func (o *SignOnPolicyAction) GetSms() SignOnPolicyActionMFAAllOfSms`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *SignOnPolicyAction) GetSmsOk() (*SignOnPolicyActionMFAAllOfSms, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *SignOnPolicyAction) SetSms(v SignOnPolicyActionMFAAllOfSms)`

SetSms sets Sms field to given value.

### HasSms

`func (o *SignOnPolicyAction) HasSms() bool`

HasSms returns a boolean if a field has been set.

### GetVoice

`func (o *SignOnPolicyAction) GetVoice() SignOnPolicyActionMFAAllOfVoice`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *SignOnPolicyAction) GetVoiceOk() (*SignOnPolicyActionMFAAllOfVoice, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *SignOnPolicyAction) SetVoice(v SignOnPolicyActionMFAAllOfVoice)`

SetVoice sets Voice field to given value.

### HasVoice

`func (o *SignOnPolicyAction) HasVoice() bool`

HasVoice returns a boolean if a field has been set.

### GetApplications

`func (o *SignOnPolicyAction) GetApplications() []SignOnPolicyActionMFAAllOfApplications`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *SignOnPolicyAction) GetApplicationsOk() (*[]SignOnPolicyActionMFAAllOfApplications, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *SignOnPolicyAction) SetApplications(v []SignOnPolicyActionMFAAllOfApplications)`

SetApplications sets Applications field to given value.

### HasApplications

`func (o *SignOnPolicyAction) HasApplications() bool`

HasApplications returns a boolean if a field has been set.

### GetDeviceAuthenticationPolicy

`func (o *SignOnPolicyAction) GetDeviceAuthenticationPolicy() SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy`

GetDeviceAuthenticationPolicy returns the DeviceAuthenticationPolicy field if non-nil, zero value otherwise.

### GetDeviceAuthenticationPolicyOk

`func (o *SignOnPolicyAction) GetDeviceAuthenticationPolicyOk() (*SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy, bool)`

GetDeviceAuthenticationPolicyOk returns a tuple with the DeviceAuthenticationPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAuthenticationPolicy

`func (o *SignOnPolicyAction) SetDeviceAuthenticationPolicy(v SignOnPolicyActionMFAAllOfDeviceAuthenticationPolicy)`

SetDeviceAuthenticationPolicy sets DeviceAuthenticationPolicy field to given value.

### HasDeviceAuthenticationPolicy

`func (o *SignOnPolicyAction) HasDeviceAuthenticationPolicy() bool`

HasDeviceAuthenticationPolicy returns a boolean if a field has been set.

### GetNoDevicesMode

`func (o *SignOnPolicyAction) GetNoDevicesMode() EnumSignOnPolicyNoDeviceMode`

GetNoDevicesMode returns the NoDevicesMode field if non-nil, zero value otherwise.

### GetNoDevicesModeOk

`func (o *SignOnPolicyAction) GetNoDevicesModeOk() (*EnumSignOnPolicyNoDeviceMode, bool)`

GetNoDevicesModeOk returns a tuple with the NoDevicesMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoDevicesMode

`func (o *SignOnPolicyAction) SetNoDevicesMode(v EnumSignOnPolicyNoDeviceMode)`

SetNoDevicesMode sets NoDevicesMode field to given value.

### HasNoDevicesMode

`func (o *SignOnPolicyAction) HasNoDevicesMode() bool`

HasNoDevicesMode returns a boolean if a field has been set.

### GetDiscoveryRules

`func (o *SignOnPolicyAction) GetDiscoveryRules() []SignOnPolicyActionIDFirstAllOfDiscoveryRules`

GetDiscoveryRules returns the DiscoveryRules field if non-nil, zero value otherwise.

### GetDiscoveryRulesOk

`func (o *SignOnPolicyAction) GetDiscoveryRulesOk() (*[]SignOnPolicyActionIDFirstAllOfDiscoveryRules, bool)`

GetDiscoveryRulesOk returns a tuple with the DiscoveryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryRules

`func (o *SignOnPolicyAction) SetDiscoveryRules(v []SignOnPolicyActionIDFirstAllOfDiscoveryRules)`

SetDiscoveryRules sets DiscoveryRules field to given value.

### HasDiscoveryRules

`func (o *SignOnPolicyAction) HasDiscoveryRules() bool`

HasDiscoveryRules returns a boolean if a field has been set.

### GetAcrValues

`func (o *SignOnPolicyAction) GetAcrValues() string`

GetAcrValues returns the AcrValues field if non-nil, zero value otherwise.

### GetAcrValuesOk

`func (o *SignOnPolicyAction) GetAcrValuesOk() (*string, bool)`

GetAcrValuesOk returns a tuple with the AcrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrValues

`func (o *SignOnPolicyAction) SetAcrValues(v string)`

SetAcrValues sets AcrValues field to given value.

### HasAcrValues

`func (o *SignOnPolicyAction) HasAcrValues() bool`

HasAcrValues returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *SignOnPolicyAction) GetIdentityProvider() SignOnPolicyActionIDPAllOfIdentityProvider`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *SignOnPolicyAction) GetIdentityProviderOk() (*SignOnPolicyActionIDPAllOfIdentityProvider, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *SignOnPolicyAction) SetIdentityProvider(v SignOnPolicyActionIDPAllOfIdentityProvider)`

SetIdentityProvider sets IdentityProvider field to given value.


### GetPassUserContext

`func (o *SignOnPolicyAction) GetPassUserContext() bool`

GetPassUserContext returns the PassUserContext field if non-nil, zero value otherwise.

### GetPassUserContextOk

`func (o *SignOnPolicyAction) GetPassUserContextOk() (*bool, bool)`

GetPassUserContextOk returns a tuple with the PassUserContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassUserContext

`func (o *SignOnPolicyAction) SetPassUserContext(v bool)`

SetPassUserContext sets PassUserContext field to given value.

### HasPassUserContext

`func (o *SignOnPolicyAction) HasPassUserContext() bool`

HasPassUserContext returns a boolean if a field has been set.

### GetAgreement

`func (o *SignOnPolicyAction) GetAgreement() SignOnPolicyActionAgreementAllOfAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *SignOnPolicyAction) GetAgreementOk() (*SignOnPolicyActionAgreementAllOfAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *SignOnPolicyAction) SetAgreement(v SignOnPolicyActionAgreementAllOfAgreement)`

SetAgreement sets Agreement field to given value.


### GetDisableDeclineOption

`func (o *SignOnPolicyAction) GetDisableDeclineOption() bool`

GetDisableDeclineOption returns the DisableDeclineOption field if non-nil, zero value otherwise.

### GetDisableDeclineOptionOk

`func (o *SignOnPolicyAction) GetDisableDeclineOptionOk() (*bool, bool)`

GetDisableDeclineOptionOk returns a tuple with the DisableDeclineOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableDeclineOption

`func (o *SignOnPolicyAction) SetDisableDeclineOption(v bool)`

SetDisableDeclineOption sets DisableDeclineOption field to given value.

### HasDisableDeclineOption

`func (o *SignOnPolicyAction) HasDisableDeclineOption() bool`

HasDisableDeclineOption returns a boolean if a field has been set.

### GetAttributes

`func (o *SignOnPolicyAction) GetAttributes() []SignOnPolicyActionProgressiveProfilingAllOfAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SignOnPolicyAction) GetAttributesOk() (*[]SignOnPolicyActionProgressiveProfilingAllOfAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SignOnPolicyAction) SetAttributes(v []SignOnPolicyActionProgressiveProfilingAllOfAttributes)`

SetAttributes sets Attributes field to given value.


### GetPreventMultiplePromptsPerFlow

`func (o *SignOnPolicyAction) GetPreventMultiplePromptsPerFlow() bool`

GetPreventMultiplePromptsPerFlow returns the PreventMultiplePromptsPerFlow field if non-nil, zero value otherwise.

### GetPreventMultiplePromptsPerFlowOk

`func (o *SignOnPolicyAction) GetPreventMultiplePromptsPerFlowOk() (*bool, bool)`

GetPreventMultiplePromptsPerFlowOk returns a tuple with the PreventMultiplePromptsPerFlow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreventMultiplePromptsPerFlow

`func (o *SignOnPolicyAction) SetPreventMultiplePromptsPerFlow(v bool)`

SetPreventMultiplePromptsPerFlow sets PreventMultiplePromptsPerFlow field to given value.


### GetPromptIntervalSeconds

`func (o *SignOnPolicyAction) GetPromptIntervalSeconds() int32`

GetPromptIntervalSeconds returns the PromptIntervalSeconds field if non-nil, zero value otherwise.

### GetPromptIntervalSecondsOk

`func (o *SignOnPolicyAction) GetPromptIntervalSecondsOk() (*int32, bool)`

GetPromptIntervalSecondsOk returns a tuple with the PromptIntervalSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptIntervalSeconds

`func (o *SignOnPolicyAction) SetPromptIntervalSeconds(v int32)`

SetPromptIntervalSeconds sets PromptIntervalSeconds field to given value.


### GetPromptText

`func (o *SignOnPolicyAction) GetPromptText() string`

GetPromptText returns the PromptText field if non-nil, zero value otherwise.

### GetPromptTextOk

`func (o *SignOnPolicyAction) GetPromptTextOk() (*string, bool)`

GetPromptTextOk returns a tuple with the PromptText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptText

`func (o *SignOnPolicyAction) SetPromptText(v string)`

SetPromptText sets PromptText field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


