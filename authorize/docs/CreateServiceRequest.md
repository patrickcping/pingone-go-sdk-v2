# CreateServiceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Embedded** | Pointer to **map[string]map[string]interface{}** | HAL embedded resources | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier | [optional] [readonly] 
**Version** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**FullName** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Parent** | Pointer to [**AuthorizeEditorDataReferenceObjectDTO**](AuthorizeEditorDataReferenceObjectDTO.md) |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**CacheSettings** | Pointer to [**AuthorizeEditorDataCacheSettingsDTO**](AuthorizeEditorDataCacheSettingsDTO.md) |  | [optional] 
**ServiceType** | **string** |  | 
**Processor** | Pointer to [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | [optional] 
**ValueType** | [**AuthorizeEditorDataValueTypeDTO**](AuthorizeEditorDataValueTypeDTO.md) |  | 
**ServiceSettings** | [**AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO**](AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO.md) |  | 

## Methods

### NewCreateServiceRequest

`func NewCreateServiceRequest(name string, serviceType string, valueType AuthorizeEditorDataValueTypeDTO, serviceSettings AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO, ) *CreateServiceRequest`

NewCreateServiceRequest instantiates a new CreateServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateServiceRequestWithDefaults

`func NewCreateServiceRequestWithDefaults() *CreateServiceRequest`

NewCreateServiceRequestWithDefaults instantiates a new CreateServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CreateServiceRequest) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateServiceRequest) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateServiceRequest) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateServiceRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *CreateServiceRequest) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *CreateServiceRequest) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *CreateServiceRequest) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *CreateServiceRequest) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateServiceRequest) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateServiceRequest) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateServiceRequest) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateServiceRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *CreateServiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateServiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateServiceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateServiceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *CreateServiceRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CreateServiceRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CreateServiceRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CreateServiceRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *CreateServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *CreateServiceRequest) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *CreateServiceRequest) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *CreateServiceRequest) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *CreateServiceRequest) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *CreateServiceRequest) GetParent() AuthorizeEditorDataReferenceObjectDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *CreateServiceRequest) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *CreateServiceRequest) SetParent(v AuthorizeEditorDataReferenceObjectDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *CreateServiceRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *CreateServiceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateServiceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateServiceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateServiceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCacheSettings

`func (o *CreateServiceRequest) GetCacheSettings() AuthorizeEditorDataCacheSettingsDTO`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *CreateServiceRequest) GetCacheSettingsOk() (*AuthorizeEditorDataCacheSettingsDTO, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *CreateServiceRequest) SetCacheSettings(v AuthorizeEditorDataCacheSettingsDTO)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *CreateServiceRequest) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.

### GetServiceType

`func (o *CreateServiceRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *CreateServiceRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *CreateServiceRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetProcessor

`func (o *CreateServiceRequest) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *CreateServiceRequest) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *CreateServiceRequest) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *CreateServiceRequest) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetValueType

`func (o *CreateServiceRequest) GetValueType() AuthorizeEditorDataValueTypeDTO`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *CreateServiceRequest) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *CreateServiceRequest) SetValueType(v AuthorizeEditorDataValueTypeDTO)`

SetValueType sets ValueType field to given value.


### GetServiceSettings

`func (o *CreateServiceRequest) GetServiceSettings() AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO`

GetServiceSettings returns the ServiceSettings field if non-nil, zero value otherwise.

### GetServiceSettingsOk

`func (o *CreateServiceRequest) GetServiceSettingsOk() (*AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO, bool)`

GetServiceSettingsOk returns a tuple with the ServiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSettings

`func (o *CreateServiceRequest) SetServiceSettings(v AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO)`

SetServiceSettings sets ServiceSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


