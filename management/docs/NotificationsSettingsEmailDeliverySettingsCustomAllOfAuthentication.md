# NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Method** | [**EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationMethod**](EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationMethod.md) |  | 
**AuthToken** | Pointer to **string** | If you specified &#x60;BEARER&#x60; as the authentication method, use authToken to provide the bearer token to use. | [optional] 
**Username** | Pointer to **string** | If you specified &#x60;BASIC&#x60; as the authentication method, use username to provide the username for authenticating with the email provider. | [optional] 
**Password** | Pointer to **string** | If you specified &#x60;BASIC&#x60; as the authentication method, use password to provide the password for authenticating with the email provider. | [optional] 
**AuthUrl** | Pointer to **string** | The URL of the authorization server that issues the access token for the email provider. Required when &#x60;authentication.method&#x3D;OAUTH2&#x60; | [optional] 
**GrantType** | Pointer to [**EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationGrantType**](EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationGrantType.md) |  | [optional] [default to ENUMNOTIFICATIONSSETTINGSEMAILDELIVERYSETTINGSCUSTOMAUTHENTICATIONGRANTTYPE_CLIENT_CREDENTIALS]
**Assertion** | Pointer to **string** | The JWT assertion used to request the access token from the authorization server. Must be a valid JWT. Required when &#x60;authentication.grantType&#x3D;JWT_BEARER&#x60; | [optional] 
**ClientId** | Pointer to **string** | The client ID used to request the access token from the authorization server. Required when &#x60;authentication.grantType&#x3D;CLIENT_CREDENTIALS&#x60; | [optional] 
**ClientSecret** | Pointer to **string** | The client secret used to request the access token from the authorization server. Required when &#x60;authentication.grantType&#x3D;CLIENT_CREDENTIALS&#x60; | [optional] 
**Scopes** | Pointer to **[]string** | An array of scopes to request in the access token from the authorization server, for example, &#x60;mail.send&#x60;. | [optional] 
**HeaderName** | Pointer to **string** | The name of the custom header used to authenticate requests to the email provider. Required when &#x60;authentication.method&#x3D;CUSTOM_HEADER&#x60; | [optional] 
**HeaderValue** | Pointer to **string** | The value of the custom header used to authenticate requests to the email provider. Required when &#x60;authentication.method&#x3D;CUSTOM_HEADER&#x60; | [optional] 
**ClientAuthenticationMethod** | Pointer to [**EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationClientAuthenticationMethod**](EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationClientAuthenticationMethod.md) |  | [optional] 

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

### GetAuthUrl

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.

### HasAuthUrl

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasAuthUrl() bool`

HasAuthUrl returns a boolean if a field has been set.

### GetGrantType

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetGrantType() EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationGrantType`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetGrantTypeOk() (*EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationGrantType, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetGrantType(v EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationGrantType)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### GetAssertion

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetAssertion() string`

GetAssertion returns the Assertion field if non-nil, zero value otherwise.

### GetAssertionOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetAssertionOk() (*string, bool)`

GetAssertionOk returns a tuple with the Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertion

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetAssertion(v string)`

SetAssertion sets Assertion field to given value.

### HasAssertion

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasAssertion() bool`

HasAssertion returns a boolean if a field has been set.

### GetClientId

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetScopes

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetHeaderName

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetHeaderName() string`

GetHeaderName returns the HeaderName field if non-nil, zero value otherwise.

### GetHeaderNameOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetHeaderNameOk() (*string, bool)`

GetHeaderNameOk returns a tuple with the HeaderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderName

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetHeaderName(v string)`

SetHeaderName sets HeaderName field to given value.

### HasHeaderName

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasHeaderName() bool`

HasHeaderName returns a boolean if a field has been set.

### GetHeaderValue

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetHeaderValue() string`

GetHeaderValue returns the HeaderValue field if non-nil, zero value otherwise.

### GetHeaderValueOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetHeaderValueOk() (*string, bool)`

GetHeaderValueOk returns a tuple with the HeaderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderValue

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetHeaderValue(v string)`

SetHeaderValue sets HeaderValue field to given value.

### HasHeaderValue

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasHeaderValue() bool`

HasHeaderValue returns a boolean if a field has been set.

### GetClientAuthenticationMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetClientAuthenticationMethod() EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationClientAuthenticationMethod`

GetClientAuthenticationMethod returns the ClientAuthenticationMethod field if non-nil, zero value otherwise.

### GetClientAuthenticationMethodOk

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) GetClientAuthenticationMethodOk() (*EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationClientAuthenticationMethod, bool)`

GetClientAuthenticationMethodOk returns a tuple with the ClientAuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientAuthenticationMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) SetClientAuthenticationMethod(v EnumNotificationsSettingsEmailDeliverySettingsCustomAuthenticationClientAuthenticationMethod)`

SetClientAuthenticationMethod sets ClientAuthenticationMethod field to given value.

### HasClientAuthenticationMethod

`func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfAuthentication) HasClientAuthenticationMethod() bool`

HasClientAuthenticationMethod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


