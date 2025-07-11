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

// checks if the IntegrationVersionSAMLAllOfThirdPartyMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationVersionSAMLAllOfThirdPartyMetadata{}

// IntegrationVersionSAMLAllOfThirdPartyMetadata The third-party IdP metadata.
type IntegrationVersionSAMLAllOfThirdPartyMetadata struct {
	// URL of IdP's SAML metadata XML.
	Href *string `json:"href,omitempty"`
}

// NewIntegrationVersionSAMLAllOfThirdPartyMetadata instantiates a new IntegrationVersionSAMLAllOfThirdPartyMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationVersionSAMLAllOfThirdPartyMetadata() *IntegrationVersionSAMLAllOfThirdPartyMetadata {
	this := IntegrationVersionSAMLAllOfThirdPartyMetadata{}
	return &this
}

// NewIntegrationVersionSAMLAllOfThirdPartyMetadataWithDefaults instantiates a new IntegrationVersionSAMLAllOfThirdPartyMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationVersionSAMLAllOfThirdPartyMetadataWithDefaults() *IntegrationVersionSAMLAllOfThirdPartyMetadata {
	this := IntegrationVersionSAMLAllOfThirdPartyMetadata{}
	return &this
}

// GetHref returns the Href field value if set, zero value otherwise.
func (o *IntegrationVersionSAMLAllOfThirdPartyMetadata) GetHref() string {
	if o == nil || IsNil(o.Href) {
		var ret string
		return ret
	}
	return *o.Href
}

// GetHrefOk returns a tuple with the Href field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionSAMLAllOfThirdPartyMetadata) GetHrefOk() (*string, bool) {
	if o == nil || IsNil(o.Href) {
		return nil, false
	}
	return o.Href, true
}

// HasHref returns a boolean if a field has been set.
func (o *IntegrationVersionSAMLAllOfThirdPartyMetadata) HasHref() bool {
	if o != nil && !IsNil(o.Href) {
		return true
	}

	return false
}

// SetHref gets a reference to the given string and assigns it to the Href field.
func (o *IntegrationVersionSAMLAllOfThirdPartyMetadata) SetHref(v string) {
	o.Href = &v
}

func (o IntegrationVersionSAMLAllOfThirdPartyMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationVersionSAMLAllOfThirdPartyMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Href) {
		toSerialize["href"] = o.Href
	}
	return toSerialize, nil
}

type NullableIntegrationVersionSAMLAllOfThirdPartyMetadata struct {
	value *IntegrationVersionSAMLAllOfThirdPartyMetadata
	isSet bool
}

func (v NullableIntegrationVersionSAMLAllOfThirdPartyMetadata) Get() *IntegrationVersionSAMLAllOfThirdPartyMetadata {
	return v.value
}

func (v *NullableIntegrationVersionSAMLAllOfThirdPartyMetadata) Set(val *IntegrationVersionSAMLAllOfThirdPartyMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationVersionSAMLAllOfThirdPartyMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationVersionSAMLAllOfThirdPartyMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationVersionSAMLAllOfThirdPartyMetadata(val *IntegrationVersionSAMLAllOfThirdPartyMetadata) *NullableIntegrationVersionSAMLAllOfThirdPartyMetadata {
	return &NullableIntegrationVersionSAMLAllOfThirdPartyMetadata{value: val, isSet: true}
}

func (v NullableIntegrationVersionSAMLAllOfThirdPartyMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationVersionSAMLAllOfThirdPartyMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
