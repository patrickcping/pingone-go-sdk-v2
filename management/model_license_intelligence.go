/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the LicenseIntelligence type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseIntelligence{}

// LicenseIntelligence struct for LicenseIntelligence
type LicenseIntelligence struct {
	// A read-only boolean that specifies whether to use the intelligence geo-velocity feature. For `TRIAL` (unpaid) licenses, the default value is true. For `ADMIN`, `GLOBAL`, `RISK`, and `MFARISK`, the default value is true.
	AllowGeoVelocity *bool `json:"allowGeoVelocity,omitempty"`
	// A read-only boolean that specifies whether to use the intelligence anonymous network detection feature. For `TRIAL` (unpaid) licenses, the default value is true. For `ADMIN`, `GLOBAL`, `RISK`, and `MFARISK`, the default value is true.
	AllowAnonymousNetworkDetection *bool `json:"allowAnonymousNetworkDetection,omitempty"`
	// A read-only boolean that specifies whether to use the intelligence IP reputation feature. For `TRIAL` (unpaid) licenses, the default value is true. For `ADMIN`, `GLOBAL`, `RISK`, and `MFARISK`, the default value is true.
	AllowReputation *bool `json:"allowReputation,omitempty"`
	// A read-only boolean that specifies whether the customer has opted in to allow user and event behavior analytics (UEBA) data collection.
	AllowDataConsent *bool `json:"allowDataConsent,omitempty"`
	// A read-only boolean that specifies whether your license permits you to configure risk features such as sign-on policies that include rules to detect anomalous changes to your locations (such as impossible travel). This capability is supported for TRIAL, RISK, and MFARISK license packages. Note, The sharing of user data to enable our machine-learning engine, which is integral to PingOne Risk, is captured in the license property license.intelligence.allowDataConsent, but it is not set to true by default in any license package. This license capability always requires active consent by the customer before it can be enabled, and if consent is given, then it allows the full scope of intelligence features included in PingOne Risk (and PingOne Risk plus MFA).
	AllowRisk               *bool `json:"allowRisk,omitempty"`
	AllowAdvancedPredictors *bool `json:"allowAdvancedPredictors,omitempty"`
}

// NewLicenseIntelligence instantiates a new LicenseIntelligence object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseIntelligence() *LicenseIntelligence {
	this := LicenseIntelligence{}
	return &this
}

// NewLicenseIntelligenceWithDefaults instantiates a new LicenseIntelligence object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseIntelligenceWithDefaults() *LicenseIntelligence {
	this := LicenseIntelligence{}
	return &this
}

// GetAllowGeoVelocity returns the AllowGeoVelocity field value if set, zero value otherwise.
func (o *LicenseIntelligence) GetAllowGeoVelocity() bool {
	if o == nil || IsNil(o.AllowGeoVelocity) {
		var ret bool
		return ret
	}
	return *o.AllowGeoVelocity
}

// GetAllowGeoVelocityOk returns a tuple with the AllowGeoVelocity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIntelligence) GetAllowGeoVelocityOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowGeoVelocity) {
		return nil, false
	}
	return o.AllowGeoVelocity, true
}

// HasAllowGeoVelocity returns a boolean if a field has been set.
func (o *LicenseIntelligence) HasAllowGeoVelocity() bool {
	if o != nil && !IsNil(o.AllowGeoVelocity) {
		return true
	}

	return false
}

// SetAllowGeoVelocity gets a reference to the given bool and assigns it to the AllowGeoVelocity field.
func (o *LicenseIntelligence) SetAllowGeoVelocity(v bool) {
	o.AllowGeoVelocity = &v
}

// GetAllowAnonymousNetworkDetection returns the AllowAnonymousNetworkDetection field value if set, zero value otherwise.
func (o *LicenseIntelligence) GetAllowAnonymousNetworkDetection() bool {
	if o == nil || IsNil(o.AllowAnonymousNetworkDetection) {
		var ret bool
		return ret
	}
	return *o.AllowAnonymousNetworkDetection
}

// GetAllowAnonymousNetworkDetectionOk returns a tuple with the AllowAnonymousNetworkDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIntelligence) GetAllowAnonymousNetworkDetectionOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAnonymousNetworkDetection) {
		return nil, false
	}
	return o.AllowAnonymousNetworkDetection, true
}

// HasAllowAnonymousNetworkDetection returns a boolean if a field has been set.
func (o *LicenseIntelligence) HasAllowAnonymousNetworkDetection() bool {
	if o != nil && !IsNil(o.AllowAnonymousNetworkDetection) {
		return true
	}

	return false
}

// SetAllowAnonymousNetworkDetection gets a reference to the given bool and assigns it to the AllowAnonymousNetworkDetection field.
func (o *LicenseIntelligence) SetAllowAnonymousNetworkDetection(v bool) {
	o.AllowAnonymousNetworkDetection = &v
}

// GetAllowReputation returns the AllowReputation field value if set, zero value otherwise.
func (o *LicenseIntelligence) GetAllowReputation() bool {
	if o == nil || IsNil(o.AllowReputation) {
		var ret bool
		return ret
	}
	return *o.AllowReputation
}

