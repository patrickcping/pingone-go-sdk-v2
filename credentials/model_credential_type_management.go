/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the CredentialTypeManagement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialTypeManagement{}

// CredentialTypeManagement struct for CredentialTypeManagement
type CredentialTypeManagement struct {
	Mode *EnumCredentialTypeManagementMode `json:"mode,omitempty"`
}

// NewCredentialTypeManagement instantiates a new CredentialTypeManagement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialTypeManagement() *CredentialTypeManagement {
	this := CredentialTypeManagement{}
	return &this
}

// NewCredentialTypeManagementWithDefaults instantiates a new CredentialTypeManagement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialTypeManagementWithDefaults() *CredentialTypeManagement {
	this := CredentialTypeManagement{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *CredentialTypeManagement) GetMode() EnumCredentialTypeManagementMode {
	if o == nil || IsNil(o.Mode) {
		var ret EnumCredentialTypeManagementMode
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialTypeManagement) GetModeOk() (*EnumCredentialTypeManagementMode, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *CredentialTypeManagement) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given EnumCredentialTypeManagementMode and assigns it to the Mode field.
func (o *CredentialTypeManagement) SetMode(v EnumCredentialTypeManagementMode) {
	o.Mode = &v
}

func (o CredentialTypeManagement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialTypeManagement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	return toSerialize, nil
}

type NullableCredentialTypeManagement struct {
	value *CredentialTypeManagement
	isSet bool
}

func (v NullableCredentialTypeManagement) Get() *CredentialTypeManagement {
	return v.value
}

func (v *NullableCredentialTypeManagement) Set(val *CredentialTypeManagement) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialTypeManagement) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialTypeManagement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialTypeManagement(val *CredentialTypeManagement) *NullableCredentialTypeManagement {
	return &NullableCredentialTypeManagement{value: val, isSet: true}
}

func (v NullableCredentialTypeManagement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialTypeManagement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
