# IdentityProviderAppleAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | A string that specifies the application ID from Apple. This is the identifier obtained after registering a services ID in the Apple developer portal. This is a required property. | 
**ClientSecretSigningKey** | **string** | A string that specifies the private key that is used to generate a client secret. This is a required property. | 
**KeyId** | **string** | A 10-character string that Apple uses to identify an authentication key. This is a required property. | 
**TeamId** | **string** | A 10-character string that Apple uses to identify teams. This is a required property. | 

## Methods

### NewIdentityProviderAppleAllOf

`func NewIdentityProviderAppleAllOf(clientId string, clientSecretSigningKey string, keyId string, teamId string, ) *IdentityProviderAppleAllOf`

NewIdentityProviderAppleAllOf instantiates a new IdentityProviderAppleAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderAppleAllOfWithDefaults

`func NewIdentityProviderAppleAllOfWithDefaults() *IdentityProviderAppleAllOf`

NewIdentityProviderAppleAllOfWithDefaults instantiates a new IdentityProviderAppleAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IdentityProviderAppleAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityProviderAppleAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityProviderAppleAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecretSigningKey

`func (o *IdentityProviderAppleAllOf) GetClientSecretSigningKey() string`

GetClientSecretSigningKey returns the ClientSecretSigningKey field if non-nil, zero value otherwise.

### GetClientSecretSigningKeyOk

`func (o *IdentityProviderAppleAllOf) GetClientSecretSigningKeyOk() (*string, bool)`

GetClientSecretSigningKeyOk returns a tuple with the ClientSecretSigningKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecretSigningKey

`func (o *IdentityProviderAppleAllOf) SetClientSecretSigningKey(v string)`

SetClientSecretSigningKey sets ClientSecretSigningKey field to given value.


### GetKeyId

`func (o *IdentityProviderAppleAllOf) GetKeyId() string`

GetKeyId returns the KeyId field if non-nil, zero value otherwise.

### GetKeyIdOk

`func (o *IdentityProviderAppleAllOf) GetKeyIdOk() (*string, bool)`

GetKeyIdOk returns a tuple with the KeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyId

`func (o *IdentityProviderAppleAllOf) SetKeyId(v string)`

SetKeyId sets KeyId field to given value.


### GetTeamId

`func (o *IdentityProviderAppleAllOf) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *IdentityProviderAppleAllOf) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *IdentityProviderAppleAllOf) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


