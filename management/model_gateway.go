/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// Gateway struct for Gateway
type Gateway struct {
	// A string that specifies the instance ID of the gateway. The gateway instance ID is created by the gateway when it starts up.
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Credentials []GatewayCredential `json:"credentials,omitempty"`
	// A string that specifies the resource name, which must be provided and must be unique within an environment. Valid characters are any Unicode letter, mark, numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen.
	Name string `json:"name"`
	// A string that specifies the description of the resource.
	Description *string `json:"description,omitempty"`
	Type EnumGatewayType `json:"type"`
	// A boolean that specifies whether the gateway is enabled. This is a required property.
	Enabled bool `json:"enabled"`
	SupportedVersions *GatewaySupportedVersions `json:"supportedVersions,omitempty"`
}

// NewGateway instantiates a new Gateway object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGateway(name string, type_ EnumGatewayType, enabled bool) *Gateway {
	this := Gateway{}
	this.Name = name
	this.Type = type_
	this.Enabled = enabled
	return &this
}

// NewGatewayWithDefaults instantiates a new Gateway object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayWithDefaults() *Gateway {
	this := Gateway{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Gateway) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Gateway) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Gateway) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Gateway) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Gateway) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Gateway) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *Gateway) GetCredentials() []GatewayCredential {
	if o == nil || o.Credentials == nil {
		var ret []GatewayCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetCredentialsOk() ([]GatewayCredential, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *Gateway) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []GatewayCredential and assigns it to the Credentials field.
func (o *Gateway) SetCredentials(v []GatewayCredential) {
	o.Credentials = v
}

// GetName returns the Name field value
func (o *Gateway) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Gateway) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Gateway) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Gateway) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Gateway) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Gateway) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *Gateway) GetType() EnumGatewayType {
	if o == nil {
		var ret EnumGatewayType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Gateway) GetTypeOk() (*EnumGatewayType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Gateway) SetType(v EnumGatewayType) {
	o.Type = v
}

// GetEnabled returns the Enabled field value
func (o *Gateway) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Gateway) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Gateway) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSupportedVersions returns the SupportedVersions field value if set, zero value otherwise.
func (o *Gateway) GetSupportedVersions() GatewaySupportedVersions {
	if o == nil || o.SupportedVersions == nil {
		var ret GatewaySupportedVersions
		return ret
	}
	return *o.SupportedVersions
}

// GetSupportedVersionsOk returns a tuple with the SupportedVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Gateway) GetSupportedVersionsOk() (*GatewaySupportedVersions, bool) {
	if o == nil || o.SupportedVersions == nil {
		return nil, false
	}
	return o.SupportedVersions, true
}

// HasSupportedVersions returns a boolean if a field has been set.
func (o *Gateway) HasSupportedVersions() bool {
	if o != nil && o.SupportedVersions != nil {
		return true
	}

	return false
}

// SetSupportedVersions gets a reference to the given GatewaySupportedVersions and assigns it to the SupportedVersions field.
func (o *Gateway) SetSupportedVersions(v GatewaySupportedVersions) {
	o.SupportedVersions = &v
}

func (o Gateway) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.SupportedVersions != nil {
		toSerialize["supportedVersions"] = o.SupportedVersions
	}
	return json.Marshal(toSerialize)
}

type NullableGateway struct {
	value *Gateway
	isSet bool
}

func (v NullableGateway) Get() *Gateway {
	return v.value
}

func (v *NullableGateway) Set(val *Gateway) {
	v.value = val
	v.isSet = true
}

func (v NullableGateway) IsSet() bool {
	return v.isSet
}

func (v *NullableGateway) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGateway(val *Gateway) *NullableGateway {
	return &NullableGateway{value: val, isSet: true}
}

func (v NullableGateway) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGateway) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


