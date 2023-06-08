# FormFieldRadio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**AttributeDisabled** | Pointer to **bool** | A boolean that specifies whether the linked directory attribute is disabled. | [optional] [readonly] 
**Key** | **string** | A string that specifies an identifier for the field component. | 
**Label** | Pointer to **string** | A string of escaped JSON that is designed to store a series of text and translatable keys. | [optional] 
**LabelMode** | Pointer to [**EnumFormElementLabelMode**](EnumFormElementLabelMode.md) |  | [optional] 
**Required** | **bool** | A boolean that specifies whether the field is required. | 
**OtherOptionEnabled** | Pointer to **bool** | A boolean that specifies whether the end user can type an entry that is not in a predefined list. | [optional] 
**OtherOptionKey** | Pointer to **string** | A string that specifies whether the form identifies that the choice is a custom choice not from a predefined list. | [optional] 
**OtherOptionlabel** | Pointer to **string** | A string that specifies the label for a custom or \&quot;other\&quot; choice in a list. | [optional] 
**OtherOptionInputlabel** | Pointer to **string** | A string that specifies the label for the other option in drop-down controls. | [optional] 
**OtherOptionAttributeDisabled** | Pointer to **bool** | A boolean that specifies whether the directory attribute option is disabled. Set to true if it references a PingOne directory attribute. | [optional] 
**Layout** | [**EnumFormElementLayout**](EnumFormElementLayout.md) |  | 
**Options** | **[]string** | An array of strings that specifies the unique list of options. This is a required property when the type is &#x60;RADIO&#x60;, &#x60;CHECKBOX&#x60;, or &#x60;DROPDOWN&#x60;. | 
**Validation** | Pointer to [**FormElementValidation**](FormElementValidation.md) |  | [optional] 

## Methods

### NewFormFieldRadio

`func NewFormFieldRadio(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, required bool, layout EnumFormElementLayout, options []string, ) *FormFieldRadio`

NewFormFieldRadio instantiates a new FormFieldRadio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldRadioWithDefaults

`func NewFormFieldRadioWithDefaults() *FormFieldRadio`

NewFormFieldRadioWithDefaults instantiates a new FormFieldRadio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldRadio) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldRadio) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldRadio) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldRadio) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldRadio) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldRadio) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetAttributeDisabled

`func (o *FormFieldRadio) GetAttributeDisabled() bool`

GetAttributeDisabled returns the AttributeDisabled field if non-nil, zero value otherwise.

### GetAttributeDisabledOk

`func (o *FormFieldRadio) GetAttributeDisabledOk() (*bool, bool)`

GetAttributeDisabledOk returns a tuple with the AttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDisabled

`func (o *FormFieldRadio) SetAttributeDisabled(v bool)`

SetAttributeDisabled sets AttributeDisabled field to given value.

### HasAttributeDisabled

`func (o *FormFieldRadio) HasAttributeDisabled() bool`

HasAttributeDisabled returns a boolean if a field has been set.

### GetKey

`func (o *FormFieldRadio) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldRadio) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldRadio) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormFieldRadio) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldRadio) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldRadio) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormFieldRadio) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLabelMode

`func (o *FormFieldRadio) GetLabelMode() EnumFormElementLabelMode`

GetLabelMode returns the LabelMode field if non-nil, zero value otherwise.

### GetLabelModeOk

`func (o *FormFieldRadio) GetLabelModeOk() (*EnumFormElementLabelMode, bool)`

GetLabelModeOk returns a tuple with the LabelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMode

`func (o *FormFieldRadio) SetLabelMode(v EnumFormElementLabelMode)`

SetLabelMode sets LabelMode field to given value.

### HasLabelMode

`func (o *FormFieldRadio) HasLabelMode() bool`

HasLabelMode returns a boolean if a field has been set.

### GetRequired

`func (o *FormFieldRadio) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldRadio) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldRadio) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetOtherOptionEnabled

`func (o *FormFieldRadio) GetOtherOptionEnabled() bool`

GetOtherOptionEnabled returns the OtherOptionEnabled field if non-nil, zero value otherwise.

### GetOtherOptionEnabledOk

