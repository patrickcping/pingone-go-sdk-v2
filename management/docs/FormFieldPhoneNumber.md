# FormFieldPhoneNumber

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

### NewFormFieldPhoneNumber

`func NewFormFieldPhoneNumber(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, ) *FormFieldPhoneNumber`

NewFormFieldPhoneNumber instantiates a new FormFieldPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldPhoneNumberWithDefaults

`func NewFormFieldPhoneNumberWithDefaults() *FormFieldPhoneNumber`

NewFormFieldPhoneNumberWithDefaults instantiates a new FormFieldPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldPhoneNumber) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldPhoneNumber) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldPhoneNumber) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldPhoneNumber) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldPhoneNumber) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldPhoneNumber) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldPhoneNumber) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldPhoneNumber) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldPhoneNumber) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldPhoneNumber) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAttributeDisabled

`func (o *FormFieldPhoneNumber) GetAttributeDisabled() bool`

GetAttributeDisabled returns the AttributeDisabled field if non-nil, zero value otherwise.

### GetAttributeDisabledOk

`func (o *FormFieldPhoneNumber) GetAttributeDisabledOk() (*bool, bool)`

GetAttributeDisabledOk returns a tuple with the AttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDisabled

`func (o *FormFieldPhoneNumber) SetAttributeDisabled(v bool)`

SetAttributeDisabled sets AttributeDisabled field to given value.

### HasAttributeDisabled

`func (o *FormFieldPhoneNumber) HasAttributeDisabled() bool`

HasAttributeDisabled returns a boolean if a field has been set.

### GetKey

`func (o *FormFieldPhoneNumber) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldPhoneNumber) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldPhoneNumber) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormFieldPhoneNumber) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldPhoneNumber) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldPhoneNumber) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelMode

`func (o *FormFieldPhoneNumber) GetLabelMode() EnumFormElementLabelMode`

GetLabelMode returns the LabelMode field if non-nil, zero value otherwise.

### GetLabelModeOk

`func (o *FormFieldPhoneNumber) GetLabelModeOk() (*EnumFormElementLabelMode, bool)`

GetLabelModeOk returns a tuple with the LabelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMode

`func (o *FormFieldPhoneNumber) SetLabelMode(v EnumFormElementLabelMode)`

SetLabelMode sets LabelMode field to given value.

### HasLabelMode

`func (o *FormFieldPhoneNumber) HasLabelMode() bool`

HasLabelMode returns a boolean if a field has been set.

### GetRequired

`func (o *FormFieldPhoneNumber) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldPhoneNumber) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldPhoneNumber) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormFieldPhoneNumber) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetOtherOptionEnabled

`func (o *FormFieldPhoneNumber) GetOtherOptionEnabled() bool`

GetOtherOptionEnabled returns the OtherOptionEnabled field if non-nil, zero value otherwise.

### GetOtherOptionEnabledOk

`func (o *FormFieldPhoneNumber) GetOtherOptionEnabledOk() (*bool, bool)`

GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionEnabled

`func (o *FormFieldPhoneNumber) SetOtherOptionEnabled(v bool)`

SetOtherOptionEnabled sets OtherOptionEnabled field to given value.

### HasOtherOptionEnabled

`func (o *FormFieldPhoneNumber) HasOtherOptionEnabled() bool`

HasOtherOptionEnabled returns a boolean if a field has been set.

### GetOtherOptionKey

`func (o *FormFieldPhoneNumber) GetOtherOptionKey() string`

GetOtherOptionKey returns the OtherOptionKey field if non-nil, zero value otherwise.

### GetOtherOptionKeyOk

`func (o *FormFieldPhoneNumber) GetOtherOptionKeyOk() (*string, bool)`

GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionKey

`func (o *FormFieldPhoneNumber) SetOtherOptionKey(v string)`

SetOtherOptionKey sets OtherOptionKey field to given value.

### HasOtherOptionKey

`func (o *FormFieldPhoneNumber) HasOtherOptionKey() bool`

HasOtherOptionKey returns a boolean if a field has been set.

### GetOtherOptionLabel

`func (o *FormFieldPhoneNumber) GetOtherOptionLabel() string`

GetOtherOptionLabel returns the OtherOptionLabel field if non-nil, zero value otherwise.

### GetOtherOptionLabelOk

`func (o *FormFieldPhoneNumber) GetOtherOptionLabelOk() (*string, bool)`

GetOtherOptionLabelOk returns a tuple with the OtherOptionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionLabel

`func (o *FormFieldPhoneNumber) SetOtherOptionLabel(v string)`

SetOtherOptionLabel sets OtherOptionLabel field to given value.

### HasOtherOptionLabel

`func (o *FormFieldPhoneNumber) HasOtherOptionLabel() bool`

HasOtherOptionLabel returns a boolean if a field has been set.

### GetOtherOptionInputLabel

`func (o *FormFieldPhoneNumber) GetOtherOptionInputLabel() string`

GetOtherOptionInputLabel returns the OtherOptionInputLabel field if non-nil, zero value otherwise.

### GetOtherOptionInputLabelOk

