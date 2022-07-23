/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// ApplicationOIDCAllOfMobileIntegrityDetection struct for ApplicationOIDCAllOfMobileIntegrityDetection
type ApplicationOIDCAllOfMobileIntegrityDetection struct {
	Mode *EnumEnabledStatus `json:"mode,omitempty"`
	CacheDuration *ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration `json:"cacheDuration,omitempty"`
}

// NewApplicationOIDCAllOfMobileIntegrityDetection instantiates a new ApplicationOIDCAllOfMobileIntegrityDetection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationOIDCAllOfMobileIntegrityDetection() *ApplicationOIDCAllOfMobileIntegrityDetection {
	this := ApplicationOIDCAllOfMobileIntegrityDetection{}
	return &this
}

// NewApplicationOIDCAllOfMobileIntegrityDetectionWithDefaults instantiates a new ApplicationOIDCAllOfMobileIntegrityDetection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationOIDCAllOfMobileIntegrityDetectionWithDefaults() *ApplicationOIDCAllOfMobileIntegrityDetection {
	this := ApplicationOIDCAllOfMobileIntegrityDetection{}
	return &this
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetMode() EnumEnabledStatus {
	if o == nil || o.Mode == nil {
		var ret EnumEnabledStatus
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetModeOk() (*EnumEnabledStatus, bool) {
	if o == nil || o.Mode == nil {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobileIntegrityDetection) HasMode() bool {
	if o != nil && o.Mode != nil {
		return true
	}

	return false
}

// SetMode gets a reference to the given EnumEnabledStatus and assigns it to the Mode field.
func (o *ApplicationOIDCAllOfMobileIntegrityDetection) SetMode(v EnumEnabledStatus) {
	o.Mode = &v
}

// GetCacheDuration returns the CacheDuration field value if set, zero value otherwise.
func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetCacheDuration() ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration {
	if o == nil || o.CacheDuration == nil {
		var ret ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration
		return ret
	}
	return *o.CacheDuration
}

// GetCacheDurationOk returns a tuple with the CacheDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationOIDCAllOfMobileIntegrityDetection) GetCacheDurationOk() (*ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration, bool) {
	if o == nil || o.CacheDuration == nil {
		return nil, false
	}
	return o.CacheDuration, true
}

// HasCacheDuration returns a boolean if a field has been set.
func (o *ApplicationOIDCAllOfMobileIntegrityDetection) HasCacheDuration() bool {
	if o != nil && o.CacheDuration != nil {
		return true
	}

	return false
}

// SetCacheDuration gets a reference to the given ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration and assigns it to the CacheDuration field.
func (o *ApplicationOIDCAllOfMobileIntegrityDetection) SetCacheDuration(v ApplicationOIDCAllOfMobileIntegrityDetectionCacheDuration) {
	o.CacheDuration = &v
}

func (o ApplicationOIDCAllOfMobileIntegrityDetection) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Mode != nil {
		toSerialize["mode"] = o.Mode
	}
	if o.CacheDuration != nil {
		toSerialize["cacheDuration"] = o.CacheDuration
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationOIDCAllOfMobileIntegrityDetection struct {
	value *ApplicationOIDCAllOfMobileIntegrityDetection
	isSet bool
}

func (v NullableApplicationOIDCAllOfMobileIntegrityDetection) Get() *ApplicationOIDCAllOfMobileIntegrityDetection {
	return v.value
}

func (v *NullableApplicationOIDCAllOfMobileIntegrityDetection) Set(val *ApplicationOIDCAllOfMobileIntegrityDetection) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationOIDCAllOfMobileIntegrityDetection) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationOIDCAllOfMobileIntegrityDetection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationOIDCAllOfMobileIntegrityDetection(val *ApplicationOIDCAllOfMobileIntegrityDetection) *NullableApplicationOIDCAllOfMobileIntegrityDetection {
	return &NullableApplicationOIDCAllOfMobileIntegrityDetection{value: val, isSet: true}
}

func (v NullableApplicationOIDCAllOfMobileIntegrityDetection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationOIDCAllOfMobileIntegrityDetection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


