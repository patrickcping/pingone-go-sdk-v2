# FormField

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
**Layout** | [**EnumFormElementLayout**](EnumFormElementLayout.md) |  | 
**Options** | [**[]AuthenticationDevice**](AuthenticationDevice.md) | List of devices available for authentication. Must not be empty. | 
**Validation** | [**FormElementValidation**](FormElementValidation.md) |  | 
**ShowPasswordRequirements** | Pointer to **bool** | A boolean that specifies whether the password requirements are displayed. | [optional] 
**LabelPasswordVerify** | Pointer to **string** | A string that when a second field for verifies password is used, this poperty specifies the field label for that verify field. | [optional] 
**DefaultCountryCode** | Pointer to **string** | The country code to default the country code selection field to (two character country code). | [optional] 
**CountryCodeLabel** | Pointer to **string** | Label for the country code field. | [optional] 
**ValidatePhoneNumber** | Pointer to **bool** | Whether to validate the phone number input. | [optional] 
**ShowExtension** | Pointer to **bool** | Whether to show an extension field. | [optional] 
**ExtensionLabel** | Pointer to **string** | Label for the extension field. This is required and must not be blank if showExtension is true. | [optional] 
**Content** | Pointer to **string** | A string that specifies the field content. | [optional] 
**Icon** | Pointer to [**FormItemWithIconAllOfIcon**](FormItemWithIconAllOfIcon.md) |  | [optional] 
**Styles** | Pointer to [**FormSocialLoginButtonAllOfStyles**](FormSocialLoginButtonAllOfStyles.md) |  | [optional] 
**Size** | [**EnumFormItemSize**](EnumFormItemSize.md) |  | 
**Theme** | [**EnumFormRecaptchaV2Theme**](EnumFormRecaptchaV2Theme.md) |  | 
**Alignment** | [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | 
**FallbackText** | Pointer to **string** | A string that specifies the text label for fallback under the QR code. | [optional] 
**IdpName** | **string** | A string that specifies the external identity provider name. | 
**IdpType** | [**EnumFormSocialLoginIdpType**](EnumFormSocialLoginIdpType.md) |  | 
**IdpId** | **string** | A string that specifies the external identity provider&#39;s ID. | 
**IdpEnabled** | Pointer to **bool** | A boolean that specifies whether the external identity provider is enabled. | [optional] 
**IconSrc** | Pointer to **string** | A string that specifies the icon image URL to be displayed on the button. | [optional] 
**PollingAppearance** | [**EnumFormPollingAppearance**](EnumFormPollingAppearance.md) |  | 
**Trigger** | [**EnumFormFIDO2Trigger**](EnumFormFIDO2Trigger.md) |  | 
**Action** | [**EnumFormFIDO2Action**](EnumFormFIDO2Action.md) |  | 
**Appearance** | [**EnumFormSingleCheckboxAppearance**](EnumFormSingleCheckboxAppearance.md) |  | 
**ErrorMessage** | Pointer to **string** | A string that specifies the message to display if validation fails. | [optional] 
**InputType** | [**EnumFormAgreementInputType**](EnumFormAgreementInputType.md) |  | 
**TitleEnabled** | **bool** | Specifies whether the title is enabled. | 
**Agreement** | Pointer to [**FormAgreementAllOfAgreement**](FormAgreementAllOfAgreement.md) |  | [optional] 

## Methods

### NewFormField

`func NewFormField(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, layout EnumFormElementLayout, options []AuthenticationDevice, validation FormElementValidation, size EnumFormItemSize, theme EnumFormRecaptchaV2Theme, alignment EnumFormItemAlignment, idpName string, idpType EnumFormSocialLoginIdpType, idpId string, pollingAppearance EnumFormPollingAppearance, trigger EnumFormFIDO2Trigger, action EnumFormFIDO2Action, appearance EnumFormSingleCheckboxAppearance, inputType EnumFormAgreementInputType, titleEnabled bool, ) *FormField`

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


### GetVisibility

`func (o *FormField) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormField) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormField) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormField) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

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

### HasRequired

`func (o *FormField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

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

### GetOtherOptionLabel

`func (o *FormField) GetOtherOptionLabel() string`

GetOtherOptionLabel returns the OtherOptionLabel field if non-nil, zero value otherwise.

### GetOtherOptionLabelOk

`func (o *FormField) GetOtherOptionLabelOk() (*string, bool)`

GetOtherOptionLabelOk returns a tuple with the OtherOptionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionLabel

`func (o *FormField) SetOtherOptionLabel(v string)`

SetOtherOptionLabel sets OtherOptionLabel field to given value.

### HasOtherOptionLabel

`func (o *FormField) HasOtherOptionLabel() bool`

HasOtherOptionLabel returns a boolean if a field has been set.

### GetOtherOptionInputLabel

`func (o *FormField) GetOtherOptionInputLabel() string`

GetOtherOptionInputLabel returns the OtherOptionInputLabel field if non-nil, zero value otherwise.

### GetOtherOptionInputLabelOk

`func (o *FormField) GetOtherOptionInputLabelOk() (*string, bool)`

GetOtherOptionInputLabelOk returns a tuple with the OtherOptionInputLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherOptionInputLabel

`func (o *FormField) SetOtherOptionInputLabel(v string)`

SetOtherOptionInputLabel sets OtherOptionInputLabel field to given value.

### HasOtherOptionInputLabel

`func (o *FormField) HasOtherOptionInputLabel() bool`

HasOtherOptionInputLabel returns a boolean if a field has been set.

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

`func (o *FormField) GetOptions() []AuthenticationDevice`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormField) GetOptionsOk() (*[]AuthenticationDevice, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormField) SetOptions(v []AuthenticationDevice)`

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


### GetShowPasswordRequirements

`func (o *FormField) GetShowPasswordRequirements() bool`

GetShowPasswordRequirements returns the ShowPasswordRequirements field if non-nil, zero value otherwise.

### GetShowPasswordRequirementsOk

`func (o *FormField) GetShowPasswordRequirementsOk() (*bool, bool)`

GetShowPasswordRequirementsOk returns a tuple with the ShowPasswordRequirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowPasswordRequirements

`func (o *FormField) SetShowPasswordRequirements(v bool)`

SetShowPasswordRequirements sets ShowPasswordRequirements field to given value.

### HasShowPasswordRequirements

`func (o *FormField) HasShowPasswordRequirements() bool`

HasShowPasswordRequirements returns a boolean if a field has been set.

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

### GetDefaultCountryCode

`func (o *FormField) GetDefaultCountryCode() string`

GetDefaultCountryCode returns the DefaultCountryCode field if non-nil, zero value otherwise.

### GetDefaultCountryCodeOk

`func (o *FormField) GetDefaultCountryCodeOk() (*string, bool)`

GetDefaultCountryCodeOk returns a tuple with the DefaultCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCountryCode

`func (o *FormField) SetDefaultCountryCode(v string)`

SetDefaultCountryCode sets DefaultCountryCode field to given value.

### HasDefaultCountryCode

`func (o *FormField) HasDefaultCountryCode() bool`

HasDefaultCountryCode returns a boolean if a field has been set.

### GetCountryCodeLabel

`func (o *FormField) GetCountryCodeLabel() string`

GetCountryCodeLabel returns the CountryCodeLabel field if non-nil, zero value otherwise.

### GetCountryCodeLabelOk

`func (o *FormField) GetCountryCodeLabelOk() (*string, bool)`

GetCountryCodeLabelOk returns a tuple with the CountryCodeLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCodeLabel

`func (o *FormField) SetCountryCodeLabel(v string)`

SetCountryCodeLabel sets CountryCodeLabel field to given value.

### HasCountryCodeLabel

`func (o *FormField) HasCountryCodeLabel() bool`

HasCountryCodeLabel returns a boolean if a field has been set.

### GetValidatePhoneNumber

`func (o *FormField) GetValidatePhoneNumber() bool`

GetValidatePhoneNumber returns the ValidatePhoneNumber field if non-nil, zero value otherwise.

### GetValidatePhoneNumberOk

`func (o *FormField) GetValidatePhoneNumberOk() (*bool, bool)`

GetValidatePhoneNumberOk returns a tuple with the ValidatePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidatePhoneNumber

`func (o *FormField) SetValidatePhoneNumber(v bool)`

SetValidatePhoneNumber sets ValidatePhoneNumber field to given value.

### HasValidatePhoneNumber

`func (o *FormField) HasValidatePhoneNumber() bool`

HasValidatePhoneNumber returns a boolean if a field has been set.

### GetShowExtension

`func (o *FormField) GetShowExtension() bool`

GetShowExtension returns the ShowExtension field if non-nil, zero value otherwise.

### GetShowExtensionOk

`func (o *FormField) GetShowExtensionOk() (*bool, bool)`

GetShowExtensionOk returns a tuple with the ShowExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowExtension

`func (o *FormField) SetShowExtension(v bool)`

SetShowExtension sets ShowExtension field to given value.

### HasShowExtension

`func (o *FormField) HasShowExtension() bool`

HasShowExtension returns a boolean if a field has been set.

### GetExtensionLabel

`func (o *FormField) GetExtensionLabel() string`

GetExtensionLabel returns the ExtensionLabel field if non-nil, zero value otherwise.

### GetExtensionLabelOk

`func (o *FormField) GetExtensionLabelOk() (*string, bool)`

GetExtensionLabelOk returns a tuple with the ExtensionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionLabel

`func (o *FormField) SetExtensionLabel(v string)`

SetExtensionLabel sets ExtensionLabel field to given value.

### HasExtensionLabel

`func (o *FormField) HasExtensionLabel() bool`

HasExtensionLabel returns a boolean if a field has been set.

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

### GetIcon

`func (o *FormField) GetIcon() FormItemWithIconAllOfIcon`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *FormField) GetIconOk() (*FormItemWithIconAllOfIcon, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *FormField) SetIcon(v FormItemWithIconAllOfIcon)`

SetIcon sets Icon field to given value.

### HasIcon

`func (o *FormField) HasIcon() bool`

HasIcon returns a boolean if a field has been set.

### GetStyles

`func (o *FormField) GetStyles() FormSocialLoginButtonAllOfStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormField) GetStylesOk() (*FormSocialLoginButtonAllOfStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormField) SetStyles(v FormSocialLoginButtonAllOfStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormField) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### GetSize

`func (o *FormField) GetSize() EnumFormItemSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FormField) GetSizeOk() (*EnumFormItemSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FormField) SetSize(v EnumFormItemSize)`

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


### GetFallbackText

`func (o *FormField) GetFallbackText() string`

GetFallbackText returns the FallbackText field if non-nil, zero value otherwise.

### GetFallbackTextOk

`func (o *FormField) GetFallbackTextOk() (*string, bool)`

GetFallbackTextOk returns a tuple with the FallbackText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackText

`func (o *FormField) SetFallbackText(v string)`

SetFallbackText sets FallbackText field to given value.

### HasFallbackText

`func (o *FormField) HasFallbackText() bool`

HasFallbackText returns a boolean if a field has been set.

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

### HasIdpEnabled

`func (o *FormField) HasIdpEnabled() bool`

HasIdpEnabled returns a boolean if a field has been set.

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

### HasIconSrc

`func (o *FormField) HasIconSrc() bool`

HasIconSrc returns a boolean if a field has been set.

### GetPollingAppearance

`func (o *FormField) GetPollingAppearance() EnumFormPollingAppearance`

GetPollingAppearance returns the PollingAppearance field if non-nil, zero value otherwise.

### GetPollingAppearanceOk

`func (o *FormField) GetPollingAppearanceOk() (*EnumFormPollingAppearance, bool)`

GetPollingAppearanceOk returns a tuple with the PollingAppearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingAppearance

`func (o *FormField) SetPollingAppearance(v EnumFormPollingAppearance)`

SetPollingAppearance sets PollingAppearance field to given value.


### GetTrigger

`func (o *FormField) GetTrigger() EnumFormFIDO2Trigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *FormField) GetTriggerOk() (*EnumFormFIDO2Trigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *FormField) SetTrigger(v EnumFormFIDO2Trigger)`

SetTrigger sets Trigger field to given value.


### GetAction

`func (o *FormField) GetAction() EnumFormFIDO2Action`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *FormField) GetActionOk() (*EnumFormFIDO2Action, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *FormField) SetAction(v EnumFormFIDO2Action)`

