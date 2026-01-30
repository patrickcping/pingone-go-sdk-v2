# FormFieldFlowLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**Key** | **string** | A string that specifies an identifier for the field component. | 
**Label** | **string** | A string that specifies the link label. | 
**Styles** | Pointer to [**FormFlowLinkStyles**](FormFlowLinkStyles.md) |  | [optional] 

## Methods

### NewFormFieldFlowLink

`func NewFormFieldFlowLink(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, ) *FormFieldFlowLink`

NewFormFieldFlowLink instantiates a new FormFieldFlowLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldFlowLinkWithDefaults

`func NewFormFieldFlowLinkWithDefaults() *FormFieldFlowLink`

NewFormFieldFlowLinkWithDefaults instantiates a new FormFieldFlowLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldFlowLink) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldFlowLink) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldFlowLink) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldFlowLink) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldFlowLink) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldFlowLink) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldFlowLink) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldFlowLink) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldFlowLink) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldFlowLink) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetKey

`func (o *FormFieldFlowLink) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldFlowLink) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldFlowLink) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormFieldFlowLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldFlowLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldFlowLink) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStyles

`func (o *FormFieldFlowLink) GetStyles() FormFlowLinkStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormFieldFlowLink) GetStylesOk() (*FormFlowLinkStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormFieldFlowLink) SetStyles(v FormFlowLinkStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormFieldFlowLink) HasStyles() bool`

HasStyles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


