/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
	"time"
)

// checks if the CredentialIssuanceRuleStagedChangeStagedChanges type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleStagedChangeStagedChanges{}

// CredentialIssuanceRuleStagedChangeStagedChanges struct for CredentialIssuanceRuleStagedChangeStagedChanges
type CredentialIssuanceRuleStagedChangeStagedChanges struct {
	Action *EnumCredentialIssuanceRuleAutomationMethod `json:"action,omitempty"`
	// A string that specifies the date and time the change was staged by the service.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	CredentialType *CredentialIssuanceRuleStagedChangeStagedChangesCredentialType `json:"credentialType,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	IssuanceRule *CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule `json:"issuanceRule,omitempty"`
	// A boolean that specifies whether or not the staged change is scheduled.
	Scheduled *bool `json:"scheduled,omitempty"`
	User *CredentialIssuanceRuleStagedChangeStagedChangesUser `json:"user,omitempty"`
}

// NewCredentialIssuanceRuleStagedChangeStagedChanges instantiates a new CredentialIssuanceRuleStagedChangeStagedChanges object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleStagedChangeStagedChanges() *CredentialIssuanceRuleStagedChangeStagedChanges {
	this := CredentialIssuanceRuleStagedChangeStagedChanges{}
	return &this
}

// NewCredentialIssuanceRuleStagedChangeStagedChangesWithDefaults instantiates a new CredentialIssuanceRuleStagedChangeStagedChanges object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleStagedChangeStagedChangesWithDefaults() *CredentialIssuanceRuleStagedChangeStagedChanges {
	this := CredentialIssuanceRuleStagedChangeStagedChanges{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetAction() EnumCredentialIssuanceRuleAutomationMethod {
	if o == nil || IsNil(o.Action) {
		var ret EnumCredentialIssuanceRuleAutomationMethod
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetActionOk() (*EnumCredentialIssuanceRuleAutomationMethod, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given EnumCredentialIssuanceRuleAutomationMethod and assigns it to the Action field.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetAction(v EnumCredentialIssuanceRuleAutomationMethod) {
	o.Action = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetCredentialType returns the CredentialType field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetCredentialType() CredentialIssuanceRuleStagedChangeStagedChangesCredentialType {
	if o == nil || IsNil(o.CredentialType) {
		var ret CredentialIssuanceRuleStagedChangeStagedChangesCredentialType
		return ret
	}
	return *o.CredentialType
}

// GetCredentialTypeOk returns a tuple with the CredentialType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetCredentialTypeOk() (*CredentialIssuanceRuleStagedChangeStagedChangesCredentialType, bool) {
	if o == nil || IsNil(o.CredentialType) {
		return nil, false
	}
	return o.CredentialType, true
}

// HasCredentialType returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasCredentialType() bool {
	if o != nil && !IsNil(o.CredentialType) {
		return true
	}

	return false
}

// SetCredentialType gets a reference to the given CredentialIssuanceRuleStagedChangeStagedChangesCredentialType and assigns it to the CredentialType field.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetCredentialType(v CredentialIssuanceRuleStagedChangeStagedChangesCredentialType) {
	o.CredentialType = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIssuanceRule returns the IssuanceRule field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetIssuanceRule() CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule {
	if o == nil || IsNil(o.IssuanceRule) {
		var ret CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule
		return ret
	}
	return *o.IssuanceRule
}

// GetIssuanceRuleOk returns a tuple with the IssuanceRule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetIssuanceRuleOk() (*CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule, bool) {
	if o == nil || IsNil(o.IssuanceRule) {
		return nil, false
	}
	return o.IssuanceRule, true
}

// HasIssuanceRule returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasIssuanceRule() bool {
	if o != nil && !IsNil(o.IssuanceRule) {
		return true
	}

	return false
}

// SetIssuanceRule gets a reference to the given CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule and assigns it to the IssuanceRule field.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetIssuanceRule(v CredentialIssuanceRuleStagedChangeStagedChangesIssuanceRule) {
	o.IssuanceRule = &v
}

// GetScheduled returns the Scheduled field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetScheduled() bool {
	if o == nil || IsNil(o.Scheduled) {
		var ret bool
		return ret
	}
	return *o.Scheduled
}

// GetScheduledOk returns a tuple with the Scheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetScheduledOk() (*bool, bool) {
	if o == nil || IsNil(o.Scheduled) {
		return nil, false
	}
	return o.Scheduled, true
}

// HasScheduled returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasScheduled() bool {
	if o != nil && !IsNil(o.Scheduled) {
		return true
	}

	return false
}

// SetScheduled gets a reference to the given bool and assigns it to the Scheduled field.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetScheduled(v bool) {
	o.Scheduled = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetUser() CredentialIssuanceRuleStagedChangeStagedChangesUser {
	if o == nil || IsNil(o.User) {
		var ret CredentialIssuanceRuleStagedChangeStagedChangesUser
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) GetUserOk() (*CredentialIssuanceRuleStagedChangeStagedChangesUser, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given CredentialIssuanceRuleStagedChangeStagedChangesUser and assigns it to the User field.
func (o *CredentialIssuanceRuleStagedChangeStagedChanges) SetUser(v CredentialIssuanceRuleStagedChangeStagedChangesUser) {
	o.User = &v
}

func (o CredentialIssuanceRuleStagedChangeStagedChanges) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleStagedChangeStagedChanges) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.CredentialType) {
		toSerialize["credentialType"] = o.CredentialType
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.IssuanceRule) {
		toSerialize["issuanceRule"] = o.IssuanceRule
	}
	if !IsNil(o.Scheduled) {
		toSerialize["scheduled"] = o.Scheduled
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleStagedChangeStagedChanges struct {
	value *CredentialIssuanceRuleStagedChangeStagedChanges
	isSet bool
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChanges) Get() *CredentialIssuanceRuleStagedChangeStagedChanges {
	return v.value
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChanges) Set(val *CredentialIssuanceRuleStagedChangeStagedChanges) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChanges) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChanges) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleStagedChangeStagedChanges(val *CredentialIssuanceRuleStagedChangeStagedChanges) *NullableCredentialIssuanceRuleStagedChangeStagedChanges {
	return &NullableCredentialIssuanceRuleStagedChangeStagedChanges{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleStagedChangeStagedChanges) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleStagedChangeStagedChanges) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


