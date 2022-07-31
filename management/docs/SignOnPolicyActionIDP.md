# SignOnPolicyActionIDP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcrValues** | Pointer to **string** | A string that designates the sign-on policies included in the authorization flow request. Options can include the PingOne predefined sign-on policies, Single_Factor and Multi_Factor, or any custom defined sign-on policy names. Sign-on policy names should be listed in order of preference, and they must be assigned to the application. This property can be configured on the identity provider action and is passed to the identity provider if the identity provider is of type &#x60;SAML&#x60; or &#x60;OPENID_CONNECT&#x60;. | [optional] 
**IdentityProvider** | [**SignOnPolicyActionIDPIdentityProvider**](SignOnPolicyActionIDPIdentityProvider.md) |  | 
**PassUserContext** | Pointer to **bool** | A boolean that specifies whether to pass in a login hint to the identity provider on the authentication request. Based on user context, the login hint is set if (1) the user is set on the flow, and (2) the user already has an account link for the identity provider. If both of these conditions are true, then the user is sent to the identity provider with a login hint equal to their externalId for the identity provider (saved on the account link). If these conditions are not true, then the API checks see if there is an OIDC login hint on the flow. If so, that login hint is used. If none of these conditions are true, the login hint parameter is not included on the authorization request to the identity provider. | [optional] 
**Registration** | Pointer to [**SignOnPolicyActionIDPRegistration**](SignOnPolicyActionIDPRegistration.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionIDP

`func NewSignOnPolicyActionIDP(identityProvider SignOnPolicyActionIDPIdentityProvider, ) *SignOnPolicyActionIDP`

NewSignOnPolicyActionIDP instantiates a new SignOnPolicyActionIDP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionIDPWithDefaults

`func NewSignOnPolicyActionIDPWithDefaults() *SignOnPolicyActionIDP`

NewSignOnPolicyActionIDPWithDefaults instantiates a new SignOnPolicyActionIDP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *SignOnPolicyActionIDP) GetIdentityProvider() SignOnPolicyActionIDPIdentityProvider`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *SignOnPolicyActionIDP) GetIdentityProviderOk() (*SignOnPolicyActionIDPIdentityProvider, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *SignOnPolicyActionIDP) SetIdentityProvider(v SignOnPolicyActionIDPIdentityProvider)`

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

`func (o *SignOnPolicyActionIDP) GetRegistration() SignOnPolicyActionIDPRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SignOnPolicyActionIDP) GetRegistrationOk() (*SignOnPolicyActionIDPRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SignOnPolicyActionIDP) SetRegistration(v SignOnPolicyActionIDPRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SignOnPolicyActionIDP) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


