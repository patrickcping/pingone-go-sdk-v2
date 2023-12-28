# FormSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A string that specifies an identifier for the field component. | [optional] 
**Label** | **string** | A string that specifies the button label. | 
**Styles** | Pointer to [**FormStyles**](FormStyles.md) |  | [optional] 

## Methods

### NewFormSubmit

`func NewFormSubmit(label string, ) *FormSubmit`

NewFormSubmit instantiates a new FormSubmit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormSubmitWithDefaults

`func NewFormSubmitWithDefaults() *FormSubmit`

NewFormSubmitWithDefaults instantiates a new FormSubmit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FormSubmit) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormSubmit) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormSubmit) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FormSubmit) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLabel

`func (o *FormSubmit) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormSubmit) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormSubmit) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStyles

`func (o *FormSubmit) GetStyles() FormStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormSubmit) GetStylesOk() (*FormStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormSubmit) SetStyles(v FormStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormSubmit) HasStyles() bool`

HasStyles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


