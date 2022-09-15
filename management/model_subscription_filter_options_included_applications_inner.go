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

// SubscriptionFilterOptionsIncludedApplicationsInner struct for SubscriptionFilterOptionsIncludedApplicationsInner
type SubscriptionFilterOptionsIncludedApplicationsInner struct {
	Id string `json:"id"`
}

// NewSubscriptionFilterOptionsIncludedApplicationsInner instantiates a new SubscriptionFilterOptionsIncludedApplicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionFilterOptionsIncludedApplicationsInner(id string) *SubscriptionFilterOptionsIncludedApplicationsInner {
	this := SubscriptionFilterOptionsIncludedApplicationsInner{}
	this.Id = id
	return &this
}

// NewSubscriptionFilterOptionsIncludedApplicationsInnerWithDefaults instantiates a new SubscriptionFilterOptionsIncludedApplicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionFilterOptionsIncludedApplicationsInnerWithDefaults() *SubscriptionFilterOptionsIncludedApplicationsInner {
	this := SubscriptionFilterOptionsIncludedApplicationsInner{}
	return &this
}

// GetId returns the Id field value
func (o *SubscriptionFilterOptionsIncludedApplicationsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubscriptionFilterOptionsIncludedApplicationsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubscriptionFilterOptionsIncludedApplicationsInner) SetId(v string) {
	o.Id = v
}

func (o SubscriptionFilterOptionsIncludedApplicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableSubscriptionFilterOptionsIncludedApplicationsInner struct {
	value *SubscriptionFilterOptionsIncludedApplicationsInner
	isSet bool
}

func (v NullableSubscriptionFilterOptionsIncludedApplicationsInner) Get() *SubscriptionFilterOptionsIncludedApplicationsInner {
	return v.value
}

func (v *NullableSubscriptionFilterOptionsIncludedApplicationsInner) Set(val *SubscriptionFilterOptionsIncludedApplicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionFilterOptionsIncludedApplicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionFilterOptionsIncludedApplicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionFilterOptionsIncludedApplicationsInner(val *SubscriptionFilterOptionsIncludedApplicationsInner) *NullableSubscriptionFilterOptionsIncludedApplicationsInner {
	return &NullableSubscriptionFilterOptionsIncludedApplicationsInner{value: val, isSet: true}
}

func (v NullableSubscriptionFilterOptionsIncludedApplicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionFilterOptionsIncludedApplicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

