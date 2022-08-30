# GatewayCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the auto-generated ID for this credential. This is the JWT&#39;s jti claim. This is a required property. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | A date that specifies the date the credential was created in Coordinated Universal Time (UTC). This is a required property. | [optional] [readonly] 
**LastUsedAt** | Pointer to **string** | A date that specifies the date the credential was last used in UTC. This is a required property. | [optional] [readonly] 
**Credential** | Pointer to **string** | A string that specifies the signed JWT for the gateway credential. This property is present only when the gateway credential is created. | [optional] [readonly] 

## Methods

### NewGatewayCredential

`func NewGatewayCredential() *GatewayCredential`

NewGatewayCredential instantiates a new GatewayCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCredentialWithDefaults

`func NewGatewayCredentialWithDefaults() *GatewayCredential`

NewGatewayCredentialWithDefaults instantiates a new GatewayCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayCredential) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayCredential) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayCredential) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayCredential) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GatewayCredential) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GatewayCredential) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GatewayCredential) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GatewayCredential) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastUsedAt

`func (o *GatewayCredential) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *GatewayCredential) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *GatewayCredential) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *GatewayCredential) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### GetCredential

`func (o *GatewayCredential) GetCredential() string`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GatewayCredential) GetCredentialOk() (*string, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GatewayCredential) SetCredential(v string)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GatewayCredential) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


