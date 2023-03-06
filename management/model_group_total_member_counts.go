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

// checks if the GroupTotalMemberCounts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GroupTotalMemberCounts{}

// GroupTotalMemberCounts An object containing a users (int) property. This property lists the total number of users added to the group. You must use GET /environments/{environmentID}/groups with the include=totalMemberCounts query parameter to retrieve this property. This property is not returned with a list of groups.
type GroupTotalMemberCounts struct {
	// Number of users with direct membership
	Users *int32 `json:"users,omitempty"`
}

// NewGroupTotalMemberCounts instantiates a new GroupTotalMemberCounts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupTotalMemberCounts() *GroupTotalMemberCounts {
	this := GroupTotalMemberCounts{}
	return &this
}

// NewGroupTotalMemberCountsWithDefaults instantiates a new GroupTotalMemberCounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupTotalMemberCountsWithDefaults() *GroupTotalMemberCounts {
	this := GroupTotalMemberCounts{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *GroupTotalMemberCounts) GetUsers() int32 {
	if o == nil || IsNil(o.Users) {
		var ret int32
		return ret
	}
	return *o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupTotalMemberCounts) GetUsersOk() (*int32, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *GroupTotalMemberCounts) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given int32 and assigns it to the Users field.
func (o *GroupTotalMemberCounts) SetUsers(v int32) {
	o.Users = &v
}

func (o GroupTotalMemberCounts) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GroupTotalMemberCounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableGroupTotalMemberCounts struct {
	value *GroupTotalMemberCounts
	isSet bool
}

func (v NullableGroupTotalMemberCounts) Get() *GroupTotalMemberCounts {
	return v.value
}

func (v *NullableGroupTotalMemberCounts) Set(val *GroupTotalMemberCounts) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupTotalMemberCounts) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupTotalMemberCounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupTotalMemberCounts(val *GroupTotalMemberCounts) *NullableGroupTotalMemberCounts {
	return &NullableGroupTotalMemberCounts{value: val, isSet: true}
}

func (v NullableGroupTotalMemberCounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupTotalMemberCounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


