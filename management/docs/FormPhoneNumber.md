# FormPhoneNumber

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
**DefaultCountryCode** | Pointer to **string** | The country code to default the country code selection field to (two character country code). | [optional] 
**CountryCodeLabel** | Pointer to **string** | Label for the country code field. | [optional] 
**ValidatePhoneNumber** | Pointer to **bool** | Whether to validate the phone number input. | [optional] 
**ShowExtension** | Pointer to **bool** | Whether to show an extension field. | [optional] 
**ExtensionLabel** | Pointer to **string** | Label for the extension field. This is required and must not be blank if showExtension is true. | [optional] 

## Methods

### NewFormPhoneNumber

`func NewFormPhoneNumber(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, ) *FormPhoneNumber`

NewFormPhoneNumber instantiates a new FormPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormPhoneNumberWithDefaults

`func NewFormPhoneNumberWithDefaults() *FormPhoneNumber`

NewFormPhoneNumberWithDefaults instantiates a new FormPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormPhoneNumber) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormPhoneNumber) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormPhoneNumber) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormPhoneNumber) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormPhoneNumber) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormPhoneNumber) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormPhoneNumber) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormPhoneNumber) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormPhoneNumber) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormPhoneNumber) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAttributeDisabled

`func (o *FormPhoneNumber) GetAttributeDisabled() bool`

GetAttributeDisabled returns the AttributeDisabled field if non-nil, zero value otherwise.

### GetAttributeDisabledOk

`func (o *FormPhoneNumber) GetAttributeDisabledOk() (*bool, bool)`

GetAttributeDisabledOk returns a tuple with the AttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDisabled

`func (o *FormPhoneNumber) SetAttributeDisabled(v bool)`

SetAttributeDisabled sets AttributeDisabled field to given value.

### HasAttributeDisabled

`func (o *FormPhoneNumber) HasAttributeDisabled() bool`

HasAttributeDisabled returns a boolean if a field has been set.

### GetKey

`func (o *FormPhoneNumber) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormPhoneNumber) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormPhoneNumber) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormPhoneNumber) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormPhoneNumber) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormPhoneNumber) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelMode

`func (o *FormPhoneNumber) GetLabelMode() EnumFormElementLabelMode`

GetLabelMode returns the LabelMode field if non-nil, zero value otherwise.

### GetLabelModeOk

`func (o *FormPhoneNumber) GetLabelModeOk() (*EnumFormElementLabelMode, bool)`

GetLabelModeOk returns a tuple with the LabelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMode

`func (o *FormPhoneNumber) SetLabelMode(v EnumFormElementLabelMode)`

SetLabelMode sets LabelMode field to given value.

### HasLabelMode

`func (o *FormPhoneNumber) HasLabelMode() bool`

HasLabelMode returns a boolean if a field has been set.

### GetRequired

`func (o *FormPhoneNumber) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormPhoneNumber) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormPhoneNumber) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormPhoneNumber) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetOtherOptionEnabled

`func (o *FormPhoneNumber) GetOtherOptionEnabled() bool`

GetOtherOptionEnabled returns the OtherOptionEnabled field if non-nil, zero value otherwise.

### GetOtherOptionEnabledOk

`func (o *FormPhoneNumber) GetOtherOptionEnabledOk() (*bool, bool)`

GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionEnabled

`func (o *FormPhoneNumber) SetOtherOptionEnabled(v bool)`

SetOtherOptionEnabled sets OtherOptionEnabled field to given value.

### HasOtherOptionEnabled

`func (o *FormPhoneNumber) HasOtherOptionEnabled() bool`

HasOtherOptionEnabled returns a boolean if a field has been set.

### GetOtherOptionKey

`func (o *FormPhoneNumber) GetOtherOptionKey() string`

GetOtherOptionKey returns the OtherOptionKey field if non-nil, zero value otherwise.

### GetOtherOptionKeyOk

`func (o *FormPhoneNumber) GetOtherOptionKeyOk() (*string, bool)`

GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionKey

`func (o *FormPhoneNumber) SetOtherOptionKey(v string)`

SetOtherOptionKey sets OtherOptionKey field to given value.

### HasOtherOptionKey

`func (o *FormPhoneNumber) HasOtherOptionKey() bool`

HasOtherOptionKey returns a boolean if a field has been set.

### GetOtherOptionLabel

`func (o *FormPhoneNumber) GetOtherOptionLabel() string`

GetOtherOptionLabel returns the OtherOptionLabel field if non-nil, zero value otherwise.

### GetOtherOptionLabelOk

`func (o *FormPhoneNumber) GetOtherOptionLabelOk() (*string, bool)`

GetOtherOptionLabelOk returns a tuple with the OtherOptionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionLabel

`func (o *FormPhoneNumber) SetOtherOptionLabel(v string)`

SetOtherOptionLabel sets OtherOptionLabel field to given value.

### HasOtherOptionLabel

`func (o *FormPhoneNumber) HasOtherOptionLabel() bool`

HasOtherOptionLabel returns a boolean if a field has been set.

### GetOtherOptionInputLabel

