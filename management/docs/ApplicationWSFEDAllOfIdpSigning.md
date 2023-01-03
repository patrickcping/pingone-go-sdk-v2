# ApplicationWSFEDAllOfIdpSigning

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Algorithm** | **string** | The signature algorithm to be used for signing. Algorithms upported RSA_SHA 256, 384, 512, and ECDSA_SHA 1, 224, 256, 384, 512. | 
**Key** | [**ApplicationWSFEDAllOfIdpSigningKey**](ApplicationWSFEDAllOfIdpSigningKey.md) |  | 

## Methods

### NewApplicationWSFEDAllOfIdpSigning

`func NewApplicationWSFEDAllOfIdpSigning(algorithm string, key ApplicationWSFEDAllOfIdpSigningKey, ) *ApplicationWSFEDAllOfIdpSigning`

NewApplicationWSFEDAllOfIdpSigning instantiates a new ApplicationWSFEDAllOfIdpSigning object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWSFEDAllOfIdpSigningWithDefaults

`func NewApplicationWSFEDAllOfIdpSigningWithDefaults() *ApplicationWSFEDAllOfIdpSigning`

NewApplicationWSFEDAllOfIdpSigningWithDefaults instantiates a new ApplicationWSFEDAllOfIdpSigning object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlgorithm

`func (o *ApplicationWSFEDAllOfIdpSigning) GetAlgorithm() string`

GetAlgorithm returns the Algorithm field if non-nil, zero value otherwise.

### GetAlgorithmOk

`func (o *ApplicationWSFEDAllOfIdpSigning) GetAlgorithmOk() (*string, bool)`

GetAlgorithmOk returns a tuple with the Algorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlgorithm

`func (o *ApplicationWSFEDAllOfIdpSigning) SetAlgorithm(v string)`

SetAlgorithm sets Algorithm field to given value.


### GetKey

`func (o *ApplicationWSFEDAllOfIdpSigning) GetKey() ApplicationWSFEDAllOfIdpSigningKey`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ApplicationWSFEDAllOfIdpSigning) GetKeyOk() (*ApplicationWSFEDAllOfIdpSigningKey, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ApplicationWSFEDAllOfIdpSigning) SetKey(v ApplicationWSFEDAllOfIdpSigningKey)`

SetKey sets Key field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


