# NotificationsSettingsPhoneDeliverySettingsCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Provider** | [**EnumNotificationsSettingsPhoneDeliverySettingsProvider**](EnumNotificationsSettingsPhoneDeliverySettingsProvider.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Name** | **string** | The customer provider&#39;s name. | 
**Requests** | [**[]NotificationsSettingsPhoneDeliverySettingsCustomRequest**](NotificationsSettingsPhoneDeliverySettingsCustomRequest.md) |  | 
**Authentication** | [**NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication**](NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication.md) |  | 
**Numbers** | Pointer to [**[]NotificationsSettingsPhoneDeliverySettingsCustomNumbers**](NotificationsSettingsPhoneDeliverySettingsCustomNumbers.md) |  | [optional] 

## Methods

### NewNotificationsSettingsPhoneDeliverySettingsCustom

`func NewNotificationsSettingsPhoneDeliverySettingsCustom(provider EnumNotificationsSettingsPhoneDeliverySettingsProvider, name string, requests []NotificationsSettingsPhoneDeliverySettingsCustomRequest, authentication NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication, ) *NotificationsSettingsPhoneDeliverySettingsCustom`

NewNotificationsSettingsPhoneDeliverySettingsCustom instantiates a new NotificationsSettingsPhoneDeliverySettingsCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsPhoneDeliverySettingsCustomWithDefaults

`func NewNotificationsSettingsPhoneDeliverySettingsCustomWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustom`

NewNotificationsSettingsPhoneDeliverySettingsCustomWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProvider

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetProvider() EnumNotificationsSettingsPhoneDeliverySettingsProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetProviderOk() (*EnumNotificationsSettingsPhoneDeliverySettingsProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetProvider(v EnumNotificationsSettingsPhoneDeliverySettingsProvider)`

SetProvider sets Provider field to given value.


### GetCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetName(v string)`

SetName sets Name field to given value.


### GetRequests

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetRequests() []NotificationsSettingsPhoneDeliverySettingsCustomRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetRequestsOk() (*[]NotificationsSettingsPhoneDeliverySettingsCustomRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetRequests(v []NotificationsSettingsPhoneDeliverySettingsCustomRequest)`

SetRequests sets Requests field to given value.


### GetAuthentication

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetAuthentication() NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetAuthenticationOk() (*NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetAuthentication(v NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication)`

SetAuthentication sets Authentication field to given value.


### GetNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetNumbers() []NotificationsSettingsPhoneDeliverySettingsCustomNumbers`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) GetNumbersOk() (*[]NotificationsSettingsPhoneDeliverySettingsCustomNumbers, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) SetNumbers(v []NotificationsSettingsPhoneDeliverySettingsCustomNumbers)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsCustom) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


