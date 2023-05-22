# FormFieldRadioAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | [**EnumFormElementLayout**](EnumFormElementLayout.md) |  | 
**Options** | **[]string** | An array of strings that specifies the unique list of options. This is a required property when the type is &#x60;RADIO&#x60;, &#x60;CHECKBOX&#x60;, or &#x60;DROPDOWN&#x60;. | 
**Validation** | Pointer to [**FormElementValidation**](FormElementValidation.md) |  | [optional] 

## Methods

### NewFormFieldRadioAllOf

`func NewFormFieldRadioAllOf(layout EnumFormElementLayout, options []string, ) *FormFieldRadioAllOf`

NewFormFieldRadioAllOf instantiates a new FormFieldRadioAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldRadioAllOfWithDefaults

`func NewFormFieldRadioAllOfWithDefaults() *FormFieldRadioAllOf`

NewFormFieldRadioAllOfWithDefaults instantiates a new FormFieldRadioAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayout

`func (o *FormFieldRadioAllOf) GetLayout() EnumFormElementLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *FormFieldRadioAllOf) GetLayoutOk() (*EnumFormElementLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *FormFieldRadioAllOf) SetLayout(v EnumFormElementLayout)`

SetLayout sets Layout field to given value.


### GetOptions

`func (o *FormFieldRadioAllOf) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldRadioAllOf) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldRadioAllOf) SetOptions(v []string)`

SetOptions sets Options field to given value.


### GetValidation

`func (o *FormFieldRadioAllOf) GetValidation() FormElementValidation`

GetValidation returns the Validation field if non-nil, zero value otherwise.

### GetValidationOk

`func (o *FormFieldRadioAllOf) GetValidationOk() (*FormElementValidation, bool)`

GetValidationOk returns a tuple with the Validation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidation

`func (o *FormFieldRadioAllOf) SetValidation(v FormElementValidation)`

SetValidation sets Validation field to given value.

### HasValidation

`func (o *FormFieldRadioAllOf) HasValidation() bool`

HasValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


