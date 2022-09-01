# ApplicationSAMLAllOfIdpSigning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | [**ApplicationSAMLAllOfIdpSigningKey**](ApplicationSAMLAllOfIdpSigningKey.md) |  | 
**Algorithm** | Pointer to **string** | A string that specifies the IdP Signing key algorithm. | [optional] [readonly] 

## Methods

### NewApplicationSAMLAllOfIdpSigning

`func NewApplicationSAMLAllOfIdpSigning(key ApplicationSAMLAllOfIdpSigningKey, ) *ApplicationSAMLAllOfIdpSigning`

NewApplicationSAMLAllOfIdpSigning instantiates a new ApplicationSAMLAllOfIdpSigning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSAMLAllOfIdpSigningWithDefaults

`func NewApplicationSAMLAllOfIdpSigningWithDefaults() *ApplicationSAMLAllOfIdpSigning`

NewApplicationSAMLAllOfIdpSigningWithDefaults instantiates a new ApplicationSAMLAllOfIdpSigning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *ApplicationSAMLAllOfIdpSigning) GetKey() ApplicationSAMLAllOfIdpSigningKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationSAMLAllOfIdpSigning) GetKeyOk() (*ApplicationSAMLAllOfIdpSigningKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationSAMLAllOfIdpSigning) SetKey(v ApplicationSAMLAllOfIdpSigningKey)`

SetKey sets Key field to given value.


### GetAlgorithm

`func (o *ApplicationSAMLAllOfIdpSigning) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ApplicationSAMLAllOfIdpSigning) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ApplicationSAMLAllOfIdpSigning) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.

### HasAlgorithm

`func (o *ApplicationSAMLAllOfIdpSigning) HasAlgorithm() bool`

HasAlgorithm returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


