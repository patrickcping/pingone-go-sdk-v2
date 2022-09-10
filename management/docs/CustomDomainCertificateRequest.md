# CustomDomainCertificateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | A string that specifies the PEM-encoded certificate to import. This is a required property. The following validation is performed on the certificate: \\ It must not be expired. \\ It MUST not be self signed. \\ The custom domain name MUST match one of the subject alternative name (SAN) values on the certificate.  | 
**IntermediateCertificates** | Pointer to **string** | A string that specifies the PEM-encoded certificate chain. | [optional] 
**PrivateKey** | **string** | A string that specifies the PEM-encoded, unencrypted private key that matches the certificate&#39;s public key. This is a required property. | 

## Methods

### NewCustomDomainCertificateRequest

`func NewCustomDomainCertificateRequest(certificate string, privateKey string, ) *CustomDomainCertificateRequest`

NewCustomDomainCertificateRequest instantiates a new CustomDomainCertificateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainCertificateRequestWithDefaults

`func NewCustomDomainCertificateRequestWithDefaults() *CustomDomainCertificateRequest`

NewCustomDomainCertificateRequestWithDefaults instantiates a new CustomDomainCertificateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CustomDomainCertificateRequest) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CustomDomainCertificateRequest) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CustomDomainCertificateRequest) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetIntermediateCertificates

`func (o *CustomDomainCertificateRequest) GetIntermediateCertificates() string`

GetIntermediateCertificates returns the IntermediateCertificates field if non-nil, zero value otherwise.

### GetIntermediateCertificatesOk

`func (o *CustomDomainCertificateRequest) GetIntermediateCertificatesOk() (*string, bool)`

GetIntermediateCertificatesOk returns a tuple with the IntermediateCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateCertificates

`func (o *CustomDomainCertificateRequest) SetIntermediateCertificates(v string)`

SetIntermediateCertificates sets IntermediateCertificates field to given value.

### HasIntermediateCertificates

`func (o *CustomDomainCertificateRequest) HasIntermediateCertificates() bool`

HasIntermediateCertificates returns a boolean if a field has been set.

### GetPrivateKey

`func (o *CustomDomainCertificateRequest) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *CustomDomainCertificateRequest) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *CustomDomainCertificateRequest) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


