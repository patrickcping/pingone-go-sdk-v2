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

// checks if the BrandingSettingsLogo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingSettingsLogo{}

// BrandingSettingsLogo struct for BrandingSettingsLogo
type BrandingSettingsLogo struct {
	// The URL or fully qualified path to the logo file used for branding.
	Href string `json:"href"`
	// The ID of the logo image.
	Id string `json:"id"`
}

type _BrandingSettingsLogo BrandingSettingsLogo

// NewBrandingSettingsLogo instantiates a new BrandingSettingsLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingSettingsLogo(href string, id string) *BrandingSettingsLogo {
	this := BrandingSettingsLogo{}
	this.Href = href
	this.Id = id
	return &this
}

// NewBrandingSettingsLogoWithDefaults instantiates a new BrandingSettingsLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingSettingsLogoWithDefaults() *BrandingSettingsLogo {
	this := BrandingSettingsLogo{}
	return &this
}

// GetHref returns the Href field value
func (o *BrandingSettingsLogo) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *BrandingSettingsLogo) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *BrandingSettingsLogo) SetHref(v string) {
	o.Href = v
}

// GetId returns the Id field value
func (o *BrandingSettingsLogo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BrandingSettingsLogo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BrandingSettingsLogo) SetId(v string) {
	o.Id = v
}

func (o BrandingSettingsLogo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingSettingsLogo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *BrandingSettingsLogo) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"href",
		"id",
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

	varBrandingSettingsLogo := _BrandingSettingsLogo{}

	err = json.Unmarshal(bytes, &varBrandingSettingsLogo)

	if err != nil {
		return err
	}

	*o = BrandingSettingsLogo(varBrandingSettingsLogo)

	return err
}

type NullableBrandingSettingsLogo struct {
	value *BrandingSettingsLogo
	isSet bool
}

func (v NullableBrandingSettingsLogo) Get() *BrandingSettingsLogo {
	return v.value
}

func (v *NullableBrandingSettingsLogo) Set(val *BrandingSettingsLogo) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingSettingsLogo) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingSettingsLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingSettingsLogo(val *BrandingSettingsLogo) *NullableBrandingSettingsLogo {
	return &NullableBrandingSettingsLogo{value: val, isSet: true}
}

func (v NullableBrandingSettingsLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingSettingsLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


