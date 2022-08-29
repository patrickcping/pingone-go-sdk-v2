# IdentityProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Description** | Pointer to **string** | The description of the IdP. | [optional] 
**Enabled** | **bool** | The current enabled state of the IdP. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Icon** | Pointer to [**IdentityProviderCommonIcon**](IdentityProviderCommonIcon.md) |  | [optional] 
**Id** | Pointer to **string** | The resource ID. | [optional] [readonly] 
**LoginButtonIcon** | Pointer to [**IdentityProviderCommonLoginButtonIcon**](IdentityProviderCommonLoginButtonIcon.md) |  | [optional] 
**Name** | **string** | The name of the IdP. | 
**Registration** | Pointer to [**IdentityProviderCommonRegistration**](IdentityProviderCommonRegistration.md) |  | [optional] 
**Type** | [**EnumIdentityProviderExt**](EnumIdentityProviderExt.md) |  | 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**AppId** | **string** | A string that specifies the application ID from Facebook. This is a required property. | 
**AppSecret** | **string** | A string that specifies the application secret from Facebook. This is a required property. | 
**ClientId** | **string** | A string that specifies the application ID from PayPal. This is a required property. | 
**ClientSecret** | **string** | A string that specifies the application secret from PayPal. This is a required property. | 
**AuthorizationEndpoint** | **string** | A string that specifies the the OIDC identity provider&#39;s authorization endpoint. This value must be a URL that uses https. This is a required property. | 
**DiscoveryEndpoint** | Pointer to **string** | A string that specifies the OIDC identity provider&#39;s discovery endpoint. This value must be a URL that uses https. | [optional] 
**Issuer** | **string** | A string that specifies the issuer to which the authentication is sent for the OIDC identity provider. This value must be a URL that uses https. This is a required property. | 
**JwksEndpoint** | **string** | A string that specifies the OIDC identity provider&#39;s jwks endpoint. This value must be a URL that uses https. This is a required property. | 
**Scopes** | **[]string** | An array that specifies the scopes to include in the authentication request to the OIDC identity provider. This is a required property. | 
**TokenEndpoint** | **string** | A string that specifies the OIDC identity provider&#39;s token endpoint. This is a required property. | 
**TokenEndpointAuthMethod** | [**EnumIdentityProviderOIDCTokenAuthMethod**](EnumIdentityProviderOIDCTokenAuthMethod.md) |  | [default to ENUMIDENTITYPROVIDEROIDCTOKENAUTHMETHOD_CLIENT_SECRET_BASIC]
**UserInfoEndpoint** | Pointer to **string** | A string that specifies the OIDC identity provider&#39;s userInfo endpoint. | [optional] 
**ClientSecretSigningKey** | **string** | A string that specifies the private key that is used to generate a client secret. This is a required property. | 
**KeyId** | **string** | A 10-character string that Apple uses to identify an authentication key. This is a required property. | 
**TeamId** | **string** | A 10-character string that Apple uses to identify teams. This is a required property. | 
**ClientEnvironment** | **string** | A string that specifies the PayPal environment. Options are sandbox, and live. This is a required property. | 
**AuthnRequestSigned** | Pointer to **bool** | A boolean that specifies whether the SAML authentication request will be signed when sending to the identity provider. Set this to true if the external IDP is included in an authentication policy to be used by applications that are accessed using a mix of default URLS and custom Domains URLs. | [optional] 
**IdpEntityId** | Pointer to **string** | A string that specifies the entity ID URI that is checked against the issuerId tag in the incoming response. | [optional] 
**IdpVerification** | Pointer to [**IdentityProviderSAMLAllOfIdpVerification**](IdentityProviderSAMLAllOfIdpVerification.md) |  | [optional] 
**SpEntityId** | Pointer to **string** | A string that specifies the service provider&#39;s entity ID, used to look up the application. | [optional] 
**SpSigning** | Pointer to [**IdentityProviderSAMLAllOfSpSigning**](IdentityProviderSAMLAllOfSpSigning.md) |  | [optional] 
**SsoBinding** | Pointer to [**EnumIdentityProviderSAMLSSOBinding**](EnumIdentityProviderSAMLSSOBinding.md) |  | [optional] 
**SsoEndpoint** | Pointer to **string** | A string that specifies the SSO endpoint for the authentication request. | [optional] 

## Methods

### NewIdentityProvider

