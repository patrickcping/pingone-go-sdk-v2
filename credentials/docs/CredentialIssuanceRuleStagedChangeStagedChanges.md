# CredentialIssuanceRuleStagedChangeStagedChanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to [**EnumCredentialIssuanceRuleAutomationMethod**](EnumCredentialIssuanceRuleAutomationMethod.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | A string that specifies the date and time the change was staged by the service. | [optional] 
**CredentialType** | Pointer to [**CredentialIssuanceRuleStagedChangeStagedChangesCredentialType**](CredentialIssuanceRuleStagedChangeStagedChangesCredentialType.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**IssuanceRule** | Pointer to [**CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule**](CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule.md) |  | [optional] 
**Scheduled** | Pointer to **bool** | A boolean that specifies whether or not the staged change is scheduled. | [optional] 
**User** | Pointer to [**CredentialIssuanceRuleStagedChangeStagedChangesUser**](CredentialIssuanceRuleStagedChangeStagedChangesUser.md) |  | [optional] 

## Methods

### NewCredentialIssuanceRuleStagedChangeStagedChanges

`func NewCredentialIssuanceRuleStagedChangeStagedChanges() *CredentialIssuanceRuleStagedChangeStagedChanges`

NewCredentialIssuanceRuleStagedChangeStagedChanges instantiates a new CredentialIssuanceRuleStagedChangeStagedChanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialIssuanceRuleStagedChangeStagedChangesWithDefaults

`func NewCredentialIssuanceRuleStagedChangeStagedChangesWithDefaults() *CredentialIssuanceRuleStagedChangeStagedChanges`

NewCredentialIssuanceRuleStagedChangeStagedChangesWithDefaults instantiates a new CredentialIssuanceRuleStagedChangeStagedChanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetAction() EnumCredentialIssuanceRuleAutomationMethod`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetActionOk() (*EnumCredentialIssuanceRuleAutomationMethod, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetAction(v EnumCredentialIssuanceRuleAutomationMethod)`

SetAction sets Action field to given value.

### HasAction

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCredentialType

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetCredentialType() CredentialIssuanceRuleStagedChangeStagedChangesCredentialType`

GetCredentialType returns the CredentialType field if non-nil, zero value otherwise.

### GetCredentialTypeOk

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetCredentialTypeOk() (*CredentialIssuanceRuleStagedChangeStagedChangesCredentialType, bool)`

GetCredentialTypeOk returns a tuple with the CredentialType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialType

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetCredentialType(v CredentialIssuanceRuleStagedChangeStagedChangesCredentialType)`

SetCredentialType sets CredentialType field to given value.

### HasCredentialType

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasCredentialType() bool`

HasCredentialType returns a boolean if a field has been set.

### GetEnvironment

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIssuanceRule

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetIssuanceRule() CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule`

GetIssuanceRule returns the IssuanceRule field if non-nil, zero value otherwise.

### GetIssuanceRuleOk

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetIssuanceRuleOk() (*CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule, bool)`

GetIssuanceRuleOk returns a tuple with the IssuanceRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuanceRule

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetIssuanceRule(v CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule)`

SetIssuanceRule sets IssuanceRule field to given value.

### HasIssuanceRule

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasIssuanceRule() bool`

HasIssuanceRule returns a boolean if a field has been set.

### GetScheduled

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetScheduled() bool`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetScheduledOk() (*bool, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetScheduled(v bool)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.

### GetUser

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetUser() CredentialIssuanceRuleStagedChangeStagedChangesUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetUserOk() (*CredentialIssuanceRuleStagedChangeStagedChangesUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetUser(v CredentialIssuanceRuleStagedChangeStagedChangesUser)`

SetUser sets User field to given value.

### HasUser

`func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


