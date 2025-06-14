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

// checks if the IntegrationVersionCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationVersionCommon{}

// IntegrationVersionCommon struct for IntegrationVersionCommon
type IntegrationVersionCommon struct {
	Configuration *IntegrationVersionCommonConfiguration `json:"configuration,omitempty"`
	// Unicode characters. The description of this integration metadata version.
	Description *string `json:"description,omitempty"`
	// The platform-generated ID of this integration metadata version.
	Id          *string                              `json:"id,omitempty"`
	Integration *IntegrationVersionCommonIntegration `json:"integration,omitempty"`
	// A unique name for the integration metadata version.
	Name string `json:"name"`
	// A unique number for the integration version.
	Number string                      `json:"number"`
	Type   *EnumIntegrationVersionType `json:"type,omitempty"`
}

// NewIntegrationVersionCommon instantiates a new IntegrationVersionCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationVersionCommon(name string, number string) *IntegrationVersionCommon {
	this := IntegrationVersionCommon{}
	this.Name = name
	this.Number = number
	return &this
}

// NewIntegrationVersionCommonWithDefaults instantiates a new IntegrationVersionCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationVersionCommonWithDefaults() *IntegrationVersionCommon {
	this := IntegrationVersionCommon{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *IntegrationVersionCommon) GetConfiguration() IntegrationVersionCommonConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret IntegrationVersionCommonConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionCommon) GetConfigurationOk() (*IntegrationVersionCommonConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *IntegrationVersionCommon) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given IntegrationVersionCommonConfiguration and assigns it to the Configuration field.
func (o *IntegrationVersionCommon) SetConfiguration(v IntegrationVersionCommonConfiguration) {
	o.Configuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationVersionCommon) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionCommon) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationVersionCommon) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationVersionCommon) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationVersionCommon) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionCommon) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationVersionCommon) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationVersionCommon) SetId(v string) {
	o.Id = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *IntegrationVersionCommon) GetIntegration() IntegrationVersionCommonIntegration {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationVersionCommonIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionCommon) GetIntegrationOk() (*IntegrationVersionCommonIntegration, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *IntegrationVersionCommon) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationVersionCommonIntegration and assigns it to the Integration field.
func (o *IntegrationVersionCommon) SetIntegration(v IntegrationVersionCommonIntegration) {
	o.Integration = &v
}

// GetName returns the Name field value
func (o *IntegrationVersionCommon) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegrationVersionCommon) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntegrationVersionCommon) SetName(v string) {
	o.Name = v
}

// GetNumber returns the Number field value
func (o *IntegrationVersionCommon) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *IntegrationVersionCommon) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *IntegrationVersionCommon) SetNumber(v string) {
	o.Number = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegrationVersionCommon) GetType() EnumIntegrationVersionType {
	if o == nil || IsNil(o.Type) {
		var ret EnumIntegrationVersionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionCommon) GetTypeOk() (*EnumIntegrationVersionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegrationVersionCommon) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumIntegrationVersionType and assigns it to the Type field.
func (o *IntegrationVersionCommon) SetType(v EnumIntegrationVersionType) {
	o.Type = &v
}

func (o IntegrationVersionCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationVersionCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Integration) {
		toSerialize["integration"] = o.Integration
	}
	toSerialize["name"] = o.Name
	toSerialize["number"] = o.Number
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableIntegrationVersionCommon struct {
	value *IntegrationVersionCommon
	isSet bool
}

func (v NullableIntegrationVersionCommon) Get() *IntegrationVersionCommon {
	return v.value
}

func (v *NullableIntegrationVersionCommon) Set(val *IntegrationVersionCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationVersionCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationVersionCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationVersionCommon(val *IntegrationVersionCommon) *NullableIntegrationVersionCommon {
	return &NullableIntegrationVersionCommon{value: val, isSet: true}
}

func (v NullableIntegrationVersionCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationVersionCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
