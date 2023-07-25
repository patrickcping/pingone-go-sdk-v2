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

// checks if the DeviceAuthenticationPolicyMobileApplicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicyMobileApplicationsInner{}

// DeviceAuthenticationPolicyMobileApplicationsInner struct for DeviceAuthenticationPolicyMobileApplicationsInner
type DeviceAuthenticationPolicyMobileApplicationsInner struct {
	// The application's ID.
	Id string `json:"id"`
	Push *DeviceAuthenticationPolicyMobileApplicationsInnerPush `json:"push,omitempty"`
	PushTimeout *DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout `json:"pushTimeout,omitempty"`
	PairingKeyLifetime *DeviceAuthenticationPolicyMobileApplicationsInnerPairingKeyLifetime `json:"pairingKeyLifetime,omitempty"`
	PushLimit *DeviceAuthenticationPolicyMobileApplicationsInnerPushLimit `json:"pushLimit,omitempty"`
	Otp *DeviceAuthenticationPolicyMobileApplicationsInnerOtp `json:"otp,omitempty"`
	DeviceAuthorization *DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization `json:"deviceAuthorization,omitempty"`
	AutoEnrollment *DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment `json:"autoEnrollment,omitempty"`
	// You can set `pairingDisabled` to true to prevent users from pairing new devices with the relevant method. You can use this option if you want to phase out an existing authentication method but want to allow users to continue using the method for authentication for existing devices.
	PairingDisabled *bool `json:"pairingDisabled,omitempty"`
	IntegrityDetection *EnumMFADevicePolicyMobileIntegrityDetection `json:"integrityDetection,omitempty"`
}

// NewDeviceAuthenticationPolicyMobileApplicationsInner instantiates a new DeviceAuthenticationPolicyMobileApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicyMobileApplicationsInner(id string) *DeviceAuthenticationPolicyMobileApplicationsInner {
	this := DeviceAuthenticationPolicyMobileApplicationsInner{}
	this.Id = id
	return &this
}

// NewDeviceAuthenticationPolicyMobileApplicationsInnerWithDefaults instantiates a new DeviceAuthenticationPolicyMobileApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyMobileApplicationsInnerWithDefaults() *DeviceAuthenticationPolicyMobileApplicationsInner {
	this := DeviceAuthenticationPolicyMobileApplicationsInner{}
	return &this
}

// GetId returns the Id field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetId(v string) {
	o.Id = v
}

// GetPush returns the Push field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPush() DeviceAuthenticationPolicyMobileApplicationsInnerPush {
	if o == nil || IsNil(o.Push) {
		var ret DeviceAuthenticationPolicyMobileApplicationsInnerPush
		return ret
	}
	return *o.Push
}

// GetPushOk returns a tuple with the Push field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPushOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerPush, bool) {
	if o == nil || IsNil(o.Push) {
		return nil, false
	}
	return o.Push, true
}

// HasPush returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasPush() bool {
	if o != nil && !IsNil(o.Push) {
		return true
	}

	return false
}

// SetPush gets a reference to the given DeviceAuthenticationPolicyMobileApplicationsInnerPush and assigns it to the Push field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetPush(v DeviceAuthenticationPolicyMobileApplicationsInnerPush) {
	o.Push = &v
}

// GetPushTimeout returns the PushTimeout field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPushTimeout() DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout {
	if o == nil || IsNil(o.PushTimeout) {
		var ret DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout
		return ret
	}
	return *o.PushTimeout
}

// GetPushTimeoutOk returns a tuple with the PushTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPushTimeoutOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout, bool) {
	if o == nil || IsNil(o.PushTimeout) {
		return nil, false
	}
	return o.PushTimeout, true
}

// HasPushTimeout returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasPushTimeout() bool {
	if o != nil && !IsNil(o.PushTimeout) {
		return true
	}

	return false
}

// SetPushTimeout gets a reference to the given DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout and assigns it to the PushTimeout field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetPushTimeout(v DeviceAuthenticationPolicyMobileApplicationsInnerPushTimeout) {
	o.PushTimeout = &v
}

