# NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Provider** | [**EnumNotificationsSettingsPhoneDeliverySettingsProvider**](EnumNotificationsSettingsPhoneDeliverySettingsProvider.md) |  | 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Sid** | **string** | The public ID of the Twilio account. Relevant to Twilio only.  | 
**AuthToken** | **string** | The secret key of the Twilio or Syniverse account. | 
**Numbers** | Pointer to [**[]NotificationsSettingsPhoneDeliverySettingsCustomNumbers**](NotificationsSettingsPhoneDeliverySettingsCustomNumbers.md) |  | [optional] 

## Methods

### NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse

`func NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse(provider EnumNotificationsSettingsPhoneDeliverySettingsProvider, sid string, authToken string, ) *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse`

NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverse instantiates a new NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseWithDefaults

`func NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseWithDefaults() *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse`

NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProvider

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetProvider() EnumNotificationsSettingsPhoneDeliverySettingsProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetProviderOk() (*EnumNotificationsSettingsPhoneDeliverySettingsProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetProvider(v EnumNotificationsSettingsPhoneDeliverySettingsProvider)`

SetProvider sets Provider field to given value.


### GetCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetSid

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetSid(v string)`

SetSid sets Sid field to given value.


### GetAuthToken

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.


### GetNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetNumbers() []NotificationsSettingsPhoneDeliverySettingsCustomNumbers`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) GetNumbersOk() (*[]NotificationsSettingsPhoneDeliverySettingsCustomNumbers, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) SetNumbers(v []NotificationsSettingsPhoneDeliverySettingsCustomNumbers)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverse) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


