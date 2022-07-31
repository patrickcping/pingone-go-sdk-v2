# SignOnPolicyActionIDFirst

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**Condition** | Pointer to [**SignOnPolicyActionCommonCondition**](SignOnPolicyActionCommonCondition.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the sign-on policy assignment resource’s unique identifier. | [optional] [readonly] 
**Priority** | **int32** | An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property. | 
**SignOnPolicy** | Pointer to [**SignOnPolicyActionCommonSignOnPolicy**](SignOnPolicyActionCommonSignOnPolicy.md) |  | [optional] 
**Type** | [**EnumSignOnPolicyType**](EnumSignOnPolicyType.md) |  | 
**ConfirmIdentityProviderAttributes** | Pointer to **bool** | A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user&#39;s profile during account creation. This is an optional property. If omitted, the default value is set to false. | [optional] [default to false]
**DiscoveryRules** | Pointer to [**[]SignOnPolicyActionIDFirstAllOfDiscoveryRules**](SignOnPolicyActionIDFirstAllOfDiscoveryRules.md) | The list of IDP discovery rules that are evaluated in order when no user is associated with the user identifier. The maximum number of rules is 100. The condition on which this identity provider is used to authenticate the user is expressed using the PingOne policy condition language | [optional] 
**EnforceLockoutForIdentityProviders** | Pointer to **bool** | A boolean that if set to true and if the user&#39;s account is locked (the account.canAuthenticate attribute is set to false), then social sign on with an external identity provider is prevented. | [optional] 
**Recovery** | Pointer to [**SignOnPolicyActionLoginAllOfRecovery**](SignOnPolicyActionLoginAllOfRecovery.md) |  | [optional] 
**Registration** | Pointer to [**SignOnPolicyActionLoginAllOfRegistration**](SignOnPolicyActionLoginAllOfRegistration.md) |  | [optional] 
**SocialProviders** | Pointer to [**[]SignOnPolicyActionLoginAllOfSocialProviders**](SignOnPolicyActionLoginAllOfSocialProviders.md) | An array of strings that specifies the IDs of the identity providers that can be used for the social login sign-on flow. | [optional] 

## Methods

### NewSignOnPolicyActionIDFirst

`func NewSignOnPolicyActionIDFirst(priority int32, type_ EnumSignOnPolicyType, ) *SignOnPolicyActionIDFirst`

NewSignOnPolicyActionIDFirst instantiates a new SignOnPolicyActionIDFirst object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionIDFirstWithDefaults

`func NewSignOnPolicyActionIDFirstWithDefaults() *SignOnPolicyActionIDFirst`

NewSignOnPolicyActionIDFirstWithDefaults instantiates a new SignOnPolicyActionIDFirst object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SignOnPolicyActionIDFirst) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SignOnPolicyActionIDFirst) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SignOnPolicyActionIDFirst) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SignOnPolicyActionIDFirst) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCondition

`func (o *SignOnPolicyActionIDFirst) GetCondition() SignOnPolicyActionCommonCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SignOnPolicyActionIDFirst) GetConditionOk() (*SignOnPolicyActionCommonCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SignOnPolicyActionIDFirst) SetCondition(v SignOnPolicyActionCommonCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SignOnPolicyActionIDFirst) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEnvironment

`func (o *SignOnPolicyActionIDFirst) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SignOnPolicyActionIDFirst) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SignOnPolicyActionIDFirst) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SignOnPolicyActionIDFirst) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SignOnPolicyActionIDFirst) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyActionIDFirst) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyActionIDFirst) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignOnPolicyActionIDFirst) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *SignOnPolicyActionIDFirst) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SignOnPolicyActionIDFirst) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SignOnPolicyActionIDFirst) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSignOnPolicy

`func (o *SignOnPolicyActionIDFirst) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy`

GetSignOnPolicy returns the SignOnPolicy field if non-nil, zero value otherwise.

### GetSignOnPolicyOk

`func (o *SignOnPolicyActionIDFirst) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool)`

GetSignOnPolicyOk returns a tuple with the SignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicy

`func (o *SignOnPolicyActionIDFirst) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy)`

SetSignOnPolicy sets SignOnPolicy field to given value.

### HasSignOnPolicy

`func (o *SignOnPolicyActionIDFirst) HasSignOnPolicy() bool`

HasSignOnPolicy returns a boolean if a field has been set.

### GetType

`func (o *SignOnPolicyActionIDFirst) GetType() EnumSignOnPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignOnPolicyActionIDFirst) GetTypeOk() (*EnumSignOnPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignOnPolicyActionIDFirst) SetType(v EnumSignOnPolicyType)`

SetType sets Type field to given value.


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

`func (o *SignOnPolicyActionIDFirst) GetDiscoveryRules() []SignOnPolicyActionIDFirstAllOfDiscoveryRules`

GetDiscoveryRules returns the DiscoveryRules field if non-nil, zero value otherwise.

### GetDiscoveryRulesOk

`func (o *SignOnPolicyActionIDFirst) GetDiscoveryRulesOk() (*[]SignOnPolicyActionIDFirstAllOfDiscoveryRules, bool)`

GetDiscoveryRulesOk returns a tuple with the DiscoveryRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryRules

`func (o *SignOnPolicyActionIDFirst) SetDiscoveryRules(v []SignOnPolicyActionIDFirstAllOfDiscoveryRules)`

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

`func (o *SignOnPolicyActionIDFirst) GetRecovery() SignOnPolicyActionLoginAllOfRecovery`

GetRecovery returns the Recovery field if non-nil, zero value otherwise.

### GetRecoveryOk

`func (o *SignOnPolicyActionIDFirst) GetRecoveryOk() (*SignOnPolicyActionLoginAllOfRecovery, bool)`

GetRecoveryOk returns a tuple with the Recovery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecovery

`func (o *SignOnPolicyActionIDFirst) SetRecovery(v SignOnPolicyActionLoginAllOfRecovery)`

SetRecovery sets Recovery field to given value.

### HasRecovery

`func (o *SignOnPolicyActionIDFirst) HasRecovery() bool`

HasRecovery returns a boolean if a field has been set.

### GetRegistration

`func (o *SignOnPolicyActionIDFirst) GetRegistration() SignOnPolicyActionLoginAllOfRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SignOnPolicyActionIDFirst) GetRegistrationOk() (*SignOnPolicyActionLoginAllOfRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SignOnPolicyActionIDFirst) SetRegistration(v SignOnPolicyActionLoginAllOfRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SignOnPolicyActionIDFirst) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetSocialProviders

`func (o *SignOnPolicyActionIDFirst) GetSocialProviders() []SignOnPolicyActionLoginAllOfSocialProviders`

GetSocialProviders returns the SocialProviders field if non-nil, zero value otherwise.

### GetSocialProvidersOk

`func (o *SignOnPolicyActionIDFirst) GetSocialProvidersOk() (*[]SignOnPolicyActionLoginAllOfSocialProviders, bool)`

GetSocialProvidersOk returns a tuple with the SocialProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSocialProviders

`func (o *SignOnPolicyActionIDFirst) SetSocialProviders(v []SignOnPolicyActionLoginAllOfSocialProviders)`

SetSocialProviders sets SocialProviders field to given value.

### HasSocialProviders

`func (o *SignOnPolicyActionIDFirst) HasSocialProviders() bool`

HasSocialProviders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


