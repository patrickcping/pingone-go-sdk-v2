# FormFieldPolling

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**PollingAppearance** | [**EnumFormPollingAppearance**](EnumFormPollingAppearance.md) |  | 
**Size** | [**EnumFormItemSize**](EnumFormItemSize.md) |  | 

## Methods

### NewFormFieldPolling

`func NewFormFieldPolling(type_ EnumFormFieldType, position FormFieldCommonPosition, pollingAppearance EnumFormPollingAppearance, size EnumFormItemSize, ) *FormFieldPolling`

NewFormFieldPolling instantiates a new FormFieldPolling object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldPollingWithDefaults

`func NewFormFieldPollingWithDefaults() *FormFieldPolling`

NewFormFieldPollingWithDefaults instantiates a new FormFieldPolling object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldPolling) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldPolling) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldPolling) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldPolling) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldPolling) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldPolling) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldPolling) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldPolling) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldPolling) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldPolling) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetPollingAppearance

`func (o *FormFieldPolling) GetPollingAppearance() EnumFormPollingAppearance`

GetPollingAppearance returns the PollingAppearance field if non-nil, zero value otherwise.

### GetPollingAppearanceOk

`func (o *FormFieldPolling) GetPollingAppearanceOk() (*EnumFormPollingAppearance, bool)`

GetPollingAppearanceOk returns a tuple with the PollingAppearance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingAppearance

`func (o *FormFieldPolling) SetPollingAppearance(v EnumFormPollingAppearance)`

SetPollingAppearance sets PollingAppearance field to given value.


### GetSize

`func (o *FormFieldPolling) GetSize() EnumFormItemSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FormFieldPolling) GetSizeOk() (*EnumFormItemSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FormFieldPolling) SetSize(v EnumFormItemSize)`

SetSize sets Size field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


