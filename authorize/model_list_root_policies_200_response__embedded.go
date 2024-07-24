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

// checks if the ListRootPolicies200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListRootPolicies200ResponseEmbedded{}

// ListRootPolicies200ResponseEmbedded struct for ListRootPolicies200ResponseEmbedded
type ListRootPolicies200ResponseEmbedded struct {
	AuthorizationPolicies []AuthorizeEditorDataPoliciesReferenceablePolicyDTO `json:"authorizationPolicies,omitempty"`
}

// NewListRootPolicies200ResponseEmbedded instantiates a new ListRootPolicies200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListRootPolicies200ResponseEmbedded() *ListRootPolicies200ResponseEmbedded {
	this := ListRootPolicies200ResponseEmbedded{}
	return &this
}

// NewListRootPolicies200ResponseEmbeddedWithDefaults instantiates a new ListRootPolicies200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListRootPolicies200ResponseEmbeddedWithDefaults() *ListRootPolicies200ResponseEmbedded {
	this := ListRootPolicies200ResponseEmbedded{}
	return &this
}

// GetAuthorizationPolicies returns the AuthorizationPolicies field value if set, zero value otherwise.
func (o *ListRootPolicies200ResponseEmbedded) GetAuthorizationPolicies() []AuthorizeEditorDataPoliciesReferenceablePolicyDTO {
	if o == nil || IsNil(o.AuthorizationPolicies) {
		var ret []AuthorizeEditorDataPoliciesReferenceablePolicyDTO
		return ret
	}
	return o.AuthorizationPolicies
}

// GetAuthorizationPoliciesOk returns a tuple with the AuthorizationPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListRootPolicies200ResponseEmbedded) GetAuthorizationPoliciesOk() ([]AuthorizeEditorDataPoliciesReferenceablePolicyDTO, bool) {
	if o == nil || IsNil(o.AuthorizationPolicies) {
		return nil, false
	}
	return o.AuthorizationPolicies, true
}

// HasAuthorizationPolicies returns a boolean if a field has been set.
func (o *ListRootPolicies200ResponseEmbedded) HasAuthorizationPolicies() bool {
	if o != nil && !IsNil(o.AuthorizationPolicies) {
		return true
	}

	return false
}

// SetAuthorizationPolicies gets a reference to the given []AuthorizeEditorDataPoliciesReferenceablePolicyDTO and assigns it to the AuthorizationPolicies field.
func (o *ListRootPolicies200ResponseEmbedded) SetAuthorizationPolicies(v []AuthorizeEditorDataPoliciesReferenceablePolicyDTO) {
	o.AuthorizationPolicies = v
}

func (o ListRootPolicies200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListRootPolicies200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationPolicies) {
		toSerialize["authorizationPolicies"] = o.AuthorizationPolicies
	}
	return toSerialize, nil
}

type NullableListRootPolicies200ResponseEmbedded struct {
	value *ListRootPolicies200ResponseEmbedded
	isSet bool
}

func (v NullableListRootPolicies200ResponseEmbedded) Get() *ListRootPolicies200ResponseEmbedded {
	return v.value
}

func (v *NullableListRootPolicies200ResponseEmbedded) Set(val *ListRootPolicies200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListRootPolicies200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListRootPolicies200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListRootPolicies200ResponseEmbedded(val *ListRootPolicies200ResponseEmbedded) *NullableListRootPolicies200ResponseEmbedded {
	return &NullableListRootPolicies200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListRootPolicies200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListRootPolicies200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

