# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The time the key resource expires.The date and time at which the subscription resource was created (ISO 8601 format). | [optional] [readonly] 
**Enabled** | **bool** | A boolean that specifies whether a created or updated subscription should be active or suspended. A suspended state (&#x60;\&quot;enabled\&quot;:false&#x60;) accumulates all matched events, but these events are not delivered until the subscription becomes active again (&#x60;\&quot;enabled\&quot;:true&#x60;). For suspended subscriptions, events accumulate for a maximum of two weeks. Events older than two weeks are deleted. Restarted subscriptions receive the saved events (up to two weeks from the restart date). This is a required property. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**FilterOptions** | [**SubscriptionFilterOptions**](SubscriptionFilterOptions.md) |  | 
**Format** | [**EnumSubscriptionFormat**](EnumSubscriptionFormat.md) |  | 
**Id** | Pointer to **string** | A string that specifies the user resource’s unique identifier. | [optional] [readonly] 
**HttpEndpoint** | [**SubscriptionHttpEndpoint**](SubscriptionHttpEndpoint.md) |  | 
**Name** | **string** | A string that specifies the subscription name. This is a required property. | 
**UpdatedAt** | Pointer to **time.Time** | The date and time at which the subscription resource was last updated (ISO 8601 format). | [optional] [readonly] 
**VerifyTlsCertificates** | **bool** | A boolean that specifies whether a certificates should be verified. If this property&#39;s value is set to false, then all certificates are trusted. (Setting this property&#39;s value to false introduces a security risk.) This is a required property. | 

## Methods

### NewSubscription

`func NewSubscription(enabled bool, filterOptions SubscriptionFilterOptions, format EnumSubscriptionFormat, httpEndpoint SubscriptionHttpEndpoint, name string, verifyTlsCertificates bool, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *Subscription) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Subscription) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Subscription) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Subscription) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetEnabled

`func (o *Subscription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Subscription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Subscription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *Subscription) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Subscription) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Subscription) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Subscription) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFilterOptions

`func (o *Subscription) GetFilterOptions() SubscriptionFilterOptions`

GetFilterOptions returns the FilterOptions field if non-nil, zero value otherwise.

### GetFilterOptionsOk

`func (o *Subscription) GetFilterOptionsOk() (*SubscriptionFilterOptions, bool)`

GetFilterOptionsOk returns a tuple with the FilterOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterOptions

`func (o *Subscription) SetFilterOptions(v SubscriptionFilterOptions)`

SetFilterOptions sets FilterOptions field to given value.


### GetFormat

`func (o *Subscription) GetFormat() EnumSubscriptionFormat`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *Subscription) GetFormatOk() (*EnumSubscriptionFormat, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *Subscription) SetFormat(v EnumSubscriptionFormat)`

SetFormat sets Format field to given value.


### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHttpEndpoint

`func (o *Subscription) GetHttpEndpoint() SubscriptionHttpEndpoint`

GetHttpEndpoint returns the HttpEndpoint field if non-nil, zero value otherwise.

### GetHttpEndpointOk

`func (o *Subscription) GetHttpEndpointOk() (*SubscriptionHttpEndpoint, bool)`

GetHttpEndpointOk returns a tuple with the HttpEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpEndpoint

`func (o *Subscription) SetHttpEndpoint(v SubscriptionHttpEndpoint)`

SetHttpEndpoint sets HttpEndpoint field to given value.


### GetName

`func (o *Subscription) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Subscription) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Subscription) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *Subscription) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Subscription) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Subscription) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetVerifyTlsCertificates

`func (o *Subscription) GetVerifyTlsCertificates() bool`

GetVerifyTlsCertificates returns the VerifyTlsCertificates field if non-nil, zero value otherwise.

### GetVerifyTlsCertificatesOk

`func (o *Subscription) GetVerifyTlsCertificatesOk() (*bool, bool)`

GetVerifyTlsCertificatesOk returns a tuple with the VerifyTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifyTlsCertificates

`func (o *Subscription) SetVerifyTlsCertificates(v bool)`

SetVerifyTlsCertificates sets VerifyTlsCertificates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


