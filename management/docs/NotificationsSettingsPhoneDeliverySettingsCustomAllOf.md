# NotificationsSettingsPhoneDeliverySettingsCustomAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The customer provider&#39;s name. | 
**Requests** | [**[]NotificationsSettingsPhoneDeliverySettingsCustomRequest**](NotificationsSettingsPhoneDeliverySettingsCustomRequest.md) |  | 
**Authentication** | [**NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication**](NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication.md) |  | 
**Numbers** | Pointer to [**[]NotificationsSettingsPhoneDeliverySettingsCustomNumbers**](NotificationsSettingsPhoneDeliverySettingsCustomNumbers.md) |  | [optional] 

## Methods

### NewNotificationsSettingsPhoneDeliverySettingsCustomAllOf

`func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOf(name string, requests []NotificationsSettingsPhoneDeliverySettingsCustomRequest, authentication NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication, ) *NotificationsSettingsPhoneDeliverySettingsCustomAllOf`

NewNotificationsSettingsPhoneDeliverySettingsCustomAllOf instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfWithDefaults

`func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustomAllOf`

NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) SetName(v string)`

SetName sets Name field to given value.


### GetRequests

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetRequests() []NotificationsSettingsPhoneDeliverySettingsCustomRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetRequestsOk() (*[]NotificationsSettingsPhoneDeliverySettingsCustomRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) SetRequests(v []NotificationsSettingsPhoneDeliverySettingsCustomRequest)`

SetRequests sets Requests field to given value.


### GetAuthentication

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetAuthentication() NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetAuthenticationOk() (*NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) SetAuthentication(v NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication)`

SetAuthentication sets Authentication field to given value.


### GetNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetNumbers() []NotificationsSettingsPhoneDeliverySettingsCustomNumbers`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) GetNumbersOk() (*[]NotificationsSettingsPhoneDeliverySettingsCustomNumbers, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) SetNumbers(v []NotificationsSettingsPhoneDeliverySettingsCustomNumbers)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOf) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


