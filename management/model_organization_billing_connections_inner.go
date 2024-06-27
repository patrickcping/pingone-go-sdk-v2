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

// checks if the OrganizationBillingConnectionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrganizationBillingConnectionsInner{}

// OrganizationBillingConnectionsInner struct for OrganizationBillingConnectionsInner
type OrganizationBillingConnectionsInner struct {
	// A string that specifies the list of the BillingConnection resource IDs for the organization.
	Id *string `json:"id,omitempty"`
}

// NewOrganizationBillingConnectionsInner instantiates a new OrganizationBillingConnectionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrganizationBillingConnectionsInner() *OrganizationBillingConnectionsInner {
	this := OrganizationBillingConnectionsInner{}
	return &this
}

// NewOrganizationBillingConnectionsInnerWithDefaults instantiates a new OrganizationBillingConnectionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrganizationBillingConnectionsInnerWithDefaults() *OrganizationBillingConnectionsInner {
	this := OrganizationBillingConnectionsInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OrganizationBillingConnectionsInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrganizationBillingConnectionsInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OrganizationBillingConnectionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *OrganizationBillingConnectionsInner) SetId(v string) {
	o.Id = &v
}

func (o OrganizationBillingConnectionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrganizationBillingConnectionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableOrganizationBillingConnectionsInner struct {
	value *OrganizationBillingConnectionsInner
	isSet bool
}

func (v NullableOrganizationBillingConnectionsInner) Get() *OrganizationBillingConnectionsInner {
	return v.value
}

func (v *NullableOrganizationBillingConnectionsInner) Set(val *OrganizationBillingConnectionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrganizationBillingConnectionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrganizationBillingConnectionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrganizationBillingConnectionsInner(val *OrganizationBillingConnectionsInner) *NullableOrganizationBillingConnectionsInner {
	return &NullableOrganizationBillingConnectionsInner{value: val, isSet: true}
}

func (v NullableOrganizationBillingConnectionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrganizationBillingConnectionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


