# IdentityProviderSAMLAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewIdentityProviderSAMLAllOf

`func NewIdentityProviderSAMLAllOf(idpEntityId string, idpVerification IdentityProviderSAMLAllOfIdpVerification, spEntityId string, ssoBinding EnumIdentityProviderSAMLSSOBinding, ssoEndpoint string, ) *IdentityProviderSAMLAllOf`

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


### GetSloBinding

`func (o *IdentityProviderSAMLAllOf) GetSloBinding() EnumIdentityProviderSAMLSLOBinding`

GetSloBinding returns the SloBinding field if non-nil, zero value otherwise.

### GetSloBindingOk

`func (o *IdentityProviderSAMLAllOf) GetSloBindingOk() (*EnumIdentityProviderSAMLSLOBinding, bool)`

GetSloBindingOk returns a tuple with the SloBinding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloBinding

`func (o *IdentityProviderSAMLAllOf) SetSloBinding(v EnumIdentityProviderSAMLSLOBinding)`

SetSloBinding sets SloBinding field to given value.

### HasSloBinding

`func (o *IdentityProviderSAMLAllOf) HasSloBinding() bool`

HasSloBinding returns a boolean if a field has been set.

### GetSloEndpoint

`func (o *IdentityProviderSAMLAllOf) GetSloEndpoint() string`

GetSloEndpoint returns the SloEndpoint field if non-nil, zero value otherwise.

### GetSloEndpointOk

`func (o *IdentityProviderSAMLAllOf) GetSloEndpointOk() (*string, bool)`

GetSloEndpointOk returns a tuple with the SloEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloEndpoint

`func (o *IdentityProviderSAMLAllOf) SetSloEndpoint(v string)`

SetSloEndpoint sets SloEndpoint field to given value.

### HasSloEndpoint

`func (o *IdentityProviderSAMLAllOf) HasSloEndpoint() bool`

HasSloEndpoint returns a boolean if a field has been set.

### GetSloResponseEndpoint

`func (o *IdentityProviderSAMLAllOf) GetSloResponseEndpoint() string`

GetSloResponseEndpoint returns the SloResponseEndpoint field if non-nil, zero value otherwise.

### GetSloResponseEndpointOk

`func (o *IdentityProviderSAMLAllOf) GetSloResponseEndpointOk() (*string, bool)`

GetSloResponseEndpointOk returns a tuple with the SloResponseEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloResponseEndpoint

`func (o *IdentityProviderSAMLAllOf) SetSloResponseEndpoint(v string)`

SetSloResponseEndpoint sets SloResponseEndpoint field to given value.

### HasSloResponseEndpoint

`func (o *IdentityProviderSAMLAllOf) HasSloResponseEndpoint() bool`

HasSloResponseEndpoint returns a boolean if a field has been set.

### GetSloWindow

`func (o *IdentityProviderSAMLAllOf) GetSloWindow() int32`

GetSloWindow returns the SloWindow field if non-nil, zero value otherwise.

### GetSloWindowOk

`func (o *IdentityProviderSAMLAllOf) GetSloWindowOk() (*int32, bool)`

GetSloWindowOk returns a tuple with the SloWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSloWindow

`func (o *IdentityProviderSAMLAllOf) SetSloWindow(v int32)`

SetSloWindow sets SloWindow field to given value.

### HasSloWindow

`func (o *IdentityProviderSAMLAllOf) HasSloWindow() bool`

HasSloWindow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


