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

// TemplateContentPushAllOf struct for TemplateContentPushAllOf
type TemplateContentPushAllOf struct {
	// The push title (maximum 200 characters). If supported, this can include variables.
	Title *string `json:"title,omitempty"`
}

// NewTemplateContentPushAllOf instantiates a new TemplateContentPushAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateContentPushAllOf() *TemplateContentPushAllOf {
	this := TemplateContentPushAllOf{}
	return &this
}

// NewTemplateContentPushAllOfWithDefaults instantiates a new TemplateContentPushAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateContentPushAllOfWithDefaults() *TemplateContentPushAllOf {
	this := TemplateContentPushAllOf{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *TemplateContentPushAllOf) GetTitle() string {
	if o == nil || isNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateContentPushAllOf) GetTitleOk() (*string, bool) {
	if o == nil || isNil(o.Title) {
    return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *TemplateContentPushAllOf) HasTitle() bool {
	if o != nil && !isNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *TemplateContentPushAllOf) SetTitle(v string) {
	o.Title = &v
}

func (o TemplateContentPushAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	return json.Marshal(toSerialize)
}

type NullableTemplateContentPushAllOf struct {
	value *TemplateContentPushAllOf
	isSet bool
}

func (v NullableTemplateContentPushAllOf) Get() *TemplateContentPushAllOf {
	return v.value
}

func (v *NullableTemplateContentPushAllOf) Set(val *TemplateContentPushAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateContentPushAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateContentPushAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateContentPushAllOf(val *TemplateContentPushAllOf) *NullableTemplateContentPushAllOf {
	return &NullableTemplateContentPushAllOf{value: val, isSet: true}
}

func (v NullableTemplateContentPushAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateContentPushAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

