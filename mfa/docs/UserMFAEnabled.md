# UserMFAEnabled

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**MfaEnabled** | **bool** | Whether multi-factor authentication is enabled. This attribute is set to &#x60;false&#x60; by default when the user is created. | 

## Methods

### NewUserMFAEnabled

`func NewUserMFAEnabled(mfaEnabled bool, ) *UserMFAEnabled`

NewUserMFAEnabled instantiates a new UserMFAEnabled object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserMFAEnabledWithDefaults

`func NewUserMFAEnabledWithDefaults() *UserMFAEnabled`

NewUserMFAEnabledWithDefaults instantiates a new UserMFAEnabled object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *UserMFAEnabled) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *UserMFAEnabled) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *UserMFAEnabled) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *UserMFAEnabled) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetMfaEnabled

`func (o *UserMFAEnabled) GetMfaEnabled() bool`

GetMfaEnabled returns the MfaEnabled field if non-nil, zero value otherwise.

### GetMfaEnabledOk

`func (o *UserMFAEnabled) GetMfaEnabledOk() (*bool, bool)`

GetMfaEnabledOk returns a tuple with the MfaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMfaEnabled

`func (o *UserMFAEnabled) SetMfaEnabled(v bool)`

SetMfaEnabled sets MfaEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


