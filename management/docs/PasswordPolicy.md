# PasswordPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**BypassPolicy** | Pointer to **bool** | Determines whether the password policy for a user will be ignored. If this property is omitted from a CREATE Password Policy request, its value is set to false. | [optional] [default to false]
**CreatedAt** | Pointer to **string** | The date and time the resource was created (format ISO-8061). | [optional] [readonly] 
**CurrentPassword** | Pointer to **string** | The current password to be verified before the new password is set. Required for self-change when the user already has a password (the user whose password is being changed is the same as the actor in the access token). | [optional] 
**Default** | Pointer to **bool** | Indicates whether this password policy is enforced within the environment. When set to true, all other password policies are set to false. | [optional] 
**Description** | Pointer to **string** | Specifies the brief description of the password policy. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**ExcludesCommonlyUsed** | **bool** | Set this to true to ensure the password is not one of the commonly used passwords. | 
**ExcludesProfileData** | **bool** | Set this to true to ensure the password is not an exact match for the value of any attribute in the user’s profile, such as name, phone number, or address. | 
**History** | Pointer to [**PasswordPolicyHistory**](PasswordPolicyHistory.md) |  | [optional] 
**Id** | Pointer to **string** | The password resource’s unique identifier. | [optional] [readonly] 
**Length** | Pointer to [**PasswordPolicyLength**](PasswordPolicyLength.md) |  | [optional] 
**Lockout** | Pointer to [**PasswordPolicyLockout**](PasswordPolicyLockout.md) |  | [optional] 
**MaxAgeDays** | Pointer to **int32** | The maximum number of days the same password can be used before it must be changed. The value must be a positive, non-zero integer.  The value must be greater than the sum of minAgeDays (if set) + 21 (the expiration warning interval for passwords). | [optional] 
**MaxRepeatedCharacters** | Pointer to **int32** | The maximum number of repeated characters allowed. This property is not enforced when not present. | [optional] 
**MinAgeDays** | Pointer to **int32** | The minimum number of days a password must be used before changing. The value must be a positive, non-zero integer. This property is not enforced when not present. | [optional] 
**MinCharacters** | Pointer to [**PasswordPolicyMinCharacters**](PasswordPolicyMinCharacters.md) |  | [optional] 
**MinComplexity** | Pointer to **int32** | The minimum complexity of the password based on the concept of password haystacks. The value is the number of days required to exhaust the entire search space during a brute force attack. This property is not enforced when not present. | [optional] 
**MinUniqueCharacters** | Pointer to **int32** | The minimum number of unique characters required. This property is not enforced when not present. | [optional] 
**Name** | **string** | The name of the password policy. This value must be unique within the environment. | 
**NewPassword** | Pointer to **string** | The new password (must satisfy all requirements). | [optional] 
**NotSimilarToCurrent** | **bool** | Set this to true to ensure that the proposed password is not too similar to the user&#39;s current password based on the Levenshtein distance algorithm. The value of this parameter is evaluated only for password change actions in which the user enters both the current and the new password. By design, PingOne does not know the user&#39;s current password. | 
**PopulationCount** | Pointer to **int32** | Returned in the response. The number of populations associated with the password policy. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The date and time the resource was last updated (format ISO-8061). | [optional] [readonly] 

## Methods

### NewPasswordPolicy

`func NewPasswordPolicy(excludesCommonlyUsed bool, excludesProfileData bool, name string, notSimilarToCurrent bool, ) *PasswordPolicy`

NewPasswordPolicy instantiates a new PasswordPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordPolicyWithDefaults

`func NewPasswordPolicyWithDefaults() *PasswordPolicy`

NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *PasswordPolicy) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *PasswordPolicy) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *PasswordPolicy) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *PasswordPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetBypassPolicy

`func (o *PasswordPolicy) GetBypassPolicy() bool`

GetBypassPolicy returns the BypassPolicy field if non-nil, zero value otherwise.

### GetBypassPolicyOk

`func (o *PasswordPolicy) GetBypassPolicyOk() (*bool, bool)`

GetBypassPolicyOk returns a tuple with the BypassPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBypassPolicy

`func (o *PasswordPolicy) SetBypassPolicy(v bool)`

SetBypassPolicy sets BypassPolicy field to given value.

### HasBypassPolicy

`func (o *PasswordPolicy) HasBypassPolicy() bool`

HasBypassPolicy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PasswordPolicy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PasswordPolicy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PasswordPolicy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PasswordPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCurrentPassword

`func (o *PasswordPolicy) GetCurrentPassword() string`

GetCurrentPassword returns the CurrentPassword field if non-nil, zero value otherwise.

