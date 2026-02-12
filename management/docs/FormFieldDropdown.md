# FormFieldDropdown

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**AttributeDisabled** | Pointer to **bool** | A boolean that specifies whether the linked directory attribute is disabled. | [optional] [readonly] 
**Key** | **string** | A string that specifies an identifier for the field component. | 
**Label** | **string** | A string of escaped JSON that is designed to store a series of text and translatable keys. | 
**LabelMode** | Pointer to [**EnumFormElementLabelMode**](EnumFormElementLabelMode.md) |  | [optional] 
**Required** | Pointer to **bool** | A boolean that specifies whether the field is required. | [optional] 
**OtherOptionEnabled** | Pointer to **bool** | A boolean that specifies whether the end user can type an entry that is not in a predefined list. | [optional] 
**OtherOptionKey** | Pointer to **string** | A string that specifies whether the form identifies that the choice is a custom choice not from a predefined list. | [optional] 
**OtherOptionLabel** | Pointer to **string** | A string that specifies the label for a custom or \&quot;other\&quot; choice in a list. | [optional] 
**OtherOptionInputLabel** | Pointer to **string** | A string that specifies the label for the other option in drop-down controls. | [optional] 
**OtherOptionAttributeDisabled** | Pointer to **bool** | A boolean that specifies whether the directory attribute option is disabled. Set to true if it references a PingOne directory attribute. | [optional] 
**Layout** | Pointer to [**EnumFormElementLayout**](EnumFormElementLayout.md) |  | [optional] 
**Options** | [**[]FormElementOption**](FormElementOption.md) | An array of objects (label/value pairs) that specifies the unique list of options. This is a required property when the type is &#x60;RADIO&#x60;, &#x60;CHECKBOX&#x60;, or &#x60;DROPDOWN&#x60;. | 
**Validation** | Pointer to [**FormElementValidation**](FormElementValidation.md) |  | [optional] 

## Methods

### NewFormFieldDropdown

`func NewFormFieldDropdown(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, options []FormElementOption, ) *FormFieldDropdown`

NewFormFieldDropdown instantiates a new FormFieldDropdown object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldDropdownWithDefaults

`func NewFormFieldDropdownWithDefaults() *FormFieldDropdown`

NewFormFieldDropdownWithDefaults instantiates a new FormFieldDropdown object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldDropdown) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldDropdown) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldDropdown) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldDropdown) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldDropdown) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldDropdown) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldDropdown) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldDropdown) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldDropdown) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldDropdown) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAttributeDisabled

`func (o *FormFieldDropdown) GetAttributeDisabled() bool`

GetAttributeDisabled returns the AttributeDisabled field if non-nil, zero value otherwise.

### GetAttributeDisabledOk

`func (o *FormFieldDropdown) GetAttributeDisabledOk() (*bool, bool)`

GetAttributeDisabledOk returns a tuple with the AttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDisabled

`func (o *FormFieldDropdown) SetAttributeDisabled(v bool)`

SetAttributeDisabled sets AttributeDisabled field to given value.

### HasAttributeDisabled

`func (o *FormFieldDropdown) HasAttributeDisabled() bool`

HasAttributeDisabled returns a boolean if a field has been set.

### GetKey

`func (o *FormFieldDropdown) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldDropdown) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldDropdown) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormFieldDropdown) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldDropdown) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldDropdown) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelMode

`func (o *FormFieldDropdown) GetLabelMode() EnumFormElementLabelMode`

GetLabelMode returns the LabelMode field if non-nil, zero value otherwise.

### GetLabelModeOk

`func (o *FormFieldDropdown) GetLabelModeOk() (*EnumFormElementLabelMode, bool)`

GetLabelModeOk returns a tuple with the LabelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMode

`func (o *FormFieldDropdown) SetLabelMode(v EnumFormElementLabelMode)`

SetLabelMode sets LabelMode field to given value.

### HasLabelMode

`func (o *FormFieldDropdown) HasLabelMode() bool`

HasLabelMode returns a boolean if a field has been set.

### GetRequired

`func (o *FormFieldDropdown) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldDropdown) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldDropdown) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormFieldDropdown) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetOtherOptionEnabled

`func (o *FormFieldDropdown) GetOtherOptionEnabled() bool`

GetOtherOptionEnabled returns the OtherOptionEnabled field if non-nil, zero value otherwise.

