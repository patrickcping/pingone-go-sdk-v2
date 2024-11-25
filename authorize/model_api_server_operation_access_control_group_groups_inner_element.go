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

// checks if the APIServerOperationAccessControlGroupGroupsInnerElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerOperationAccessControlGroupGroupsInnerElement{}

// APIServerOperationAccessControlGroupGroupsInnerElement The ID of the group, wrapped in an object, for future extensibility. This is a required property if `operations.value.accessControl.group` is set.
type APIServerOperationAccessControlGroupGroupsInnerElement struct {
	// A UUID that specifies the group ID. This is a required property if `accessControl.group` is set.
	Id string `json:"id"`
}

// NewAPIServerOperationAccessControlGroupGroupsInnerElement instantiates a new APIServerOperationAccessControlGroupGroupsInnerElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerOperationAccessControlGroupGroupsInnerElement(id string) *APIServerOperationAccessControlGroupGroupsInnerElement {
	this := APIServerOperationAccessControlGroupGroupsInnerElement{}
	this.Id = id
	return &this
}

// NewAPIServerOperationAccessControlGroupGroupsInnerElementWithDefaults instantiates a new APIServerOperationAccessControlGroupGroupsInnerElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerOperationAccessControlGroupGroupsInnerElementWithDefaults() *APIServerOperationAccessControlGroupGroupsInnerElement {
	this := APIServerOperationAccessControlGroupGroupsInnerElement{}
	return &this
}

// GetId returns the Id field value
func (o *APIServerOperationAccessControlGroupGroupsInnerElement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *APIServerOperationAccessControlGroupGroupsInnerElement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *APIServerOperationAccessControlGroupGroupsInnerElement) SetId(v string) {
	o.Id = v
}

func (o APIServerOperationAccessControlGroupGroupsInnerElement) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerOperationAccessControlGroupGroupsInnerElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAPIServerOperationAccessControlGroupGroupsInnerElement struct {
	value *APIServerOperationAccessControlGroupGroupsInnerElement
	isSet bool
}

func (v NullableAPIServerOperationAccessControlGroupGroupsInnerElement) Get() *APIServerOperationAccessControlGroupGroupsInnerElement {
	return v.value
}

func (v *NullableAPIServerOperationAccessControlGroupGroupsInnerElement) Set(val *APIServerOperationAccessControlGroupGroupsInnerElement) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerOperationAccessControlGroupGroupsInnerElement) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerOperationAccessControlGroupGroupsInnerElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerOperationAccessControlGroupGroupsInnerElement(val *APIServerOperationAccessControlGroupGroupsInnerElement) *NullableAPIServerOperationAccessControlGroupGroupsInnerElement {
	return &NullableAPIServerOperationAccessControlGroupGroupsInnerElement{value: val, isSet: true}
}

func (v NullableAPIServerOperationAccessControlGroupGroupsInnerElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerOperationAccessControlGroupGroupsInnerElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

