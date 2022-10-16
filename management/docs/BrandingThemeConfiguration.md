# BrandingThemeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundColor** | Pointer to **string** | The background color for the theme. It must be a valid hexadecimal color code, and it is a required property when configuration.backgroundType is set to COLOR. | [optional] 
**BackgroundType** | [**EnumBrandingThemeBackgroundType**](EnumBrandingThemeBackgroundType.md) |  | 
**BackgroundImage** | Pointer to [**BrandingThemeConfigurationBackgroundImage**](BrandingThemeConfigurationBackgroundImage.md) |  | [optional] 
**BodyTextColor** | **string** | The body text color for the theme. It must be a valid hexadecimal color code. | 
**ButtonColor** | **string** | The button color for the theme. It must be a valid hexadecimal color code. | 
**ButtonTextColor** | **string** | The button text color for the branding theme. It must be a valid hexadecimal color code. | 
**CardColor** | **string** | The card color for the branding theme. It must be a valid hexadecimal color code. | 
**Footer** | Pointer to **string** | The footer of the branding theme. | [optional] 
**HeadingTextColor** | **string** | The heading text color for the branding theme. It must be a valid hexadecimal color code. | 
**LinkTextColor** | **string** | The hyperlink text color for the branding theme. It must be a valid hexadecimal color code. | 
**Logo** | Pointer to [**BrandingThemeConfigurationLogo**](BrandingThemeConfigurationLogo.md) |  | [optional] 
**LogoType** | [**EnumBrandingLogoType**](EnumBrandingLogoType.md) |  | 
**Name** | Pointer to **string** | The name of the branding theme. | [optional] 

## Methods

### NewBrandingThemeConfiguration

`func NewBrandingThemeConfiguration(backgroundType EnumBrandingThemeBackgroundType, bodyTextColor string, buttonColor string, buttonTextColor string, cardColor string, headingTextColor string, linkTextColor string, logoType EnumBrandingLogoType, ) *BrandingThemeConfiguration`

NewBrandingThemeConfiguration instantiates a new BrandingThemeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBrandingThemeConfigurationWithDefaults

`func NewBrandingThemeConfigurationWithDefaults() *BrandingThemeConfiguration`

NewBrandingThemeConfigurationWithDefaults instantiates a new BrandingThemeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackgroundColor

