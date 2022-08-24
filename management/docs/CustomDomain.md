# CustomDomain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to [**CustomDomainCertificate**](CustomDomainCertificate.md) |  | [optional] 
**CanonicalName** | Pointer to **string** | A string that specifies the domain name that should be used as the value of the CNAME record in the customer’s DNS. | [optional] [readonly] 
**DomainName** | **string** | A string that specifies the resource name, which must be provided and must be unique within an environment (for example, auth.shopco.com). This is a required property. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Status** | Pointer to [**EnumCustomDomainStatus**](EnumCustomDomainStatus.md) |  | [optional] 

## Methods

### NewCustomDomain

`func NewCustomDomain(domainName string, ) *CustomDomain`

NewCustomDomain instantiates a new CustomDomain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomDomainWithDefaults

`func NewCustomDomainWithDefaults() *CustomDomain`

NewCustomDomainWithDefaults instantiates a new CustomDomain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CustomDomain) GetCertificate() CustomDomainCertificate`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CustomDomain) GetCertificateOk() (*CustomDomainCertificate, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CustomDomain) SetCertificate(v CustomDomainCertificate)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CustomDomain) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetCanonicalName

`func (o *CustomDomain) GetCanonicalName() string`

GetCanonicalName returns the CanonicalName field if non-nil, zero value otherwise.

### GetCanonicalNameOk

`func (o *CustomDomain) GetCanonicalNameOk() (*string, bool)`

GetCanonicalNameOk returns a tuple with the CanonicalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalName

`func (o *CustomDomain) SetCanonicalName(v string)`

SetCanonicalName sets CanonicalName field to given value.

### HasCanonicalName

`func (o *CustomDomain) HasCanonicalName() bool`

HasCanonicalName returns a boolean if a field has been set.

### GetDomainName

`func (o *CustomDomain) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *CustomDomain) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *CustomDomain) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.


### GetEnvironment

`func (o *CustomDomain) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CustomDomain) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CustomDomain) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CustomDomain) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *CustomDomain) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomDomain) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomDomain) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomDomain) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *CustomDomain) GetStatus() EnumCustomDomainStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomDomain) GetStatusOk() (*EnumCustomDomainStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomDomain) SetStatus(v EnumCustomDomainStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomDomain) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


