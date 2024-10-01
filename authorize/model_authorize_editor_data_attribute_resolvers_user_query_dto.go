/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"fmt"
)

// AuthorizeEditorDataAttributeResolversUserQueryDTO - struct for AuthorizeEditorDataAttributeResolversUserQueryDTO
type AuthorizeEditorDataAttributeResolversUserQueryDTO struct {
	AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO *AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO
}

// AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTOAsAuthorizeEditorDataAttributeResolversUserQueryDTO is a convenience function that returns AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO wrapped in AuthorizeEditorDataAttributeResolversUserQueryDTO
func AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTOAsAuthorizeEditorDataAttributeResolversUserQueryDTO(v *AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO) AuthorizeEditorDataAttributeResolversUserQueryDTO {
	return AuthorizeEditorDataAttributeResolversUserQueryDTO{
		AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AuthorizeEditorDataAttributeResolversUserQueryDTO) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO
	err = newStrictDecoder(data).Decode(&dst.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO)
	if err == nil {
		jsonAuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO, _ := json.Marshal(dst.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO)
		if string(jsonAuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO) == "{}" { // empty struct
			dst.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO = nil
		} else {
			match++
		}
	} else {
		dst.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AuthorizeEditorDataAttributeResolversUserQueryDTO)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AuthorizeEditorDataAttributeResolversUserQueryDTO)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AuthorizeEditorDataAttributeResolversUserQueryDTO) MarshalJSON() ([]byte, error) {
	if src.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO != nil {
		return json.Marshal(&src.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AuthorizeEditorDataAttributeResolversUserQueryDTO) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO != nil {
		return obj.AuthorizeEditorDataAttributeResolversUserQueryUserIdQueryDTO
	}

	// all schemas are nil
	return nil
}

type NullableAuthorizeEditorDataAttributeResolversUserQueryDTO struct {
	value *AuthorizeEditorDataAttributeResolversUserQueryDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataAttributeResolversUserQueryDTO) Get() *AuthorizeEditorDataAttributeResolversUserQueryDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataAttributeResolversUserQueryDTO) Set(val *AuthorizeEditorDataAttributeResolversUserQueryDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataAttributeResolversUserQueryDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataAttributeResolversUserQueryDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataAttributeResolversUserQueryDTO(val *AuthorizeEditorDataAttributeResolversUserQueryDTO) *NullableAuthorizeEditorDataAttributeResolversUserQueryDTO {
	return &NullableAuthorizeEditorDataAttributeResolversUserQueryDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataAttributeResolversUserQueryDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataAttributeResolversUserQueryDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


