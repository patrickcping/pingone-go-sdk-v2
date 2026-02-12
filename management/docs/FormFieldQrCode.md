# FormFieldQrCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**Alignment** | [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | 
**Size** | [**EnumFormItemSize**](EnumFormItemSize.md) |  | 
**FallbackText** | Pointer to **string** | A string that specifies the text label for fallback under the QR code. | [optional] 

## Methods

### NewFormFieldQrCode

`func NewFormFieldQrCode(type_ EnumFormFieldType, position FormFieldCommonPosition, alignment EnumFormItemAlignment, size EnumFormItemSize, ) *FormFieldQrCode`

NewFormFieldQrCode instantiates a new FormFieldQrCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldQrCodeWithDefaults

`func NewFormFieldQrCodeWithDefaults() *FormFieldQrCode`

NewFormFieldQrCodeWithDefaults instantiates a new FormFieldQrCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldQrCode) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldQrCode) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldQrCode) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldQrCode) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldQrCode) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldQrCode) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldQrCode) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldQrCode) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldQrCode) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldQrCode) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetAlignment

`func (o *FormFieldQrCode) GetAlignment() EnumFormItemAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *FormFieldQrCode) GetAlignmentOk() (*EnumFormItemAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *FormFieldQrCode) SetAlignment(v EnumFormItemAlignment)`

SetAlignment sets Alignment field to given value.


### GetSize

`func (o *FormFieldQrCode) GetSize() EnumFormItemSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FormFieldQrCode) GetSizeOk() (*EnumFormItemSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FormFieldQrCode) SetSize(v EnumFormItemSize)`

SetSize sets Size field to given value.


### GetFallbackText

`func (o *FormFieldQrCode) GetFallbackText() string`

GetFallbackText returns the FallbackText field if non-nil, zero value otherwise.

### GetFallbackTextOk

`func (o *FormFieldQrCode) GetFallbackTextOk() (*string, bool)`

GetFallbackTextOk returns a tuple with the FallbackText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackText

`func (o *FormFieldQrCode) SetFallbackText(v string)`

SetFallbackText sets FallbackText field to given value.

### HasFallbackText

`func (o *FormFieldQrCode) HasFallbackText() bool`

HasFallbackText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