`func (o *FormFieldRadio) GetOtherOptionEnabledOk() (*bool, bool)`

GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionEnabled

`func (o *FormFieldRadio) SetOtherOptionEnabled(v bool)`

SetOtherOptionEnabled sets OtherOptionEnabled field to given value.

### HasOtherOptionEnabled

`func (o *FormFieldRadio) HasOtherOptionEnabled() bool`

HasOtherOptionEnabled returns a boolean if a field has been set.

### GetOtherOptionKey

`func (o *FormFieldRadio) GetOtherOptionKey() string`

GetOtherOptionKey returns the OtherOptionKey field if non-nil, zero value otherwise.

### GetOtherOptionKeyOk

`func (o *FormFieldRadio) GetOtherOptionKeyOk() (*string, bool)`

GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionKey

`func (o *FormFieldRadio) SetOtherOptionKey(v string)`

SetOtherOptionKey sets OtherOptionKey field to given value.

### HasOtherOptionKey

`func (o *FormFieldRadio) HasOtherOptionKey() bool`

HasOtherOptionKey returns a boolean if a field has been set.

### GetOtherOptionlabel

`func (o *FormFieldRadio) GetOtherOptionlabel() string`

GetOtherOptionlabel returns the OtherOptionlabel field if non-nil, zero value otherwise.

### GetOtherOptionlabelOk

`func (o *FormFieldRadio) GetOtherOptionlabelOk() (*string, bool)`

GetOtherOptionlabelOk returns a tuple with the OtherOptionlabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionlabel

`func (o *FormFieldRadio) SetOtherOptionlabel(v string)`

SetOtherOptionlabel sets OtherOptionlabel field to given value.

### HasOtherOptionlabel

`func (o *FormFieldRadio) HasOtherOptionlabel() bool`

HasOtherOptionlabel returns a boolean if a field has been set.

### GetOtherOptionInputlabel

`func (o *FormFieldRadio) GetOtherOptionInputlabel() string`

GetOtherOptionInputlabel returns the OtherOptionInputlabel field if non-nil, zero value otherwise.

### GetOtherOptionInputlabelOk

`func (o *FormFieldRadio) GetOtherOptionInputlabelOk() (*string, bool)`

GetOtherOptionInputlabelOk returns a tuple with the OtherOptionInputlabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionInputlabel

`func (o *FormFieldRadio) SetOtherOptionInputlabel(v string)`

SetOtherOptionInputlabel sets OtherOptionInputlabel field to given value.

### HasOtherOptionInputlabel

`func (o *FormFieldRadio) HasOtherOptionInputlabel() bool`

HasOtherOptionInputlabel returns a boolean if a field has been set.

### GetOtherOptionAttributeDisabled

`func (o *FormFieldRadio) GetOtherOptionAttributeDisabled() bool`

GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field if non-nil, zero value otherwise.

### GetOtherOptionAttributeDisabledOk

`func (o *FormFieldRadio) GetOtherOptionAttributeDisabledOk() (*bool, bool)`

GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionAttributeDisabled

`func (o *FormFieldRadio) SetOtherOptionAttributeDisabled(v bool)`

SetOtherOptionAttributeDisabled sets OtherOptionAttributeDisabled field to given value.

### HasOtherOptionAttributeDisabled

`func (o *FormFieldRadio) HasOtherOptionAttributeDisabled() bool`

HasOtherOptionAttributeDisabled returns a boolean if a field has been set.

### GetLayout

`func (o *FormFieldRadio) GetLayout() EnumFormElementLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FormFieldRadio) GetLayoutOk() (*EnumFormElementLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FormFieldRadio) SetLayout(v EnumFormElementLayout)`

SetLayout sets Layout field to given value.


### GetOptions

`func (o *FormFieldRadio) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldRadio) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldRadio) SetOptions(v []string)`

SetOptions sets Options field to given value.


### GetValidation

`func (o *FormFieldRadio) GetValidation() FormElementValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *FormFieldRadio) GetValidationOk() (*FormElementValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *FormFieldRadio) SetValidation(v FormElementValidation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *FormFieldRadio) HasValidation() bool`

HasValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

