# CredentialTypeOnDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RevokeIssuedCredentials** | Pointer to **bool** | A boolean that specifies whether a user&#39;s issued verifiable credentials are automatically revoked when a credential type, user, or environment is deleted. | [optional] 

## Methods

### NewCredentialTypeOnDelete

`func NewCredentialTypeOnDelete() *CredentialTypeOnDelete`

NewCredentialTypeOnDelete instantiates a new CredentialTypeOnDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeOnDeleteWithDefaults

`func NewCredentialTypeOnDeleteWithDefaults() *CredentialTypeOnDelete`

NewCredentialTypeOnDeleteWithDefaults instantiates a new CredentialTypeOnDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevokeIssuedCredentials

`func (o *CredentialTypeOnDelete) GetRevokeIssuedCredentials() bool`

GetRevokeIssuedCredentials returns the RevokeIssuedCredentials field if non-nil, zero value otherwise.

### GetRevokeIssuedCredentialsOk

`func (o *CredentialTypeOnDelete) GetRevokeIssuedCredentialsOk() (*bool, bool)`

GetRevokeIssuedCredentialsOk returns a tuple with the RevokeIssuedCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokeIssuedCredentials

`func (o *CredentialTypeOnDelete) SetRevokeIssuedCredentials(v bool)`

SetRevokeIssuedCredentials sets RevokeIssuedCredentials field to given value.

### HasRevokeIssuedCredentials

`func (o *CredentialTypeOnDelete) HasRevokeIssuedCredentials() bool`

HasRevokeIssuedCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