// GetPairingKeyLifetime returns the PairingKeyLifetime field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPairingKeyLifetime() DeviceAuthenticationPolicyMobileApplicationsInnerPairingKeyLifetime {
	if o == nil || IsNil(o.PairingKeyLifetime) {
		var ret DeviceAuthenticationPolicyMobileApplicationsInnerPairingKeyLifetime
		return ret
	}
	return *o.PairingKeyLifetime
}

// GetPairingKeyLifetimeOk returns a tuple with the PairingKeyLifetime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPairingKeyLifetimeOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerPairingKeyLifetime, bool) {
	if o == nil || IsNil(o.PairingKeyLifetime) {
		return nil, false
	}
	return o.PairingKeyLifetime, true
}

// HasPairingKeyLifetime returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasPairingKeyLifetime() bool {
	if o != nil && !IsNil(o.PairingKeyLifetime) {
		return true
	}

	return false
}

// SetPairingKeyLifetime gets a reference to the given DeviceAuthenticationPolicyMobileApplicationsInnerPairingKeyLifetime and assigns it to the PairingKeyLifetime field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetPairingKeyLifetime(v DeviceAuthenticationPolicyMobileApplicationsInnerPairingKeyLifetime) {
	o.PairingKeyLifetime = &v
}

// GetPushLimit returns the PushLimit field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPushLimit() DeviceAuthenticationPolicyMobileApplicationsInnerPushLimit {
	if o == nil || IsNil(o.PushLimit) {
		var ret DeviceAuthenticationPolicyMobileApplicationsInnerPushLimit
		return ret
	}
	return *o.PushLimit
}

// GetPushLimitOk returns a tuple with the PushLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPushLimitOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerPushLimit, bool) {
	if o == nil || IsNil(o.PushLimit) {
		return nil, false
	}
	return o.PushLimit, true
}

// HasPushLimit returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasPushLimit() bool {
	if o != nil && !IsNil(o.PushLimit) {
		return true
	}

	return false
}

// SetPushLimit gets a reference to the given DeviceAuthenticationPolicyMobileApplicationsInnerPushLimit and assigns it to the PushLimit field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetPushLimit(v DeviceAuthenticationPolicyMobileApplicationsInnerPushLimit) {
	o.PushLimit = &v
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetOtp() DeviceAuthenticationPolicyMobileApplicationsInnerOtp {
	if o == nil || IsNil(o.Otp) {
		var ret DeviceAuthenticationPolicyMobileApplicationsInnerOtp
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetOtpOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerOtp, bool) {
	if o == nil || IsNil(o.Otp) {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasOtp() bool {
	if o != nil && !IsNil(o.Otp) {
		return true
	}

	return false
}

// SetOtp gets a reference to the given DeviceAuthenticationPolicyMobileApplicationsInnerOtp and assigns it to the Otp field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetOtp(v DeviceAuthenticationPolicyMobileApplicationsInnerOtp) {
	o.Otp = &v
}

// GetDeviceAuthorization returns the DeviceAuthorization field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetDeviceAuthorization() DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization {
	if o == nil || IsNil(o.DeviceAuthorization) {
		var ret DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization
		return ret
	}
	return *o.DeviceAuthorization
}

// GetDeviceAuthorizationOk returns a tuple with the DeviceAuthorization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetDeviceAuthorizationOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization, bool) {
	if o == nil || IsNil(o.DeviceAuthorization) {
		return nil, false
	}
	return o.DeviceAuthorization, true
}

// HasDeviceAuthorization returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasDeviceAuthorization() bool {
	if o != nil && !IsNil(o.DeviceAuthorization) {
		return true
	}

	return false
}

// SetDeviceAuthorization gets a reference to the given DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization and assigns it to the DeviceAuthorization field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetDeviceAuthorization(v DeviceAuthenticationPolicyMobileApplicationsInnerDeviceAuthorization) {
	o.DeviceAuthorization = &v
}

// GetAutoEnrollment returns the AutoEnrollment field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetAutoEnrollment() DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment {
	if o == nil || IsNil(o.AutoEnrollment) {
		var ret DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment
		return ret
	}
	return *o.AutoEnrollment
}

// GetAutoEnrollmentOk returns a tuple with the AutoEnrollment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetAutoEnrollmentOk() (*DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment, bool) {
	if o == nil || IsNil(o.AutoEnrollment) {
		return nil, false
	}
	return o.AutoEnrollment, true
}

