/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// checks if the FlowPolicyAssignmentFlowPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlowPolicyAssignmentFlowPolicy{}

// FlowPolicyAssignmentFlowPolicy An object that contains details of the flow policy resource.
type FlowPolicyAssignmentFlowPolicy struct {
	// A string that specifies the flow policy resource ID associated with the flow policy assignment.
	Id string `json:"id"`
}

type _FlowPolicyAssignmentFlowPolicy FlowPolicyAssignmentFlowPolicy

// NewFlowPolicyAssignmentFlowPolicy instantiates a new FlowPolicyAssignmentFlowPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlowPolicyAssignmentFlowPolicy(id string) *FlowPolicyAssignmentFlowPolicy {
	this := FlowPolicyAssignmentFlowPolicy{}
	this.Id = id
	return &this
}

// NewFlowPolicyAssignmentFlowPolicyWithDefaults instantiates a new FlowPolicyAssignmentFlowPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlowPolicyAssignmentFlowPolicyWithDefaults() *FlowPolicyAssignmentFlowPolicy {
	this := FlowPolicyAssignmentFlowPolicy{}
	return &this
}

// GetId returns the Id field value
func (o *FlowPolicyAssignmentFlowPolicy) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FlowPolicyAssignmentFlowPolicy) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FlowPolicyAssignmentFlowPolicy) SetId(v string) {
	o.Id = v
}

func (o FlowPolicyAssignmentFlowPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlowPolicyAssignmentFlowPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *FlowPolicyAssignmentFlowPolicy) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFlowPolicyAssignmentFlowPolicy := _FlowPolicyAssignmentFlowPolicy{}

	err = json.Unmarshal(bytes, &varFlowPolicyAssignmentFlowPolicy)

	if err != nil {
		return err
	}

	*o = FlowPolicyAssignmentFlowPolicy(varFlowPolicyAssignmentFlowPolicy)

	return err
}

type NullableFlowPolicyAssignmentFlowPolicy struct {
	value *FlowPolicyAssignmentFlowPolicy
	isSet bool
}

func (v NullableFlowPolicyAssignmentFlowPolicy) Get() *FlowPolicyAssignmentFlowPolicy {
	return v.value
}

func (v *NullableFlowPolicyAssignmentFlowPolicy) Set(val *FlowPolicyAssignmentFlowPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableFlowPolicyAssignmentFlowPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableFlowPolicyAssignmentFlowPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlowPolicyAssignmentFlowPolicy(val *FlowPolicyAssignmentFlowPolicy) *NullableFlowPolicyAssignmentFlowPolicy {
	return &NullableFlowPolicyAssignmentFlowPolicy{value: val, isSet: true}
}

func (v NullableFlowPolicyAssignmentFlowPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlowPolicyAssignmentFlowPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


