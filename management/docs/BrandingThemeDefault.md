# BrandingThemeDefault

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Default** | **bool** | A boolean to specify whether the theme is the default in the environment | 

## Methods

### NewBrandingThemeDefault

`func NewBrandingThemeDefault(default_ bool, ) *BrandingThemeDefault`

NewBrandingThemeDefault instantiates a new BrandingThemeDefault object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingThemeDefaultWithDefaults

`func NewBrandingThemeDefaultWithDefaults() *BrandingThemeDefault`

NewBrandingThemeDefaultWithDefaults instantiates a new BrandingThemeDefault object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *BrandingThemeDefault) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BrandingThemeDefault) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BrandingThemeDefault) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *BrandingThemeDefault) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDefault

`func (o *BrandingThemeDefault) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *BrandingThemeDefault) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *BrandingThemeDefault) SetDefault(v bool)`

SetDefault sets Default field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


