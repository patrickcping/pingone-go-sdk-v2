# NotificationsPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | The name to use for the notification policy. Must be unique among the notification policies in the environment. | 
**Default** | Pointer to **bool** | Indication of whether this policy is the default notification policy for the environment. If the parameter is not provided, the value used is &#x60;false&#x60; | [optional] [default to false]
**CountryLimit** | Pointer to [**NotificationsPolicyCountryLimit**](NotificationsPolicyCountryLimit.md) |  | [optional] 
**Quotas** | [**[]NotificationsPolicyQuotasInner**](NotificationsPolicyQuotasInner.md) | Collection of objects that define the SMS/Voice limits. Each object contain the following elements- &#x60;type&#x60;, &#x60;deliveryMethods&#x60;, &#x60;total&#x60;. Currently, a policy can contain ony one such object. Note that instead of &#x60;total&#x60;, you can use the pair of fields- &#x60;claimed&#x60; and &#x60;unclaimed&#x60;. | 

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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