`func NewIdentityProvider(enabled bool, name string, type_ EnumIdentityProviderExt, appId string, appSecret string, clientId string, clientSecret string, authorizationEndpoint string, issuer string, jwksEndpoint string, scopes []string, tokenEndpoint string, tokenEndpointAuthMethod EnumIdentityProviderOIDCTokenAuthMethod, clientSecretSigningKey string, keyId string, teamId string, clientEnvironment string, ) *IdentityProvider`

NewIdentityProvider instantiates a new IdentityProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderWithDefaults

`func NewIdentityProviderWithDefaults() *IdentityProvider`

NewIdentityProviderWithDefaults instantiates a new IdentityProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IdentityProvider) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityProvider) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityProvider) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityProvider) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityProvider) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProvider) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProvider) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProvider) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityProvider) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityProvider) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityProvider) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *IdentityProvider) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IdentityProvider) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IdentityProvider) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *IdentityProvider) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIcon

`func (o *IdentityProvider) GetIcon() IdentityProviderCommonIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IdentityProvider) GetIconOk() (*IdentityProviderCommonIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IdentityProvider) SetIcon(v IdentityProviderCommonIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IdentityProvider) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *IdentityProvider) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProvider) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProvider) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginButtonIcon

`func (o *IdentityProvider) GetLoginButtonIcon() IdentityProviderCommonLoginButtonIcon`

GetLoginButtonIcon returns the LoginButtonIcon field if non-nil, zero value otherwise.

### GetLoginButtonIconOk

`func (o *IdentityProvider) GetLoginButtonIconOk() (*IdentityProviderCommonLoginButtonIcon, bool)`

GetLoginButtonIconOk returns a tuple with the LoginButtonIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonIcon

`func (o *IdentityProvider) SetLoginButtonIcon(v IdentityProviderCommonLoginButtonIcon)`

SetLoginButtonIcon sets LoginButtonIcon field to given value.

### HasLoginButtonIcon

`func (o *IdentityProvider) HasLoginButtonIcon() bool`

HasLoginButtonIcon returns a boolean if a field has been set.

### GetName

`func (o *IdentityProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProvider) SetName(v string)`

SetName sets Name field to given value.


### GetRegistration

`func (o *IdentityProvider) GetRegistration() IdentityProviderCommonRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *IdentityProvider) GetRegistrationOk() (*IdentityProviderCommonRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *IdentityProvider) SetRegistration(v IdentityProviderCommonRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *IdentityProvider) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetType

`func (o *IdentityProvider) GetType() EnumIdentityProviderExt`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityProvider) GetTypeOk() (*EnumIdentityProviderExt, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityProvider) SetType(v EnumIdentityProviderExt)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *IdentityProvider) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityProvider) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityProvider) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityProvider) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityProvider) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityProvider) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityProvider) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityProvider) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAppId

`func (o *IdentityProvider) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *IdentityProvider) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *IdentityProvider) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetAppSecret

`func (o *IdentityProvider) GetAppSecret() string`

GetAppSecret returns the AppSecret field if non-nil, zero value otherwise.

### GetAppSecretOk

`func (o *IdentityProvider) GetAppSecretOk() (*string, bool)`

GetAppSecretOk returns a tuple with the AppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppSecret

`func (o *IdentityProvider) SetAppSecret(v string)`

SetAppSecret sets AppSecret field to given value.


### GetClientId

`func (o *IdentityProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IdentityProvider) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityProvider) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityProvider) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetAuthorizationEndpoint

`func (o *IdentityProvider) GetAuthorizationEndpoint() string`

GetAuthorizationEndpoint returns the AuthorizationEndpoint field if non-nil, zero value otherwise.

### GetAuthorizationEndpointOk

`func (o *IdentityProvider) GetAuthorizationEndpointOk() (*string, bool)`

GetAuthorizationEndpointOk returns a tuple with the AuthorizationEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationEndpoint

`func (o *IdentityProvider) SetAuthorizationEndpoint(v string)`

SetAuthorizationEndpoint sets AuthorizationEndpoint field to given value.


### GetDiscoveryEndpoint

`func (o *IdentityProvider) GetDiscoveryEndpoint() string`

GetDiscoveryEndpoint returns the DiscoveryEndpoint field if non-nil, zero value otherwise.

