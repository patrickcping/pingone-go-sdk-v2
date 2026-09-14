# BrandingThemeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackgroundColor** | Pointer to **string** | The background color for the theme. It must be a valid hexadecimal color code, and it is a required property when configuration.backgroundType is set to COLOR. | [optional] 
**BackgroundType** | [**EnumBrandingThemeBackgroundType**](EnumBrandingThemeBackgroundType.md) |  | 
**BackgroundImage** | Pointer to [**BrandingThemeConfigurationBackgroundImage**](BrandingThemeConfigurationBackgroundImage.md) |  | [optional] 
**ApplicationBackgroundColor** | Pointer to **string** | The application background color for the theme. It must be a valid hexadecimal color code. Note that this property is not used by DaVinci forms. | [optional] 
**BodyTextColor** | **string** | The body text color for the theme. It must be a valid hexadecimal color code. | 
**BodyTextSize** | Pointer to **string** | The body text size for the theme. | [optional] 
**BodyTextWeight** | Pointer to **string** | The body text weight for the theme (range from 100-900). | [optional] 
**ButtonColor** | **string** | The button color for the theme. It must be a valid hexadecimal color code. | 
**ButtonBorderColor** | Pointer to **string** | The button border color for the theme. It must be a valid hexadecimal color code. | [optional] 
**ButtonCornerRadius** | Pointer to **string** | The button corner radius for the theme (range from 0-25px). | [optional] 
**ButtonBorderWidth** | Pointer to **string** | The button border width for the theme (range from 0-4px). | [optional] 
**ButtonHoverStateBorderColor** | Pointer to **string** | The button hover state border color for the theme. It must be a valid hexadecimal color code. | [optional] 
**ButtonHoverStateFillColor** | Pointer to **string** | The button hover state fill color for the theme. It must be a valid hexadecimal color code. | [optional] 
**ButtonHoverStateTextColor** | Pointer to **string** | The button hover state text color for the theme. It must be a valid hexadecimal color code. | [optional] 
**ButtonTextColor** | **string** | The button text color for the branding theme. It must be a valid hexadecimal color code. | 
**ButtonTextSize** | Pointer to **string** | The button text size for the theme. | [optional] 
**ButtonTextWeight** | Pointer to **string** | The button text weight for the theme (range from 100-900). | [optional] 
**CardColor** | **string** | The card color for the branding theme. It must be a valid hexadecimal color code. | 
**CardBorderColor** | Pointer to **string** | The card border color for the theme. It must be a valid hexadecimal color code. | [optional] 
**CardBorderWidth** | Pointer to **string** | The card border width for the theme (range from 0-4px). | [optional] 
**CardCornerRadius** | Pointer to **string** | The card corner radius for the theme (range from 0-25px). | [optional] 
**CardHorizontalAlignment** | Pointer to **string** | The card horizontal alignment for the theme. | [optional] 
**CardLogoAlignment** | Pointer to **string** | The card logo alignment for the theme. | [optional] 
**CardShadow** | Pointer to **string** | The card shadow for the theme. | [optional] 
**CardVerticalAlignment** | Pointer to **string** | The card vertical alignment for the theme. | [optional] 
**FocusRectangleColor** | Pointer to **string** | The focus rectangle color for the theme. It must be a valid hexadecimal color code. | [optional] 
**Footer** | Pointer to **string** | The footer of the branding theme. | [optional] 
**FooterLocalized** | Pointer to [**BrandingThemeConfigurationLocalizedText**](BrandingThemeConfigurationLocalizedText.md) |  | [optional] 
**ForegroundHighlightColor** | Pointer to **string** | For PingOne Neo verification presentation screen, the highlight color of the foreground object for the branding theme. It must be a valid hexadecimal color code. Defaults to | [optional] 
**ForegroundMainColor** | Pointer to **string** | For PingOne Neo verification presentation screen, the outline color of the foreground object for the branding theme. It must be a valid hexadecimal color code. Defaults to | [optional] 
**GlobalFont** | Pointer to **string** | The global font for the theme. The default value is &#x60;Helvetica Neue, Helvetica, sans-serif&#x60;. | [optional] 
**Header** | Pointer to **string** | The header for the theme. For example, \&quot;&lt;h1&gt;Welcome to PingOne&lt;h1&gt;\&quot;. | [optional] 
**HeaderBackgroundColor** | Pointer to **string** | The header background color for the theme. It must be a valid hexadecimal color code. Note that this property is not used by DaVinci forms. | [optional] 
**HeaderLocalized** | Pointer to [**BrandingThemeConfigurationLocalizedText**](BrandingThemeConfigurationLocalizedText.md) |  | [optional] 
**HeadingTextColor** | **string** | The heading text color for the branding theme. It must be a valid hexadecimal color code. | 
**InputBorderWidth** | Pointer to **string** | The input border width for the theme (range from 0-4px). | [optional] 
**InputBoxBorderColor** | Pointer to **string** | The input box border color for the theme. It must be a valid hexadecimal color code. | [optional] 
**InputCornerRadius** | Pointer to **string** | The input corner radius for the theme (range from 0-25px). | [optional] 
**InputLabelPosition** | Pointer to **string** | The input label position for the theme. | [optional] 
**InputLabelTextColor** | Pointer to **string** | The input label text color for the theme. It must be a valid hexadecimal color code. | [optional] 
**InputLabelTextSize** | Pointer to **string** | The input label text size for the theme. | [optional] 
**InputLabelTextWeight** | Pointer to **string** | The input label text weight for the theme (range from 100-900). | [optional] 
**InputValueTextColor** | Pointer to **string** | The input value text color for the theme. It must be a valid hexadecimal color code. | [optional] 
**InputValueTextSize** | Pointer to **string** | The input value text size for the theme. | [optional] 
**InputValueTextWeight** | Pointer to **string** | The input value text weight for the theme (range from 100-900). | [optional] 
**LinkTextColor** | **string** | The hyperlink text color for the branding theme. It must be a valid hexadecimal color code. | 
**LinkTextHoverColor** | Pointer to **string** | The link text hover color for the theme. It must be a valid hexadecimal color code. | [optional] 
**LinkTextSize** | Pointer to **string** | The link text size for the theme. | [optional] 
**LinkTextWeight** | Pointer to **string** | The link text weight for the theme (range from 100-900). | [optional] 
**Logo** | Pointer to [**BrandingThemeConfigurationLogo**](BrandingThemeConfigurationLogo.md) |  | [optional] 
**LogoHeight** | Pointer to **string** | The logo height assigned to the image (range from 0-100px). | [optional] 
**LogoType** | [**EnumBrandingLogoType**](EnumBrandingLogoType.md) |  | 
**Name** | Pointer to **string** | The name of the branding theme. | [optional] 
**SchemaVersion** | Pointer to **int32** | The schema version number for the theme. | [optional] [readonly] 
**SubTitleTextColor** | Pointer to **string** | The subtitle text color for the theme. It must be a valid hexadecimal color code. | [optional] 
**SubTitleTextSize** | Pointer to **string** | The subtitle text size for the theme. | [optional] 
**SubTitleTextWeight** | Pointer to **string** | The subtitle text weight for the theme (range from 100-900). | [optional] 
**TitleTextColor** | Pointer to **string** | The title text color for the theme. It must be a valid hexadecimal color code. | [optional] 
**TitleTextSize** | Pointer to **string** | The title text size for the theme. | [optional] 
**TitleTextWeight** | Pointer to **string** | The title text weight for the theme (range from 100-900). | [optional] 

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

