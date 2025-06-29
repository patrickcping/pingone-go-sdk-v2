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

// checks if the SignOnPolicyActionPingIDWinLoginPasswordless type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionPingIDWinLoginPasswordless{}

// SignOnPolicyActionPingIDWinLoginPasswordless struct for SignOnPolicyActionPingIDWinLoginPasswordless
type SignOnPolicyActionPingIDWinLoginPasswordless struct {
	Links       *map[string]LinksHATEOASValue               `json:"_links,omitempty"`
	Condition   *SignOnPolicyActionCommonConditionOrOrInner `json:"condition,omitempty"`
	Environment *ObjectEnvironment                          `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// An integer that specifies the order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first. This is a required property.
	Priority            int32                                                                `json:"priority"`
	SignOnPolicy        *SignOnPolicyActionCommonSignOnPolicy                                `json:"signOnPolicy,omitempty"`
	Type                EnumSignOnPolicyType                                                 `json:"type"`
	UniqueUserAttribute SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute `json:"uniqueUserAttribute"`
	OfflineMode         SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode         `json:"offlineMode"`
}

// NewSignOnPolicyActionPingIDWinLoginPasswordless instantiates a new SignOnPolicyActionPingIDWinLoginPasswordless object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionPingIDWinLoginPasswordless(priority int32, type_ EnumSignOnPolicyType, uniqueUserAttribute SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute, offlineMode SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) *SignOnPolicyActionPingIDWinLoginPasswordless {
	this := SignOnPolicyActionPingIDWinLoginPasswordless{}
	this.Priority = priority
	this.Type = type_
	this.UniqueUserAttribute = uniqueUserAttribute
	this.OfflineMode = offlineMode
	return &this
}

// NewSignOnPolicyActionPingIDWinLoginPasswordlessWithDefaults instantiates a new SignOnPolicyActionPingIDWinLoginPasswordless object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionPingIDWinLoginPasswordlessWithDefaults() *SignOnPolicyActionPingIDWinLoginPasswordless {
	this := SignOnPolicyActionPingIDWinLoginPasswordless{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetCondition() SignOnPolicyActionCommonConditionOrOrInner {
	if o == nil || IsNil(o.Condition) {
		var ret SignOnPolicyActionCommonConditionOrOrInner
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetConditionOk() (*SignOnPolicyActionCommonConditionOrOrInner, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given SignOnPolicyActionCommonConditionOrOrInner and assigns it to the Condition field.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetCondition(v SignOnPolicyActionCommonConditionOrOrInner) {
	o.Condition = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value if set, zero value otherwise.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil || IsNil(o.SignOnPolicy) {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}
	return *o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil || IsNil(o.SignOnPolicy) {
		return nil, false
	}
	return o.SignOnPolicy, true
}

// HasSignOnPolicy returns a boolean if a field has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) HasSignOnPolicy() bool {
	if o != nil && !IsNil(o.SignOnPolicy) {
		return true
	}

	return false
}

// SetSignOnPolicy gets a reference to the given SignOnPolicyActionCommonSignOnPolicy and assigns it to the SignOnPolicy field.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = &v
}

// GetType returns the Type field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetType() EnumSignOnPolicyType {
	if o == nil {
		var ret EnumSignOnPolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetTypeOk() (*EnumSignOnPolicyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetType(v EnumSignOnPolicyType) {
	o.Type = v
}

// GetUniqueUserAttribute returns the UniqueUserAttribute field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetUniqueUserAttribute() SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute {
	if o == nil {
		var ret SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute
		return ret
	}

	return o.UniqueUserAttribute
}

// GetUniqueUserAttributeOk returns a tuple with the UniqueUserAttribute field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetUniqueUserAttributeOk() (*SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UniqueUserAttribute, true
}

// SetUniqueUserAttribute sets field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetUniqueUserAttribute(v SignOnPolicyActionPingIDWinLoginPasswordlessAllOfUniqueUserAttribute) {
	o.UniqueUserAttribute = v
}

// GetOfflineMode returns the OfflineMode field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetOfflineMode() SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode {
	if o == nil {
		var ret SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode
		return ret
	}

	return o.OfflineMode
}

// GetOfflineModeOk returns a tuple with the OfflineMode field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) GetOfflineModeOk() (*SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfflineMode, true
}

// SetOfflineMode sets field value
func (o *SignOnPolicyActionPingIDWinLoginPasswordless) SetOfflineMode(v SignOnPolicyActionPingIDWinLoginPasswordlessAllOfOfflineMode) {
	o.OfflineMode = v
}

func (o SignOnPolicyActionPingIDWinLoginPasswordless) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionPingIDWinLoginPasswordless) ToMap() (map[string]interface{}, error) {
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
	toSerialize["uniqueUserAttribute"] = o.UniqueUserAttribute
	toSerialize["offlineMode"] = o.OfflineMode
	return toSerialize, nil
}

type NullableSignOnPolicyActionPingIDWinLoginPasswordless struct {
	value *SignOnPolicyActionPingIDWinLoginPasswordless
	isSet bool
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordless) Get() *SignOnPolicyActionPingIDWinLoginPasswordless {
	return v.value
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordless) Set(val *SignOnPolicyActionPingIDWinLoginPasswordless) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordless) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordless) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionPingIDWinLoginPasswordless(val *SignOnPolicyActionPingIDWinLoginPasswordless) *NullableSignOnPolicyActionPingIDWinLoginPasswordless {
	return &NullableSignOnPolicyActionPingIDWinLoginPasswordless{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionPingIDWinLoginPasswordless) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionPingIDWinLoginPasswordless) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
