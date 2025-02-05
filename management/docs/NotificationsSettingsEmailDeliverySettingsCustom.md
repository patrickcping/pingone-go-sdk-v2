# NotificationsSettingsEmailDeliverySettingsCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Protocol** | [**EnumNotificationsSettingsEmailDeliverySettingsProtocol**](EnumNotificationsSettingsEmailDeliverySettingsProtocol.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 
**Authentication** | [**NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication**](NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication.md) |  | 
**From** | Pointer to [**NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom**](NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom.md) |  | [optional] 
**Name** | **string** | Name to use to identify the provider. | 
**Provider** | Pointer to [**EnumNotificationsSettingsEmailDeliverySettingsCustomProvider**](EnumNotificationsSettingsEmailDeliverySettingsCustomProvider.md) |  | [optional] 
**ReplyTo** | Pointer to [**NotificationsSettingsEmailDeliverySettingsCustomAllOfReplyTo**](NotificationsSettingsEmailDeliverySettingsCustomAllOfReplyTo.md) |  | [optional] 
**Requests** | [**[]NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests**](NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests.md) | Contains the object that is used to configure the API requests sent to the email provider. | 

## Methods

### NewNotificationsSettingsEmailDeliverySettingsCustom

`func NewNotificationsSettingsEmailDeliverySettingsCustom(protocol EnumNotificationsSettingsEmailDeliverySettingsProtocol, authentication NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication, name string, requests []NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests, ) *NotificationsSettingsEmailDeliverySettingsCustom`

NewNotificationsSettingsEmailDeliverySettingsCustom instantiates a new NotificationsSettingsEmailDeliverySettingsCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsEmailDeliverySettingsCustomWithDefaults

`func NewNotificationsSettingsEmailDeliverySettingsCustomWithDefaults() *NotificationsSettingsEmailDeliverySettingsCustom`

NewNotificationsSettingsEmailDeliverySettingsCustomWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettingsCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetProtocol

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetProtocol() EnumNotificationsSettingsEmailDeliverySettingsProtocol`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetProtocolOk() (*EnumNotificationsSettingsEmailDeliverySettingsProtocol, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetProtocol(v EnumNotificationsSettingsEmailDeliverySettingsProtocol)`

SetProtocol sets Protocol field to given value.


### GetUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetAuthentication

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetAuthentication() NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication`

GetAuthentication returns the Authentication field if non-nil, zero value otherwise.

### GetAuthenticationOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetAuthenticationOk() (*NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication, bool)`

GetAuthenticationOk returns a tuple with the Authentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthentication

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetAuthentication(v NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication)`

SetAuthentication sets Authentication field to given value.


### GetFrom

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetFrom() NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetFromOk() (*NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetFrom(v NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetName

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetProvider() EnumNotificationsSettingsEmailDeliverySettingsCustomProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetProviderOk() (*EnumNotificationsSettingsEmailDeliverySettingsCustomProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetProvider(v EnumNotificationsSettingsEmailDeliverySettingsCustomProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetReplyTo

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetReplyTo() NotificationsSettingsEmailDeliverySettingsCustomAllOfReplyTo`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetReplyToOk() (*NotificationsSettingsEmailDeliverySettingsCustomAllOfReplyTo, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetReplyTo(v NotificationsSettingsEmailDeliverySettingsCustomAllOfReplyTo)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetRequests

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetRequests() []NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) GetRequestsOk() (*[]NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *NotificationsSettingsEmailDeliverySettingsCustom) SetRequests(v []NotificationsSettingsEmailDeliverySettingsCustomAllOfRequests)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


