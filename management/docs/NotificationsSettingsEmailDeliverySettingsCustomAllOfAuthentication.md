# NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationMethod**](EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationMethod.md) |  | 
**AuthToken** | Pointer to **string** | If you specified &#x60;BEARER&#x60; as the authentication method, use authToken to provide the bearer token to use. | [optional] 
**Username** | Pointer to **string** | If you specified &#x60;BASIC&#x60; as the authentication method, use username to provide the username for authenticating with the email provider. | [optional] 
**Password** | Pointer to **string** | If you specified &#x60;BASIC&#x60; as the authentication method, use password to provide the password for authenticating with the email provider. | [optional] 

## Methods

### NewNotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication

`func NewNotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication(method EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationMethod, ) *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication`

NewNotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication instantiates a new NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotificationsSettingsEmailDeliverySettingsCustomAllOfAuthenticationWithDefaults

`func NewNotificationsSettingsEmailDeliverySettingsCustomAllOfAuthenticationWithDefaults() *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication`

NewNotificationsSettingsEmailDeliverySettingsCustomAllOfAuthenticationWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetMethod() EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationMethod`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetMethodOk() (*EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationMethod, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetMethod(v EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationMethod)`

SetMethod sets Method field to given value.


### GetAuthToken

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetUsername

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


