# AuthorizeEditorDataServicesConnectorServiceDefinitionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Processor** | Pointer to [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | [optional] 
**ValueType** | [**AuthorizeEditorDataValueTypeDTO**](AuthorizeEditorDataValueTypeDTO.md) |  | 
**ServiceSettings** | [**AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO**](AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO

`func NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO(valueType AuthorizeEditorDataValueTypeDTO, serviceSettings AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO, ) *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO`

NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTO instantiates a new AuthorizeEditorDataServicesConnectorServiceDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTOWithDefaults

`func NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTOWithDefaults() *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO`

NewAuthorizeEditorDataServicesConnectorServiceDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataServicesConnectorServiceDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetProcessor

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetValueType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetValueType() AuthorizeEditorDataValueTypeDTO`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetValueType(v AuthorizeEditorDataValueTypeDTO)`

SetValueType sets ValueType field to given value.


### GetServiceSettings

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetServiceSettings() AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO`

GetServiceSettings returns the ServiceSettings field if non-nil, zero value otherwise.

### GetServiceSettingsOk

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) GetServiceSettingsOk() (*AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO, bool)`

GetServiceSettingsOk returns a tuple with the ServiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSettings

`func (o *AuthorizeEditorDataServicesConnectorServiceDefinitionDTO) SetServiceSettings(v AuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTO)`

SetServiceSettings sets ServiceSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


