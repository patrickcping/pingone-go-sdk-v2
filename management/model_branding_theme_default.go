/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// checks if the BrandingThemeDefault type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingThemeDefault{}

// BrandingThemeDefault struct for BrandingThemeDefault
type BrandingThemeDefault struct {
	Links *LinksHATEOAS `json:"_links,omitempty"`
	// A boolean to specify whether the theme is the default in the environment
	Default bool `json:"default"`
}

type _BrandingThemeDefault BrandingThemeDefault

// NewBrandingThemeDefault instantiates a new BrandingThemeDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingThemeDefault(default_ bool) *BrandingThemeDefault {
	this := BrandingThemeDefault{}
	this.Default = default_
	return &this
}

// NewBrandingThemeDefaultWithDefaults instantiates a new BrandingThemeDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingThemeDefaultWithDefaults() *BrandingThemeDefault {
	this := BrandingThemeDefault{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BrandingThemeDefault) GetLinks() LinksHATEOAS {
	if o == nil || IsNil(o.Links) {
		var ret LinksHATEOAS
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingThemeDefault) GetLinksOk() (*LinksHATEOAS, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BrandingThemeDefault) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given LinksHATEOAS and assigns it to the Links field.
func (o *BrandingThemeDefault) SetLinks(v LinksHATEOAS) {
	o.Links = &v
}

// GetDefault returns the Default field value
func (o *BrandingThemeDefault) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *BrandingThemeDefault) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *BrandingThemeDefault) SetDefault(v bool) {
	o.Default = v
}

func (o BrandingThemeDefault) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingThemeDefault) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	toSerialize["default"] = o.Default
	return toSerialize, nil
}

func (o *BrandingThemeDefault) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varBrandingThemeDefault := _BrandingThemeDefault{}

	err = json.Unmarshal(bytes, &varBrandingThemeDefault)

	if err != nil {
		return err
	}

	*o = BrandingThemeDefault(varBrandingThemeDefault)

	return err
}

type NullableBrandingThemeDefault struct {
	value *BrandingThemeDefault
	isSet bool
}

func (v NullableBrandingThemeDefault) Get() *BrandingThemeDefault {
	return v.value
}

func (v *NullableBrandingThemeDefault) Set(val *BrandingThemeDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingThemeDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingThemeDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingThemeDefault(val *BrandingThemeDefault) *NullableBrandingThemeDefault {
	return &NullableBrandingThemeDefault{value: val, isSet: true}
}

func (v NullableBrandingThemeDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingThemeDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


