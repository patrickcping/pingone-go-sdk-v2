# FormSocialLoginButtonAllOfStyles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | Pointer to [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | [optional] 
**TextColor** | Pointer to **string** | A string that specifies the button text color. The value must be a valid hexadecimal color. | [optional] 
**Enabled** | Pointer to **bool** | A boolean that specifies whether the button is enabled. | [optional] 
**Height** | Pointer to **int32** | An integer that specifies the button height. | [optional] 
**Width** | Pointer to **int32** | An integer that specifies the button width. | [optional] 
**WidthUnit** | Pointer to [**EnumFormStylesWidthUnit**](EnumFormStylesWidthUnit.md) |  | [optional] 
**Padding** | Pointer to [**FormStylesPadding**](FormStylesPadding.md) |  | [optional] 
**BackgroundColor** | Pointer to **string** | A string that specifies the button background color. The value must be a valid hexadecimal color. | [optional] 
**BorderColor** | Pointer to **string** | A string that specifies the button border color. The value must be a valid hexadecimal color. | [optional] 
**DisplayDefaultThemeButtonBackgroundColor** | Pointer to **bool** | A boolean that specifies whether the button uses the default theme’s background color. | [optional] 
**DisplayDefaultThemeButtonBorderColor** | Pointer to **bool** | A boolean that specifies whether the button uses the default theme’s border color. | [optional] 
**DisplayDefaultThemeButtonTextColor** | Pointer to **bool** | A boolean that specifies whether the button uses the default theme’s text color. | [optional] 

## Methods

### NewFormSocialLoginButtonAllOfStyles

`func NewFormSocialLoginButtonAllOfStyles() *FormSocialLoginButtonAllOfStyles`

NewFormSocialLoginButtonAllOfStyles instantiates a new FormSocialLoginButtonAllOfStyles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSocialLoginButtonAllOfStylesWithDefaults

`func NewFormSocialLoginButtonAllOfStylesWithDefaults() *FormSocialLoginButtonAllOfStyles`

NewFormSocialLoginButtonAllOfStylesWithDefaults instantiates a new FormSocialLoginButtonAllOfStyles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignment

`func (o *FormSocialLoginButtonAllOfStyles) GetAlignment() EnumFormItemAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *FormSocialLoginButtonAllOfStyles) GetAlignmentOk() (*EnumFormItemAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *FormSocialLoginButtonAllOfStyles) SetAlignment(v EnumFormItemAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *FormSocialLoginButtonAllOfStyles) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetTextColor

`func (o *FormSocialLoginButtonAllOfStyles) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *FormSocialLoginButtonAllOfStyles) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *FormSocialLoginButtonAllOfStyles) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *FormSocialLoginButtonAllOfStyles) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.

### GetEnabled

`func (o *FormSocialLoginButtonAllOfStyles) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FormSocialLoginButtonAllOfStyles) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FormSocialLoginButtonAllOfStyles) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FormSocialLoginButtonAllOfStyles) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHeight

`func (o *FormSocialLoginButtonAllOfStyles) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *FormSocialLoginButtonAllOfStyles) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *FormSocialLoginButtonAllOfStyles) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *FormSocialLoginButtonAllOfStyles) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetWidth

`func (o *FormSocialLoginButtonAllOfStyles) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FormSocialLoginButtonAllOfStyles) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FormSocialLoginButtonAllOfStyles) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *FormSocialLoginButtonAllOfStyles) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetWidthUnit

`func (o *FormSocialLoginButtonAllOfStyles) GetWidthUnit() EnumFormStylesWidthUnit`

GetWidthUnit returns the WidthUnit field if non-nil, zero value otherwise.

### GetWidthUnitOk

`func (o *FormSocialLoginButtonAllOfStyles) GetWidthUnitOk() (*EnumFormStylesWidthUnit, bool)`

GetWidthUnitOk returns a tuple with the WidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthUnit

`func (o *FormSocialLoginButtonAllOfStyles) SetWidthUnit(v EnumFormStylesWidthUnit)`

SetWidthUnit sets WidthUnit field to given value.

### HasWidthUnit

`func (o *FormSocialLoginButtonAllOfStyles) HasWidthUnit() bool`

HasWidthUnit returns a boolean if a field has been set.

### GetPadding

`func (o *FormSocialLoginButtonAllOfStyles) GetPadding() FormStylesPadding`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *FormSocialLoginButtonAllOfStyles) GetPaddingOk() (*FormStylesPadding, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *FormSocialLoginButtonAllOfStyles) SetPadding(v FormStylesPadding)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *FormSocialLoginButtonAllOfStyles) HasPadding() bool`

