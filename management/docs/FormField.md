# FormField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**AttributeDisabled** | Pointer to **bool** | A boolean that specifies whether the linked directory attribute is disabled. | [optional] [readonly] 
**Key** | **string** | A string that specifies an identifier for the field component. | 
**Label** | **string** | A string that specifies the social login button label. | 
**LabelMode** | Pointer to [**EnumFormElementLabelMode**](EnumFormElementLabelMode.md) |  | [optional] 
**Required** | **bool** | A boolean that specifies whether the field is required. | 
**OtherOptionEnabled** | Pointer to **bool** | A boolean that specifies whether the end user can type an entry that is not in a predefined list. | [optional] 
**OtherOptionKey** | Pointer to **string** | A string that specifies whether the form identifies that the choice is a custom choice not from a predefined list. | [optional] 
**OtherOptionlabel** | Pointer to **string** | A string that specifies the label for a custom or \&quot;other\&quot; choice in a list. | [optional] 
**OtherOptionInputlabel** | Pointer to **string** | A string that specifies the label for the other option in drop-down controls. | [optional] 
**OtherOptionAttributeDisabled** | Pointer to **bool** | A boolean that specifies whether the directory attribute option is disabled. Set to true if it references a PingOne directory attribute. | [optional] 
**Layout** | [**EnumFormElementLayout**](EnumFormElementLayout.md) |  | 
**Options** | **[]string** | An array of strings that specifies the unique list of options. This is a required property when the type is &#x60;RADIO&#x60;, &#x60;CHECKBOX&#x60;, or &#x60;DROPDOWN&#x60;. | 
**Validation** | [**FormElementValidation**](FormElementValidation.md) |  | 
**LabelPasswordVerify** | Pointer to **string** | A string that when a second field for verifies password is used, this poperty specifies the field label for that verify field. | [optional] 
**Value** | **string** | A string that specifies the value of the field if this option is selected. | 
**Content** | Pointer to **string** | A string that specifies the field content. | [optional] 
**Styles** | Pointer to [**FormSocialLoginButtonStyles**](FormSocialLoginButtonStyles.md) |  | [optional] 
**Size** | [**EnumFormRecaptchaV2Size**](EnumFormRecaptchaV2Size.md) |  | 
**Theme** | [**EnumFormRecaptchaV2Theme**](EnumFormRecaptchaV2Theme.md) |  | 
**Alignment** | [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | 
**QrCodeType** | [**EnumFormQrCodeType**](EnumFormQrCodeType.md) |  | 
**ShowBorder** | **bool** | A boolean that specifies the border visibility. | 
**IdpType** | [**EnumFormSocialLoginIdpType**](EnumFormSocialLoginIdpType.md) |  | 
**IdpName** | **string** | A string that specifies the external identity provider name. | 
**IdpId** | **string** | A string that specifies the external identity provider&#39;s ID. | 
**IdpEnabled** | **bool** | A boolean that specifies whether the external identity provider is enabled. | 
**IconSrc** | **string** | A string that specifies the HTTP link (URL format) for the external identity provider&#39;s icon. | 
**Width** | Pointer to **int32** | An integer that specifies the button width. Set as a percentage. | [optional] 

## Methods

### NewFormField

`func NewFormField(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, required bool, layout EnumFormElementLayout, options []string, validation FormElementValidation, value string, size EnumFormRecaptchaV2Size, theme EnumFormRecaptchaV2Theme, alignment EnumFormItemAlignment, qrCodeType EnumFormQrCodeType, showBorder bool, idpType EnumFormSocialLoginIdpType, idpName string, idpId string, idpEnabled bool, iconSrc string, ) *FormField`

NewFormField instantiates a new FormField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldWithDefaults

`func NewFormFieldWithDefaults() *FormField`

NewFormFieldWithDefaults instantiates a new FormField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormField) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormField) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormField) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormField) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormField) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormField) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetAttributeDisabled

`func (o *FormField) GetAttributeDisabled() bool`

GetAttributeDisabled returns the AttributeDisabled field if non-nil, zero value otherwise.

### GetAttributeDisabledOk

`func (o *FormField) GetAttributeDisabledOk() (*bool, bool)`

GetAttributeDisabledOk returns a tuple with the AttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeDisabled

`func (o *FormField) SetAttributeDisabled(v bool)`

SetAttributeDisabled sets AttributeDisabled field to given value.

### HasAttributeDisabled

`func (o *FormField) HasAttributeDisabled() bool`

HasAttributeDisabled returns a boolean if a field has been set.

### GetKey