`func (o *BrandingThemeConfiguration) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *BrandingThemeConfiguration) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *BrandingThemeConfiguration) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *BrandingThemeConfiguration) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetBackgroundType

`func (o *BrandingThemeConfiguration) GetBackgroundType() EnumBrandingThemeBackgroundType`

GetBackgroundType returns the BackgroundType field if non-nil, zero value otherwise.

### GetBackgroundTypeOk

`func (o *BrandingThemeConfiguration) GetBackgroundTypeOk() (*EnumBrandingThemeBackgroundType, bool)`

GetBackgroundTypeOk returns a tuple with the BackgroundType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundType

`func (o *BrandingThemeConfiguration) SetBackgroundType(v EnumBrandingThemeBackgroundType)`

SetBackgroundType sets BackgroundType field to given value.


### GetBackgroundImage

`func (o *BrandingThemeConfiguration) GetBackgroundImage() BrandingThemeConfigurationBackgroundImage`

GetBackgroundImage returns the BackgroundImage field if non-nil, zero value otherwise.

### GetBackgroundImageOk

`func (o *BrandingThemeConfiguration) GetBackgroundImageOk() (*BrandingThemeConfigurationBackgroundImage, bool)`

GetBackgroundImageOk returns a tuple with the BackgroundImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundImage

`func (o *BrandingThemeConfiguration) SetBackgroundImage(v BrandingThemeConfigurationBackgroundImage)`

SetBackgroundImage sets BackgroundImage field to given value.

### HasBackgroundImage

`func (o *BrandingThemeConfiguration) HasBackgroundImage() bool`

HasBackgroundImage returns a boolean if a field has been set.

### GetBodyTextColor

`func (o *BrandingThemeConfiguration) GetBodyTextColor() string`

GetBodyTextColor returns the BodyTextColor field if non-nil, zero value otherwise.

### GetBodyTextColorOk

`func (o *BrandingThemeConfiguration) GetBodyTextColorOk() (*string, bool)`

GetBodyTextColorOk returns a tuple with the BodyTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTextColor

`func (o *BrandingThemeConfiguration) SetBodyTextColor(v string)`

SetBodyTextColor sets BodyTextColor field to given value.


### GetButtonColor

`func (o *BrandingThemeConfiguration) GetButtonColor() string`

GetButtonColor returns the ButtonColor field if non-nil, zero value otherwise.

### GetButtonColorOk

`func (o *BrandingThemeConfiguration) GetButtonColorOk() (*string, bool)`

GetButtonColorOk returns a tuple with the ButtonColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonColor

`func (o *BrandingThemeConfiguration) SetButtonColor(v string)`

SetButtonColor sets ButtonColor field to given value.


### GetButtonTextColor

`func (o *BrandingThemeConfiguration) GetButtonTextColor() string`

GetButtonTextColor returns the ButtonTextColor field if non-nil, zero value otherwise.

### GetButtonTextColorOk

`func (o *BrandingThemeConfiguration) GetButtonTextColorOk() (*string, bool)`

GetButtonTextColorOk returns a tuple with the ButtonTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTextColor

`func (o *BrandingThemeConfiguration) SetButtonTextColor(v string)`

SetButtonTextColor sets ButtonTextColor field to given value.


### GetCardColor

`func (o *BrandingThemeConfiguration) GetCardColor() string`

GetCardColor returns the CardColor field if non-nil, zero value otherwise.

### GetCardColorOk

`func (o *BrandingThemeConfiguration) GetCardColorOk() (*string, bool)`

GetCardColorOk returns a tuple with the CardColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardColor

`func (o *BrandingThemeConfiguration) SetCardColor(v string)`

SetCardColor sets CardColor field to given value.


### GetFooter

`func (o *BrandingThemeConfiguration) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *BrandingThemeConfiguration) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *BrandingThemeConfiguration) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *BrandingThemeConfiguration) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetHeadingTextColor

`func (o *BrandingThemeConfiguration) GetHeadingTextColor() string`

GetHeadingTextColor returns the HeadingTextColor field if non-nil, zero value otherwise.

### GetHeadingTextColorOk

`func (o *BrandingThemeConfiguration) GetHeadingTextColorOk() (*string, bool)`

GetHeadingTextColorOk returns a tuple with the HeadingTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadingTextColor

`func (o *BrandingThemeConfiguration) SetHeadingTextColor(v string)`

SetHeadingTextColor sets HeadingTextColor field to given value.


### GetLinkTextColor

`func (o *BrandingThemeConfiguration) GetLinkTextColor() string`

GetLinkTextColor returns the LinkTextColor field if non-nil, zero value otherwise.

### GetLinkTextColorOk

`func (o *BrandingThemeConfiguration) GetLinkTextColorOk() (*string, bool)`

GetLinkTextColorOk returns a tuple with the LinkTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTextColor

`func (o *BrandingThemeConfiguration) SetLinkTextColor(v string)`

SetLinkTextColor sets LinkTextColor field to given value.


### GetLogo

`func (o *BrandingThemeConfiguration) GetLogo() BrandingThemeConfigurationLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *BrandingThemeConfiguration) GetLogoOk() (*BrandingThemeConfigurationLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *BrandingThemeConfiguration) SetLogo(v BrandingThemeConfigurationLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *BrandingThemeConfiguration) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLogoType

`func (o *BrandingThemeConfiguration) GetLogoType() EnumBrandingLogoType`

GetLogoType returns the LogoType field if non-nil, zero value otherwise.

### GetLogoTypeOk

`func (o *BrandingThemeConfiguration) GetLogoTypeOk() (*EnumBrandingLogoType, bool)`

GetLogoTypeOk returns a tuple with the LogoType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoType

`func (o *BrandingThemeConfiguration) SetLogoType(v EnumBrandingLogoType)`

SetLogoType sets LogoType field to given value.


### GetName

`func (o *BrandingThemeConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BrandingThemeConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BrandingThemeConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BrandingThemeConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


