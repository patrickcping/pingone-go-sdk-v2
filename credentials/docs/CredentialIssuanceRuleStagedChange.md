# CredentialIssuanceRuleStagedChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issue** | Pointer to **[]string** | An array that specifies one or more identifiers (UUIDs) of users whose credentials are in an issue action state and should be issued. | [optional] 
**Revoke** | Pointer to **[]string** | An array that specifies one or more identifiers (UUIDs) of users whose credentials are in an revoke action state and should be issued. | [optional] 
**StagedChanges** | Pointer to [**CredentialIssuanceRuleStagedChangeStagedChanges**](CredentialIssuanceRuleStagedChangeStagedChanges.md) |  | [optional] 
**Update** | Pointer to **[]string** | An array that specifies one or more identifiers (UUIDs) of users whose credentials are in an update action state and should be issued. | [optional] 

## Methods

### NewCredentialIssuanceRuleStagedChange

`func NewCredentialIssuanceRuleStagedChange() *CredentialIssuanceRuleStagedChange`

NewCredentialIssuanceRuleStagedChange instantiates a new CredentialIssuanceRuleStagedChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuanceRuleStagedChangeWithDefaults

`func NewCredentialIssuanceRuleStagedChangeWithDefaults() *CredentialIssuanceRuleStagedChange`

NewCredentialIssuanceRuleStagedChangeWithDefaults instantiates a new CredentialIssuanceRuleStagedChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssue

`func (o *CredentialIssuanceRuleStagedChange) GetIssue() []string`

GetIssue returns the Issue field if non-nil, zero value otherwise.

### GetIssueOk

`func (o *CredentialIssuanceRuleStagedChange) GetIssueOk() (*[]string, bool)`

GetIssueOk returns a tuple with the Issue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssue

`func (o *CredentialIssuanceRuleStagedChange) SetIssue(v []string)`

SetIssue sets Issue field to given value.

### HasIssue

`func (o *CredentialIssuanceRuleStagedChange) HasIssue() bool`

HasIssue returns a boolean if a field has been set.

### GetRevoke

`func (o *CredentialIssuanceRuleStagedChange) GetRevoke() []string`

GetRevoke returns the Revoke field if non-nil, zero value otherwise.

### GetRevokeOk

`func (o *CredentialIssuanceRuleStagedChange) GetRevokeOk() (*[]string, bool)`

GetRevokeOk returns a tuple with the Revoke field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoke

`func (o *CredentialIssuanceRuleStagedChange) SetRevoke(v []string)`

SetRevoke sets Revoke field to given value.

### HasRevoke

`func (o *CredentialIssuanceRuleStagedChange) HasRevoke() bool`

HasRevoke returns a boolean if a field has been set.

### GetStagedChanges

`func (o *CredentialIssuanceRuleStagedChange) GetStagedChanges() CredentialIssuanceRuleStagedChangeStagedChanges`

GetStagedChanges returns the StagedChanges field if non-nil, zero value otherwise.

### GetStagedChangesOk

`func (o *CredentialIssuanceRuleStagedChange) GetStagedChangesOk() (*CredentialIssuanceRuleStagedChangeStagedChanges, bool)`

GetStagedChangesOk returns a tuple with the StagedChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStagedChanges

`func (o *CredentialIssuanceRuleStagedChange) SetStagedChanges(v CredentialIssuanceRuleStagedChangeStagedChanges)`

SetStagedChanges sets StagedChanges field to given value.

### HasStagedChanges

`func (o *CredentialIssuanceRuleStagedChange) HasStagedChanges() bool`

HasStagedChanges returns a boolean if a field has been set.

### GetUpdate

`func (o *CredentialIssuanceRuleStagedChange) GetUpdate() []string`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *CredentialIssuanceRuleStagedChange) GetUpdateOk() (*[]string, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *CredentialIssuanceRuleStagedChange) SetUpdate(v []string)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *CredentialIssuanceRuleStagedChange) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


