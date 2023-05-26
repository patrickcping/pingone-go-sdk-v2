# FormFieldPassword

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
**Layout** | Pointer to [**EnumFormElementLayout**](EnumFormElementLayout.md) |  | [optional] 
**Options** | Pointer to **[]string** | An array of strings that specifies the unique list of options. This is a required property when the type is &#x60;RADIO&#x60;, &#x60;CHECKBOX&#x60;, or &#x60;DROPDOWN&#x60;. | [optional] 
**Validation** | Pointer to [**FormElementValidation**](FormElementValidation.md) |  | [optional] 

## Methods

### NewFormFieldPassword

`func NewFormFieldPassword(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, required bool, ) *FormFieldPassword`

NewFormFieldPassword instantiates a new FormFieldPassword object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldPasswordWithDefaults

`func NewFormFieldPasswordWithDefaults() *FormFieldPassword`

NewFormFieldPasswordWithDefaults instantiates a new FormFieldPassword object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldPassword) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldPassword) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldPassword) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldPassword) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldPassword) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldPassword) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetAttributeDisabled

`func (o *FormFieldPassword) GetAttributeDisabled() bool`

GetAttributeDisabled returns the AttributeDisabled field if non-nil, zero value otherwise.

### GetAttributeDisabledOk

`func (o *FormFieldPassword) GetAttributeDisabledOk() (*bool, bool)`

GetAttributeDisabledOk returns a tuple with the AttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDisabled

`func (o *FormFieldPassword) SetAttributeDisabled(v bool)`

SetAttributeDisabled sets AttributeDisabled field to given value.

### HasAttributeDisabled

`func (o *FormFieldPassword) HasAttributeDisabled() bool`

HasAttributeDisabled returns a boolean if a field has been set.

### GetKey

`func (o *FormFieldPassword) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldPassword) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldPassword) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormFieldPassword) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldPassword) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldPassword) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *FormFieldPassword) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLabelMode

`func (o *FormFieldPassword) GetLabelMode() EnumFormElementLabelMode`

GetLabelMode returns the LabelMode field if non-nil, zero value otherwise.

### GetLabelModeOk

`func (o *FormFieldPassword) GetLabelModeOk() (*EnumFormElementLabelMode, bool)`

GetLabelModeOk returns a tuple with the LabelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMode

`func (o *FormFieldPassword) SetLabelMode(v EnumFormElementLabelMode)`

SetLabelMode sets LabelMode field to given value.

### HasLabelMode

`func (o *FormFieldPassword) HasLabelMode() bool`

HasLabelMode returns a boolean if a field has been set.

### GetRequired

`func (o *FormFieldPassword) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldPassword) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldPassword) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetOtherOptionEnabled

`func (o *FormFieldPassword) GetOtherOptionEnabled() bool`

GetOtherOptionEnabled returns the OtherOptionEnabled field if non-nil, zero value otherwise.

### GetOtherOptionEnabledOk

`func (o *FormFieldPassword) GetOtherOptionEnabledOk() (*bool, bool)`

GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionEnabled

`func (o *FormFieldPassword) SetOtherOptionEnabled(v bool)`

SetOtherOptionEnabled sets OtherOptionEnabled field to given value.

### HasOtherOptionEnabled

`func (o *FormFieldPassword) HasOtherOptionEnabled() bool`

HasOtherOptionEnabled returns a boolean if a field has been set.

### GetOtherOptionKey

`func (o *FormFieldPassword) GetOtherOptionKey() string`

GetOtherOptionKey returns the OtherOptionKey field if non-nil, zero value otherwise.

### GetOtherOptionKeyOk

`func (o *FormFieldPassword) GetOtherOptionKeyOk() (*string, bool)`

GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionKey

`func (o *FormFieldPassword) SetOtherOptionKey(v string)`

SetOtherOptionKey sets OtherOptionKey field to given value.

### HasOtherOptionKey

`func (o *FormFieldPassword) HasOtherOptionKey() bool`

HasOtherOptionKey returns a boolean if a field has been set.

### GetOtherOptionlabel

`func (o *FormFieldPassword) GetOtherOptionlabel() string`

GetOtherOptionlabel returns the OtherOptionlabel field if non-nil, zero value otherwise.

### GetOtherOptionlabelOk

`func (o *FormFieldPassword) GetOtherOptionlabelOk() (*string, bool)`

GetOtherOptionlabelOk returns a tuple with the OtherOptionlabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionlabel

`func (o *FormFieldPassword) SetOtherOptionlabel(v string)`

SetOtherOptionlabel sets OtherOptionlabel field to given value.

### HasOtherOptionlabel

`func (o *FormFieldPassword) HasOtherOptionlabel() bool`

HasOtherOptionlabel returns a boolean if a field has been set.

### GetOtherOptionInputlabel

`func (o *FormFieldPassword) GetOtherOptionInputlabel() string`

GetOtherOptionInputlabel returns the OtherOptionInputlabel field if non-nil, zero value otherwise.

### GetOtherOptionInputlabelOk

`func (o *FormFieldPassword) GetOtherOptionInputlabelOk() (*string, bool)`

GetOtherOptionInputlabelOk returns a tuple with the OtherOptionInputlabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionInputlabel

`func (o *FormFieldPassword) SetOtherOptionInputlabel(v string)`

SetOtherOptionInputlabel sets OtherOptionInputlabel field to given value.

### HasOtherOptionInputlabel

`func (o *FormFieldPassword) HasOtherOptionInputlabel() bool`

HasOtherOptionInputlabel returns a boolean if a field has been set.

### GetOtherOptionAttributeDisabled

`func (o *FormFieldPassword) GetOtherOptionAttributeDisabled() bool`

GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field if non-nil, zero value otherwise.

### GetOtherOptionAttributeDisabledOk

`func (o *FormFieldPassword) GetOtherOptionAttributeDisabledOk() (*bool, bool)`

GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionAttributeDisabled

`func (o *FormFieldPassword) SetOtherOptionAttributeDisabled(v bool)`

SetOtherOptionAttributeDisabled sets OtherOptionAttributeDisabled field to given value.

### HasOtherOptionAttributeDisabled

`func (o *FormFieldPassword) HasOtherOptionAttributeDisabled() bool`

HasOtherOptionAttributeDisabled returns a boolean if a field has been set.

### GetLayout

`func (o *FormFieldPassword) GetLayout() EnumFormElementLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FormFieldPassword) GetLayoutOk() (*EnumFormElementLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FormFieldPassword) SetLayout(v EnumFormElementLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *FormFieldPassword) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetOptions

`func (o *FormFieldPassword) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldPassword) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldPassword) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FormFieldPassword) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetValidation

`func (o *FormFieldPassword) GetValidation() FormElementValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *FormFieldPassword) GetValidationOk() (*FormElementValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *FormFieldPassword) SetValidation(v FormElementValidation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *FormFieldPassword) HasValidation() bool`

HasValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

