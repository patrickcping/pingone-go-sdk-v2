# CredentialIssuanceRuleUsageInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**CredentialIssuanceRuleUsageInnerUser**](CredentialIssuanceRuleUsageInnerUser.md) |  | [optional] 
**Credential** | Pointer to [**CredentialIssuanceRuleUsageInnerCredential**](CredentialIssuanceRuleUsageInnerCredential.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | A string representing the date and time the credential was issued by the service. | [optional] [readonly] 

## Methods

### NewCredentialIssuanceRuleUsageInner

`func NewCredentialIssuanceRuleUsageInner() *CredentialIssuanceRuleUsageInner`

NewCredentialIssuanceRuleUsageInner instantiates a new CredentialIssuanceRuleUsageInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuanceRuleUsageInnerWithDefaults

`func NewCredentialIssuanceRuleUsageInnerWithDefaults() *CredentialIssuanceRuleUsageInner`

NewCredentialIssuanceRuleUsageInnerWithDefaults instantiates a new CredentialIssuanceRuleUsageInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *CredentialIssuanceRuleUsageInner) GetUser() CredentialIssuanceRuleUsageInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CredentialIssuanceRuleUsageInner) GetUserOk() (*CredentialIssuanceRuleUsageInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CredentialIssuanceRuleUsageInner) SetUser(v CredentialIssuanceRuleUsageInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *CredentialIssuanceRuleUsageInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetCredential

`func (o *CredentialIssuanceRuleUsageInner) GetCredential() CredentialIssuanceRuleUsageInnerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *CredentialIssuanceRuleUsageInner) GetCredentialOk() (*CredentialIssuanceRuleUsageInnerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *CredentialIssuanceRuleUsageInner) SetCredential(v CredentialIssuanceRuleUsageInnerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *CredentialIssuanceRuleUsageInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CredentialIssuanceRuleUsageInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialIssuanceRuleUsageInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialIssuanceRuleUsageInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialIssuanceRuleUsageInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


