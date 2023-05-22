# FormQrCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QrCodeType** | [**EnumFormQrCodeType**](EnumFormQrCodeType.md) |  | 
**Alignment** | [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | 
**ShowBorder** | **bool** | A boolean that specifies the border visibility. | 

## Methods

### NewFormQrCode

`func NewFormQrCode(qrCodeType EnumFormQrCodeType, alignment EnumFormItemAlignment, showBorder bool, ) *FormQrCode`

NewFormQrCode instantiates a new FormQrCode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormQrCodeWithDefaults

`func NewFormQrCodeWithDefaults() *FormQrCode`

NewFormQrCodeWithDefaults instantiates a new FormQrCode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQrCodeType

`func (o *FormQrCode) GetQrCodeType() EnumFormQrCodeType`

GetQrCodeType returns the QrCodeType field if non-nil, zero value otherwise.

### GetQrCodeTypeOk

`func (o *FormQrCode) GetQrCodeTypeOk() (*EnumFormQrCodeType, bool)`

GetQrCodeTypeOk returns a tuple with the QrCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeType

`func (o *FormQrCode) SetQrCodeType(v EnumFormQrCodeType)`

SetQrCodeType sets QrCodeType field to given value.


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


### GetShowBorder

`func (o *FormQrCode) GetShowBorder() bool`

GetShowBorder returns the ShowBorder field if non-nil, zero value otherwise.

### GetShowBorderOk

`func (o *FormQrCode) GetShowBorderOk() (*bool, bool)`

GetShowBorderOk returns a tuple with the ShowBorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBorder

`func (o *FormQrCode) SetShowBorder(v bool)`

SetShowBorder sets ShowBorder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