`func (o *FormPhoneNumber) GetOtherOptionInputLabel() string`

GetOtherOptionInputLabel returns the OtherOptionInputLabel field if non-nil, zero value otherwise.

### GetOtherOptionInputLabelOk

`func (o *FormPhoneNumber) GetOtherOptionInputLabelOk() (*string, bool)`

GetOtherOptionInputLabelOk returns a tuple with the OtherOptionInputLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionInputLabel

`func (o *FormPhoneNumber) SetOtherOptionInputLabel(v string)`

SetOtherOptionInputLabel sets OtherOptionInputLabel field to given value.

### HasOtherOptionInputLabel

`func (o *FormPhoneNumber) HasOtherOptionInputLabel() bool`

HasOtherOptionInputLabel returns a boolean if a field has been set.

### GetOtherOptionAttributeDisabled

`func (o *FormPhoneNumber) GetOtherOptionAttributeDisabled() bool`

GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field if non-nil, zero value otherwise.

### GetOtherOptionAttributeDisabledOk

`func (o *FormPhoneNumber) GetOtherOptionAttributeDisabledOk() (*bool, bool)`

GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionAttributeDisabled

`func (o *FormPhoneNumber) SetOtherOptionAttributeDisabled(v bool)`

SetOtherOptionAttributeDisabled sets OtherOptionAttributeDisabled field to given value.

### HasOtherOptionAttributeDisabled

`func (o *FormPhoneNumber) HasOtherOptionAttributeDisabled() bool`

HasOtherOptionAttributeDisabled returns a boolean if a field has been set.

### GetDefaultCountryCode

`func (o *FormPhoneNumber) GetDefaultCountryCode() string`

GetDefaultCountryCode returns the DefaultCountryCode field if non-nil, zero value otherwise.

### GetDefaultCountryCodeOk

`func (o *FormPhoneNumber) GetDefaultCountryCodeOk() (*string, bool)`

GetDefaultCountryCodeOk returns a tuple with the DefaultCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCountryCode

`func (o *FormPhoneNumber) SetDefaultCountryCode(v string)`

SetDefaultCountryCode sets DefaultCountryCode field to given value.

### HasDefaultCountryCode

`func (o *FormPhoneNumber) HasDefaultCountryCode() bool`

HasDefaultCountryCode returns a boolean if a field has been set.

### GetCountryCodeLabel

`func (o *FormPhoneNumber) GetCountryCodeLabel() string`

GetCountryCodeLabel returns the CountryCodeLabel field if non-nil, zero value otherwise.

### GetCountryCodeLabelOk

`func (o *FormPhoneNumber) GetCountryCodeLabelOk() (*string, bool)`

GetCountryCodeLabelOk returns a tuple with the CountryCodeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeLabel

`func (o *FormPhoneNumber) SetCountryCodeLabel(v string)`

SetCountryCodeLabel sets CountryCodeLabel field to given value.

### HasCountryCodeLabel

`func (o *FormPhoneNumber) HasCountryCodeLabel() bool`

HasCountryCodeLabel returns a boolean if a field has been set.

### GetValidatePhoneNumber

`func (o *FormPhoneNumber) GetValidatePhoneNumber() bool`

GetValidatePhoneNumber returns the ValidatePhoneNumber field if non-nil, zero value otherwise.

### GetValidatePhoneNumberOk

`func (o *FormPhoneNumber) GetValidatePhoneNumberOk() (*bool, bool)`

GetValidatePhoneNumberOk returns a tuple with the ValidatePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatePhoneNumber

`func (o *FormPhoneNumber) SetValidatePhoneNumber(v bool)`

SetValidatePhoneNumber sets ValidatePhoneNumber field to given value.

### HasValidatePhoneNumber

`func (o *FormPhoneNumber) HasValidatePhoneNumber() bool`

HasValidatePhoneNumber returns a boolean if a field has been set.

### GetShowExtension

`func (o *FormPhoneNumber) GetShowExtension() bool`

GetShowExtension returns the ShowExtension field if non-nil, zero value otherwise.

### GetShowExtensionOk

`func (o *FormPhoneNumber) GetShowExtensionOk() (*bool, bool)`

GetShowExtensionOk returns a tuple with the ShowExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowExtension

`func (o *FormPhoneNumber) SetShowExtension(v bool)`

SetShowExtension sets ShowExtension field to given value.

### HasShowExtension

`func (o *FormPhoneNumber) HasShowExtension() bool`

HasShowExtension returns a boolean if a field has been set.

### GetExtensionLabel

`func (o *FormPhoneNumber) GetExtensionLabel() string`

GetExtensionLabel returns the ExtensionLabel field if non-nil, zero value otherwise.

### GetExtensionLabelOk

`func (o *FormPhoneNumber) GetExtensionLabelOk() (*string, bool)`

GetExtensionLabelOk returns a tuple with the ExtensionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionLabel

`func (o *FormPhoneNumber) SetExtensionLabel(v string)`

SetExtensionLabel sets ExtensionLabel field to given value.

### HasExtensionLabel

`func (o *FormPhoneNumber) HasExtensionLabel() bool`

HasExtensionLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


