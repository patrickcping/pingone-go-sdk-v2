# ApplicationSAMLAllOfVirtualServerIdSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Indicates whether the virtual server ID or IDs specified are to be used. Defaults to &#x60;false&#x60;. | [optional] [default to false]
**VirtualServerIds** | Pointer to [**[]ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds**](ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds.md) | Required if &#x60;enabled&#x60; is &#x60;true&#x60;. Contains the list of virtual ID or IDs to be used. | [optional] 

## Methods

### NewApplicationSAMLAllOfVirtualServerIdSettings

`func NewApplicationSAMLAllOfVirtualServerIdSettings() *ApplicationSAMLAllOfVirtualServerIdSettings`

NewApplicationSAMLAllOfVirtualServerIdSettings instantiates a new ApplicationSAMLAllOfVirtualServerIdSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSAMLAllOfVirtualServerIdSettingsWithDefaults

`func NewApplicationSAMLAllOfVirtualServerIdSettingsWithDefaults() *ApplicationSAMLAllOfVirtualServerIdSettings`

NewApplicationSAMLAllOfVirtualServerIdSettingsWithDefaults instantiates a new ApplicationSAMLAllOfVirtualServerIdSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ApplicationSAMLAllOfVirtualServerIdSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ApplicationSAMLAllOfVirtualServerIdSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ApplicationSAMLAllOfVirtualServerIdSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ApplicationSAMLAllOfVirtualServerIdSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVirtualServerIds

`func (o *ApplicationSAMLAllOfVirtualServerIdSettings) GetVirtualServerIds() []ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds`

GetVirtualServerIds returns the VirtualServerIds field if non-nil, zero value otherwise.

### GetVirtualServerIdsOk

`func (o *ApplicationSAMLAllOfVirtualServerIdSettings) GetVirtualServerIdsOk() (*[]ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds, bool)`

GetVirtualServerIdsOk returns a tuple with the VirtualServerIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualServerIds

`func (o *ApplicationSAMLAllOfVirtualServerIdSettings) SetVirtualServerIds(v []ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds)`

SetVirtualServerIds sets VirtualServerIds field to given value.

### HasVirtualServerIds

`func (o *ApplicationSAMLAllOfVirtualServerIdSettings) HasVirtualServerIds() bool`

HasVirtualServerIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


