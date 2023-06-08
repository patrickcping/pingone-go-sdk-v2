# FormFieldEmptyField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Content** | Pointer to **string** | A string that specifies the field content. | [optional] 

## Methods

### NewFormFieldEmptyField

`func NewFormFieldEmptyField(type_ EnumFormFieldType, position FormFieldCommonPosition, ) *FormFieldEmptyField`

NewFormFieldEmptyField instantiates a new FormFieldEmptyField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldEmptyFieldWithDefaults

`func NewFormFieldEmptyFieldWithDefaults() *FormFieldEmptyField`

NewFormFieldEmptyFieldWithDefaults instantiates a new FormFieldEmptyField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldEmptyField) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldEmptyField) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldEmptyField) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldEmptyField) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldEmptyField) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldEmptyField) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetContent

`func (o *FormFieldEmptyField) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FormFieldEmptyField) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FormFieldEmptyField) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FormFieldEmptyField) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


