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

// checks if the MFASettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MFASettings{}

// MFASettings struct for MFASettings
type MFASettings struct {
	Links       *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	Environment *ObjectEnvironment            `json:"environment,omitempty"`
	// Deprecated
	Authentication  *MFASettingsAuthentication  `json:"authentication,omitempty"`
	Lockout         *MFASettingsLockout         `json:"lockout,omitempty"`
	Pairing         MFASettingsPairing          `json:"pairing"`
	PhoneExtensions *MFASettingsPhoneExtensions `json:"phoneExtensions,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *time.Time        `json:"updatedAt,omitempty"`
	Users     *MFASettingsUsers `json:"users,omitempty"`
}

// NewMFASettings instantiates a new MFASettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMFASettings(pairing MFASettingsPairing) *MFASettings {
	this := MFASettings{}
	this.Pairing = pairing
	return &this
}

// NewMFASettingsWithDefaults instantiates a new MFASettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMFASettingsWithDefaults() *MFASettings {
	this := MFASettings{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *MFASettings) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFASettings) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *MFASettings) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *MFASettings) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *MFASettings) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFASettings) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *MFASettings) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *MFASettings) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetAuthentication returns the Authentication field value if set, zero value otherwise.
// Deprecated
func (o *MFASettings) GetAuthentication() MFASettingsAuthentication {
	if o == nil || IsNil(o.Authentication) {
		var ret MFASettingsAuthentication
		return ret
	}
	return *o.Authentication
}

// GetAuthenticationOk returns a tuple with the Authentication field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *MFASettings) GetAuthenticationOk() (*MFASettingsAuthentication, bool) {
	if o == nil || IsNil(o.Authentication) {
		return nil, false
	}
	return o.Authentication, true
}

// HasAuthentication returns a boolean if a field has been set.
func (o *MFASettings) HasAuthentication() bool {
	if o != nil && !IsNil(o.Authentication) {
		return true
	}

	return false
}

// SetAuthentication gets a reference to the given MFASettingsAuthentication and assigns it to the Authentication field.
// Deprecated
func (o *MFASettings) SetAuthentication(v MFASettingsAuthentication) {
	o.Authentication = &v
}

// GetLockout returns the Lockout field value if set, zero value otherwise.
func (o *MFASettings) GetLockout() MFASettingsLockout {
	if o == nil || IsNil(o.Lockout) {
		var ret MFASettingsLockout
		return ret
	}
	return *o.Lockout
}

// GetLockoutOk returns a tuple with the Lockout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFASettings) GetLockoutOk() (*MFASettingsLockout, bool) {
	if o == nil || IsNil(o.Lockout) {
		return nil, false
	}
	return o.Lockout, true
}

// HasLockout returns a boolean if a field has been set.
func (o *MFASettings) HasLockout() bool {
	if o != nil && !IsNil(o.Lockout) {
		return true
	}

	return false
}

// SetLockout gets a reference to the given MFASettingsLockout and assigns it to the Lockout field.
func (o *MFASettings) SetLockout(v MFASettingsLockout) {
	o.Lockout = &v
}

// GetPairing returns the Pairing field value
func (o *MFASettings) GetPairing() MFASettingsPairing {
	if o == nil {
		var ret MFASettingsPairing
		return ret
	}

	return o.Pairing
}

// GetPairingOk returns a tuple with the Pairing field value
// and a boolean to check if the value has been set.
func (o *MFASettings) GetPairingOk() (*MFASettingsPairing, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pairing, true
}

// SetPairing sets field value
func (o *MFASettings) SetPairing(v MFASettingsPairing) {
	o.Pairing = v
}

// GetPhoneExtensions returns the PhoneExtensions field value if set, zero value otherwise.
func (o *MFASettings) GetPhoneExtensions() MFASettingsPhoneExtensions {
	if o == nil || IsNil(o.PhoneExtensions) {
		var ret MFASettingsPhoneExtensions
		return ret
	}
	return *o.PhoneExtensions
}

// GetPhoneExtensionsOk returns a tuple with the PhoneExtensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFASettings) GetPhoneExtensionsOk() (*MFASettingsPhoneExtensions, bool) {
	if o == nil || IsNil(o.PhoneExtensions) {
		return nil, false
	}
	return o.PhoneExtensions, true
}

// HasPhoneExtensions returns a boolean if a field has been set.
func (o *MFASettings) HasPhoneExtensions() bool {
	if o != nil && !IsNil(o.PhoneExtensions) {
		return true
	}

	return false
}

// SetPhoneExtensions gets a reference to the given MFASettingsPhoneExtensions and assigns it to the PhoneExtensions field.
func (o *MFASettings) SetPhoneExtensions(v MFASettingsPhoneExtensions) {
	o.PhoneExtensions = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MFASettings) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFASettings) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MFASettings) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MFASettings) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *MFASettings) GetUsers() MFASettingsUsers {
	if o == nil || IsNil(o.Users) {
		var ret MFASettingsUsers
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MFASettings) GetUsersOk() (*MFASettingsUsers, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *MFASettings) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given MFASettingsUsers and assigns it to the Users field.
func (o *MFASettings) SetUsers(v MFASettingsUsers) {
	o.Users = &v
}

func (o MFASettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MFASettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Authentication) {
		toSerialize["authentication"] = o.Authentication
	}
	if !IsNil(o.Lockout) {
		toSerialize["lockout"] = o.Lockout
	}
	toSerialize["pairing"] = o.Pairing
	if !IsNil(o.PhoneExtensions) {
		toSerialize["phoneExtensions"] = o.PhoneExtensions
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableMFASettings struct {
	value *MFASettings
	isSet bool
}

func (v NullableMFASettings) Get() *MFASettings {
	return v.value
}

func (v *NullableMFASettings) Set(val *MFASettings) {
	v.value = val
	v.isSet = true
}

func (v NullableMFASettings) IsSet() bool {
	return v.isSet
}

func (v *NullableMFASettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMFASettings(val *MFASettings) *NullableMFASettings {
	return &NullableMFASettings{value: val, isSet: true}
}

func (v NullableMFASettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMFASettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
