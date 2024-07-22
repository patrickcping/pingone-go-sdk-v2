# IntegrationVersionCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to [**IntegrationVersionCommonConfiguration**](IntegrationVersionCommonConfiguration.md) |  | [optional] 
**Description** | Pointer to **string** | Unicode characters. The description of this integration metadata version. | [optional] 
**Id** | Pointer to **string** | The platform-generated ID of this integration metadata version. | [optional] [readonly] 
**Integration** | Pointer to [**IntegrationVersionCommonIntegration**](IntegrationVersionCommonIntegration.md) |  | [optional] 
**Name** | **string** | A unique name for the integration metadata version. | 
**Number** | **string** | A unique number for the integration version. | 
**Type** | Pointer to [**EnumIntegrationVersionType**](EnumIntegrationVersionType.md) |  | [optional] 

## Methods

### NewIntegrationVersionCommon

`func NewIntegrationVersionCommon(name string, number string, ) *IntegrationVersionCommon`

NewIntegrationVersionCommon instantiates a new IntegrationVersionCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationVersionCommonWithDefaults

`func NewIntegrationVersionCommonWithDefaults() *IntegrationVersionCommon`

NewIntegrationVersionCommonWithDefaults instantiates a new IntegrationVersionCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *IntegrationVersionCommon) GetConfiguration() IntegrationVersionCommonConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *IntegrationVersionCommon) GetConfigurationOk() (*IntegrationVersionCommonConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *IntegrationVersionCommon) SetConfiguration(v IntegrationVersionCommonConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *IntegrationVersionCommon) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetDescription

`func (o *IntegrationVersionCommon) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IntegrationVersionCommon) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IntegrationVersionCommon) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IntegrationVersionCommon) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *IntegrationVersionCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationVersionCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationVersionCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationVersionCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIntegration

`func (o *IntegrationVersionCommon) GetIntegration() IntegrationVersionCommonIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *IntegrationVersionCommon) GetIntegrationOk() (*IntegrationVersionCommonIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *IntegrationVersionCommon) SetIntegration(v IntegrationVersionCommonIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *IntegrationVersionCommon) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetName

`func (o *IntegrationVersionCommon) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationVersionCommon) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationVersionCommon) SetName(v string)`

SetName sets Name field to given value.


### GetNumber

`func (o *IntegrationVersionCommon) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *IntegrationVersionCommon) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *IntegrationVersionCommon) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *IntegrationVersionCommon) GetType() EnumIntegrationVersionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationVersionCommon) GetTypeOk() (*EnumIntegrationVersionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationVersionCommon) SetType(v EnumIntegrationVersionType)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationVersionCommon) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


