# CreateService201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**FullName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | [optional] 
**Type** | Pointer to [**EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType**](EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType.md) |  | [optional] 
**CacheSettings** | Pointer to [**AuthorizeEditorDataCacheSettingsDTO**](AuthorizeEditorDataCacheSettingsDTO.md) |  | [optional] 
**ServiceType** | [**EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType**](EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType.md) |  | 
**Processor** | Pointer to [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | [optional] 
**ValueType** | [**AuthorizeEditorDataValueTypeDTO**](AuthorizeEditorDataValueTypeDTO.md) |  | 
**ServiceSettings** | [**AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO**](AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO.md) |  | 

## Methods

### NewCreateService201Response

`func NewCreateService201Response(name string, serviceType EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType, valueType AuthorizeEditorDataValueTypeDTO, serviceSettings AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO, ) *CreateService201Response`

NewCreateService201Response instantiates a new CreateService201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateService201ResponseWithDefaults

`func NewCreateService201ResponseWithDefaults() *CreateService201Response`

NewCreateService201ResponseWithDefaults instantiates a new CreateService201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CreateService201Response) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateService201Response) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateService201Response) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateService201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateService201Response) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateService201Response) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateService201Response) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateService201Response) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *CreateService201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateService201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateService201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateService201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *CreateService201Response) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateService201Response) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateService201Response) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateService201Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *CreateService201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateService201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateService201Response) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *CreateService201Response) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CreateService201Response) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CreateService201Response) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *CreateService201Response) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateService201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateService201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateService201Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateService201Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *CreateService201Response) GetParent() AuthorizeEditorDataReferenceObjectDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CreateService201Response) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CreateService201Response) SetParent(v AuthorizeEditorDataReferenceObjectDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CreateService201Response) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *CreateService201Response) GetType() EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateService201Response) GetTypeOk() (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateService201Response) SetType(v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOType)`

SetType sets Type field to given value.

### HasType

`func (o *CreateService201Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCacheSettings

`func (o *CreateService201Response) GetCacheSettings() AuthorizeEditorDataCacheSettingsDTO`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *CreateService201Response) GetCacheSettingsOk() (*AuthorizeEditorDataCacheSettingsDTO, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *CreateService201Response) SetCacheSettings(v AuthorizeEditorDataCacheSettingsDTO)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *CreateService201Response) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.

### GetServiceType

`func (o *CreateService201Response) GetServiceType() EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CreateService201Response) GetServiceTypeOk() (*EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CreateService201Response) SetServiceType(v EnumAuthorizeEditorDataDefinitionsServiceDefinitionDTOServiceType)`

SetServiceType sets ServiceType field to given value.


### GetProcessor

`func (o *CreateService201Response) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *CreateService201Response) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *CreateService201Response) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *CreateService201Response) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetValueType

`func (o *CreateService201Response) GetValueType() AuthorizeEditorDataValueTypeDTO`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreateService201Response) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreateService201Response) SetValueType(v AuthorizeEditorDataValueTypeDTO)`

SetValueType sets ValueType field to given value.


### GetServiceSettings

`func (o *CreateService201Response) GetServiceSettings() AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO`

GetServiceSettings returns the ServiceSettings field if non-nil, zero value otherwise.

### GetServiceSettingsOk

`func (o *CreateService201Response) GetServiceSettingsOk() (*AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO, bool)`

GetServiceSettingsOk returns a tuple with the ServiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSettings

`func (o *CreateService201Response) SetServiceSettings(v AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO)`

SetServiceSettings sets ServiceSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


