# FormFieldFormAgreement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**InputType** | [**EnumFormAgreementInputType**](EnumFormAgreementInputType.md) |  | 
**TitleEnabled** | **bool** | Specifies whether the title is enabled. | 
**Agreement** | Pointer to [**FormAgreementAgreement**](FormAgreementAgreement.md) |  | [optional] 

## Methods

### NewFormFieldFormAgreement

`func NewFormFieldFormAgreement(type_ EnumFormFieldType, position FormFieldCommonPosition, inputType EnumFormAgreementInputType, titleEnabled bool, ) *FormFieldFormAgreement`

NewFormFieldFormAgreement instantiates a new FormFieldFormAgreement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldFormAgreementWithDefaults

`func NewFormFieldFormAgreementWithDefaults() *FormFieldFormAgreement`

NewFormFieldFormAgreementWithDefaults instantiates a new FormFieldFormAgreement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldFormAgreement) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldFormAgreement) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldFormAgreement) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldFormAgreement) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldFormAgreement) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldFormAgreement) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldFormAgreement) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldFormAgreement) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldFormAgreement) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldFormAgreement) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetInputType

`func (o *FormFieldFormAgreement) GetInputType() EnumFormAgreementInputType`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *FormFieldFormAgreement) GetInputTypeOk() (*EnumFormAgreementInputType, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *FormFieldFormAgreement) SetInputType(v EnumFormAgreementInputType)`

SetInputType sets InputType field to given value.


### GetTitleEnabled

`func (o *FormFieldFormAgreement) GetTitleEnabled() bool`

GetTitleEnabled returns the TitleEnabled field if non-nil, zero value otherwise.

### GetTitleEnabledOk

`func (o *FormFieldFormAgreement) GetTitleEnabledOk() (*bool, bool)`

GetTitleEnabledOk returns a tuple with the TitleEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleEnabled

`func (o *FormFieldFormAgreement) SetTitleEnabled(v bool)`

SetTitleEnabled sets TitleEnabled field to given value.


### GetAgreement

`func (o *FormFieldFormAgreement) GetAgreement() FormAgreementAgreement`

GetAgreement returns the Agreement field if non-nil, zero value otherwise.

### GetAgreementOk

`func (o *FormFieldFormAgreement) GetAgreementOk() (*FormAgreementAgreement, bool)`

GetAgreementOk returns a tuple with the Agreement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgreement

`func (o *FormFieldFormAgreement) SetAgreement(v FormAgreementAgreement)`

SetAgreement sets Agreement field to given value.

### HasAgreement

`func (o *FormFieldFormAgreement) HasAgreement() bool`

HasAgreement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


