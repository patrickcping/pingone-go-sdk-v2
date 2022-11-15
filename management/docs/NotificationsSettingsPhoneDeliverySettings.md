# NotificationsSettingsPhoneDeliverySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Provider** | [**EnumNotificationsSettingsPhoneDeliverySettingsProvider**](EnumNotificationsSettingsPhoneDeliverySettingsProvider.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewNotificationsSettingsPhoneDeliverySettings

`func NewNotificationsSettingsPhoneDeliverySettings(provider EnumNotificationsSettingsPhoneDeliverySettingsProvider, ) *NotificationsSettingsPhoneDeliverySettings`

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


