# NotificationsPolicyProviderConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]NotificationsPolicyProviderConfigurationConditionsInner**](NotificationsPolicyProviderConfigurationConditionsInner.md) | Each element in the conditions array represents the provider fallback order to use for a specific group of countries for the specified methods (SMS and/or Voice). Note that in addition to any country-specific orders specified, the array must contain a fallback order without the countries array in the object. This order is used for all countries that were not specified in one of the other fallback orders defined. | [optional] 

## Methods

### NewNotificationsPolicyProviderConfiguration

`func NewNotificationsPolicyProviderConfiguration() *NotificationsPolicyProviderConfiguration`

NewNotificationsPolicyProviderConfiguration instantiates a new NotificationsPolicyProviderConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsPolicyProviderConfigurationWithDefaults

`func NewNotificationsPolicyProviderConfigurationWithDefaults() *NotificationsPolicyProviderConfiguration`

NewNotificationsPolicyProviderConfigurationWithDefaults instantiates a new NotificationsPolicyProviderConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *NotificationsPolicyProviderConfiguration) GetConditions() []NotificationsPolicyProviderConfigurationConditionsInner`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *NotificationsPolicyProviderConfiguration) GetConditionsOk() (*[]NotificationsPolicyProviderConfigurationConditionsInner, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *NotificationsPolicyProviderConfiguration) SetConditions(v []NotificationsPolicyProviderConfigurationConditionsInner)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *NotificationsPolicyProviderConfiguration) HasConditions() bool`

HasConditions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


