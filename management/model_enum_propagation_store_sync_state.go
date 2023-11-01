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

// EnumPropagationStoreSyncState The current state of synchronization with a propagation store or stores. Options are either SYNCING or FAILED.
type EnumPropagationStoreSyncState string

// List of EnumPropagationStoreSyncState
const (
	ENUMPROPAGATIONSTORESYNCSTATE_SYNCING EnumPropagationStoreSyncState = "SYNCING"
	ENUMPROPAGATIONSTORESYNCSTATE_FAILED EnumPropagationStoreSyncState = "FAILED"
)

// All allowed values of EnumPropagationStoreSyncState enum
var AllowedEnumPropagationStoreSyncStateEnumValues = []EnumPropagationStoreSyncState{
	"SYNCING",
	"FAILED",
}

func (v *EnumPropagationStoreSyncState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumPropagationStoreSyncState(value)
	for _, existing := range AllowedEnumPropagationStoreSyncStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumPropagationStoreSyncState(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumPropagationStoreSyncStateFromValue returns a pointer to a valid EnumPropagationStoreSyncState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumPropagationStoreSyncStateFromValue(v string) (*EnumPropagationStoreSyncState, error) {
	ev := EnumPropagationStoreSyncState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumPropagationStoreSyncState: valid values are %v", v, AllowedEnumPropagationStoreSyncStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumPropagationStoreSyncState) IsValid() bool {
	for _, existing := range AllowedEnumPropagationStoreSyncStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumPropagationStoreSyncState value
func (v EnumPropagationStoreSyncState) Ptr() *EnumPropagationStoreSyncState {
	return &v
}

type NullableEnumPropagationStoreSyncState struct {
	value *EnumPropagationStoreSyncState
	isSet bool
}

func (v NullableEnumPropagationStoreSyncState) Get() *EnumPropagationStoreSyncState {
	return v.value
}

func (v *NullableEnumPropagationStoreSyncState) Set(val *EnumPropagationStoreSyncState) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumPropagationStoreSyncState) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumPropagationStoreSyncState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumPropagationStoreSyncState(val *EnumPropagationStoreSyncState) *NullableEnumPropagationStoreSyncState {
	return &NullableEnumPropagationStoreSyncState{value: val, isSet: true}
}

func (v NullableEnumPropagationStoreSyncState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumPropagationStoreSyncState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

