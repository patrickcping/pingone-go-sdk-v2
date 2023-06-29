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

// EnumCustomDomainPostHeader the model 'EnumCustomDomainPostHeader'
type EnumCustomDomainPostHeader string

// List of EnumCustomDomainPostHeader
const (
	ENUMCUSTOMDOMAINPOSTHEADER_CERTIFICATE_IMPORTJSON EnumCustomDomainPostHeader = "application/vnd.pingidentity.certificate.import+json"
	ENUMCUSTOMDOMAINPOSTHEADER_DOMAIN_NAME_VERIFYJSON EnumCustomDomainPostHeader = "application/vnd.pingidentity.domainName.verify+json"
)

// All allowed values of EnumCustomDomainPostHeader enum
var AllowedEnumCustomDomainPostHeaderEnumValues = []EnumCustomDomainPostHeader{
	"application/vnd.pingidentity.certificate.import+json",
	"application/vnd.pingidentity.domainName.verify+json",
}

func (v *EnumCustomDomainPostHeader) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumCustomDomainPostHeader(value)
	for _, existing := range AllowedEnumCustomDomainPostHeaderEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EnumCustomDomainPostHeader", value)
}

// NewEnumCustomDomainPostHeaderFromValue returns a pointer to a valid EnumCustomDomainPostHeader
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumCustomDomainPostHeaderFromValue(v string) (*EnumCustomDomainPostHeader, error) {
	ev := EnumCustomDomainPostHeader(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumCustomDomainPostHeader: valid values are %v", v, AllowedEnumCustomDomainPostHeaderEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumCustomDomainPostHeader) IsValid() bool {
	for _, existing := range AllowedEnumCustomDomainPostHeaderEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumCustomDomainPostHeader value
func (v EnumCustomDomainPostHeader) Ptr() *EnumCustomDomainPostHeader {
	return &v
}

type NullableEnumCustomDomainPostHeader struct {
	value *EnumCustomDomainPostHeader
	isSet bool
}

func (v NullableEnumCustomDomainPostHeader) Get() *EnumCustomDomainPostHeader {
	return v.value
}

func (v *NullableEnumCustomDomainPostHeader) Set(val *EnumCustomDomainPostHeader) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumCustomDomainPostHeader) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumCustomDomainPostHeader) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumCustomDomainPostHeader(val *EnumCustomDomainPostHeader) *NullableEnumCustomDomainPostHeader {
	return &NullableEnumCustomDomainPostHeader{value: val, isSet: true}
}

func (v NullableEnumCustomDomainPostHeader) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumCustomDomainPostHeader) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

