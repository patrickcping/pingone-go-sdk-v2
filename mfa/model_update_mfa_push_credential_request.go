/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2021-10-17
*/

package mfa

import (
	"encoding/json"
	"fmt"
)

// UpdateMFAPushCredentialRequest - struct for UpdateMFAPushCredentialRequest
type UpdateMFAPushCredentialRequest struct {
	MFAPushCredentialAPNS      *MFAPushCredentialAPNS
	MFAPushCredentialFCM       *MFAPushCredentialFCM
	MFAPushCredentialFCMHTTPV1 *MFAPushCredentialFCMHTTPV1
	MFAPushCredentialHMS       *MFAPushCredentialHMS
}

// MFAPushCredentialAPNSAsUpdateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialAPNS wrapped in UpdateMFAPushCredentialRequest
func MFAPushCredentialAPNSAsUpdateMFAPushCredentialRequest(v *MFAPushCredentialAPNS) UpdateMFAPushCredentialRequest {
	return UpdateMFAPushCredentialRequest{
		MFAPushCredentialAPNS: v,
	}
}

// MFAPushCredentialFCMAsUpdateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialFCM wrapped in UpdateMFAPushCredentialRequest
func MFAPushCredentialFCMAsUpdateMFAPushCredentialRequest(v *MFAPushCredentialFCM) UpdateMFAPushCredentialRequest {
	return UpdateMFAPushCredentialRequest{
		MFAPushCredentialFCM: v,
	}
}

// MFAPushCredentialFCMHTTPV1AsUpdateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialFCMHTTPV1 wrapped in UpdateMFAPushCredentialRequest
func MFAPushCredentialFCMHTTPV1AsUpdateMFAPushCredentialRequest(v *MFAPushCredentialFCMHTTPV1) UpdateMFAPushCredentialRequest {
	return UpdateMFAPushCredentialRequest{
		MFAPushCredentialFCMHTTPV1: v,
	}
}

// MFAPushCredentialHMSAsUpdateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialHMS wrapped in UpdateMFAPushCredentialRequest
func MFAPushCredentialHMSAsUpdateMFAPushCredentialRequest(v *MFAPushCredentialHMS) UpdateMFAPushCredentialRequest {
	return UpdateMFAPushCredentialRequest{
		MFAPushCredentialHMS: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateMFAPushCredentialRequest) UnmarshalJSON(data []byte) error {

	var common MFAPushCredential

	if err := json.Unmarshal(data, &common); err != nil { // simple model
		return err
	}

	dst.MFAPushCredentialFCM = nil
	dst.MFAPushCredentialFCMHTTPV1 = nil
	dst.MFAPushCredentialAPNS = nil
	dst.MFAPushCredentialHMS = nil

	switch common.GetType() {
	case ENUMMFAPUSHCREDENTIALATTRTYPE_FCM:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialFCM); err != nil { // simple model
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_FCM_HTTP_V1:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialFCMHTTPV1); err != nil { // simple model
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_APNS:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialAPNS); err != nil { // simple model
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_HMS:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialHMS); err != nil { // simple model
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(UpdateMFAPushCredentialRequest)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
	if src.MFAPushCredentialAPNS != nil {
		return json.Marshal(&src.MFAPushCredentialAPNS)
	}

	if src.MFAPushCredentialFCM != nil {
		return json.Marshal(&src.MFAPushCredentialFCM)
	}

	if src.MFAPushCredentialFCMHTTPV1 != nil {
		return json.Marshal(&src.MFAPushCredentialFCMHTTPV1)
	}

	if src.MFAPushCredentialHMS != nil {
		return json.Marshal(&src.MFAPushCredentialHMS)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateMFAPushCredentialRequest) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.MFAPushCredentialAPNS != nil {
		return obj.MFAPushCredentialAPNS
	}

	if obj.MFAPushCredentialFCM != nil {
		return obj.MFAPushCredentialFCM
	}

	if obj.MFAPushCredentialFCMHTTPV1 != nil {
		return obj.MFAPushCredentialFCMHTTPV1
	}

	if obj.MFAPushCredentialHMS != nil {
		return obj.MFAPushCredentialHMS
	}

	// all schemas are nil
	return nil
}

type NullableUpdateMFAPushCredentialRequest struct {
	value *UpdateMFAPushCredentialRequest
	isSet bool
}

func (v NullableUpdateMFAPushCredentialRequest) Get() *UpdateMFAPushCredentialRequest {
	return v.value
}

func (v *NullableUpdateMFAPushCredentialRequest) Set(val *UpdateMFAPushCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMFAPushCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMFAPushCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMFAPushCredentialRequest(val *UpdateMFAPushCredentialRequest) *NullableUpdateMFAPushCredentialRequest {
	return &NullableUpdateMFAPushCredentialRequest{value: val, isSet: true}
}

func (v NullableUpdateMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMFAPushCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
