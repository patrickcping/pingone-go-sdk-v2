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

// checks if the IdentityProviderCommonIcon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProviderCommonIcon{}

// IdentityProviderCommonIcon struct for IdentityProviderCommonIcon
type IdentityProviderCommonIcon struct {
	// The ID for the IdP icon.
	Id *string `json:"id,omitempty"`
	// The HREF for the IdP icon.
	Href *string `json:"href,omitempty"`
}

// NewIdentityProviderCommonIcon instantiates a new IdentityProviderCommonIcon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderCommonIcon() *IdentityProviderCommonIcon {
	this := IdentityProviderCommonIcon{}
	return &this
}

// NewIdentityProviderCommonIconWithDefaults instantiates a new IdentityProviderCommonIcon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderCommonIconWithDefaults() *IdentityProviderCommonIcon {
	this := IdentityProviderCommonIcon{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProviderCommonIcon) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommonIcon) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProviderCommonIcon) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProviderCommonIcon) SetId(v string) {
	o.Id = &v
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *IdentityProviderCommonIcon) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderCommonIcon) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *IdentityProviderCommonIcon) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *IdentityProviderCommonIcon) SetHref(v string) {
	o.Href = &v
}

func (o IdentityProviderCommonIcon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProviderCommonIcon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableIdentityProviderCommonIcon struct {
	value *IdentityProviderCommonIcon
	isSet bool
}

func (v NullableIdentityProviderCommonIcon) Get() *IdentityProviderCommonIcon {
	return v.value
}

func (v *NullableIdentityProviderCommonIcon) Set(val *IdentityProviderCommonIcon) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderCommonIcon) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderCommonIcon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderCommonIcon(val *IdentityProviderCommonIcon) *NullableIdentityProviderCommonIcon {
	return &NullableIdentityProviderCommonIcon{value: val, isSet: true}
}

func (v NullableIdentityProviderCommonIcon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderCommonIcon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
