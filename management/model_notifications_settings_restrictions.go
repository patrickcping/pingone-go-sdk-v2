/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the NotificationsSettingsRestrictions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettingsRestrictions{}

// NotificationsSettingsRestrictions struct for NotificationsSettingsRestrictions
type NotificationsSettingsRestrictions struct {
	SmsVoiceQuota NotificationsSettingsRestrictionsSmsVoiceQuota `json:"smsVoiceQuota"`
}

// NewNotificationsSettingsRestrictions instantiates a new NotificationsSettingsRestrictions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsRestrictions(smsVoiceQuota NotificationsSettingsRestrictionsSmsVoiceQuota) *NotificationsSettingsRestrictions {
	this := NotificationsSettingsRestrictions{}
	this.SmsVoiceQuota = smsVoiceQuota
	return &this
}

// NewNotificationsSettingsRestrictionsWithDefaults instantiates a new NotificationsSettingsRestrictions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsRestrictionsWithDefaults() *NotificationsSettingsRestrictions {
	this := NotificationsSettingsRestrictions{}
	return &this
}

// GetSmsVoiceQuota returns the SmsVoiceQuota field value
func (o *NotificationsSettingsRestrictions) GetSmsVoiceQuota() NotificationsSettingsRestrictionsSmsVoiceQuota {
	if o == nil {
		var ret NotificationsSettingsRestrictionsSmsVoiceQuota
		return ret
	}

	return o.SmsVoiceQuota
}

// GetSmsVoiceQuotaOk returns a tuple with the SmsVoiceQuota field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsRestrictions) GetSmsVoiceQuotaOk() (*NotificationsSettingsRestrictionsSmsVoiceQuota, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SmsVoiceQuota, true
}

// SetSmsVoiceQuota sets field value
func (o *NotificationsSettingsRestrictions) SetSmsVoiceQuota(v NotificationsSettingsRestrictionsSmsVoiceQuota) {
	o.SmsVoiceQuota = v
}

func (o NotificationsSettingsRestrictions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettingsRestrictions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["smsVoiceQuota"] = o.SmsVoiceQuota
	return toSerialize, nil
}

type NullableNotificationsSettingsRestrictions struct {
	value *NotificationsSettingsRestrictions
	isSet bool
}

func (v NullableNotificationsSettingsRestrictions) Get() *NotificationsSettingsRestrictions {
	return v.value
}

func (v *NullableNotificationsSettingsRestrictions) Set(val *NotificationsSettingsRestrictions) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsRestrictions) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsRestrictions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsRestrictions(val *NotificationsSettingsRestrictions) *NullableNotificationsSettingsRestrictions {
	return &NullableNotificationsSettingsRestrictions{value: val, isSet: true}
}

func (v NullableNotificationsSettingsRestrictions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsRestrictions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