// GetAllowReputationOk returns a tuple with the AllowReputation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIntelligence) GetAllowReputationOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowReputation) {
		return nil, false
	}
	return o.AllowReputation, true
}

// HasAllowReputation returns a boolean if a field has been set.
func (o *LicenseIntelligence) HasAllowReputation() bool {
	if o != nil && !IsNil(o.AllowReputation) {
		return true
	}

	return false
}

// SetAllowReputation gets a reference to the given bool and assigns it to the AllowReputation field.
func (o *LicenseIntelligence) SetAllowReputation(v bool) {
	o.AllowReputation = &v
}

// GetAllowDataConsent returns the AllowDataConsent field value if set, zero value otherwise.
func (o *LicenseIntelligence) GetAllowDataConsent() bool {
	if o == nil || IsNil(o.AllowDataConsent) {
		var ret bool
		return ret
	}
	return *o.AllowDataConsent
}

// GetAllowDataConsentOk returns a tuple with the AllowDataConsent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIntelligence) GetAllowDataConsentOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowDataConsent) {
		return nil, false
	}
	return o.AllowDataConsent, true
}

// HasAllowDataConsent returns a boolean if a field has been set.
func (o *LicenseIntelligence) HasAllowDataConsent() bool {
	if o != nil && !IsNil(o.AllowDataConsent) {
		return true
	}

	return false
}

// SetAllowDataConsent gets a reference to the given bool and assigns it to the AllowDataConsent field.
func (o *LicenseIntelligence) SetAllowDataConsent(v bool) {
	o.AllowDataConsent = &v
}

// GetAllowRisk returns the AllowRisk field value if set, zero value otherwise.
func (o *LicenseIntelligence) GetAllowRisk() bool {
	if o == nil || IsNil(o.AllowRisk) {
		var ret bool
		return ret
	}
	return *o.AllowRisk
}

// GetAllowRiskOk returns a tuple with the AllowRisk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIntelligence) GetAllowRiskOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowRisk) {
		return nil, false
	}
	return o.AllowRisk, true
}

// HasAllowRisk returns a boolean if a field has been set.
func (o *LicenseIntelligence) HasAllowRisk() bool {
	if o != nil && !IsNil(o.AllowRisk) {
		return true
	}

	return false
}

// SetAllowRisk gets a reference to the given bool and assigns it to the AllowRisk field.
func (o *LicenseIntelligence) SetAllowRisk(v bool) {
	o.AllowRisk = &v
}

// GetAllowAdvancedPredictors returns the AllowAdvancedPredictors field value if set, zero value otherwise.
func (o *LicenseIntelligence) GetAllowAdvancedPredictors() bool {
	if o == nil || IsNil(o.AllowAdvancedPredictors) {
		var ret bool
		return ret
	}
	return *o.AllowAdvancedPredictors
}

// GetAllowAdvancedPredictorsOk returns a tuple with the AllowAdvancedPredictors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseIntelligence) GetAllowAdvancedPredictorsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowAdvancedPredictors) {
		return nil, false
	}
	return o.AllowAdvancedPredictors, true
}

// HasAllowAdvancedPredictors returns a boolean if a field has been set.
func (o *LicenseIntelligence) HasAllowAdvancedPredictors() bool {
	if o != nil && !IsNil(o.AllowAdvancedPredictors) {
		return true
	}

	return false
}

// SetAllowAdvancedPredictors gets a reference to the given bool and assigns it to the AllowAdvancedPredictors field.
func (o *LicenseIntelligence) SetAllowAdvancedPredictors(v bool) {
	o.AllowAdvancedPredictors = &v
}

func (o LicenseIntelligence) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseIntelligence) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowGeoVelocity) {
		toSerialize["allowGeoVelocity"] = o.AllowGeoVelocity
	}
	if !IsNil(o.AllowAnonymousNetworkDetection) {
		toSerialize["allowAnonymousNetworkDetection"] = o.AllowAnonymousNetworkDetection
	}
	if !IsNil(o.AllowReputation) {
		toSerialize["allowReputation"] = o.AllowReputation
	}
	if !IsNil(o.AllowDataConsent) {
		toSerialize["allowDataConsent"] = o.AllowDataConsent
	}
	if !IsNil(o.AllowRisk) {
		toSerialize["allowRisk"] = o.AllowRisk
	}
	if !IsNil(o.AllowAdvancedPredictors) {
		toSerialize["allowAdvancedPredictors"] = o.AllowAdvancedPredictors
	}
	return toSerialize, nil
}

type NullableLicenseIntelligence struct {
	value *LicenseIntelligence
	isSet bool
}

func (v NullableLicenseIntelligence) Get() *LicenseIntelligence {
	return v.value
}

func (v *NullableLicenseIntelligence) Set(val *LicenseIntelligence) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseIntelligence) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseIntelligence) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseIntelligence(val *LicenseIntelligence) *NullableLicenseIntelligence {
	return &NullableLicenseIntelligence{value: val, isSet: true}
}

func (v NullableLicenseIntelligence) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseIntelligence) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
