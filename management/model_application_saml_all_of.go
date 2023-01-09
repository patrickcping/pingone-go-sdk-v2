/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// ApplicationSAMLAllOf struct for ApplicationSAMLAllOf
type ApplicationSAMLAllOf struct {
	// A string that specifies the custom home page URL for the application.
	HomePageUrl *string `json:"homePageUrl,omitempty"`
	// A string that specifies the Assertion Consumer Service URLs. The first URL in the list is used as default (there must be at least one URL). This is a required property.
	AcsUrls []string `json:"acsUrls"`
	// An integer that specifies the assertion validity duration in seconds. This is a required property.
	AssertionDuration int32 `json:"assertionDuration"`
	// A boolean that specifies whether the SAML assertion itself should be signed. The default value is true.
	AssertionSigned *bool `json:"assertionSigned,omitempty"`
	IdpSigning *ApplicationSAMLAllOfIdpSigning `json:"idpSigning,omitempty"`
	// A string that specifies the format of the Subject NameID attibute in the SAML assertion
	NameIdFormat *string `json:"nameIdFormat,omitempty"`
	// A boolean that specifies whether the SAML assertion response itself should be signed. The default value is False.
	ResponseSigned *bool `json:"responseSigned,omitempty"`
	SloBinding *EnumApplicationSAMLSloBinding `json:"sloBinding,omitempty"`
	// A string that specifies the logout endpoint URL. This is an optional property. However, if a sloEndpoint logout endpoint URL is not defined, logout actions result in an error.
	SloEndpoint *string `json:"sloEndpoint,omitempty"`
	// A string that specifies the endpoint URL to submit the logout response. If a value is not provided, the sloEndpoint property value is used to submit SLO response.
	SloResponseEndpoint *string `json:"sloResponseEndpoint,omitempty"`
	// A string that specifies the service provider entity ID used to lookup the application. This is a required property and is unique within the environment.
	SpEntityId string `json:"spEntityId"`
	SpVerification *ApplicationSAMLAllOfSpVerification `json:"spVerification,omitempty"`
}

// NewApplicationSAMLAllOf instantiates a new ApplicationSAMLAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationSAMLAllOf(acsUrls []string, assertionDuration int32, spEntityId string) *ApplicationSAMLAllOf {
	this := ApplicationSAMLAllOf{}
	this.AcsUrls = acsUrls
	this.AssertionDuration = assertionDuration
	this.SpEntityId = spEntityId
	return &this
}

// NewApplicationSAMLAllOfWithDefaults instantiates a new ApplicationSAMLAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationSAMLAllOfWithDefaults() *ApplicationSAMLAllOf {
	this := ApplicationSAMLAllOf{}
	return &this
}

// GetHomePageUrl returns the HomePageUrl field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetHomePageUrl() string {
	if o == nil || isNil(o.HomePageUrl) {
		var ret string
		return ret
	}
	return *o.HomePageUrl
}

// GetHomePageUrlOk returns a tuple with the HomePageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetHomePageUrlOk() (*string, bool) {
	if o == nil || isNil(o.HomePageUrl) {
    return nil, false
	}
	return o.HomePageUrl, true
}

// HasHomePageUrl returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasHomePageUrl() bool {
	if o != nil && !isNil(o.HomePageUrl) {
		return true
	}

	return false
}

// SetHomePageUrl gets a reference to the given string and assigns it to the HomePageUrl field.
func (o *ApplicationSAMLAllOf) SetHomePageUrl(v string) {
	o.HomePageUrl = &v
}

// GetAcsUrls returns the AcsUrls field value
func (o *ApplicationSAMLAllOf) GetAcsUrls() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AcsUrls
}

// GetAcsUrlsOk returns a tuple with the AcsUrls field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetAcsUrlsOk() ([]string, bool) {
	if o == nil {
    return nil, false
	}
	return o.AcsUrls, true
}

// SetAcsUrls sets field value
func (o *ApplicationSAMLAllOf) SetAcsUrls(v []string) {
	o.AcsUrls = v
}

// GetAssertionDuration returns the AssertionDuration field value
func (o *ApplicationSAMLAllOf) GetAssertionDuration() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AssertionDuration
}

// GetAssertionDurationOk returns a tuple with the AssertionDuration field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetAssertionDurationOk() (*int32, bool) {
	if o == nil {
    return nil, false
	}
	return &o.AssertionDuration, true
}

// SetAssertionDuration sets field value
func (o *ApplicationSAMLAllOf) SetAssertionDuration(v int32) {
	o.AssertionDuration = v
}

// GetAssertionSigned returns the AssertionSigned field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetAssertionSigned() bool {
	if o == nil || isNil(o.AssertionSigned) {
		var ret bool
		return ret
	}
	return *o.AssertionSigned
}

