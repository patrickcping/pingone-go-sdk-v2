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

// checks if the UserPhoto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPhoto{}

// UserPhoto struct for UserPhoto
type UserPhoto struct {
	// A string that specifies the URI that is a uniform resource locator (as defined in Section 1.1.3 of RFC 3986) that points to a resource location representing the user’s image. This can be removed from a user by setting the photo attribute to null. If provided, the resource must be a file (for example, a GIF, JPEG, or PNG image file) rather than a web page containing an image. It must be a valid URL that starts with the HTTP or HTTPS scheme.
	Href string `json:"href"`
}

type _UserPhoto UserPhoto

// NewUserPhoto instantiates a new UserPhoto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPhoto(href string) *UserPhoto {
	this := UserPhoto{}
	this.Href = href
	return &this
}

// NewUserPhotoWithDefaults instantiates a new UserPhoto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPhotoWithDefaults() *UserPhoto {
	this := UserPhoto{}
	return &this
}

// GetHref returns the Href field value
func (o *UserPhoto) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *UserPhoto) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *UserPhoto) SetHref(v string) {
	o.Href = v
}

func (o UserPhoto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPhoto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	return toSerialize, nil
}

func (o *UserPhoto) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUserPhoto := _UserPhoto{}

	err = json.Unmarshal(bytes, &varUserPhoto)

	if err != nil {
		return err
	}

	*o = UserPhoto(varUserPhoto)

	return err
}

type NullableUserPhoto struct {
	value *UserPhoto
	isSet bool
}

func (v NullableUserPhoto) Get() *UserPhoto {
	return v.value
}

func (v *NullableUserPhoto) Set(val *UserPhoto) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPhoto) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPhoto(val *UserPhoto) *NullableUserPhoto {
	return &NullableUserPhoto{value: val, isSet: true}
}

func (v NullableUserPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


