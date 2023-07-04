/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"fmt"
)

// DeviceAuthenticationPolicyPost - struct for DeviceAuthenticationPolicyPost
type DeviceAuthenticationPolicyPost struct {
	DeviceAuthenticationPolicy *DeviceAuthenticationPolicy
	DeviceAuthenticationPolicyMigrate *DeviceAuthenticationPolicyMigrate
}

// DeviceAuthenticationPolicyAsDeviceAuthenticationPolicyPost is a convenience function that returns DeviceAuthenticationPolicy wrapped in DeviceAuthenticationPolicyPost
func DeviceAuthenticationPolicyAsDeviceAuthenticationPolicyPost(v *DeviceAuthenticationPolicy) DeviceAuthenticationPolicyPost {
	return DeviceAuthenticationPolicyPost{
		DeviceAuthenticationPolicy: v,
	}
}

// DeviceAuthenticationPolicyMigrateAsDeviceAuthenticationPolicyPost is a convenience function that returns DeviceAuthenticationPolicyMigrate wrapped in DeviceAuthenticationPolicyPost
func DeviceAuthenticationPolicyMigrateAsDeviceAuthenticationPolicyPost(v *DeviceAuthenticationPolicyMigrate) DeviceAuthenticationPolicyPost {
	return DeviceAuthenticationPolicyPost{
		DeviceAuthenticationPolicyMigrate: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DeviceAuthenticationPolicyPost) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into DeviceAuthenticationPolicy
	err = json.Unmarshal(data, &dst.DeviceAuthenticationPolicy)
	if err == nil {
		if v, ok := dst.DeviceAuthenticationPolicy.GetNameOk(); ok && v != nil && *v != "" {
			match++
		} else {
			dst.DeviceAuthenticationPolicy = nil
		}
	} else {
		dst.DeviceAuthenticationPolicy = nil
	}

	// try to unmarshal data into DeviceAuthenticationPolicyMigrate
	err = json.Unmarshal(data, &dst.DeviceAuthenticationPolicyMigrate)
	if err == nil {
		if v, ok := dst.DeviceAuthenticationPolicyMigrate.GetMigrationDataOk(); ok && len(v) > 0 {
			match++
		} else {
			dst.DeviceAuthenticationPolicyMigrate = nil
		}
	} else {
		dst.DeviceAuthenticationPolicyMigrate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.DeviceAuthenticationPolicy = nil
		dst.DeviceAuthenticationPolicyMigrate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DeviceAuthenticationPolicyPost)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DeviceAuthenticationPolicyPost)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DeviceAuthenticationPolicyPost) MarshalJSON() ([]byte, error) {
	if src.DeviceAuthenticationPolicy != nil {
		return json.Marshal(&src.DeviceAuthenticationPolicy)
	}

	if src.DeviceAuthenticationPolicyMigrate != nil {
		return json.Marshal(&src.DeviceAuthenticationPolicyMigrate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DeviceAuthenticationPolicyPost) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.DeviceAuthenticationPolicy != nil {
		return obj.DeviceAuthenticationPolicy
	}

	if obj.DeviceAuthenticationPolicyMigrate != nil {
		return obj.DeviceAuthenticationPolicyMigrate
	}

	// all schemas are nil
	return nil
}

type NullableDeviceAuthenticationPolicyPost struct {
	value *DeviceAuthenticationPolicyPost
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyPost) Get() *DeviceAuthenticationPolicyPost {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyPost) Set(val *DeviceAuthenticationPolicyPost) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyPost) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyPost(val *DeviceAuthenticationPolicyPost) *NullableDeviceAuthenticationPolicyPost {
	return &NullableDeviceAuthenticationPolicyPost{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


