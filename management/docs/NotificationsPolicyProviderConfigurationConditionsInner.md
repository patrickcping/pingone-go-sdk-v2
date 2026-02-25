# NotificationsPolicyProviderConfigurationConditionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeliveryMethods** | Pointer to [**[]EnumNotificationsPolicyProviderConfigurationConditionsDeliveryMethods**](EnumNotificationsPolicyProviderConfigurationConditionsDeliveryMethods.md) | The authentication methods that the fallback order should be used for. The array can contain VOICE, SMS, or both. | [optional] 
**Countries** | Pointer to **[]string** | The countries for which the fallback order should be used. Use the two-letter country codes from ISO 3166-1. | [optional] 
**FallbackChain** | Pointer to [**[]NotificationsPolicyProviderConfigurationConditionsInnerFallbackChainInner**](NotificationsPolicyProviderConfigurationConditionsInnerFallbackChainInner.md) | An array containing the IDs of the defined custom providers, in the order that they should be used if available. | [optional] 

## Methods

### NewNotificationsPolicyProviderConfigurationConditionsInner

`func NewNotificationsPolicyProviderConfigurationConditionsInner() *NotificationsPolicyProviderConfigurationConditionsInner`

NewNotificationsPolicyProviderConfigurationConditionsInner instantiates a new NotificationsPolicyProviderConfigurationConditionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsPolicyProviderConfigurationConditionsInnerWithDefaults

`func NewNotificationsPolicyProviderConfigurationConditionsInnerWithDefaults() *NotificationsPolicyProviderConfigurationConditionsInner`

NewNotificationsPolicyProviderConfigurationConditionsInnerWithDefaults instantiates a new NotificationsPolicyProviderConfigurationConditionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeliveryMethods

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) GetDeliveryMethods() []EnumNotificationsPolicyProviderConfigurationConditionsDeliveryMethods`

GetDeliveryMethods returns the DeliveryMethods field if non-nil, zero value otherwise.

### GetDeliveryMethodsOk

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) GetDeliveryMethodsOk() (*[]EnumNotificationsPolicyProviderConfigurationConditionsDeliveryMethods, bool)`

GetDeliveryMethodsOk returns a tuple with the DeliveryMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethods

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) SetDeliveryMethods(v []EnumNotificationsPolicyProviderConfigurationConditionsDeliveryMethods)`

SetDeliveryMethods sets DeliveryMethods field to given value.

### HasDeliveryMethods

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) HasDeliveryMethods() bool`

HasDeliveryMethods returns a boolean if a field has been set.

### GetCountries

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetFallbackChain

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) GetFallbackChain() []NotificationsPolicyProviderConfigurationConditionsInnerFallbackChainInner`

GetFallbackChain returns the FallbackChain field if non-nil, zero value otherwise.

### GetFallbackChainOk

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) GetFallbackChainOk() (*[]NotificationsPolicyProviderConfigurationConditionsInnerFallbackChainInner, bool)`

GetFallbackChainOk returns a tuple with the FallbackChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackChain

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) SetFallbackChain(v []NotificationsPolicyProviderConfigurationConditionsInnerFallbackChainInner)`

SetFallbackChain sets FallbackChain field to given value.

### HasFallbackChain

`func (o *NotificationsPolicyProviderConfigurationConditionsInner) HasFallbackChain() bool`

HasFallbackChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


