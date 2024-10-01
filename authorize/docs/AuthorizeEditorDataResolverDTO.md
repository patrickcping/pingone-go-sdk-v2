# AuthorizeEditorDataResolverDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Condition** | Pointer to [**AuthorizeEditorDataConditionDTO**](AuthorizeEditorDataConditionDTO.md) |  | [optional] 
**Processor** | Pointer to [**AuthorizeEditorDataProcessorDTO**](AuthorizeEditorDataProcessorDTO.md) |  | [optional] 
**Type** | [**EnumAuthorizeEditorDataResolverDTOType**](EnumAuthorizeEditorDataResolverDTOType.md) |  | 

## Methods

### NewAuthorizeEditorDataResolverDTO

`func NewAuthorizeEditorDataResolverDTO(type_ EnumAuthorizeEditorDataResolverDTOType, ) *AuthorizeEditorDataResolverDTO`

NewAuthorizeEditorDataResolverDTO instantiates a new AuthorizeEditorDataResolverDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorizeEditorDataResolverDTOWithDefaults

`func NewAuthorizeEditorDataResolverDTOWithDefaults() *AuthorizeEditorDataResolverDTO`

NewAuthorizeEditorDataResolverDTOWithDefaults instantiates a new AuthorizeEditorDataResolverDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AuthorizeEditorDataResolverDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AuthorizeEditorDataResolverDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AuthorizeEditorDataResolverDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AuthorizeEditorDataResolverDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCondition

`func (o *AuthorizeEditorDataResolverDTO) GetCondition() AuthorizeEditorDataConditionDTO`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *AuthorizeEditorDataResolverDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *AuthorizeEditorDataResolverDTO) SetCondition(v AuthorizeEditorDataConditionDTO)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *AuthorizeEditorDataResolverDTO) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### GetProcessor

`func (o *AuthorizeEditorDataResolverDTO) GetProcessor() AuthorizeEditorDataProcessorDTO`

GetProcessor returns the Processor field if non-nil, zero value otherwise.

### GetProcessorOk

`func (o *AuthorizeEditorDataResolverDTO) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool)`

GetProcessorOk returns a tuple with the Processor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessor

`func (o *AuthorizeEditorDataResolverDTO) SetProcessor(v AuthorizeEditorDataProcessorDTO)`

SetProcessor sets Processor field to given value.

### HasProcessor

`func (o *AuthorizeEditorDataResolverDTO) HasProcessor() bool`

HasProcessor returns a boolean if a field has been set.

### GetType

`func (o *AuthorizeEditorDataResolverDTO) GetType() EnumAuthorizeEditorDataResolverDTOType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuthorizeEditorDataResolverDTO) GetTypeOk() (*EnumAuthorizeEditorDataResolverDTOType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuthorizeEditorDataResolverDTO) SetType(v EnumAuthorizeEditorDataResolverDTOType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