### GetApplicationBackgroundColor

`func (o *BrandingThemeConfiguration) GetApplicationBackgroundColor() string`

GetApplicationBackgroundColor returns the ApplicationBackgroundColor field if non-nil, zero value otherwise.

### GetApplicationBackgroundColorOk

`func (o *BrandingThemeConfiguration) GetApplicationBackgroundColorOk() (*string, bool)`

GetApplicationBackgroundColorOk returns a tuple with the ApplicationBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationBackgroundColor

`func (o *BrandingThemeConfiguration) SetApplicationBackgroundColor(v string)`

SetApplicationBackgroundColor sets ApplicationBackgroundColor field to given value.

### HasApplicationBackgroundColor

`func (o *BrandingThemeConfiguration) HasApplicationBackgroundColor() bool`

HasApplicationBackgroundColor returns a boolean if a field has been set.

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


### GetBodyTextSize

`func (o *BrandingThemeConfiguration) GetBodyTextSize() string`

GetBodyTextSize returns the BodyTextSize field if non-nil, zero value otherwise.

### GetBodyTextSizeOk

`func (o *BrandingThemeConfiguration) GetBodyTextSizeOk() (*string, bool)`

GetBodyTextSizeOk returns a tuple with the BodyTextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTextSize

`func (o *BrandingThemeConfiguration) SetBodyTextSize(v string)`

SetBodyTextSize sets BodyTextSize field to given value.

### HasBodyTextSize

`func (o *BrandingThemeConfiguration) HasBodyTextSize() bool`

HasBodyTextSize returns a boolean if a field has been set.

### GetBodyTextWeight

`func (o *BrandingThemeConfiguration) GetBodyTextWeight() string`

GetBodyTextWeight returns the BodyTextWeight field if non-nil, zero value otherwise.

### GetBodyTextWeightOk

`func (o *BrandingThemeConfiguration) GetBodyTextWeightOk() (*string, bool)`

GetBodyTextWeightOk returns a tuple with the BodyTextWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBodyTextWeight

`func (o *BrandingThemeConfiguration) SetBodyTextWeight(v string)`

SetBodyTextWeight sets BodyTextWeight field to given value.

### HasBodyTextWeight

`func (o *BrandingThemeConfiguration) HasBodyTextWeight() bool`

HasBodyTextWeight returns a boolean if a field has been set.

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


### GetButtonBorderColor

`func (o *BrandingThemeConfiguration) GetButtonBorderColor() string`

GetButtonBorderColor returns the ButtonBorderColor field if non-nil, zero value otherwise.

### GetButtonBorderColorOk

`func (o *BrandingThemeConfiguration) GetButtonBorderColorOk() (*string, bool)`

GetButtonBorderColorOk returns a tuple with the ButtonBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonBorderColor

`func (o *BrandingThemeConfiguration) SetButtonBorderColor(v string)`

SetButtonBorderColor sets ButtonBorderColor field to given value.

### HasButtonBorderColor

`func (o *BrandingThemeConfiguration) HasButtonBorderColor() bool`

HasButtonBorderColor returns a boolean if a field has been set.

### GetButtonCornerRadius

`func (o *BrandingThemeConfiguration) GetButtonCornerRadius() string`

GetButtonCornerRadius returns the ButtonCornerRadius field if non-nil, zero value otherwise.

