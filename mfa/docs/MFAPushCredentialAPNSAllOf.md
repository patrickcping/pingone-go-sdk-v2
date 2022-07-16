# MFAPushCredentialAPNSAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TeamId** | **string** | A string that Apple uses as an identifier to identify teams. | 
**Token** | **string** | A string that Apple uses as the authentication token signing key to securely connect to APNS. This is a p8 file with a private key format. | 

## Methods

### NewMFAPushCredentialAPNSAllOf

`func NewMFAPushCredentialAPNSAllOf(teamId string, token string, ) *MFAPushCredentialAPNSAllOf`

NewMFAPushCredentialAPNSAllOf instantiates a new MFAPushCredentialAPNSAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMFAPushCredentialAPNSAllOfWithDefaults

`func NewMFAPushCredentialAPNSAllOfWithDefaults() *MFAPushCredentialAPNSAllOf`

NewMFAPushCredentialAPNSAllOfWithDefaults instantiates a new MFAPushCredentialAPNSAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeamId

`func (o *MFAPushCredentialAPNSAllOf) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *MFAPushCredentialAPNSAllOf) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *MFAPushCredentialAPNSAllOf) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetToken

`func (o *MFAPushCredentialAPNSAllOf) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *MFAPushCredentialAPNSAllOf) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *MFAPushCredentialAPNSAllOf) SetToken(v string)`

SetToken sets Token field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


