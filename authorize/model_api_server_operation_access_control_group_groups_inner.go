/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the APIServerOperationAccessControlGroupGroupsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerOperationAccessControlGroupGroupsInner{}

// APIServerOperationAccessControlGroupGroupsInner struct for APIServerOperationAccessControlGroupGroupsInner
type APIServerOperationAccessControlGroupGroupsInner struct {
	// A UUID that specifies the group ID. This is a required property if `accessControl.group` is set.
	Id string `json:"id"`
}

// NewAPIServerOperationAccessControlGroupGroupsInner instantiates a new APIServerOperationAccessControlGroupGroupsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerOperationAccessControlGroupGroupsInner(id string) *APIServerOperationAccessControlGroupGroupsInner {
	this := APIServerOperationAccessControlGroupGroupsInner{}
	this.Id = id
	return &this
}

// NewAPIServerOperationAccessControlGroupGroupsInnerWithDefaults instantiates a new APIServerOperationAccessControlGroupGroupsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerOperationAccessControlGroupGroupsInnerWithDefaults() *APIServerOperationAccessControlGroupGroupsInner {
	this := APIServerOperationAccessControlGroupGroupsInner{}
	return &this
}

// GetId returns the Id field value
func (o *APIServerOperationAccessControlGroupGroupsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *APIServerOperationAccessControlGroupGroupsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *APIServerOperationAccessControlGroupGroupsInner) SetId(v string) {
	o.Id = v
}

func (o APIServerOperationAccessControlGroupGroupsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerOperationAccessControlGroupGroupsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAPIServerOperationAccessControlGroupGroupsInner struct {
	value *APIServerOperationAccessControlGroupGroupsInner
	isSet bool
}

func (v NullableAPIServerOperationAccessControlGroupGroupsInner) Get() *APIServerOperationAccessControlGroupGroupsInner {
	return v.value
}

func (v *NullableAPIServerOperationAccessControlGroupGroupsInner) Set(val *APIServerOperationAccessControlGroupGroupsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerOperationAccessControlGroupGroupsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerOperationAccessControlGroupGroupsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerOperationAccessControlGroupGroupsInner(val *APIServerOperationAccessControlGroupGroupsInner) *NullableAPIServerOperationAccessControlGroupGroupsInner {
	return &NullableAPIServerOperationAccessControlGroupGroupsInner{value: val, isSet: true}
}

func (v NullableAPIServerOperationAccessControlGroupGroupsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerOperationAccessControlGroupGroupsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


