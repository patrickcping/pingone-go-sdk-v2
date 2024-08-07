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

// checks if the IntegrationVersionSAMLThirdParty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationVersionSAMLThirdParty{}

// IntegrationVersionSAMLThirdParty Third-party IdP information.
type IntegrationVersionSAMLThirdParty struct {
	Metadata *IntegrationVersionSAMLThirdPartyMetadata `json:"metadata,omitempty"`
	Instructions *IntegrationVersionSAMLThirdPartyInstructions `json:"instructions,omitempty"`
}

// NewIntegrationVersionSAMLThirdParty instantiates a new IntegrationVersionSAMLThirdParty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationVersionSAMLThirdParty() *IntegrationVersionSAMLThirdParty {
	this := IntegrationVersionSAMLThirdParty{}
	return &this
}

// NewIntegrationVersionSAMLThirdPartyWithDefaults instantiates a new IntegrationVersionSAMLThirdParty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationVersionSAMLThirdPartyWithDefaults() *IntegrationVersionSAMLThirdParty {
	this := IntegrationVersionSAMLThirdParty{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *IntegrationVersionSAMLThirdParty) GetMetadata() IntegrationVersionSAMLThirdPartyMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret IntegrationVersionSAMLThirdPartyMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionSAMLThirdParty) GetMetadataOk() (*IntegrationVersionSAMLThirdPartyMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *IntegrationVersionSAMLThirdParty) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given IntegrationVersionSAMLThirdPartyMetadata and assigns it to the Metadata field.
func (o *IntegrationVersionSAMLThirdParty) SetMetadata(v IntegrationVersionSAMLThirdPartyMetadata) {
	o.Metadata = &v
}

// GetInstructions returns the Instructions field value if set, zero value otherwise.
func (o *IntegrationVersionSAMLThirdParty) GetInstructions() IntegrationVersionSAMLThirdPartyInstructions {
	if o == nil || IsNil(o.Instructions) {
		var ret IntegrationVersionSAMLThirdPartyInstructions
		return ret
	}
	return *o.Instructions
}

// GetInstructionsOk returns a tuple with the Instructions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionSAMLThirdParty) GetInstructionsOk() (*IntegrationVersionSAMLThirdPartyInstructions, bool) {
	if o == nil || IsNil(o.Instructions) {
		return nil, false
	}
	return o.Instructions, true
}

// HasInstructions returns a boolean if a field has been set.
func (o *IntegrationVersionSAMLThirdParty) HasInstructions() bool {
	if o != nil && !IsNil(o.Instructions) {
		return true
	}

	return false
}

// SetInstructions gets a reference to the given IntegrationVersionSAMLThirdPartyInstructions and assigns it to the Instructions field.
func (o *IntegrationVersionSAMLThirdParty) SetInstructions(v IntegrationVersionSAMLThirdPartyInstructions) {
	o.Instructions = &v
}

func (o IntegrationVersionSAMLThirdParty) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationVersionSAMLThirdParty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Instructions) {
		toSerialize["instructions"] = o.Instructions
	}
	return toSerialize, nil
}

type NullableIntegrationVersionSAMLThirdParty struct {
	value *IntegrationVersionSAMLThirdParty
	isSet bool
}

func (v NullableIntegrationVersionSAMLThirdParty) Get() *IntegrationVersionSAMLThirdParty {
	return v.value
}

func (v *NullableIntegrationVersionSAMLThirdParty) Set(val *IntegrationVersionSAMLThirdParty) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationVersionSAMLThirdParty) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationVersionSAMLThirdParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationVersionSAMLThirdParty(val *IntegrationVersionSAMLThirdParty) *NullableIntegrationVersionSAMLThirdParty {
	return &NullableIntegrationVersionSAMLThirdParty{value: val, isSet: true}
}

func (v NullableIntegrationVersionSAMLThirdParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationVersionSAMLThirdParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


