# UpdateServiceRequest

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

### NewUpdateServiceRequest

`func NewUpdateServiceRequest(name string, serviceType string, valueType AuthorizeEditorDataValueTypeDTO, serviceSettings AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO, ) *UpdateServiceRequest`

NewUpdateServiceRequest instantiates a new UpdateServiceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServiceRequestWithDefaults

`func NewUpdateServiceRequestWithDefaults() *UpdateServiceRequest`

NewUpdateServiceRequestWithDefaults instantiates a new UpdateServiceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UpdateServiceRequest) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UpdateServiceRequest) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UpdateServiceRequest) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UpdateServiceRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEmbedded

`func (o *UpdateServiceRequest) GetEmbedded() map[string]map[string]interface{}`

GetEmbedded returns the Embedded field if non-nil, zero value otherwise.

### GetEmbeddedOk

`func (o *UpdateServiceRequest) GetEmbeddedOk() (*map[string]map[string]interface{}, bool)`

GetEmbeddedOk returns a tuple with the Embedded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbedded

`func (o *UpdateServiceRequest) SetEmbedded(v map[string]map[string]interface{})`

SetEmbedded sets Embedded field to given value.

### HasEmbedded

`func (o *UpdateServiceRequest) HasEmbedded() bool`

HasEmbedded returns a boolean if a field has been set.

### GetEnvironment

`func (o *UpdateServiceRequest) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *UpdateServiceRequest) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *UpdateServiceRequest) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *UpdateServiceRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *UpdateServiceRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateServiceRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateServiceRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateServiceRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateServiceRequest) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateServiceRequest) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateServiceRequest) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateServiceRequest) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetName

`func (o *UpdateServiceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServiceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServiceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetFullName

`func (o *UpdateServiceRequest) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *UpdateServiceRequest) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *UpdateServiceRequest) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *UpdateServiceRequest) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateServiceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServiceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServiceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServiceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *UpdateServiceRequest) GetParent() AuthorizeEditorDataReferenceObjectDTO`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *UpdateServiceRequest) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *UpdateServiceRequest) SetParent(v AuthorizeEditorDataReferenceObjectDTO)`

SetParent sets Parent field to given value.

### HasParent

`func (o *UpdateServiceRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetType

`func (o *UpdateServiceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateServiceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateServiceRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateServiceRequest) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCacheSettings

`func (o *UpdateServiceRequest) GetCacheSettings() AuthorizeEditorDataCacheSettingsDTO`

GetCacheSettings returns the CacheSettings field if non-nil, zero value otherwise.

### GetCacheSettingsOk

`func (o *UpdateServiceRequest) GetCacheSettingsOk() (*AuthorizeEditorDataCacheSettingsDTO, bool)`

GetCacheSettingsOk returns a tuple with the CacheSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheSettings

`func (o *UpdateServiceRequest) SetCacheSettings(v AuthorizeEditorDataCacheSettingsDTO)`

SetCacheSettings sets CacheSettings field to given value.

### HasCacheSettings

`func (o *UpdateServiceRequest) HasCacheSettings() bool`

HasCacheSettings returns a boolean if a field has been set.

### GetServiceType

`func (o *UpdateServiceRequest) GetServiceType() string`

GetServiceType returns the ServiceType field if non-nil, zero value otherwise.

### GetServiceTypeOk

`func (o *UpdateServiceRequest) GetServiceTypeOk() (*string, bool)`

GetServiceTypeOk returns a tuple with the ServiceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceType

`func (o *UpdateServiceRequest) SetServiceType(v string)`

SetServiceType sets ServiceType field to given value.


### GetProcessor

`func (o *UpdateServiceRequest) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *UpdateServiceRequest) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *UpdateServiceRequest) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *UpdateServiceRequest) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetValueType

`func (o *UpdateServiceRequest) GetValueType() AuthorizeEditorDataValueTypeDTO`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *UpdateServiceRequest) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *UpdateServiceRequest) SetValueType(v AuthorizeEditorDataValueTypeDTO)`

SetValueType sets ValueType field to given value.


### GetServiceSettings

`func (o *UpdateServiceRequest) GetServiceSettings() AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO`

GetServiceSettings returns the ServiceSettings field if non-nil, zero value otherwise.

### GetServiceSettingsOk

`func (o *UpdateServiceRequest) GetServiceSettingsOk() (*AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO, bool)`

GetServiceSettingsOk returns a tuple with the ServiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSettings

`func (o *UpdateServiceRequest) SetServiceSettings(v AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO)`

SetServiceSettings sets ServiceSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


