# IdentityProviderSAMLAllOfSpSigning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | Pointer to [**EnumIdentityProviderSAMLSigningAlgorithm**](EnumIdentityProviderSAMLSigningAlgorithm.md) |  | [optional] 
**Key** | [**IdentityProviderSAMLAllOfSpSigningKey**](IdentityProviderSAMLAllOfSpSigningKey.md) |  | 

## Methods

### NewIdentityProviderSAMLAllOfSpSigning

`func NewIdentityProviderSAMLAllOfSpSigning(key IdentityProviderSAMLAllOfSpSigningKey, ) *IdentityProviderSAMLAllOfSpSigning`

NewIdentityProviderSAMLAllOfSpSigning instantiates a new IdentityProviderSAMLAllOfSpSigning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderSAMLAllOfSpSigningWithDefaults

`func NewIdentityProviderSAMLAllOfSpSigningWithDefaults() *IdentityProviderSAMLAllOfSpSigning`

NewIdentityProviderSAMLAllOfSpSigningWithDefaults instantiates a new IdentityProviderSAMLAllOfSpSigning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *IdentityProviderSAMLAllOfSpSigning) GetAlgorithm() EnumIdentityProviderSAMLSigningAlgorithm`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *IdentityProviderSAMLAllOfSpSigning) GetAlgorithmOk() (*EnumIdentityProviderSAMLSigningAlgorithm, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *IdentityProviderSAMLAllOfSpSigning) SetAlgorithm(v EnumIdentityProviderSAMLSigningAlgorithm)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *IdentityProviderSAMLAllOfSpSigning) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.

### GetKey

`func (o *IdentityProviderSAMLAllOfSpSigning) GetKey() IdentityProviderSAMLAllOfSpSigningKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IdentityProviderSAMLAllOfSpSigning) GetKeyOk() (*IdentityProviderSAMLAllOfSpSigningKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IdentityProviderSAMLAllOfSpSigning) SetKey(v IdentityProviderSAMLAllOfSpSigningKey)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


