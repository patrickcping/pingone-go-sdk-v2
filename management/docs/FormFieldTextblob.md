# FormFieldTextblob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**Content** | Pointer to **string** | A string that specifies the field content. | [optional] 

## Methods

### NewFormFieldTextblob

`func NewFormFieldTextblob(type_ EnumFormFieldType, position FormFieldCommonPosition, ) *FormFieldTextblob`

NewFormFieldTextblob instantiates a new FormFieldTextblob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldTextblobWithDefaults

`func NewFormFieldTextblobWithDefaults() *FormFieldTextblob`

NewFormFieldTextblobWithDefaults instantiates a new FormFieldTextblob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldTextblob) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldTextblob) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldTextblob) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldTextblob) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldTextblob) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldTextblob) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldTextblob) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldTextblob) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldTextblob) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldTextblob) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetContent

`func (o *FormFieldTextblob) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *FormFieldTextblob) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *FormFieldTextblob) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *FormFieldTextblob) HasContent() bool`

HasContent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


