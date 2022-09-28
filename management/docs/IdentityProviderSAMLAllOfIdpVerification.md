# IdentityProviderSAMLAllOfIdpVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | [**[]IdentityProviderSAMLAllOfIdpVerificationCertificates**](IdentityProviderSAMLAllOfIdpVerificationCertificates.md) | A array that specifies the identity provider&#39;s certificate IDs used to verify the signature on the signed assertion from the identity provider. Signing is done with a private key and verified with a public key. | 

## Methods

### NewIdentityProviderSAMLAllOfIdpVerification

`func NewIdentityProviderSAMLAllOfIdpVerification(certificates []IdentityProviderSAMLAllOfIdpVerificationCertificates, ) *IdentityProviderSAMLAllOfIdpVerification`

NewIdentityProviderSAMLAllOfIdpVerification instantiates a new IdentityProviderSAMLAllOfIdpVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderSAMLAllOfIdpVerificationWithDefaults

`func NewIdentityProviderSAMLAllOfIdpVerificationWithDefaults() *IdentityProviderSAMLAllOfIdpVerification`

NewIdentityProviderSAMLAllOfIdpVerificationWithDefaults instantiates a new IdentityProviderSAMLAllOfIdpVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *IdentityProviderSAMLAllOfIdpVerification) GetCertificates() []IdentityProviderSAMLAllOfIdpVerificationCertificates`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *IdentityProviderSAMLAllOfIdpVerification) GetCertificatesOk() (*[]IdentityProviderSAMLAllOfIdpVerificationCertificates, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *IdentityProviderSAMLAllOfIdpVerification) SetCertificates(v []IdentityProviderSAMLAllOfIdpVerificationCertificates)`

SetCertificates sets Certificates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


