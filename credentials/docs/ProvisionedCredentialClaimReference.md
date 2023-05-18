# ProvisionedCredentialClaimReference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **int32** |  | [optional] 
**Issuer** | Pointer to **string** |  | [optional] 
**Subject** | Pointer to **string** |  | [optional] 
**Holder** | Pointer to **string** |  | [optional] 
**CreateDate** | Pointer to **time.Time** |  | [optional] [readonly] 
**DataJson** | Pointer to **string** |  | [optional] 
**DataSignature** | Pointer to **string** |  | [optional] 
**DataHash** | Pointer to **string** |  | [optional] 
**PartitionId** | Pointer to **string** |  | [optional] 

## Methods

### NewProvisionedCredentialClaimReference

`func NewProvisionedCredentialClaimReference() *ProvisionedCredentialClaimReference`

NewProvisionedCredentialClaimReference instantiates a new ProvisionedCredentialClaimReference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProvisionedCredentialClaimReferenceWithDefaults

`func NewProvisionedCredentialClaimReferenceWithDefaults() *ProvisionedCredentialClaimReference`

NewProvisionedCredentialClaimReferenceWithDefaults instantiates a new ProvisionedCredentialClaimReference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProvisionedCredentialClaimReference) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProvisionedCredentialClaimReference) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProvisionedCredentialClaimReference) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProvisionedCredentialClaimReference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVersion

`func (o *ProvisionedCredentialClaimReference) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ProvisionedCredentialClaimReference) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ProvisionedCredentialClaimReference) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ProvisionedCredentialClaimReference) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetIssuer

`func (o *ProvisionedCredentialClaimReference) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *ProvisionedCredentialClaimReference) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *ProvisionedCredentialClaimReference) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.

### HasIssuer

`func (o *ProvisionedCredentialClaimReference) HasIssuer() bool`

HasIssuer returns a boolean if a field has been set.

### GetSubject

`func (o *ProvisionedCredentialClaimReference) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *ProvisionedCredentialClaimReference) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *ProvisionedCredentialClaimReference) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *ProvisionedCredentialClaimReference) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetHolder

`func (o *ProvisionedCredentialClaimReference) GetHolder() string`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *ProvisionedCredentialClaimReference) GetHolderOk() (*string, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *ProvisionedCredentialClaimReference) SetHolder(v string)`

SetHolder sets Holder field to given value.

### HasHolder

`func (o *ProvisionedCredentialClaimReference) HasHolder() bool`

HasHolder returns a boolean if a field has been set.

### GetCreateDate

`func (o *ProvisionedCredentialClaimReference) GetCreateDate() time.Time`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *ProvisionedCredentialClaimReference) GetCreateDateOk() (*time.Time, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *ProvisionedCredentialClaimReference) SetCreateDate(v time.Time)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *ProvisionedCredentialClaimReference) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetDataJson

`func (o *ProvisionedCredentialClaimReference) GetDataJson() string`

GetDataJson returns the DataJson field if non-nil, zero value otherwise.

### GetDataJsonOk

`func (o *ProvisionedCredentialClaimReference) GetDataJsonOk() (*string, bool)`

GetDataJsonOk returns a tuple with the DataJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataJson

`func (o *ProvisionedCredentialClaimReference) SetDataJson(v string)`

SetDataJson sets DataJson field to given value.

### HasDataJson

`func (o *ProvisionedCredentialClaimReference) HasDataJson() bool`

HasDataJson returns a boolean if a field has been set.

### GetDataSignature

`func (o *ProvisionedCredentialClaimReference) GetDataSignature() string`

GetDataSignature returns the DataSignature field if non-nil, zero value otherwise.

### GetDataSignatureOk

`func (o *ProvisionedCredentialClaimReference) GetDataSignatureOk() (*string, bool)`

GetDataSignatureOk returns a tuple with the DataSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataSignature

`func (o *ProvisionedCredentialClaimReference) SetDataSignature(v string)`

SetDataSignature sets DataSignature field to given value.

### HasDataSignature

`func (o *ProvisionedCredentialClaimReference) HasDataSignature() bool`

HasDataSignature returns a boolean if a field has been set.

### GetDataHash

`func (o *ProvisionedCredentialClaimReference) GetDataHash() string`

GetDataHash returns the DataHash field if non-nil, zero value otherwise.

### GetDataHashOk

`func (o *ProvisionedCredentialClaimReference) GetDataHashOk() (*string, bool)`

GetDataHashOk returns a tuple with the DataHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataHash

`func (o *ProvisionedCredentialClaimReference) SetDataHash(v string)`

SetDataHash sets DataHash field to given value.

### HasDataHash

`func (o *ProvisionedCredentialClaimReference) HasDataHash() bool`

HasDataHash returns a boolean if a field has been set.

### GetPartitionId

`func (o *ProvisionedCredentialClaimReference) GetPartitionId() string`

GetPartitionId returns the PartitionId field if non-nil, zero value otherwise.

### GetPartitionIdOk

`func (o *ProvisionedCredentialClaimReference) GetPartitionIdOk() (*string, bool)`

GetPartitionIdOk returns a tuple with the PartitionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionId

`func (o *ProvisionedCredentialClaimReference) SetPartitionId(v string)`

SetPartitionId sets PartitionId field to given value.

### HasPartitionId

`func (o *ProvisionedCredentialClaimReference) HasPartitionId() bool`

HasPartitionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


