# IdentityProviderSAML

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the IdP. | [optional] 
**Enabled** | [**EnumEnabledStatus**](EnumEnabledStatus.md) |  | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Icon** | Pointer to [**IdentityProviderCommonIcon**](IdentityProviderCommonIcon.md) |  | [optional] 
**Id** | Pointer to **string** | The resource ID. | [optional] [readonly] 
**LoginButtonIcon** | Pointer to [**IdentityProviderCommonLoginButtonIcon**](IdentityProviderCommonLoginButtonIcon.md) |  | [optional] 
**Name** | **string** | The name of the IdP. | 
**Registration** | Pointer to [**IdentityProviderCommonRegistration**](IdentityProviderCommonRegistration.md) |  | [optional] 
**Type** | [**EnumIdentityProviderExt**](EnumIdentityProviderExt.md) |  | 
**AuthnRequestSigned** | Pointer to **bool** | A boolean that specifies whether the SAML authentication request will be signed when sending to the identity provider. Set this to true if the external IDP is included in an authentication policy to be used by applications that are accessed using a mix of default URLS and custom Domains URLs. | [optional] 
**IdpEntityId** | Pointer to **string** | A string that specifies the entity ID URI that is checked against the issuerId tag in the incoming response. | [optional] 
**IdpVerification** | Pointer to [**IdentityProviderSAMLAllOfIdpVerification**](IdentityProviderSAMLAllOfIdpVerification.md) |  | [optional] 
**SpEntityId** | Pointer to **string** | A string that specifies the service provider&#39;s entity ID, used to look up the application. | [optional] 
**SpSigning** | Pointer to [**IdentityProviderSAMLAllOfSpSigning**](IdentityProviderSAMLAllOfSpSigning.md) |  | [optional] 
**SsoBinding** | Pointer to [**EnumIdentityProviderSAMLSSOBinding**](EnumIdentityProviderSAMLSSOBinding.md) |  | [optional] 
**SsoEndpoint** | Pointer to **string** | A string that specifies the SSO endpoint for the authentication request. | [optional] 

## Methods

### NewIdentityProviderSAML

`func NewIdentityProviderSAML(enabled EnumEnabledStatus, name string, type_ EnumIdentityProviderExt, ) *IdentityProviderSAML`

NewIdentityProviderSAML instantiates a new IdentityProviderSAML object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderSAMLWithDefaults

`func NewIdentityProviderSAMLWithDefaults() *IdentityProviderSAML`

NewIdentityProviderSAMLWithDefaults instantiates a new IdentityProviderSAML object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

`func (o *IdentityProviderSAML) GetEnabled() EnumEnabledStatus`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IdentityProviderSAML) GetEnabledOk() (*EnumEnabledStatus, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IdentityProviderSAML) SetEnabled(v EnumEnabledStatus)`

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

### HasIdpEntityId

`func (o *IdentityProviderSAML) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

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

### HasIdpVerification

`func (o *IdentityProviderSAML) HasIdpVerification() bool`

HasIdpVerification returns a boolean if a field has been set.

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

### HasSpEntityId

`func (o *IdentityProviderSAML) HasSpEntityId() bool`

HasSpEntityId returns a boolean if a field has been set.

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

### HasSsoBinding

`func (o *IdentityProviderSAML) HasSsoBinding() bool`

HasSsoBinding returns a boolean if a field has been set.

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

### HasSsoEndpoint

`func (o *IdentityProviderSAML) HasSsoEndpoint() bool`

HasSsoEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