HasPadding returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *FormSocialLoginButtonAllOfStyles) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *FormSocialLoginButtonAllOfStyles) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *FormSocialLoginButtonAllOfStyles) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *FormSocialLoginButtonAllOfStyles) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetBorderColor

`func (o *FormSocialLoginButtonAllOfStyles) GetBorderColor() string`

GetBorderColor returns the BorderColor field if non-nil, zero value otherwise.

### GetBorderColorOk

`func (o *FormSocialLoginButtonAllOfStyles) GetBorderColorOk() (*string, bool)`

GetBorderColorOk returns a tuple with the BorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColor

`func (o *FormSocialLoginButtonAllOfStyles) SetBorderColor(v string)`

SetBorderColor sets BorderColor field to given value.

### HasBorderColor

`func (o *FormSocialLoginButtonAllOfStyles) HasBorderColor() bool`

HasBorderColor returns a boolean if a field has been set.

### GetDisplayDefaultThemeButtonBackgroundColor

`func (o *FormSocialLoginButtonAllOfStyles) GetDisplayDefaultThemeButtonBackgroundColor() bool`

GetDisplayDefaultThemeButtonBackgroundColor returns the DisplayDefaultThemeButtonBackgroundColor field if non-nil, zero value otherwise.

### GetDisplayDefaultThemeButtonBackgroundColorOk

`func (o *FormSocialLoginButtonAllOfStyles) GetDisplayDefaultThemeButtonBackgroundColorOk() (*bool, bool)`

GetDisplayDefaultThemeButtonBackgroundColorOk returns a tuple with the DisplayDefaultThemeButtonBackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDefaultThemeButtonBackgroundColor

`func (o *FormSocialLoginButtonAllOfStyles) SetDisplayDefaultThemeButtonBackgroundColor(v bool)`

SetDisplayDefaultThemeButtonBackgroundColor sets DisplayDefaultThemeButtonBackgroundColor field to given value.

### HasDisplayDefaultThemeButtonBackgroundColor

`func (o *FormSocialLoginButtonAllOfStyles) HasDisplayDefaultThemeButtonBackgroundColor() bool`

HasDisplayDefaultThemeButtonBackgroundColor returns a boolean if a field has been set.

### GetDisplayDefaultThemeButtonBorderColor

`func (o *FormSocialLoginButtonAllOfStyles) GetDisplayDefaultThemeButtonBorderColor() bool`

GetDisplayDefaultThemeButtonBorderColor returns the DisplayDefaultThemeButtonBorderColor field if non-nil, zero value otherwise.

### GetDisplayDefaultThemeButtonBorderColorOk

`func (o *FormSocialLoginButtonAllOfStyles) GetDisplayDefaultThemeButtonBorderColorOk() (*bool, bool)`

GetDisplayDefaultThemeButtonBorderColorOk returns a tuple with the DisplayDefaultThemeButtonBorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDefaultThemeButtonBorderColor

`func (o *FormSocialLoginButtonAllOfStyles) SetDisplayDefaultThemeButtonBorderColor(v bool)`

SetDisplayDefaultThemeButtonBorderColor sets DisplayDefaultThemeButtonBorderColor field to given value.

### HasDisplayDefaultThemeButtonBorderColor

`func (o *FormSocialLoginButtonAllOfStyles) HasDisplayDefaultThemeButtonBorderColor() bool`

HasDisplayDefaultThemeButtonBorderColor returns a boolean if a field has been set.

### GetDisplayDefaultThemeButtonTextColor

`func (o *FormSocialLoginButtonAllOfStyles) GetDisplayDefaultThemeButtonTextColor() bool`

GetDisplayDefaultThemeButtonTextColor returns the DisplayDefaultThemeButtonTextColor field if non-nil, zero value otherwise.

### GetDisplayDefaultThemeButtonTextColorOk

`func (o *FormSocialLoginButtonAllOfStyles) GetDisplayDefaultThemeButtonTextColorOk() (*bool, bool)`

GetDisplayDefaultThemeButtonTextColorOk returns a tuple with the DisplayDefaultThemeButtonTextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayDefaultThemeButtonTextColor

`func (o *FormSocialLoginButtonAllOfStyles) SetDisplayDefaultThemeButtonTextColor(v bool)`

SetDisplayDefaultThemeButtonTextColor sets DisplayDefaultThemeButtonTextColor field to given value.

### HasDisplayDefaultThemeButtonTextColor

`func (o *FormSocialLoginButtonAllOfStyles) HasDisplayDefaultThemeButtonTextColor() bool`

HasDisplayDefaultThemeButtonTextColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


