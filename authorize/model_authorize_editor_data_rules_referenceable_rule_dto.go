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

// checks if the AuthorizeEditorDataRulesReferenceableRuleDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataRulesReferenceableRuleDTO{}

// AuthorizeEditorDataRulesReferenceableRuleDTO struct for AuthorizeEditorDataRulesReferenceableRuleDTO
type AuthorizeEditorDataRulesReferenceableRuleDTO struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// HAL embedded resources
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The resource's unique identifier
	Id string `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Statements *[]map[string]interface{} `json:"statements,omitempty"`
	Condition *AuthorizeEditorDataConditionDTO `json:"condition,omitempty"`
	EffectSettings AuthorizeEditorDataRulesEffectSettingsDTO `json:"effectSettings"`
	Version string `json:"version"`
}

// NewAuthorizeEditorDataRulesReferenceableRuleDTO instantiates a new AuthorizeEditorDataRulesReferenceableRuleDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataRulesReferenceableRuleDTO(id string, name string, effectSettings AuthorizeEditorDataRulesEffectSettingsDTO, version string) *AuthorizeEditorDataRulesReferenceableRuleDTO {
	this := AuthorizeEditorDataRulesReferenceableRuleDTO{}
	this.Id = id
	this.Name = name
	this.EffectSettings = effectSettings
	this.Version = version
	return &this
}

// NewAuthorizeEditorDataRulesReferenceableRuleDTOWithDefaults instantiates a new AuthorizeEditorDataRulesReferenceableRuleDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataRulesReferenceableRuleDTOWithDefaults() *AuthorizeEditorDataRulesReferenceableRuleDTO {
	this := AuthorizeEditorDataRulesReferenceableRuleDTO{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatements returns the Statements field value if set, zero value otherwise.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetStatements() []map[string]interface{} {
	if o == nil || IsNil(o.Statements) {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetStatementsOk() (*[]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Statements) {
		return nil, false
	}
	return o.Statements, true
}

// HasStatements returns a boolean if a field has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasStatements() bool {
	if o != nil && !IsNil(o.Statements) {
		return true
	}

	return false
}

// SetStatements gets a reference to the given []map[string]interface{} and assigns it to the Statements field.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetStatements(v []map[string]interface{}) {
	o.Statements = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetCondition() AuthorizeEditorDataConditionDTO {
	if o == nil || IsNil(o.Condition) {
		var ret AuthorizeEditorDataConditionDTO
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given AuthorizeEditorDataConditionDTO and assigns it to the Condition field.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetCondition(v AuthorizeEditorDataConditionDTO) {
	o.Condition = &v
}

// GetEffectSettings returns the EffectSettings field value
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEffectSettings() AuthorizeEditorDataRulesEffectSettingsDTO {
	if o == nil {
		var ret AuthorizeEditorDataRulesEffectSettingsDTO
		return ret
	}

	return o.EffectSettings
}

// GetEffectSettingsOk returns a tuple with the EffectSettings field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetEffectSettingsOk() (*AuthorizeEditorDataRulesEffectSettingsDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EffectSettings, true
}

// SetEffectSettings sets field value
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetEffectSettings(v AuthorizeEditorDataRulesEffectSettingsDTO) {
	o.EffectSettings = v
}

// GetVersion returns the Version field value
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AuthorizeEditorDataRulesReferenceableRuleDTO) SetVersion(v string) {
	o.Version = v
}

func (o AuthorizeEditorDataRulesReferenceableRuleDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataRulesReferenceableRuleDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Statements) {
		toSerialize["statements"] = o.Statements
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	toSerialize["effectSettings"] = o.EffectSettings
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableAuthorizeEditorDataRulesReferenceableRuleDTO struct {
	value *AuthorizeEditorDataRulesReferenceableRuleDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataRulesReferenceableRuleDTO) Get() *AuthorizeEditorDataRulesReferenceableRuleDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataRulesReferenceableRuleDTO) Set(val *AuthorizeEditorDataRulesReferenceableRuleDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataRulesReferenceableRuleDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataRulesReferenceableRuleDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataRulesReferenceableRuleDTO(val *AuthorizeEditorDataRulesReferenceableRuleDTO) *NullableAuthorizeEditorDataRulesReferenceableRuleDTO {
	return &NullableAuthorizeEditorDataRulesReferenceableRuleDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataRulesReferenceableRuleDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataRulesReferenceableRuleDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

