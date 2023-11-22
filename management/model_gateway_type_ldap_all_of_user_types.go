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

// checks if the GatewayTypeLDAPAllOfUserTypes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayTypeLDAPAllOfUserTypes{}

// GatewayTypeLDAPAllOfUserTypes struct for GatewayTypeLDAPAllOfUserTypes
type GatewayTypeLDAPAllOfUserTypes struct {
	// Defaults to false if this property isn't specified in the request. If false, the user cannot change the password in the remote LDAP directory. In this case, operations for forgotten passwords or resetting of passwords are not available to a user referencing this gateway.
	AllowPasswordChanges *bool `json:"allowPasswordChanges,omitempty"`
	// The UUID of the user type. This correlates to the password.external.gateway.userType.id User property.
	Id *string `json:"id,omitempty"`
	// The name of the user type.
	Name string `json:"name"`
	NewUserLookup *GatewayTypeLDAPAllOfNewUserLookup `json:"newUserLookup,omitempty"`
	// A map of key/value entries used to persist the external LDAP directory attributes.
	OrderedCorrelationAttributes []string `json:"orderedCorrelationAttributes"`
	PasswordAuthority EnumGatewayPasswordAuthority `json:"passwordAuthority"`
	// The LDAP base domain name (DN) for this user type.
	SearchBaseDn string `json:"searchBaseDn"`
}

type _GatewayTypeLDAPAllOfUserTypes GatewayTypeLDAPAllOfUserTypes

// NewGatewayTypeLDAPAllOfUserTypes instantiates a new GatewayTypeLDAPAllOfUserTypes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayTypeLDAPAllOfUserTypes(name string, orderedCorrelationAttributes []string, passwordAuthority EnumGatewayPasswordAuthority, searchBaseDn string) *GatewayTypeLDAPAllOfUserTypes {
	this := GatewayTypeLDAPAllOfUserTypes{}
	this.Name = name
	this.OrderedCorrelationAttributes = orderedCorrelationAttributes
	this.PasswordAuthority = passwordAuthority
	this.SearchBaseDn = searchBaseDn
	return &this
}

// NewGatewayTypeLDAPAllOfUserTypesWithDefaults instantiates a new GatewayTypeLDAPAllOfUserTypes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayTypeLDAPAllOfUserTypesWithDefaults() *GatewayTypeLDAPAllOfUserTypes {
	this := GatewayTypeLDAPAllOfUserTypes{}
	return &this
}

// GetAllowPasswordChanges returns the AllowPasswordChanges field value if set, zero value otherwise.
func (o *GatewayTypeLDAPAllOfUserTypes) GetAllowPasswordChanges() bool {
	if o == nil || IsNil(o.AllowPasswordChanges) {
		var ret bool
		return ret
	}
	return *o.AllowPasswordChanges
}

// GetAllowPasswordChangesOk returns a tuple with the AllowPasswordChanges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) GetAllowPasswordChangesOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowPasswordChanges) {
		return nil, false
	}
	return o.AllowPasswordChanges, true
}

// HasAllowPasswordChanges returns a boolean if a field has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) HasAllowPasswordChanges() bool {
	if o != nil && !IsNil(o.AllowPasswordChanges) {
		return true
	}

	return false
}

// SetAllowPasswordChanges gets a reference to the given bool and assigns it to the AllowPasswordChanges field.
func (o *GatewayTypeLDAPAllOfUserTypes) SetAllowPasswordChanges(v bool) {
	o.AllowPasswordChanges = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayTypeLDAPAllOfUserTypes) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GatewayTypeLDAPAllOfUserTypes) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *GatewayTypeLDAPAllOfUserTypes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayTypeLDAPAllOfUserTypes) SetName(v string) {
	o.Name = v
}

// GetNewUserLookup returns the NewUserLookup field value if set, zero value otherwise.
func (o *GatewayTypeLDAPAllOfUserTypes) GetNewUserLookup() GatewayTypeLDAPAllOfNewUserLookup {
	if o == nil || IsNil(o.NewUserLookup) {
		var ret GatewayTypeLDAPAllOfNewUserLookup
		return ret
	}
	return *o.NewUserLookup
}

// GetNewUserLookupOk returns a tuple with the NewUserLookup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) GetNewUserLookupOk() (*GatewayTypeLDAPAllOfNewUserLookup, bool) {
	if o == nil || IsNil(o.NewUserLookup) {
		return nil, false
	}
	return o.NewUserLookup, true
}

