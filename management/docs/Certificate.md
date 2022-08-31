# Certificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | [**EnumCertificateKeyAlgorithm**](EnumCertificateKeyAlgorithm.md) |  | 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**Default** | Pointer to **bool** | Specifies whether this is the default key for the specified environment. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**ExpiresAt** | Pointer to **time.Time** | The time the key resource expires. | [optional] [readonly] 
**Id** | Pointer to **string** | Specifies the resource’s unique identifier. | [optional] [readonly] 
**IssuerDN** | Pointer to **string** | Specifies the distinguished name of the certificate issuer. | [optional] 
**KeyLength** | **int32** | Specifies the key length. For RSA keys, options are 2048, 3072, and 7680. For elliptical curve (EC) keys, options are 224, 256, and 384. | 
**Name** | **string** | Specifies the resource name. | 
**Organization** | Pointer to [**ObjectOrganization**](ObjectOrganization.md) |  | [optional] 
**SerialNumber** | Pointer to **int32** | Specifies the serial number of the key or certificate. | [optional] 
**SignatureAlgorithm** | [**EnumCertificateKeySignagureAlgorithm**](EnumCertificateKeySignagureAlgorithm.md) |  | 
**StartsAt** | Pointer to **time.Time** | The time the validity period starts. | [optional] [readonly] 
**Status** | Pointer to [**EnumCertificateKeyStatus**](EnumCertificateKeyStatus.md) |  | [optional] 
**SubjectDN** | **string** | Specifies the distinguished name of the subject being secured. | 
**UsageType** | [**EnumCertificateKeyUsageType**](EnumCertificateKeyUsageType.md) |  | 
**ValidityPeriod** | **int32** | Specifies the number of days the key is valid. | 

## Methods

### NewCertificate

`func NewCertificate(algorithm EnumCertificateKeyAlgorithm, keyLength int32, name string, signatureAlgorithm EnumCertificateKeySignagureAlgorithm, subjectDN string, usageType EnumCertificateKeyUsageType, validityPeriod int32, ) *Certificate`

NewCertificate instantiates a new Certificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateWithDefaults

`func NewCertificateWithDefaults() *Certificate`

NewCertificateWithDefaults instantiates a new Certificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *Certificate) GetAlgorithm() EnumCertificateKeyAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *Certificate) GetAlgorithmOk() (*EnumCertificateKeyAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *Certificate) SetAlgorithm(v EnumCertificateKeyAlgorithm)`

SetAlgorithm sets Algorithm field to given value.


### GetCreatedAt

`func (o *Certificate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Certificate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Certificate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Certificate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *Certificate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *Certificate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *Certificate) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *Certificate) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnvironment

`func (o *Certificate) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Certificate) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Certificate) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Certificate) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExpiresAt

`func (o *Certificate) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *Certificate) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *Certificate) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *Certificate) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetId

`func (o *Certificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Certificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Certificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Certificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuerDN

`func (o *Certificate) GetIssuerDN() string`

GetIssuerDN returns the IssuerDN field if non-nil, zero value otherwise.

### GetIssuerDNOk

`func (o *Certificate) GetIssuerDNOk() (*string, bool)`

GetIssuerDNOk returns a tuple with the IssuerDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerDN

`func (o *Certificate) SetIssuerDN(v string)`

SetIssuerDN sets IssuerDN field to given value.

### HasIssuerDN

`func (o *Certificate) HasIssuerDN() bool`

HasIssuerDN returns a boolean if a field has been set.

### GetKeyLength

`func (o *Certificate) GetKeyLength() int32`

GetKeyLength returns the KeyLength field if non-nil, zero value otherwise.

### GetKeyLengthOk

`func (o *Certificate) GetKeyLengthOk() (*int32, bool)`

GetKeyLengthOk returns a tuple with the KeyLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyLength

`func (o *Certificate) SetKeyLength(v int32)`

SetKeyLength sets KeyLength field to given value.


### GetName

`func (o *Certificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Certificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Certificate) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *Certificate) GetOrganization() ObjectOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Certificate) GetOrganizationOk() (*ObjectOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Certificate) SetOrganization(v ObjectOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Certificate) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetSerialNumber

`func (o *Certificate) GetSerialNumber() int32`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *Certificate) GetSerialNumberOk() (*int32, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *Certificate) SetSerialNumber(v int32)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *Certificate) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetSignatureAlgorithm

`func (o *Certificate) GetSignatureAlgorithm() EnumCertificateKeySignagureAlgorithm`

GetSignatureAlgorithm returns the SignatureAlgorithm field if non-nil, zero value otherwise.

### GetSignatureAlgorithmOk

`func (o *Certificate) GetSignatureAlgorithmOk() (*EnumCertificateKeySignagureAlgorithm, bool)`

GetSignatureAlgorithmOk returns a tuple with the SignatureAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignatureAlgorithm

`func (o *Certificate) SetSignatureAlgorithm(v EnumCertificateKeySignagureAlgorithm)`

SetSignatureAlgorithm sets SignatureAlgorithm field to given value.


### GetStartsAt

`func (o *Certificate) GetStartsAt() time.Time`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *Certificate) GetStartsAtOk() (*time.Time, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *Certificate) SetStartsAt(v time.Time)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *Certificate) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetStatus

`func (o *Certificate) GetStatus() EnumCertificateKeyStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Certificate) GetStatusOk() (*EnumCertificateKeyStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Certificate) SetStatus(v EnumCertificateKeyStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Certificate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubjectDN

`func (o *Certificate) GetSubjectDN() string`

GetSubjectDN returns the SubjectDN field if non-nil, zero value otherwise.

### GetSubjectDNOk

`func (o *Certificate) GetSubjectDNOk() (*string, bool)`

GetSubjectDNOk returns a tuple with the SubjectDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectDN

`func (o *Certificate) SetSubjectDN(v string)`

SetSubjectDN sets SubjectDN field to given value.


### GetUsageType

`func (o *Certificate) GetUsageType() EnumCertificateKeyUsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *Certificate) GetUsageTypeOk() (*EnumCertificateKeyUsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *Certificate) SetUsageType(v EnumCertificateKeyUsageType)`

SetUsageType sets UsageType field to given value.


### GetValidityPeriod

`func (o *Certificate) GetValidityPeriod() int32`

GetValidityPeriod returns the ValidityPeriod field if non-nil, zero value otherwise.

### GetValidityPeriodOk

`func (o *Certificate) GetValidityPeriodOk() (*int32, bool)`

GetValidityPeriodOk returns a tuple with the ValidityPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidityPeriod

`func (o *Certificate) SetValidityPeriod(v int32)`

SetValidityPeriod sets ValidityPeriod field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


