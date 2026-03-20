# FormSubmit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | A string that specifies the button label. | 
**Styles** | Pointer to [**FormFlowButtonStyles**](FormFlowButtonStyles.md) |  | [optional] 

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

`func (o *FormSubmit) GetStyles() FormFlowButtonStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormSubmit) GetStylesOk() (*FormFlowButtonStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormSubmit) SetStyles(v FormFlowButtonStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormSubmit) HasStyles() bool`

HasStyles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


