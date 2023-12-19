# FormFieldCommonPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Row** | **int32** | An integer that specifies the number of rows (maximum number is 50). | 
**Col** | **int32** | An integer that specifies the number of columns (min &#x3D; 0; max &#x3D; 4). | 
**Width** | Pointer to **int32** | An integer that specifies the width of the field (in percentage). | [optional] 

## Methods

### NewFormFieldCommonPosition

`func NewFormFieldCommonPosition(row int32, col int32, ) *FormFieldCommonPosition`

NewFormFieldCommonPosition instantiates a new FormFieldCommonPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldCommonPositionWithDefaults

`func NewFormFieldCommonPositionWithDefaults() *FormFieldCommonPosition`

NewFormFieldCommonPositionWithDefaults instantiates a new FormFieldCommonPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRow

`func (o *FormFieldCommonPosition) GetRow() int32`

GetRow returns the Row field if non-nil, zero value otherwise.

### GetRowOk

`func (o *FormFieldCommonPosition) GetRowOk() (*int32, bool)`

GetRowOk returns a tuple with the Row field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRow

`func (o *FormFieldCommonPosition) SetRow(v int32)`

SetRow sets Row field to given value.


### GetCol

`func (o *FormFieldCommonPosition) GetCol() int32`

GetCol returns the Col field if non-nil, zero value otherwise.

### GetColOk

`func (o *FormFieldCommonPosition) GetColOk() (*int32, bool)`

GetColOk returns a tuple with the Col field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCol

`func (o *FormFieldCommonPosition) SetCol(v int32)`

SetCol sets Col field to given value.


### GetWidth

`func (o *FormFieldCommonPosition) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *FormFieldCommonPosition) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *FormFieldCommonPosition) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *FormFieldCommonPosition) HasWidth() bool`

HasWidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


