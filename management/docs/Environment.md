# Environment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillOfMaterials** | Pointer to [**BillOfMaterials**](BillOfMaterials.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the population. | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**License** | [**EnvironmentLicense**](EnvironmentLicense.md) |  | 
**Name** | **string** | A string that specifies the environment name, which must be provided and must be unique within an organization. | 
**Organization** | Pointer to [**EnvironmentOrganization**](EnvironmentOrganization.md) |  | [optional] 
**Region** | [**EnumRegionCode**](EnumRegionCode.md) |  | 
**Type** | [**EnumEnvironmentType**](EnumEnvironmentType.md) |  | 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewEnvironment

`func NewEnvironment(license EnvironmentLicense, name string, region EnumRegionCode, type_ EnumEnvironmentType, ) *Environment`

NewEnvironment instantiates a new Environment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnvironmentWithDefaults

`func NewEnvironmentWithDefaults() *Environment`

NewEnvironmentWithDefaults instantiates a new Environment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillOfMaterials

`func (o *Environment) GetBillOfMaterials() BillOfMaterials`

GetBillOfMaterials returns the BillOfMaterials field if non-nil, zero value otherwise.

### GetBillOfMaterialsOk

`func (o *Environment) GetBillOfMaterialsOk() (*BillOfMaterials, bool)`

GetBillOfMaterialsOk returns a tuple with the BillOfMaterials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillOfMaterials

`func (o *Environment) SetBillOfMaterials(v BillOfMaterials)`

SetBillOfMaterials sets BillOfMaterials field to given value.

### HasBillOfMaterials

`func (o *Environment) HasBillOfMaterials() bool`

HasBillOfMaterials returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Environment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Environment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Environment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Environment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDescription

`func (o *Environment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Environment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Environment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Environment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Environment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Environment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Environment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Environment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLicense

`func (o *Environment) GetLicense() EnvironmentLicense`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *Environment) GetLicenseOk() (*EnvironmentLicense, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *Environment) SetLicense(v EnvironmentLicense)`

SetLicense sets License field to given value.


### GetName

`func (o *Environment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Environment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Environment) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *Environment) GetOrganization() EnvironmentOrganization`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Environment) GetOrganizationOk() (*EnvironmentOrganization, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Environment) SetOrganization(v EnvironmentOrganization)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Environment) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetRegion

`func (o *Environment) GetRegion() EnumRegionCode`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Environment) GetRegionOk() (*EnumRegionCode, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Environment) SetRegion(v EnumRegionCode)`

SetRegion sets Region field to given value.


### GetType

`func (o *Environment) GetType() EnumEnvironmentType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Environment) GetTypeOk() (*EnumEnvironmentType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Environment) SetType(v EnumEnvironmentType)`

SetType sets Type field to given value.


### GetUpdatedAt

`func (o *Environment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Environment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Environment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Environment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


