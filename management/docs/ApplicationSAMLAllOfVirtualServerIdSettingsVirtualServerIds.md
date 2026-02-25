# ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VsId** | **string** | This must be a valid SAML entity ID. | 
**Default** | Pointer to **bool** | Indicates whether the virtual server identified by the associated &#x60;vsId&#x60; is to be used as the default virtual server. | [optional] 

## Methods

### NewApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds

`func NewApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds(vsId string, ) *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds`

NewApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds instantiates a new ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIdsWithDefaults

`func NewApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIdsWithDefaults() *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds`

NewApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIdsWithDefaults instantiates a new ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVsId

`func (o *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds) GetVsId() string`

GetVsId returns the VsId field if non-nil, zero value otherwise.

### GetVsIdOk

`func (o *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds) GetVsIdOk() (*string, bool)`

GetVsIdOk returns a tuple with the VsId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsId

`func (o *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds) SetVsId(v string)`

SetVsId sets VsId field to given value.


### GetDefault

`func (o *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *ApplicationSAMLAllOfVirtualServerIdSettingsVirtualServerIds) HasDefault() bool`

HasDefault returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


