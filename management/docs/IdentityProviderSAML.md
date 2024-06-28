# IdentityProviderSAML

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
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
**AuthnRequestSigned** | Pointer to **bool** | A boolean that specifies whether the SAML authentication request will be signed when sending to the identity provider. Set this to true if the external IDP is included in an authentication policy to be used by applications that are accessed using a mix of default URLS and custom Domains URLs. | [optional] 
**IdpEntityId** | **string** | A string that specifies the entity ID URI that is checked against the issuerId tag in the incoming response. | 
**IdpVerification** | [**IdentityProviderSAMLAllOfIdpVerification**](IdentityProviderSAMLAllOfIdpVerification.md) |  | 
**SpEntityId** | **string** | A string that specifies the service provider&#39;s entity ID, used to look up the application. | 
**SpSigning** | Pointer to [**IdentityProviderSAMLAllOfSpSigning**](IdentityProviderSAMLAllOfSpSigning.md) |  | [optional] 
**SsoBinding** | [**EnumIdentityProviderSAMLSSOBinding**](EnumIdentityProviderSAMLSSOBinding.md) |  | 
**SsoEndpoint** | **string** | A string that specifies the SSO endpoint for the authentication request. | 
**SloBinding** | Pointer to [**EnumIdentityProviderSAMLSLOBinding**](EnumIdentityProviderSAMLSLOBinding.md) |  | [optional] [default to ENUMIDENTITYPROVIDERSAMLSLOBINDING_POST]
**SloEndpoint** | Pointer to **string** | The logout endpoint URL. This is an optional property. However, if a &#x60;sloEndpoint&#x60; logout endpoint URL is not defined, logout actions result in an error. | [optional] 
**SloResponseEndpoint** | Pointer to **string** | The endpoint URL to submit the logout response. If a value is not provided, the &#x60;sloEndpoint&#x60; property value is used to submit SLO response. | [optional] 
**SloWindow** | Pointer to **int32** | Defines how long PingOne can exchange logout messages with the application, specifically a &#x60;LogoutRequest&#x60; from the application, since the initial request. PingOne can also send a &#x60;LogoutRequest&#x60; to the application when a single logout is initiated by the user from other session participants, such as an application or identity provider. This setting is per application. The SLO logout is separate from the user session logout that revokes all tokens. | [optional] 

## Methods

### NewIdentityProviderSAML

`func NewIdentityProviderSAML(enabled bool, name string, type_ EnumIdentityProviderExt, idpEntityId string, idpVerification IdentityProviderSAMLAllOfIdpVerification, spEntityId string, ssoBinding EnumIdentityProviderSAMLSSOBinding, ssoEndpoint string, ) *IdentityProviderSAML`

NewIdentityProviderSAML instantiates a new IdentityProviderSAML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderSAMLWithDefaults

`func NewIdentityProviderSAMLWithDefaults() *IdentityProviderSAML`

NewIdentityProviderSAMLWithDefaults instantiates a new IdentityProviderSAML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *IdentityProviderSAML) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *IdentityProviderSAML) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *IdentityProviderSAML) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *IdentityProviderSAML) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *IdentityProviderSAML) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IdentityProviderSAML) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IdentityProviderSAML) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IdentityProviderSAML) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *IdentityProviderSAML) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityProviderSAML) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityProviderSAML) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *IdentityProviderSAML) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IdentityProviderSAML) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IdentityProviderSAML) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *IdentityProviderSAML) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIcon

`func (o *IdentityProviderSAML) GetIcon() IdentityProviderCommonIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *IdentityProviderSAML) GetIconOk() (*IdentityProviderCommonIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *IdentityProviderSAML) SetIcon(v IdentityProviderCommonIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *IdentityProviderSAML) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetId

`func (o *IdentityProviderSAML) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderSAML) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderSAML) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderSAML) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLoginButtonIcon

`func (o *IdentityProviderSAML) GetLoginButtonIcon() IdentityProviderCommonLoginButtonIcon`

GetLoginButtonIcon returns the LoginButtonIcon field if non-nil, zero value otherwise.

### GetLoginButtonIconOk

`func (o *IdentityProviderSAML) GetLoginButtonIconOk() (*IdentityProviderCommonLoginButtonIcon, bool)`

GetLoginButtonIconOk returns a tuple with the LoginButtonIcon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginButtonIcon

