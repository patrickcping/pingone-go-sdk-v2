/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"time"
)

// checks if the RiskPredictorCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCommon{}

// RiskPredictorCommon struct for RiskPredictorCommon
type RiskPredictorCommon struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id          *string            `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights.
	Name string `json:"name"`
	// A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details).
	CompactName string            `json:"compactName"`
	Type        EnumPredictorType `json:"type"`
	// A string type. This specifies the description of the risk predictor. Maximum length is 1024 characters.
	Description *string `json:"description,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Indicates whether PingOne Risk is licensed for the environment.
	Licensed *bool `json:"licensed,omitempty"`
	// A boolean to indicate whether the predictor is deletable in the environment.
	Deletable *bool                         `json:"deletable,omitempty"`
	Default   *RiskPredictorCommonDefault   `json:"default,omitempty"`
	Condition *RiskPredictorCommonCondition `json:"condition,omitempty"`
}

// NewRiskPredictorCommon instantiates a new RiskPredictorCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCommon(name string, compactName string, type_ EnumPredictorType) *RiskPredictorCommon {
	this := RiskPredictorCommon{}
	this.Name = name
	this.CompactName = compactName
	this.Type = type_
	return &this
}

// NewRiskPredictorCommonWithDefaults instantiates a new RiskPredictorCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCommonWithDefaults() *RiskPredictorCommon {
	this := RiskPredictorCommon{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *RiskPredictorCommon) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskPredictorCommon) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *RiskPredictorCommon) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetName returns the Name field value
func (o *RiskPredictorCommon) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskPredictorCommon) SetName(v string) {
	o.Name = v
}

// GetCompactName returns the CompactName field value
func (o *RiskPredictorCommon) GetCompactName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompactName
}

// GetCompactNameOk returns a tuple with the CompactName field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetCompactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompactName, true
}

// SetCompactName sets field value
func (o *RiskPredictorCommon) SetCompactName(v string) {
	o.CompactName = v
}

// GetType returns the Type field value
func (o *RiskPredictorCommon) GetType() EnumPredictorType {
	if o == nil {
		var ret EnumPredictorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetTypeOk() (*EnumPredictorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RiskPredictorCommon) SetType(v EnumPredictorType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskPredictorCommon) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RiskPredictorCommon) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RiskPredictorCommon) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetLicensed() bool {
	if o == nil || IsNil(o.Licensed) {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetLicensedOk() (*bool, bool) {
	if o == nil || IsNil(o.Licensed) {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasLicensed() bool {
	if o != nil && !IsNil(o.Licensed) {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *RiskPredictorCommon) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetDeletable returns the Deletable field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetDeletable() bool {
	if o == nil || IsNil(o.Deletable) {
		var ret bool
		return ret
	}
	return *o.Deletable
}

// GetDeletableOk returns a tuple with the Deletable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetDeletableOk() (*bool, bool) {
	if o == nil || IsNil(o.Deletable) {
		return nil, false
	}
	return o.Deletable, true
}

// HasDeletable returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasDeletable() bool {
	if o != nil && !IsNil(o.Deletable) {
		return true
	}

	return false
}

// SetDeletable gets a reference to the given bool and assigns it to the Deletable field.
func (o *RiskPredictorCommon) SetDeletable(v bool) {
	o.Deletable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetDefault() RiskPredictorCommonDefault {
	if o == nil || IsNil(o.Default) {
		var ret RiskPredictorCommonDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetDefaultOk() (*RiskPredictorCommonDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given RiskPredictorCommonDefault and assigns it to the Default field.
func (o *RiskPredictorCommon) SetDefault(v RiskPredictorCommonDefault) {
	o.Default = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *RiskPredictorCommon) GetCondition() RiskPredictorCommonCondition {
	if o == nil || IsNil(o.Condition) {
		var ret RiskPredictorCommonCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCommon) GetConditionOk() (*RiskPredictorCommonCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *RiskPredictorCommon) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given RiskPredictorCommonCondition and assigns it to the Condition field.
func (o *RiskPredictorCommon) SetCondition(v RiskPredictorCommonCondition) {
	o.Condition = &v
}

func (o RiskPredictorCommon) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["name"] = o.Name
	toSerialize["compactName"] = o.CompactName
	toSerialize["type"] = o.Type
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Licensed) {
		toSerialize["licensed"] = o.Licensed
	}
	if !IsNil(o.Deletable) {
		toSerialize["deletable"] = o.Deletable
	}
	if !IsNil(o.Default) {
		toSerialize["default"] = o.Default
	}
	if !IsNil(o.Condition) {
		toSerialize["condition"] = o.Condition
	}
	return toSerialize, nil
}

type NullableRiskPredictorCommon struct {
	value *RiskPredictorCommon
	isSet bool
}

func (v NullableRiskPredictorCommon) Get() *RiskPredictorCommon {
	return v.value
}

func (v *NullableRiskPredictorCommon) Set(val *RiskPredictorCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCommon(val *RiskPredictorCommon) *NullableRiskPredictorCommon {
	return &NullableRiskPredictorCommon{value: val, isSet: true}
}

func (v NullableRiskPredictorCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
