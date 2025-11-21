# NotificationsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**CooldownConfiguration** | Pointer to [**NotificationsPolicyCooldownConfiguration**](NotificationsPolicyCooldownConfiguration.md) |  | [optional] 
**CountryLimit** | Pointer to [**NotificationsPolicyCountryLimit**](NotificationsPolicyCountryLimit.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Default** | Pointer to **bool** | Indication of whether this policy is the default notification policy for the environment. If the parameter is not provided, the value used is &#x60;false&#x60; | [optional] [default to false]
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | The name to use for the notification policy. Must be unique among the notification policies in the environment. | 
**ProviderConfiguration** | Pointer to [**NotificationsPolicyProviderConfiguration**](NotificationsPolicyProviderConfiguration.md) |  | [optional] 
**Quotas** | [**[]NotificationsPolicyQuotasInner**](NotificationsPolicyQuotasInner.md) | Collection of objects that define the SMS/Voice limits. Each object contain the following elements- &#x60;type&#x60;, &#x60;deliveryMethods&#x60;, &#x60;total&#x60;. Currently, a policy can contain ony one such object. Note that instead of &#x60;total&#x60;, you can use the pair of fields- &#x60;claimed&#x60; and &#x60;unclaimed&#x60;. | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewNotificationsPolicy

`func NewNotificationsPolicy(name string, quotas []NotificationsPolicyQuotasInner, ) *NotificationsPolicy`

NewNotificationsPolicy instantiates a new NotificationsPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsPolicyWithDefaults

`func NewNotificationsPolicyWithDefaults() *NotificationsPolicy`

NewNotificationsPolicyWithDefaults instantiates a new NotificationsPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *NotificationsPolicy) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationsPolicy) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationsPolicy) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotificationsPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCooldownConfiguration

`func (o *NotificationsPolicy) GetCooldownConfiguration() NotificationsPolicyCooldownConfiguration`

GetCooldownConfiguration returns the CooldownConfiguration field if non-nil, zero value otherwise.

### GetCooldownConfigurationOk

`func (o *NotificationsPolicy) GetCooldownConfigurationOk() (*NotificationsPolicyCooldownConfiguration, bool)`

GetCooldownConfigurationOk returns a tuple with the CooldownConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCooldownConfiguration

`func (o *NotificationsPolicy) SetCooldownConfiguration(v NotificationsPolicyCooldownConfiguration)`

SetCooldownConfiguration sets CooldownConfiguration field to given value.

### HasCooldownConfiguration

`func (o *NotificationsPolicy) HasCooldownConfiguration() bool`

HasCooldownConfiguration returns a boolean if a field has been set.

### GetCountryLimit

`func (o *NotificationsPolicy) GetCountryLimit() NotificationsPolicyCountryLimit`

GetCountryLimit returns the CountryLimit field if non-nil, zero value otherwise.

### GetCountryLimitOk

`func (o *NotificationsPolicy) GetCountryLimitOk() (*NotificationsPolicyCountryLimit, bool)`

GetCountryLimitOk returns a tuple with the CountryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryLimit

`func (o *NotificationsPolicy) SetCountryLimit(v NotificationsPolicyCountryLimit)`

SetCountryLimit sets CountryLimit field to given value.

### HasCountryLimit

`func (o *NotificationsPolicy) HasCountryLimit() bool`

HasCountryLimit returns a boolean if a field has been set.

### GetCreatedAt

`func (o *NotificationsPolicy) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationsPolicy) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationsPolicy) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationsPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *NotificationsPolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *NotificationsPolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *NotificationsPolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *NotificationsPolicy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *NotificationsPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NotificationsPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationsPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationsPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetProviderConfiguration

`func (o *NotificationsPolicy) GetProviderConfiguration() NotificationsPolicyProviderConfiguration`

GetProviderConfiguration returns the ProviderConfiguration field if non-nil, zero value otherwise.

### GetProviderConfigurationOk

`func (o *NotificationsPolicy) GetProviderConfigurationOk() (*NotificationsPolicyProviderConfiguration, bool)`

GetProviderConfigurationOk returns a tuple with the ProviderConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderConfiguration

`func (o *NotificationsPolicy) SetProviderConfiguration(v NotificationsPolicyProviderConfiguration)`

SetProviderConfiguration sets ProviderConfiguration field to given value.

### HasProviderConfiguration

`func (o *NotificationsPolicy) HasProviderConfiguration() bool`

HasProviderConfiguration returns a boolean if a field has been set.

### GetQuotas

`func (o *NotificationsPolicy) GetQuotas() []NotificationsPolicyQuotasInner`

GetQuotas returns the Quotas field if non-nil, zero value otherwise.

### GetQuotasOk

`func (o *NotificationsPolicy) GetQuotasOk() (*[]NotificationsPolicyQuotasInner, bool)`

GetQuotasOk returns a tuple with the Quotas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuotas

`func (o *NotificationsPolicy) SetQuotas(v []NotificationsPolicyQuotasInner)`

SetQuotas sets Quotas field to given value.


### GetUpdatedAt

`func (o *NotificationsPolicy) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsPolicy) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsPolicy) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