### GetDiscoveryEndpointOk

`func (o *IdentityProvider) GetDiscoveryEndpointOk() (*string, bool)`

GetDiscoveryEndpointOk returns a tuple with the DiscoveryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEndpoint

`func (o *IdentityProvider) SetDiscoveryEndpoint(v string)`

SetDiscoveryEndpoint sets DiscoveryEndpoint field to given value.

### HasDiscoveryEndpoint

`func (o *IdentityProvider) HasDiscoveryEndpoint() bool`

HasDiscoveryEndpoint returns a boolean if a field has been set.

### GetIssuer

`func (o *IdentityProvider) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *IdentityProvider) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *IdentityProvider) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetJwksEndpoint

`func (o *IdentityProvider) GetJwksEndpoint() string`

GetJwksEndpoint returns the JwksEndpoint field if non-nil, zero value otherwise.

### GetJwksEndpointOk

`func (o *IdentityProvider) GetJwksEndpointOk() (*string, bool)`

GetJwksEndpointOk returns a tuple with the JwksEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJwksEndpoint

`func (o *IdentityProvider) SetJwksEndpoint(v string)`

SetJwksEndpoint sets JwksEndpoint field to given value.


### GetScopes

`func (o *IdentityProvider) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IdentityProvider) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IdentityProvider) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetTokenEndpoint

`func (o *IdentityProvider) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *IdentityProvider) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *IdentityProvider) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetTokenEndpointAuthMethod

`func (o *IdentityProvider) GetTokenEndpointAuthMethod() EnumIdentityProviderOIDCTokenAuthMethod`

GetTokenEndpointAuthMethod returns the TokenEndpointAuthMethod field if non-nil, zero value otherwise.

### GetTokenEndpointAuthMethodOk

`func (o *IdentityProvider) GetTokenEndpointAuthMethodOk() (*EnumIdentityProviderOIDCTokenAuthMethod, bool)`

GetTokenEndpointAuthMethodOk returns a tuple with the TokenEndpointAuthMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpointAuthMethod

`func (o *IdentityProvider) SetTokenEndpointAuthMethod(v EnumIdentityProviderOIDCTokenAuthMethod)`

SetTokenEndpointAuthMethod sets TokenEndpointAuthMethod field to given value.


### GetUserInfoEndpoint

`func (o *IdentityProvider) GetUserInfoEndpoint() string`

GetUserInfoEndpoint returns the UserInfoEndpoint field if non-nil, zero value otherwise.

### GetUserInfoEndpointOk

`func (o *IdentityProvider) GetUserInfoEndpointOk() (*string, bool)`

GetUserInfoEndpointOk returns a tuple with the UserInfoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoEndpoint

`func (o *IdentityProvider) SetUserInfoEndpoint(v string)`

SetUserInfoEndpoint sets UserInfoEndpoint field to given value.

### HasUserInfoEndpoint

`func (o *IdentityProvider) HasUserInfoEndpoint() bool`

HasUserInfoEndpoint returns a boolean if a field has been set.

### GetClientSecretSigningKey

`func (o *IdentityProvider) GetClientSecretSigningKey() string`

GetClientSecretSigningKey returns the ClientSecretSigningKey field if non-nil, zero value otherwise.

### GetClientSecretSigningKeyOk

`func (o *IdentityProvider) GetClientSecretSigningKeyOk() (*string, bool)`

GetClientSecretSigningKeyOk returns a tuple with the ClientSecretSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretSigningKey

`func (o *IdentityProvider) SetClientSecretSigningKey(v string)`

SetClientSecretSigningKey sets ClientSecretSigningKey field to given value.


### GetKeyId

`func (o *IdentityProvider) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *IdentityProvider) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *IdentityProvider) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetTeamId

