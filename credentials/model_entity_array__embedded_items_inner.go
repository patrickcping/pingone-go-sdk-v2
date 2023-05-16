/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-03-30
*/

package credentials

import (
	"encoding/json"
	"fmt"
)

// EntityArrayEmbeddedItemsInner struct for EntityArrayEmbeddedItemsInner
type EntityArrayEmbeddedItemsInner struct {
	CredentialType *CredentialType
	UserCredential *UserCredential
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EntityArrayEmbeddedItemsInner) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CredentialType
	err = json.Unmarshal(data, &dst.CredentialType)
	if err == nil {
		jsonCredentialType, _ := json.Marshal(dst.CredentialType)
		if string(jsonCredentialType) == "{}" { // empty struct
			dst.CredentialType = nil
		} else {
			if dst.CredentialType.CardDesignTemplate != "" {
				return nil // data stored in dst.CredentialType, return on the first match
			} else {
				dst.CredentialType = nil
			}
		}
	} else {
		dst.CredentialType = nil
	}

	// try to unmarshal JSON data into UserCredential
	err = json.Unmarshal(data, &dst.UserCredential)
	if err == nil {
		jsonUserCredential, _ := json.Marshal(dst.UserCredential)
		if string(jsonUserCredential) == "{}" { // empty struct
			dst.UserCredential = nil
		} else {
			if dst.UserCredential.User.Id != "" {
				return nil // data stored in dst.UserCredential, return on the first match
			} else {
				dst.UserCredential = nil
			}
		}
	} else {
		dst.UserCredential = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EntityArrayEmbeddedItemsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EntityArrayEmbeddedItemsInner) MarshalJSON() ([]byte, error) {
	if src.CredentialType != nil {
		return json.Marshal(&src.CredentialType)
	}

	if src.UserCredential != nil {
		return json.Marshal(&src.UserCredential)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEntityArrayEmbeddedItemsInner struct {
	value *EntityArrayEmbeddedItemsInner
	isSet bool
}

func (v NullableEntityArrayEmbeddedItemsInner) Get() *EntityArrayEmbeddedItemsInner {
	return v.value
}

func (v *NullableEntityArrayEmbeddedItemsInner) Set(val *EntityArrayEmbeddedItemsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbeddedItemsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbeddedItemsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbeddedItemsInner(val *EntityArrayEmbeddedItemsInner) *NullableEntityArrayEmbeddedItemsInner {
	return &NullableEntityArrayEmbeddedItemsInner{value: val, isSet: true}
}

func (v NullableEntityArrayEmbeddedItemsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbeddedItemsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