`func (o *FormField) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormField) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormField) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormField) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLabelMode

`func (o *FormField) GetLabelMode() EnumFormElementLabelMode`

GetLabelMode returns the LabelMode field if non-nil, zero value otherwise.

### GetLabelModeOk

`func (o *FormField) GetLabelModeOk() (*EnumFormElementLabelMode, bool)`

GetLabelModeOk returns a tuple with the LabelMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelMode

`func (o *FormField) SetLabelMode(v EnumFormElementLabelMode)`

SetLabelMode sets LabelMode field to given value.

### HasLabelMode

`func (o *FormField) HasLabelMode() bool`

HasLabelMode returns a boolean if a field has been set.

### GetRequired

`func (o *FormField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetOtherOptionEnabled

`func (o *FormField) GetOtherOptionEnabled() bool`

GetOtherOptionEnabled returns the OtherOptionEnabled field if non-nil, zero value otherwise.

### GetOtherOptionEnabledOk

`func (o *FormField) GetOtherOptionEnabledOk() (*bool, bool)`

GetOtherOptionEnabledOk returns a tuple with the OtherOptionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionEnabled

`func (o *FormField) SetOtherOptionEnabled(v bool)`

SetOtherOptionEnabled sets OtherOptionEnabled field to given value.

### HasOtherOptionEnabled

`func (o *FormField) HasOtherOptionEnabled() bool`

HasOtherOptionEnabled returns a boolean if a field has been set.

### GetOtherOptionKey

`func (o *FormField) GetOtherOptionKey() string`

GetOtherOptionKey returns the OtherOptionKey field if non-nil, zero value otherwise.

### GetOtherOptionKeyOk

`func (o *FormField) GetOtherOptionKeyOk() (*string, bool)`

GetOtherOptionKeyOk returns a tuple with the OtherOptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionKey

`func (o *FormField) SetOtherOptionKey(v string)`

SetOtherOptionKey sets OtherOptionKey field to given value.

### HasOtherOptionKey

`func (o *FormField) HasOtherOptionKey() bool`

HasOtherOptionKey returns a boolean if a field has been set.

### GetOtherOptionlabel

`func (o *FormField) GetOtherOptionlabel() string`

GetOtherOptionlabel returns the OtherOptionlabel field if non-nil, zero value otherwise.

### GetOtherOptionlabelOk

`func (o *FormField) GetOtherOptionlabelOk() (*string, bool)`

GetOtherOptionlabelOk returns a tuple with the OtherOptionlabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionlabel

`func (o *FormField) SetOtherOptionlabel(v string)`

SetOtherOptionlabel sets OtherOptionlabel field to given value.

### HasOtherOptionlabel

`func (o *FormField) HasOtherOptionlabel() bool`

HasOtherOptionlabel returns a boolean if a field has been set.

### GetOtherOptionInputlabel

`func (o *FormField) GetOtherOptionInputlabel() string`

GetOtherOptionInputlabel returns the OtherOptionInputlabel field if non-nil, zero value otherwise.

### GetOtherOptionInputlabelOk

`func (o *FormField) GetOtherOptionInputlabelOk() (*string, bool)`

GetOtherOptionInputlabelOk returns a tuple with the OtherOptionInputlabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionInputlabel

`func (o *FormField) SetOtherOptionInputlabel(v string)`

SetOtherOptionInputlabel sets OtherOptionInputlabel field to given value.

### HasOtherOptionInputlabel

`func (o *FormField) HasOtherOptionInputlabel() bool`

HasOtherOptionInputlabel returns a boolean if a field has been set.

### GetOtherOptionAttributeDisabled

`func (o *FormField) GetOtherOptionAttributeDisabled() bool`

GetOtherOptionAttributeDisabled returns the OtherOptionAttributeDisabled field if non-nil, zero value otherwise.

### GetOtherOptionAttributeDisabledOk

`func (o *FormField) GetOtherOptionAttributeDisabledOk() (*bool, bool)`

GetOtherOptionAttributeDisabledOk returns a tuple with the OtherOptionAttributeDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionAttributeDisabled

`func (o *FormField) SetOtherOptionAttributeDisabled(v bool)`

SetOtherOptionAttributeDisabled sets OtherOptionAttributeDisabled field to given value.

### HasOtherOptionAttributeDisabled

`func (o *FormField) HasOtherOptionAttributeDisabled() bool`

HasOtherOptionAttributeDisabled returns a boolean if a field has been set.

### GetLayout

`func (o *FormField) GetLayout() EnumFormElementLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FormField) GetLayoutOk() (*EnumFormElementLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FormField) SetLayout(v EnumFormElementLayout)`

SetLayout sets Layout field to given value.


### GetOptions

`func (o *FormField) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormField) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormField) SetOptions(v []string)`

SetOptions sets Options field to given value.


### GetValidation

`func (o *FormField) GetValidation() FormElementValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *FormField) GetValidationOk() (*FormElementValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *FormField) SetValidation(v FormElementValidation)`

SetValidation sets Validation field to given value.


### GetLabelPasswordVerify

`func (o *FormField) GetLabelPasswordVerify() string`

GetLabelPasswordVerify returns the LabelPasswordVerify field if non-nil, zero value otherwise.

### GetLabelPasswordVerifyOk

`func (o *FormField) GetLabelPasswordVerifyOk() (*string, bool)`

GetLabelPasswordVerifyOk returns a tuple with the LabelPasswordVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPasswordVerify

`func (o *FormField) SetLabelPasswordVerify(v string)`

SetLabelPasswordVerify sets LabelPasswordVerify field to given value.

### HasLabelPasswordVerify

`func (o *FormField) HasLabelPasswordVerify() bool`

HasLabelPasswordVerify returns a boolean if a field has been set.

### GetValue

`func (o *FormField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FormField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FormField) SetValue(v string)`

SetValue sets Value field to given value.


### GetContent

`func (o *FormField) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FormField) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FormField) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FormField) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetStyles

`func (o *FormField) GetStyles() FormSocialLoginButtonStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormField) GetStylesOk() (*FormSocialLoginButtonStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormField) SetStyles(v FormSocialLoginButtonStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormField) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### GetSize

