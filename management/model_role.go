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

// checks if the Role type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Role{}

// Role struct for Role
type Role struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// A set of strings that specifies the scopes to which the role applies.
	ApplicableTo []EnumRoleAssignmentScopeType `json:"applicableTo,omitempty"`
	// A string that specifies the description of the role.
	Description *string `json:"description,omitempty"`
	// A string that specifies the ID of the role.
	Id *string `json:"id,omitempty"`
	Name *EnumRoleName `json:"name,omitempty"`
	// A set of permissions assigned to the role.
	Permissions []RolePermissionsInner `json:"permissions,omitempty"`
}

// NewRole instantiates a new Role object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRole() *Role {
	this := Role{}
	return &this
}

// NewRoleWithDefaults instantiates a new Role object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleWithDefaults() *Role {
	this := Role{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Role) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Role) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *Role) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetApplicableTo returns the ApplicableTo field value if set, zero value otherwise.
func (o *Role) GetApplicableTo() []EnumRoleAssignmentScopeType {
	if o == nil || IsNil(o.ApplicableTo) {
		var ret []EnumRoleAssignmentScopeType
		return ret
	}
	return o.ApplicableTo
}

// GetApplicableToOk returns a tuple with the ApplicableTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetApplicableToOk() ([]EnumRoleAssignmentScopeType, bool) {
	if o == nil || IsNil(o.ApplicableTo) {
		return nil, false
	}
	return o.ApplicableTo, true
}

// HasApplicableTo returns a boolean if a field has been set.
func (o *Role) HasApplicableTo() bool {
	if o != nil && !IsNil(o.ApplicableTo) {
		return true
	}

	return false
}

// SetApplicableTo gets a reference to the given []EnumRoleAssignmentScopeType and assigns it to the ApplicableTo field.
func (o *Role) SetApplicableTo(v []EnumRoleAssignmentScopeType) {
	o.ApplicableTo = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Role) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Role) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Role) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Role) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Role) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Role) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Role) GetName() EnumRoleName {
	if o == nil || IsNil(o.Name) {
		var ret EnumRoleName
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetNameOk() (*EnumRoleName, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Role) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given EnumRoleName and assigns it to the Name field.
func (o *Role) SetName(v EnumRoleName) {
	o.Name = &v
}

// GetPermissions returns the Permissions field value if set, zero value otherwise.
func (o *Role) GetPermissions() []RolePermissionsInner {
	if o == nil || IsNil(o.Permissions) {
		var ret []RolePermissionsInner
		return ret
	}
	return o.Permissions
}

// GetPermissionsOk returns a tuple with the Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Role) GetPermissionsOk() ([]RolePermissionsInner, bool) {
	if o == nil || IsNil(o.Permissions) {
		return nil, false
	}
	return o.Permissions, true
}

// HasPermissions returns a boolean if a field has been set.
func (o *Role) HasPermissions() bool {
	if o != nil && !IsNil(o.Permissions) {
		return true
	}

	return false
}

// SetPermissions gets a reference to the given []RolePermissionsInner and assigns it to the Permissions field.
func (o *Role) SetPermissions(v []RolePermissionsInner) {
	o.Permissions = v
}

func (o Role) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Role) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.ApplicableTo) {
		toSerialize["applicableTo"] = o.ApplicableTo
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Permissions) {
		toSerialize["permissions"] = o.Permissions
	}
	return toSerialize, nil
}

type NullableRole struct {
	value *Role
	isSet bool
}

func (v NullableRole) Get() *Role {
	return v.value
}

func (v *NullableRole) Set(val *Role) {
	v.value = val
	v.isSet = true
}

func (v NullableRole) IsSet() bool {
	return v.isSet
}

func (v *NullableRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRole(val *Role) *NullableRole {
	return &NullableRole{value: val, isSet: true}
}

func (v NullableRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


