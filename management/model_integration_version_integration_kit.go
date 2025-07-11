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

// checks if the IntegrationVersionIntegrationKit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationVersionIntegrationKit{}

// IntegrationVersionIntegrationKit struct for IntegrationVersionIntegrationKit
type IntegrationVersionIntegrationKit struct {
	Configuration *IntegrationVersionCommonConfiguration `json:"configuration,omitempty"`
	// Unicode characters. The description of this integration metadata version.
	Description *string `json:"description,omitempty"`
	// The platform-generated ID of this integration metadata version.
	Id          *string                              `json:"id,omitempty"`
	Integration *IntegrationVersionCommonIntegration `json:"integration,omitempty"`
	// A unique name for the integration metadata version.
	Name string `json:"name"`
	// Unique number for the integration version.
	Number string                      `json:"number"`
	Type   *EnumIntegrationVersionType `json:"type,omitempty"`
	// Absolute URL to the documentation.
	DocumentationUrl *string `json:"documentationUrl,omitempty"`
	// EOL support date in the form yyyy-mm-dd.
	EndOfLifeOn    *string                                              `json:"endOfLifeOn,omitempty"`
	IntegratedWith *IntegrationVersionIntegrationKitAllOfIntegratedWith `json:"integratedWith,omitempty"`
	// Release date in the form yyyy-mm-dd.
	ReleasedOn string `json:"releasedOn"`
}

// NewIntegrationVersionIntegrationKit instantiates a new IntegrationVersionIntegrationKit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationVersionIntegrationKit(name string, number string, releasedOn string) *IntegrationVersionIntegrationKit {
	this := IntegrationVersionIntegrationKit{}
	this.Name = name
	this.Number = number
	this.ReleasedOn = releasedOn
	return &this
}

// NewIntegrationVersionIntegrationKitWithDefaults instantiates a new IntegrationVersionIntegrationKit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationVersionIntegrationKitWithDefaults() *IntegrationVersionIntegrationKit {
	this := IntegrationVersionIntegrationKit{}
	return &this
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *IntegrationVersionIntegrationKit) GetConfiguration() IntegrationVersionCommonConfiguration {
	if o == nil || IsNil(o.Configuration) {
		var ret IntegrationVersionCommonConfiguration
		return ret
	}
	return *o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetConfigurationOk() (*IntegrationVersionCommonConfiguration, bool) {
	if o == nil || IsNil(o.Configuration) {
		return nil, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *IntegrationVersionIntegrationKit) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given IntegrationVersionCommonConfiguration and assigns it to the Configuration field.
func (o *IntegrationVersionIntegrationKit) SetConfiguration(v IntegrationVersionCommonConfiguration) {
	o.Configuration = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IntegrationVersionIntegrationKit) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IntegrationVersionIntegrationKit) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IntegrationVersionIntegrationKit) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IntegrationVersionIntegrationKit) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IntegrationVersionIntegrationKit) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IntegrationVersionIntegrationKit) SetId(v string) {
	o.Id = &v
}

// GetIntegration returns the Integration field value if set, zero value otherwise.
func (o *IntegrationVersionIntegrationKit) GetIntegration() IntegrationVersionCommonIntegration {
	if o == nil || IsNil(o.Integration) {
		var ret IntegrationVersionCommonIntegration
		return ret
	}
	return *o.Integration
}

// GetIntegrationOk returns a tuple with the Integration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetIntegrationOk() (*IntegrationVersionCommonIntegration, bool) {
	if o == nil || IsNil(o.Integration) {
		return nil, false
	}
	return o.Integration, true
}

// HasIntegration returns a boolean if a field has been set.
func (o *IntegrationVersionIntegrationKit) HasIntegration() bool {
	if o != nil && !IsNil(o.Integration) {
		return true
	}

	return false
}

// SetIntegration gets a reference to the given IntegrationVersionCommonIntegration and assigns it to the Integration field.
func (o *IntegrationVersionIntegrationKit) SetIntegration(v IntegrationVersionCommonIntegration) {
	o.Integration = &v
}

// GetName returns the Name field value
func (o *IntegrationVersionIntegrationKit) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IntegrationVersionIntegrationKit) SetName(v string) {
	o.Name = v
}

// GetNumber returns the Number field value
func (o *IntegrationVersionIntegrationKit) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *IntegrationVersionIntegrationKit) SetNumber(v string) {
	o.Number = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *IntegrationVersionIntegrationKit) GetType() EnumIntegrationVersionType {
	if o == nil || IsNil(o.Type) {
		var ret EnumIntegrationVersionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetTypeOk() (*EnumIntegrationVersionType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *IntegrationVersionIntegrationKit) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given EnumIntegrationVersionType and assigns it to the Type field.
func (o *IntegrationVersionIntegrationKit) SetType(v EnumIntegrationVersionType) {
	o.Type = &v
}

// GetDocumentationUrl returns the DocumentationUrl field value if set, zero value otherwise.
func (o *IntegrationVersionIntegrationKit) GetDocumentationUrl() string {
	if o == nil || IsNil(o.DocumentationUrl) {
		var ret string
		return ret
	}
	return *o.DocumentationUrl
}

// GetDocumentationUrlOk returns a tuple with the DocumentationUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetDocumentationUrlOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentationUrl) {
		return nil, false
	}
	return o.DocumentationUrl, true
}

