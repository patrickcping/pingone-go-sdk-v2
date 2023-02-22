# NotificationsSettingsEmailDeliverySettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | Pointer to **string** | A string that specifies the organization&#39;s SMTP server. | [optional] 
**Port** | Pointer to **int32** | An integer that specifies the port used by the organization&#39;s SMTP server to send emails (default &#x60;465&#x60;). Note that the protocol used depends upon the port specified. If you specify port &#x60;25&#x60;, &#x60;587&#x60;, or &#x60;2525&#x60;, SMTP with &#x60;STARTTLS&#x60; is used. Otherwise, &#x60;SMTPS&#x60; is used. | [optional] [default to 465]
**Protocol** | Pointer to **string** | A string that specifies the organization&#39;s SMTP server&#39;s protocol. | [optional] [readonly] 
**Username** | Pointer to **string** | A string that specifies the organization&#39;s SMTP server&#39;s username. | [optional] 
**Password** | Pointer to **string** | A string that specifies the organization&#39;s SMTP server&#39;s password. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**From** | Pointer to [**NotificationsSettingsEmailDeliverySettingsFrom**](NotificationsSettingsEmailDeliverySettingsFrom.md) |  | [optional] 
**ReplyTo** | Pointer to [**NotificationsSettingsEmailDeliverySettingsReplyTo**](NotificationsSettingsEmailDeliverySettingsReplyTo.md) |  | [optional] 

## Methods

### NewNotificationsSettingsEmailDeliverySettings

`func NewNotificationsSettingsEmailDeliverySettings() *NotificationsSettingsEmailDeliverySettings`

NewNotificationsSettingsEmailDeliverySettings instantiates a new NotificationsSettingsEmailDeliverySettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsEmailDeliverySettingsWithDefaults

`func NewNotificationsSettingsEmailDeliverySettingsWithDefaults() *NotificationsSettingsEmailDeliverySettings`

NewNotificationsSettingsEmailDeliverySettingsWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *NotificationsSettingsEmailDeliverySettings) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *NotificationsSettingsEmailDeliverySettings) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *NotificationsSettingsEmailDeliverySettings) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *NotificationsSettingsEmailDeliverySettings) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetPort

`func (o *NotificationsSettingsEmailDeliverySettings) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NotificationsSettingsEmailDeliverySettings) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NotificationsSettingsEmailDeliverySettings) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *NotificationsSettingsEmailDeliverySettings) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetProtocol

`func (o *NotificationsSettingsEmailDeliverySettings) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NotificationsSettingsEmailDeliverySettings) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NotificationsSettingsEmailDeliverySettings) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *NotificationsSettingsEmailDeliverySettings) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationsSettingsEmailDeliverySettings) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationsSettingsEmailDeliverySettings) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationsSettingsEmailDeliverySettings) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationsSettingsEmailDeliverySettings) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationsSettingsEmailDeliverySettings) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationsSettingsEmailDeliverySettings) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationsSettingsEmailDeliverySettings) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationsSettingsEmailDeliverySettings) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsSettingsEmailDeliverySettings) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsSettingsEmailDeliverySettings) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsSettingsEmailDeliverySettings) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsSettingsEmailDeliverySettings) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetFrom

`func (o *NotificationsSettingsEmailDeliverySettings) GetFrom() NotificationsSettingsEmailDeliverySettingsFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *NotificationsSettingsEmailDeliverySettings) GetFromOk() (*NotificationsSettingsEmailDeliverySettingsFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *NotificationsSettingsEmailDeliverySettings) SetFrom(v NotificationsSettingsEmailDeliverySettingsFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *NotificationsSettingsEmailDeliverySettings) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetReplyTo

`func (o *NotificationsSettingsEmailDeliverySettings) GetReplyTo() NotificationsSettingsEmailDeliverySettingsReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *NotificationsSettingsEmailDeliverySettings) GetReplyToOk() (*NotificationsSettingsEmailDeliverySettingsReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *NotificationsSettingsEmailDeliverySettings) SetReplyTo(v NotificationsSettingsEmailDeliverySettingsReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *NotificationsSettingsEmailDeliverySettings) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


