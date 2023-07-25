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

// checks if the ApplicationExternalLinkAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationExternalLinkAllOf{}

// ApplicationExternalLinkAllOf struct for ApplicationExternalLinkAllOf
type ApplicationExternalLinkAllOf struct {
	// A string that specifies the custom home page URL for the application.
	HomePageUrl string `json:"homePageUrl"`
}

// NewApplicationExternalLinkAllOf instantiates a new ApplicationExternalLinkAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationExternalLinkAllOf(homePageUrl string) *ApplicationExternalLinkAllOf {
	this := ApplicationExternalLinkAllOf{}
	this.HomePageUrl = homePageUrl
	return &this
}

// NewApplicationExternalLinkAllOfWithDefaults instantiates a new ApplicationExternalLinkAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationExternalLinkAllOfWithDefaults() *ApplicationExternalLinkAllOf {
	this := ApplicationExternalLinkAllOf{}
	return &this
}

// GetHomePageUrl returns the HomePageUrl field value
func (o *ApplicationExternalLinkAllOf) GetHomePageUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HomePageUrl
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value
// and a boolean to check if the value has been set.
func (o *ApplicationExternalLinkAllOf) GetHomePageUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HomePageUrl, true
}

// SetHomePageUrl sets field value
func (o *ApplicationExternalLinkAllOf) SetHomePageUrl(v string) {
	o.HomePageUrl = v
}

func (o ApplicationExternalLinkAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationExternalLinkAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["homePageUrl"] = o.HomePageUrl
	return toSerialize, nil
}

type NullableApplicationExternalLinkAllOf struct {
	value *ApplicationExternalLinkAllOf
	isSet bool
}

func (v NullableApplicationExternalLinkAllOf) Get() *ApplicationExternalLinkAllOf {
	return v.value
}

func (v *NullableApplicationExternalLinkAllOf) Set(val *ApplicationExternalLinkAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationExternalLinkAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationExternalLinkAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationExternalLinkAllOf(val *ApplicationExternalLinkAllOf) *NullableApplicationExternalLinkAllOf {
	return &NullableApplicationExternalLinkAllOf{value: val, isSet: true}
}

func (v NullableApplicationExternalLinkAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationExternalLinkAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


