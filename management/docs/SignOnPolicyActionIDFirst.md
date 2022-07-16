# SignOnPolicyActionIDFirst

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmIdentityProviderAttributes** | Pointer to **bool** | A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user&#39;s profile during account creation. This is an optional property. If omitted, the default value is set to false. | [optional] [default to false]
**DiscoveryRules** | Pointer to [**[]SignOnPolicyActionIDFirstDiscoveryRulesInner**](SignOnPolicyActionIDFirstDiscoveryRulesInner.md) | The list of IDP discovery rules that are evaluated in order when no user is associated with the user identifier. The maximum number of rules is 100. The condition on which this identity provider is used to authenticate the user is expressed using the PingOne policy condition language | [optional] 
**EnforceLockoutForIdentityProviders** | Pointer to **bool** | A boolean that if set to true and if the user&#39;s account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented. | [optional] 
**Recovery** | Pointer to [**SignOnPolicyActionLoginRecovery**](SignOnPolicyActionLoginRecovery.md) |  | [optional] 
**Registration** | Pointer to [**SignOnPolicyActionLoginRegistration**](SignOnPolicyActionLoginRegistration.md) |  | [optional] 
**SocialProviders** | Pointer to [**[]SignOnPolicyActionLoginSocialProvidersInner**](SignOnPolicyActionLoginSocialProvidersInner.md) | An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow. | [optional] 

## Methods

### NewSignOnPolicyActionIDFirst

`func NewSignOnPolicyActionIDFirst() *SignOnPolicyActionIDFirst`

NewSignOnPolicyActionIDFirst instantiates a new SignOnPolicyActionIDFirst object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionIDFirstWithDefaults

`func NewSignOnPolicyActionIDFirstWithDefaults() *SignOnPolicyActionIDFirst`

NewSignOnPolicyActionIDFirstWithDefaults instantiates a new SignOnPolicyActionIDFirst object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirst) GetConfirmIdentityProviderAttributes() bool`

GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field if non-nil, zero value otherwise.

### GetConfirmIdentityProviderAttributesOk

`func (o *SignOnPolicyActionIDFirst) GetConfirmIdentityProviderAttributesOk() (*bool, bool)`

GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirst) SetConfirmIdentityProviderAttributes(v bool)`

SetConfirmIdentityProviderAttributes sets ConfirmIdentityProviderAttributes field to given value.

### HasConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirst) HasConfirmIdentityProviderAttributes() bool`

HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.

### GetDiscoveryRules

`func (o *SignOnPolicyActionIDFirst) GetDiscoveryRules() []SignOnPolicyActionIDFirstDiscoveryRulesInner`

GetDiscoveryRules returns the DiscoveryRules field if non-nil, zero value otherwise.

### GetDiscoveryRulesOk

`func (o *SignOnPolicyActionIDFirst) GetDiscoveryRulesOk() (*[]SignOnPolicyActionIDFirstDiscoveryRulesInner, bool)`

GetDiscoveryRulesOk returns a tuple with the DiscoveryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryRules

`func (o *SignOnPolicyActionIDFirst) SetDiscoveryRules(v []SignOnPolicyActionIDFirstDiscoveryRulesInner)`

SetDiscoveryRules sets DiscoveryRules field to given value.

### HasDiscoveryRules

`func (o *SignOnPolicyActionIDFirst) HasDiscoveryRules() bool`

HasDiscoveryRules returns a boolean if a field has been set.

### GetEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionIDFirst) GetEnforceLockoutForIdentityProviders() bool`

GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field if non-nil, zero value otherwise.

### GetEnforceLockoutForIdentityProvidersOk

`func (o *SignOnPolicyActionIDFirst) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool)`

GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionIDFirst) SetEnforceLockoutForIdentityProviders(v bool)`

SetEnforceLockoutForIdentityProviders sets EnforceLockoutForIdentityProviders field to given value.

### HasEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionIDFirst) HasEnforceLockoutForIdentityProviders() bool`

HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.

### GetRecovery

`func (o *SignOnPolicyActionIDFirst) GetRecovery() SignOnPolicyActionLoginRecovery`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *SignOnPolicyActionIDFirst) GetRecoveryOk() (*SignOnPolicyActionLoginRecovery, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *SignOnPolicyActionIDFirst) SetRecovery(v SignOnPolicyActionLoginRecovery)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *SignOnPolicyActionIDFirst) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.

### GetRegistration

`func (o *SignOnPolicyActionIDFirst) GetRegistration() SignOnPolicyActionLoginRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SignOnPolicyActionIDFirst) GetRegistrationOk() (*SignOnPolicyActionLoginRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SignOnPolicyActionIDFirst) SetRegistration(v SignOnPolicyActionLoginRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SignOnPolicyActionIDFirst) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetSocialProviders

`func (o *SignOnPolicyActionIDFirst) GetSocialProviders() []SignOnPolicyActionLoginSocialProvidersInner`

GetSocialProviders returns the SocialProviders field if non-nil, zero value otherwise.

### GetSocialProvidersOk

`func (o *SignOnPolicyActionIDFirst) GetSocialProvidersOk() (*[]SignOnPolicyActionLoginSocialProvidersInner, bool)`

GetSocialProvidersOk returns a tuple with the SocialProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProviders

`func (o *SignOnPolicyActionIDFirst) SetSocialProviders(v []SignOnPolicyActionLoginSocialProvidersInner)`

SetSocialProviders sets SocialProviders field to given value.

### HasSocialProviders

`func (o *SignOnPolicyActionIDFirst) HasSocialProviders() bool`

HasSocialProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


