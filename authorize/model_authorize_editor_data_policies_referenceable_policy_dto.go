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

// checks if the AuthorizeEditorDataPoliciesReferenceablePolicyDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataPoliciesReferenceablePolicyDTO{}

// AuthorizeEditorDataPoliciesReferenceablePolicyDTO struct for AuthorizeEditorDataPoliciesReferenceablePolicyDTO
type AuthorizeEditorDataPoliciesReferenceablePolicyDTO struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// HAL embedded resources
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The resource's unique identifier
	Id string `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Statements []map[string]interface{} `json:"statements,omitempty"`
	Condition *AuthorizeEditorDataConditionDTO `json:"condition,omitempty"`
	CombiningAlgorithm AuthorizeEditorDataPoliciesCombiningAlgorithmDTO `json:"combiningAlgorithm"`
	Children []map[string]interface{} `json:"children,omitempty"`
	RepetitionSettings *AuthorizeEditorDataPoliciesRepetitionSettingsDTO `json:"repetitionSettings,omitempty"`
	ManagedEntity *AuthorizeEditorDataManagedEntityDTO `json:"managedEntity,omitempty"`
	Version string `json:"version"`
}

// NewAuthorizeEditorDataPoliciesReferenceablePolicyDTO instantiates a new AuthorizeEditorDataPoliciesReferenceablePolicyDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataPoliciesReferenceablePolicyDTO(id string, name string, combiningAlgorithm AuthorizeEditorDataPoliciesCombiningAlgorithmDTO, version string) *AuthorizeEditorDataPoliciesReferenceablePolicyDTO {
	this := AuthorizeEditorDataPoliciesReferenceablePolicyDTO{}
	this.Id = id
	this.Name = name
	this.CombiningAlgorithm = combiningAlgorithm
	this.Version = version
	return &this
}

// NewAuthorizeEditorDataPoliciesReferenceablePolicyDTOWithDefaults instantiates a new AuthorizeEditorDataPoliciesReferenceablePolicyDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataPoliciesReferenceablePolicyDTOWithDefaults() *AuthorizeEditorDataPoliciesReferenceablePolicyDTO {
	this := AuthorizeEditorDataPoliciesReferenceablePolicyDTO{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatements returns the Statements field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetStatements() []map[string]interface{} {
	if o == nil || IsNil(o.Statements) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Statements
}

// GetStatementsOk returns a tuple with the Statements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetStatementsOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Statements) {
		return nil, false
	}
	return o.Statements, true
}

// HasStatements returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasStatements() bool {
	if o != nil && !IsNil(o.Statements) {
		return true
	}

	return false
}

// SetStatements gets a reference to the given []map[string]interface{} and assigns it to the Statements field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetStatements(v []map[string]interface{}) {
	o.Statements = v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetCondition() AuthorizeEditorDataConditionDTO {
	if o == nil || IsNil(o.Condition) {
		var ret AuthorizeEditorDataConditionDTO
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetConditionOk() (*AuthorizeEditorDataConditionDTO, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given AuthorizeEditorDataConditionDTO and assigns it to the Condition field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetCondition(v AuthorizeEditorDataConditionDTO) {
	o.Condition = &v
}

// GetCombiningAlgorithm returns the CombiningAlgorithm field value
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetCombiningAlgorithm() AuthorizeEditorDataPoliciesCombiningAlgorithmDTO {
	if o == nil {
		var ret AuthorizeEditorDataPoliciesCombiningAlgorithmDTO
		return ret
	}

	return o.CombiningAlgorithm
}

// GetCombiningAlgorithmOk returns a tuple with the CombiningAlgorithm field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetCombiningAlgorithmOk() (*AuthorizeEditorDataPoliciesCombiningAlgorithmDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CombiningAlgorithm, true
}

// SetCombiningAlgorithm sets field value
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetCombiningAlgorithm(v AuthorizeEditorDataPoliciesCombiningAlgorithmDTO) {
	o.CombiningAlgorithm = v
}

// GetChildren returns the Children field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetChildren() []map[string]interface{} {
	if o == nil || IsNil(o.Children) {
		var ret []map[string]interface{}
		return ret
	}
	return o.Children
}

// GetChildrenOk returns a tuple with the Children field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetChildrenOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Children) {
		return nil, false
	}
	return o.Children, true
}

// HasChildren returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasChildren() bool {
	if o != nil && !IsNil(o.Children) {
		return true
	}

	return false
}

// SetChildren gets a reference to the given []map[string]interface{} and assigns it to the Children field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetChildren(v []map[string]interface{}) {
	o.Children = v
}

// GetRepetitionSettings returns the RepetitionSettings field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetRepetitionSettings() AuthorizeEditorDataPoliciesRepetitionSettingsDTO {
	if o == nil || IsNil(o.RepetitionSettings) {
		var ret AuthorizeEditorDataPoliciesRepetitionSettingsDTO
		return ret
	}
	return *o.RepetitionSettings
}

// GetRepetitionSettingsOk returns a tuple with the RepetitionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetRepetitionSettingsOk() (*AuthorizeEditorDataPoliciesRepetitionSettingsDTO, bool) {
	if o == nil || IsNil(o.RepetitionSettings) {
		return nil, false
	}
	return o.RepetitionSettings, true
}

// HasRepetitionSettings returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasRepetitionSettings() bool {
	if o != nil && !IsNil(o.RepetitionSettings) {
		return true
	}

	return false
}

// SetRepetitionSettings gets a reference to the given AuthorizeEditorDataPoliciesRepetitionSettingsDTO and assigns it to the RepetitionSettings field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetRepetitionSettings(v AuthorizeEditorDataPoliciesRepetitionSettingsDTO) {
	o.RepetitionSettings = &v
}

// GetManagedEntity returns the ManagedEntity field value if set, zero value otherwise.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetManagedEntity() AuthorizeEditorDataManagedEntityDTO {
	if o == nil || IsNil(o.ManagedEntity) {
		var ret AuthorizeEditorDataManagedEntityDTO
		return ret
	}
	return *o.ManagedEntity
}

// GetManagedEntityOk returns a tuple with the ManagedEntity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetManagedEntityOk() (*AuthorizeEditorDataManagedEntityDTO, bool) {
	if o == nil || IsNil(o.ManagedEntity) {
		return nil, false
	}
	return o.ManagedEntity, true
}

// HasManagedEntity returns a boolean if a field has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) HasManagedEntity() bool {
	if o != nil && !IsNil(o.ManagedEntity) {
		return true
	}

	return false
}

// SetManagedEntity gets a reference to the given AuthorizeEditorDataManagedEntityDTO and assigns it to the ManagedEntity field.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetManagedEntity(v AuthorizeEditorDataManagedEntityDTO) {
	o.ManagedEntity = &v
}

// GetVersion returns the Version field value
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) SetVersion(v string) {
	o.Version = v
}

func (o AuthorizeEditorDataPoliciesReferenceablePolicyDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataPoliciesReferenceablePolicyDTO) ToMap() (map[string]interface{}, error) {
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
	toSerialize["combiningAlgorithm"] = o.CombiningAlgorithm
	if !IsNil(o.Children) {
		toSerialize["children"] = o.Children
	}
	if !IsNil(o.RepetitionSettings) {
		toSerialize["repetitionSettings"] = o.RepetitionSettings
	}
	if !IsNil(o.ManagedEntity) {
		toSerialize["managedEntity"] = o.ManagedEntity
	}
	toSerialize["version"] = o.Version
	return toSerialize, nil
}

type NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO struct {
	value *AuthorizeEditorDataPoliciesReferenceablePolicyDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO) Get() *AuthorizeEditorDataPoliciesReferenceablePolicyDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO) Set(val *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO(val *AuthorizeEditorDataPoliciesReferenceablePolicyDTO) *NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO {
	return &NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataPoliciesReferenceablePolicyDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


