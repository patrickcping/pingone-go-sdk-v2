# SubscriptionFilterOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludedActionTypes** | **[]string** | A non-empty array that specifies the list of action types that should be matched for the subscription. This is a required property. | 
**IncludedApplications** | Pointer to [**[]SubscriptionFilterOptionsIncludedApplicationsInner**](SubscriptionFilterOptionsIncludedApplicationsInner.md) | An array that specifies the list of applications (by ID) whose events are monitored by the subscription (maximum of 10 IDs in the array). This is an optional property. If a list of applications is not provided, events are monitored for all applications in the environment. | [optional] 
**IncludedPopulations** | Pointer to [**[]SubscriptionFilterOptionsIncludedApplicationsInner**](SubscriptionFilterOptionsIncludedApplicationsInner.md) | An array that specifies the list of populations (by ID) whose events are monitored by the subscription (maximum of 10 IDs in the array). This property matches events for users in the specified populations, as opposed to events generated in which the user in one of the populations is the actor. This is an optional property. | [optional] 
**IncludedTags** | Pointer to [**[]EnumSubscriptionFilterIncludedTags**](EnumSubscriptionFilterIncludedTags.md) | An array of tags that events must have to be monitored by the subscription. If tags are not specified, all events are monitored. Currently, the available tags are &#x60;adminIdentityEvent&#x60;. Identifies the event as the action of an administrator on other administrators. | [optional] 
**IpAddressExposed** | Pointer to **bool** | Whether the IP address of an actor should be present in the &#x60;source&#x60; section of the event. | [optional] [default to false]
**UserAgentExposed** | Pointer to **bool** | Whether the User-Agent HTTP header of an event should be present in the &#x60;source&#x60; section of the event. | [optional] [default to false]

## Methods

### NewSubscriptionFilterOptions

`func NewSubscriptionFilterOptions(includedActionTypes []string, ) *SubscriptionFilterOptions`

NewSubscriptionFilterOptions instantiates a new SubscriptionFilterOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionFilterOptionsWithDefaults

`func NewSubscriptionFilterOptionsWithDefaults() *SubscriptionFilterOptions`

NewSubscriptionFilterOptionsWithDefaults instantiates a new SubscriptionFilterOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludedActionTypes

`func (o *SubscriptionFilterOptions) GetIncludedActionTypes() []string`

GetIncludedActionTypes returns the IncludedActionTypes field if non-nil, zero value otherwise.

### GetIncludedActionTypesOk

`func (o *SubscriptionFilterOptions) GetIncludedActionTypesOk() (*[]string, bool)`

GetIncludedActionTypesOk returns a tuple with the IncludedActionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedActionTypes

`func (o *SubscriptionFilterOptions) SetIncludedActionTypes(v []string)`

SetIncludedActionTypes sets IncludedActionTypes field to given value.


### GetIncludedApplications

`func (o *SubscriptionFilterOptions) GetIncludedApplications() []SubscriptionFilterOptionsIncludedApplicationsInner`

GetIncludedApplications returns the IncludedApplications field if non-nil, zero value otherwise.

### GetIncludedApplicationsOk

`func (o *SubscriptionFilterOptions) GetIncludedApplicationsOk() (*[]SubscriptionFilterOptionsIncludedApplicationsInner, bool)`

GetIncludedApplicationsOk returns a tuple with the IncludedApplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedApplications

`func (o *SubscriptionFilterOptions) SetIncludedApplications(v []SubscriptionFilterOptionsIncludedApplicationsInner)`

SetIncludedApplications sets IncludedApplications field to given value.

### HasIncludedApplications

`func (o *SubscriptionFilterOptions) HasIncludedApplications() bool`

HasIncludedApplications returns a boolean if a field has been set.

### GetIncludedPopulations

`func (o *SubscriptionFilterOptions) GetIncludedPopulations() []SubscriptionFilterOptionsIncludedApplicationsInner`

GetIncludedPopulations returns the IncludedPopulations field if non-nil, zero value otherwise.

### GetIncludedPopulationsOk

`func (o *SubscriptionFilterOptions) GetIncludedPopulationsOk() (*[]SubscriptionFilterOptionsIncludedApplicationsInner, bool)`

GetIncludedPopulationsOk returns a tuple with the IncludedPopulations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedPopulations

`func (o *SubscriptionFilterOptions) SetIncludedPopulations(v []SubscriptionFilterOptionsIncludedApplicationsInner)`

SetIncludedPopulations sets IncludedPopulations field to given value.

### HasIncludedPopulations

`func (o *SubscriptionFilterOptions) HasIncludedPopulations() bool`

HasIncludedPopulations returns a boolean if a field has been set.

### GetIncludedTags

`func (o *SubscriptionFilterOptions) GetIncludedTags() []EnumSubscriptionFilterIncludedTags`

GetIncludedTags returns the IncludedTags field if non-nil, zero value otherwise.

### GetIncludedTagsOk

`func (o *SubscriptionFilterOptions) GetIncludedTagsOk() (*[]EnumSubscriptionFilterIncludedTags, bool)`

GetIncludedTagsOk returns a tuple with the IncludedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludedTags

`func (o *SubscriptionFilterOptions) SetIncludedTags(v []EnumSubscriptionFilterIncludedTags)`

SetIncludedTags sets IncludedTags field to given value.

### HasIncludedTags

`func (o *SubscriptionFilterOptions) HasIncludedTags() bool`

HasIncludedTags returns a boolean if a field has been set.

### GetIpAddressExposed

`func (o *SubscriptionFilterOptions) GetIpAddressExposed() bool`

GetIpAddressExposed returns the IpAddressExposed field if non-nil, zero value otherwise.

### GetIpAddressExposedOk

`func (o *SubscriptionFilterOptions) GetIpAddressExposedOk() (*bool, bool)`

GetIpAddressExposedOk returns a tuple with the IpAddressExposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddressExposed

`func (o *SubscriptionFilterOptions) SetIpAddressExposed(v bool)`

SetIpAddressExposed sets IpAddressExposed field to given value.

### HasIpAddressExposed

`func (o *SubscriptionFilterOptions) HasIpAddressExposed() bool`

HasIpAddressExposed returns a boolean if a field has been set.

### GetUserAgentExposed

`func (o *SubscriptionFilterOptions) GetUserAgentExposed() bool`

GetUserAgentExposed returns the UserAgentExposed field if non-nil, zero value otherwise.

### GetUserAgentExposedOk

`func (o *SubscriptionFilterOptions) GetUserAgentExposedOk() (*bool, bool)`

GetUserAgentExposedOk returns a tuple with the UserAgentExposed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentExposed

`func (o *SubscriptionFilterOptions) SetUserAgentExposed(v bool)`

SetUserAgentExposed sets UserAgentExposed field to given value.

### HasUserAgentExposed

`func (o *SubscriptionFilterOptions) HasUserAgentExposed() bool`

HasUserAgentExposed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


