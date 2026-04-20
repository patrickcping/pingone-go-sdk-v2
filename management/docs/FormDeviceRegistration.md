# FormDeviceRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | A string that specifies an identifier for the field component. | 
**Label** | **string** | Label for the device registration field. | 
**Options** | [**[]RegistrationDevice**](RegistrationDevice.md) | List of devices available for registration. Must not be empty. | 
**Required** | Pointer to **bool** | Whether the field is required or not. | [optional] 

## Methods

### NewFormDeviceRegistration

`func NewFormDeviceRegistration(key string, label string, options []RegistrationDevice, ) *FormDeviceRegistration`

NewFormDeviceRegistration instantiates a new FormDeviceRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDeviceRegistrationWithDefaults

`func NewFormDeviceRegistrationWithDefaults() *FormDeviceRegistration`

NewFormDeviceRegistrationWithDefaults instantiates a new FormDeviceRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *FormDeviceRegistration) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FormDeviceRegistration) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FormDeviceRegistration) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *FormDeviceRegistration) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormDeviceRegistration) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormDeviceRegistration) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOptions

`func (o *FormDeviceRegistration) GetOptions() []RegistrationDevice`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormDeviceRegistration) GetOptionsOk() (*[]RegistrationDevice, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormDeviceRegistration) SetOptions(v []RegistrationDevice)`

SetOptions sets Options field to given value.


### GetRequired

`func (o *FormDeviceRegistration) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormDeviceRegistration) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormDeviceRegistration) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormDeviceRegistration) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