### GetButtonCornerRadiusOk

`func (o *BrandingThemeConfiguration) GetButtonCornerRadiusOk() (*string, bool)`

GetButtonCornerRadiusOk returns a tuple with the ButtonCornerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonCornerRadius

`func (o *BrandingThemeConfiguration) SetButtonCornerRadius(v string)`

SetButtonCornerRadius sets ButtonCornerRadius field to given value.

### HasButtonCornerRadius

`func (o *BrandingThemeConfiguration) HasButtonCornerRadius() bool`

HasButtonCornerRadius returns a boolean if a field has been set.

### GetButtonBorderWidth

`func (o *BrandingThemeConfiguration) GetButtonBorderWidth() string`

GetButtonBorderWidth returns the ButtonBorderWidth field if non-nil, zero value otherwise.

### GetButtonBorderWidthOk

`func (o *BrandingThemeConfiguration) GetButtonBorderWidthOk() (*string, bool)`

GetButtonBorderWidthOk returns a tuple with the ButtonBorderWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonBorderWidth

`func (o *BrandingThemeConfiguration) SetButtonBorderWidth(v string)`

SetButtonBorderWidth sets ButtonBorderWidth field to given value.

### HasButtonBorderWidth

`func (o *BrandingThemeConfiguration) HasButtonBorderWidth() bool`

HasButtonBorderWidth returns a boolean if a field has been set.

### GetButtonHoverStateBorderColor

`func (o *BrandingThemeConfiguration) GetButtonHoverStateBorderColor() string`

GetButtonHoverStateBorderColor returns the ButtonHoverStateBorderColor field if non-nil, zero value otherwise.

### GetButtonHoverStateBorderColorOk

`func (o *BrandingThemeConfiguration) GetButtonHoverStateBorderColorOk() (*string, bool)`

GetButtonHoverStateBorderColorOk returns a tuple with the ButtonHoverStateBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonHoverStateBorderColor

`func (o *BrandingThemeConfiguration) SetButtonHoverStateBorderColor(v string)`

SetButtonHoverStateBorderColor sets ButtonHoverStateBorderColor field to given value.

### HasButtonHoverStateBorderColor

`func (o *BrandingThemeConfiguration) HasButtonHoverStateBorderColor() bool`

HasButtonHoverStateBorderColor returns a boolean if a field has been set.

### GetButtonHoverStateFillColor

`func (o *BrandingThemeConfiguration) GetButtonHoverStateFillColor() string`

GetButtonHoverStateFillColor returns the ButtonHoverStateFillColor field if non-nil, zero value otherwise.

### GetButtonHoverStateFillColorOk

`func (o *BrandingThemeConfiguration) GetButtonHoverStateFillColorOk() (*string, bool)`

GetButtonHoverStateFillColorOk returns a tuple with the ButtonHoverStateFillColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonHoverStateFillColor

`func (o *BrandingThemeConfiguration) SetButtonHoverStateFillColor(v string)`

SetButtonHoverStateFillColor sets ButtonHoverStateFillColor field to given value.

### HasButtonHoverStateFillColor

`func (o *BrandingThemeConfiguration) HasButtonHoverStateFillColor() bool`

HasButtonHoverStateFillColor returns a boolean if a field has been set.

### GetButtonHoverStateTextColor

`func (o *BrandingThemeConfiguration) GetButtonHoverStateTextColor() string`

GetButtonHoverStateTextColor returns the ButtonHoverStateTextColor field if non-nil, zero value otherwise.

### GetButtonHoverStateTextColorOk

`func (o *BrandingThemeConfiguration) GetButtonHoverStateTextColorOk() (*string, bool)`

GetButtonHoverStateTextColorOk returns a tuple with the ButtonHoverStateTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonHoverStateTextColor

`func (o *BrandingThemeConfiguration) SetButtonHoverStateTextColor(v string)`

SetButtonHoverStateTextColor sets ButtonHoverStateTextColor field to given value.

### HasButtonHoverStateTextColor

`func (o *BrandingThemeConfiguration) HasButtonHoverStateTextColor() bool`

HasButtonHoverStateTextColor returns a boolean if a field has been set.

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


### GetButtonTextSize

`func (o *BrandingThemeConfiguration) GetButtonTextSize() string`

GetButtonTextSize returns the ButtonTextSize field if non-nil, zero value otherwise.

### GetButtonTextSizeOk

`func (o *BrandingThemeConfiguration) GetButtonTextSizeOk() (*string, bool)`

GetButtonTextSizeOk returns a tuple with the ButtonTextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTextSize

`func (o *BrandingThemeConfiguration) SetButtonTextSize(v string)`

SetButtonTextSize sets ButtonTextSize field to given value.

### HasButtonTextSize

`func (o *BrandingThemeConfiguration) HasButtonTextSize() bool`

HasButtonTextSize returns a boolean if a field has been set.

### GetButtonTextWeight

`func (o *BrandingThemeConfiguration) GetButtonTextWeight() string`

GetButtonTextWeight returns the ButtonTextWeight field if non-nil, zero value otherwise.

### GetButtonTextWeightOk

`func (o *BrandingThemeConfiguration) GetButtonTextWeightOk() (*string, bool)`

GetButtonTextWeightOk returns a tuple with the ButtonTextWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonTextWeight

`func (o *BrandingThemeConfiguration) SetButtonTextWeight(v string)`

SetButtonTextWeight sets ButtonTextWeight field to given value.

