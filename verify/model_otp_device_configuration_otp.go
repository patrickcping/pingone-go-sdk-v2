/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
)

// checks if the OTPDeviceConfigurationOtp type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTPDeviceConfigurationOtp{}

// OTPDeviceConfigurationOtp struct for OTPDeviceConfigurationOtp
type OTPDeviceConfigurationOtp struct {
	Attempts     OTPDeviceConfigurationOtpAttempts     `json:"attempts"`
	Deliveries   OTPDeviceConfigurationOtpDeliveries   `json:"deliveries"`
	LifeTime     OTPDeviceConfigurationOtpLifeTime     `json:"lifeTime"`
	Notification OTPDeviceConfigurationOtpNotification `json:"notification"`
}

// NewOTPDeviceConfigurationOtp instantiates a new OTPDeviceConfigurationOtp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTPDeviceConfigurationOtp(attempts OTPDeviceConfigurationOtpAttempts, deliveries OTPDeviceConfigurationOtpDeliveries, lifeTime OTPDeviceConfigurationOtpLifeTime, notification OTPDeviceConfigurationOtpNotification) *OTPDeviceConfigurationOtp {
	this := OTPDeviceConfigurationOtp{}
	this.Attempts = attempts
	this.Deliveries = deliveries
	this.LifeTime = lifeTime
	this.Notification = notification
	return &this
}

// NewOTPDeviceConfigurationOtpWithDefaults instantiates a new OTPDeviceConfigurationOtp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTPDeviceConfigurationOtpWithDefaults() *OTPDeviceConfigurationOtp {
	this := OTPDeviceConfigurationOtp{}
	return &this
}

// GetAttempts returns the Attempts field value
func (o *OTPDeviceConfigurationOtp) GetAttempts() OTPDeviceConfigurationOtpAttempts {
	if o == nil {
		var ret OTPDeviceConfigurationOtpAttempts
		return ret
	}

	return o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfigurationOtp) GetAttemptsOk() (*OTPDeviceConfigurationOtpAttempts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attempts, true
}

// SetAttempts sets field value
func (o *OTPDeviceConfigurationOtp) SetAttempts(v OTPDeviceConfigurationOtpAttempts) {
	o.Attempts = v
}

// GetDeliveries returns the Deliveries field value
func (o *OTPDeviceConfigurationOtp) GetDeliveries() OTPDeviceConfigurationOtpDeliveries {
	if o == nil {
		var ret OTPDeviceConfigurationOtpDeliveries
		return ret
	}

	return o.Deliveries
}

// GetDeliveriesOk returns a tuple with the Deliveries field value
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfigurationOtp) GetDeliveriesOk() (*OTPDeviceConfigurationOtpDeliveries, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deliveries, true
}

// SetDeliveries sets field value
func (o *OTPDeviceConfigurationOtp) SetDeliveries(v OTPDeviceConfigurationOtpDeliveries) {
	o.Deliveries = v
}

// GetLifeTime returns the LifeTime field value
func (o *OTPDeviceConfigurationOtp) GetLifeTime() OTPDeviceConfigurationOtpLifeTime {
	if o == nil {
		var ret OTPDeviceConfigurationOtpLifeTime
		return ret
	}

	return o.LifeTime
}

// GetLifeTimeOk returns a tuple with the LifeTime field value
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfigurationOtp) GetLifeTimeOk() (*OTPDeviceConfigurationOtpLifeTime, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LifeTime, true
}

// SetLifeTime sets field value
func (o *OTPDeviceConfigurationOtp) SetLifeTime(v OTPDeviceConfigurationOtpLifeTime) {
	o.LifeTime = v
}

// GetNotification returns the Notification field value
func (o *OTPDeviceConfigurationOtp) GetNotification() OTPDeviceConfigurationOtpNotification {
	if o == nil {
		var ret OTPDeviceConfigurationOtpNotification
		return ret
	}

	return o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfigurationOtp) GetNotificationOk() (*OTPDeviceConfigurationOtpNotification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notification, true
}

// SetNotification sets field value
func (o *OTPDeviceConfigurationOtp) SetNotification(v OTPDeviceConfigurationOtpNotification) {
	o.Notification = v
}

func (o OTPDeviceConfigurationOtp) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTPDeviceConfigurationOtp) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["attempts"] = o.Attempts
	toSerialize["deliveries"] = o.Deliveries
	toSerialize["lifeTime"] = o.LifeTime
	toSerialize["notification"] = o.Notification
	return toSerialize, nil
}

type NullableOTPDeviceConfigurationOtp struct {
	value *OTPDeviceConfigurationOtp
	isSet bool
}

func (v NullableOTPDeviceConfigurationOtp) Get() *OTPDeviceConfigurationOtp {
	return v.value
}

func (v *NullableOTPDeviceConfigurationOtp) Set(val *OTPDeviceConfigurationOtp) {
	v.value = val
	v.isSet = true
}

func (v NullableOTPDeviceConfigurationOtp) IsSet() bool {
	return v.isSet
}

func (v *NullableOTPDeviceConfigurationOtp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTPDeviceConfigurationOtp(val *OTPDeviceConfigurationOtp) *NullableOTPDeviceConfigurationOtp {
	return &NullableOTPDeviceConfigurationOtp{value: val, isSet: true}
}

func (v NullableOTPDeviceConfigurationOtp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTPDeviceConfigurationOtp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
