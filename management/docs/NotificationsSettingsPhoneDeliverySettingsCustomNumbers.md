# NotificationsSettingsPhoneDeliverySettingsCustomNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | **string** | The phone number, toll-free number or short code. | 
**Type** | [**EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType**](EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType.md) |  | 
**Selected** | Pointer to **bool** | Specifies whether the number is selected by the admin for sending messages. | [optional] 
**Available** | Pointer to **bool** | Specifies whether the number is currently available in the provider account. | [optional] 
**Capabilities** | Pointer to [**[]EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability**](EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability.md) | A collection of the phone delivery service capabilities. | [optional] 
**SupportedCountries** | Pointer to **[]string** | Specifies the &#x60;number&#x60;&#39;s supported countries for notification recipients, depending on the phone number type: &#x60;SHORT_CODE&#x60;: A collection containing a single 2-character ISO country code, for example, &#x60;US&#x60;, &#x60;GB&#x60;, &#x60;CA&#x60;. If the custom provider is of &#x60;type&#x3D;CUSTOM_PROVIDER&#x60;, &#x60;supportedCountries&#x60; must not be empty or null. For other custom provider types, if &#x60;supportedCountries&#x60; is null (empty is not supported), the specified short code number can only be used to dispatch notifications to United States recipient numbers. &#x60;TOLL_FREE&#x60;: A collection of valid 2-character country ISO codes, for example, &#x60;US&#x60;, &#x60;GB&#x60;, &#x60;CA&#x60;. If the custom provider is of &#x60;type&#x3D;CUSTOM_PROVIDER&#x60;, &#x60;supportedCountries&#x60; must not be empty or null. For other custom provider types, if &#x60;supportedCountries&#x60; is null (empty is not supported), the specified toll-free number can only be used to dispatch notifications to United States recipient numbers. &#x60;PHONE_NUMBER&#x60;: &#x60;supportedCountries&#x60; can not be specified. If an SMS template has an alphanumeric &#x60;sender&#x60; ID and also has short code, the &#x60;sender&#x60; ID will be used for destination countries that support both alphanumeric senders and short codes. For Unites States and Canada that don&#39;t support alphanumeric sender IDs, a short code will be used if both an alphanumeric sender and a short code are specified.  | [optional] 

## Methods

### NewNotificationsSettingsPhoneDeliverySettingsCustomNumbers

`func NewNotificationsSettingsPhoneDeliverySettingsCustomNumbers(number string, type_ EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType, ) *NotificationsSettingsPhoneDeliverySettingsCustomNumbers`

NewNotificationsSettingsPhoneDeliverySettingsCustomNumbers instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsPhoneDeliverySettingsCustomNumbersWithDefaults

`func NewNotificationsSettingsPhoneDeliverySettingsCustomNumbersWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustomNumbers`

NewNotificationsSettingsPhoneDeliverySettingsCustomNumbersWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetNumber(v string)`

SetNumber sets Number field to given value.


### GetType

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetType() EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetTypeOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetType(v EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType)`

SetType sets Type field to given value.


### GetSelected

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetSelected() bool`

GetSelected returns the Selected field if non-nil, zero value otherwise.

### GetSelectedOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetSelectedOk() (*bool, bool)`

GetSelectedOk returns a tuple with the Selected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelected

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetSelected(v bool)`

SetSelected sets Selected field to given value.

### HasSelected

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) HasSelected() bool`

HasSelected returns a boolean if a field has been set.

### GetAvailable

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetCapabilities

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetCapabilities() []EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetCapabilitiesOk() (*[]EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetCapabilities(v []EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetSupportedCountries

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetSupportedCountries() []string`

GetSupportedCountries returns the SupportedCountries field if non-nil, zero value otherwise.

### GetSupportedCountriesOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetSupportedCountriesOk() (*[]string, bool)`

GetSupportedCountriesOk returns a tuple with the SupportedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedCountries

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetSupportedCountries(v []string)`

SetSupportedCountries sets SupportedCountries field to given value.

### HasSupportedCountries

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) HasSupportedCountries() bool`

HasSupportedCountries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


