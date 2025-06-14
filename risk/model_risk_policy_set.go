/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"time"
)

// checks if the RiskPolicySet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPolicySet{}

// RiskPolicySet struct for RiskPolicySet
type RiskPolicySet struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// The time the resource was created (format ISO-8061).
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A boolean that specifies whether this risk policy set is the environment's default risk policy set, which is used whenever an explicit policySet ID is not specified in the risk policy evaluation request. If this property is not specified, the value defaults to false, and this risk policy set is not regarded as the default risk policy set for the environment. When this property is set to true (in PUT or POST requests), the default property of all other risk policy sets in the environment is set to false.
	Default       *bool                       `json:"default,omitempty"`
	DefaultResult *RiskPolicySetDefaultResult `json:"defaultResult,omitempty"`
	// A string that specifies a description for this policy set. This is an optional property. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, punctuation character, or space. Maximum size is 1024 characters.
	Description *string            `json:"description,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies a name for this policy set. Valid characters consist of any Unicode letter, mark (for example, accent, umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. Maximum size is 256 characters.
	Name string `json:"name"`
	// An array of policies related to this policy set.
	RiskPolicies []RiskPolicy `json:"riskPolicies,omitempty"`
	// The time the resource was last updated (format ISO-8061).
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The IDs for the predictors to evaluate in this policy set. In POST and PUT requests, if this property is null, all of the licensed predictors are used.
	EvaluatedPredictors []RiskPolicySetEvaluatedPredictorsInner `json:"evaluatedPredictors,omitempty"`
	// An array of triggers related to this policy set.
	Triggers []RiskPolicySetTriggersInner `json:"triggers,omitempty"`
}

// NewRiskPolicySet instantiates a new RiskPolicySet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPolicySet(name string) *RiskPolicySet {
	this := RiskPolicySet{}
	this.Name = name
	return &this
}

// NewRiskPolicySetWithDefaults instantiates a new RiskPolicySet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPolicySetWithDefaults() *RiskPolicySet {
	this := RiskPolicySet{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RiskPolicySet) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RiskPolicySet) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *RiskPolicySet) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskPolicySet) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskPolicySet) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RiskPolicySet) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RiskPolicySet) GetDefault() bool {
	if o == nil || IsNil(o.Default) {
		var ret bool
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RiskPolicySet) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given bool and assigns it to the Default field.
func (o *RiskPolicySet) SetDefault(v bool) {
	o.Default = &v
}

// GetDefaultResult returns the DefaultResult field value if set, zero value otherwise.
func (o *RiskPolicySet) GetDefaultResult() RiskPolicySetDefaultResult {
	if o == nil || IsNil(o.DefaultResult) {
		var ret RiskPolicySetDefaultResult
		return ret
	}
	return *o.DefaultResult
}

// GetDefaultResultOk returns a tuple with the DefaultResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetDefaultResultOk() (*RiskPolicySetDefaultResult, bool) {
	if o == nil || IsNil(o.DefaultResult) {
		return nil, false
	}
	return o.DefaultResult, true
}

// HasDefaultResult returns a boolean if a field has been set.
func (o *RiskPolicySet) HasDefaultResult() bool {
	if o != nil && !IsNil(o.DefaultResult) {
		return true
	}

	return false
}

// SetDefaultResult gets a reference to the given RiskPolicySetDefaultResult and assigns it to the DefaultResult field.
func (o *RiskPolicySet) SetDefaultResult(v RiskPolicySetDefaultResult) {
	o.DefaultResult = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskPolicySet) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskPolicySet) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskPolicySet) SetDescription(v string) {
	o.Description = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *RiskPolicySet) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *RiskPolicySet) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *RiskPolicySet) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskPolicySet) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskPolicySet) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskPolicySet) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RiskPolicySet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskPolicySet) SetName(v string) {
	o.Name = v
}

// GetRiskPolicies returns the RiskPolicies field value if set, zero value otherwise.
func (o *RiskPolicySet) GetRiskPolicies() []RiskPolicy {
	if o == nil || IsNil(o.RiskPolicies) {
		var ret []RiskPolicy
		return ret
	}
	return o.RiskPolicies
}

// GetRiskPoliciesOk returns a tuple with the RiskPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetRiskPoliciesOk() ([]RiskPolicy, bool) {
	if o == nil || IsNil(o.RiskPolicies) {
		return nil, false
	}
	return o.RiskPolicies, true
}

// HasRiskPolicies returns a boolean if a field has been set.
func (o *RiskPolicySet) HasRiskPolicies() bool {
	if o != nil && !IsNil(o.RiskPolicies) {
		return true
	}

	return false
}

// SetRiskPolicies gets a reference to the given []RiskPolicy and assigns it to the RiskPolicies field.
func (o *RiskPolicySet) SetRiskPolicies(v []RiskPolicy) {
	o.RiskPolicies = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskPolicySet) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskPolicySet) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RiskPolicySet) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetEvaluatedPredictors returns the EvaluatedPredictors field value if set, zero value otherwise.
func (o *RiskPolicySet) GetEvaluatedPredictors() []RiskPolicySetEvaluatedPredictorsInner {
	if o == nil || IsNil(o.EvaluatedPredictors) {
		var ret []RiskPolicySetEvaluatedPredictorsInner
		return ret
	}
	return o.EvaluatedPredictors
}

// GetEvaluatedPredictorsOk returns a tuple with the EvaluatedPredictors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetEvaluatedPredictorsOk() ([]RiskPolicySetEvaluatedPredictorsInner, bool) {
	if o == nil || IsNil(o.EvaluatedPredictors) {
		return nil, false
	}
	return o.EvaluatedPredictors, true
}

// HasEvaluatedPredictors returns a boolean if a field has been set.
func (o *RiskPolicySet) HasEvaluatedPredictors() bool {
	if o != nil && !IsNil(o.EvaluatedPredictors) {
		return true
	}

	return false
}

// SetEvaluatedPredictors gets a reference to the given []RiskPolicySetEvaluatedPredictorsInner and assigns it to the EvaluatedPredictors field.
func (o *RiskPolicySet) SetEvaluatedPredictors(v []RiskPolicySetEvaluatedPredictorsInner) {
	o.EvaluatedPredictors = v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *RiskPolicySet) GetTriggers() []RiskPolicySetTriggersInner {
	if o == nil || IsNil(o.Triggers) {
		var ret []RiskPolicySetTriggersInner
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPolicySet) GetTriggersOk() ([]RiskPolicySetTriggersInner, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *RiskPolicySet) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []RiskPolicySetTriggersInner and assigns it to the Triggers field.
func (o *RiskPolicySet) SetTriggers(v []RiskPolicySetTriggersInner) {
	o.Triggers = v
}

func (o RiskPolicySet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPolicySet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.DefaultResult) {
		toSerialize["defaultResult"] = o.DefaultResult
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.RiskPolicies) {
		toSerialize["riskPolicies"] = o.RiskPolicies
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.EvaluatedPredictors) {
		toSerialize["evaluatedPredictors"] = o.EvaluatedPredictors
	}
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	return toSerialize, nil
}

type NullableRiskPolicySet struct {
	value *RiskPolicySet
	isSet bool
}

func (v NullableRiskPolicySet) Get() *RiskPolicySet {
	return v.value
}

func (v *NullableRiskPolicySet) Set(val *RiskPolicySet) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPolicySet) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPolicySet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPolicySet(val *RiskPolicySet) *NullableRiskPolicySet {
	return &NullableRiskPolicySet{value: val, isSet: true}
}

func (v NullableRiskPolicySet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPolicySet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