### GetCurrentPasswordOk

`func (o *PasswordPolicy) GetCurrentPasswordOk() (*string, bool)`

GetCurrentPasswordOk returns a tuple with the CurrentPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPassword

`func (o *PasswordPolicy) SetCurrentPassword(v string)`

SetCurrentPassword sets CurrentPassword field to given value.

### HasCurrentPassword

`func (o *PasswordPolicy) HasCurrentPassword() bool`

HasCurrentPassword returns a boolean if a field has been set.

### GetDefault

`func (o *PasswordPolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *PasswordPolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *PasswordPolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *PasswordPolicy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *PasswordPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PasswordPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PasswordPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PasswordPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *PasswordPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *PasswordPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *PasswordPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *PasswordPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetExcludesCommonlyUsed

`func (o *PasswordPolicy) GetExcludesCommonlyUsed() bool`

GetExcludesCommonlyUsed returns the ExcludesCommonlyUsed field if non-nil, zero value otherwise.

### GetExcludesCommonlyUsedOk

`func (o *PasswordPolicy) GetExcludesCommonlyUsedOk() (*bool, bool)`

GetExcludesCommonlyUsedOk returns a tuple with the ExcludesCommonlyUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludesCommonlyUsed

`func (o *PasswordPolicy) SetExcludesCommonlyUsed(v bool)`

SetExcludesCommonlyUsed sets ExcludesCommonlyUsed field to given value.


### GetExcludesProfileData

`func (o *PasswordPolicy) GetExcludesProfileData() bool`

GetExcludesProfileData returns the ExcludesProfileData field if non-nil, zero value otherwise.

### GetExcludesProfileDataOk

`func (o *PasswordPolicy) GetExcludesProfileDataOk() (*bool, bool)`

GetExcludesProfileDataOk returns a tuple with the ExcludesProfileData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludesProfileData

`func (o *PasswordPolicy) SetExcludesProfileData(v bool)`

SetExcludesProfileData sets ExcludesProfileData field to given value.


### GetHistory

`func (o *PasswordPolicy) GetHistory() PasswordPolicyHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *PasswordPolicy) GetHistoryOk() (*PasswordPolicyHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *PasswordPolicy) SetHistory(v PasswordPolicyHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *PasswordPolicy) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetId

`func (o *PasswordPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PasswordPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PasswordPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PasswordPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLength

`func (o *PasswordPolicy) GetLength() PasswordPolicyLength`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *PasswordPolicy) GetLengthOk() (*PasswordPolicyLength, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *PasswordPolicy) SetLength(v PasswordPolicyLength)`

SetLength sets Length field to given value.

### HasLength

