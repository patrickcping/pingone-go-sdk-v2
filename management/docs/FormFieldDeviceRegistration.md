# FormFieldDeviceRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**Key** | **string** | A string that specifies an identifier for the field component. | 
**Label** | **string** | Label for the device registration field. | 
**Options** | [**[]RegistrationDevice**](RegistrationDevice.md) | List of devices available for registration. Must not be empty. | 
**Required** | Pointer to **bool** | Whether the field is required or not. | [optional] 

## Methods

### NewFormFieldDeviceRegistration

`func NewFormFieldDeviceRegistration(type_ EnumFormFieldType, position FormFieldCommonPosition, key string, label string, options []RegistrationDevice, ) *FormFieldDeviceRegistration`

NewFormFieldDeviceRegistration instantiates a new FormFieldDeviceRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldDeviceRegistrationWithDefaults

`func NewFormFieldDeviceRegistrationWithDefaults() *FormFieldDeviceRegistration`

NewFormFieldDeviceRegistrationWithDefaults instantiates a new FormFieldDeviceRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldDeviceRegistration) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldDeviceRegistration) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldDeviceRegistration) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldDeviceRegistration) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldDeviceRegistration) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldDeviceRegistration) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldDeviceRegistration) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldDeviceRegistration) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldDeviceRegistration) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldDeviceRegistration) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetKey

`func (o *FormFieldDeviceRegistration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormFieldDeviceRegistration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormFieldDeviceRegistration) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormFieldDeviceRegistration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldDeviceRegistration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldDeviceRegistration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOptions

`func (o *FormFieldDeviceRegistration) GetOptions() []RegistrationDevice`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldDeviceRegistration) GetOptionsOk() (*[]RegistrationDevice, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldDeviceRegistration) SetOptions(v []RegistrationDevice)`

SetOptions sets Options field to given value.


### GetRequired

`func (o *FormFieldDeviceRegistration) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldDeviceRegistration) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldDeviceRegistration) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormFieldDeviceRegistration) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


