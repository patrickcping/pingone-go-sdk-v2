# NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersType**](EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersType.md) |  | 
**Selected** | **bool** | Specifies whether the number is selected by the admin for sending messages. | 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Number** | **string** | The phone number, toll-free number or short code. | 
**Available** | **bool** | Specifies whether the number is currently available in the provider account. | 
**Capabilities** | [**[]EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersCapability**](EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersCapability.md) | A collection of the phone delivery service capabilities. | 
**SupportedCountries** | **[]string** | Specifies the number&#39;s supported countries for notification recipients, depending on the phone number &#x60;type&#x60;: - &#x60;SHORT_CODE&#x60;: A collection containing a single 2-character ISO country code, for example, &#x60;US&#x60;, &#x60;GB&#x60;, &#x60;CA&#x60;. If the custom provider is of &#x60;type&#x3D;CUSTOM_PROVIDER&#x60;, &#x60;supportedCountries&#x60; must not be empty or null. For other custom provider types, if supportedCountries is null (empty is not supported), the specified short code number can only be used to dispatch notifications to United States recipient numbers. - &#x60;TOLL_FREE&#x60;: A collection of valid 2-character country ISO codes, for example, &#x60;US&#x60;, &#x60;GB&#x60;, &#x60;CA&#x60;. If the custom provider is of &#x60;type&#x3D;CUSTOM_PROVIDER&#x60;, &#x60;supportedCountries&#x60; must not be empty or null. For other custom provider types, if &#x60;supportedCountries&#x60; is null (empty is not supported), the specified toll-free number can only be used to dispatch notifications to United States recipient numbers. - &#x60;PHONE_NUMBER&#x60;: &#x60;supportedCountries&#x60; can not be specified. If an SMS template has an alphanumeric sender ID and also has short code, the sender ID will be used for destination countries that support both alphanumeric senders and short codes. For Unites States and Canada that don&#39;t support alphanumeric sender IDs, a short code will be used if both an alphanumeric sender and a short code are specified.  | 

## Methods

### NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers

`func NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers(type_ EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersType, selected bool, number string, available bool, capabilities []EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersCapability, supportedCountries []string, ) *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers`

NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers instantiates a new NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbersWithDefaults

`func NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbersWithDefaults() *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers`

NewNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbersWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetType() EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetTypeOk() (*EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) SetType(v EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersType)`

SetType sets Type field to given value.


### GetSelected

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) SetSelected(v bool)`

SetSelected sets Selected field to given value.


### GetCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetNumber

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetAvailable

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetCapabilities

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetCapabilities() []EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetCapabilitiesOk() (*[]EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) SetCapabilities(v []EnumNotificationsSettingsPhoneDeliverySettingsTwilioSyniverseNumbersCapability)`

SetCapabilities sets Capabilities field to given value.


### GetSupportedCountries

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetSupportedCountries() []string`

GetSupportedCountries returns the SupportedCountries field if non-nil, zero value otherwise.

### GetSupportedCountriesOk

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) GetSupportedCountriesOk() (*[]string, bool)`

GetSupportedCountriesOk returns a tuple with the SupportedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCountries

`func (o *NotificationsSettingsPhoneDeliverySettingsTwilioSyniverseAllOfNumbers) SetSupportedCountries(v []string)`

SetSupportedCountries sets SupportedCountries field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


