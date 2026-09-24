# NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod**](EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthMethod.md) |  | 
**Username** | Pointer to **string** | The username for the custom provider account. Required when &#x60;authentication.method&#x3D;BASIC&#x60; | [optional] 
**Password** | Pointer to **string** | The password for the custom provider account. Required when &#x60;authentication.method&#x3D;BASIC&#x60; | [optional] 
**AuthToken** | Pointer to **string** | The authentication token for the custom provider account.  Required when &#x60;authentication.method&#x3D;BEARER&#x60; | [optional] 
**AuthUrl** | Pointer to **string** | The URL of the authorization server that issues the access token for the custom provider account. Required when &#x60;authentication.method&#x3D;OAUTH2&#x60; | [optional] 
**GrantType** | Pointer to [**EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthGrantType**](EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthGrantType.md) |  | [optional] [default to ENUMNOTIFICATIONSSETTINGSPHONEDELIVERYSETTINGSCUSTOMAUTHGRANTTYPE_CLIENT_CREDENTIALS]
**Assertion** | Pointer to **string** | The JWT assertion used to request the access token from the authorization server. Must be a valid JWT. Required when &#x60;authentication.grantType&#x3D;JWT_BEARER&#x60; | [optional] 
**ClientId** | Pointer to **string** | The client ID used to request the access token from the authorization server. Required when &#x60;authentication.grantType&#x3D;CLIENT_CREDENTIALS&#x60; | [optional] 
**ClientSecret** | Pointer to **string** | The client secret used to request the access token from the authorization server. Required when &#x60;authentication.grantType&#x3D;CLIENT_CREDENTIALS&#x60; | [optional] 
**Scopes** | Pointer to **[]string** | An array of scopes to request in the access token from the authorization server, for example, &#x60;sms:send&#x60;, &#x60;voice:send&#x60;. | [optional] 
**HeaderName** | Pointer to **string** | The name of the custom header used to authenticate requests to the custom provider. Required when &#x60;authentication.method&#x3D;CUSTOM_HEADER&#x60; | [optional] 
**HeaderValue** | Pointer to **string** | The value of the custom header used to authenticate requests to the custom provider. Required when &#x60;authentication.method&#x3D;CUSTOM_HEADER&#x60; | [optional] 
**ClientAuthenticationMethod** | Pointer to [**EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthClientAuthenticationMethod**](EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthClientAuthenticationMethod.md) |  | [optional] 

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

### GetAuthToken

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetAuthToken() string`

GetAuthToken returns the AuthToken field if non-nil, zero value otherwise.

### GetAuthTokenOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetAuthTokenOk() (*string, bool)`

GetAuthTokenOk returns a tuple with the AuthToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthToken

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetAuthToken(v string)`

SetAuthToken sets AuthToken field to given value.

### HasAuthToken

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasAuthToken() bool`

HasAuthToken returns a boolean if a field has been set.

### GetAuthUrl

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.

### HasAuthUrl

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasAuthUrl() bool`

HasAuthUrl returns a boolean if a field has been set.

### GetGrantType

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetGrantType() EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthGrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetGrantTypeOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthGrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetGrantType(v EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthGrantType)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetAssertion

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetAssertion() string`

GetAssertion returns the Assertion field if non-nil, zero value otherwise.

### GetAssertionOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetAssertionOk() (*string, bool)`

GetAssertionOk returns a tuple with the Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertion

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetAssertion(v string)`

SetAssertion sets Assertion field to given value.

### HasAssertion

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasAssertion() bool`

HasAssertion returns a boolean if a field has been set.

### GetClientId

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScopes

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetHeaderName

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### GetHeaderValue

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### GetClientAuthenticationMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetClientAuthenticationMethod() EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthClientAuthenticationMethod`

GetClientAuthenticationMethod returns the ClientAuthenticationMethod field if non-nil, zero value otherwise.

### GetClientAuthenticationMethodOk

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) GetClientAuthenticationMethodOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthClientAuthenticationMethod, bool)`

GetClientAuthenticationMethodOk returns a tuple with the ClientAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) SetClientAuthenticationMethod(v EnumNotificationsSettingsPhoneDeliverySettingsCustomAuthClientAuthenticationMethod)`

SetClientAuthenticationMethod sets ClientAuthenticationMethod field to given value.

### HasClientAuthenticationMethod

`func (o *NotificationsSettingsPhoneDeliverySettingsCustomAllOfAuthentication) HasClientAuthenticationMethod() bool`

HasClientAuthenticationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


