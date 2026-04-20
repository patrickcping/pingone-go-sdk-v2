# FormDeviceAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** | Label for the device authentication field. | 
**Options** | [**[]AuthenticationDevice**](AuthenticationDevice.md) | List of devices available for authentication. Must not be empty. | 
**Required** | Pointer to **bool** | Whether the field is required or not. | [optional] 

## Methods

### NewFormDeviceAuthentication

`func NewFormDeviceAuthentication(label string, options []AuthenticationDevice, ) *FormDeviceAuthentication`

NewFormDeviceAuthentication instantiates a new FormDeviceAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFormDeviceAuthenticationWithDefaults

`func NewFormDeviceAuthenticationWithDefaults() *FormDeviceAuthentication`

NewFormDeviceAuthenticationWithDefaults instantiates a new FormDeviceAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *FormDeviceAuthentication) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FormDeviceAuthentication) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FormDeviceAuthentication) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetOptions

`func (o *FormDeviceAuthentication) GetOptions() []AuthenticationDevice`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FormDeviceAuthentication) GetOptionsOk() (*[]AuthenticationDevice, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FormDeviceAuthentication) SetOptions(v []AuthenticationDevice)`

SetOptions sets Options field to given value.


### GetRequired

`func (o *FormDeviceAuthentication) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FormDeviceAuthentication) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FormDeviceAuthentication) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FormDeviceAuthentication) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