### HasButtonTextWeight

`func (o *BrandingThemeConfiguration) HasButtonTextWeight() bool`

HasButtonTextWeight returns a boolean if a field has been set.

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


### GetCardBorderColor

`func (o *BrandingThemeConfiguration) GetCardBorderColor() string`

GetCardBorderColor returns the CardBorderColor field if non-nil, zero value otherwise.

### GetCardBorderColorOk

`func (o *BrandingThemeConfiguration) GetCardBorderColorOk() (*string, bool)`

GetCardBorderColorOk returns a tuple with the CardBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBorderColor

`func (o *BrandingThemeConfiguration) SetCardBorderColor(v string)`

SetCardBorderColor sets CardBorderColor field to given value.

### HasCardBorderColor

`func (o *BrandingThemeConfiguration) HasCardBorderColor() bool`

HasCardBorderColor returns a boolean if a field has been set.

### GetCardBorderWidth

`func (o *BrandingThemeConfiguration) GetCardBorderWidth() string`

GetCardBorderWidth returns the CardBorderWidth field if non-nil, zero value otherwise.

### GetCardBorderWidthOk

`func (o *BrandingThemeConfiguration) GetCardBorderWidthOk() (*string, bool)`

GetCardBorderWidthOk returns a tuple with the CardBorderWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBorderWidth

`func (o *BrandingThemeConfiguration) SetCardBorderWidth(v string)`

SetCardBorderWidth sets CardBorderWidth field to given value.

### HasCardBorderWidth

`func (o *BrandingThemeConfiguration) HasCardBorderWidth() bool`

HasCardBorderWidth returns a boolean if a field has been set.

### GetCardCornerRadius

`func (o *BrandingThemeConfiguration) GetCardCornerRadius() string`

GetCardCornerRadius returns the CardCornerRadius field if non-nil, zero value otherwise.

### GetCardCornerRadiusOk

`func (o *BrandingThemeConfiguration) GetCardCornerRadiusOk() (*string, bool)`

GetCardCornerRadiusOk returns a tuple with the CardCornerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardCornerRadius

`func (o *BrandingThemeConfiguration) SetCardCornerRadius(v string)`

SetCardCornerRadius sets CardCornerRadius field to given value.

### HasCardCornerRadius

`func (o *BrandingThemeConfiguration) HasCardCornerRadius() bool`

HasCardCornerRadius returns a boolean if a field has been set.

### GetCardHorizontalAlignment

`func (o *BrandingThemeConfiguration) GetCardHorizontalAlignment() string`

GetCardHorizontalAlignment returns the CardHorizontalAlignment field if non-nil, zero value otherwise.

### GetCardHorizontalAlignmentOk

`func (o *BrandingThemeConfiguration) GetCardHorizontalAlignmentOk() (*string, bool)`

GetCardHorizontalAlignmentOk returns a tuple with the CardHorizontalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardHorizontalAlignment

`func (o *BrandingThemeConfiguration) SetCardHorizontalAlignment(v string)`

SetCardHorizontalAlignment sets CardHorizontalAlignment field to given value.

### HasCardHorizontalAlignment

`func (o *BrandingThemeConfiguration) HasCardHorizontalAlignment() bool`

HasCardHorizontalAlignment returns a boolean if a field has been set.

### GetCardLogoAlignment

`func (o *BrandingThemeConfiguration) GetCardLogoAlignment() string`

GetCardLogoAlignment returns the CardLogoAlignment field if non-nil, zero value otherwise.

### GetCardLogoAlignmentOk

`func (o *BrandingThemeConfiguration) GetCardLogoAlignmentOk() (*string, bool)`

GetCardLogoAlignmentOk returns a tuple with the CardLogoAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardLogoAlignment

`func (o *BrandingThemeConfiguration) SetCardLogoAlignment(v string)`

SetCardLogoAlignment sets CardLogoAlignment field to given value.

### HasCardLogoAlignment

`func (o *BrandingThemeConfiguration) HasCardLogoAlignment() bool`

HasCardLogoAlignment returns a boolean if a field has been set.

### GetCardShadow

`func (o *BrandingThemeConfiguration) GetCardShadow() string`

GetCardShadow returns the CardShadow field if non-nil, zero value otherwise.

### GetCardShadowOk

`func (o *BrandingThemeConfiguration) GetCardShadowOk() (*string, bool)`

GetCardShadowOk returns a tuple with the CardShadow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardShadow

`func (o *BrandingThemeConfiguration) SetCardShadow(v string)`

SetCardShadow sets CardShadow field to given value.

### HasCardShadow

`func (o *BrandingThemeConfiguration) HasCardShadow() bool`

HasCardShadow returns a boolean if a field has been set.

### GetCardVerticalAlignment

`func (o *BrandingThemeConfiguration) GetCardVerticalAlignment() string`

GetCardVerticalAlignment returns the CardVerticalAlignment field if non-nil, zero value otherwise.

### GetCardVerticalAlignmentOk

`func (o *BrandingThemeConfiguration) GetCardVerticalAlignmentOk() (*string, bool)`

GetCardVerticalAlignmentOk returns a tuple with the CardVerticalAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardVerticalAlignment

`func (o *BrandingThemeConfiguration) SetCardVerticalAlignment(v string)`

SetCardVerticalAlignment sets CardVerticalAlignment field to given value.

### HasCardVerticalAlignment

