/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
)

// checks if the FIDO2PolicyMdsAuthenticatorsRequirements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FIDO2PolicyMdsAuthenticatorsRequirements{}

// FIDO2PolicyMdsAuthenticatorsRequirements Used to specify whether attestation is requested from the authenticator, and whether this information is used to restrict authenticator usage.
type FIDO2PolicyMdsAuthenticatorsRequirements struct {
	// If you set `mdsAuthenticatorsRequirements.option` to `SPECIFIC`, use this array to specify the authenticators that you want to allow.
	AllowedAuthenticators []FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner `json:"allowedAuthenticators,omitempty"`
	// Set to true if you want the device characteristics related to attestation to be checked again at each authentication attempt and not just once during registration. Set to false to have them checked only at registration.
	EnforceDuringAuthentication bool `json:"enforceDuringAuthentication"`
	Option EnumFIDO2PolicyMDSAuthenticatorOption `json:"option"`
}

// NewFIDO2PolicyMdsAuthenticatorsRequirements instantiates a new FIDO2PolicyMdsAuthenticatorsRequirements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFIDO2PolicyMdsAuthenticatorsRequirements(enforceDuringAuthentication bool, option EnumFIDO2PolicyMDSAuthenticatorOption) *FIDO2PolicyMdsAuthenticatorsRequirements {
	this := FIDO2PolicyMdsAuthenticatorsRequirements{}
	this.EnforceDuringAuthentication = enforceDuringAuthentication
	this.Option = option
	return &this
}

// NewFIDO2PolicyMdsAuthenticatorsRequirementsWithDefaults instantiates a new FIDO2PolicyMdsAuthenticatorsRequirements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFIDO2PolicyMdsAuthenticatorsRequirementsWithDefaults() *FIDO2PolicyMdsAuthenticatorsRequirements {
	this := FIDO2PolicyMdsAuthenticatorsRequirements{}
	return &this
}

// GetAllowedAuthenticators returns the AllowedAuthenticators field value if set, zero value otherwise.
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetAllowedAuthenticators() []FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner {
	if o == nil || IsNil(o.AllowedAuthenticators) {
		var ret []FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner
		return ret
	}
	return o.AllowedAuthenticators
}

// GetAllowedAuthenticatorsOk returns a tuple with the AllowedAuthenticators field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetAllowedAuthenticatorsOk() ([]FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner, bool) {
	if o == nil || IsNil(o.AllowedAuthenticators) {
		return nil, false
	}
	return o.AllowedAuthenticators, true
}

// HasAllowedAuthenticators returns a boolean if a field has been set.
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) HasAllowedAuthenticators() bool {
	if o != nil && !IsNil(o.AllowedAuthenticators) {
		return true
	}

	return false
}

// SetAllowedAuthenticators gets a reference to the given []FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner and assigns it to the AllowedAuthenticators field.
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) SetAllowedAuthenticators(v []FIDO2PolicyMdsAuthenticatorsRequirementsAllowedAuthenticatorsInner) {
	o.AllowedAuthenticators = v
}

// GetEnforceDuringAuthentication returns the EnforceDuringAuthentication field value
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetEnforceDuringAuthentication() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.EnforceDuringAuthentication
}

// GetEnforceDuringAuthenticationOk returns a tuple with the EnforceDuringAuthentication field value
// and a boolean to check if the value has been set.
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetEnforceDuringAuthenticationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnforceDuringAuthentication, true
}

// SetEnforceDuringAuthentication sets field value
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) SetEnforceDuringAuthentication(v bool) {
	o.EnforceDuringAuthentication = v
}

// GetOption returns the Option field value
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetOption() EnumFIDO2PolicyMDSAuthenticatorOption {
	if o == nil {
		var ret EnumFIDO2PolicyMDSAuthenticatorOption
		return ret
	}

	return o.Option
}

// GetOptionOk returns a tuple with the Option field value
// and a boolean to check if the value has been set.
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) GetOptionOk() (*EnumFIDO2PolicyMDSAuthenticatorOption, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Option, true
}

// SetOption sets field value
func (o *FIDO2PolicyMdsAuthenticatorsRequirements) SetOption(v EnumFIDO2PolicyMDSAuthenticatorOption) {
	o.Option = v
}

func (o FIDO2PolicyMdsAuthenticatorsRequirements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FIDO2PolicyMdsAuthenticatorsRequirements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedAuthenticators) {
		toSerialize["allowedAuthenticators"] = o.AllowedAuthenticators
	}
	toSerialize["enforceDuringAuthentication"] = o.EnforceDuringAuthentication
	toSerialize["option"] = o.Option
	return toSerialize, nil
}

type NullableFIDO2PolicyMdsAuthenticatorsRequirements struct {
	value *FIDO2PolicyMdsAuthenticatorsRequirements
	isSet bool
}

func (v NullableFIDO2PolicyMdsAuthenticatorsRequirements) Get() *FIDO2PolicyMdsAuthenticatorsRequirements {
	return v.value
}

func (v *NullableFIDO2PolicyMdsAuthenticatorsRequirements) Set(val *FIDO2PolicyMdsAuthenticatorsRequirements) {
	v.value = val
	v.isSet = true
}

func (v NullableFIDO2PolicyMdsAuthenticatorsRequirements) IsSet() bool {
	return v.isSet
}

func (v *NullableFIDO2PolicyMdsAuthenticatorsRequirements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFIDO2PolicyMdsAuthenticatorsRequirements(val *FIDO2PolicyMdsAuthenticatorsRequirements) *NullableFIDO2PolicyMdsAuthenticatorsRequirements {
	return &NullableFIDO2PolicyMdsAuthenticatorsRequirements{value: val, isSet: true}
}

func (v NullableFIDO2PolicyMdsAuthenticatorsRequirements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFIDO2PolicyMdsAuthenticatorsRequirements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


