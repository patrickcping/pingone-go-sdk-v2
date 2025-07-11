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

// checks if the SignOnPolicyActionProgressiveProfiling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionProgressiveProfiling{}

// SignOnPolicyActionProgressiveProfiling struct for SignOnPolicyActionProgressiveProfiling
type SignOnPolicyActionProgressiveProfiling struct {
	Links       *map[string]LinksHATEOASValue               `json:"_links,omitempty"`
	Condition   *SignOnPolicyActionCommonConditionOrOrInner `json:"condition,omitempty"`
	Environment *ObjectEnvironment                          `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property.
	Priority     int32                                                   `json:"priority"`
	SignOnPolicy *SignOnPolicyActionCommonSignOnPolicy                   `json:"signOnPolicy,omitempty"`
	Type         EnumSignOnPolicyType                                    `json:"type"`
	Attributes   []SignOnPolicyActionProgressiveProfilingAllOfAttributes `json:"attributes"`
	// A boolean that specifies whether the progressive profiling action will not be executed if another progressive profiling action has already been executed during the flow. This property is required.
	PreventMultiplePromptsPerFlow bool `json:"preventMultiplePromptsPerFlow"`
	// An integer that specifies how often to prompt the user to provide profile data for the configured attributes for which they do not have values. This property is required.
	PromptIntervalSeconds int32 `json:"promptIntervalSeconds"`
	// A string that specifies text to display to the user when prompting for attribute values. This property is required.
	PromptText string `json:"promptText"`
}

// NewSignOnPolicyActionProgressiveProfiling instantiates a new SignOnPolicyActionProgressiveProfiling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionProgressiveProfiling(priority int32, type_ EnumSignOnPolicyType, attributes []SignOnPolicyActionProgressiveProfilingAllOfAttributes, preventMultiplePromptsPerFlow bool, promptIntervalSeconds int32, promptText string) *SignOnPolicyActionProgressiveProfiling {
	this := SignOnPolicyActionProgressiveProfiling{}
	this.Priority = priority
	this.Type = type_
	this.Attributes = attributes
	this.PreventMultiplePromptsPerFlow = preventMultiplePromptsPerFlow
	this.PromptIntervalSeconds = promptIntervalSeconds
	this.PromptText = promptText
	return &this
}

// NewSignOnPolicyActionProgressiveProfilingWithDefaults instantiates a new SignOnPolicyActionProgressiveProfiling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionProgressiveProfilingWithDefaults() *SignOnPolicyActionProgressiveProfiling {
	this := SignOnPolicyActionProgressiveProfiling{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignOnPolicyActionProgressiveProfiling) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SignOnPolicyActionProgressiveProfiling) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *SignOnPolicyActionProgressiveProfiling) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SignOnPolicyActionProgressiveProfiling) GetCondition() SignOnPolicyActionCommonConditionOrOrInner {
	if o == nil || IsNil(o.Condition) {
		var ret SignOnPolicyActionCommonConditionOrOrInner
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SignOnPolicyActionProgressiveProfiling) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given SignOnPolicyActionCommonConditionOrOrInner and assigns it to the Condition field.
func (o *SignOnPolicyActionProgressiveProfiling) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner) {
	o.Condition = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyActionProgressiveProfiling) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyActionProgressiveProfiling) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyActionProgressiveProfiling) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyActionProgressiveProfiling) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyActionProgressiveProfiling) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyActionProgressiveProfiling) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyActionProgressiveProfiling) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyActionProgressiveProfiling) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value if set, zero value otherwise.
func (o *SignOnPolicyActionProgressiveProfiling) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil || IsNil(o.SignOnPolicy) {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}
	return *o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil || IsNil(o.SignOnPolicy) {
		return nil, false
	}
	return o.SignOnPolicy, true
}

// HasSignOnPolicy returns a boolean if a field has been set.
func (o *SignOnPolicyActionProgressiveProfiling) HasSignOnPolicy() bool {
	if o != nil && !IsNil(o.SignOnPolicy) {
		return true
	}

	return false
}

// SetSignOnPolicy gets a reference to the given SignOnPolicyActionCommonSignOnPolicy and assigns it to the SignOnPolicy field.
func (o *SignOnPolicyActionProgressiveProfiling) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = &v
}

// GetType returns the Type field value
func (o *SignOnPolicyActionProgressiveProfiling) GetType() EnumSignOnPolicyType {
	if o == nil {
		var ret EnumSignOnPolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetTypeOk() (*EnumSignOnPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignOnPolicyActionProgressiveProfiling) SetType(v EnumSignOnPolicyType) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SignOnPolicyActionProgressiveProfiling) GetAttributes() []SignOnPolicyActionProgressiveProfilingAllOfAttributes {
	if o == nil {
		var ret []SignOnPolicyActionProgressiveProfilingAllOfAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetAttributesOk() ([]SignOnPolicyActionProgressiveProfilingAllOfAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *SignOnPolicyActionProgressiveProfiling) SetAttributes(v []SignOnPolicyActionProgressiveProfilingAllOfAttributes) {
	o.Attributes = v
}

// GetPreventMultiplePromptsPerFlow returns the PreventMultiplePromptsPerFlow field value
func (o *SignOnPolicyActionProgressiveProfiling) GetPreventMultiplePromptsPerFlow() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PreventMultiplePromptsPerFlow
}

// GetPreventMultiplePromptsPerFlowOk returns a tuple with the PreventMultiplePromptsPerFlow field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetPreventMultiplePromptsPerFlowOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PreventMultiplePromptsPerFlow, true
}

// SetPreventMultiplePromptsPerFlow sets field value
func (o *SignOnPolicyActionProgressiveProfiling) SetPreventMultiplePromptsPerFlow(v bool) {
	o.PreventMultiplePromptsPerFlow = v
}

// GetPromptIntervalSeconds returns the PromptIntervalSeconds field value
func (o *SignOnPolicyActionProgressiveProfiling) GetPromptIntervalSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PromptIntervalSeconds
}

// GetPromptIntervalSecondsOk returns a tuple with the PromptIntervalSeconds field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetPromptIntervalSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptIntervalSeconds, true
}

// SetPromptIntervalSeconds sets field value
func (o *SignOnPolicyActionProgressiveProfiling) SetPromptIntervalSeconds(v int32) {
	o.PromptIntervalSeconds = v
}

// GetPromptText returns the PromptText field value
func (o *SignOnPolicyActionProgressiveProfiling) GetPromptText() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PromptText
}

// GetPromptTextOk returns a tuple with the PromptText field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionProgressiveProfiling) GetPromptTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PromptText, true
}

// SetPromptText sets field value
func (o *SignOnPolicyActionProgressiveProfiling) SetPromptText(v string) {
	o.PromptText = v
}

func (o SignOnPolicyActionProgressiveProfiling) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionProgressiveProfiling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["priority"] = o.Priority
	if !IsNil(o.SignOnPolicy) {
		toSerialize["signOnPolicy"] = o.SignOnPolicy
	}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["preventMultiplePromptsPerFlow"] = o.PreventMultiplePromptsPerFlow
	toSerialize["promptIntervalSeconds"] = o.PromptIntervalSeconds
	toSerialize["promptText"] = o.PromptText
	return toSerialize, nil
}

type NullableSignOnPolicyActionProgressiveProfiling struct {
	value *SignOnPolicyActionProgressiveProfiling
	isSet bool
}

func (v NullableSignOnPolicyActionProgressiveProfiling) Get() *SignOnPolicyActionProgressiveProfiling {
	return v.value
}

func (v *NullableSignOnPolicyActionProgressiveProfiling) Set(val *SignOnPolicyActionProgressiveProfiling) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionProgressiveProfiling) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionProgressiveProfiling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionProgressiveProfiling(val *SignOnPolicyActionProgressiveProfiling) *NullableSignOnPolicyActionProgressiveProfiling {
	return &NullableSignOnPolicyActionProgressiveProfiling{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionProgressiveProfiling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionProgressiveProfiling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
