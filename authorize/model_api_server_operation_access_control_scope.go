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

// checks if the APIServerOperationAccessControlScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerOperationAccessControlScope{}

// APIServerOperationAccessControlScope Defines the scope membership requirements for the operation.
type APIServerOperationAccessControlScope struct {
	MatchType *EnumAPIServerOperationMatchType `json:"matchType,omitempty"`
	// A list of scopes that define the access requirements for the operation. The client must be authorized with `ANY` or `ALL` of the scopes to be granted access, depending on the `matchType` configuration.
	Scopes []APIServerOperationAccessControlScopeScopesInner `json:"scopes"`
}

// NewAPIServerOperationAccessControlScope instantiates a new APIServerOperationAccessControlScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerOperationAccessControlScope(scopes []APIServerOperationAccessControlScopeScopesInner) *APIServerOperationAccessControlScope {
	this := APIServerOperationAccessControlScope{}
	this.Scopes = scopes
	return &this
}

// NewAPIServerOperationAccessControlScopeWithDefaults instantiates a new APIServerOperationAccessControlScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerOperationAccessControlScopeWithDefaults() *APIServerOperationAccessControlScope {
	this := APIServerOperationAccessControlScope{}
	return &this
}

// GetMatchType returns the MatchType field value if set, zero value otherwise.
func (o *APIServerOperationAccessControlScope) GetMatchType() EnumAPIServerOperationMatchType {
	if o == nil || IsNil(o.MatchType) {
		var ret EnumAPIServerOperationMatchType
		return ret
	}
	return *o.MatchType
}

// GetMatchTypeOk returns a tuple with the MatchType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerOperationAccessControlScope) GetMatchTypeOk() (*EnumAPIServerOperationMatchType, bool) {
	if o == nil || IsNil(o.MatchType) {
		return nil, false
	}
	return o.MatchType, true
}

// HasMatchType returns a boolean if a field has been set.
func (o *APIServerOperationAccessControlScope) HasMatchType() bool {
	if o != nil && !IsNil(o.MatchType) {
		return true
	}

	return false
}

// SetMatchType gets a reference to the given EnumAPIServerOperationMatchType and assigns it to the MatchType field.
func (o *APIServerOperationAccessControlScope) SetMatchType(v EnumAPIServerOperationMatchType) {
	o.MatchType = &v
}

// GetScopes returns the Scopes field value
func (o *APIServerOperationAccessControlScope) GetScopes() []APIServerOperationAccessControlScopeScopesInner {
	if o == nil {
		var ret []APIServerOperationAccessControlScopeScopesInner
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *APIServerOperationAccessControlScope) GetScopesOk() ([]APIServerOperationAccessControlScopeScopesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *APIServerOperationAccessControlScope) SetScopes(v []APIServerOperationAccessControlScopeScopesInner) {
	o.Scopes = v
}

func (o APIServerOperationAccessControlScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerOperationAccessControlScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MatchType) {
		toSerialize["matchType"] = o.MatchType
	}
	toSerialize["scopes"] = o.Scopes
	return toSerialize, nil
}

type NullableAPIServerOperationAccessControlScope struct {
	value *APIServerOperationAccessControlScope
	isSet bool
}

func (v NullableAPIServerOperationAccessControlScope) Get() *APIServerOperationAccessControlScope {
	return v.value
}

func (v *NullableAPIServerOperationAccessControlScope) Set(val *APIServerOperationAccessControlScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerOperationAccessControlScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerOperationAccessControlScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerOperationAccessControlScope(val *APIServerOperationAccessControlScope) *NullableAPIServerOperationAccessControlScope {
	return &NullableAPIServerOperationAccessControlScope{value: val, isSet: true}
}

func (v NullableAPIServerOperationAccessControlScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerOperationAccessControlScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

