# FormFieldQrCode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**QrCodeType** | [**EnumFormQrCodeType**](EnumFormQrCodeType.md) |  | 
**Alignment** | [**EnumFormItemAlignment**](EnumFormItemAlignment.md) |  | 
**ShowBorder** | **bool** | A boolean that specifies the border visibility. | 

## Methods

### NewFormFieldQrCode

`func NewFormFieldQrCode(type_ EnumFormFieldType, position FormFieldCommonPosition, qrCodeType EnumFormQrCodeType, alignment EnumFormItemAlignment, showBorder bool, ) *FormFieldQrCode`

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


### GetQrCodeType

`func (o *FormFieldQrCode) GetQrCodeType() EnumFormQrCodeType`

GetQrCodeType returns the QrCodeType field if non-nil, zero value otherwise.

### GetQrCodeTypeOk

`func (o *FormFieldQrCode) GetQrCodeTypeOk() (*EnumFormQrCodeType, bool)`

GetQrCodeTypeOk returns a tuple with the QrCodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeType

`func (o *FormFieldQrCode) SetQrCodeType(v EnumFormQrCodeType)`

SetQrCodeType sets QrCodeType field to given value.


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


### GetShowBorder

`func (o *FormFieldQrCode) GetShowBorder() bool`

GetShowBorder returns the ShowBorder field if non-nil, zero value otherwise.

### GetShowBorderOk

`func (o *FormFieldQrCode) GetShowBorderOk() (*bool, bool)`

GetShowBorderOk returns a tuple with the ShowBorder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowBorder

`func (o *FormFieldQrCode) SetShowBorder(v bool)`

SetShowBorder sets ShowBorder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


