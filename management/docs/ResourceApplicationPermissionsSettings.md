# ResourceApplicationPermissionsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimEnabled** | Pointer to **bool** | A setting to enable application permission claims in the access token. If this property is omitted, the value is set to &#x60;false&#x60;. | [optional] 

## Methods

### NewResourceApplicationPermissionsSettings

`func NewResourceApplicationPermissionsSettings() *ResourceApplicationPermissionsSettings`

NewResourceApplicationPermissionsSettings instantiates a new ResourceApplicationPermissionsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceApplicationPermissionsSettingsWithDefaults

`func NewResourceApplicationPermissionsSettingsWithDefaults() *ResourceApplicationPermissionsSettings`

NewResourceApplicationPermissionsSettingsWithDefaults instantiates a new ResourceApplicationPermissionsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimEnabled

`func (o *ResourceApplicationPermissionsSettings) GetClaimEnabled() bool`

GetClaimEnabled returns the ClaimEnabled field if non-nil, zero value otherwise.

### GetClaimEnabledOk

`func (o *ResourceApplicationPermissionsSettings) GetClaimEnabledOk() (*bool, bool)`

GetClaimEnabledOk returns a tuple with the ClaimEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimEnabled

`func (o *ResourceApplicationPermissionsSettings) SetClaimEnabled(v bool)`

SetClaimEnabled sets ClaimEnabled field to given value.

### HasClaimEnabled

`func (o *ResourceApplicationPermissionsSettings) HasClaimEnabled() bool`

HasClaimEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


