# NotificationsSettingsPhoneDeliverySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Provider** | [**EnumNotificationsSettingsPhoneDeliverySettingsProvider**](EnumNotificationsSettingsPhoneDeliverySettingsProvider.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Name** | **string** | The customer provider&#39;s name. | 
**Requests** | [**[]NotificationsSettingsPhoneDeliverySettingsCustomRequest**](NotificationsSettingsPhoneDeliverySettingsCustomRequest.md) |  | 
**Authentication** | [**NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication**](NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication.md) |  | 
**Numbers** | Pointer to [**[]NotificationsSettingsPhoneDeliverySettingsCustomNumbers**](NotificationsSettingsPhoneDeliverySettingsCustomNumbers.md) |  | [optional] 
**Sid** | **string** | The public ID of the Twilio account. Relevant to Twilio only.  | 
**AuthToken** | **string** | The secret key of the Twilio or Syniverse account. | 

## Methods

### NewNotificationsSettingsPhoneDeliverySettings

`func NewNotificationsSettingsPhoneDeliverySettings(provider EnumNotificationsSettingsPhoneDeliverySettingsProvider, name string, requests []NotificationsSettingsPhoneDeliverySettingsCustomRequest, authentication NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication, sid string, authToken string, ) *NotificationsSettingsPhoneDeliverySettings`

NewNotificationsSettingsPhoneDeliverySettings instantiates a new NotificationsSettingsPhoneDeliverySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsPhoneDeliverySettingsWithDefaults

`func NewNotificationsSettingsPhoneDeliverySettingsWithDefaults() *NotificationsSettingsPhoneDeliverySettings`

NewNotificationsSettingsPhoneDeliverySettingsWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NotificationsSettingsPhoneDeliverySettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsSettingsPhoneDeliverySettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsSettingsPhoneDeliverySettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettings) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettings) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettings) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProvider

`func (o *NotificationsSettingsPhoneDeliverySettings) GetProvider() EnumNotificationsSettingsPhoneDeliverySettingsProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetProviderOk() (*EnumNotificationsSettingsPhoneDeliverySettingsProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationsSettingsPhoneDeliverySettings) SetProvider(v EnumNotificationsSettingsPhoneDeliverySettingsProvider)`

SetProvider sets Provider field to given value.


### GetCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettings) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettings) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettings) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettings) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettings) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettings) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *NotificationsSettingsPhoneDeliverySettings) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationsSettingsPhoneDeliverySettings) SetName(v string)`

SetName sets Name field to given value.


### GetRequests

`func (o *NotificationsSettingsPhoneDeliverySettings) GetRequests() []NotificationsSettingsPhoneDeliverySettingsCustomRequest`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetRequestsOk() (*[]NotificationsSettingsPhoneDeliverySettingsCustomRequest, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *NotificationsSettingsPhoneDeliverySettings) SetRequests(v []NotificationsSettingsPhoneDeliverySettingsCustomRequest)`

SetRequests sets Requests field to given value.


### GetAuthentication

`func (o *NotificationsSettingsPhoneDeliverySettings) GetAuthentication() NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetAuthenticationOk() (*NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NotificationsSettingsPhoneDeliverySettings) SetAuthentication(v NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication)`

SetAuthentication sets Authentication field to given value.


### GetNumbers

`func (o *NotificationsSettingsPhoneDeliverySettings) GetNumbers() []NotificationsSettingsPhoneDeliverySettingsCustomNumbers`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetNumbersOk() (*[]NotificationsSettingsPhoneDeliverySettingsCustomNumbers, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *NotificationsSettingsPhoneDeliverySettings) SetNumbers(v []NotificationsSettingsPhoneDeliverySettingsCustomNumbers)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *NotificationsSettingsPhoneDeliverySettings) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetSid

`func (o *NotificationsSettingsPhoneDeliverySettings) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NotificationsSettingsPhoneDeliverySettings) SetSid(v string)`

SetSid sets Sid field to given value.


### GetAuthToken

`func (o *NotificationsSettingsPhoneDeliverySettings) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *NotificationsSettingsPhoneDeliverySettings) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *NotificationsSettingsPhoneDeliverySettings) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


