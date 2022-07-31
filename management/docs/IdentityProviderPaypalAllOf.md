# IdentityProviderPaypalAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | A string that specifies the application ID from PayPal. This is a required property. | 
**ClientSecret** | **string** | A string that specifies the application secret from PayPal. This is a required property. | 
**ClientEnvironment** | **string** | A string that specifies the PayPal environment. Options are sandbox, and live. This is a required property. | 

## Methods

### NewIdentityProviderPaypalAllOf

`func NewIdentityProviderPaypalAllOf(clientId string, clientSecret string, clientEnvironment string, ) *IdentityProviderPaypalAllOf`

NewIdentityProviderPaypalAllOf instantiates a new IdentityProviderPaypalAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderPaypalAllOfWithDefaults

`func NewIdentityProviderPaypalAllOfWithDefaults() *IdentityProviderPaypalAllOf`

NewIdentityProviderPaypalAllOfWithDefaults instantiates a new IdentityProviderPaypalAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IdentityProviderPaypalAllOf) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IdentityProviderPaypalAllOf) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IdentityProviderPaypalAllOf) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *IdentityProviderPaypalAllOf) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *IdentityProviderPaypalAllOf) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *IdentityProviderPaypalAllOf) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetClientEnvironment

`func (o *IdentityProviderPaypalAllOf) GetClientEnvironment() string`

GetClientEnvironment returns the ClientEnvironment field if non-nil, zero value otherwise.

### GetClientEnvironmentOk

`func (o *IdentityProviderPaypalAllOf) GetClientEnvironmentOk() (*string, bool)`

GetClientEnvironmentOk returns a tuple with the ClientEnvironment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientEnvironment

`func (o *IdentityProviderPaypalAllOf) SetClientEnvironment(v string)`

SetClientEnvironment sets ClientEnvironment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


