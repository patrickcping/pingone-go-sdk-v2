# BrandingTheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Configuration** | [**BrandingThemeConfiguration**](BrandingThemeConfiguration.md) |  | 
**Default** | **bool** | Specifies whether this theme is the environment&#39;s default branding configuration. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | Specifies the resource’s unique identifier. | [optional] [readonly] 
**Template** | [**EnumBrandingThemeTemplate**](EnumBrandingThemeTemplate.md) |  | 

## Methods

### NewBrandingTheme

`func NewBrandingTheme(configuration BrandingThemeConfiguration, default_ bool, template EnumBrandingThemeTemplate, ) *BrandingTheme`

NewBrandingTheme instantiates a new BrandingTheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingThemeWithDefaults

`func NewBrandingThemeWithDefaults() *BrandingTheme`

NewBrandingThemeWithDefaults instantiates a new BrandingTheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BrandingTheme) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrandingTheme) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrandingTheme) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrandingTheme) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetConfiguration

`func (o *BrandingTheme) GetConfiguration() BrandingThemeConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *BrandingTheme) GetConfigurationOk() (*BrandingThemeConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *BrandingTheme) SetConfiguration(v BrandingThemeConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetDefault

`func (o *BrandingTheme) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BrandingTheme) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BrandingTheme) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetEnvironment

`func (o *BrandingTheme) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *BrandingTheme) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *BrandingTheme) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *BrandingTheme) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *BrandingTheme) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BrandingTheme) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BrandingTheme) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BrandingTheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTemplate

`func (o *BrandingTheme) GetTemplate() EnumBrandingThemeTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *BrandingTheme) GetTemplateOk() (*EnumBrandingThemeTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *BrandingTheme) SetTemplate(v EnumBrandingThemeTemplate)`

SetTemplate sets Template field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


