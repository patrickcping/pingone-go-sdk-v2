/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"fmt"
)

// CreateMFAPushCredential201Response - struct for CreateMFAPushCredential201Response
type CreateMFAPushCredential201Response struct {
	MFAPushCredential *MFAPushCredential
	MFAPushCredentialAPNS *MFAPushCredentialAPNS
}

// MFAPushCredentialAsCreateMFAPushCredential201Response is a convenience function that returns MFAPushCredential wrapped in CreateMFAPushCredential201Response
func MFAPushCredentialAsCreateMFAPushCredential201Response(v *MFAPushCredential) CreateMFAPushCredential201Response {
	return CreateMFAPushCredential201Response{
		MFAPushCredential: v,
	}
}

// MFAPushCredentialAPNSAsCreateMFAPushCredential201Response is a convenience function that returns MFAPushCredentialAPNS wrapped in CreateMFAPushCredential201Response
func MFAPushCredentialAPNSAsCreateMFAPushCredential201Response(v *MFAPushCredentialAPNS) CreateMFAPushCredential201Response {
	return CreateMFAPushCredential201Response{
		MFAPushCredentialAPNS: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateMFAPushCredential201Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MFAPushCredential
	err = newStrictDecoder(data).Decode(&dst.MFAPushCredential)
	if err == nil {
		jsonMFAPushCredential, _ := json.Marshal(dst.MFAPushCredential)
		if string(jsonMFAPushCredential) == "{}" { // empty struct
			dst.MFAPushCredential = nil
		} else {
			match++
		}
	} else {
		dst.MFAPushCredential = nil
	}

	// try to unmarshal data into MFAPushCredentialAPNS
	err = newStrictDecoder(data).Decode(&dst.MFAPushCredentialAPNS)
	if err == nil {
		jsonMFAPushCredentialAPNS, _ := json.Marshal(dst.MFAPushCredentialAPNS)
		if string(jsonMFAPushCredentialAPNS) == "{}" { // empty struct
			dst.MFAPushCredentialAPNS = nil
		} else {
			match++
		}
	} else {
		dst.MFAPushCredentialAPNS = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MFAPushCredential = nil
		dst.MFAPushCredentialAPNS = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(CreateMFAPushCredential201Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(CreateMFAPushCredential201Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateMFAPushCredential201Response) MarshalJSON() ([]byte, error) {
	if src.MFAPushCredential != nil {
		return json.Marshal(&src.MFAPushCredential)
	}

	if src.MFAPushCredentialAPNS != nil {
		return json.Marshal(&src.MFAPushCredentialAPNS)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *CreateMFAPushCredential201Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MFAPushCredential != nil {
		return obj.MFAPushCredential
	}

	if obj.MFAPushCredentialAPNS != nil {
		return obj.MFAPushCredentialAPNS
	}

	// all schemas are nil
	return nil
}

type NullableCreateMFAPushCredential201Response struct {
	value *CreateMFAPushCredential201Response
	isSet bool
}

func (v NullableCreateMFAPushCredential201Response) Get() *CreateMFAPushCredential201Response {
	return v.value
}

func (v *NullableCreateMFAPushCredential201Response) Set(val *CreateMFAPushCredential201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMFAPushCredential201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMFAPushCredential201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMFAPushCredential201Response(val *CreateMFAPushCredential201Response) *NullableCreateMFAPushCredential201Response {
	return &NullableCreateMFAPushCredential201Response{value: val, isSet: true}
}

func (v NullableCreateMFAPushCredential201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMFAPushCredential201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


