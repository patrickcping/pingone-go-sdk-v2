# NotificationsSettingsEmailDeliverySettingsSMTP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Protocol** | Pointer to [**EnumNotificationsSettingsEmailDeliverySettingsProtocol**](EnumNotificationsSettingsEmailDeliverySettingsProtocol.md) |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Host** | Pointer to **string** | A string that specifies the organization&#39;s SMTP server. | [optional] 
**Port** | Pointer to **int32** | An integer that specifies the port used by the organization&#39;s SMTP server to send emails (default &#x60;465&#x60;). Note that the protocol used depends upon the port specified. If you specify port &#x60;25&#x60;, &#x60;587&#x60;, or &#x60;2525&#x60;, SMTP with &#x60;STARTTLS&#x60; is used. Otherwise, &#x60;SMTPS&#x60; is used. | [optional] [default to 465]
**Username** | Pointer to **string** | A string that specifies the organization&#39;s SMTP server&#39;s username. | [optional] 
**Password** | Pointer to **string** | A string that specifies the organization&#39;s SMTP server&#39;s password. | [optional] 
**From** | Pointer to [**NotificationsSettingsEmailDeliverySettingsSMTPAllOfFrom**](NotificationsSettingsEmailDeliverySettingsSMTPAllOfFrom.md) |  | [optional] 
**ReplyTo** | Pointer to [**NotificationsSettingsEmailDeliverySettingsSMTPAllOfReplyTo**](NotificationsSettingsEmailDeliverySettingsSMTPAllOfReplyTo.md) |  | [optional] 

## Methods

### NewNotificationsSettingsEmailDeliverySettingsSMTP

`func NewNotificationsSettingsEmailDeliverySettingsSMTP() *NotificationsSettingsEmailDeliverySettingsSMTP`

NewNotificationsSettingsEmailDeliverySettingsSMTP instantiates a new NotificationsSettingsEmailDeliverySettingsSMTP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsEmailDeliverySettingsSMTPWithDefaults

`func NewNotificationsSettingsEmailDeliverySettingsSMTPWithDefaults() *NotificationsSettingsEmailDeliverySettingsSMTP`

NewNotificationsSettingsEmailDeliverySettingsSMTPWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettingsSMTP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProtocol

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetProtocol() EnumNotificationsSettingsEmailDeliverySettingsProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetProtocolOk() (*EnumNotificationsSettingsEmailDeliverySettingsProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetProtocol(v EnumNotificationsSettingsEmailDeliverySettingsProtocol)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetHost

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetFrom

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetFrom() NotificationsSettingsEmailDeliverySettingsSMTPAllOfFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetFromOk() (*NotificationsSettingsEmailDeliverySettingsSMTPAllOfFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetFrom(v NotificationsSettingsEmailDeliverySettingsSMTPAllOfFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetReplyTo() NotificationsSettingsEmailDeliverySettingsSMTPAllOfReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) GetReplyToOk() (*NotificationsSettingsEmailDeliverySettingsSMTPAllOfReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) SetReplyTo(v NotificationsSettingsEmailDeliverySettingsSMTPAllOfReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *NotificationsSettingsEmailDeliverySettingsSMTP) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


