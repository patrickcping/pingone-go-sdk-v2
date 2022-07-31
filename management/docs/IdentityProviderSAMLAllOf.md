# IdentityProviderSAMLAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthnRequestSigned** | Pointer to **bool** | A boolean that specifies whether the SAML authentication request will be signed when sending to the identity provider. Set this to true if the external IDP is included in an authentication policy to be used by applications that are accessed using a mix of default URLS and custom Domains URLs. | [optional] 
**IdpEntityId** | Pointer to **string** | A string that specifies the entity ID URI that is checked against the issuerId tag in the incoming response. | [optional] 
**IdpVerification** | Pointer to [**IdentityProviderSAMLAllOfIdpVerification**](IdentityProviderSAMLAllOfIdpVerification.md) |  | [optional] 
**SpEntityId** | Pointer to **string** | A string that specifies the service provider&#39;s entity ID, used to look up the application. | [optional] 
**SpSigning** | Pointer to [**IdentityProviderSAMLAllOfSpSigning**](IdentityProviderSAMLAllOfSpSigning.md) |  | [optional] 
**SsoBinding** | Pointer to [**EnumIdentityProviderSAMLSSOBinding**](EnumIdentityProviderSAMLSSOBinding.md) |  | [optional] 
**SsoEndpoint** | Pointer to **string** | A string that specifies the SSO endpoint for the authentication request. | [optional] 

## Methods

### NewIdentityProviderSAMLAllOf

`func NewIdentityProviderSAMLAllOf() *IdentityProviderSAMLAllOf`

NewIdentityProviderSAMLAllOf instantiates a new IdentityProviderSAMLAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderSAMLAllOfWithDefaults

`func NewIdentityProviderSAMLAllOfWithDefaults() *IdentityProviderSAMLAllOf`

NewIdentityProviderSAMLAllOfWithDefaults instantiates a new IdentityProviderSAMLAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthnRequestSigned

`func (o *IdentityProviderSAMLAllOf) GetAuthnRequestSigned() bool`

GetAuthnRequestSigned returns the AuthnRequestSigned field if non-nil, zero value otherwise.

### GetAuthnRequestSignedOk

`func (o *IdentityProviderSAMLAllOf) GetAuthnRequestSignedOk() (*bool, bool)`

GetAuthnRequestSignedOk returns a tuple with the AuthnRequestSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnRequestSigned

`func (o *IdentityProviderSAMLAllOf) SetAuthnRequestSigned(v bool)`

SetAuthnRequestSigned sets AuthnRequestSigned field to given value.

### HasAuthnRequestSigned

`func (o *IdentityProviderSAMLAllOf) HasAuthnRequestSigned() bool`

HasAuthnRequestSigned returns a boolean if a field has been set.

### GetIdpEntityId

`func (o *IdentityProviderSAMLAllOf) GetIdpEntityId() string`

GetIdpEntityId returns the IdpEntityId field if non-nil, zero value otherwise.

### GetIdpEntityIdOk

`func (o *IdentityProviderSAMLAllOf) GetIdpEntityIdOk() (*string, bool)`

GetIdpEntityIdOk returns a tuple with the IdpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEntityId

`func (o *IdentityProviderSAMLAllOf) SetIdpEntityId(v string)`

SetIdpEntityId sets IdpEntityId field to given value.

### HasIdpEntityId

`func (o *IdentityProviderSAMLAllOf) HasIdpEntityId() bool`

HasIdpEntityId returns a boolean if a field has been set.

### GetIdpVerification

`func (o *IdentityProviderSAMLAllOf) GetIdpVerification() IdentityProviderSAMLAllOfIdpVerification`

GetIdpVerification returns the IdpVerification field if non-nil, zero value otherwise.

### GetIdpVerificationOk

`func (o *IdentityProviderSAMLAllOf) GetIdpVerificationOk() (*IdentityProviderSAMLAllOfIdpVerification, bool)`

GetIdpVerificationOk returns a tuple with the IdpVerification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpVerification

`func (o *IdentityProviderSAMLAllOf) SetIdpVerification(v IdentityProviderSAMLAllOfIdpVerification)`

SetIdpVerification sets IdpVerification field to given value.

### HasIdpVerification

`func (o *IdentityProviderSAMLAllOf) HasIdpVerification() bool`

HasIdpVerification returns a boolean if a field has been set.

### GetSpEntityId

`func (o *IdentityProviderSAMLAllOf) GetSpEntityId() string`

GetSpEntityId returns the SpEntityId field if non-nil, zero value otherwise.

### GetSpEntityIdOk

`func (o *IdentityProviderSAMLAllOf) GetSpEntityIdOk() (*string, bool)`

GetSpEntityIdOk returns a tuple with the SpEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpEntityId

`func (o *IdentityProviderSAMLAllOf) SetSpEntityId(v string)`

SetSpEntityId sets SpEntityId field to given value.

### HasSpEntityId

`func (o *IdentityProviderSAMLAllOf) HasSpEntityId() bool`

HasSpEntityId returns a boolean if a field has been set.

### GetSpSigning

`func (o *IdentityProviderSAMLAllOf) GetSpSigning() IdentityProviderSAMLAllOfSpSigning`

GetSpSigning returns the SpSigning field if non-nil, zero value otherwise.

### GetSpSigningOk

`func (o *IdentityProviderSAMLAllOf) GetSpSigningOk() (*IdentityProviderSAMLAllOfSpSigning, bool)`

GetSpSigningOk returns a tuple with the SpSigning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpSigning

`func (o *IdentityProviderSAMLAllOf) SetSpSigning(v IdentityProviderSAMLAllOfSpSigning)`

SetSpSigning sets SpSigning field to given value.

### HasSpSigning

`func (o *IdentityProviderSAMLAllOf) HasSpSigning() bool`

HasSpSigning returns a boolean if a field has been set.

### GetSsoBinding

`func (o *IdentityProviderSAMLAllOf) GetSsoBinding() EnumIdentityProviderSAMLSSOBinding`

GetSsoBinding returns the SsoBinding field if non-nil, zero value otherwise.

### GetSsoBindingOk

`func (o *IdentityProviderSAMLAllOf) GetSsoBindingOk() (*EnumIdentityProviderSAMLSSOBinding, bool)`

GetSsoBindingOk returns a tuple with the SsoBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoBinding

`func (o *IdentityProviderSAMLAllOf) SetSsoBinding(v EnumIdentityProviderSAMLSSOBinding)`

SetSsoBinding sets SsoBinding field to given value.

### HasSsoBinding

`func (o *IdentityProviderSAMLAllOf) HasSsoBinding() bool`

HasSsoBinding returns a boolean if a field has been set.

### GetSsoEndpoint

`func (o *IdentityProviderSAMLAllOf) GetSsoEndpoint() string`

GetSsoEndpoint returns the SsoEndpoint field if non-nil, zero value otherwise.

### GetSsoEndpointOk

`func (o *IdentityProviderSAMLAllOf) GetSsoEndpointOk() (*string, bool)`

GetSsoEndpointOk returns a tuple with the SsoEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsoEndpoint

`func (o *IdentityProviderSAMLAllOf) SetSsoEndpoint(v string)`

SetSsoEndpoint sets SsoEndpoint field to given value.

### HasSsoEndpoint

`func (o *IdentityProviderSAMLAllOf) HasSsoEndpoint() bool`

HasSsoEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


