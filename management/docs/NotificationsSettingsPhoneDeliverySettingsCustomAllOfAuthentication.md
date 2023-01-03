# NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod**](EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod.md) |  | 
**Username** | Pointer to **string** | The username for the custom provider account. Required when &#x60;authentication.method&#x3D;BASIC&#x60; | [optional] 
**Password** | Pointer to **string** | The password for the custom provider account. Required when &#x60;authentication.method&#x3D;BASIC&#x60; | [optional] 
**Token** | Pointer to **string** | The authentication token for the custom provider account.  Required when &#x60;authentication.method&#x3D;BEARER&#x60; | [optional] 

## Methods

### NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication

`func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication(method EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod, ) *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication`

NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthenticationWithDefaults

`func NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthenticationWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication`

NewNotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthenticationWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetMethod() EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetMethodOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetMethod(v EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod)`

SetMethod sets Method field to given value.


### GetUsername

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetToken

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasToken() bool`

HasToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