`func (o *FormFieldPhoneNumber) GetOtherOptionInputLabelOk() (*string, bool)`

GetOtherOptionInputLabelOk returns a tuple with the OtherOptionInputLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionInputLabel

`func (o *FormFieldPhoneNumber) SetOtherOptionInputLabel(v string)`

SetOtherOptionInputLabel sets OtherOptionInputLabel field to given value.

### HasOtherOptionInputLabel

`func (o *FormFieldPhoneNumber) HasOtherOptionInputLabel() bool`

HasOtherOptionInputLabel returns a boolean if a field has been set.

### GetOtherOptionAttributeDisabled

`func (o *FormFieldPhoneNumber) GetOtherOptionAttributeDisabled() bool`

GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field if non-nil, zero value otherwise.

### GetOtherOptionAttributeDisabledOk

`func (o *FormFieldPhoneNumber) GetOtherOptionAttributeDisabledOk() (*bool, bool)`

GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionAttributeDisabled

`func (o *FormFieldPhoneNumber) SetOtherOptionAttributeDisabled(v bool)`

SetOtherOptionAttributeDisabled sets OtherOptionAttributeDisabled field to given value.

### HasOtherOptionAttributeDisabled

`func (o *FormFieldPhoneNumber) HasOtherOptionAttributeDisabled() bool`

HasOtherOptionAttributeDisabled returns a boolean if a field has been set.

### GetDefaultCountryCode

`func (o *FormFieldPhoneNumber) GetDefaultCountryCode() string`

GetDefaultCountryCode returns the DefaultCountryCode field if non-nil, zero value otherwise.

### GetDefaultCountryCodeOk

`func (o *FormFieldPhoneNumber) GetDefaultCountryCodeOk() (*string, bool)`

GetDefaultCountryCodeOk returns a tuple with the DefaultCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCountryCode

`func (o *FormFieldPhoneNumber) SetDefaultCountryCode(v string)`

SetDefaultCountryCode sets DefaultCountryCode field to given value.

### HasDefaultCountryCode

`func (o *FormFieldPhoneNumber) HasDefaultCountryCode() bool`

HasDefaultCountryCode returns a boolean if a field has been set.

### GetCountryCodeLabel

`func (o *FormFieldPhoneNumber) GetCountryCodeLabel() string`

GetCountryCodeLabel returns the CountryCodeLabel field if non-nil, zero value otherwise.

### GetCountryCodeLabelOk

`func (o *FormFieldPhoneNumber) GetCountryCodeLabelOk() (*string, bool)`

GetCountryCodeLabelOk returns a tuple with the CountryCodeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeLabel

`func (o *FormFieldPhoneNumber) SetCountryCodeLabel(v string)`

SetCountryCodeLabel sets CountryCodeLabel field to given value.

### HasCountryCodeLabel

`func (o *FormFieldPhoneNumber) HasCountryCodeLabel() bool`

HasCountryCodeLabel returns a boolean if a field has been set.

### GetValidatePhoneNumber

`func (o *FormFieldPhoneNumber) GetValidatePhoneNumber() bool`

GetValidatePhoneNumber returns the ValidatePhoneNumber field if non-nil, zero value otherwise.

### GetValidatePhoneNumberOk

`func (o *FormFieldPhoneNumber) GetValidatePhoneNumberOk() (*bool, bool)`

GetValidatePhoneNumberOk returns a tuple with the ValidatePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatePhoneNumber

`func (o *FormFieldPhoneNumber) SetValidatePhoneNumber(v bool)`

SetValidatePhoneNumber sets ValidatePhoneNumber field to given value.

### HasValidatePhoneNumber

`func (o *FormFieldPhoneNumber) HasValidatePhoneNumber() bool`

HasValidatePhoneNumber returns a boolean if a field has been set.

### GetShowExtension

`func (o *FormFieldPhoneNumber) GetShowExtension() bool`

GetShowExtension returns the ShowExtension field if non-nil, zero value otherwise.

### GetShowExtensionOk

`func (o *FormFieldPhoneNumber) GetShowExtensionOk() (*bool, bool)`

GetShowExtensionOk returns a tuple with the ShowExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowExtension

`func (o *FormFieldPhoneNumber) SetShowExtension(v bool)`

SetShowExtension sets ShowExtension field to given value.

### HasShowExtension

`func (o *FormFieldPhoneNumber) HasShowExtension() bool`

HasShowExtension returns a boolean if a field has been set.

### GetExtensionLabel

`func (o *FormFieldPhoneNumber) GetExtensionLabel() string`

GetExtensionLabel returns the ExtensionLabel field if non-nil, zero value otherwise.

### GetExtensionLabelOk

`func (o *FormFieldPhoneNumber) GetExtensionLabelOk() (*string, bool)`

GetExtensionLabelOk returns a tuple with the ExtensionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionLabel

`func (o *FormFieldPhoneNumber) SetExtensionLabel(v string)`

SetExtensionLabel sets ExtensionLabel field to given value.

### HasExtensionLabel

`func (o *FormFieldPhoneNumber) HasExtensionLabel() bool`

HasExtensionLabel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


