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

// CreateMFAPushCredentialRequest - struct for CreateMFAPushCredentialRequest
type CreateMFAPushCredentialRequest struct {
	MFAPushCredentialAPNS      *MFAPushCredentialAPNS
	MFAPushCredentialFCM       *MFAPushCredentialFCM
	MFAPushCredentialFCMHTTPV1 *MFAPushCredentialFCMHTTPV1
	MFAPushCredentialHMS       *MFAPushCredentialHMS
}

// MFAPushCredentialAPNSAsCreateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialAPNS wrapped in CreateMFAPushCredentialRequest
func MFAPushCredentialAPNSAsCreateMFAPushCredentialRequest(v *MFAPushCredentialAPNS) CreateMFAPushCredentialRequest {
	return CreateMFAPushCredentialRequest{
		MFAPushCredentialAPNS: v,
	}
}

// MFAPushCredentialFCMAsCreateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialFCM wrapped in CreateMFAPushCredentialRequest
func MFAPushCredentialFCMAsCreateMFAPushCredentialRequest(v *MFAPushCredentialFCM) CreateMFAPushCredentialRequest {
	return CreateMFAPushCredentialRequest{
		MFAPushCredentialFCM: v,
	}
}

// MFAPushCredentialFCMHTTPV1AsCreateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialFCMHTTPV1 wrapped in CreateMFAPushCredentialRequest
func MFAPushCredentialFCMHTTPV1AsCreateMFAPushCredentialRequest(v *MFAPushCredentialFCMHTTPV1) CreateMFAPushCredentialRequest {
	return CreateMFAPushCredentialRequest{
		MFAPushCredentialFCMHTTPV1: v,
	}
}

// MFAPushCredentialHMSAsCreateMFAPushCredentialRequest is a convenience function that returns MFAPushCredentialHMS wrapped in CreateMFAPushCredentialRequest
func MFAPushCredentialHMSAsCreateMFAPushCredentialRequest(v *MFAPushCredentialHMS) CreateMFAPushCredentialRequest {
	return CreateMFAPushCredentialRequest{
		MFAPushCredentialHMS: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *CreateMFAPushCredentialRequest) UnmarshalJSON(data []byte) error {

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
		return fmt.Errorf("Data failed to match schemas in anyOf(CreateMFAPushCredential201Response)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src CreateMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
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
func (obj *CreateMFAPushCredentialRequest) GetActualInstance() interface{} {
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

type NullableCreateMFAPushCredentialRequest struct {
	value *CreateMFAPushCredentialRequest
	isSet bool
}

func (v NullableCreateMFAPushCredentialRequest) Get() *CreateMFAPushCredentialRequest {
	return v.value
}

func (v *NullableCreateMFAPushCredentialRequest) Set(val *CreateMFAPushCredentialRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMFAPushCredentialRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMFAPushCredentialRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMFAPushCredentialRequest(val *CreateMFAPushCredentialRequest) *NullableCreateMFAPushCredentialRequest {
	return &NullableCreateMFAPushCredentialRequest{value: val, isSet: true}
}

func (v NullableCreateMFAPushCredentialRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMFAPushCredentialRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
