# NotificationsSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**DeliveryMode** | Pointer to [**EnumNotificationsSettingsDeliveryMode**](EnumNotificationsSettingsDeliveryMode.md) |  | [optional] 
**Restrictions** | Pointer to [**NotificationsSettingsRestrictions**](NotificationsSettingsRestrictions.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**SmsProvidersFallbackChain** | Pointer to **[]string** | A list which represents the execution order of different SMS/Voice providers configured for the environment. The providers and their accounts’ configurations are represented in the list by the ID of the corresponding &#x60;PhoneDeliverySettings&#x60; resource. The only provider which is not represented by a &#x60;phoneDeliverySettingsID&#x60; is the PingOne Twilio provider. The PingOne Twilio provider is represented by the &#x60;PINGONE_TWILIO&#x60; string. If the &#x60;smsProvidersFallbackChain&#x60; list is empty, an SMS or voice message will be sent using the default Ping Twilio account. Otherwise, an SMS or voice message will be sent using the first provider in the list. If the server fails to queue the message using that provider, it will use the next provider in the list to try to send the message. This process will go on until there are no more providers in the list. If the server failed to send the message using all providers, the notification status is set to &#x60;FAILED&#x60;. | [optional] 
**From** | Pointer to [**NotificationsSettingsFrom**](NotificationsSettingsFrom.md) |  | [optional] 
**ReplyTo** | Pointer to [**NotificationsSettingsReplyTo**](NotificationsSettingsReplyTo.md) |  | [optional] 
**Whitelist** | Pointer to [**[]NotificationsSettingsWhitelistInner**](NotificationsSettingsWhitelistInner.md) |  | [optional] 

## Methods

### NewNotificationsSettings

`func NewNotificationsSettings() *NotificationsSettings`

NewNotificationsSettings instantiates a new NotificationsSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsWithDefaults

`func NewNotificationsSettingsWithDefaults() *NotificationsSettings`

NewNotificationsSettingsWithDefaults instantiates a new NotificationsSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdatedAt

`func (o *NotificationsSettings) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsSettings) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsSettings) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsSettings) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsSettings) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsSettings) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsSettings) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsSettings) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetDeliveryMode

`func (o *NotificationsSettings) GetDeliveryMode() EnumNotificationsSettingsDeliveryMode`

GetDeliveryMode returns the DeliveryMode field if non-nil, zero value otherwise.

### GetDeliveryModeOk

`func (o *NotificationsSettings) GetDeliveryModeOk() (*EnumNotificationsSettingsDeliveryMode, bool)`

GetDeliveryModeOk returns a tuple with the DeliveryMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryMode

`func (o *NotificationsSettings) SetDeliveryMode(v EnumNotificationsSettingsDeliveryMode)`

SetDeliveryMode sets DeliveryMode field to given value.

### HasDeliveryMode

`func (o *NotificationsSettings) HasDeliveryMode() bool`

HasDeliveryMode returns a boolean if a field has been set.

### GetRestrictions

`func (o *NotificationsSettings) GetRestrictions() NotificationsSettingsRestrictions`

GetRestrictions returns the Restrictions field if non-nil, zero value otherwise.

### GetRestrictionsOk

`func (o *NotificationsSettings) GetRestrictionsOk() (*NotificationsSettingsRestrictions, bool)`

GetRestrictionsOk returns a tuple with the Restrictions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictions

`func (o *NotificationsSettings) SetRestrictions(v NotificationsSettingsRestrictions)`

SetRestrictions sets Restrictions field to given value.

### HasRestrictions

`func (o *NotificationsSettings) HasRestrictions() bool`

HasRestrictions returns a boolean if a field has been set.

### GetId

`func (o *NotificationsSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSmsProvidersFallbackChain

`func (o *NotificationsSettings) GetSmsProvidersFallbackChain() []string`

GetSmsProvidersFallbackChain returns the SmsProvidersFallbackChain field if non-nil, zero value otherwise.

### GetSmsProvidersFallbackChainOk

`func (o *NotificationsSettings) GetSmsProvidersFallbackChainOk() (*[]string, bool)`

GetSmsProvidersFallbackChainOk returns a tuple with the SmsProvidersFallbackChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmsProvidersFallbackChain

`func (o *NotificationsSettings) SetSmsProvidersFallbackChain(v []string)`

SetSmsProvidersFallbackChain sets SmsProvidersFallbackChain field to given value.

### HasSmsProvidersFallbackChain

`func (o *NotificationsSettings) HasSmsProvidersFallbackChain() bool`

HasSmsProvidersFallbackChain returns a boolean if a field has been set.

### GetFrom

`func (o *NotificationsSettings) GetFrom() NotificationsSettingsFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *NotificationsSettings) GetFromOk() (*NotificationsSettingsFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *NotificationsSettings) SetFrom(v NotificationsSettingsFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *NotificationsSettings) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *NotificationsSettings) GetReplyTo() NotificationsSettingsReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *NotificationsSettings) GetReplyToOk() (*NotificationsSettingsReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *NotificationsSettings) SetReplyTo(v NotificationsSettingsReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *NotificationsSettings) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetWhitelist

`func (o *NotificationsSettings) GetWhitelist() []NotificationsSettingsWhitelistInner`

GetWhitelist returns the Whitelist field if non-nil, zero value otherwise.

### GetWhitelistOk

`func (o *NotificationsSettings) GetWhitelistOk() (*[]NotificationsSettingsWhitelistInner, bool)`

GetWhitelistOk returns a tuple with the Whitelist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhitelist

`func (o *NotificationsSettings) SetWhitelist(v []NotificationsSettingsWhitelistInner)`

SetWhitelist sets Whitelist field to given value.

### HasWhitelist

`func (o *NotificationsSettings) HasWhitelist() bool`

HasWhitelist returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


