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

// MFAPushCredentialRequest - struct for MFAPushCredentialRequest
type MFAPushCredentialRequest struct {
	MFAPushCredentialAPNS *MFAPushCredentialAPNS
	MFAPushCredentialFCM *MFAPushCredentialFCM
	MFAPushCredentialFCMHTTPV1 *MFAPushCredentialFCMHTTPV1
	MFAPushCredentialHMS *MFAPushCredentialHMS
}

// MFAPushCredentialAPNSAsMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialAPNS wrapped in MFAPushCredentialRequest
func MFAPushCredentialAPNSAsMFAPushCredentialRequest(v *MFAPushCredentialAPNS) MFAPushCredentialRequest {
	return MFAPushCredentialRequest{
		MFAPushCredentialAPNS: v,
	}
}

// MFAPushCredentialFCMAsMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialFCM wrapped in MFAPushCredentialRequest
func MFAPushCredentialFCMAsMFAPushCredentialRequest(v *MFAPushCredentialFCM) MFAPushCredentialRequest {
	return MFAPushCredentialRequest{
		MFAPushCredentialFCM: v,
	}
}

// MFAPushCredentialFCMHTTPV1AsMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialFCMHTTPV1 wrapped in MFAPushCredentialRequest
func MFAPushCredentialFCMHTTPV1AsMFAPushCredentialRequest(v *MFAPushCredentialFCMHTTPV1) MFAPushCredentialRequest {
	return MFAPushCredentialRequest{
		MFAPushCredentialFCMHTTPV1: v,
	}
}

// MFAPushCredentialHMSAsMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialHMS wrapped in MFAPushCredentialRequest
func MFAPushCredentialHMSAsMFAPushCredentialRequest(v *MFAPushCredentialHMS) MFAPushCredentialRequest {
	return MFAPushCredentialRequest{
		MFAPushCredentialHMS: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MFAPushCredentialRequest) UnmarshalJSON(data []byte) error {

	var common MFAPushCredential

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.MFAPushCredentialAPNS = nil
	dst.MFAPushCredentialFCM = nil
	dst.MFAPushCredentialFCMHTTPV1 = nil
	dst.MFAPushCredentialHMS = nil

	objType := common.GetType()

	if !objType.IsValid() {
		return nil
	}

	switch objType {
	case ENUMMFAPUSHCREDENTIALATTRTYPE_APNS:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialAPNS); err != nil {
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_FCM:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialFCM); err != nil {
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_HMS:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialHMS); err != nil {
			return err
		}
	case ENUMMFAPUSHCREDENTIALATTRTYPE_FCM_HTTP_V1:
		if err := json.Unmarshal(data, &dst.MFAPushCredentialFCMHTTPV1); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(MFAPushCredentialRequest)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
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
func (obj *MFAPushCredentialRequest) GetActualInstance() (interface{}) {
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

type NullableMFAPushCredentialRequest struct {
	value *MFAPushCredentialRequest
	isSet bool
}

func (v NullableMFAPushCredentialRequest) Get() *MFAPushCredentialRequest {
	return v.value
}

func (v *NullableMFAPushCredentialRequest) Set(val *MFAPushCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMFAPushCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMFAPushCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFAPushCredentialRequest(val *MFAPushCredentialRequest) *NullableMFAPushCredentialRequest {
	return &NullableMFAPushCredentialRequest{value: val, isSet: true}
}

func (v NullableMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFAPushCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


