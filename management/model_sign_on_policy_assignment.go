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

// checks if the SignOnPolicyAssignment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyAssignment{}

// SignOnPolicyAssignment struct for SignOnPolicyAssignment
type SignOnPolicyAssignment struct {
	Application *ObjectApplication `json:"application,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the sign-on policy assignment resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// The order in which the policy referenced by this assignment is evaluated during an authentication flow relative to other policies. An assignment with a lower priority will be evaluated first.
	Priority int32 `json:"priority"`
	SignOnPolicy SignOnPolicyActionCommonSignOnPolicy `json:"signOnPolicy"`
}

// NewSignOnPolicyAssignment instantiates a new SignOnPolicyAssignment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyAssignment(priority int32, signOnPolicy SignOnPolicyActionCommonSignOnPolicy) *SignOnPolicyAssignment {
	this := SignOnPolicyAssignment{}
	this.Priority = priority
	this.SignOnPolicy = signOnPolicy
	return &this
}

// NewSignOnPolicyAssignmentWithDefaults instantiates a new SignOnPolicyAssignment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyAssignmentWithDefaults() *SignOnPolicyAssignment {
	this := SignOnPolicyAssignment{}
	return &this
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *SignOnPolicyAssignment) GetApplication() ObjectApplication {
	if o == nil || IsNil(o.Application) {
		var ret ObjectApplication
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAssignment) GetApplicationOk() (*ObjectApplication, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *SignOnPolicyAssignment) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given ObjectApplication and assigns it to the Application field.
func (o *SignOnPolicyAssignment) SetApplication(v ObjectApplication) {
	o.Application = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SignOnPolicyAssignment) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAssignment) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SignOnPolicyAssignment) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SignOnPolicyAssignment) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SignOnPolicyAssignment) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAssignment) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SignOnPolicyAssignment) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SignOnPolicyAssignment) SetId(v string) {
	o.Id = &v
}

// GetPriority returns the Priority field value
func (o *SignOnPolicyAssignment) GetPriority() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAssignment) GetPriorityOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Priority, true
}

// SetPriority sets field value
func (o *SignOnPolicyAssignment) SetPriority(v int32) {
	o.Priority = v
}

// GetSignOnPolicy returns the SignOnPolicy field value
func (o *SignOnPolicyAssignment) GetSignOnPolicy() SignOnPolicyActionCommonSignOnPolicy {
	if o == nil {
		var ret SignOnPolicyActionCommonSignOnPolicy
		return ret
	}

	return o.SignOnPolicy
}

// GetSignOnPolicyOk returns a tuple with the SignOnPolicy field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyAssignment) GetSignOnPolicyOk() (*SignOnPolicyActionCommonSignOnPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignOnPolicy, true
}

// SetSignOnPolicy sets field value
func (o *SignOnPolicyAssignment) SetSignOnPolicy(v SignOnPolicyActionCommonSignOnPolicy) {
	o.SignOnPolicy = v
}

func (o SignOnPolicyAssignment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyAssignment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	toSerialize["priority"] = o.Priority
	toSerialize["signOnPolicy"] = o.SignOnPolicy
	return toSerialize, nil
}

type NullableSignOnPolicyAssignment struct {
	value *SignOnPolicyAssignment
	isSet bool
}

func (v NullableSignOnPolicyAssignment) Get() *SignOnPolicyAssignment {
	return v.value
}

func (v *NullableSignOnPolicyAssignment) Set(val *SignOnPolicyAssignment) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyAssignment) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyAssignment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyAssignment(val *SignOnPolicyAssignment) *NullableSignOnPolicyAssignment {
	return &NullableSignOnPolicyAssignment{value: val, isSet: true}
}

func (v NullableSignOnPolicyAssignment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyAssignment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


