# UpdateDomain200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | [**CustomDomainCertificate**](CustomDomainCertificate.md) |  | 
**IntermediateCertificates** | Pointer to **string** | A string that specifies the PEM-encoded certificate chain. | [optional] 
**PrivateKey** | **string** | A string that specifies the PEM-encoded, unencrypted private key that matches the certificate&#39;s public key. This is a required property. | 
**CanonicalName** | Pointer to **string** | A string that specifies the domain name that should be used as the value of the CNAME record in the customer’s DNS. | [optional] [readonly] 
**DomainName** | **string** | A string that specifies the resource name, which must be provided and must be unique within an environment (for example, auth.shopco.com). This is a required property. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Status** | Pointer to [**EnumCustomDomainStatus**](EnumCustomDomainStatus.md) |  | [optional] 

## Methods

### NewUpdateDomain200Response

`func NewUpdateDomain200Response(certificate CustomDomainCertificate, privateKey string, domainName string, ) *UpdateDomain200Response`

NewUpdateDomain200Response instantiates a new UpdateDomain200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDomain200ResponseWithDefaults

`func NewUpdateDomain200ResponseWithDefaults() *UpdateDomain200Response`

NewUpdateDomain200ResponseWithDefaults instantiates a new UpdateDomain200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *UpdateDomain200Response) GetCertificate() CustomDomainCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *UpdateDomain200Response) GetCertificateOk() (*CustomDomainCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *UpdateDomain200Response) SetCertificate(v CustomDomainCertificate)`

SetCertificate sets Certificate field to given value.


### GetIntermediateCertificates

`func (o *UpdateDomain200Response) GetIntermediateCertificates() string`

GetIntermediateCertificates returns the IntermediateCertificates field if non-nil, zero value otherwise.

### GetIntermediateCertificatesOk

`func (o *UpdateDomain200Response) GetIntermediateCertificatesOk() (*string, bool)`

GetIntermediateCertificatesOk returns a tuple with the IntermediateCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntermediateCertificates

`func (o *UpdateDomain200Response) SetIntermediateCertificates(v string)`

SetIntermediateCertificates sets IntermediateCertificates field to given value.

### HasIntermediateCertificates

`func (o *UpdateDomain200Response) HasIntermediateCertificates() bool`

HasIntermediateCertificates returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UpdateDomain200Response) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UpdateDomain200Response) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UpdateDomain200Response) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.


### GetCanonicalName

`func (o *UpdateDomain200Response) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *UpdateDomain200Response) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *UpdateDomain200Response) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *UpdateDomain200Response) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### GetDomainName

`func (o *UpdateDomain200Response) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UpdateDomain200Response) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UpdateDomain200Response) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetEnvironment

`func (o *UpdateDomain200Response) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateDomain200Response) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateDomain200Response) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateDomain200Response) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *UpdateDomain200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateDomain200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateDomain200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateDomain200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateDomain200Response) GetStatus() EnumCustomDomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateDomain200Response) GetStatusOk() (*EnumCustomDomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateDomain200Response) SetStatus(v EnumCustomDomainStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateDomain200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