`func (o *PasswordPolicy) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetLockout

`func (o *PasswordPolicy) GetLockout() PasswordPolicyLockout`

GetLockout returns the Lockout field if non-nil, zero value otherwise.

### GetLockoutOk

`func (o *PasswordPolicy) GetLockoutOk() (*PasswordPolicyLockout, bool)`

GetLockoutOk returns a tuple with the Lockout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockout

`func (o *PasswordPolicy) SetLockout(v PasswordPolicyLockout)`

SetLockout sets Lockout field to given value.

### HasLockout

`func (o *PasswordPolicy) HasLockout() bool`

HasLockout returns a boolean if a field has been set.

### GetMaxAgeDays

`func (o *PasswordPolicy) GetMaxAgeDays() int32`

GetMaxAgeDays returns the MaxAgeDays field if non-nil, zero value otherwise.

### GetMaxAgeDaysOk

`func (o *PasswordPolicy) GetMaxAgeDaysOk() (*int32, bool)`

GetMaxAgeDaysOk returns a tuple with the MaxAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAgeDays

`func (o *PasswordPolicy) SetMaxAgeDays(v int32)`

SetMaxAgeDays sets MaxAgeDays field to given value.

### HasMaxAgeDays

`func (o *PasswordPolicy) HasMaxAgeDays() bool`

HasMaxAgeDays returns a boolean if a field has been set.

### GetMaxRepeatedCharacters

`func (o *PasswordPolicy) GetMaxRepeatedCharacters() int32`

GetMaxRepeatedCharacters returns the MaxRepeatedCharacters field if non-nil, zero value otherwise.

### GetMaxRepeatedCharactersOk

`func (o *PasswordPolicy) GetMaxRepeatedCharactersOk() (*int32, bool)`

GetMaxRepeatedCharactersOk returns a tuple with the MaxRepeatedCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRepeatedCharacters

`func (o *PasswordPolicy) SetMaxRepeatedCharacters(v int32)`

SetMaxRepeatedCharacters sets MaxRepeatedCharacters field to given value.

### HasMaxRepeatedCharacters

`func (o *PasswordPolicy) HasMaxRepeatedCharacters() bool`

HasMaxRepeatedCharacters returns a boolean if a field has been set.

### GetMinAgeDays

`func (o *PasswordPolicy) GetMinAgeDays() int32`

GetMinAgeDays returns the MinAgeDays field if non-nil, zero value otherwise.

### GetMinAgeDaysOk

`func (o *PasswordPolicy) GetMinAgeDaysOk() (*int32, bool)`

GetMinAgeDaysOk returns a tuple with the MinAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAgeDays

`func (o *PasswordPolicy) SetMinAgeDays(v int32)`

SetMinAgeDays sets MinAgeDays field to given value.

### HasMinAgeDays

`func (o *PasswordPolicy) HasMinAgeDays() bool`

HasMinAgeDays returns a boolean if a field has been set.

### GetMinCharacters

`func (o *PasswordPolicy) GetMinCharacters() PasswordPolicyMinCharacters`

GetMinCharacters returns the MinCharacters field if non-nil, zero value otherwise.

### GetMinCharactersOk

`func (o *PasswordPolicy) GetMinCharactersOk() (*PasswordPolicyMinCharacters, bool)`

GetMinCharactersOk returns a tuple with the MinCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCharacters

`func (o *PasswordPolicy) SetMinCharacters(v PasswordPolicyMinCharacters)`

SetMinCharacters sets MinCharacters field to given value.

### HasMinCharacters

`func (o *PasswordPolicy) HasMinCharacters() bool`

HasMinCharacters returns a boolean if a field has been set.

### GetMinComplexity

`func (o *PasswordPolicy) GetMinComplexity() int32`

GetMinComplexity returns the MinComplexity field if non-nil, zero value otherwise.

### GetMinComplexityOk

`func (o *PasswordPolicy) GetMinComplexityOk() (*int32, bool)`

GetMinComplexityOk returns a tuple with the MinComplexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinComplexity

`func (o *PasswordPolicy) SetMinComplexity(v int32)`

SetMinComplexity sets MinComplexity field to given value.

### HasMinComplexity

`func (o *PasswordPolicy) HasMinComplexity() bool`

HasMinComplexity returns a boolean if a field has been set.

### GetMinUniqueCharacters

`func (o *PasswordPolicy) GetMinUniqueCharacters() int32`

GetMinUniqueCharacters returns the MinUniqueCharacters field if non-nil, zero value otherwise.

### GetMinUniqueCharactersOk

`func (o *PasswordPolicy) GetMinUniqueCharactersOk() (*int32, bool)`

GetMinUniqueCharactersOk returns a tuple with the MinUniqueCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinUniqueCharacters

`func (o *PasswordPolicy) SetMinUniqueCharacters(v int32)`

SetMinUniqueCharacters sets MinUniqueCharacters field to given value.

### HasMinUniqueCharacters

`func (o *PasswordPolicy) HasMinUniqueCharacters() bool`

HasMinUniqueCharacters returns a boolean if a field has been set.

### GetName

`func (o *PasswordPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PasswordPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PasswordPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetNewPassword

`func (o *PasswordPolicy) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *PasswordPolicy) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *PasswordPolicy) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *PasswordPolicy) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.

### GetNotSimilarToCurrent

`func (o *PasswordPolicy) GetNotSimilarToCurrent() bool`

GetNotSimilarToCurrent returns the NotSimilarToCurrent field if non-nil, zero value otherwise.

### GetNotSimilarToCurrentOk

`func (o *PasswordPolicy) GetNotSimilarToCurrentOk() (*bool, bool)`

GetNotSimilarToCurrentOk returns a tuple with the NotSimilarToCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotSimilarToCurrent

`func (o *PasswordPolicy) SetNotSimilarToCurrent(v bool)`

SetNotSimilarToCurrent sets NotSimilarToCurrent field to given value.


### GetPopulationCount

`func (o *PasswordPolicy) GetPopulationCount() int32`

GetPopulationCount returns the PopulationCount field if non-nil, zero value otherwise.

### GetPopulationCountOk

`func (o *PasswordPolicy) GetPopulationCountOk() (*int32, bool)`

GetPopulationCountOk returns a tuple with the PopulationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationCount

`func (o *PasswordPolicy) SetPopulationCount(v int32)`

SetPopulationCount sets PopulationCount field to given value.

### HasPopulationCount

`func (o *PasswordPolicy) HasPopulationCount() bool`

HasPopulationCount returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PasswordPolicy) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PasswordPolicy) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PasswordPolicy) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PasswordPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