`func (o *BrandingThemeConfiguration) HasCardVerticalAlignment() bool`

HasCardVerticalAlignment returns a boolean if a field has been set.

### GetFocusRectangleColor

`func (o *BrandingThemeConfiguration) GetFocusRectangleColor() string`

GetFocusRectangleColor returns the FocusRectangleColor field if non-nil, zero value otherwise.

### GetFocusRectangleColorOk

`func (o *BrandingThemeConfiguration) GetFocusRectangleColorOk() (*string, bool)`

GetFocusRectangleColorOk returns a tuple with the FocusRectangleColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFocusRectangleColor

`func (o *BrandingThemeConfiguration) SetFocusRectangleColor(v string)`

SetFocusRectangleColor sets FocusRectangleColor field to given value.

### HasFocusRectangleColor

`func (o *BrandingThemeConfiguration) HasFocusRectangleColor() bool`

HasFocusRectangleColor returns a boolean if a field has been set.

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

### GetFooterLocalized

`func (o *BrandingThemeConfiguration) GetFooterLocalized() BrandingThemeConfigurationLocalizedText`

GetFooterLocalized returns the FooterLocalized field if non-nil, zero value otherwise.

### GetFooterLocalizedOk

`func (o *BrandingThemeConfiguration) GetFooterLocalizedOk() (*BrandingThemeConfigurationLocalizedText, bool)`

GetFooterLocalizedOk returns a tuple with the FooterLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterLocalized

`func (o *BrandingThemeConfiguration) SetFooterLocalized(v BrandingThemeConfigurationLocalizedText)`

SetFooterLocalized sets FooterLocalized field to given value.

### HasFooterLocalized

`func (o *BrandingThemeConfiguration) HasFooterLocalized() bool`

HasFooterLocalized returns a boolean if a field has been set.

### GetForegroundHighlightColor

`func (o *BrandingThemeConfiguration) GetForegroundHighlightColor() string`

GetForegroundHighlightColor returns the ForegroundHighlightColor field if non-nil, zero value otherwise.

### GetForegroundHighlightColorOk

`func (o *BrandingThemeConfiguration) GetForegroundHighlightColorOk() (*string, bool)`

GetForegroundHighlightColorOk returns a tuple with the ForegroundHighlightColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForegroundHighlightColor

`func (o *BrandingThemeConfiguration) SetForegroundHighlightColor(v string)`

SetForegroundHighlightColor sets ForegroundHighlightColor field to given value.

### HasForegroundHighlightColor

`func (o *BrandingThemeConfiguration) HasForegroundHighlightColor() bool`

HasForegroundHighlightColor returns a boolean if a field has been set.

### GetForegroundMainColor

`func (o *BrandingThemeConfiguration) GetForegroundMainColor() string`

GetForegroundMainColor returns the ForegroundMainColor field if non-nil, zero value otherwise.

### GetForegroundMainColorOk

`func (o *BrandingThemeConfiguration) GetForegroundMainColorOk() (*string, bool)`

GetForegroundMainColorOk returns a tuple with the ForegroundMainColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForegroundMainColor

`func (o *BrandingThemeConfiguration) SetForegroundMainColor(v string)`

SetForegroundMainColor sets ForegroundMainColor field to given value.

### HasForegroundMainColor

`func (o *BrandingThemeConfiguration) HasForegroundMainColor() bool`

HasForegroundMainColor returns a boolean if a field has been set.

### GetGlobalFont

`func (o *BrandingThemeConfiguration) GetGlobalFont() string`

GetGlobalFont returns the GlobalFont field if non-nil, zero value otherwise.

### GetGlobalFontOk

`func (o *BrandingThemeConfiguration) GetGlobalFontOk() (*string, bool)`

GetGlobalFontOk returns a tuple with the GlobalFont field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalFont

`func (o *BrandingThemeConfiguration) SetGlobalFont(v string)`

SetGlobalFont sets GlobalFont field to given value.

### HasGlobalFont

`func (o *BrandingThemeConfiguration) HasGlobalFont() bool`

HasGlobalFont returns a boolean if a field has been set.

### GetHeader

