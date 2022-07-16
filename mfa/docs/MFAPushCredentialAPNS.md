# MFAPushCredentialAPNS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A string that specifies the type of the push credentials. Mandatory. Valid values APNS, FCM | 
**Key** | **string** | A string that Apple uses as an identifier to identify an authentication key.  Mandatory. | 
**TeamId** | **string** | A string that Apple uses as an identifier to identify teams. | 
**Token** | **string** | A string that Apple uses as the authentication token signing key to securely connect to APNS. This is a p8 file with a private key format. | 

## Methods

### NewMFAPushCredentialAPNS

`func NewMFAPushCredentialAPNS(type_ string, key string, teamId string, token string, ) *MFAPushCredentialAPNS`

NewMFAPushCredentialAPNS instantiates a new MFAPushCredentialAPNS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAPushCredentialAPNSWithDefaults

`func NewMFAPushCredentialAPNSWithDefaults() *MFAPushCredentialAPNS`

NewMFAPushCredentialAPNSWithDefaults instantiates a new MFAPushCredentialAPNS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MFAPushCredentialAPNS) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MFAPushCredentialAPNS) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MFAPushCredentialAPNS) SetType(v string)`

SetType sets Type field to given value.


### GetKey

`func (o *MFAPushCredentialAPNS) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MFAPushCredentialAPNS) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MFAPushCredentialAPNS) SetKey(v string)`

SetKey sets Key field to given value.


### GetTeamId

`func (o *MFAPushCredentialAPNS) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MFAPushCredentialAPNS) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MFAPushCredentialAPNS) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetToken

`func (o *MFAPushCredentialAPNS) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MFAPushCredentialAPNS) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MFAPushCredentialAPNS) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


