# SignOnPolicyActionIDFirstAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfirmIdentityProviderAttributes** | Pointer to **bool** | A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user&#39;s profile during account creation. This is an optional property. If omitted, the default value is set to false. | [optional] [default to false]
**DiscoveryRules** | Pointer to [**[]SignOnPolicyActionIDFirstAllOfDiscoveryRules**](SignOnPolicyActionIDFirstAllOfDiscoveryRules.md) | The list of IDP discovery rules that are evaluated in order when no user is associated with the user identifier. The maximum number of rules is 100. The condition on which this identity provider is used to authenticate the user is expressed using the PingOne policy condition language | [optional] 
**EnforceLockoutForIdentityProviders** | Pointer to **bool** | A boolean that if set to true and if the user&#39;s account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented. | [optional] 
**Recovery** | Pointer to [**SignOnPolicyActionLoginAllOfRecovery**](SignOnPolicyActionLoginAllOfRecovery.md) |  | [optional] 
**Registration** | Pointer to [**SignOnPolicyActionLoginAllOfRegistration**](SignOnPolicyActionLoginAllOfRegistration.md) |  | [optional] 
**SocialProviders** | Pointer to [**[]SignOnPolicyActionLoginAllOfSocialProviders**](SignOnPolicyActionLoginAllOfSocialProviders.md) | An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow. | [optional] 

## Methods

### NewSignOnPolicyActionIDFirstAllOf

`func NewSignOnPolicyActionIDFirstAllOf() *SignOnPolicyActionIDFirstAllOf`

NewSignOnPolicyActionIDFirstAllOf instantiates a new SignOnPolicyActionIDFirstAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionIDFirstAllOfWithDefaults

`func NewSignOnPolicyActionIDFirstAllOfWithDefaults() *SignOnPolicyActionIDFirstAllOf`

NewSignOnPolicyActionIDFirstAllOfWithDefaults instantiates a new SignOnPolicyActionIDFirstAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirstAllOf) GetConfirmIdentityProviderAttributes() bool`

GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field if non-nil, zero value otherwise.

### GetConfirmIdentityProviderAttributesOk

`func (o *SignOnPolicyActionIDFirstAllOf) GetConfirmIdentityProviderAttributesOk() (*bool, bool)`

GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirstAllOf) SetConfirmIdentityProviderAttributes(v bool)`

SetConfirmIdentityProviderAttributes sets ConfirmIdentityProviderAttributes field to given value.

### HasConfirmIdentityProviderAttributes

`func (o *SignOnPolicyActionIDFirstAllOf) HasConfirmIdentityProviderAttributes() bool`

HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.

### GetDiscoveryRules

`func (o *SignOnPolicyActionIDFirstAllOf) GetDiscoveryRules() []SignOnPolicyActionIDFirstAllOfDiscoveryRules`

GetDiscoveryRules returns the DiscoveryRules field if non-nil, zero value otherwise.

### GetDiscoveryRulesOk

`func (o *SignOnPolicyActionIDFirstAllOf) GetDiscoveryRulesOk() (*[]SignOnPolicyActionIDFirstAllOfDiscoveryRules, bool)`

GetDiscoveryRulesOk returns a tuple with the DiscoveryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryRules

`func (o *SignOnPolicyActionIDFirstAllOf) SetDiscoveryRules(v []SignOnPolicyActionIDFirstAllOfDiscoveryRules)`

SetDiscoveryRules sets DiscoveryRules field to given value.

### HasDiscoveryRules

`func (o *SignOnPolicyActionIDFirstAllOf) HasDiscoveryRules() bool`

HasDiscoveryRules returns a boolean if a field has been set.

### GetEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionIDFirstAllOf) GetEnforceLockoutForIdentityProviders() bool`

GetEnforceLockoutForIdentityProviders returns the EnforceLockoutForIdentityProviders field if non-nil, zero value otherwise.

### GetEnforceLockoutForIdentityProvidersOk

`func (o *SignOnPolicyActionIDFirstAllOf) GetEnforceLockoutForIdentityProvidersOk() (*bool, bool)`

GetEnforceLockoutForIdentityProvidersOk returns a tuple with the EnforceLockoutForIdentityProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionIDFirstAllOf) SetEnforceLockoutForIdentityProviders(v bool)`

SetEnforceLockoutForIdentityProviders sets EnforceLockoutForIdentityProviders field to given value.

### HasEnforceLockoutForIdentityProviders

`func (o *SignOnPolicyActionIDFirstAllOf) HasEnforceLockoutForIdentityProviders() bool`

HasEnforceLockoutForIdentityProviders returns a boolean if a field has been set.

### GetRecovery

`func (o *SignOnPolicyActionIDFirstAllOf) GetRecovery() SignOnPolicyActionLoginAllOfRecovery`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *SignOnPolicyActionIDFirstAllOf) GetRecoveryOk() (*SignOnPolicyActionLoginAllOfRecovery, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *SignOnPolicyActionIDFirstAllOf) SetRecovery(v SignOnPolicyActionLoginAllOfRecovery)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *SignOnPolicyActionIDFirstAllOf) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.

### GetRegistration

`func (o *SignOnPolicyActionIDFirstAllOf) GetRegistration() SignOnPolicyActionLoginAllOfRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SignOnPolicyActionIDFirstAllOf) GetRegistrationOk() (*SignOnPolicyActionLoginAllOfRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SignOnPolicyActionIDFirstAllOf) SetRegistration(v SignOnPolicyActionLoginAllOfRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SignOnPolicyActionIDFirstAllOf) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetSocialProviders

`func (o *SignOnPolicyActionIDFirstAllOf) GetSocialProviders() []SignOnPolicyActionLoginAllOfSocialProviders`

GetSocialProviders returns the SocialProviders field if non-nil, zero value otherwise.

### GetSocialProvidersOk

`func (o *SignOnPolicyActionIDFirstAllOf) GetSocialProvidersOk() (*[]SignOnPolicyActionLoginAllOfSocialProviders, bool)`

GetSocialProvidersOk returns a tuple with the SocialProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProviders

`func (o *SignOnPolicyActionIDFirstAllOf) SetSocialProviders(v []SignOnPolicyActionLoginAllOfSocialProviders)`

SetSocialProviders sets SocialProviders field to given value.

### HasSocialProviders

`func (o *SignOnPolicyActionIDFirstAllOf) HasSocialProviders() bool`

HasSocialProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