// HasNewUserLookup returns a boolean if a field has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) HasNewUserLookup() bool {
	if o != nil && !IsNil(o.NewUserLookup) {
		return true
	}

	return false
}

// SetNewUserLookup gets a reference to the given GatewayTypeLDAPAllOfNewUserLookup and assigns it to the NewUserLookup field.
func (o *GatewayTypeLDAPAllOfUserTypes) SetNewUserLookup(v GatewayTypeLDAPAllOfNewUserLookup) {
	o.NewUserLookup = &v
}

// GetOrderedCorrelationAttributes returns the OrderedCorrelationAttributes field value
func (o *GatewayTypeLDAPAllOfUserTypes) GetOrderedCorrelationAttributes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.OrderedCorrelationAttributes
}

// GetOrderedCorrelationAttributesOk returns a tuple with the OrderedCorrelationAttributes field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) GetOrderedCorrelationAttributesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OrderedCorrelationAttributes, true
}

// SetOrderedCorrelationAttributes sets field value
func (o *GatewayTypeLDAPAllOfUserTypes) SetOrderedCorrelationAttributes(v []string) {
	o.OrderedCorrelationAttributes = v
}

// GetPasswordAuthority returns the PasswordAuthority field value
func (o *GatewayTypeLDAPAllOfUserTypes) GetPasswordAuthority() EnumGatewayPasswordAuthority {
	if o == nil {
		var ret EnumGatewayPasswordAuthority
		return ret
	}

	return o.PasswordAuthority
}

// GetPasswordAuthorityOk returns a tuple with the PasswordAuthority field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) GetPasswordAuthorityOk() (*EnumGatewayPasswordAuthority, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PasswordAuthority, true
}

// SetPasswordAuthority sets field value
func (o *GatewayTypeLDAPAllOfUserTypes) SetPasswordAuthority(v EnumGatewayPasswordAuthority) {
	o.PasswordAuthority = v
}

// GetSearchBaseDn returns the SearchBaseDn field value
func (o *GatewayTypeLDAPAllOfUserTypes) GetSearchBaseDn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SearchBaseDn
}

// GetSearchBaseDnOk returns a tuple with the SearchBaseDn field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfUserTypes) GetSearchBaseDnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SearchBaseDn, true
}

// SetSearchBaseDn sets field value
func (o *GatewayTypeLDAPAllOfUserTypes) SetSearchBaseDn(v string) {
	o.SearchBaseDn = v
}

func (o GatewayTypeLDAPAllOfUserTypes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayTypeLDAPAllOfUserTypes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowPasswordChanges) {
		toSerialize["allowPasswordChanges"] = o.AllowPasswordChanges
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.NewUserLookup) {
		toSerialize["newUserLookup"] = o.NewUserLookup
	}
	toSerialize["orderedCorrelationAttributes"] = o.OrderedCorrelationAttributes
	toSerialize["passwordAuthority"] = o.PasswordAuthority
	toSerialize["searchBaseDn"] = o.SearchBaseDn
	return toSerialize, nil
}

func (o *GatewayTypeLDAPAllOfUserTypes) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"orderedCorrelationAttributes",
		"passwordAuthority",
		"searchBaseDn",
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

	varGatewayTypeLDAPAllOfUserTypes := _GatewayTypeLDAPAllOfUserTypes{}

	err = json.Unmarshal(bytes, &varGatewayTypeLDAPAllOfUserTypes)

	if err != nil {
		return err
	}

	*o = GatewayTypeLDAPAllOfUserTypes(varGatewayTypeLDAPAllOfUserTypes)

	return err
}

type NullableGatewayTypeLDAPAllOfUserTypes struct {
	value *GatewayTypeLDAPAllOfUserTypes
	isSet bool
}

func (v NullableGatewayTypeLDAPAllOfUserTypes) Get() *GatewayTypeLDAPAllOfUserTypes {
	return v.value
}

func (v *NullableGatewayTypeLDAPAllOfUserTypes) Set(val *GatewayTypeLDAPAllOfUserTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTypeLDAPAllOfUserTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTypeLDAPAllOfUserTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTypeLDAPAllOfUserTypes(val *GatewayTypeLDAPAllOfUserTypes) *NullableGatewayTypeLDAPAllOfUserTypes {
	return &NullableGatewayTypeLDAPAllOfUserTypes{value: val, isSet: true}
}

func (v NullableGatewayTypeLDAPAllOfUserTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTypeLDAPAllOfUserTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


