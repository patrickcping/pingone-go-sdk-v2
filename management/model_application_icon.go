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

// checks if the ApplicationIcon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationIcon{}

// ApplicationIcon The HREF and the ID for the application icon.
type ApplicationIcon struct {
	Id string `json:"id"`
	Href string `json:"href"`
}

type _ApplicationIcon ApplicationIcon

// NewApplicationIcon instantiates a new ApplicationIcon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationIcon(id string, href string) *ApplicationIcon {
	this := ApplicationIcon{}
	this.Id = id
	this.Href = href
	return &this
}

// NewApplicationIconWithDefaults instantiates a new ApplicationIcon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationIconWithDefaults() *ApplicationIcon {
	this := ApplicationIcon{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationIcon) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationIcon) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationIcon) SetId(v string) {
	o.Id = v
}

// GetHref returns the Href field value
func (o *ApplicationIcon) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *ApplicationIcon) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *ApplicationIcon) SetHref(v string) {
	o.Href = v
}

func (o ApplicationIcon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationIcon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["href"] = o.Href
	return toSerialize, nil
}

func (o *ApplicationIcon) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"href",
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

	varApplicationIcon := _ApplicationIcon{}

	err = json.Unmarshal(bytes, &varApplicationIcon)

	if err != nil {
		return err
	}

	*o = ApplicationIcon(varApplicationIcon)

	return err
}

type NullableApplicationIcon struct {
	value *ApplicationIcon
	isSet bool
}

func (v NullableApplicationIcon) Get() *ApplicationIcon {
	return v.value
}

func (v *NullableApplicationIcon) Set(val *ApplicationIcon) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationIcon) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationIcon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationIcon(val *ApplicationIcon) *NullableApplicationIcon {
	return &NullableApplicationIcon{value: val, isSet: true}
}

func (v NullableApplicationIcon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationIcon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


