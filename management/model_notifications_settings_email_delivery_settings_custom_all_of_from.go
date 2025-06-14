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

// checks if the NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom{}

// NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom Contains the \"from\" information to use for the notifications.
type NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom struct {
	// The \"from\" name to use in the notifications that are sent.
	Name *string `json:"name,omitempty"`
	// The \"from\" address to use in the notifications that are sent.
	Address *string `json:"address,omitempty"`
}

// NewNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom instantiates a new NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom() *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom {
	this := NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom{}
	return &this
}

// NewNotificationsSettingsEmailDeliverySettingsCustomAllOfFromWithDefaults instantiates a new NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsEmailDeliverySettingsCustomAllOfFromWithDefaults() *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom {
	this := NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) GetAddress() string {
	if o == nil || IsNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) GetAddressOk() (*string, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) SetAddress(v string) {
	o.Address = &v
}

func (o NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return toSerialize, nil
}

type NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom struct {
	value *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom
	isSet bool
}

func (v NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) Get() *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom {
	return v.value
}

func (v *NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) Set(val *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom(val *NotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) *NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom {
	return &NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom{value: val, isSet: true}
}

func (v NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsEmailDeliverySettingsCustomAllOfFrom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