`func (o *FormField) GetSize() EnumFormRecaptchaV2Size`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FormField) GetSizeOk() (*EnumFormRecaptchaV2Size, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FormField) SetSize(v EnumFormRecaptchaV2Size)`

SetSize sets Size field to given value.


### GetTheme

`func (o *FormField) GetTheme() EnumFormRecaptchaV2Theme`

GetTheme returns the Theme field if non-nil, zero value otherwise.

### GetThemeOk

`func (o *FormField) GetThemeOk() (*EnumFormRecaptchaV2Theme, bool)`

GetThemeOk returns a tuple with the Theme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTheme

`func (o *FormField) SetTheme(v EnumFormRecaptchaV2Theme)`

SetTheme sets Theme field to given value.


### GetAlignment

`func (o *FormField) GetAlignment() EnumFormItemAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *FormField) GetAlignmentOk() (*EnumFormItemAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *FormField) SetAlignment(v EnumFormItemAlignment)`

SetAlignment sets Alignment field to given value.


### GetQrCodeType

`func (o *FormField) GetQrCodeType() EnumFormQrCodeType`

GetQrCodeType returns the QrCodeType field if non-nil, zero value otherwise.

### GetQrCodeTypeOk

`func (o *FormField) GetQrCodeTypeOk() (*EnumFormQrCodeType, bool)`

GetQrCodeTypeOk returns a tuple with the QrCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeType

`func (o *FormField) SetQrCodeType(v EnumFormQrCodeType)`

SetQrCodeType sets QrCodeType field to given value.


### GetShowBorder

`func (o *FormField) GetShowBorder() bool`

GetShowBorder returns the ShowBorder field if non-nil, zero value otherwise.

### GetShowBorderOk

`func (o *FormField) GetShowBorderOk() (*bool, bool)`

GetShowBorderOk returns a tuple with the ShowBorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBorder

`func (o *FormField) SetShowBorder(v bool)`

SetShowBorder sets ShowBorder field to given value.


### GetIdpType

`func (o *FormField) GetIdpType() EnumFormSocialLoginIdpType`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *FormField) GetIdpTypeOk() (*EnumFormSocialLoginIdpType, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *FormField) SetIdpType(v EnumFormSocialLoginIdpType)`

SetIdpType sets IdpType field to given value.


### GetIdpName

`func (o *FormField) GetIdpName() string`

GetIdpName returns the IdpName field if non-nil, zero value otherwise.

### GetIdpNameOk

`func (o *FormField) GetIdpNameOk() (*string, bool)`

GetIdpNameOk returns a tuple with the IdpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpName

`func (o *FormField) SetIdpName(v string)`

SetIdpName sets IdpName field to given value.


### GetIdpId

`func (o *FormField) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *FormField) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *FormField) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.


### GetIdpEnabled

`func (o *FormField) GetIdpEnabled() bool`

GetIdpEnabled returns the IdpEnabled field if non-nil, zero value otherwise.

### GetIdpEnabledOk

`func (o *FormField) GetIdpEnabledOk() (*bool, bool)`

GetIdpEnabledOk returns a tuple with the IdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEnabled

`func (o *FormField) SetIdpEnabled(v bool)`

SetIdpEnabled sets IdpEnabled field to given value.


### GetIconSrc

`func (o *FormField) GetIconSrc() string`

GetIconSrc returns the IconSrc field if non-nil, zero value otherwise.

### GetIconSrcOk

`func (o *FormField) GetIconSrcOk() (*string, bool)`

GetIconSrcOk returns a tuple with the IconSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSrc

`func (o *FormField) SetIconSrc(v string)`

SetIconSrc sets IconSrc field to given value.


### GetWidth

`func (o *FormField) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FormField) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FormField) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *FormField) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