// GetAssertionSignedOk returns a tuple with the AssertionSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetAssertionSignedOk() (*bool, bool) {
	if o == nil || isNil(o.AssertionSigned) {
    return nil, false
	}
	return o.AssertionSigned, true
}

// HasAssertionSigned returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasAssertionSigned() bool {
	if o != nil && !isNil(o.AssertionSigned) {
		return true
	}

	return false
}

// SetAssertionSigned gets a reference to the given bool and assigns it to the AssertionSigned field.
func (o *ApplicationSAMLAllOf) SetAssertionSigned(v bool) {
	o.AssertionSigned = &v
}

// GetIdpSigning returns the IdpSigning field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetIdpSigning() ApplicationSAMLAllOfIdpSigning {
	if o == nil || isNil(o.IdpSigning) {
		var ret ApplicationSAMLAllOfIdpSigning
		return ret
	}
	return *o.IdpSigning
}

// GetIdpSigningOk returns a tuple with the IdpSigning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetIdpSigningOk() (*ApplicationSAMLAllOfIdpSigning, bool) {
	if o == nil || isNil(o.IdpSigning) {
    return nil, false
	}
	return o.IdpSigning, true
}

// HasIdpSigning returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasIdpSigning() bool {
	if o != nil && !isNil(o.IdpSigning) {
		return true
	}

	return false
}

// SetIdpSigning gets a reference to the given ApplicationSAMLAllOfIdpSigning and assigns it to the IdpSigning field.
func (o *ApplicationSAMLAllOf) SetIdpSigning(v ApplicationSAMLAllOfIdpSigning) {
	o.IdpSigning = &v
}

// GetNameIdFormat returns the NameIdFormat field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetNameIdFormat() string {
	if o == nil || isNil(o.NameIdFormat) {
		var ret string
		return ret
	}
	return *o.NameIdFormat
}

// GetNameIdFormatOk returns a tuple with the NameIdFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetNameIdFormatOk() (*string, bool) {
	if o == nil || isNil(o.NameIdFormat) {
    return nil, false
	}
	return o.NameIdFormat, true
}

// HasNameIdFormat returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasNameIdFormat() bool {
	if o != nil && !isNil(o.NameIdFormat) {
		return true
	}

	return false
}

// SetNameIdFormat gets a reference to the given string and assigns it to the NameIdFormat field.
func (o *ApplicationSAMLAllOf) SetNameIdFormat(v string) {
	o.NameIdFormat = &v
}

// GetResponseSigned returns the ResponseSigned field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetResponseSigned() bool {
	if o == nil || isNil(o.ResponseSigned) {
		var ret bool
		return ret
	}
	return *o.ResponseSigned
}

// GetResponseSignedOk returns a tuple with the ResponseSigned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetResponseSignedOk() (*bool, bool) {
	if o == nil || isNil(o.ResponseSigned) {
    return nil, false
	}
	return o.ResponseSigned, true
}

// HasResponseSigned returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasResponseSigned() bool {
	if o != nil && !isNil(o.ResponseSigned) {
		return true
	}

	return false
}

// SetResponseSigned gets a reference to the given bool and assigns it to the ResponseSigned field.
func (o *ApplicationSAMLAllOf) SetResponseSigned(v bool) {
	o.ResponseSigned = &v
}

// GetSloBinding returns the SloBinding field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetSloBinding() EnumApplicationSAMLSloBinding {
	if o == nil || isNil(o.SloBinding) {
		var ret EnumApplicationSAMLSloBinding
		return ret
	}
	return *o.SloBinding
}

// GetSloBindingOk returns a tuple with the SloBinding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetSloBindingOk() (*EnumApplicationSAMLSloBinding, bool) {
	if o == nil || isNil(o.SloBinding) {
    return nil, false
	}
	return o.SloBinding, true
}

// HasSloBinding returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasSloBinding() bool {
	if o != nil && !isNil(o.SloBinding) {
		return true
	}

	return false
}

// SetSloBinding gets a reference to the given EnumApplicationSAMLSloBinding and assigns it to the SloBinding field.
func (o *ApplicationSAMLAllOf) SetSloBinding(v EnumApplicationSAMLSloBinding) {
	o.SloBinding = &v
}

// GetSloEndpoint returns the SloEndpoint field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetSloEndpoint() string {
	if o == nil || isNil(o.SloEndpoint) {
		var ret string
		return ret
	}
	return *o.SloEndpoint
}

// GetSloEndpointOk returns a tuple with the SloEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetSloEndpointOk() (*string, bool) {
	if o == nil || isNil(o.SloEndpoint) {
    return nil, false
	}
	return o.SloEndpoint, true
}

// HasSloEndpoint returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasSloEndpoint() bool {
	if o != nil && !isNil(o.SloEndpoint) {
		return true
	}

	return false
}

