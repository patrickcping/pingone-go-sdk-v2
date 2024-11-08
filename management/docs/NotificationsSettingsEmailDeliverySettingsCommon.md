# NotificationsSettingsEmailDeliverySettingsCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewNotificationsSettingsEmailDeliverySettingsCommon

`func NewNotificationsSettingsEmailDeliverySettingsCommon() *NotificationsSettingsEmailDeliverySettingsCommon`

NewNotificationsSettingsEmailDeliverySettingsCommon instantiates a new NotificationsSettingsEmailDeliverySettingsCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsEmailDeliverySettingsCommonWithDefaults

`func NewNotificationsSettingsEmailDeliverySettingsCommonWithDefaults() *NotificationsSettingsEmailDeliverySettingsCommon`

NewNotificationsSettingsEmailDeliverySettingsCommonWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettingsCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsCommon) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


