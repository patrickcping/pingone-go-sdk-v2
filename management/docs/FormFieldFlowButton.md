# FormFieldFlowButton

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**Key** | **string** | A string that specifies an identifier for the field component. | 
**Label** | **string** | A string that specifies the button label. | 
**Styles** | Pointer to [**FormStyles**](FormStyles.md) |  | [optional] 

## Methods

### NewFormFieldFlowButton

`func NewFormFieldFlowButton(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, ) *FormFieldFlowButton`

NewFormFieldFlowButton instantiates a new FormFieldFlowButton object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldFlowButtonWithDefaults

`func NewFormFieldFlowButtonWithDefaults() *FormFieldFlowButton`

NewFormFieldFlowButtonWithDefaults instantiates a new FormFieldFlowButton object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldFlowButton) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldFlowButton) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldFlowButton) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldFlowButton) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldFlowButton) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldFlowButton) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldFlowButton) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldFlowButton) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldFlowButton) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldFlowButton) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetKey

`func (o *FormFieldFlowButton) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldFlowButton) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldFlowButton) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormFieldFlowButton) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldFlowButton) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldFlowButton) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStyles

`func (o *FormFieldFlowButton) GetStyles() FormStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormFieldFlowButton) GetStylesOk() (*FormStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormFieldFlowButton) SetStyles(v FormStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormFieldFlowButton) HasStyles() bool`

HasStyles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


