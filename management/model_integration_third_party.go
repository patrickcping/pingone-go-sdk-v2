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

// checks if the IntegrationThirdParty type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationThirdParty{}

// IntegrationThirdParty Metadata that defines the third party related to this integration.
type IntegrationThirdParty struct {
	// Name of the third party company for the listing.
	CompanyName string `json:"companyName"`
	// Names of the third party products for the integration.
	Products []string `json:"products,omitempty"`
}

// NewIntegrationThirdParty instantiates a new IntegrationThirdParty object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationThirdParty(companyName string) *IntegrationThirdParty {
	this := IntegrationThirdParty{}
	this.CompanyName = companyName
	return &this
}

// NewIntegrationThirdPartyWithDefaults instantiates a new IntegrationThirdParty object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationThirdPartyWithDefaults() *IntegrationThirdParty {
	this := IntegrationThirdParty{}
	return &this
}

// GetCompanyName returns the CompanyName field value
func (o *IntegrationThirdParty) GetCompanyName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value
// and a boolean to check if the value has been set.
func (o *IntegrationThirdParty) GetCompanyNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompanyName, true
}

// SetCompanyName sets field value
func (o *IntegrationThirdParty) SetCompanyName(v string) {
	o.CompanyName = v
}

// GetProducts returns the Products field value if set, zero value otherwise.
func (o *IntegrationThirdParty) GetProducts() []string {
	if o == nil || IsNil(o.Products) {
		var ret []string
		return ret
	}
	return o.Products
}

// GetProductsOk returns a tuple with the Products field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationThirdParty) GetProductsOk() ([]string, bool) {
	if o == nil || IsNil(o.Products) {
		return nil, false
	}
	return o.Products, true
}

// HasProducts returns a boolean if a field has been set.
func (o *IntegrationThirdParty) HasProducts() bool {
	if o != nil && !IsNil(o.Products) {
		return true
	}

	return false
}

// SetProducts gets a reference to the given []string and assigns it to the Products field.
func (o *IntegrationThirdParty) SetProducts(v []string) {
	o.Products = v
}

func (o IntegrationThirdParty) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationThirdParty) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["companyName"] = o.CompanyName
	if !IsNil(o.Products) {
		toSerialize["products"] = o.Products
	}
	return toSerialize, nil
}

type NullableIntegrationThirdParty struct {
	value *IntegrationThirdParty
	isSet bool
}

func (v NullableIntegrationThirdParty) Get() *IntegrationThirdParty {
	return v.value
}

func (v *NullableIntegrationThirdParty) Set(val *IntegrationThirdParty) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationThirdParty) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationThirdParty) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationThirdParty(val *IntegrationThirdParty) *NullableIntegrationThirdParty {
	return &NullableIntegrationThirdParty{value: val, isSet: true}
}

func (v NullableIntegrationThirdParty) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationThirdParty) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