`func (o *IdentityProviderSAML) SetLoginButtonIcon(v IdentityProviderCommonLoginButtonIcon)`

SetLoginButtonIcon sets LoginButtonIcon field to given value.

### HasLoginButtonIcon

`func (o *IdentityProviderSAML) HasLoginButtonIcon() bool`

HasLoginButtonIcon returns a boolean if a field has been set.

### GetName

`func (o *IdentityProviderSAML) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderSAML) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderSAML) SetName(v string)`

SetName sets Name field to given value.


### GetRegistration

`func (o *IdentityProviderSAML) GetRegistration() IdentityProviderCommonRegistration`

GetRegistration returns the Registration field if non-nil, zero value otherwise.

### GetRegistrationOk

`func (o *IdentityProviderSAML) GetRegistrationOk() (*IdentityProviderCommonRegistration, bool)`

GetRegistrationOk returns a tuple with the Registration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistration

`func (o *IdentityProviderSAML) SetRegistration(v IdentityProviderCommonRegistration)`

SetRegistration sets Registration field to given value.

### HasRegistration

`func (o *IdentityProviderSAML) HasRegistration() bool`

HasRegistration returns a boolean if a field has been set.

### GetType

`func (o *IdentityProviderSAML) GetType() EnumIdentityProviderExt`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IdentityProviderSAML) GetTypeOk() (*EnumIdentityProviderExt, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IdentityProviderSAML) SetType(v EnumIdentityProviderExt)`

SetType sets Type field to given value.


### GetCreatedAt

`func (o *IdentityProviderSAML) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityProviderSAML) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityProviderSAML) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityProviderSAML) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityProviderSAML) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityProviderSAML) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityProviderSAML) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityProviderSAML) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAuthnRequestSigned

`func (o *IdentityProviderSAML) GetAuthnRequestSigned() bool`

GetAuthnRequestSigned returns the AuthnRequestSigned field if non-nil, zero value otherwise.

### GetAuthnRequestSignedOk

`func (o *IdentityProviderSAML) GetAuthnRequestSignedOk() (*bool, bool)`

GetAuthnRequestSignedOk returns a tuple with the AuthnRequestSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnRequestSigned

`func (o *IdentityProviderSAML) SetAuthnRequestSigned(v bool)`

SetAuthnRequestSigned sets AuthnRequestSigned field to given value.

### HasAuthnRequestSigned

`func (o *IdentityProviderSAML) HasAuthnRequestSigned() bool`

HasAuthnRequestSigned returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IdentityProviderSAML) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IdentityProviderSAML) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IdentityProviderSAML) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.


### GetIdpVerification

`func (o *IdentityProviderSAML) GetIdpVerification() IdentityProviderSAMLAllOfIdpVerification`

GetIdpVerification returns the IdpVerification field if non-nil, zero value otherwise.

### GetIdpVerificationOk

`func (o *IdentityProviderSAML) GetIdpVerificationOk() (*IdentityProviderSAMLAllOfIdpVerification, bool)`

GetIdpVerificationOk returns a tuple with the IdpVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpVerification

`func (o *IdentityProviderSAML) SetIdpVerification(v IdentityProviderSAMLAllOfIdpVerification)`

SetIdpVerification sets IdpVerification field to given value.


### GetSpEntityId

`func (o *IdentityProviderSAML) GetSpEntityId() string`

GetSpEntityId returns the SpEntityId field if non-nil, zero value otherwise.

### GetSpEntityIdOk

`func (o *IdentityProviderSAML) GetSpEntityIdOk() (*string, bool)`

GetSpEntityIdOk returns a tuple with the SpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEntityId

`func (o *IdentityProviderSAML) SetSpEntityId(v string)`

SetSpEntityId sets SpEntityId field to given value.


### GetSpSigning

`func (o *IdentityProviderSAML) GetSpSigning() IdentityProviderSAMLAllOfSpSigning`

GetSpSigning returns the SpSigning field if non-nil, zero value otherwise.

### GetSpSigningOk

`func (o *IdentityProviderSAML) GetSpSigningOk() (*IdentityProviderSAMLAllOfSpSigning, bool)`

GetSpSigningOk returns a tuple with the SpSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpSigning

