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

// checks if the SignOnPolicyActionLoginAllOfRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignOnPolicyActionLoginAllOfRegistration{}

// SignOnPolicyActionLoginAllOfRegistration Specifies the account registration options.
type SignOnPolicyActionLoginAllOfRegistration struct {
	// A boolean that specifies the enabled/disabled state of the policy action. The default is disabled when creating a new policy. When enabled, it allows the use of the new user registration flow. This attribute should be set to true when implementing the social login sign-on policy option.
	Enabled    bool                                                `json:"enabled"`
	External   *SignOnPolicyActionLoginAllOfRegistrationExternal   `json:"external,omitempty"`
	Population *SignOnPolicyActionLoginAllOfRegistrationPopulation `json:"population,omitempty"`
	// A boolean that specifies whether users must confirm data returned from an identity provider prior to registration. Users can modify the data and omit non-required attributes. Modified attributes are added to the user's profile during account creation. This is an optional property. If omitted, the default value is set to false.
	ConfirmIdentityProviderAttributes *bool `json:"confirmIdentityProviderAttributes,omitempty"`
}

// NewSignOnPolicyActionLoginAllOfRegistration instantiates a new SignOnPolicyActionLoginAllOfRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignOnPolicyActionLoginAllOfRegistration(enabled bool) *SignOnPolicyActionLoginAllOfRegistration {
	this := SignOnPolicyActionLoginAllOfRegistration{}
	this.Enabled = enabled
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	return &this
}

// NewSignOnPolicyActionLoginAllOfRegistrationWithDefaults instantiates a new SignOnPolicyActionLoginAllOfRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignOnPolicyActionLoginAllOfRegistrationWithDefaults() *SignOnPolicyActionLoginAllOfRegistration {
	this := SignOnPolicyActionLoginAllOfRegistration{}
	var confirmIdentityProviderAttributes bool = false
	this.ConfirmIdentityProviderAttributes = &confirmIdentityProviderAttributes
	return &this
}

// GetEnabled returns the Enabled field value
func (o *SignOnPolicyActionLoginAllOfRegistration) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOfRegistration) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SignOnPolicyActionLoginAllOfRegistration) SetEnabled(v bool) {
	o.Enabled = v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *SignOnPolicyActionLoginAllOfRegistration) GetExternal() SignOnPolicyActionLoginAllOfRegistrationExternal {
	if o == nil || IsNil(o.External) {
		var ret SignOnPolicyActionLoginAllOfRegistrationExternal
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOfRegistration) GetExternalOk() (*SignOnPolicyActionLoginAllOfRegistrationExternal, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *SignOnPolicyActionLoginAllOfRegistration) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given SignOnPolicyActionLoginAllOfRegistrationExternal and assigns it to the External field.
func (o *SignOnPolicyActionLoginAllOfRegistration) SetExternal(v SignOnPolicyActionLoginAllOfRegistrationExternal) {
	o.External = &v
}

// GetPopulation returns the Population field value if set, zero value otherwise.
func (o *SignOnPolicyActionLoginAllOfRegistration) GetPopulation() SignOnPolicyActionLoginAllOfRegistrationPopulation {
	if o == nil || IsNil(o.Population) {
		var ret SignOnPolicyActionLoginAllOfRegistrationPopulation
		return ret
	}
	return *o.Population
}

// GetPopulationOk returns a tuple with the Population field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOfRegistration) GetPopulationOk() (*SignOnPolicyActionLoginAllOfRegistrationPopulation, bool) {
	if o == nil || IsNil(o.Population) {
		return nil, false
	}
	return o.Population, true
}

// HasPopulation returns a boolean if a field has been set.
func (o *SignOnPolicyActionLoginAllOfRegistration) HasPopulation() bool {
	if o != nil && !IsNil(o.Population) {
		return true
	}

	return false
}

// SetPopulation gets a reference to the given SignOnPolicyActionLoginAllOfRegistrationPopulation and assigns it to the Population field.
func (o *SignOnPolicyActionLoginAllOfRegistration) SetPopulation(v SignOnPolicyActionLoginAllOfRegistrationPopulation) {
	o.Population = &v
}

// GetConfirmIdentityProviderAttributes returns the ConfirmIdentityProviderAttributes field value if set, zero value otherwise.
func (o *SignOnPolicyActionLoginAllOfRegistration) GetConfirmIdentityProviderAttributes() bool {
	if o == nil || IsNil(o.ConfirmIdentityProviderAttributes) {
		var ret bool
		return ret
	}
	return *o.ConfirmIdentityProviderAttributes
}

// GetConfirmIdentityProviderAttributesOk returns a tuple with the ConfirmIdentityProviderAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignOnPolicyActionLoginAllOfRegistration) GetConfirmIdentityProviderAttributesOk() (*bool, bool) {
	if o == nil || IsNil(o.ConfirmIdentityProviderAttributes) {
		return nil, false
	}
	return o.ConfirmIdentityProviderAttributes, true
}

// HasConfirmIdentityProviderAttributes returns a boolean if a field has been set.
func (o *SignOnPolicyActionLoginAllOfRegistration) HasConfirmIdentityProviderAttributes() bool {
	if o != nil && !IsNil(o.ConfirmIdentityProviderAttributes) {
		return true
	}

	return false
}

// SetConfirmIdentityProviderAttributes gets a reference to the given bool and assigns it to the ConfirmIdentityProviderAttributes field.
func (o *SignOnPolicyActionLoginAllOfRegistration) SetConfirmIdentityProviderAttributes(v bool) {
	o.ConfirmIdentityProviderAttributes = &v
}

func (o SignOnPolicyActionLoginAllOfRegistration) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignOnPolicyActionLoginAllOfRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.Population) {
		toSerialize["population"] = o.Population
	}
	if !IsNil(o.ConfirmIdentityProviderAttributes) {
		toSerialize["confirmIdentityProviderAttributes"] = o.ConfirmIdentityProviderAttributes
	}
	return toSerialize, nil
}

type NullableSignOnPolicyActionLoginAllOfRegistration struct {
	value *SignOnPolicyActionLoginAllOfRegistration
	isSet bool
}

func (v NullableSignOnPolicyActionLoginAllOfRegistration) Get() *SignOnPolicyActionLoginAllOfRegistration {
	return v.value
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistration) Set(val *SignOnPolicyActionLoginAllOfRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableSignOnPolicyActionLoginAllOfRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignOnPolicyActionLoginAllOfRegistration(val *SignOnPolicyActionLoginAllOfRegistration) *NullableSignOnPolicyActionLoginAllOfRegistration {
	return &NullableSignOnPolicyActionLoginAllOfRegistration{value: val, isSet: true}
}

func (v NullableSignOnPolicyActionLoginAllOfRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignOnPolicyActionLoginAllOfRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