// HasDocumentationUrl returns a boolean if a field has been set.
func (o *IntegrationVersionIntegrationKit) HasDocumentationUrl() bool {
	if o != nil && !IsNil(o.DocumentationUrl) {
		return true
	}

	return false
}

// SetDocumentationUrl gets a reference to the given string and assigns it to the DocumentationUrl field.
func (o *IntegrationVersionIntegrationKit) SetDocumentationUrl(v string) {
	o.DocumentationUrl = &v
}

// GetEndOfLifeOn returns the EndOfLifeOn field value if set, zero value otherwise.
func (o *IntegrationVersionIntegrationKit) GetEndOfLifeOn() string {
	if o == nil || IsNil(o.EndOfLifeOn) {
		var ret string
		return ret
	}
	return *o.EndOfLifeOn
}

// GetEndOfLifeOnOk returns a tuple with the EndOfLifeOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetEndOfLifeOnOk() (*string, bool) {
	if o == nil || IsNil(o.EndOfLifeOn) {
		return nil, false
	}
	return o.EndOfLifeOn, true
}

// HasEndOfLifeOn returns a boolean if a field has been set.
func (o *IntegrationVersionIntegrationKit) HasEndOfLifeOn() bool {
	if o != nil && !IsNil(o.EndOfLifeOn) {
		return true
	}

	return false
}

// SetEndOfLifeOn gets a reference to the given string and assigns it to the EndOfLifeOn field.
func (o *IntegrationVersionIntegrationKit) SetEndOfLifeOn(v string) {
	o.EndOfLifeOn = &v
}

// GetIntegratedWith returns the IntegratedWith field value if set, zero value otherwise.
func (o *IntegrationVersionIntegrationKit) GetIntegratedWith() IntegrationVersionIntegrationKitAllOfIntegratedWith {
	if o == nil || IsNil(o.IntegratedWith) {
		var ret IntegrationVersionIntegrationKitAllOfIntegratedWith
		return ret
	}
	return *o.IntegratedWith
}

// GetIntegratedWithOk returns a tuple with the IntegratedWith field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetIntegratedWithOk() (*IntegrationVersionIntegrationKitAllOfIntegratedWith, bool) {
	if o == nil || IsNil(o.IntegratedWith) {
		return nil, false
	}
	return o.IntegratedWith, true
}

// HasIntegratedWith returns a boolean if a field has been set.
func (o *IntegrationVersionIntegrationKit) HasIntegratedWith() bool {
	if o != nil && !IsNil(o.IntegratedWith) {
		return true
	}

	return false
}

// SetIntegratedWith gets a reference to the given IntegrationVersionIntegrationKitAllOfIntegratedWith and assigns it to the IntegratedWith field.
func (o *IntegrationVersionIntegrationKit) SetIntegratedWith(v IntegrationVersionIntegrationKitAllOfIntegratedWith) {
	o.IntegratedWith = &v
}

// GetReleasedOn returns the ReleasedOn field value
func (o *IntegrationVersionIntegrationKit) GetReleasedOn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReleasedOn
}

// GetReleasedOnOk returns a tuple with the ReleasedOn field value
// and a boolean to check if the value has been set.
func (o *IntegrationVersionIntegrationKit) GetReleasedOnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReleasedOn, true
}

// SetReleasedOn sets field value
func (o *IntegrationVersionIntegrationKit) SetReleasedOn(v string) {
	o.ReleasedOn = v
}

func (o IntegrationVersionIntegrationKit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationVersionIntegrationKit) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.DocumentationUrl) {
		toSerialize["documentationUrl"] = o.DocumentationUrl
	}
	if !IsNil(o.EndOfLifeOn) {
		toSerialize["endOfLifeOn"] = o.EndOfLifeOn
	}
	if !IsNil(o.IntegratedWith) {
		toSerialize["integratedWith"] = o.IntegratedWith
	}
	toSerialize["releasedOn"] = o.ReleasedOn
	return toSerialize, nil
}

type NullableIntegrationVersionIntegrationKit struct {
	value *IntegrationVersionIntegrationKit
	isSet bool
}

func (v NullableIntegrationVersionIntegrationKit) Get() *IntegrationVersionIntegrationKit {
	return v.value
}

func (v *NullableIntegrationVersionIntegrationKit) Set(val *IntegrationVersionIntegrationKit) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationVersionIntegrationKit) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationVersionIntegrationKit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationVersionIntegrationKit(val *IntegrationVersionIntegrationKit) *NullableIntegrationVersionIntegrationKit {
	return &NullableIntegrationVersionIntegrationKit{value: val, isSet: true}
}

func (v NullableIntegrationVersionIntegrationKit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationVersionIntegrationKit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
