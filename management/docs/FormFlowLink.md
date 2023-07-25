# FormFlowLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A string that specifies an identifier for the field component. | 
**Label** | **string** | A string that specifies the link label. | 
**Styles** | Pointer to [**FormFlowLinkStyles**](FormFlowLinkStyles.md) |  | [optional] 

## Methods

### NewFormFlowLink

`func NewFormFlowLink(key string, label string, ) *FormFlowLink`

NewFormFlowLink instantiates a new FormFlowLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFlowLinkWithDefaults

`func NewFormFlowLinkWithDefaults() *FormFlowLink`

NewFormFlowLinkWithDefaults instantiates a new FormFlowLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FormFlowLink) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFlowLink) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFlowLink) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormFlowLink) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFlowLink) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFlowLink) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetStyles

`func (o *FormFlowLink) GetStyles() FormFlowLinkStyles`

GetStyles returns the Styles field if non-nil, zero value otherwise.

### GetStylesOk

`func (o *FormFlowLink) GetStylesOk() (*FormFlowLinkStyles, bool)`

GetStylesOk returns a tuple with the Styles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyles

`func (o *FormFlowLink) SetStyles(v FormFlowLinkStyles)`

SetStyles sets Styles field to given value.

### HasStyles

`func (o *FormFlowLink) HasStyles() bool`

HasStyles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


