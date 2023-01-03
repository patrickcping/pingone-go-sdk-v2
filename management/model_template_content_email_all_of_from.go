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

// TemplateContentEmailAllOfFrom struct for TemplateContentEmailAllOfFrom
type TemplateContentEmailAllOfFrom struct {
	// The email's sender name. If the environment uses the Ping Identity email sender, the name \"PingOne\" is used. You can configure other email sender names per environment. See [Note](https://apidocs.pingidentity.com/pingone/platform/v1/api/#from-replyTo-note) for details. 
	Name *string `json:"name,omitempty"`
	// The sender email address. If the environment uses the Ping Identity email sender, or if the address field is empty, the address \"noreply@pingidentity.com\" is used. You can configure other email sender addresses per environment. See [Note](https://apidocs.pingidentity.com/pingone/platform/v1/api/#from-replyTo-note) for details. 
	Address *string `json:"address,omitempty"`
}

// NewTemplateContentEmailAllOfFrom instantiates a new TemplateContentEmailAllOfFrom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateContentEmailAllOfFrom() *TemplateContentEmailAllOfFrom {
	this := TemplateContentEmailAllOfFrom{}
	var name string = "PingOne"
	this.Name = &name
	var address string = "noreply@pingidentity.com"
	this.Address = &address
	return &this
}

// NewTemplateContentEmailAllOfFromWithDefaults instantiates a new TemplateContentEmailAllOfFrom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateContentEmailAllOfFromWithDefaults() *TemplateContentEmailAllOfFrom {
	this := TemplateContentEmailAllOfFrom{}
	var name string = "PingOne"
	this.Name = &name
	var address string = "noreply@pingidentity.com"
	this.Address = &address
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateContentEmailAllOfFrom) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmailAllOfFrom) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateContentEmailAllOfFrom) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateContentEmailAllOfFrom) SetName(v string) {
	o.Name = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *TemplateContentEmailAllOfFrom) GetAddress() string {
	if o == nil || isNil(o.Address) {
		var ret string
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentEmailAllOfFrom) GetAddressOk() (*string, bool) {
	if o == nil || isNil(o.Address) {
    return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *TemplateContentEmailAllOfFrom) HasAddress() bool {
	if o != nil && !isNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given string and assigns it to the Address field.
func (o *TemplateContentEmailAllOfFrom) SetAddress(v string) {
	o.Address = &v
}

func (o TemplateContentEmailAllOfFrom) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateContentEmailAllOfFrom struct {
	value *TemplateContentEmailAllOfFrom
	isSet bool
}

func (v NullableTemplateContentEmailAllOfFrom) Get() *TemplateContentEmailAllOfFrom {
	return v.value
}

func (v *NullableTemplateContentEmailAllOfFrom) Set(val *TemplateContentEmailAllOfFrom) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateContentEmailAllOfFrom) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateContentEmailAllOfFrom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateContentEmailAllOfFrom(val *TemplateContentEmailAllOfFrom) *NullableTemplateContentEmailAllOfFrom {
	return &NullableTemplateContentEmailAllOfFrom{value: val, isSet: true}
}

func (v NullableTemplateContentEmailAllOfFrom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateContentEmailAllOfFrom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