`func (o *IdentityProvider) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *IdentityProvider) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *IdentityProvider) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetClientEnvironment

`func (o *IdentityProvider) GetClientEnvironment() string`

GetClientEnvironment returns the ClientEnvironment field if non-nil, zero value otherwise.

### GetClientEnvironmentOk

`func (o *IdentityProvider) GetClientEnvironmentOk() (*string, bool)`

GetClientEnvironmentOk returns a tuple with the ClientEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEnvironment

`func (o *IdentityProvider) SetClientEnvironment(v string)`

SetClientEnvironment sets ClientEnvironment field to given value.


### GetAuthnRequestSigned

`func (o *IdentityProvider) GetAuthnRequestSigned() bool`

GetAuthnRequestSigned returns the AuthnRequestSigned field if non-nil, zero value otherwise.

### GetAuthnRequestSignedOk

`func (o *IdentityProvider) GetAuthnRequestSignedOk() (*bool, bool)`

GetAuthnRequestSignedOk returns a tuple with the AuthnRequestSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnRequestSigned

`func (o *IdentityProvider) SetAuthnRequestSigned(v bool)`

SetAuthnRequestSigned sets AuthnRequestSigned field to given value.

### HasAuthnRequestSigned

`func (o *IdentityProvider) HasAuthnRequestSigned() bool`

HasAuthnRequestSigned returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IdentityProvider) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IdentityProvider) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IdentityProvider) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *IdentityProvider) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetIdpVerification

`func (o *IdentityProvider) GetIdpVerification() IdentityProviderSAMLAllOfIdpVerification`

GetIdpVerification returns the IdpVerification field if non-nil, zero value otherwise.

### GetIdpVerificationOk

`func (o *IdentityProvider) GetIdpVerificationOk() (*IdentityProviderSAMLAllOfIdpVerification, bool)`

GetIdpVerificationOk returns a tuple with the IdpVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpVerification

`func (o *IdentityProvider) SetIdpVerification(v IdentityProviderSAMLAllOfIdpVerification)`

SetIdpVerification sets IdpVerification field to given value.

### HasIdpVerification

`func (o *IdentityProvider) HasIdpVerification() bool`

HasIdpVerification returns a boolean if a field has been set.

### GetSpEntityId

`func (o *IdentityProvider) GetSpEntityId() string`

GetSpEntityId returns the SpEntityId field if non-nil, zero value otherwise.

### GetSpEntityIdOk

`func (o *IdentityProvider) GetSpEntityIdOk() (*string, bool)`

GetSpEntityIdOk returns a tuple with the SpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEntityId

`func (o *IdentityProvider) SetSpEntityId(v string)`

SetSpEntityId sets SpEntityId field to given value.

### HasSpEntityId

`func (o *IdentityProvider) HasSpEntityId() bool`

HasSpEntityId returns a boolean if a field has been set.

### GetSpSigning

`func (o *IdentityProvider) GetSpSigning() IdentityProviderSAMLAllOfSpSigning`

GetSpSigning returns the SpSigning field if non-nil, zero value otherwise.

### GetSpSigningOk

`func (o *IdentityProvider) GetSpSigningOk() (*IdentityProviderSAMLAllOfSpSigning, bool)`

GetSpSigningOk returns a tuple with the SpSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpSigning

`func (o *IdentityProvider) SetSpSigning(v IdentityProviderSAMLAllOfSpSigning)`

SetSpSigning sets SpSigning field to given value.

### HasSpSigning

`func (o *IdentityProvider) HasSpSigning() bool`

HasSpSigning returns a boolean if a field has been set.

### GetSsoBinding

`func (o *IdentityProvider) GetSsoBinding() EnumIdentityProviderSAMLSSOBinding`

GetSsoBinding returns the SsoBinding field if non-nil, zero value otherwise.

### GetSsoBindingOk

`func (o *IdentityProvider) GetSsoBindingOk() (*EnumIdentityProviderSAMLSSOBinding, bool)`

GetSsoBindingOk returns a tuple with the SsoBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBinding

`func (o *IdentityProvider) SetSsoBinding(v EnumIdentityProviderSAMLSSOBinding)`

SetSsoBinding sets SsoBinding field to given value.

### HasSsoBinding

`func (o *IdentityProvider) HasSsoBinding() bool`

HasSsoBinding returns a boolean if a field has been set.

### GetSsoEndpoint

`func (o *IdentityProvider) GetSsoEndpoint() string`

GetSsoEndpoint returns the SsoEndpoint field if non-nil, zero value otherwise.

### GetSsoEndpointOk

`func (o *IdentityProvider) GetSsoEndpointOk() (*string, bool)`

GetSsoEndpointOk returns a tuple with the SsoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEndpoint

`func (o *IdentityProvider) SetSsoEndpoint(v string)`

SetSsoEndpoint sets SsoEndpoint field to given value.

### HasSsoEndpoint

`func (o *IdentityProvider) HasSsoEndpoint() bool`

HasSsoEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


