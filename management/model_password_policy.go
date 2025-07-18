/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the PasswordPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PasswordPolicy{}

// PasswordPolicy struct for PasswordPolicy
type PasswordPolicy struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// Determines whether the password policy for a user will be ignored. If this property is omitted from a CREATE Password Policy request, its value is set to false.
	BypassPolicy *bool `json:"bypassPolicy,omitempty"`
	// The date and time the resource was created (format ISO-8061).
	CreatedAt *string `json:"createdAt,omitempty"`
	// The current password to be verified before the new password is set. Required for self-change when the user already has a password (the user whose password is being changed is the same as the actor in the access token).
	CurrentPassword *string `json:"currentPassword,omitempty"`
	// Indicates whether this password policy is enforced within the environment. When set to true, all other password policies are set to false.
	Default *bool `json:"default,omitempty"`
	// Specifies the brief description of the password policy.
	Description *string            `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// Set this to true to ensure the password is not one of the commonly used passwords.
	ExcludesCommonlyUsed bool `json:"excludesCommonlyUsed"`
	// Set this to true to ensure the password is not an exact match for the value of any attribute in the user’s profile, such as name, phone number, or address.
	ExcludesProfileData bool                   `json:"excludesProfileData"`
	History             *PasswordPolicyHistory `json:"history,omitempty"`
	// The password resource’s unique identifier.
	Id      *string                `json:"id,omitempty"`
	Length  *PasswordPolicyLength  `json:"length,omitempty"`
	Lockout *PasswordPolicyLockout `json:"lockout,omitempty"`
	// The maximum number of days the same password can be used before it must be changed. The value must be a positive, non-zero integer.  The value must be greater than the sum of minAgeDays (if set) + 21 (the expiration warning interval for passwords).
	MaxAgeDays *int32 `json:"maxAgeDays,omitempty"`
	// The maximum number of repeated characters allowed. This property is not enforced when not present.
	MaxRepeatedCharacters *int32 `json:"maxRepeatedCharacters,omitempty"`
	// The minimum number of days a password must be used before changing. The value must be a positive, non-zero integer. This property is not enforced when not present.
	MinAgeDays    *int32                       `json:"minAgeDays,omitempty"`
	MinCharacters *PasswordPolicyMinCharacters `json:"minCharacters,omitempty"`
	// The minimum complexity of the password based on the concept of password haystacks. The value is the number of days required to exhaust the entire search space during a brute force attack. This property is not enforced when not present.
	MinComplexity *int32 `json:"minComplexity,omitempty"`
	// The minimum number of unique characters required. This property is not enforced when not present.
	MinUniqueCharacters *int32 `json:"minUniqueCharacters,omitempty"`
	// The name of the password policy. This value must be unique within the environment.
	Name string `json:"name"`
	// The new password (must satisfy all requirements).
	NewPassword *string `json:"newPassword,omitempty"`
	// Set this to true to ensure that the proposed password is not too similar to the user's current password based on the Levenshtein distance algorithm. The value of this parameter is evaluated only for password change actions in which the user enters both the current and the new password. By design, PingOne does not know the user's current password.
	NotSimilarToCurrent bool `json:"notSimilarToCurrent"`
	// Returned in the response. The number of populations associated with the password policy.
	PopulationCount *int32 `json:"populationCount,omitempty"`
	// The date and time the resource was last updated (format ISO-8061).
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewPasswordPolicy instantiates a new PasswordPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPasswordPolicy(excludesCommonlyUsed bool, excludesProfileData bool, name string, notSimilarToCurrent bool) *PasswordPolicy {
	this := PasswordPolicy{}
	var bypassPolicy bool = false
	this.BypassPolicy = &bypassPolicy
	this.ExcludesCommonlyUsed = excludesCommonlyUsed
	this.ExcludesProfileData = excludesProfileData
	this.Name = name
	this.NotSimilarToCurrent = notSimilarToCurrent
	return &this
}

// NewPasswordPolicyWithDefaults instantiates a new PasswordPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordPolicyWithDefaults() *PasswordPolicy {
	this := PasswordPolicy{}
	var bypassPolicy bool = false
	this.BypassPolicy = &bypassPolicy
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PasswordPolicy) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PasswordPolicy) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *PasswordPolicy) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetBypassPolicy returns the BypassPolicy field value if set, zero value otherwise.
func (o *PasswordPolicy) GetBypassPolicy() bool {
	if o == nil || IsNil(o.BypassPolicy) {
		var ret bool
		return ret
	}
	return *o.BypassPolicy
}

// GetBypassPolicyOk returns a tuple with the BypassPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetBypassPolicyOk() (*bool, bool) {
	if o == nil || IsNil(o.BypassPolicy) {
		return nil, false
	}
	return o.BypassPolicy, true
}

// HasBypassPolicy returns a boolean if a field has been set.
func (o *PasswordPolicy) HasBypassPolicy() bool {
	if o != nil && !IsNil(o.BypassPolicy) {
		return true
	}

	return false
}

// SetBypassPolicy gets a reference to the given bool and assigns it to the BypassPolicy field.
func (o *PasswordPolicy) SetBypassPolicy(v bool) {
	o.BypassPolicy = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PasswordPolicy) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PasswordPolicy) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *PasswordPolicy) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetCurrentPassword returns the CurrentPassword field value if set, zero value otherwise.
func (o *PasswordPolicy) GetCurrentPassword() string {
	if o == nil || IsNil(o.CurrentPassword) {
		var ret string
		return ret
	}
	return *o.CurrentPassword
}

// GetCurrentPasswordOk returns a tuple with the CurrentPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetCurrentPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentPassword) {
		return nil, false
	}
	return o.CurrentPassword, true
}

// HasCurrentPassword returns a boolean if a field has been set.
func (o *PasswordPolicy) HasCurrentPassword() bool {
	if o != nil && !IsNil(o.CurrentPassword) {
		return true
	}

	return false
}

// SetCurrentPassword gets a reference to the given string and assigns it to the CurrentPassword field.
func (o *PasswordPolicy) SetCurrentPassword(v string) {
	o.CurrentPassword = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *PasswordPolicy) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *PasswordPolicy) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *PasswordPolicy) SetDefault(v bool) {
	o.Default = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PasswordPolicy) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PasswordPolicy) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PasswordPolicy) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *PasswordPolicy) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *PasswordPolicy) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *PasswordPolicy) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetExcludesCommonlyUsed returns the ExcludesCommonlyUsed field value
func (o *PasswordPolicy) GetExcludesCommonlyUsed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExcludesCommonlyUsed
}

// GetExcludesCommonlyUsedOk returns a tuple with the ExcludesCommonlyUsed field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetExcludesCommonlyUsedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExcludesCommonlyUsed, true
}

// SetExcludesCommonlyUsed sets field value
func (o *PasswordPolicy) SetExcludesCommonlyUsed(v bool) {
	o.ExcludesCommonlyUsed = v
}

// GetExcludesProfileData returns the ExcludesProfileData field value
func (o *PasswordPolicy) GetExcludesProfileData() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExcludesProfileData
}

// GetExcludesProfileDataOk returns a tuple with the ExcludesProfileData field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetExcludesProfileDataOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExcludesProfileData, true
}

// SetExcludesProfileData sets field value
func (o *PasswordPolicy) SetExcludesProfileData(v bool) {
	o.ExcludesProfileData = v
}

// GetHistory returns the History field value if set, zero value otherwise.
func (o *PasswordPolicy) GetHistory() PasswordPolicyHistory {
	if o == nil || IsNil(o.History) {
		var ret PasswordPolicyHistory
		return ret
	}
	return *o.History
}

// GetHistoryOk returns a tuple with the History field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetHistoryOk() (*PasswordPolicyHistory, bool) {
	if o == nil || IsNil(o.History) {
		return nil, false
	}
	return o.History, true
}

// HasHistory returns a boolean if a field has been set.
func (o *PasswordPolicy) HasHistory() bool {
	if o != nil && !IsNil(o.History) {
		return true
	}

	return false
}

// SetHistory gets a reference to the given PasswordPolicyHistory and assigns it to the History field.
func (o *PasswordPolicy) SetHistory(v PasswordPolicyHistory) {
	o.History = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PasswordPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PasswordPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PasswordPolicy) SetId(v string) {
	o.Id = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *PasswordPolicy) GetLength() PasswordPolicyLength {
	if o == nil || IsNil(o.Length) {
		var ret PasswordPolicyLength
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetLengthOk() (*PasswordPolicyLength, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *PasswordPolicy) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given PasswordPolicyLength and assigns it to the Length field.
func (o *PasswordPolicy) SetLength(v PasswordPolicyLength) {
	o.Length = &v
}

// GetLockout returns the Lockout field value if set, zero value otherwise.
func (o *PasswordPolicy) GetLockout() PasswordPolicyLockout {
	if o == nil || IsNil(o.Lockout) {
		var ret PasswordPolicyLockout
		return ret
	}
	return *o.Lockout
}

// GetLockoutOk returns a tuple with the Lockout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetLockoutOk() (*PasswordPolicyLockout, bool) {
	if o == nil || IsNil(o.Lockout) {
		return nil, false
	}
	return o.Lockout, true
}

// HasLockout returns a boolean if a field has been set.
func (o *PasswordPolicy) HasLockout() bool {
	if o != nil && !IsNil(o.Lockout) {
		return true
	}

	return false
}

// SetLockout gets a reference to the given PasswordPolicyLockout and assigns it to the Lockout field.
func (o *PasswordPolicy) SetLockout(v PasswordPolicyLockout) {
	o.Lockout = &v
}

// GetMaxAgeDays returns the MaxAgeDays field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMaxAgeDays() int32 {
	if o == nil || IsNil(o.MaxAgeDays) {
		var ret int32
		return ret
	}
	return *o.MaxAgeDays
}

// GetMaxAgeDaysOk returns a tuple with the MaxAgeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMaxAgeDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxAgeDays) {
		return nil, false
	}
	return o.MaxAgeDays, true
}

// HasMaxAgeDays returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMaxAgeDays() bool {
	if o != nil && !IsNil(o.MaxAgeDays) {
		return true
	}

	return false
}

// SetMaxAgeDays gets a reference to the given int32 and assigns it to the MaxAgeDays field.
func (o *PasswordPolicy) SetMaxAgeDays(v int32) {
	o.MaxAgeDays = &v
}

// GetMaxRepeatedCharacters returns the MaxRepeatedCharacters field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMaxRepeatedCharacters() int32 {
	if o == nil || IsNil(o.MaxRepeatedCharacters) {
		var ret int32
		return ret
	}
	return *o.MaxRepeatedCharacters
}

// GetMaxRepeatedCharactersOk returns a tuple with the MaxRepeatedCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMaxRepeatedCharactersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxRepeatedCharacters) {
		return nil, false
	}
	return o.MaxRepeatedCharacters, true
}

// HasMaxRepeatedCharacters returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMaxRepeatedCharacters() bool {
	if o != nil && !IsNil(o.MaxRepeatedCharacters) {
		return true
	}

	return false
}

// SetMaxRepeatedCharacters gets a reference to the given int32 and assigns it to the MaxRepeatedCharacters field.
func (o *PasswordPolicy) SetMaxRepeatedCharacters(v int32) {
	o.MaxRepeatedCharacters = &v
}

// GetMinAgeDays returns the MinAgeDays field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMinAgeDays() int32 {
	if o == nil || IsNil(o.MinAgeDays) {
		var ret int32
		return ret
	}
	return *o.MinAgeDays
}

// GetMinAgeDaysOk returns a tuple with the MinAgeDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMinAgeDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.MinAgeDays) {
		return nil, false
	}
	return o.MinAgeDays, true
}

// HasMinAgeDays returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMinAgeDays() bool {
	if o != nil && !IsNil(o.MinAgeDays) {
		return true
	}

	return false
}

// SetMinAgeDays gets a reference to the given int32 and assigns it to the MinAgeDays field.
func (o *PasswordPolicy) SetMinAgeDays(v int32) {
	o.MinAgeDays = &v
}

// GetMinCharacters returns the MinCharacters field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMinCharacters() PasswordPolicyMinCharacters {
	if o == nil || IsNil(o.MinCharacters) {
		var ret PasswordPolicyMinCharacters
		return ret
	}
	return *o.MinCharacters
}

// GetMinCharactersOk returns a tuple with the MinCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMinCharactersOk() (*PasswordPolicyMinCharacters, bool) {
	if o == nil || IsNil(o.MinCharacters) {
		return nil, false
	}
	return o.MinCharacters, true
}

// HasMinCharacters returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMinCharacters() bool {
	if o != nil && !IsNil(o.MinCharacters) {
		return true
	}

	return false
}

// SetMinCharacters gets a reference to the given PasswordPolicyMinCharacters and assigns it to the MinCharacters field.
func (o *PasswordPolicy) SetMinCharacters(v PasswordPolicyMinCharacters) {
	o.MinCharacters = &v
}

// GetMinComplexity returns the MinComplexity field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMinComplexity() int32 {
	if o == nil || IsNil(o.MinComplexity) {
		var ret int32
		return ret
	}
	return *o.MinComplexity
}

// GetMinComplexityOk returns a tuple with the MinComplexity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMinComplexityOk() (*int32, bool) {
	if o == nil || IsNil(o.MinComplexity) {
		return nil, false
	}
	return o.MinComplexity, true
}

// HasMinComplexity returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMinComplexity() bool {
	if o != nil && !IsNil(o.MinComplexity) {
		return true
	}

	return false
}

// SetMinComplexity gets a reference to the given int32 and assigns it to the MinComplexity field.
func (o *PasswordPolicy) SetMinComplexity(v int32) {
	o.MinComplexity = &v
}

// GetMinUniqueCharacters returns the MinUniqueCharacters field value if set, zero value otherwise.
func (o *PasswordPolicy) GetMinUniqueCharacters() int32 {
	if o == nil || IsNil(o.MinUniqueCharacters) {
		var ret int32
		return ret
	}
	return *o.MinUniqueCharacters
}

// GetMinUniqueCharactersOk returns a tuple with the MinUniqueCharacters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetMinUniqueCharactersOk() (*int32, bool) {
	if o == nil || IsNil(o.MinUniqueCharacters) {
		return nil, false
	}
	return o.MinUniqueCharacters, true
}

// HasMinUniqueCharacters returns a boolean if a field has been set.
func (o *PasswordPolicy) HasMinUniqueCharacters() bool {
	if o != nil && !IsNil(o.MinUniqueCharacters) {
		return true
	}

	return false
}

// SetMinUniqueCharacters gets a reference to the given int32 and assigns it to the MinUniqueCharacters field.
func (o *PasswordPolicy) SetMinUniqueCharacters(v int32) {
	o.MinUniqueCharacters = &v
}

// GetName returns the Name field value
func (o *PasswordPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PasswordPolicy) SetName(v string) {
	o.Name = v
}

// GetNewPassword returns the NewPassword field value if set, zero value otherwise.
func (o *PasswordPolicy) GetNewPassword() string {
	if o == nil || IsNil(o.NewPassword) {
		var ret string
		return ret
	}
	return *o.NewPassword
}

// GetNewPasswordOk returns a tuple with the NewPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetNewPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.NewPassword) {
		return nil, false
	}
	return o.NewPassword, true
}

// HasNewPassword returns a boolean if a field has been set.
func (o *PasswordPolicy) HasNewPassword() bool {
	if o != nil && !IsNil(o.NewPassword) {
		return true
	}

	return false
}

// SetNewPassword gets a reference to the given string and assigns it to the NewPassword field.
func (o *PasswordPolicy) SetNewPassword(v string) {
	o.NewPassword = &v
}

// GetNotSimilarToCurrent returns the NotSimilarToCurrent field value
func (o *PasswordPolicy) GetNotSimilarToCurrent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.NotSimilarToCurrent
}

// GetNotSimilarToCurrentOk returns a tuple with the NotSimilarToCurrent field value
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetNotSimilarToCurrentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotSimilarToCurrent, true
}

// SetNotSimilarToCurrent sets field value
func (o *PasswordPolicy) SetNotSimilarToCurrent(v bool) {
	o.NotSimilarToCurrent = v
}

// GetPopulationCount returns the PopulationCount field value if set, zero value otherwise.
func (o *PasswordPolicy) GetPopulationCount() int32 {
	if o == nil || IsNil(o.PopulationCount) {
		var ret int32
		return ret
	}
	return *o.PopulationCount
}

// GetPopulationCountOk returns a tuple with the PopulationCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetPopulationCountOk() (*int32, bool) {
	if o == nil || IsNil(o.PopulationCount) {
		return nil, false
	}
	return o.PopulationCount, true
}

// HasPopulationCount returns a boolean if a field has been set.
func (o *PasswordPolicy) HasPopulationCount() bool {
	if o != nil && !IsNil(o.PopulationCount) {
		return true
	}

	return false
}

// SetPopulationCount gets a reference to the given int32 and assigns it to the PopulationCount field.
func (o *PasswordPolicy) SetPopulationCount(v int32) {
	o.PopulationCount = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PasswordPolicy) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PasswordPolicy) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PasswordPolicy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *PasswordPolicy) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o PasswordPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.BypassPolicy) {
		toSerialize["bypassPolicy"] = o.BypassPolicy
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CurrentPassword) {
		toSerialize["currentPassword"] = o.CurrentPassword
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["excludesCommonlyUsed"] = o.ExcludesCommonlyUsed
	toSerialize["excludesProfileData"] = o.ExcludesProfileData
	if !IsNil(o.History) {
		toSerialize["history"] = o.History
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.Lockout) {
		toSerialize["lockout"] = o.Lockout
	}
	if !IsNil(o.MaxAgeDays) {
		toSerialize["maxAgeDays"] = o.MaxAgeDays
	}
	if !IsNil(o.MaxRepeatedCharacters) {
		toSerialize["maxRepeatedCharacters"] = o.MaxRepeatedCharacters
	}
	if !IsNil(o.MinAgeDays) {
		toSerialize["minAgeDays"] = o.MinAgeDays
	}
	if !IsNil(o.MinCharacters) {
		toSerialize["minCharacters"] = o.MinCharacters
	}
	if !IsNil(o.MinComplexity) {
		toSerialize["minComplexity"] = o.MinComplexity
	}
	if !IsNil(o.MinUniqueCharacters) {
		toSerialize["minUniqueCharacters"] = o.MinUniqueCharacters
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NewPassword) {
		toSerialize["newPassword"] = o.NewPassword
	}
	toSerialize["notSimilarToCurrent"] = o.NotSimilarToCurrent
	if !IsNil(o.PopulationCount) {
		toSerialize["populationCount"] = o.PopulationCount
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullablePasswordPolicy struct {
	value *PasswordPolicy
	isSet bool
}

func (v NullablePasswordPolicy) Get() *PasswordPolicy {
	return v.value
}

func (v *NullablePasswordPolicy) Set(val *PasswordPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePasswordPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordPolicy(val *PasswordPolicy) *NullablePasswordPolicy {
	return &NullablePasswordPolicy{value: val, isSet: true}
}

func (v NullablePasswordPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
