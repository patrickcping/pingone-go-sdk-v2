# MFASettingsPhoneExtensions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Set to &#x60;true&#x60; to allow one-time passwords to be delivered via voice to phone numbers that include extensions. Set to &#x60;false&#x60; to disable support for phone numbers with extensions. By default, support for extensions is disabled. | [optional] 

## Methods

### NewMFASettingsPhoneExtensions

`func NewMFASettingsPhoneExtensions() *MFASettingsPhoneExtensions`

NewMFASettingsPhoneExtensions instantiates a new MFASettingsPhoneExtensions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFASettingsPhoneExtensionsWithDefaults

`func NewMFASettingsPhoneExtensionsWithDefaults() *MFASettingsPhoneExtensions`

NewMFASettingsPhoneExtensionsWithDefaults instantiates a new MFASettingsPhoneExtensions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MFASettingsPhoneExtensions) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MFASettingsPhoneExtensions) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MFASettingsPhoneExtensions) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MFASettingsPhoneExtensions) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


