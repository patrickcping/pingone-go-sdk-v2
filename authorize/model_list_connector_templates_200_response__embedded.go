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

// checks if the ListConnectorTemplates200ResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListConnectorTemplates200ResponseEmbedded{}

// ListConnectorTemplates200ResponseEmbedded struct for ListConnectorTemplates200ResponseEmbedded
type ListConnectorTemplates200ResponseEmbedded struct {
	AuthorizationConnectorTemplates []AuthorizeEditorDataConnectorsConnectorTemplateDTO `json:"authorizationConnectorTemplates,omitempty"`
}

// NewListConnectorTemplates200ResponseEmbedded instantiates a new ListConnectorTemplates200ResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorTemplates200ResponseEmbedded() *ListConnectorTemplates200ResponseEmbedded {
	this := ListConnectorTemplates200ResponseEmbedded{}
	return &this
}

// NewListConnectorTemplates200ResponseEmbeddedWithDefaults instantiates a new ListConnectorTemplates200ResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorTemplates200ResponseEmbeddedWithDefaults() *ListConnectorTemplates200ResponseEmbedded {
	this := ListConnectorTemplates200ResponseEmbedded{}
	return &this
}

// GetAuthorizationConnectorTemplates returns the AuthorizationConnectorTemplates field value if set, zero value otherwise.
func (o *ListConnectorTemplates200ResponseEmbedded) GetAuthorizationConnectorTemplates() []AuthorizeEditorDataConnectorsConnectorTemplateDTO {
	if o == nil || IsNil(o.AuthorizationConnectorTemplates) {
		var ret []AuthorizeEditorDataConnectorsConnectorTemplateDTO
		return ret
	}
	return o.AuthorizationConnectorTemplates
}

// GetAuthorizationConnectorTemplatesOk returns a tuple with the AuthorizationConnectorTemplates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorTemplates200ResponseEmbedded) GetAuthorizationConnectorTemplatesOk() ([]AuthorizeEditorDataConnectorsConnectorTemplateDTO, bool) {
	if o == nil || IsNil(o.AuthorizationConnectorTemplates) {
		return nil, false
	}
	return o.AuthorizationConnectorTemplates, true
}

// HasAuthorizationConnectorTemplates returns a boolean if a field has been set.
func (o *ListConnectorTemplates200ResponseEmbedded) HasAuthorizationConnectorTemplates() bool {
	if o != nil && !IsNil(o.AuthorizationConnectorTemplates) {
		return true
	}

	return false
}

// SetAuthorizationConnectorTemplates gets a reference to the given []AuthorizeEditorDataConnectorsConnectorTemplateDTO and assigns it to the AuthorizationConnectorTemplates field.
func (o *ListConnectorTemplates200ResponseEmbedded) SetAuthorizationConnectorTemplates(v []AuthorizeEditorDataConnectorsConnectorTemplateDTO) {
	o.AuthorizationConnectorTemplates = v
}

func (o ListConnectorTemplates200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListConnectorTemplates200ResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthorizationConnectorTemplates) {
		toSerialize["authorizationConnectorTemplates"] = o.AuthorizationConnectorTemplates
	}
	return toSerialize, nil
}

type NullableListConnectorTemplates200ResponseEmbedded struct {
	value *ListConnectorTemplates200ResponseEmbedded
	isSet bool
}

func (v NullableListConnectorTemplates200ResponseEmbedded) Get() *ListConnectorTemplates200ResponseEmbedded {
	return v.value
}

func (v *NullableListConnectorTemplates200ResponseEmbedded) Set(val *ListConnectorTemplates200ResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorTemplates200ResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorTemplates200ResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorTemplates200ResponseEmbedded(val *ListConnectorTemplates200ResponseEmbedded) *NullableListConnectorTemplates200ResponseEmbedded {
	return &NullableListConnectorTemplates200ResponseEmbedded{value: val, isSet: true}
}

func (v NullableListConnectorTemplates200ResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorTemplates200ResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