SetAction sets Action field to given value.


### GetAppearance

`func (o *FormField) GetAppearance() EnumFormSingleCheckboxAppearance`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *FormField) GetAppearanceOk() (*EnumFormSingleCheckboxAppearance, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *FormField) SetAppearance(v EnumFormSingleCheckboxAppearance)`

SetAppearance sets Appearance field to given value.


### GetErrorMessage

`func (o *FormField) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FormField) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FormField) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FormField) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetInputType

`func (o *FormField) GetInputType() EnumFormAgreementInputType`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *FormField) GetInputTypeOk() (*EnumFormAgreementInputType, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *FormField) SetInputType(v EnumFormAgreementInputType)`

SetInputType sets InputType field to given value.


### GetTitleEnabled

`func (o *FormField) GetTitleEnabled() bool`

GetTitleEnabled returns the TitleEnabled field if non-nil, zero value otherwise.

### GetTitleEnabledOk

`func (o *FormField) GetTitleEnabledOk() (*bool, bool)`

GetTitleEnabledOk returns a tuple with the TitleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleEnabled

`func (o *FormField) SetTitleEnabled(v bool)`

SetTitleEnabled sets TitleEnabled field to given value.


### GetAgreement

`func (o *FormField) GetAgreement() FormAgreementAllOfAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *FormField) GetAgreementOk() (*FormAgreementAllOfAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *FormField) SetAgreement(v FormAgreementAllOfAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *FormField) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


