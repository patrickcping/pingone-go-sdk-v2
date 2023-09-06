# ApplicationSAMLAllOfSpVerification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthnRequestSigned** | Pointer to **bool** | Whether the Authn Request signing should be enforced. Default is &#x60;false&#x60;. | [optional] [default to false]
**Certificates** | [**[]ApplicationSAMLAllOfSpVerificationCertificates**](ApplicationSAMLAllOfSpVerificationCertificates.md) |  | 

## Methods

### NewApplicationSAMLAllOfSpVerification

`func NewApplicationSAMLAllOfSpVerification(certificates []ApplicationSAMLAllOfSpVerificationCertificates, ) *ApplicationSAMLAllOfSpVerification`

NewApplicationSAMLAllOfSpVerification instantiates a new ApplicationSAMLAllOfSpVerification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSAMLAllOfSpVerificationWithDefaults

`func NewApplicationSAMLAllOfSpVerificationWithDefaults() *ApplicationSAMLAllOfSpVerification`

NewApplicationSAMLAllOfSpVerificationWithDefaults instantiates a new ApplicationSAMLAllOfSpVerification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthnRequestSigned

`func (o *ApplicationSAMLAllOfSpVerification) GetAuthnRequestSigned() bool`

GetAuthnRequestSigned returns the AuthnRequestSigned field if non-nil, zero value otherwise.

### GetAuthnRequestSignedOk

`func (o *ApplicationSAMLAllOfSpVerification) GetAuthnRequestSignedOk() (*bool, bool)`

GetAuthnRequestSignedOk returns a tuple with the AuthnRequestSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthnRequestSigned

`func (o *ApplicationSAMLAllOfSpVerification) SetAuthnRequestSigned(v bool)`

SetAuthnRequestSigned sets AuthnRequestSigned field to given value.

### HasAuthnRequestSigned

`func (o *ApplicationSAMLAllOfSpVerification) HasAuthnRequestSigned() bool`

HasAuthnRequestSigned returns a boolean if a field has been set.

### GetCertificates

`func (o *ApplicationSAMLAllOfSpVerification) GetCertificates() []ApplicationSAMLAllOfSpVerificationCertificates`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *ApplicationSAMLAllOfSpVerification) GetCertificatesOk() (*[]ApplicationSAMLAllOfSpVerificationCertificates, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *ApplicationSAMLAllOfSpVerification) SetCertificates(v []ApplicationSAMLAllOfSpVerificationCertificates)`

SetCertificates sets Certificates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


