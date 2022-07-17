# CreateMFAPushCredentialRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**EnumMFAPushCredentialAttrType**](EnumMFAPushCredentialAttrType.md) |  | 
**Key** | **string** | A string that Apple uses as an identifier to identify an authentication key.  Mandatory. | 
**TeamId** | **string** | A string that Apple uses as an identifier to identify teams. | 
**Token** | **string** | A string that Apple uses as the authentication token signing key to securely connect to APNS. This is a p8 file with a private key format. | 

## Methods

### NewCreateMFAPushCredentialRequest

`func NewCreateMFAPushCredentialRequest(type_ EnumMFAPushCredentialAttrType, key string, teamId string, token string, ) *CreateMFAPushCredentialRequest`

NewCreateMFAPushCredentialRequest instantiates a new CreateMFAPushCredentialRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMFAPushCredentialRequestWithDefaults

`func NewCreateMFAPushCredentialRequestWithDefaults() *CreateMFAPushCredentialRequest`

NewCreateMFAPushCredentialRequestWithDefaults instantiates a new CreateMFAPushCredentialRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CreateMFAPushCredentialRequest) GetType() EnumMFAPushCredentialAttrType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMFAPushCredentialRequest) GetTypeOk() (*EnumMFAPushCredentialAttrType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMFAPushCredentialRequest) SetType(v EnumMFAPushCredentialAttrType)`

SetType sets Type field to given value.


### GetKey

`func (o *CreateMFAPushCredentialRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *CreateMFAPushCredentialRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *CreateMFAPushCredentialRequest) SetKey(v string)`

SetKey sets Key field to given value.


### GetTeamId

`func (o *CreateMFAPushCredentialRequest) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CreateMFAPushCredentialRequest) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CreateMFAPushCredentialRequest) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetToken

`func (o *CreateMFAPushCredentialRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CreateMFAPushCredentialRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CreateMFAPushCredentialRequest) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


