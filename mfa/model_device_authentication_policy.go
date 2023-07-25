/*
PingOne Platform API - PingOne MFA

The PingOne Platform API covering the PingOne MFA service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package mfa

import (
	"encoding/json"
	"time"
)

// checks if the DeviceAuthenticationPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceAuthenticationPolicy{}

// DeviceAuthenticationPolicy struct for DeviceAuthenticationPolicy
type DeviceAuthenticationPolicy struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// Device authentication policy's UUID.
	Id *string `json:"id,omitempty"`
	// Device authentication policy's name.
	Name string `json:"name"`
	NewDeviceNotification *EnumMFADevicePolicyNewDeviceNotification `json:"newDeviceNotification,omitempty"`
	Authentication *DeviceAuthenticationPolicyAuthentication `json:"authentication,omitempty"`
	Sms DeviceAuthenticationPolicyOfflineDevice `json:"sms"`
	Voice DeviceAuthenticationPolicyOfflineDevice `json:"voice"`
	Email DeviceAuthenticationPolicyOfflineDevice `json:"email"`
	Fido2 *DeviceAuthenticationPolicyFido2 `json:"fido2,omitempty"`
	Mobile DeviceAuthenticationPolicyMobile `json:"mobile"`
	Totp DeviceAuthenticationPolicyTotp `json:"totp"`
	// Deprecated
	SecurityKey *DeviceAuthenticationPolicyFIDODevice `json:"securityKey,omitempty"`
	// Deprecated
	Platform *DeviceAuthenticationPolicyFIDODevice `json:"platform,omitempty"`
	// A boolean that specifies whether the policy is the default for the environment.
	Default bool `json:"default"`
	// Deprecated
	ForSignOnPolicy bool `json:"forSignOnPolicy"`
	// The time the resource was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
}

// NewDeviceAuthenticationPolicy instantiates a new DeviceAuthenticationPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthenticationPolicy(name string, sms DeviceAuthenticationPolicyOfflineDevice, voice DeviceAuthenticationPolicyOfflineDevice, email DeviceAuthenticationPolicyOfflineDevice, mobile DeviceAuthenticationPolicyMobile, totp DeviceAuthenticationPolicyTotp, default_ bool, forSignOnPolicy bool) *DeviceAuthenticationPolicy {
	this := DeviceAuthenticationPolicy{}
	this.Name = name
	this.Sms = sms
	this.Voice = voice
	this.Email = email
	this.Mobile = mobile
	this.Totp = totp
	this.Default = default_
	this.ForSignOnPolicy = forSignOnPolicy
	return &this
}

// NewDeviceAuthenticationPolicyWithDefaults instantiates a new DeviceAuthenticationPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthenticationPolicyWithDefaults() *DeviceAuthenticationPolicy {
	this := DeviceAuthenticationPolicy{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *DeviceAuthenticationPolicy) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *DeviceAuthenticationPolicy) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DeviceAuthenticationPolicy) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *DeviceAuthenticationPolicy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DeviceAuthenticationPolicy) SetName(v string) {
	o.Name = v
}

// GetNewDeviceNotification returns the NewDeviceNotification field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetNewDeviceNotification() EnumMFADevicePolicyNewDeviceNotification {
	if o == nil || IsNil(o.NewDeviceNotification) {
		var ret EnumMFADevicePolicyNewDeviceNotification
		return ret
	}
	return *o.NewDeviceNotification
}

// GetNewDeviceNotificationOk returns a tuple with the NewDeviceNotification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetNewDeviceNotificationOk() (*EnumMFADevicePolicyNewDeviceNotification, bool) {
	if o == nil || IsNil(o.NewDeviceNotification) {
		return nil, false
	}
	return o.NewDeviceNotification, true
}

// HasNewDeviceNotification returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasNewDeviceNotification() bool {
	if o != nil && !IsNil(o.NewDeviceNotification) {
		return true
	}

	return false
}

// SetNewDeviceNotification gets a reference to the given EnumMFADevicePolicyNewDeviceNotification and assigns it to the NewDeviceNotification field.
func (o *DeviceAuthenticationPolicy) SetNewDeviceNotification(v EnumMFADevicePolicyNewDeviceNotification) {
	o.NewDeviceNotification = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetAuthentication() DeviceAuthenticationPolicyAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret DeviceAuthenticationPolicyAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetAuthenticationOk() (*DeviceAuthenticationPolicyAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given DeviceAuthenticationPolicyAuthentication and assigns it to the Authentication field.
func (o *DeviceAuthenticationPolicy) SetAuthentication(v DeviceAuthenticationPolicyAuthentication) {
	o.Authentication = &v
}

// GetSms returns the Sms field value
func (o *DeviceAuthenticationPolicy) GetSms() DeviceAuthenticationPolicyOfflineDevice {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDevice
		return ret
	}

	return o.Sms
}

// GetSmsOk returns a tuple with the Sms field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetSmsOk() (*DeviceAuthenticationPolicyOfflineDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sms, true
}

// SetSms sets field value
func (o *DeviceAuthenticationPolicy) SetSms(v DeviceAuthenticationPolicyOfflineDevice) {
	o.Sms = v
}

// GetVoice returns the Voice field value
func (o *DeviceAuthenticationPolicy) GetVoice() DeviceAuthenticationPolicyOfflineDevice {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDevice
		return ret
	}

	return o.Voice
}

// GetVoiceOk returns a tuple with the Voice field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetVoiceOk() (*DeviceAuthenticationPolicyOfflineDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Voice, true
}

// SetVoice sets field value
func (o *DeviceAuthenticationPolicy) SetVoice(v DeviceAuthenticationPolicyOfflineDevice) {
	o.Voice = v
}

// GetEmail returns the Email field value
func (o *DeviceAuthenticationPolicy) GetEmail() DeviceAuthenticationPolicyOfflineDevice {
	if o == nil {
		var ret DeviceAuthenticationPolicyOfflineDevice
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetEmailOk() (*DeviceAuthenticationPolicyOfflineDevice, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *DeviceAuthenticationPolicy) SetEmail(v DeviceAuthenticationPolicyOfflineDevice) {
	o.Email = v
}

// GetFido2 returns the Fido2 field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetFido2() DeviceAuthenticationPolicyFido2 {
	if o == nil || IsNil(o.Fido2) {
		var ret DeviceAuthenticationPolicyFido2
		return ret
	}
	return *o.Fido2
}

// GetFido2Ok returns a tuple with the Fido2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetFido2Ok() (*DeviceAuthenticationPolicyFido2, bool) {
	if o == nil || IsNil(o.Fido2) {
		return nil, false
	}
	return o.Fido2, true
}

// HasFido2 returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasFido2() bool {
	if o != nil && !IsNil(o.Fido2) {
		return true
	}

	return false
}

// SetFido2 gets a reference to the given DeviceAuthenticationPolicyFido2 and assigns it to the Fido2 field.
func (o *DeviceAuthenticationPolicy) SetFido2(v DeviceAuthenticationPolicyFido2) {
	o.Fido2 = &v
}

// GetMobile returns the Mobile field value
func (o *DeviceAuthenticationPolicy) GetMobile() DeviceAuthenticationPolicyMobile {
	if o == nil {
		var ret DeviceAuthenticationPolicyMobile
		return ret
	}

	return o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetMobileOk() (*DeviceAuthenticationPolicyMobile, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mobile, true
}

// SetMobile sets field value
func (o *DeviceAuthenticationPolicy) SetMobile(v DeviceAuthenticationPolicyMobile) {
	o.Mobile = v
}

// GetTotp returns the Totp field value
func (o *DeviceAuthenticationPolicy) GetTotp() DeviceAuthenticationPolicyTotp {
	if o == nil {
		var ret DeviceAuthenticationPolicyTotp
		return ret
	}

	return o.Totp
}

// GetTotpOk returns a tuple with the Totp field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetTotpOk() (*DeviceAuthenticationPolicyTotp, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Totp, true
}

// SetTotp sets field value
func (o *DeviceAuthenticationPolicy) SetTotp(v DeviceAuthenticationPolicyTotp) {
	o.Totp = v
}

// GetSecurityKey returns the SecurityKey field value if set, zero value otherwise.
// Deprecated
func (o *DeviceAuthenticationPolicy) GetSecurityKey() DeviceAuthenticationPolicyFIDODevice {
	if o == nil || IsNil(o.SecurityKey) {
		var ret DeviceAuthenticationPolicyFIDODevice
		return ret
	}
	return *o.SecurityKey
}

// GetSecurityKeyOk returns a tuple with the SecurityKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DeviceAuthenticationPolicy) GetSecurityKeyOk() (*DeviceAuthenticationPolicyFIDODevice, bool) {
	if o == nil || IsNil(o.SecurityKey) {
		return nil, false
	}
	return o.SecurityKey, true
}

// HasSecurityKey returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasSecurityKey() bool {
	if o != nil && !IsNil(o.SecurityKey) {
		return true
	}

	return false
}

// SetSecurityKey gets a reference to the given DeviceAuthenticationPolicyFIDODevice and assigns it to the SecurityKey field.
// Deprecated
func (o *DeviceAuthenticationPolicy) SetSecurityKey(v DeviceAuthenticationPolicyFIDODevice) {
	o.SecurityKey = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
// Deprecated
func (o *DeviceAuthenticationPolicy) GetPlatform() DeviceAuthenticationPolicyFIDODevice {
	if o == nil || IsNil(o.Platform) {
		var ret DeviceAuthenticationPolicyFIDODevice
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *DeviceAuthenticationPolicy) GetPlatformOk() (*DeviceAuthenticationPolicyFIDODevice, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// HasPlatform returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given DeviceAuthenticationPolicyFIDODevice and assigns it to the Platform field.
// Deprecated
func (o *DeviceAuthenticationPolicy) SetPlatform(v DeviceAuthenticationPolicyFIDODevice) {
	o.Platform = &v
}

// GetDefault returns the Default field value
func (o *DeviceAuthenticationPolicy) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *DeviceAuthenticationPolicy) SetDefault(v bool) {
	o.Default = v
}

// GetForSignOnPolicy returns the ForSignOnPolicy field value
// Deprecated
func (o *DeviceAuthenticationPolicy) GetForSignOnPolicy() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ForSignOnPolicy
}

// GetForSignOnPolicyOk returns a tuple with the ForSignOnPolicy field value
// and a boolean to check if the value has been set.
// Deprecated
func (o *DeviceAuthenticationPolicy) GetForSignOnPolicyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ForSignOnPolicy, true
}

// SetForSignOnPolicy sets field value
// Deprecated
func (o *DeviceAuthenticationPolicy) SetForSignOnPolicy(v bool) {
	o.ForSignOnPolicy = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DeviceAuthenticationPolicy) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthenticationPolicy) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DeviceAuthenticationPolicy) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DeviceAuthenticationPolicy) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o DeviceAuthenticationPolicy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceAuthenticationPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	toSerialize["name"] = o.Name
	if !IsNil(o.NewDeviceNotification) {
		toSerialize["newDeviceNotification"] = o.NewDeviceNotification
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	toSerialize["sms"] = o.Sms
	toSerialize["voice"] = o.Voice
	toSerialize["email"] = o.Email
	if !IsNil(o.Fido2) {
		toSerialize["fido2"] = o.Fido2
	}
	toSerialize["mobile"] = o.Mobile
	toSerialize["totp"] = o.Totp
	if !IsNil(o.SecurityKey) {
		toSerialize["securityKey"] = o.SecurityKey
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	toSerialize["default"] = o.Default
	toSerialize["forSignOnPolicy"] = o.ForSignOnPolicy
	// skip: updatedAt is readOnly
	return toSerialize, nil
}

type NullableDeviceAuthenticationPolicy struct {
	value *DeviceAuthenticationPolicy
	isSet bool
}

func (v NullableDeviceAuthenticationPolicy) Get() *DeviceAuthenticationPolicy {
	return v.value
}

func (v *NullableDeviceAuthenticationPolicy) Set(val *DeviceAuthenticationPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthenticationPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthenticationPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthenticationPolicy(val *DeviceAuthenticationPolicy) *NullableDeviceAuthenticationPolicy {
	return &NullableDeviceAuthenticationPolicy{value: val, isSet: true}
}

func (v NullableDeviceAuthenticationPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthenticationPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


