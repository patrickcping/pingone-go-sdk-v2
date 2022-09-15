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

// UpdateApplicationRequest - struct for UpdateApplicationRequest
type UpdateApplicationRequest struct {
	ApplicationOIDC *ApplicationOIDC
	ApplicationSAML *ApplicationSAML
}

// ApplicationOIDCAsUpdateApplicationRequest is a convenience function that returns ApplicationOIDC wrapped in UpdateApplicationRequest
func ApplicationOIDCAsUpdateApplicationRequest(v *ApplicationOIDC) UpdateApplicationRequest {
	return UpdateApplicationRequest{
		ApplicationOIDC: v,
	}
}

// ApplicationSAMLAsUpdateApplicationRequest is a convenience function that returns ApplicationSAML wrapped in UpdateApplicationRequest
func ApplicationSAMLAsUpdateApplicationRequest(v *ApplicationSAML) UpdateApplicationRequest {
	return UpdateApplicationRequest{
		ApplicationSAML: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateApplicationRequest) UnmarshalJSON(data []byte) error {

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
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateApplicationRequest)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	if src.ApplicationOIDC != nil {
		return json.Marshal(&src.ApplicationOIDC)
	}

	if src.ApplicationSAML != nil {
		return json.Marshal(&src.ApplicationSAML)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateApplicationRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ApplicationOIDC != nil {
		return obj.ApplicationOIDC
	}

	if obj.ApplicationSAML != nil {
		return obj.ApplicationSAML
	}

	// all schemas are nil
	return nil
}

type NullableUpdateApplicationRequest struct {
	value *UpdateApplicationRequest
	isSet bool
}

func (v NullableUpdateApplicationRequest) Get() *UpdateApplicationRequest {
	return v.value
}

func (v *NullableUpdateApplicationRequest) Set(val *UpdateApplicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateApplicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateApplicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateApplicationRequest(val *UpdateApplicationRequest) *NullableUpdateApplicationRequest {
	return &NullableUpdateApplicationRequest{value: val, isSet: true}
}

func (v NullableUpdateApplicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateApplicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
