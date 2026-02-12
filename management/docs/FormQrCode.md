# FormQrCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alignment** | [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | 
**Size** | [**EnumFormItemSize**](EnumFormItemSize.md) |  | 
**FallbackText** | Pointer to **string** | A string that specifies the text label for fallback under the QR code. | [optional] 

## Methods

### NewFormQrCode

`func NewFormQrCode(alignment EnumFormItemAlignment, size EnumFormItemSize, ) *FormQrCode`

NewFormQrCode instantiates a new FormQrCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormQrCodeWithDefaults

`func NewFormQrCodeWithDefaults() *FormQrCode`

NewFormQrCodeWithDefaults instantiates a new FormQrCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlignment

`func (o *FormQrCode) GetAlignment() EnumFormItemAlignment`

GetAlignment returns the Alignment field if non-nil, zero value otherwise.

### GetAlignmentOk

`func (o *FormQrCode) GetAlignmentOk() (*EnumFormItemAlignment, bool)`

GetAlignmentOk returns a tuple with the Alignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlignment

`func (o *FormQrCode) SetAlignment(v EnumFormItemAlignment)`

SetAlignment sets Alignment field to given value.


### GetSize

`func (o *FormQrCode) GetSize() EnumFormItemSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FormQrCode) GetSizeOk() (*EnumFormItemSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FormQrCode) SetSize(v EnumFormItemSize)`

SetSize sets Size field to given value.


### GetFallbackText

`func (o *FormQrCode) GetFallbackText() string`

GetFallbackText returns the FallbackText field if non-nil, zero value otherwise.

### GetFallbackTextOk

`func (o *FormQrCode) GetFallbackTextOk() (*string, bool)`

GetFallbackTextOk returns a tuple with the FallbackText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackText

`func (o *FormQrCode) SetFallbackText(v string)`

SetFallbackText sets FallbackText field to given value.

### HasFallbackText

`func (o *FormQrCode) HasFallbackText() bool`

HasFallbackText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


