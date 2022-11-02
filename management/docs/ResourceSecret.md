# ResourceSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Secret** | Pointer to **string** | An auto-generated resource client secret. Possible characters are a-z, A-Z, 0-9, -, ., _, ~. The secret has a minimum length of 64 characters per SHA-512 requirements when using the HS512 algorithm to sign ID tokens using the secret as the key. | [optional] [readonly] 

## Methods

### NewResourceSecret

`func NewResourceSecret() *ResourceSecret`

NewResourceSecret instantiates a new ResourceSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceSecretWithDefaults

`func NewResourceSecretWithDefaults() *ResourceSecret`

NewResourceSecretWithDefaults instantiates a new ResourceSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvironment

`func (o *ResourceSecret) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceSecret) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceSecret) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ResourceSecret) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSecret

`func (o *ResourceSecret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ResourceSecret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ResourceSecret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ResourceSecret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


