/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OTPDeviceConfiguration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTPDeviceConfiguration{}

// OTPDeviceConfiguration struct for OTPDeviceConfiguration
type OTPDeviceConfiguration struct {
	// Controls if email or phone verification is REQUIRED, OPTIONAL, or DISABLED.
	Verify EnumVerify `json:"verify"`
	// When enabled, PingOne Verify registers the email address or phone number with PingOne MFA as a verified MFA device.
	CreateMfaDevice *bool `json:"createMfaDevice,omitempty"`
	Otp *OTPDeviceConfigurationOtp `json:"otp,omitempty"`
}

type _OTPDeviceConfiguration OTPDeviceConfiguration

// NewOTPDeviceConfiguration instantiates a new OTPDeviceConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTPDeviceConfiguration(verify EnumVerify) *OTPDeviceConfiguration {
	this := OTPDeviceConfiguration{}
	this.Verify = verify
	return &this
}

// NewOTPDeviceConfigurationWithDefaults instantiates a new OTPDeviceConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTPDeviceConfigurationWithDefaults() *OTPDeviceConfiguration {
	this := OTPDeviceConfiguration{}
	return &this
}

// GetVerify returns the Verify field value
func (o *OTPDeviceConfiguration) GetVerify() EnumVerify {
	if o == nil {
		var ret EnumVerify
		return ret
	}

	return o.Verify
}

// GetVerifyOk returns a tuple with the Verify field value
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfiguration) GetVerifyOk() (*EnumVerify, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verify, true
}

// SetVerify sets field value
func (o *OTPDeviceConfiguration) SetVerify(v EnumVerify) {
	o.Verify = v
}

// GetCreateMfaDevice returns the CreateMfaDevice field value if set, zero value otherwise.
func (o *OTPDeviceConfiguration) GetCreateMfaDevice() bool {
	if o == nil || IsNil(o.CreateMfaDevice) {
		var ret bool
		return ret
	}
	return *o.CreateMfaDevice
}

// GetCreateMfaDeviceOk returns a tuple with the CreateMfaDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfiguration) GetCreateMfaDeviceOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateMfaDevice) {
		return nil, false
	}
	return o.CreateMfaDevice, true
}

// HasCreateMfaDevice returns a boolean if a field has been set.
func (o *OTPDeviceConfiguration) HasCreateMfaDevice() bool {
	if o != nil && !IsNil(o.CreateMfaDevice) {
		return true
	}

	return false
}

// SetCreateMfaDevice gets a reference to the given bool and assigns it to the CreateMfaDevice field.
func (o *OTPDeviceConfiguration) SetCreateMfaDevice(v bool) {
	o.CreateMfaDevice = &v
}

// GetOtp returns the Otp field value if set, zero value otherwise.
func (o *OTPDeviceConfiguration) GetOtp() OTPDeviceConfigurationOtp {
	if o == nil || IsNil(o.Otp) {
		var ret OTPDeviceConfigurationOtp
		return ret
	}
	return *o.Otp
}

// GetOtpOk returns a tuple with the Otp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfiguration) GetOtpOk() (*OTPDeviceConfigurationOtp, bool) {
	if o == nil || IsNil(o.Otp) {
		return nil, false
	}
	return o.Otp, true
}

// HasOtp returns a boolean if a field has been set.
func (o *OTPDeviceConfiguration) HasOtp() bool {
	if o != nil && !IsNil(o.Otp) {
		return true
	}

	return false
}

// SetOtp gets a reference to the given OTPDeviceConfigurationOtp and assigns it to the Otp field.
func (o *OTPDeviceConfiguration) SetOtp(v OTPDeviceConfigurationOtp) {
	o.Otp = &v
}

func (o OTPDeviceConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTPDeviceConfiguration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["verify"] = o.Verify
	if !IsNil(o.CreateMfaDevice) {
		toSerialize["createMfaDevice"] = o.CreateMfaDevice
	}
	if !IsNil(o.Otp) {
		toSerialize["otp"] = o.Otp
	}
	return toSerialize, nil
}

func (o *OTPDeviceConfiguration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"verify",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varOTPDeviceConfiguration := _OTPDeviceConfiguration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOTPDeviceConfiguration)

	if err != nil {
		return err
	}

	*o = OTPDeviceConfiguration(varOTPDeviceConfiguration)

	return err
}

type NullableOTPDeviceConfiguration struct {
	value *OTPDeviceConfiguration
	isSet bool
}

func (v NullableOTPDeviceConfiguration) Get() *OTPDeviceConfiguration {
	return v.value
}

func (v *NullableOTPDeviceConfiguration) Set(val *OTPDeviceConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableOTPDeviceConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableOTPDeviceConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTPDeviceConfiguration(val *OTPDeviceConfiguration) *NullableOTPDeviceConfiguration {
	return &NullableOTPDeviceConfiguration{value: val, isSet: true}
}

func (v NullableOTPDeviceConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTPDeviceConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


