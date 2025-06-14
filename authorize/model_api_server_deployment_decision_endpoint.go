/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the APIServerDeploymentDecisionEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerDeploymentDecisionEndpoint{}

// APIServerDeploymentDecisionEndpoint struct for APIServerDeploymentDecisionEndpoint
type APIServerDeploymentDecisionEndpoint struct {
	// The UUID of the decision endpoint.
	Id *string `json:"id,omitempty"`
}

// NewAPIServerDeploymentDecisionEndpoint instantiates a new APIServerDeploymentDecisionEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerDeploymentDecisionEndpoint() *APIServerDeploymentDecisionEndpoint {
	this := APIServerDeploymentDecisionEndpoint{}
	return &this
}

// NewAPIServerDeploymentDecisionEndpointWithDefaults instantiates a new APIServerDeploymentDecisionEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerDeploymentDecisionEndpointWithDefaults() *APIServerDeploymentDecisionEndpoint {
	this := APIServerDeploymentDecisionEndpoint{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *APIServerDeploymentDecisionEndpoint) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeploymentDecisionEndpoint) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *APIServerDeploymentDecisionEndpoint) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *APIServerDeploymentDecisionEndpoint) SetId(v string) {
	o.Id = &v
}

func (o APIServerDeploymentDecisionEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerDeploymentDecisionEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAPIServerDeploymentDecisionEndpoint struct {
	value *APIServerDeploymentDecisionEndpoint
	isSet bool
}

func (v NullableAPIServerDeploymentDecisionEndpoint) Get() *APIServerDeploymentDecisionEndpoint {
	return v.value
}

func (v *NullableAPIServerDeploymentDecisionEndpoint) Set(val *APIServerDeploymentDecisionEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerDeploymentDecisionEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerDeploymentDecisionEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerDeploymentDecisionEndpoint(val *APIServerDeploymentDecisionEndpoint) *NullableAPIServerDeploymentDecisionEndpoint {
	return &NullableAPIServerDeploymentDecisionEndpoint{value: val, isSet: true}
}

func (v NullableAPIServerDeploymentDecisionEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerDeploymentDecisionEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
