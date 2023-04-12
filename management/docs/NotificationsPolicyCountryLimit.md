# NotificationsPolicyCountryLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumNotificationsPolicyCountryLimitType**](EnumNotificationsPolicyCountryLimitType.md) |  | 
**DeliveryMethods** | Pointer to [**[]EnumNotificationsPolicyCountryLimitDeliveryMethod**](EnumNotificationsPolicyCountryLimitDeliveryMethod.md) | The delivery methods that the defined limitation should be applied to. Content of the array can be &#x60;SMS&#x60;, &#x60;Voice&#x60;, or both. If the parameter is not provided, the default is &#x60;SMS&#x60; and &#x60;Voice&#x60;. | [optional] 
**Countries** | **[]string** | The countries where the specified methods should be allowed or denied. Use the two-letter country codes from ISO 3166-1. | 

## Methods

### NewNotificationsPolicyCountryLimit

`func NewNotificationsPolicyCountryLimit(type_ EnumNotificationsPolicyCountryLimitType, countries []string, ) *NotificationsPolicyCountryLimit`

NewNotificationsPolicyCountryLimit instantiates a new NotificationsPolicyCountryLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsPolicyCountryLimitWithDefaults

`func NewNotificationsPolicyCountryLimitWithDefaults() *NotificationsPolicyCountryLimit`

NewNotificationsPolicyCountryLimitWithDefaults instantiates a new NotificationsPolicyCountryLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationsPolicyCountryLimit) GetType() EnumNotificationsPolicyCountryLimitType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationsPolicyCountryLimit) GetTypeOk() (*EnumNotificationsPolicyCountryLimitType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationsPolicyCountryLimit) SetType(v EnumNotificationsPolicyCountryLimitType)`

SetType sets Type field to given value.


### GetDeliveryMethods

`func (o *NotificationsPolicyCountryLimit) GetDeliveryMethods() []EnumNotificationsPolicyCountryLimitDeliveryMethod`

GetDeliveryMethods returns the DeliveryMethods field if non-nil, zero value otherwise.

### GetDeliveryMethodsOk

`func (o *NotificationsPolicyCountryLimit) GetDeliveryMethodsOk() (*[]EnumNotificationsPolicyCountryLimitDeliveryMethod, bool)`

GetDeliveryMethodsOk returns a tuple with the DeliveryMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMethods

`func (o *NotificationsPolicyCountryLimit) SetDeliveryMethods(v []EnumNotificationsPolicyCountryLimitDeliveryMethod)`

SetDeliveryMethods sets DeliveryMethods field to given value.

### HasDeliveryMethods

`func (o *NotificationsPolicyCountryLimit) HasDeliveryMethods() bool`

HasDeliveryMethods returns a boolean if a field has been set.

### GetCountries

`func (o *NotificationsPolicyCountryLimit) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *NotificationsPolicyCountryLimit) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *NotificationsPolicyCountryLimit) SetCountries(v []string)`

SetCountries sets Countries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


