# MFAPushCredentialRequest

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

### NewMFAPushCredentialRequest

`func NewMFAPushCredentialRequest(type_ EnumMFAPushCredentialAttrType, key string, teamId string, token string, googleServiceAccountCredentials string, clientId string, clientSecret string, ) *MFAPushCredentialRequest`

NewMFAPushCredentialRequest instantiates a new MFAPushCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAPushCredentialRequestWithDefaults

`func NewMFAPushCredentialRequestWithDefaults() *MFAPushCredentialRequest`

NewMFAPushCredentialRequestWithDefaults instantiates a new MFAPushCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MFAPushCredentialRequest) GetType() EnumMFAPushCredentialAttrType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MFAPushCredentialRequest) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MFAPushCredentialRequest) SetType(v EnumMFAPushCredentialAttrType)`

SetType sets Type field to given value.


### GetKey

`func (o *MFAPushCredentialRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MFAPushCredentialRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MFAPushCredentialRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetTeamId

`func (o *MFAPushCredentialRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MFAPushCredentialRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MFAPushCredentialRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetToken

`func (o *MFAPushCredentialRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MFAPushCredentialRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MFAPushCredentialRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetGoogleServiceAccountCredentials

`func (o *MFAPushCredentialRequest) GetGoogleServiceAccountCredentials() string`

GetGoogleServiceAccountCredentials returns the GoogleServiceAccountCredentials field if non-nil, zero value otherwise.

### GetGoogleServiceAccountCredentialsOk

`func (o *MFAPushCredentialRequest) GetGoogleServiceAccountCredentialsOk() (*string, bool)`

GetGoogleServiceAccountCredentialsOk returns a tuple with the GoogleServiceAccountCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleServiceAccountCredentials

`func (o *MFAPushCredentialRequest) SetGoogleServiceAccountCredentials(v string)`

SetGoogleServiceAccountCredentials sets GoogleServiceAccountCredentials field to given value.


### GetClientId

`func (o *MFAPushCredentialRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *MFAPushCredentialRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *MFAPushCredentialRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *MFAPushCredentialRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *MFAPushCredentialRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *MFAPushCredentialRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


