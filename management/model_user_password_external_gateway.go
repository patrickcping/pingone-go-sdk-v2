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

// checks if the UserPasswordExternalGateway type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPasswordExternalGateway{}

// UserPasswordExternalGateway An object containing the gateway properties. When this is value is specified, the user’s password is managed in an external directory. You can set the user password using Create User (Import) or Update Password (Set).
type UserPasswordExternalGateway struct {
	// The UUID of the linked gateway that references the remote directory.
	Id       *string                              `json:"id,omitempty"`
	Type     *EnumGatewayType                     `json:"type,omitempty"`
	UserType *UserPasswordExternalGatewayUserType `json:"userType,omitempty"`
	// An object that maps the external LDAP directory attributes to PingOne attributes. We use the correlationAttributes values to read the attributes from the external LDAP directory and map them to the corresponding PingOne attributes.
	CorrelationAttributes map[string]interface{} `json:"correlationAttributes,omitempty"`
}

// NewUserPasswordExternalGateway instantiates a new UserPasswordExternalGateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPasswordExternalGateway() *UserPasswordExternalGateway {
	this := UserPasswordExternalGateway{}
	return &this
}

// NewUserPasswordExternalGatewayWithDefaults instantiates a new UserPasswordExternalGateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPasswordExternalGatewayWithDefaults() *UserPasswordExternalGateway {
	this := UserPasswordExternalGateway{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserPasswordExternalGateway) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPasswordExternalGateway) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserPasswordExternalGateway) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserPasswordExternalGateway) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserPasswordExternalGateway) GetType() EnumGatewayType {
	if o == nil || IsNil(o.Type) {
		var ret EnumGatewayType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPasswordExternalGateway) GetTypeOk() (*EnumGatewayType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserPasswordExternalGateway) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumGatewayType and assigns it to the Type field.
func (o *UserPasswordExternalGateway) SetType(v EnumGatewayType) {
	o.Type = &v
}

// GetUserType returns the UserType field value if set, zero value otherwise.
func (o *UserPasswordExternalGateway) GetUserType() UserPasswordExternalGatewayUserType {
	if o == nil || IsNil(o.UserType) {
		var ret UserPasswordExternalGatewayUserType
		return ret
	}
	return *o.UserType
}

// GetUserTypeOk returns a tuple with the UserType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPasswordExternalGateway) GetUserTypeOk() (*UserPasswordExternalGatewayUserType, bool) {
	if o == nil || IsNil(o.UserType) {
		return nil, false
	}
	return o.UserType, true
}

// HasUserType returns a boolean if a field has been set.
func (o *UserPasswordExternalGateway) HasUserType() bool {
	if o != nil && !IsNil(o.UserType) {
		return true
	}

	return false
}

// SetUserType gets a reference to the given UserPasswordExternalGatewayUserType and assigns it to the UserType field.
func (o *UserPasswordExternalGateway) SetUserType(v UserPasswordExternalGatewayUserType) {
	o.UserType = &v
}

// GetCorrelationAttributes returns the CorrelationAttributes field value if set, zero value otherwise.
func (o *UserPasswordExternalGateway) GetCorrelationAttributes() map[string]interface{} {
	if o == nil || IsNil(o.CorrelationAttributes) {
		var ret map[string]interface{}
		return ret
	}
	return o.CorrelationAttributes
}

// GetCorrelationAttributesOk returns a tuple with the CorrelationAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPasswordExternalGateway) GetCorrelationAttributesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CorrelationAttributes) {
		return map[string]interface{}{}, false
	}
	return o.CorrelationAttributes, true
}

// HasCorrelationAttributes returns a boolean if a field has been set.
func (o *UserPasswordExternalGateway) HasCorrelationAttributes() bool {
	if o != nil && !IsNil(o.CorrelationAttributes) {
		return true
	}

	return false
}

// SetCorrelationAttributes gets a reference to the given map[string]interface{} and assigns it to the CorrelationAttributes field.
func (o *UserPasswordExternalGateway) SetCorrelationAttributes(v map[string]interface{}) {
	o.CorrelationAttributes = v
}

func (o UserPasswordExternalGateway) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPasswordExternalGateway) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.UserType) {
		toSerialize["userType"] = o.UserType
	}
	if !IsNil(o.CorrelationAttributes) {
		toSerialize["correlationAttributes"] = o.CorrelationAttributes
	}
	return toSerialize, nil
}

type NullableUserPasswordExternalGateway struct {
	value *UserPasswordExternalGateway
	isSet bool
}

func (v NullableUserPasswordExternalGateway) Get() *UserPasswordExternalGateway {
	return v.value
}

func (v *NullableUserPasswordExternalGateway) Set(val *UserPasswordExternalGateway) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPasswordExternalGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPasswordExternalGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPasswordExternalGateway(val *UserPasswordExternalGateway) *NullableUserPasswordExternalGateway {
	return &NullableUserPasswordExternalGateway{value: val, isSet: true}
}

func (v NullableUserPasswordExternalGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPasswordExternalGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