`func (o *IdentityProviderSAML) SetSpSigning(v IdentityProviderSAMLAllOfSpSigning)`

SetSpSigning sets SpSigning field to given value.

### HasSpSigning

`func (o *IdentityProviderSAML) HasSpSigning() bool`

HasSpSigning returns a boolean if a field has been set.

### GetSsoBinding

`func (o *IdentityProviderSAML) GetSsoBinding() EnumIdentityProviderSAMLSSOBinding`

GetSsoBinding returns the SsoBinding field if non-nil, zero value otherwise.

### GetSsoBindingOk

`func (o *IdentityProviderSAML) GetSsoBindingOk() (*EnumIdentityProviderSAMLSSOBinding, bool)`

GetSsoBindingOk returns a tuple with the SsoBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBinding

`func (o *IdentityProviderSAML) SetSsoBinding(v EnumIdentityProviderSAMLSSOBinding)`

SetSsoBinding sets SsoBinding field to given value.


### GetSsoEndpoint

`func (o *IdentityProviderSAML) GetSsoEndpoint() string`

GetSsoEndpoint returns the SsoEndpoint field if non-nil, zero value otherwise.

### GetSsoEndpointOk

`func (o *IdentityProviderSAML) GetSsoEndpointOk() (*string, bool)`

GetSsoEndpointOk returns a tuple with the SsoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEndpoint

`func (o *IdentityProviderSAML) SetSsoEndpoint(v string)`

SetSsoEndpoint sets SsoEndpoint field to given value.


### GetSloBinding

`func (o *IdentityProviderSAML) GetSloBinding() EnumIdentityProviderSAMLSLOBinding`

GetSloBinding returns the SloBinding field if non-nil, zero value otherwise.

### GetSloBindingOk

`func (o *IdentityProviderSAML) GetSloBindingOk() (*EnumIdentityProviderSAMLSLOBinding, bool)`

GetSloBindingOk returns a tuple with the SloBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloBinding

`func (o *IdentityProviderSAML) SetSloBinding(v EnumIdentityProviderSAMLSLOBinding)`

SetSloBinding sets SloBinding field to given value.

### HasSloBinding

`func (o *IdentityProviderSAML) HasSloBinding() bool`

HasSloBinding returns a boolean if a field has been set.

### GetSloEndpoint

`func (o *IdentityProviderSAML) GetSloEndpoint() string`

GetSloEndpoint returns the SloEndpoint field if non-nil, zero value otherwise.

### GetSloEndpointOk

`func (o *IdentityProviderSAML) GetSloEndpointOk() (*string, bool)`

GetSloEndpointOk returns a tuple with the SloEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloEndpoint

`func (o *IdentityProviderSAML) SetSloEndpoint(v string)`

SetSloEndpoint sets SloEndpoint field to given value.

### HasSloEndpoint

`func (o *IdentityProviderSAML) HasSloEndpoint() bool`

HasSloEndpoint returns a boolean if a field has been set.

### GetSloResponseEndpoint

`func (o *IdentityProviderSAML) GetSloResponseEndpoint() string`

GetSloResponseEndpoint returns the SloResponseEndpoint field if non-nil, zero value otherwise.

### GetSloResponseEndpointOk

`func (o *IdentityProviderSAML) GetSloResponseEndpointOk() (*string, bool)`

GetSloResponseEndpointOk returns a tuple with the SloResponseEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloResponseEndpoint

`func (o *IdentityProviderSAML) SetSloResponseEndpoint(v string)`

SetSloResponseEndpoint sets SloResponseEndpoint field to given value.

### HasSloResponseEndpoint

`func (o *IdentityProviderSAML) HasSloResponseEndpoint() bool`

HasSloResponseEndpoint returns a boolean if a field has been set.

### GetSloWindow

`func (o *IdentityProviderSAML) GetSloWindow() int32`

GetSloWindow returns the SloWindow field if non-nil, zero value otherwise.

### GetSloWindowOk

`func (o *IdentityProviderSAML) GetSloWindowOk() (*int32, bool)`

GetSloWindowOk returns a tuple with the SloWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloWindow

`func (o *IdentityProviderSAML) SetSloWindow(v int32)`

SetSloWindow sets SloWindow field to given value.

### HasSloWindow

`func (o *IdentityProviderSAML) HasSloWindow() bool`

HasSloWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


