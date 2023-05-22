# FormFieldTextAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | Pointer to [**EnumFormElementLayout**](EnumFormElementLayout.md) |  | [optional] 
**Options** | Pointer to **[]string** | An array of strings that specifies the unique list of options. This is a required property when the type is &#x60;RADIO&#x60;, &#x60;CHECKBOX&#x60;, or &#x60;DROPDOWN&#x60;. | [optional] 
**Validation** | [**FormElementValidation**](FormElementValidation.md) |  | 

## Methods

### NewFormFieldTextAllOf

`func NewFormFieldTextAllOf(validation FormElementValidation, ) *FormFieldTextAllOf`

NewFormFieldTextAllOf instantiates a new FormFieldTextAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldTextAllOfWithDefaults

`func NewFormFieldTextAllOfWithDefaults() *FormFieldTextAllOf`

NewFormFieldTextAllOfWithDefaults instantiates a new FormFieldTextAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayout

`func (o *FormFieldTextAllOf) GetLayout() EnumFormElementLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FormFieldTextAllOf) GetLayoutOk() (*EnumFormElementLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FormFieldTextAllOf) SetLayout(v EnumFormElementLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *FormFieldTextAllOf) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetOptions

`func (o *FormFieldTextAllOf) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldTextAllOf) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldTextAllOf) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *FormFieldTextAllOf) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetValidation

`func (o *FormFieldTextAllOf) GetValidation() FormElementValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *FormFieldTextAllOf) GetValidationOk() (*FormElementValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *FormFieldTextAllOf) SetValidation(v FormElementValidation)`

SetValidation sets Validation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


