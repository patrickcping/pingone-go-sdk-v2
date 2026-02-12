# FormSingleCheckbox

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Appearance** | [**EnumFormSingleCheckboxAppearance**](EnumFormSingleCheckboxAppearance.md) |  | 
**ErrorMessage** | Pointer to **string** | A string that specifies the message to display if validation fails. | [optional] 
**InputType** | [**EnumFormSingleCheckboxInputType**](EnumFormSingleCheckboxInputType.md) |  | 

## Methods

### NewFormSingleCheckbox

`func NewFormSingleCheckbox(appearance EnumFormSingleCheckboxAppearance, inputType EnumFormSingleCheckboxInputType, ) *FormSingleCheckbox`

NewFormSingleCheckbox instantiates a new FormSingleCheckbox object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSingleCheckboxWithDefaults

`func NewFormSingleCheckboxWithDefaults() *FormSingleCheckbox`

NewFormSingleCheckboxWithDefaults instantiates a new FormSingleCheckbox object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppearance

`func (o *FormSingleCheckbox) GetAppearance() EnumFormSingleCheckboxAppearance`

GetAppearance returns the Appearance field if non-nil, zero value otherwise.

### GetAppearanceOk

`func (o *FormSingleCheckbox) GetAppearanceOk() (*EnumFormSingleCheckboxAppearance, bool)`

GetAppearanceOk returns a tuple with the Appearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppearance

`func (o *FormSingleCheckbox) SetAppearance(v EnumFormSingleCheckboxAppearance)`

SetAppearance sets Appearance field to given value.


### GetErrorMessage

`func (o *FormSingleCheckbox) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FormSingleCheckbox) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FormSingleCheckbox) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FormSingleCheckbox) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetInputType

`func (o *FormSingleCheckbox) GetInputType() EnumFormSingleCheckboxInputType`

GetInputType returns the InputType field if non-nil, zero value otherwise.

### GetInputTypeOk

`func (o *FormSingleCheckbox) GetInputTypeOk() (*EnumFormSingleCheckboxInputType, bool)`

GetInputTypeOk returns a tuple with the InputType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputType

`func (o *FormSingleCheckbox) SetInputType(v EnumFormSingleCheckboxInputType)`

SetInputType sets InputType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


