# FormFlowLinkStyles

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | Pointer to [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | [optional] 
**TextColor** | Pointer to **string** | A string that specifies the link text color. The value must be a valid hexadecimal color. | [optional] 
**Enabled** | Pointer to **bool** | A boolean that specifies whether the link is enabled. | [optional] 
**Padding** | Pointer to [**FormStylesPadding**](FormStylesPadding.md) |  | [optional] 

## Methods

### NewFormFlowLinkStyles

`func NewFormFlowLinkStyles() *FormFlowLinkStyles`

NewFormFlowLinkStyles instantiates a new FormFlowLinkStyles object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFlowLinkStylesWithDefaults

`func NewFormFlowLinkStylesWithDefaults() *FormFlowLinkStyles`

NewFormFlowLinkStylesWithDefaults instantiates a new FormFlowLinkStyles object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignment

`func (o *FormFlowLinkStyles) GetAlignment() EnumFormItemAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *FormFlowLinkStyles) GetAlignmentOk() (*EnumFormItemAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *FormFlowLinkStyles) SetAlignment(v EnumFormItemAlignment)`

SetAlignment sets Alignment field to given value.

### HasAlignment

`func (o *FormFlowLinkStyles) HasAlignment() bool`

HasAlignment returns a boolean if a field has been set.

### GetTextColor

`func (o *FormFlowLinkStyles) GetTextColor() string`

GetTextColor returns the TextColor field if non-nil, zero value otherwise.

### GetTextColorOk

`func (o *FormFlowLinkStyles) GetTextColorOk() (*string, bool)`

GetTextColorOk returns a tuple with the TextColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextColor

`func (o *FormFlowLinkStyles) SetTextColor(v string)`

SetTextColor sets TextColor field to given value.

### HasTextColor

`func (o *FormFlowLinkStyles) HasTextColor() bool`

HasTextColor returns a boolean if a field has been set.

### GetEnabled

`func (o *FormFlowLinkStyles) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FormFlowLinkStyles) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FormFlowLinkStyles) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FormFlowLinkStyles) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetPadding

`func (o *FormFlowLinkStyles) GetPadding() FormStylesPadding`

GetPadding returns the Padding field if non-nil, zero value otherwise.

### GetPaddingOk

`func (o *FormFlowLinkStyles) GetPaddingOk() (*FormStylesPadding, bool)`

GetPaddingOk returns a tuple with the Padding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPadding

`func (o *FormFlowLinkStyles) SetPadding(v FormStylesPadding)`

SetPadding sets Padding field to given value.

### HasPadding

`func (o *FormFlowLinkStyles) HasPadding() bool`

HasPadding returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


