# FormFieldSingleCheckbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**Appearance** | [**EnumFormSingleCheckboxAppearance**](EnumFormSingleCheckboxAppearance.md) |  | 
**ErrorMessage** | Pointer to **string** | A string that specifies the message to display if validation fails. | [optional] 
**InputType** | [**EnumFormSingleCheckboxInputType**](EnumFormSingleCheckboxInputType.md) |  | 

## Methods

### NewFormFieldSingleCheckbox

`func NewFormFieldSingleCheckbox(type_ EnumFormFieldType, position FormFieldCommonPosition, appearance EnumFormSingleCheckboxAppearance, inputType EnumFormSingleCheckboxInputType, ) *FormFieldSingleCheckbox`

NewFormFieldSingleCheckbox instantiates a new FormFieldSingleCheckbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldSingleCheckboxWithDefaults

`func NewFormFieldSingleCheckboxWithDefaults() *FormFieldSingleCheckbox`

NewFormFieldSingleCheckboxWithDefaults instantiates a new FormFieldSingleCheckbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldSingleCheckbox) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldSingleCheckbox) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldSingleCheckbox) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldSingleCheckbox) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldSingleCheckbox) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldSingleCheckbox) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldSingleCheckbox) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldSingleCheckbox) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldSingleCheckbox) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldSingleCheckbox) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAppearance

`func (o *FormFieldSingleCheckbox) GetAppearance() EnumFormSingleCheckboxAppearance`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *FormFieldSingleCheckbox) GetAppearanceOk() (*EnumFormSingleCheckboxAppearance, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *FormFieldSingleCheckbox) SetAppearance(v EnumFormSingleCheckboxAppearance)`

SetAppearance sets Appearance field to given value.


### GetErrorMessage

`func (o *FormFieldSingleCheckbox) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FormFieldSingleCheckbox) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FormFieldSingleCheckbox) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FormFieldSingleCheckbox) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetInputType

`func (o *FormFieldSingleCheckbox) GetInputType() EnumFormSingleCheckboxInputType`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *FormFieldSingleCheckbox) GetInputTypeOk() (*EnumFormSingleCheckboxInputType, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *FormFieldSingleCheckbox) SetInputType(v EnumFormSingleCheckboxInputType)`

SetInputType sets InputType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


