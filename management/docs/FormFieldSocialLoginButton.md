# FormFieldSocialLoginButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Key** | Pointer to **string** | A string that specifies an identifier for the field component. | [optional] 
**Label** | **string** | A string that specifies the social login button label. | 
**Styles** | Pointer to [**FormStyles**](FormStyles.md) |  | [optional] 
**IdpType** | [**EnumFormSocialLoginIdpType**](EnumFormSocialLoginIdpType.md) |  | 
**IdpName** | **string** | A string that specifies the external identity provider name. | 
**IdpId** | **string** | A string that specifies the external identity provider&#39;s ID. | 
**IdpEnabled** | **bool** | A boolean that specifies whether the external identity provider is enabled. | 

## Methods

### NewFormFieldSocialLoginButton

`func NewFormFieldSocialLoginButton(type_ EnumFormFieldType, position FormFieldCommonPosition, label string, idpType EnumFormSocialLoginIdpType, idpName string, idpId string, idpEnabled bool, ) *FormFieldSocialLoginButton`

NewFormFieldSocialLoginButton instantiates a new FormFieldSocialLoginButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldSocialLoginButtonWithDefaults

`func NewFormFieldSocialLoginButtonWithDefaults() *FormFieldSocialLoginButton`

NewFormFieldSocialLoginButtonWithDefaults instantiates a new FormFieldSocialLoginButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldSocialLoginButton) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldSocialLoginButton) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldSocialLoginButton) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldSocialLoginButton) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldSocialLoginButton) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldSocialLoginButton) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetKey

`func (o *FormFieldSocialLoginButton) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldSocialLoginButton) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldSocialLoginButton) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormFieldSocialLoginButton) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *FormFieldSocialLoginButton) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldSocialLoginButton) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldSocialLoginButton) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStyles

`func (o *FormFieldSocialLoginButton) GetStyles() FormStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormFieldSocialLoginButton) GetStylesOk() (*FormStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormFieldSocialLoginButton) SetStyles(v FormStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormFieldSocialLoginButton) HasStyles() bool`

HasStyles returns a boolean if a field has been set.

### GetIdpType

`func (o *FormFieldSocialLoginButton) GetIdpType() EnumFormSocialLoginIdpType`

GetIdpType returns the IdpType field if non-nil, zero value otherwise.

### GetIdpTypeOk

`func (o *FormFieldSocialLoginButton) GetIdpTypeOk() (*EnumFormSocialLoginIdpType, bool)`

GetIdpTypeOk returns a tuple with the IdpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpType

`func (o *FormFieldSocialLoginButton) SetIdpType(v EnumFormSocialLoginIdpType)`

SetIdpType sets IdpType field to given value.


### GetIdpName

`func (o *FormFieldSocialLoginButton) GetIdpName() string`

GetIdpName returns the IdpName field if non-nil, zero value otherwise.

### GetIdpNameOk

`func (o *FormFieldSocialLoginButton) GetIdpNameOk() (*string, bool)`

GetIdpNameOk returns a tuple with the IdpName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpName

`func (o *FormFieldSocialLoginButton) SetIdpName(v string)`

SetIdpName sets IdpName field to given value.


### GetIdpId

`func (o *FormFieldSocialLoginButton) GetIdpId() string`

GetIdpId returns the IdpId field if non-nil, zero value otherwise.

### GetIdpIdOk

`func (o *FormFieldSocialLoginButton) GetIdpIdOk() (*string, bool)`

GetIdpIdOk returns a tuple with the IdpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpId

`func (o *FormFieldSocialLoginButton) SetIdpId(v string)`

SetIdpId sets IdpId field to given value.


### GetIdpEnabled

`func (o *FormFieldSocialLoginButton) GetIdpEnabled() bool`

GetIdpEnabled returns the IdpEnabled field if non-nil, zero value otherwise.

### GetIdpEnabledOk

`func (o *FormFieldSocialLoginButton) GetIdpEnabledOk() (*bool, bool)`

GetIdpEnabledOk returns a tuple with the IdpEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdpEnabled

`func (o *FormFieldSocialLoginButton) SetIdpEnabled(v bool)`

SetIdpEnabled sets IdpEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


