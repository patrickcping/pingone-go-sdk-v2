# FormSocialLoginButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A string that specifies the social login button label. | 
**Styles** | Pointer to [**FormSocialLoginButtonStyles**](FormSocialLoginButtonStyles.md) |  | [optional] 
**IdpType** | [**EnumFormSocialLoginIdpType**](EnumFormSocialLoginIdpType.md) |  | 
**IdpName** | **string** | A string that specifies the external identity provider name. | 
**IdpId** | **string** | A string that specifies the external identity provider&#39;s ID. | 
**IdpEnabled** | **bool** | A boolean that specifies whether the external identity provider is enabled. | 
**IconSrc** | **string** | A string that specifies the HTTP link (URL format) for the external identity provider&#39;s icon. | 
**Width** | Pointer to **int32** | An integer that specifies the button width. Set as a percentage. | [optional] 

## Methods

### NewFormSocialLoginButton

`func NewFormSocialLoginButton(label string, idpType EnumFormSocialLoginIdpType, idpName string, idpId string, idpEnabled bool, iconSrc string, ) *FormSocialLoginButton`

NewFormSocialLoginButton instantiates a new FormSocialLoginButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSocialLoginButtonWithDefaults

`func NewFormSocialLoginButtonWithDefaults() *FormSocialLoginButton`

NewFormSocialLoginButtonWithDefaults instantiates a new FormSocialLoginButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *FormSocialLoginButton) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormSocialLoginButton) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormSocialLoginButton) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStyles

`func (o *FormSocialLoginButton) GetStyles() FormSocialLoginButtonStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormSocialLoginButton) GetStylesOk() (*FormSocialLoginButtonStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormSocialLoginButton) SetStyles(v FormSocialLoginButtonStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormSocialLoginButton) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### GetIdpType

`func (o *FormSocialLoginButton) GetIdpType() EnumFormSocialLoginIdpType`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *FormSocialLoginButton) GetIdpTypeOk() (*EnumFormSocialLoginIdpType, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *FormSocialLoginButton) SetIdpType(v EnumFormSocialLoginIdpType)`

SetIdpType sets IdpType field to given value.


### GetIdpName

`func (o *FormSocialLoginButton) GetIdpName() string`

GetIdpName returns the IdpName field if non-nil, zero value otherwise.

### GetIdpNameOk

`func (o *FormSocialLoginButton) GetIdpNameOk() (*string, bool)`

GetIdpNameOk returns a tuple with the IdpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpName

`func (o *FormSocialLoginButton) SetIdpName(v string)`

SetIdpName sets IdpName field to given value.


### GetIdpId

`func (o *FormSocialLoginButton) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *FormSocialLoginButton) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *FormSocialLoginButton) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.


### GetIdpEnabled

`func (o *FormSocialLoginButton) GetIdpEnabled() bool`

GetIdpEnabled returns the IdpEnabled field if non-nil, zero value otherwise.

### GetIdpEnabledOk

`func (o *FormSocialLoginButton) GetIdpEnabledOk() (*bool, bool)`

GetIdpEnabledOk returns a tuple with the IdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEnabled

`func (o *FormSocialLoginButton) SetIdpEnabled(v bool)`

SetIdpEnabled sets IdpEnabled field to given value.


### GetIconSrc

`func (o *FormSocialLoginButton) GetIconSrc() string`

GetIconSrc returns the IconSrc field if non-nil, zero value otherwise.

### GetIconSrcOk

`func (o *FormSocialLoginButton) GetIconSrcOk() (*string, bool)`

GetIconSrcOk returns a tuple with the IconSrc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconSrc

`func (o *FormSocialLoginButton) SetIconSrc(v string)`

SetIconSrc sets IconSrc field to given value.


### GetWidth

`func (o *FormSocialLoginButton) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FormSocialLoginButton) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FormSocialLoginButton) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *FormSocialLoginButton) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


