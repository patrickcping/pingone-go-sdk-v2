# FormFieldDeviceAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumFormFieldType**](EnumFormFieldType.md) |  | 
**Position** | [**FormFieldCommonPosition**](FormFieldCommonPosition.md) |  | 
**Visibility** | Pointer to [**FormFieldCommonVisibility**](FormFieldCommonVisibility.md) |  | [optional] 
**Label** | **string** | Label for the device authentication field. | 
**Options** | [**[]AuthenticationDevice**](AuthenticationDevice.md) | List of devices available for authentication. Must not be empty. | 
**Required** | Pointer to **bool** | Whether the field is required or not. | [optional] 

## Methods

### NewFormFieldDeviceAuthentication

`func NewFormFieldDeviceAuthentication(type_ EnumFormFieldType, position FormFieldCommonPosition, label string, options []AuthenticationDevice, ) *FormFieldDeviceAuthentication`

NewFormFieldDeviceAuthentication instantiates a new FormFieldDeviceAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormFieldDeviceAuthenticationWithDefaults

`func NewFormFieldDeviceAuthenticationWithDefaults() *FormFieldDeviceAuthentication`

NewFormFieldDeviceAuthenticationWithDefaults instantiates a new FormFieldDeviceAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FormFieldDeviceAuthentication) GetType() EnumFormFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FormFieldDeviceAuthentication) GetTypeOk() (*EnumFormFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FormFieldDeviceAuthentication) SetType(v EnumFormFieldType)`

SetType sets Type field to given value.


### GetPosition

`func (o *FormFieldDeviceAuthentication) GetPosition() FormFieldCommonPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *FormFieldDeviceAuthentication) GetPositionOk() (*FormFieldCommonPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *FormFieldDeviceAuthentication) SetPosition(v FormFieldCommonPosition)`

SetPosition sets Position field to given value.


### GetVisibility

`func (o *FormFieldDeviceAuthentication) GetVisibility() FormFieldCommonVisibility`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *FormFieldDeviceAuthentication) GetVisibilityOk() (*FormFieldCommonVisibility, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *FormFieldDeviceAuthentication) SetVisibility(v FormFieldCommonVisibility)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *FormFieldDeviceAuthentication) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetLabel

`func (o *FormFieldDeviceAuthentication) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormFieldDeviceAuthentication) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormFieldDeviceAuthentication) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOptions

`func (o *FormFieldDeviceAuthentication) GetOptions() []AuthenticationDevice`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormFieldDeviceAuthentication) GetOptionsOk() (*[]AuthenticationDevice, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormFieldDeviceAuthentication) SetOptions(v []AuthenticationDevice)`

SetOptions sets Options field to given value.


### GetRequired

`func (o *FormFieldDeviceAuthentication) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormFieldDeviceAuthentication) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormFieldDeviceAuthentication) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormFieldDeviceAuthentication) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