### GetOtherOptionEnabledOk

`func (o *FormFieldDropdown) GetOtherOptionEnabledOk() (*bool, bool)`

GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionEnabled

`func (o *FormFieldDropdown) SetOtherOptionEnabled(v bool)`

SetOtherOptionEnabled sets OtherOptionEnabled field to given value.

### HasOtherOptionEnabled

`func (o *FormFieldDropdown) HasOtherOptionEnabled() bool`

HasOtherOptionEnabled returns a boolean if a field has been set.

### GetOtherOptionKey

`func (o *FormFieldDropdown) GetOtherOptionKey() string`

GetOtherOptionKey returns the OtherOptionKey field if non-nil, zero value otherwise.

### GetOtherOptionKeyOk

`func (o *FormFieldDropdown) GetOtherOptionKeyOk() (*string, bool)`

GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionKey

`func (o *FormFieldDropdown) SetOtherOptionKey(v string)`

SetOtherOptionKey sets OtherOptionKey field to given value.

### HasOtherOptionKey

`func (o *FormFieldDropdown) HasOtherOptionKey() bool`

HasOtherOptionKey returns a boolean if a field has been set.

### GetOtherOptionLabel

`func (o *FormFieldDropdown) GetOtherOptionLabel() string`

GetOtherOptionLabel returns the OtherOptionLabel field if non-nil, zero value otherwise.

### GetOtherOptionLabelOk

`func (o *FormFieldDropdown) GetOtherOptionLabelOk() (*string, bool)`

GetOtherOptionLabelOk returns a tuple with the OtherOptionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionLabel

`func (o *FormFieldDropdown) SetOtherOptionLabel(v string)`

SetOtherOptionLabel sets OtherOptionLabel field to given value.

### HasOtherOptionLabel

`func (o *FormFieldDropdown) HasOtherOptionLabel() bool`

HasOtherOptionLabel returns a boolean if a field has been set.

### GetOtherOptionInputLabel

`func (o *FormFieldDropdown) GetOtherOptionInputLabel() string`

GetOtherOptionInputLabel returns the OtherOptionInputLabel field if non-nil, zero value otherwise.

### GetOtherOptionInputLabelOk

`func (o *FormFieldDropdown) GetOtherOptionInputLabelOk() (*string, bool)`

GetOtherOptionInputLabelOk returns a tuple with the OtherOptionInputLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionInputLabel

`func (o *FormFieldDropdown) SetOtherOptionInputLabel(v string)`

SetOtherOptionInputLabel sets OtherOptionInputLabel field to given value.

### HasOtherOptionInputLabel

`func (o *FormFieldDropdown) HasOtherOptionInputLabel() bool`

HasOtherOptionInputLabel returns a boolean if a field has been set.

### GetOtherOptionAttributeDisabled

`func (o *FormFieldDropdown) GetOtherOptionAttributeDisabled() bool`

GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field if non-nil, zero value otherwise.

### GetOtherOptionAttributeDisabledOk

`func (o *FormFieldDropdown) GetOtherOptionAttributeDisabledOk() (*bool, bool)`

GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionAttributeDisabled

`func (o *FormFieldDropdown) SetOtherOptionAttributeDisabled(v bool)`

SetOtherOptionAttributeDisabled sets OtherOptionAttributeDisabled field to given value.

### HasOtherOptionAttributeDisabled

`func (o *FormFieldDropdown) HasOtherOptionAttributeDisabled() bool`

HasOtherOptionAttributeDisabled returns a boolean if a field has been set.

### GetLayout

`func (o *FormFieldDropdown) GetLayout() EnumFormElementLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FormFieldDropdown) GetLayoutOk() (*EnumFormElementLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FormFieldDropdown) SetLayout(v EnumFormElementLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *FormFieldDropdown) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetOptions

`func (o *FormFieldDropdown) GetOptions() []FormElementOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldDropdown) GetOptionsOk() (*[]FormElementOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldDropdown) SetOptions(v []FormElementOption)`

SetOptions sets Options field to given value.


### GetValidation

`func (o *FormFieldDropdown) GetValidation() FormElementValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *FormFieldDropdown) GetValidationOk() (*FormElementValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *FormFieldDropdown) SetValidation(v FormElementValidation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *FormFieldDropdown) HasValidation() bool`

HasValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


