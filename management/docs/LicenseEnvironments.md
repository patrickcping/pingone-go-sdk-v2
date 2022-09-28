# LicenseEnvironments

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowConnections** | Pointer to **bool** | A boolean that specifies whether the license supports creation of application connections in the specified environment. | [optional] 
**AllowCustomDomain** | Pointer to **bool** | A read-only boolean that specifies whether the license supports creation of a custom domain in the specified environment. | [optional] 
**AllowCustomSchema** | Pointer to **bool** | A read-only boolean that specifies whether the license supports using custom schema attributes in the specified environment. | [optional] 
**AllowProduction** | Pointer to **bool** | A read-only boolean that specifies whether production environments are allowed. | [optional] 
**Max** | Pointer to **int32** | A read-only integer that specifies the maximum number of environments allowed. | [optional] 
**Regions** | Pointer to [**EnumRegionCode**](EnumRegionCode.md) |  | [optional] 

## Methods

### NewLicenseEnvironments

`func NewLicenseEnvironments() *LicenseEnvironments`

NewLicenseEnvironments instantiates a new LicenseEnvironments object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseEnvironmentsWithDefaults

`func NewLicenseEnvironmentsWithDefaults() *LicenseEnvironments`

NewLicenseEnvironmentsWithDefaults instantiates a new LicenseEnvironments object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowConnections

`func (o *LicenseEnvironments) GetAllowConnections() bool`

GetAllowConnections returns the AllowConnections field if non-nil, zero value otherwise.

### GetAllowConnectionsOk

`func (o *LicenseEnvironments) GetAllowConnectionsOk() (*bool, bool)`

GetAllowConnectionsOk returns a tuple with the AllowConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowConnections

`func (o *LicenseEnvironments) SetAllowConnections(v bool)`

SetAllowConnections sets AllowConnections field to given value.

### HasAllowConnections

`func (o *LicenseEnvironments) HasAllowConnections() bool`

HasAllowConnections returns a boolean if a field has been set.

### GetAllowCustomDomain

`func (o *LicenseEnvironments) GetAllowCustomDomain() bool`

GetAllowCustomDomain returns the AllowCustomDomain field if non-nil, zero value otherwise.

### GetAllowCustomDomainOk

`func (o *LicenseEnvironments) GetAllowCustomDomainOk() (*bool, bool)`

GetAllowCustomDomainOk returns a tuple with the AllowCustomDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomDomain

`func (o *LicenseEnvironments) SetAllowCustomDomain(v bool)`

SetAllowCustomDomain sets AllowCustomDomain field to given value.

### HasAllowCustomDomain

`func (o *LicenseEnvironments) HasAllowCustomDomain() bool`

HasAllowCustomDomain returns a boolean if a field has been set.

### GetAllowCustomSchema

`func (o *LicenseEnvironments) GetAllowCustomSchema() bool`

GetAllowCustomSchema returns the AllowCustomSchema field if non-nil, zero value otherwise.

### GetAllowCustomSchemaOk

`func (o *LicenseEnvironments) GetAllowCustomSchemaOk() (*bool, bool)`

GetAllowCustomSchemaOk returns a tuple with the AllowCustomSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowCustomSchema

`func (o *LicenseEnvironments) SetAllowCustomSchema(v bool)`

SetAllowCustomSchema sets AllowCustomSchema field to given value.

### HasAllowCustomSchema

`func (o *LicenseEnvironments) HasAllowCustomSchema() bool`

HasAllowCustomSchema returns a boolean if a field has been set.

### GetAllowProduction

`func (o *LicenseEnvironments) GetAllowProduction() bool`

GetAllowProduction returns the AllowProduction field if non-nil, zero value otherwise.

### GetAllowProductionOk

`func (o *LicenseEnvironments) GetAllowProductionOk() (*bool, bool)`

GetAllowProductionOk returns a tuple with the AllowProduction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowProduction

`func (o *LicenseEnvironments) SetAllowProduction(v bool)`

SetAllowProduction sets AllowProduction field to given value.

### HasAllowProduction

`func (o *LicenseEnvironments) HasAllowProduction() bool`

HasAllowProduction returns a boolean if a field has been set.

### GetMax

`func (o *LicenseEnvironments) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *LicenseEnvironments) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *LicenseEnvironments) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *LicenseEnvironments) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetRegions

`func (o *LicenseEnvironments) GetRegions() EnumRegionCode`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *LicenseEnvironments) GetRegionsOk() (*EnumRegionCode, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *LicenseEnvironments) SetRegions(v EnumRegionCode)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *LicenseEnvironments) HasRegions() bool`

HasRegions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


