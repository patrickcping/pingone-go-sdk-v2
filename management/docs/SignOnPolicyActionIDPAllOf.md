# SignOnPolicyActionIDPAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcrValues** | Pointer to **string** | A string that designates the sign-on policies included in the authorization flow request. Options can include the PingOne predefined sign-on policies, Single_Factor and Multi_Factor, or any custom defined sign-on policy names. Sign-on policy names should be listed in order of preference, and they must be assigned to the application. This property can be configured on the identity provider action and is passed to the identity provider if the identity provider is of type &#x60;SAML&#x60; or &#x60;OPENID_CONNECT&#x60;. | [optional] 
**IdentityProvider** | [**SignOnPolicyActionIDPAllOfIdentityProvider**](SignOnPolicyActionIDPAllOfIdentityProvider.md) |  | 
**PassUserContext** | Pointer to **bool** | A boolean that specifies whether to pass in a login hint to the identity provider on the authentication request. Based on user context, the login hint is set if (1) the user is set on the flow, and (2) the user already has an account link for the identity provider. If both of these conditions are true, then the user is sent to the identity provider with a login hint equal to their externalId for the identity provider (saved on the account link). If these conditions are not true, then the API checks see if there is an OIDC login hint on the flow. If so, that login hint is used. If none of these conditions are true, the login hint parameter is not included on the authorization request to the identity provider. | [optional] 
**Registration** | Pointer to [**SignOnPolicyActionIDPAllOfRegistration**](SignOnPolicyActionIDPAllOfRegistration.md) |  | [optional] 

## Methods

### NewSignOnPolicyActionIDPAllOf

`func NewSignOnPolicyActionIDPAllOf(identityProvider SignOnPolicyActionIDPAllOfIdentityProvider, ) *SignOnPolicyActionIDPAllOf`

NewSignOnPolicyActionIDPAllOf instantiates a new SignOnPolicyActionIDPAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyActionIDPAllOfWithDefaults

`func NewSignOnPolicyActionIDPAllOfWithDefaults() *SignOnPolicyActionIDPAllOf`

NewSignOnPolicyActionIDPAllOfWithDefaults instantiates a new SignOnPolicyActionIDPAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcrValues

`func (o *SignOnPolicyActionIDPAllOf) GetAcrValues() string`

GetAcrValues returns the AcrValues field if non-nil, zero value otherwise.

### GetAcrValuesOk

`func (o *SignOnPolicyActionIDPAllOf) GetAcrValuesOk() (*string, bool)`

GetAcrValuesOk returns a tuple with the AcrValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcrValues

`func (o *SignOnPolicyActionIDPAllOf) SetAcrValues(v string)`

SetAcrValues sets AcrValues field to given value.

### HasAcrValues

`func (o *SignOnPolicyActionIDPAllOf) HasAcrValues() bool`

HasAcrValues returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *SignOnPolicyActionIDPAllOf) GetIdentityProvider() SignOnPolicyActionIDPAllOfIdentityProvider`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *SignOnPolicyActionIDPAllOf) GetIdentityProviderOk() (*SignOnPolicyActionIDPAllOfIdentityProvider, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *SignOnPolicyActionIDPAllOf) SetIdentityProvider(v SignOnPolicyActionIDPAllOfIdentityProvider)`

SetIdentityProvider sets IdentityProvider field to given value.


### GetPassUserContext

`func (o *SignOnPolicyActionIDPAllOf) GetPassUserContext() bool`

GetPassUserContext returns the PassUserContext field if non-nil, zero value otherwise.

### GetPassUserContextOk

`func (o *SignOnPolicyActionIDPAllOf) GetPassUserContextOk() (*bool, bool)`

GetPassUserContextOk returns a tuple with the PassUserContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassUserContext

`func (o *SignOnPolicyActionIDPAllOf) SetPassUserContext(v bool)`

SetPassUserContext sets PassUserContext field to given value.

### HasPassUserContext

`func (o *SignOnPolicyActionIDPAllOf) HasPassUserContext() bool`

HasPassUserContext returns a boolean if a field has been set.

### GetRegistration

`func (o *SignOnPolicyActionIDPAllOf) GetRegistration() SignOnPolicyActionIDPAllOfRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *SignOnPolicyActionIDPAllOf) GetRegistrationOk() (*SignOnPolicyActionIDPAllOfRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *SignOnPolicyActionIDPAllOf) SetRegistration(v SignOnPolicyActionIDPAllOfRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *SignOnPolicyActionIDPAllOf) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


