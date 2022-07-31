# IdentityProviderOIDCAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationEndpoint** | **string** | A string that specifies the the OIDC identity provider&#39;s authorization endpoint. This value must be a URL that uses https. This is a required property. | 
**ClientId** | **string** | A string that specifies the application ID from the OIDC identity provider. This is a required property. | 
**ClientSecret** | **string** | A string that specifies the application secret from the OIDC identity provider. This is a required property. | 
**DiscoveryEndpoint** | Pointer to **string** | A string that specifies the OIDC identity provider&#39;s discovery endpoint. This value must be a URL that uses https. | [optional] 
**Issuer** | **string** | A string that specifies the issuer to which the authentication is sent for the OIDC identity provider. This value must be a URL that uses https. This is a required property. | 
**JwksEndpoint** | **string** | A string that specifies the OIDC identity provider&#39;s jwks endpoint. This value must be a URL that uses https. This is a required property. | 
**Scopes** | **[]string** | An array that specifies the scopes to include in the authentication request to the OIDC identity provider. This is a required property. | 
**TokenEndpoint** | **string** | A string that specifies the OIDC identity provider&#39;s token endpoint. This is a required property. | 
**TokenEndpointAuthMethod** | [**EnumIdentityProviderOIDCTokenAuthMethod**](EnumIdentityProviderOIDCTokenAuthMethod.md) |  | [default to ENUMIDENTITYPROVIDEROIDCTOKENAUTHMETHOD_CLIENT_SECRET_BASIC]
**UserInfoEndpoint** | Pointer to **string** | A string that specifies the OIDC identity provider&#39;s userInfo endpoint. | [optional] 

## Methods

### NewIdentityProviderOIDCAllOf

`func NewIdentityProviderOIDCAllOf(authorizationEndpoint string, clientId string, clientSecret string, issuer string, jwksEndpoint string, scopes []string, tokenEndpoint string, tokenEndpointAuthMethod EnumIdentityProviderOIDCTokenAuthMethod, ) *IdentityProviderOIDCAllOf`

NewIdentityProviderOIDCAllOf instantiates a new IdentityProviderOIDCAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderOIDCAllOfWithDefaults

`func NewIdentityProviderOIDCAllOfWithDefaults() *IdentityProviderOIDCAllOf`

NewIdentityProviderOIDCAllOfWithDefaults instantiates a new IdentityProviderOIDCAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationEndpoint

`func (o *IdentityProviderOIDCAllOf) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *IdentityProviderOIDCAllOf) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *IdentityProviderOIDCAllOf) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.


### GetClientId

`func (o *IdentityProviderOIDCAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityProviderOIDCAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityProviderOIDCAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IdentityProviderOIDCAllOf) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityProviderOIDCAllOf) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityProviderOIDCAllOf) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetDiscoveryEndpoint

`func (o *IdentityProviderOIDCAllOf) GetDiscoveryEndpoint() string`

GetDiscoveryEndpoint returns the DiscoveryEndpoint field if non-nil, zero value otherwise.

### GetDiscoveryEndpointOk

`func (o *IdentityProviderOIDCAllOf) GetDiscoveryEndpointOk() (*string, bool)`

GetDiscoveryEndpointOk returns a tuple with the DiscoveryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEndpoint

`func (o *IdentityProviderOIDCAllOf) SetDiscoveryEndpoint(v string)`

SetDiscoveryEndpoint sets DiscoveryEndpoint field to given value.

### HasDiscoveryEndpoint

`func (o *IdentityProviderOIDCAllOf) HasDiscoveryEndpoint() bool`

HasDiscoveryEndpoint returns a boolean if a field has been set.

### GetIssuer

`func (o *IdentityProviderOIDCAllOf) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IdentityProviderOIDCAllOf) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IdentityProviderOIDCAllOf) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksEndpoint

`func (o *IdentityProviderOIDCAllOf) GetJwksEndpoint() string`

GetJwksEndpoint returns the JwksEndpoint field if non-nil, zero value otherwise.

### GetJwksEndpointOk

`func (o *IdentityProviderOIDCAllOf) GetJwksEndpointOk() (*string, bool)`

GetJwksEndpointOk returns a tuple with the JwksEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksEndpoint

`func (o *IdentityProviderOIDCAllOf) SetJwksEndpoint(v string)`

SetJwksEndpoint sets JwksEndpoint field to given value.


### GetScopes

`func (o *IdentityProviderOIDCAllOf) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IdentityProviderOIDCAllOf) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IdentityProviderOIDCAllOf) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetTokenEndpoint

`func (o *IdentityProviderOIDCAllOf) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *IdentityProviderOIDCAllOf) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *IdentityProviderOIDCAllOf) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetTokenEndpointAuthMethod

`func (o *IdentityProviderOIDCAllOf) GetTokenEndpointAuthMethod() EnumIdentityProviderOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *IdentityProviderOIDCAllOf) GetTokenEndpointAuthMethodOk() (*EnumIdentityProviderOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *IdentityProviderOIDCAllOf) SetTokenEndpointAuthMethod(v EnumIdentityProviderOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetUserInfoEndpoint

`func (o *IdentityProviderOIDCAllOf) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *IdentityProviderOIDCAllOf) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *IdentityProviderOIDCAllOf) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *IdentityProviderOIDCAllOf) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


