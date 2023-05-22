# FormFieldPasswordVerify

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**LabelPasswordVerify** | Pointer to **string** | A string that when a second field for verifies password is used, this poperty specifies the field label for that verify field. | [optional] 

## Methods

### NewFormFieldPasswordVerify

`func NewFormFieldPasswordVerify(type_ EnumFormFieldType, position FormFieldCommonPosition, ) *FormFieldPasswordVerify`

NewFormFieldPasswordVerify instantiates a new FormFieldPasswordVerify object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldPasswordVerifyWithDefaults

`func NewFormFieldPasswordVerifyWithDefaults() *FormFieldPasswordVerify`

NewFormFieldPasswordVerifyWithDefaults instantiates a new FormFieldPasswordVerify object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldPasswordVerify) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldPasswordVerify) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldPasswordVerify) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldPasswordVerify) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldPasswordVerify) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldPasswordVerify) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetLabelPasswordVerify

`func (o *FormFieldPasswordVerify) GetLabelPasswordVerify() string`

GetLabelPasswordVerify returns the LabelPasswordVerify field if non-nil, zero value otherwise.

### GetLabelPasswordVerifyOk

`func (o *FormFieldPasswordVerify) GetLabelPasswordVerifyOk() (*string, bool)`

GetLabelPasswordVerifyOk returns a tuple with the LabelPasswordVerify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPasswordVerify

`func (o *FormFieldPasswordVerify) SetLabelPasswordVerify(v string)`

SetLabelPasswordVerify sets LabelPasswordVerify field to given value.

### HasLabelPasswordVerify

`func (o *FormFieldPasswordVerify) HasLabelPasswordVerify() bool`

HasLabelPasswordVerify returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


