# CertificateKeyUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **bool** | Specifies whether this is the default key for the specified environment. | 
**UsageType** | [**EnumCertificateKeyUsageType**](EnumCertificateKeyUsageType.md) |  | 

## Methods

### NewCertificateKeyUpdate

`func NewCertificateKeyUpdate(default_ bool, usageType EnumCertificateKeyUsageType, ) *CertificateKeyUpdate`

NewCertificateKeyUpdate instantiates a new CertificateKeyUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCertificateKeyUpdateWithDefaults

`func NewCertificateKeyUpdateWithDefaults() *CertificateKeyUpdate`

NewCertificateKeyUpdateWithDefaults instantiates a new CertificateKeyUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *CertificateKeyUpdate) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *CertificateKeyUpdate) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *CertificateKeyUpdate) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetUsageType

`func (o *CertificateKeyUpdate) GetUsageType() EnumCertificateKeyUsageType`

GetUsageType returns the UsageType field if non-nil, zero value otherwise.

### GetUsageTypeOk

`func (o *CertificateKeyUpdate) GetUsageTypeOk() (*EnumCertificateKeyUsageType, bool)`

GetUsageTypeOk returns a tuple with the UsageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsageType

`func (o *CertificateKeyUpdate) SetUsageType(v EnumCertificateKeyUsageType)`

SetUsageType sets UsageType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


