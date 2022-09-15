/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

package management

import (
	"encoding/json"
	"fmt"
)

// EntityArrayEmbeddedApplicationsInner struct for EntityArrayEmbeddedApplicationsInner
type EntityArrayEmbeddedApplicationsInner struct {
	ApplicationOIDC *ApplicationOIDC
	ApplicationSAML *ApplicationSAML
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EntityArrayEmbeddedApplicationsInner) UnmarshalJSON(data []byte) error {

	var common Application

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.ApplicationOIDC = nil
	dst.ApplicationSAML = nil

	switch common.GetProtocol() {
	case ENUMAPPLICATIONPROTOCOL_OPENID_CONNECT:
		if err := json.Unmarshal(data, &dst.ApplicationOIDC); err != nil { // simple model
			return err
		}
	case ENUMAPPLICATIONPROTOCOL_SAML:
		if err := json.Unmarshal(data, &dst.ApplicationSAML); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in anyOf(EntityArrayEmbeddedApplicationsInner)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EntityArrayEmbeddedApplicationsInner) MarshalJSON() ([]byte, error) {
	if src.ApplicationOIDC != nil {
		return json.Marshal(&src.ApplicationOIDC)
	}

	if src.ApplicationSAML != nil {
		return json.Marshal(&src.ApplicationSAML)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEntityArrayEmbeddedApplicationsInner struct {
	value *EntityArrayEmbeddedApplicationsInner
	isSet bool
}

func (v NullableEntityArrayEmbeddedApplicationsInner) Get() *EntityArrayEmbeddedApplicationsInner {
	return v.value
}

func (v *NullableEntityArrayEmbeddedApplicationsInner) Set(val *EntityArrayEmbeddedApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbeddedApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbeddedApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbeddedApplicationsInner(val *EntityArrayEmbeddedApplicationsInner) *NullableEntityArrayEmbeddedApplicationsInner {
	return &NullableEntityArrayEmbeddedApplicationsInner{value: val, isSet: true}
}

func (v NullableEntityArrayEmbeddedApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbeddedApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
