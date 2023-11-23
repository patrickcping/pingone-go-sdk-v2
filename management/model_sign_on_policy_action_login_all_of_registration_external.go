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

// checks if the SignOnPolicyActionLoginAllOfRegistrationExternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionLoginAllOfRegistrationExternal{}

// SignOnPolicyActionLoginAllOfRegistrationExternal struct for SignOnPolicyActionLoginAllOfRegistrationExternal
type SignOnPolicyActionLoginAllOfRegistrationExternal struct {
	// A string that specifies the link to the external identity provider's identity store. This property is set when the administrator chooses to have users register in an external identity store. This attribute can be set only when the registration.enabled property is set to false.
	Href string `json:"href"`
}

type _SignOnPolicyActionLoginAllOfRegistrationExternal SignOnPolicyActionLoginAllOfRegistrationExternal

// NewSignOnPolicyActionLoginAllOfRegistrationExternal instantiates a new SignOnPolicyActionLoginAllOfRegistrationExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLoginAllOfRegistrationExternal(href string) *SignOnPolicyActionLoginAllOfRegistrationExternal {
	this := SignOnPolicyActionLoginAllOfRegistrationExternal{}
	this.Href = href
	return &this
}

// NewSignOnPolicyActionLoginAllOfRegistrationExternalWithDefaults instantiates a new SignOnPolicyActionLoginAllOfRegistrationExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginAllOfRegistrationExternalWithDefaults() *SignOnPolicyActionLoginAllOfRegistrationExternal {
	this := SignOnPolicyActionLoginAllOfRegistrationExternal{}
	return &this
}

// GetHref returns the Href field value
func (o *SignOnPolicyActionLoginAllOfRegistrationExternal) GetHref() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Href
}

// GetHrefOk returns a tuple with the Href field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOfRegistrationExternal) GetHrefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Href, true
}

// SetHref sets field value
func (o *SignOnPolicyActionLoginAllOfRegistrationExternal) SetHref(v string) {
	o.Href = v
}

func (o SignOnPolicyActionLoginAllOfRegistrationExternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionLoginAllOfRegistrationExternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["href"] = o.Href
	return toSerialize, nil
}

func (o *SignOnPolicyActionLoginAllOfRegistrationExternal) UnmarshalJSON(bytes []byte) (err error) {
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

	varSignOnPolicyActionLoginAllOfRegistrationExternal := _SignOnPolicyActionLoginAllOfRegistrationExternal{}

	err = json.Unmarshal(bytes, &varSignOnPolicyActionLoginAllOfRegistrationExternal)

	if err != nil {
		return err
	}

	*o = SignOnPolicyActionLoginAllOfRegistrationExternal(varSignOnPolicyActionLoginAllOfRegistrationExternal)

	return err
}

type NullableSignOnPolicyActionLoginAllOfRegistrationExternal struct {
	value *SignOnPolicyActionLoginAllOfRegistrationExternal
	isSet bool
}

func (v NullableSignOnPolicyActionLoginAllOfRegistrationExternal) Get() *SignOnPolicyActionLoginAllOfRegistrationExternal {
	return v.value
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistrationExternal) Set(val *SignOnPolicyActionLoginAllOfRegistrationExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLoginAllOfRegistrationExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistrationExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLoginAllOfRegistrationExternal(val *SignOnPolicyActionLoginAllOfRegistrationExternal) *NullableSignOnPolicyActionLoginAllOfRegistrationExternal {
	return &NullableSignOnPolicyActionLoginAllOfRegistrationExternal{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLoginAllOfRegistrationExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistrationExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


