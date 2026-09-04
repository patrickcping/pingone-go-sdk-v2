# BrandingThemeConfigurationLocalizedText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Specifies whether localization is enabled for the header or footer. | [optional] 
**DefaultLanguage** | Pointer to **string** | The localization language code for the header or footer&#39;s default language (for example, \&quot;defaultLanguage\&quot; set to \&quot;en\&quot;). | [optional] 
**ContentType** | Pointer to [**EnumBrandingThemeLocalizedContentType**](EnumBrandingThemeLocalizedContentType.md) |  | [optional] 
**Content** | Pointer to [**map[string]BrandingThemeConfigurationLocalizedTextContent**](BrandingThemeConfigurationLocalizedTextContent.md) | The localization language object that specifies the localized text and content type for the language options. | [optional] 

## Methods

### NewBrandingThemeConfigurationLocalizedText

`func NewBrandingThemeConfigurationLocalizedText() *BrandingThemeConfigurationLocalizedText`

NewBrandingThemeConfigurationLocalizedText instantiates a new BrandingThemeConfigurationLocalizedText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingThemeConfigurationLocalizedTextWithDefaults

`func NewBrandingThemeConfigurationLocalizedTextWithDefaults() *BrandingThemeConfigurationLocalizedText`

NewBrandingThemeConfigurationLocalizedTextWithDefaults instantiates a new BrandingThemeConfigurationLocalizedText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BrandingThemeConfigurationLocalizedText) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BrandingThemeConfigurationLocalizedText) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BrandingThemeConfigurationLocalizedText) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BrandingThemeConfigurationLocalizedText) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefaultLanguage

`func (o *BrandingThemeConfigurationLocalizedText) GetDefaultLanguage() string`

GetDefaultLanguage returns the DefaultLanguage field if non-nil, zero value otherwise.

### GetDefaultLanguageOk

`func (o *BrandingThemeConfigurationLocalizedText) GetDefaultLanguageOk() (*string, bool)`

GetDefaultLanguageOk returns a tuple with the DefaultLanguage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLanguage

`func (o *BrandingThemeConfigurationLocalizedText) SetDefaultLanguage(v string)`

SetDefaultLanguage sets DefaultLanguage field to given value.

### HasDefaultLanguage

`func (o *BrandingThemeConfigurationLocalizedText) HasDefaultLanguage() bool`

HasDefaultLanguage returns a boolean if a field has been set.

### GetContentType

`func (o *BrandingThemeConfigurationLocalizedText) GetContentType() EnumBrandingThemeLocalizedContentType`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *BrandingThemeConfigurationLocalizedText) GetContentTypeOk() (*EnumBrandingThemeLocalizedContentType, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *BrandingThemeConfigurationLocalizedText) SetContentType(v EnumBrandingThemeLocalizedContentType)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *BrandingThemeConfigurationLocalizedText) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetContent

`func (o *BrandingThemeConfigurationLocalizedText) GetContent() map[string]BrandingThemeConfigurationLocalizedTextContent`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *BrandingThemeConfigurationLocalizedText) GetContentOk() (*map[string]BrandingThemeConfigurationLocalizedTextContent, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *BrandingThemeConfigurationLocalizedText) SetContent(v map[string]BrandingThemeConfigurationLocalizedTextContent)`

SetContent sets Content field to given value.

### HasContent

`func (o *BrandingThemeConfigurationLocalizedText) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


