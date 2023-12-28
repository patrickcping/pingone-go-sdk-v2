# FormStyles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | Pointer to [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | [optional] 
**BackgroundColor** | Pointer to **string** | A string that specifies the button background color. The value must be a valid hexadecimal color. | [optional] 
**TextColor** | Pointer to **string** | A string that specifies the button text color. The value must be a valid hexadecimal color. | [optional] 
**BorderColor** | Pointer to **string** | A string that specifies the button border color. The value must be a valid hexadecimal color. | [optional] 
**Enabled** | Pointer to **bool** | A boolean that specifies whether the button is enabled. | [optional] 
**Height** | Pointer to **int32** |  | [optional] 
**Padding** | Pointer to [**FormStylesPadding**](FormStylesPadding.md) |  | [optional] 
**Width** | Pointer to **int32** | An integer that specifies the button width. Set as a percentage. | [optional] 
**WidthUnit** | Pointer to [**EnumFormStylesWidthUnit**](EnumFormStylesWidthUnit.md) |  | [optional] 

## Methods

### NewFormStyles

`func NewFormStyles() *FormStyles`

NewFormStyles instantiates a new FormStyles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormStylesWithDefaults

`func NewFormStylesWithDefaults() *FormStyles`

NewFormStylesWithDefaults instantiates a new FormStyles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignment

`func (o *FormStyles) GetAlignment() EnumFormItemAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *FormStyles) GetAlignmentOk() (*EnumFormItemAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *FormStyles) SetAlignment(v EnumFormItemAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *FormStyles) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetBackgroundColor

`func (o *FormStyles) GetBackgroundColor() string`

GetBackgroundColor returns the BackgroundColor field if non-nil, zero value otherwise.

### GetBackgroundColorOk

`func (o *FormStyles) GetBackgroundColorOk() (*string, bool)`

GetBackgroundColorOk returns a tuple with the BackgroundColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackgroundColor

`func (o *FormStyles) SetBackgroundColor(v string)`

SetBackgroundColor sets BackgroundColor field to given value.

### HasBackgroundColor

`func (o *FormStyles) HasBackgroundColor() bool`

HasBackgroundColor returns a boolean if a field has been set.

### GetTextColor

`func (o *FormStyles) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *FormStyles) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *FormStyles) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *FormStyles) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.

### GetBorderColor

`func (o *FormStyles) GetBorderColor() string`

GetBorderColor returns the BorderColor field if non-nil, zero value otherwise.

### GetBorderColorOk

`func (o *FormStyles) GetBorderColorOk() (*string, bool)`

GetBorderColorOk returns a tuple with the BorderColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBorderColor

`func (o *FormStyles) SetBorderColor(v string)`

SetBorderColor sets BorderColor field to given value.

### HasBorderColor

`func (o *FormStyles) HasBorderColor() bool`

HasBorderColor returns a boolean if a field has been set.

### GetEnabled

`func (o *FormStyles) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FormStyles) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FormStyles) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FormStyles) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetHeight

`func (o *FormStyles) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *FormStyles) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *FormStyles) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *FormStyles) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetPadding

`func (o *FormStyles) GetPadding() FormStylesPadding`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *FormStyles) GetPaddingOk() (*FormStylesPadding, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *FormStyles) SetPadding(v FormStylesPadding)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *FormStyles) HasPadding() bool`

HasPadding returns a boolean if a field has been set.

### GetWidth

`func (o *FormStyles) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FormStyles) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FormStyles) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *FormStyles) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### GetWidthUnit

`func (o *FormStyles) GetWidthUnit() EnumFormStylesWidthUnit`

GetWidthUnit returns the WidthUnit field if non-nil, zero value otherwise.

### GetWidthUnitOk

`func (o *FormStyles) GetWidthUnitOk() (*EnumFormStylesWidthUnit, bool)`

GetWidthUnitOk returns a tuple with the WidthUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidthUnit

`func (o *FormStyles) SetWidthUnit(v EnumFormStylesWidthUnit)`

SetWidthUnit sets WidthUnit field to given value.

### HasWidthUnit

`func (o *FormStyles) HasWidthUnit() bool`

HasWidthUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