// HasAutoEnrollment returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasAutoEnrollment() bool {
	if o != nil && !IsNil(o.AutoEnrollment) {
		return true
	}

	return false
}

// SetAutoEnrollment gets a reference to the given DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment and assigns it to the AutoEnrollment field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetAutoEnrollment(v DeviceAuthenticationPolicyMobileApplicationsInnerAutoEnrollment) {
	o.AutoEnrollment = &v
}

// GetPairingDisabled returns the PairingDisabled field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPairingDisabled() bool {
	if o == nil || IsNil(o.PairingDisabled) {
		var ret bool
		return ret
	}
	return *o.PairingDisabled
}

// GetPairingDisabledOk returns a tuple with the PairingDisabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetPairingDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PairingDisabled) {
		return nil, false
	}
	return o.PairingDisabled, true
}

// HasPairingDisabled returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasPairingDisabled() bool {
	if o != nil && !IsNil(o.PairingDisabled) {
		return true
	}

	return false
}

// SetPairingDisabled gets a reference to the given bool and assigns it to the PairingDisabled field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetPairingDisabled(v bool) {
	o.PairingDisabled = &v
}

// GetIntegrityDetection returns the IntegrityDetection field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetIntegrityDetection() EnumMFADevicePolicyMobileIntegrityDetection {
	if o == nil || IsNil(o.IntegrityDetection) {
		var ret EnumMFADevicePolicyMobileIntegrityDetection
		return ret
	}
	return *o.IntegrityDetection
}

// GetIntegrityDetectionOk returns a tuple with the IntegrityDetection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) GetIntegrityDetectionOk() (*EnumMFADevicePolicyMobileIntegrityDetection, bool) {
	if o == nil || IsNil(o.IntegrityDetection) {
		return nil, false
	}
	return o.IntegrityDetection, true
}

// HasIntegrityDetection returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) HasIntegrityDetection() bool {
	if o != nil && !IsNil(o.IntegrityDetection) {
		return true
	}

	return false
}

// SetIntegrityDetection gets a reference to the given EnumMFADevicePolicyMobileIntegrityDetection and assigns it to the IntegrityDetection field.
func (o *DeviceAuthenticationPolicyMobileApplicationsInner) SetIntegrityDetection(v EnumMFADevicePolicyMobileIntegrityDetection) {
	o.IntegrityDetection = &v
}

func (o DeviceAuthenticationPolicyMobileApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicyMobileApplicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Push) {
		toSerialize["push"] = o.Push
	}
	if !IsNil(o.PushTimeout) {
		toSerialize["pushTimeout"] = o.PushTimeout
	}
	if !IsNil(o.PairingKeyLifetime) {
		toSerialize["pairingKeyLifetime"] = o.PairingKeyLifetime
	}
	if !IsNil(o.PushLimit) {
		toSerialize["pushLimit"] = o.PushLimit
	}
	if !IsNil(o.Otp) {
		toSerialize["otp"] = o.Otp
	}
	if !IsNil(o.DeviceAuthorization) {
		toSerialize["deviceAuthorization"] = o.DeviceAuthorization
	}
	if !IsNil(o.AutoEnrollment) {
		toSerialize["autoEnrollment"] = o.AutoEnrollment
	}
	if !IsNil(o.PairingDisabled) {
		toSerialize["pairingDisabled"] = o.PairingDisabled
	}
	if !IsNil(o.IntegrityDetection) {
		toSerialize["integrityDetection"] = o.IntegrityDetection
	}
	return toSerialize, nil
}

type NullableDeviceAuthenticationPolicyMobileApplicationsInner struct {
	value *DeviceAuthenticationPolicyMobileApplicationsInner
	isSet bool
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInner) Get() *DeviceAuthenticationPolicyMobileApplicationsInner {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInner) Set(val *DeviceAuthenticationPolicyMobileApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicyMobileApplicationsInner(val *DeviceAuthenticationPolicyMobileApplicationsInner) *NullableDeviceAuthenticationPolicyMobileApplicationsInner {
	return &NullableDeviceAuthenticationPolicyMobileApplicationsInner{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicyMobileApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicyMobileApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


