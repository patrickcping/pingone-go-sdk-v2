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

// checks if the LicenseVerify type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LicenseVerify{}

// LicenseVerify struct for LicenseVerify
type LicenseVerify struct {
	// A read-only boolean that specifies whether to enable the PingOne Verify push notifications feature.
	AllowPushNotifications *bool `json:"allowPushNotifications,omitempty"`
	// A read-only boolean that specifies whether to enable the PingOne Verify document matching feature.
	AllowDocumentMatch *bool `json:"allowDocumentMatch,omitempty"`
	// A read-only boolean that specifies whether to enable the PingOne Verify face matching feature.
	AllowFaceMatch *bool `json:"allowFaceMatch,omitempty"`
	// A read-only boolean that specifies whether to enable the PingOne Verify manual ID inspection feature.
	AllowManualIdInspection *bool `json:"allowManualIdInspection,omitempty"`
}

// NewLicenseVerify instantiates a new LicenseVerify object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLicenseVerify() *LicenseVerify {
	this := LicenseVerify{}
	return &this
}

// NewLicenseVerifyWithDefaults instantiates a new LicenseVerify object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLicenseVerifyWithDefaults() *LicenseVerify {
	this := LicenseVerify{}
	return &this
}

// GetAllowPushNotifications returns the AllowPushNotifications field value if set, zero value otherwise.
func (o *LicenseVerify) GetAllowPushNotifications() bool {
	if o == nil || IsNil(o.AllowPushNotifications) {
		var ret bool
		return ret
	}
	return *o.AllowPushNotifications
}

// GetAllowPushNotificationsOk returns a tuple with the AllowPushNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseVerify) GetAllowPushNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPushNotifications) {
		return nil, false
	}
	return o.AllowPushNotifications, true
}

// HasAllowPushNotifications returns a boolean if a field has been set.
func (o *LicenseVerify) HasAllowPushNotifications() bool {
	if o != nil && !IsNil(o.AllowPushNotifications) {
		return true
	}

	return false
}

// SetAllowPushNotifications gets a reference to the given bool and assigns it to the AllowPushNotifications field.
func (o *LicenseVerify) SetAllowPushNotifications(v bool) {
	o.AllowPushNotifications = &v
}

// GetAllowDocumentMatch returns the AllowDocumentMatch field value if set, zero value otherwise.
func (o *LicenseVerify) GetAllowDocumentMatch() bool {
	if o == nil || IsNil(o.AllowDocumentMatch) {
		var ret bool
		return ret
	}
	return *o.AllowDocumentMatch
}

// GetAllowDocumentMatchOk returns a tuple with the AllowDocumentMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseVerify) GetAllowDocumentMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowDocumentMatch) {
		return nil, false
	}
	return o.AllowDocumentMatch, true
}

// HasAllowDocumentMatch returns a boolean if a field has been set.
func (o *LicenseVerify) HasAllowDocumentMatch() bool {
	if o != nil && !IsNil(o.AllowDocumentMatch) {
		return true
	}

	return false
}

// SetAllowDocumentMatch gets a reference to the given bool and assigns it to the AllowDocumentMatch field.
func (o *LicenseVerify) SetAllowDocumentMatch(v bool) {
	o.AllowDocumentMatch = &v
}

// GetAllowFaceMatch returns the AllowFaceMatch field value if set, zero value otherwise.
func (o *LicenseVerify) GetAllowFaceMatch() bool {
	if o == nil || IsNil(o.AllowFaceMatch) {
		var ret bool
		return ret
	}
	return *o.AllowFaceMatch
}

// GetAllowFaceMatchOk returns a tuple with the AllowFaceMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseVerify) GetAllowFaceMatchOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowFaceMatch) {
		return nil, false
	}
	return o.AllowFaceMatch, true
}

// HasAllowFaceMatch returns a boolean if a field has been set.
func (o *LicenseVerify) HasAllowFaceMatch() bool {
	if o != nil && !IsNil(o.AllowFaceMatch) {
		return true
	}

	return false
}

// SetAllowFaceMatch gets a reference to the given bool and assigns it to the AllowFaceMatch field.
func (o *LicenseVerify) SetAllowFaceMatch(v bool) {
	o.AllowFaceMatch = &v
}

// GetAllowManualIdInspection returns the AllowManualIdInspection field value if set, zero value otherwise.
func (o *LicenseVerify) GetAllowManualIdInspection() bool {
	if o == nil || IsNil(o.AllowManualIdInspection) {
		var ret bool
		return ret
	}
	return *o.AllowManualIdInspection
}

// GetAllowManualIdInspectionOk returns a tuple with the AllowManualIdInspection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LicenseVerify) GetAllowManualIdInspectionOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowManualIdInspection) {
		return nil, false
	}
	return o.AllowManualIdInspection, true
}

// HasAllowManualIdInspection returns a boolean if a field has been set.
func (o *LicenseVerify) HasAllowManualIdInspection() bool {
	if o != nil && !IsNil(o.AllowManualIdInspection) {
		return true
	}

	return false
}

// SetAllowManualIdInspection gets a reference to the given bool and assigns it to the AllowManualIdInspection field.
func (o *LicenseVerify) SetAllowManualIdInspection(v bool) {
	o.AllowManualIdInspection = &v
}

func (o LicenseVerify) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LicenseVerify) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowPushNotifications) {
		toSerialize["allowPushNotifications"] = o.AllowPushNotifications
	}
	if !IsNil(o.AllowDocumentMatch) {
		toSerialize["allowDocumentMatch"] = o.AllowDocumentMatch
	}
	if !IsNil(o.AllowFaceMatch) {
		toSerialize["allowFaceMatch"] = o.AllowFaceMatch
	}
	if !IsNil(o.AllowManualIdInspection) {
		toSerialize["allowManualIdInspection"] = o.AllowManualIdInspection
	}
	return toSerialize, nil
}

type NullableLicenseVerify struct {
	value *LicenseVerify
	isSet bool
}

func (v NullableLicenseVerify) Get() *LicenseVerify {
	return v.value
}

func (v *NullableLicenseVerify) Set(val *LicenseVerify) {
	v.value = val
	v.isSet = true
}

func (v NullableLicenseVerify) IsSet() bool {
	return v.isSet
}

func (v *NullableLicenseVerify) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLicenseVerify(val *LicenseVerify) *NullableLicenseVerify {
	return &NullableLicenseVerify{value: val, isSet: true}
}

func (v NullableLicenseVerify) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLicenseVerify) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