`func (o *BrandingThemeConfiguration) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *BrandingThemeConfiguration) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *BrandingThemeConfiguration) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *BrandingThemeConfiguration) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetHeaderBackgroundColor

`func (o *BrandingThemeConfiguration) GetHeaderBackgroundColor() string`

GetHeaderBackgroundColor returns the HeaderBackgroundColor field if non-nil, zero value otherwise.

### GetHeaderBackgroundColorOk

`func (o *BrandingThemeConfiguration) GetHeaderBackgroundColorOk() (*string, bool)`

GetHeaderBackgroundColorOk returns a tuple with the HeaderBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderBackgroundColor

`func (o *BrandingThemeConfiguration) SetHeaderBackgroundColor(v string)`

SetHeaderBackgroundColor sets HeaderBackgroundColor field to given value.

### HasHeaderBackgroundColor

`func (o *BrandingThemeConfiguration) HasHeaderBackgroundColor() bool`

HasHeaderBackgroundColor returns a boolean if a field has been set.

### GetHeaderLocalized

`func (o *BrandingThemeConfiguration) GetHeaderLocalized() BrandingThemeConfigurationLocalizedText`

GetHeaderLocalized returns the HeaderLocalized field if non-nil, zero value otherwise.

### GetHeaderLocalizedOk

`func (o *BrandingThemeConfiguration) GetHeaderLocalizedOk() (*BrandingThemeConfigurationLocalizedText, bool)`

GetHeaderLocalizedOk returns a tuple with the HeaderLocalized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderLocalized

`func (o *BrandingThemeConfiguration) SetHeaderLocalized(v BrandingThemeConfigurationLocalizedText)`

SetHeaderLocalized sets HeaderLocalized field to given value.

### HasHeaderLocalized

`func (o *BrandingThemeConfiguration) HasHeaderLocalized() bool`

HasHeaderLocalized returns a boolean if a field has been set.

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


### GetInputBorderWidth

`func (o *BrandingThemeConfiguration) GetInputBorderWidth() string`

GetInputBorderWidth returns the InputBorderWidth field if non-nil, zero value otherwise.

### GetInputBorderWidthOk

`func (o *BrandingThemeConfiguration) GetInputBorderWidthOk() (*string, bool)`

GetInputBorderWidthOk returns a tuple with the InputBorderWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBorderWidth

`func (o *BrandingThemeConfiguration) SetInputBorderWidth(v string)`

SetInputBorderWidth sets InputBorderWidth field to given value.

### HasInputBorderWidth

`func (o *BrandingThemeConfiguration) HasInputBorderWidth() bool`

HasInputBorderWidth returns a boolean if a field has been set.

### GetInputBoxBorderColor

`func (o *BrandingThemeConfiguration) GetInputBoxBorderColor() string`

GetInputBoxBorderColor returns the InputBoxBorderColor field if non-nil, zero value otherwise.

### GetInputBoxBorderColorOk

`func (o *BrandingThemeConfiguration) GetInputBoxBorderColorOk() (*string, bool)`

GetInputBoxBorderColorOk returns a tuple with the InputBoxBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputBoxBorderColor

`func (o *BrandingThemeConfiguration) SetInputBoxBorderColor(v string)`

SetInputBoxBorderColor sets InputBoxBorderColor field to given value.

### HasInputBoxBorderColor

`func (o *BrandingThemeConfiguration) HasInputBoxBorderColor() bool`

HasInputBoxBorderColor returns a boolean if a field has been set.

### GetInputCornerRadius

`func (o *BrandingThemeConfiguration) GetInputCornerRadius() string`

GetInputCornerRadius returns the InputCornerRadius field if non-nil, zero value otherwise.

### GetInputCornerRadiusOk

`func (o *BrandingThemeConfiguration) GetInputCornerRadiusOk() (*string, bool)`

GetInputCornerRadiusOk returns a tuple with the InputCornerRadius field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputCornerRadius

`func (o *BrandingThemeConfiguration) SetInputCornerRadius(v string)`

SetInputCornerRadius sets InputCornerRadius field to given value.

### HasInputCornerRadius

`func (o *BrandingThemeConfiguration) HasInputCornerRadius() bool`

HasInputCornerRadius returns a boolean if a field has been set.

### GetInputLabelPosition

`func (o *BrandingThemeConfiguration) GetInputLabelPosition() string`

GetInputLabelPosition returns the InputLabelPosition field if non-nil, zero value otherwise.

### GetInputLabelPositionOk

`func (o *BrandingThemeConfiguration) GetInputLabelPositionOk() (*string, bool)`

GetInputLabelPositionOk returns a tuple with the InputLabelPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputLabelPosition

`func (o *BrandingThemeConfiguration) SetInputLabelPosition(v string)`

SetInputLabelPosition sets InputLabelPosition field to given value.

### HasInputLabelPosition

`func (o *BrandingThemeConfiguration) HasInputLabelPosition() bool`

HasInputLabelPosition returns a boolean if a field has been set.

### GetInputLabelTextColor

`func (o *BrandingThemeConfiguration) GetInputLabelTextColor() string`

GetInputLabelTextColor returns the InputLabelTextColor field if non-nil, zero value otherwise.

### GetInputLabelTextColorOk

`func (o *BrandingThemeConfiguration) GetInputLabelTextColorOk() (*string, bool)`

GetInputLabelTextColorOk returns a tuple with the InputLabelTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputLabelTextColor

`func (o *BrandingThemeConfiguration) SetInputLabelTextColor(v string)`

SetInputLabelTextColor sets InputLabelTextColor field to given value.

### HasInputLabelTextColor

`func (o *BrandingThemeConfiguration) HasInputLabelTextColor() bool`

HasInputLabelTextColor returns a boolean if a field has been set.

### GetInputLabelTextSize

`func (o *BrandingThemeConfiguration) GetInputLabelTextSize() string`

GetInputLabelTextSize returns the InputLabelTextSize field if non-nil, zero value otherwise.

### GetInputLabelTextSizeOk

`func (o *BrandingThemeConfiguration) GetInputLabelTextSizeOk() (*string, bool)`

GetInputLabelTextSizeOk returns a tuple with the InputLabelTextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputLabelTextSize

`func (o *BrandingThemeConfiguration) SetInputLabelTextSize(v string)`

SetInputLabelTextSize sets InputLabelTextSize field to given value.

### HasInputLabelTextSize

`func (o *BrandingThemeConfiguration) HasInputLabelTextSize() bool`

HasInputLabelTextSize returns a boolean if a field has been set.

### GetInputLabelTextWeight

`func (o *BrandingThemeConfiguration) GetInputLabelTextWeight() string`

GetInputLabelTextWeight returns the InputLabelTextWeight field if non-nil, zero value otherwise.

### GetInputLabelTextWeightOk

`func (o *BrandingThemeConfiguration) GetInputLabelTextWeightOk() (*string, bool)`

GetInputLabelTextWeightOk returns a tuple with the InputLabelTextWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputLabelTextWeight

`func (o *BrandingThemeConfiguration) SetInputLabelTextWeight(v string)`

SetInputLabelTextWeight sets InputLabelTextWeight field to given value.

### HasInputLabelTextWeight

`func (o *BrandingThemeConfiguration) HasInputLabelTextWeight() bool`

HasInputLabelTextWeight returns a boolean if a field has been set.

### GetInputValueTextColor

`func (o *BrandingThemeConfiguration) GetInputValueTextColor() string`

GetInputValueTextColor returns the InputValueTextColor field if non-nil, zero value otherwise.

### GetInputValueTextColorOk

`func (o *BrandingThemeConfiguration) GetInputValueTextColorOk() (*string, bool)`

GetInputValueTextColorOk returns a tuple with the InputValueTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValueTextColor

`func (o *BrandingThemeConfiguration) SetInputValueTextColor(v string)`

SetInputValueTextColor sets InputValueTextColor field to given value.

### HasInputValueTextColor

`func (o *BrandingThemeConfiguration) HasInputValueTextColor() bool`

HasInputValueTextColor returns a boolean if a field has been set.

### GetInputValueTextSize

`func (o *BrandingThemeConfiguration) GetInputValueTextSize() string`

GetInputValueTextSize returns the InputValueTextSize field if non-nil, zero value otherwise.

### GetInputValueTextSizeOk

`func (o *BrandingThemeConfiguration) GetInputValueTextSizeOk() (*string, bool)`

GetInputValueTextSizeOk returns a tuple with the InputValueTextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValueTextSize

`func (o *BrandingThemeConfiguration) SetInputValueTextSize(v string)`

SetInputValueTextSize sets InputValueTextSize field to given value.

### HasInputValueTextSize

`func (o *BrandingThemeConfiguration) HasInputValueTextSize() bool`

HasInputValueTextSize returns a boolean if a field has been set.

### GetInputValueTextWeight

`func (o *BrandingThemeConfiguration) GetInputValueTextWeight() string`

GetInputValueTextWeight returns the InputValueTextWeight field if non-nil, zero value otherwise.

### GetInputValueTextWeightOk

`func (o *BrandingThemeConfiguration) GetInputValueTextWeightOk() (*string, bool)`

GetInputValueTextWeightOk returns a tuple with the InputValueTextWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputValueTextWeight

`func (o *BrandingThemeConfiguration) SetInputValueTextWeight(v string)`

SetInputValueTextWeight sets InputValueTextWeight field to given value.

### HasInputValueTextWeight

`func (o *BrandingThemeConfiguration) HasInputValueTextWeight() bool`

HasInputValueTextWeight returns a boolean if a field has been set.

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


### GetLinkTextHoverColor

`func (o *BrandingThemeConfiguration) GetLinkTextHoverColor() string`

GetLinkTextHoverColor returns the LinkTextHoverColor field if non-nil, zero value otherwise.

### GetLinkTextHoverColorOk

`func (o *BrandingThemeConfiguration) GetLinkTextHoverColorOk() (*string, bool)`

GetLinkTextHoverColorOk returns a tuple with the LinkTextHoverColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTextHoverColor

`func (o *BrandingThemeConfiguration) SetLinkTextHoverColor(v string)`

SetLinkTextHoverColor sets LinkTextHoverColor field to given value.

### HasLinkTextHoverColor

`func (o *BrandingThemeConfiguration) HasLinkTextHoverColor() bool`

HasLinkTextHoverColor returns a boolean if a field has been set.

### GetLinkTextSize

`func (o *BrandingThemeConfiguration) GetLinkTextSize() string`

GetLinkTextSize returns the LinkTextSize field if non-nil, zero value otherwise.

### GetLinkTextSizeOk

`func (o *BrandingThemeConfiguration) GetLinkTextSizeOk() (*string, bool)`

GetLinkTextSizeOk returns a tuple with the LinkTextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTextSize

`func (o *BrandingThemeConfiguration) SetLinkTextSize(v string)`

SetLinkTextSize sets LinkTextSize field to given value.

### HasLinkTextSize

`func (o *BrandingThemeConfiguration) HasLinkTextSize() bool`

HasLinkTextSize returns a boolean if a field has been set.

### GetLinkTextWeight

`func (o *BrandingThemeConfiguration) GetLinkTextWeight() string`

GetLinkTextWeight returns the LinkTextWeight field if non-nil, zero value otherwise.

### GetLinkTextWeightOk

`func (o *BrandingThemeConfiguration) GetLinkTextWeightOk() (*string, bool)`

GetLinkTextWeightOk returns a tuple with the LinkTextWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTextWeight

`func (o *BrandingThemeConfiguration) SetLinkTextWeight(v string)`

SetLinkTextWeight sets LinkTextWeight field to given value.

### HasLinkTextWeight

`func (o *BrandingThemeConfiguration) HasLinkTextWeight() bool`

HasLinkTextWeight returns a boolean if a field has been set.

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

### GetLogoHeight

`func (o *BrandingThemeConfiguration) GetLogoHeight() string`

GetLogoHeight returns the LogoHeight field if non-nil, zero value otherwise.

### GetLogoHeightOk

`func (o *BrandingThemeConfiguration) GetLogoHeightOk() (*string, bool)`

GetLogoHeightOk returns a tuple with the LogoHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoHeight

`func (o *BrandingThemeConfiguration) SetLogoHeight(v string)`

SetLogoHeight sets LogoHeight field to given value.

### HasLogoHeight

`func (o *BrandingThemeConfiguration) HasLogoHeight() bool`

HasLogoHeight returns a boolean if a field has been set.

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

### GetSchemaVersion

`func (o *BrandingThemeConfiguration) GetSchemaVersion() int32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *BrandingThemeConfiguration) GetSchemaVersionOk() (*int32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *BrandingThemeConfiguration) SetSchemaVersion(v int32)`

SetSchemaVersion sets SchemaVersion field to given value.

### HasSchemaVersion

`func (o *BrandingThemeConfiguration) HasSchemaVersion() bool`

HasSchemaVersion returns a boolean if a field has been set.

### GetSubTitleTextColor

`func (o *BrandingThemeConfiguration) GetSubTitleTextColor() string`

GetSubTitleTextColor returns the SubTitleTextColor field if non-nil, zero value otherwise.

### GetSubTitleTextColorOk

`func (o *BrandingThemeConfiguration) GetSubTitleTextColorOk() (*string, bool)`

GetSubTitleTextColorOk returns a tuple with the SubTitleTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitleTextColor

`func (o *BrandingThemeConfiguration) SetSubTitleTextColor(v string)`

SetSubTitleTextColor sets SubTitleTextColor field to given value.

### HasSubTitleTextColor

`func (o *BrandingThemeConfiguration) HasSubTitleTextColor() bool`

HasSubTitleTextColor returns a boolean if a field has been set.

### GetSubTitleTextSize

`func (o *BrandingThemeConfiguration) GetSubTitleTextSize() string`

GetSubTitleTextSize returns the SubTitleTextSize field if non-nil, zero value otherwise.

### GetSubTitleTextSizeOk

`func (o *BrandingThemeConfiguration) GetSubTitleTextSizeOk() (*string, bool)`

GetSubTitleTextSizeOk returns a tuple with the SubTitleTextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitleTextSize

`func (o *BrandingThemeConfiguration) SetSubTitleTextSize(v string)`

SetSubTitleTextSize sets SubTitleTextSize field to given value.

### HasSubTitleTextSize

`func (o *BrandingThemeConfiguration) HasSubTitleTextSize() bool`

HasSubTitleTextSize returns a boolean if a field has been set.

### GetSubTitleTextWeight

`func (o *BrandingThemeConfiguration) GetSubTitleTextWeight() string`

GetSubTitleTextWeight returns the SubTitleTextWeight field if non-nil, zero value otherwise.

### GetSubTitleTextWeightOk

`func (o *BrandingThemeConfiguration) GetSubTitleTextWeightOk() (*string, bool)`

GetSubTitleTextWeightOk returns a tuple with the SubTitleTextWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTitleTextWeight

`func (o *BrandingThemeConfiguration) SetSubTitleTextWeight(v string)`

SetSubTitleTextWeight sets SubTitleTextWeight field to given value.

### HasSubTitleTextWeight

`func (o *BrandingThemeConfiguration) HasSubTitleTextWeight() bool`

HasSubTitleTextWeight returns a boolean if a field has been set.

### GetTitleTextColor

`func (o *BrandingThemeConfiguration) GetTitleTextColor() string`

GetTitleTextColor returns the TitleTextColor field if non-nil, zero value otherwise.

### GetTitleTextColorOk

`func (o *BrandingThemeConfiguration) GetTitleTextColorOk() (*string, bool)`

GetTitleTextColorOk returns a tuple with the TitleTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleTextColor

`func (o *BrandingThemeConfiguration) SetTitleTextColor(v string)`

SetTitleTextColor sets TitleTextColor field to given value.

### HasTitleTextColor

`func (o *BrandingThemeConfiguration) HasTitleTextColor() bool`

HasTitleTextColor returns a boolean if a field has been set.

### GetTitleTextSize

`func (o *BrandingThemeConfiguration) GetTitleTextSize() string`

GetTitleTextSize returns the TitleTextSize field if non-nil, zero value otherwise.

### GetTitleTextSizeOk

`func (o *BrandingThemeConfiguration) GetTitleTextSizeOk() (*string, bool)`

GetTitleTextSizeOk returns a tuple with the TitleTextSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleTextSize

`func (o *BrandingThemeConfiguration) SetTitleTextSize(v string)`

SetTitleTextSize sets TitleTextSize field to given value.

### HasTitleTextSize

`func (o *BrandingThemeConfiguration) HasTitleTextSize() bool`

HasTitleTextSize returns a boolean if a field has been set.

### GetTitleTextWeight

`func (o *BrandingThemeConfiguration) GetTitleTextWeight() string`

GetTitleTextWeight returns the TitleTextWeight field if non-nil, zero value otherwise.

### GetTitleTextWeightOk

`func (o *BrandingThemeConfiguration) GetTitleTextWeightOk() (*string, bool)`

GetTitleTextWeightOk returns a tuple with the TitleTextWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleTextWeight

`func (o *BrandingThemeConfiguration) SetTitleTextWeight(v string)`

SetTitleTextWeight sets TitleTextWeight field to given value.

### HasTitleTextWeight

`func (o *BrandingThemeConfiguration) HasTitleTextWeight() bool`

HasTitleTextWeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