// SetSloEndpoint gets a reference to the given string and assigns it to the SloEndpoint field.
func (o *ApplicationSAMLAllOf) SetSloEndpoint(v string) {
	o.SloEndpoint = &v
}

// GetSloResponseEndpoint returns the SloResponseEndpoint field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetSloResponseEndpoint() string {
	if o == nil || isNil(o.SloResponseEndpoint) {
		var ret string
		return ret
	}
	return *o.SloResponseEndpoint
}

// GetSloResponseEndpointOk returns a tuple with the SloResponseEndpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetSloResponseEndpointOk() (*string, bool) {
	if o == nil || isNil(o.SloResponseEndpoint) {
    return nil, false
	}
	return o.SloResponseEndpoint, true
}

// HasSloResponseEndpoint returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasSloResponseEndpoint() bool {
	if o != nil && !isNil(o.SloResponseEndpoint) {
		return true
	}

	return false
}

// SetSloResponseEndpoint gets a reference to the given string and assigns it to the SloResponseEndpoint field.
func (o *ApplicationSAMLAllOf) SetSloResponseEndpoint(v string) {
	o.SloResponseEndpoint = &v
}

// GetSpEntityId returns the SpEntityId field value
func (o *ApplicationSAMLAllOf) GetSpEntityId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SpEntityId
}

// GetSpEntityIdOk returns a tuple with the SpEntityId field value
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetSpEntityIdOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.SpEntityId, true
}

// SetSpEntityId sets field value
func (o *ApplicationSAMLAllOf) SetSpEntityId(v string) {
	o.SpEntityId = v
}

// GetSpVerification returns the SpVerification field value if set, zero value otherwise.
func (o *ApplicationSAMLAllOf) GetSpVerification() ApplicationSAMLAllOfSpVerification {
	if o == nil || isNil(o.SpVerification) {
		var ret ApplicationSAMLAllOfSpVerification
		return ret
	}
	return *o.SpVerification
}

// GetSpVerificationOk returns a tuple with the SpVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationSAMLAllOf) GetSpVerificationOk() (*ApplicationSAMLAllOfSpVerification, bool) {
	if o == nil || isNil(o.SpVerification) {
    return nil, false
	}
	return o.SpVerification, true
}

// HasSpVerification returns a boolean if a field has been set.
func (o *ApplicationSAMLAllOf) HasSpVerification() bool {
	if o != nil && !isNil(o.SpVerification) {
		return true
	}

	return false
}

// SetSpVerification gets a reference to the given ApplicationSAMLAllOfSpVerification and assigns it to the SpVerification field.
func (o *ApplicationSAMLAllOf) SetSpVerification(v ApplicationSAMLAllOfSpVerification) {
	o.SpVerification = &v
}

func (o ApplicationSAMLAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HomePageUrl) {
		toSerialize["homePageUrl"] = o.HomePageUrl
	}
	if true {
		toSerialize["acsUrls"] = o.AcsUrls
	}
	if true {
		toSerialize["assertionDuration"] = o.AssertionDuration
	}
	if !isNil(o.AssertionSigned) {
		toSerialize["assertionSigned"] = o.AssertionSigned
	}
	if !isNil(o.IdpSigning) {
		toSerialize["idpSigning"] = o.IdpSigning
	}
	if !isNil(o.NameIdFormat) {
		toSerialize["nameIdFormat"] = o.NameIdFormat
	}
	if !isNil(o.ResponseSigned) {
		toSerialize["responseSigned"] = o.ResponseSigned
	}
	if !isNil(o.SloBinding) {
		toSerialize["sloBinding"] = o.SloBinding
	}
	if !isNil(o.SloEndpoint) {
		toSerialize["sloEndpoint"] = o.SloEndpoint
	}
	if !isNil(o.SloResponseEndpoint) {
		toSerialize["sloResponseEndpoint"] = o.SloResponseEndpoint
	}
	if true {
		toSerialize["spEntityId"] = o.SpEntityId
	}
	if !isNil(o.SpVerification) {
		toSerialize["spVerification"] = o.SpVerification
	}
	return json.Marshal(toSerialize)
}

type NullableApplicationSAMLAllOf struct {
	value *ApplicationSAMLAllOf
	isSet bool
}

func (v NullableApplicationSAMLAllOf) Get() *ApplicationSAMLAllOf {
	return v.value
}

func (v *NullableApplicationSAMLAllOf) Set(val *ApplicationSAMLAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationSAMLAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationSAMLAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationSAMLAllOf(val *ApplicationSAMLAllOf) *NullableApplicationSAMLAllOf {
	return &NullableApplicationSAMLAllOf{value: val, isSet: true}
}

func (v NullableApplicationSAMLAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationSAMLAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


