# AuthorizeEditorDataServicesHttpServiceDefinitionDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**ValueType** | [**AuthorizeEditorDataValueTypeDTO**](AuthorizeEditorDataValueTypeDTO.md) |  | 
**ServiceSettings** | [**AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO**](AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO.md) |  | 

## Methods

### NewAuthorizeEditorDataServicesHttpServiceDefinitionDTO

`func NewAuthorizeEditorDataServicesHttpServiceDefinitionDTO(valueType AuthorizeEditorDataValueTypeDTO, serviceSettings AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO, ) *AuthorizeEditorDataServicesHttpServiceDefinitionDTO`

NewAuthorizeEditorDataServicesHttpServiceDefinitionDTO instantiates a new AuthorizeEditorDataServicesHttpServiceDefinitionDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataServicesHttpServiceDefinitionDTOWithDefaults

`func NewAuthorizeEditorDataServicesHttpServiceDefinitionDTOWithDefaults() *AuthorizeEditorDataServicesHttpServiceDefinitionDTO`

NewAuthorizeEditorDataServicesHttpServiceDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataServicesHttpServiceDefinitionDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetValueType

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) GetValueType() AuthorizeEditorDataValueTypeDTO`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) GetValueTypeOk() (*AuthorizeEditorDataValueTypeDTO, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) SetValueType(v AuthorizeEditorDataValueTypeDTO)`

SetValueType sets ValueType field to given value.


### GetServiceSettings

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) GetServiceSettings() AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO`

GetServiceSettings returns the ServiceSettings field if non-nil, zero value otherwise.

### GetServiceSettingsOk

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) GetServiceSettingsOk() (*AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO, bool)`

GetServiceSettingsOk returns a tuple with the ServiceSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceSettings

`func (o *AuthorizeEditorDataServicesHttpServiceDefinitionDTO) SetServiceSettings(v AuthorizeEditorDataServiceSettingsHttpServiceSettingsDTO)`

SetServiceSettings sets ServiceSettings field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


