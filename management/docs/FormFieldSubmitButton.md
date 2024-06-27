# FormFieldSubmitButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Key** | Pointer to **string** | A string that specifies an identifier for the field component. | [optional] 
**Label** | **string** | A string that specifies the button label. | 
**Styles** | Pointer to [**FormStyles**](FormStyles.md) |  | [optional] 

## Methods

### NewFormFieldSubmitButton

`func NewFormFieldSubmitButton(type_ EnumFormFieldType, position FormFieldCommonPosition, label string, ) *FormFieldSubmitButton`

NewFormFieldSubmitButton instantiates a new FormFieldSubmitButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldSubmitButtonWithDefaults

`func NewFormFieldSubmitButtonWithDefaults() *FormFieldSubmitButton`

NewFormFieldSubmitButtonWithDefaults instantiates a new FormFieldSubmitButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldSubmitButton) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldSubmitButton) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldSubmitButton) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldSubmitButton) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldSubmitButton) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldSubmitButton) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetKey

`func (o *FormFieldSubmitButton) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldSubmitButton) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldSubmitButton) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormFieldSubmitButton) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *FormFieldSubmitButton) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldSubmitButton) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldSubmitButton) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStyles

`func (o *FormFieldSubmitButton) GetStyles() FormStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormFieldSubmitButton) GetStylesOk() (*FormStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormFieldSubmitButton) SetStyles(v FormStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormFieldSubmitButton) HasStyles() bool`

HasStyles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


