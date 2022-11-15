# NotificationsSettingsFrom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A string that specifies the email&#39;s \&quot;from\&quot; name (relevant when the &#x60;deliveryMethod&#x60; is &#x60;Email&#x60;).  See [Note](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-settings-from-replyTo-note) for details. | [optional] 
**Address** | Pointer to **string** | A string that specifies the email&#39;s \&quot;from\&quot; address (relevant when the &#x60;deliveryMethod&#x60; is &#x60;Email&#x60;). This value must be a trusted email address.  See [Note](https://apidocs.pingidentity.com/pingone/platform/v1/api/#notifications-settings-from-replyTo-note) for details. | [optional] 

## Methods

### NewNotificationsSettingsFrom

`func NewNotificationsSettingsFrom() *NotificationsSettingsFrom`

NewNotificationsSettingsFrom instantiates a new NotificationsSettingsFrom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsFromWithDefaults

`func NewNotificationsSettingsFromWithDefaults() *NotificationsSettingsFrom`

NewNotificationsSettingsFromWithDefaults instantiates a new NotificationsSettingsFrom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NotificationsSettingsFrom) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NotificationsSettingsFrom) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NotificationsSettingsFrom) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NotificationsSettingsFrom) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *NotificationsSettingsFrom) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *NotificationsSettingsFrom) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *NotificationsSettingsFrom) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *NotificationsSettingsFrom) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


