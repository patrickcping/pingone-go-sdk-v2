# Form

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Name** | **string** | A string that specifies the form name, which must be provided and must be unique within an environment. | 
**Description** | Pointer to **string** | A string that specifies the description of the form. | [optional] 
**Category** | [**EnumFormCategory**](EnumFormCategory.md) |  | 
**Components** | [**FormComponents**](FormComponents.md) |  | 
**Cols** | Pointer to **int32** | An integer that specifies the number of columns in the form (min &#x3D; 1; max &#x3D; 4). | [optional] 
**MarkOptional** | **bool** | A boolean that specifies whether optional fields are highlighted in the rendered form. | 
**MarkRequired** | **bool** | A boolean that specifies whether required fields are highlighted in the rendered form. | 
**TranslationMethod** | Pointer to [**EnumFormTranslationMethod**](EnumFormTranslationMethod.md) |  | [optional] 
**FieldTypes** | Pointer to [**[]EnumFormFieldType**](EnumFormFieldType.md) | A read-only object that specifies the list of the FormField types in the form. | [optional] [readonly] 
**LanguageBundle** | Pointer to **map[string]string** | An object that provides a map of i18n keys to their translations. This object includes both the keys and their default translations. The PingOne language management service finds this object, and creates the new keys for translation for this form. | [optional] 
**Created** | Pointer to **time.Time** | The date the resouce was created (ISO-8061 format). | [optional] [readonly] 
**Modified** | Pointer to **time.Time** | The date the resouce was modified (ISO-8061 format). | [optional] [readonly] 

## Methods

### NewForm

`func NewForm(name string, category EnumFormCategory, components FormComponents, markOptional bool, markRequired bool, ) *Form`

NewForm instantiates a new Form object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormWithDefaults

`func NewFormWithDefaults() *Form`

NewFormWithDefaults instantiates a new Form object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *Form) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Form) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Form) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Form) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *Form) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Form) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Form) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Form) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *Form) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Form) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Form) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Form) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetName

`func (o *Form) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Form) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Form) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Form) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Form) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Form) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Form) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *Form) GetCategory() EnumFormCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Form) GetCategoryOk() (*EnumFormCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Form) SetCategory(v EnumFormCategory)`

SetCategory sets Category field to given value.


### GetComponents

`func (o *Form) GetComponents() FormComponents`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *Form) GetComponentsOk() (*FormComponents, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *Form) SetComponents(v FormComponents)`

SetComponents sets Components field to given value.


### GetCols

`func (o *Form) GetCols() int32`

GetCols returns the Cols field if non-nil, zero value otherwise.

### GetColsOk

`func (o *Form) GetColsOk() (*int32, bool)`

GetColsOk returns a tuple with the Cols field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCols

`func (o *Form) SetCols(v int32)`

SetCols sets Cols field to given value.

### HasCols

`func (o *Form) HasCols() bool`

HasCols returns a boolean if a field has been set.

### GetMarkOptional

`func (o *Form) GetMarkOptional() bool`

GetMarkOptional returns the MarkOptional field if non-nil, zero value otherwise.

### GetMarkOptionalOk

`func (o *Form) GetMarkOptionalOk() (*bool, bool)`

GetMarkOptionalOk returns a tuple with the MarkOptional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkOptional

`func (o *Form) SetMarkOptional(v bool)`

SetMarkOptional sets MarkOptional field to given value.


### GetMarkRequired

`func (o *Form) GetMarkRequired() bool`

GetMarkRequired returns the MarkRequired field if non-nil, zero value otherwise.

### GetMarkRequiredOk

`func (o *Form) GetMarkRequiredOk() (*bool, bool)`

GetMarkRequiredOk returns a tuple with the MarkRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkRequired

`func (o *Form) SetMarkRequired(v bool)`

SetMarkRequired sets MarkRequired field to given value.


### GetTranslationMethod

`func (o *Form) GetTranslationMethod() EnumFormTranslationMethod`

GetTranslationMethod returns the TranslationMethod field if non-nil, zero value otherwise.

### GetTranslationMethodOk

`func (o *Form) GetTranslationMethodOk() (*EnumFormTranslationMethod, bool)`

GetTranslationMethodOk returns a tuple with the TranslationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationMethod

`func (o *Form) SetTranslationMethod(v EnumFormTranslationMethod)`

SetTranslationMethod sets TranslationMethod field to given value.

### HasTranslationMethod

`func (o *Form) HasTranslationMethod() bool`

HasTranslationMethod returns a boolean if a field has been set.

### GetFieldTypes

`func (o *Form) GetFieldTypes() []EnumFormFieldType`

GetFieldTypes returns the FieldTypes field if non-nil, zero value otherwise.

### GetFieldTypesOk

`func (o *Form) GetFieldTypesOk() (*[]EnumFormFieldType, bool)`

GetFieldTypesOk returns a tuple with the FieldTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldTypes

`func (o *Form) SetFieldTypes(v []EnumFormFieldType)`

SetFieldTypes sets FieldTypes field to given value.

### HasFieldTypes

`func (o *Form) HasFieldTypes() bool`

HasFieldTypes returns a boolean if a field has been set.

### GetLanguageBundle

`func (o *Form) GetLanguageBundle() map[string]string`

GetLanguageBundle returns the LanguageBundle field if non-nil, zero value otherwise.

### GetLanguageBundleOk

`func (o *Form) GetLanguageBundleOk() (*map[string]string, bool)`

GetLanguageBundleOk returns a tuple with the LanguageBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageBundle

`func (o *Form) SetLanguageBundle(v map[string]string)`

SetLanguageBundle sets LanguageBundle field to given value.

### HasLanguageBundle

`func (o *Form) HasLanguageBundle() bool`

HasLanguageBundle returns a boolean if a field has been set.

### GetCreated

`func (o *Form) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Form) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Form) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Form) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetModified

`func (o *Form) GetModified() time.Time`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *Form) GetModifiedOk() (*time.Time, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *Form) SetModified(v time.Time)`

SetModified sets Modified field to given value.

### HasModified

`func (o *Form) HasModified() bool`

HasModified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


