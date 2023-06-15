# EntityArrayEmbeddedPushCredentialsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumMFAPushCredentialAttrType**](EnumMFAPushCredentialAttrType.md) |  | 
**Key** | **string** | A string that Apple uses as an identifier to identify an authentication key.  Mandatory. | 
**TeamId** | **string** | A string that Apple uses as an identifier to identify teams. | 
**Token** | **string** | A string that Apple uses as the authentication token signing key to securely connect to APNS. This is a p8 file with a private key format. | 
**GoogleServiceAccountCredentials** | **string** | Used when &#x60;type&#x60; is set to &#x60;FCM_HTTP_V1&#x60;. The value should be the contents of the JSON file that represents your Service Account Credentials. | 
**ClientId** | **string** | Used only if type is set to HMS. OAuth 2.0 Client ID from the Huawei Developers API console. | 
**ClientSecret** | **string** | Used only if type is set to HMS. The client secret associated with the OAuth 2.0 Client ID. | 

## Methods

### NewEntityArrayEmbeddedPushCredentialsInner

`func NewEntityArrayEmbeddedPushCredentialsInner(type_ EnumMFAPushCredentialAttrType, key string, teamId string, token string, googleServiceAccountCredentials string, clientId string, clientSecret string, ) *EntityArrayEmbeddedPushCredentialsInner`

NewEntityArrayEmbeddedPushCredentialsInner instantiates a new EntityArrayEmbeddedPushCredentialsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedPushCredentialsInnerWithDefaults

`func NewEntityArrayEmbeddedPushCredentialsInnerWithDefaults() *EntityArrayEmbeddedPushCredentialsInner`

NewEntityArrayEmbeddedPushCredentialsInnerWithDefaults instantiates a new EntityArrayEmbeddedPushCredentialsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetType() EnumMFAPushCredentialAttrType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityArrayEmbeddedPushCredentialsInner) SetType(v EnumMFAPushCredentialAttrType)`

SetType sets Type field to given value.


### GetKey

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EntityArrayEmbeddedPushCredentialsInner) SetKey(v string)`

SetKey sets Key field to given value.


### GetTeamId

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *EntityArrayEmbeddedPushCredentialsInner) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetToken

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *EntityArrayEmbeddedPushCredentialsInner) SetToken(v string)`

SetToken sets Token field to given value.


### GetGoogleServiceAccountCredentials

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetGoogleServiceAccountCredentials() string`

GetGoogleServiceAccountCredentials returns the GoogleServiceAccountCredentials field if non-nil, zero value otherwise.

### GetGoogleServiceAccountCredentialsOk

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetGoogleServiceAccountCredentialsOk() (*string, bool)`

GetGoogleServiceAccountCredentialsOk returns a tuple with the GoogleServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleServiceAccountCredentials

`func (o *EntityArrayEmbeddedPushCredentialsInner) SetGoogleServiceAccountCredentials(v string)`

SetGoogleServiceAccountCredentials sets GoogleServiceAccountCredentials field to given value.


### GetClientId

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *EntityArrayEmbeddedPushCredentialsInner) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *EntityArrayEmbeddedPushCredentialsInner) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *EntityArrayEmbeddedPushCredentialsInner) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


