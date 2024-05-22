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

// checks if the ApplicationRolePermission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationRolePermission{}

// ApplicationRolePermission struct for ApplicationRolePermission
type ApplicationRolePermission struct {
	// The ID of the application resource permission to associate with this role.
	Id string `json:"id"`
	Permission *ApplicationRolePermissionPermission `json:"permission,omitempty"`
}

// NewApplicationRolePermission instantiates a new ApplicationRolePermission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationRolePermission(id string) *ApplicationRolePermission {
	this := ApplicationRolePermission{}
	this.Id = id
	return &this
}

// NewApplicationRolePermissionWithDefaults instantiates a new ApplicationRolePermission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationRolePermissionWithDefaults() *ApplicationRolePermission {
	this := ApplicationRolePermission{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationRolePermission) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationRolePermission) SetId(v string) {
	o.Id = v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *ApplicationRolePermission) GetPermission() ApplicationRolePermissionPermission {
	if o == nil || IsNil(o.Permission) {
		var ret ApplicationRolePermissionPermission
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationRolePermission) GetPermissionOk() (*ApplicationRolePermissionPermission, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *ApplicationRolePermission) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given ApplicationRolePermissionPermission and assigns it to the Permission field.
func (o *ApplicationRolePermission) SetPermission(v ApplicationRolePermissionPermission) {
	o.Permission = &v
}

func (o ApplicationRolePermission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationRolePermission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	return toSerialize, nil
}

type NullableApplicationRolePermission struct {
	value *ApplicationRolePermission
	isSet bool
}

func (v NullableApplicationRolePermission) Get() *ApplicationRolePermission {
	return v.value
}

func (v *NullableApplicationRolePermission) Set(val *ApplicationRolePermission) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationRolePermission) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationRolePermission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationRolePermission(val *ApplicationRolePermission) *NullableApplicationRolePermission {
	return &NullableApplicationRolePermission{value: val, isSet: true}
}

func (v NullableApplicationRolePermission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationRolePermission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

