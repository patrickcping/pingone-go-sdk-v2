# SignOnPolicyActionIDP

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
**AcrValues** | Pointer to **string** | A string that designates the sign-on policies included in the authorization flow request. Options can include the PingOne predefined sign-on policies, Single_Factor and Multi_Factor, or any custom defined sign-on policy names. Sign-on policy names should be listed in order of preference, and they must be assigned to the application. This property can be configured on the identity provider action and is passed to the identity provider if the identity provider is of type &#x60;SAML&#x60; or &#x60;OPENID_CONNECT&#x60;. | [optional] 
**IdentityProvider** | [**SignOnPolicyActionIDPAllOfIdentityProvider**](SignOnPolicyActionIDPAllOfIdentityProvider.md) |  | 
**PassUserContext** | Pointer to **bool** | A boolean that specifies whether to pass in a login hint to the identity provider on the authentication request. Based on user context, the login hint is set if (1) the user is set on the flow, and (2) the user already has an account link for the identity provider. If both of these conditions are true, then the user is sent to the identity provider with a login hint equal to their externalId for the identity provider (saved on the account link). If these conditions are not true, then the API checks see if there is an OIDC login hint on the flow. If so, that login hint is used. If none of these conditions are true, the login hint parameter is not included on the authorization request to the identity provider. | [optional] 
**Registration** | Pointer to [**SignOnPolicyActionIDPAllOfRegistration**](SignOnPolicyActionIDPAllOfRegistration.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionIDP

`func NewSignOnPolicyActionIDP(priority int32, type_ EnumSignOnPolicyType, identityProvider SignOnPolicyActionIDPAllOfIdentityProvider, ) *SignOnPolicyActionIDP`

NewSignOnPolicyActionIDP instantiates a new SignOnPolicyActionIDP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionIDPWithDefaults

`func NewSignOnPolicyActionIDPWithDefaults() *SignOnPolicyActionIDP`

NewSignOnPolicyActionIDPWithDefaults instantiates a new SignOnPolicyActionIDP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SignOnPolicyActionIDP) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SignOnPolicyActionIDP) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SignOnPolicyActionIDP) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SignOnPolicyActionIDP) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCondition

`func (o *SignOnPolicyActionIDP) GetCondition() SignOnPolicyActionCommonCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *SignOnPolicyActionIDP) GetConditionOk() (*SignOnPolicyActionCommonCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *SignOnPolicyActionIDP) SetCondition(v SignOnPolicyActionCommonCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *SignOnPolicyActionIDP) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetEnvironment

`func (o *SignOnPolicyActionIDP) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SignOnPolicyActionIDP) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SignOnPolicyActionIDP) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SignOnPolicyActionIDP) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SignOnPolicyActionIDP) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicyActionIDP) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicyActionIDP) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignOnPolicyActionIDP) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriority

`func (o *SignOnPolicyActionIDP) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *SignOnPolicyActionIDP) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *SignOnPolicyActionIDP) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetSignOnPolicy

`func (o *SignOnPolicyActionIDP) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy`

GetSignOnPolicy returns the SignOnPolicy field if non-nil, zero value otherwise.

### GetSignOnPolicyOk

`func (o *SignOnPolicyActionIDP) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool)`

GetSignOnPolicyOk returns a tuple with the SignOnPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignOnPolicy

`func (o *SignOnPolicyActionIDP) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy)`

SetSignOnPolicy sets SignOnPolicy field to given value.

### HasSignOnPolicy

`func (o *SignOnPolicyActionIDP) HasSignOnPolicy() bool`

HasSignOnPolicy returns a boolean if a field has been set.

### GetType

`func (o *SignOnPolicyActionIDP) GetType() EnumSignOnPolicyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SignOnPolicyActionIDP) GetTypeOk() (*EnumSignOnPolicyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SignOnPolicyActionIDP) SetType(v EnumSignOnPolicyType)`

SetType sets Type field to given value.


### GetAcrValues

`func (o *SignOnPolicyActionIDP) GetAcrValues() string`

GetAcrValues returns the AcrValues field if non-nil, zero value otherwise.

### GetAcrValuesOk

`func (o *SignOnPolicyActionIDP) GetAcrValuesOk() (*string, bool)`

GetAcrValuesOk returns a tuple with the AcrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrValues

`func (o *SignOnPolicyActionIDP) SetAcrValues(v string)`

SetAcrValues sets AcrValues field to given value.

### HasAcrValues

`func (o *SignOnPolicyActionIDP) HasAcrValues() bool`

HasAcrValues returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *SignOnPolicyActionIDP) GetIdentityProvider() SignOnPolicyActionIDPAllOfIdentityProvider`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *SignOnPolicyActionIDP) GetIdentityProviderOk() (*SignOnPolicyActionIDPAllOfIdentityProvider, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *SignOnPolicyActionIDP) SetIdentityProvider(v SignOnPolicyActionIDPAllOfIdentityProvider)`

SetIdentityProvider sets IdentityProvider field to given value.


### GetPassUserContext

`func (o *SignOnPolicyActionIDP) GetPassUserContext() bool`

GetPassUserContext returns the PassUserContext field if non-nil, zero value otherwise.

### GetPassUserContextOk

`func (o *SignOnPolicyActionIDP) GetPassUserContextOk() (*bool, bool)`

GetPassUserContextOk returns a tuple with the PassUserContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassUserContext

`func (o *SignOnPolicyActionIDP) SetPassUserContext(v bool)`

SetPassUserContext sets PassUserContext field to given value.

### HasPassUserContext

`func (o *SignOnPolicyActionIDP) HasPassUserContext() bool`

HasPassUserContext returns a boolean if a field has been set.

### GetRegistration

`func (o *SignOnPolicyActionIDP) GetRegistration() SignOnPolicyActionIDPAllOfRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SignOnPolicyActionIDP) GetRegistrationOk() (*SignOnPolicyActionIDPAllOfRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SignOnPolicyActionIDP) SetRegistration(v SignOnPolicyActionIDPAllOfRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SignOnPolicyActionIDP) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


