# FormElementValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Regex** | Pointer to **string** | A string that specifies a validation regular expression. The expression must be a valid regular expression string. This is a required property when the validation type is &#x60;CUSTOM&#x60;. | [optional] 
**Type** | Pointer to [**EnumFormElementValidationType**](EnumFormElementValidationType.md) |  | [optional] 
**ErrorMessage** | Pointer to **string** | A string that specifies the error message to be displayed when the field validation fails. | [optional] 

## Methods

### NewFormElementValidation

`func NewFormElementValidation() *FormElementValidation`

NewFormElementValidation instantiates a new FormElementValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormElementValidationWithDefaults

`func NewFormElementValidationWithDefaults() *FormElementValidation`

NewFormElementValidationWithDefaults instantiates a new FormElementValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegex

`func (o *FormElementValidation) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *FormElementValidation) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *FormElementValidation) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *FormElementValidation) HasRegex() bool`

HasRegex returns a boolean if a field has been set.

### GetType

`func (o *FormElementValidation) GetType() EnumFormElementValidationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormElementValidation) GetTypeOk() (*EnumFormElementValidationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormElementValidation) SetType(v EnumFormElementValidationType)`

SetType sets Type field to given value.

### HasType

`func (o *FormElementValidation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetErrorMessage

`func (o *FormElementValidation) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *FormElementValidation) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *FormElementValidation) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *FormElementValidation) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


