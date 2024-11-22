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
	"bytes"
	"fmt"
)

// checks if the RiskPredictorCustom type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskPredictorCustom{}

// RiskPredictorCustom struct for RiskPredictorCustom
type RiskPredictorCustom struct {
	Links map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string type. A unique, friendly name for the predictor. This name is displayed in the Risk Policies UI, when the admin is asked to define the overrides and weights.
	Name string `json:"name"`
	// A string type. A unique name for the predictor. This property is immutable; it cannot be modified after initial creation. The value must be alpha-numeric, with no special characters or spaces. This name is used in the API both for policy configuration, and in the Risk Evaluation response (under details).
	CompactName string `json:"compactName"`
	Type EnumPredictorType `json:"type"`
	// A string type. This specifies the description of the risk predictor. Maximum length is 1024 characters.
	Description *string `json:"description,omitempty"`
	// The time the resource was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// The time the resource was updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// Indicates whether PingOne Risk is licensed for the environment.
	Licensed *bool `json:"licensed,omitempty"`
	// A boolean to indicate whether the predictor is deletable in the environment.
	Deletable *bool `json:"deletable,omitempty"`
	Default *RiskPredictorCommonDefault `json:"default,omitempty"`
	Condition *RiskPredictorCommonCondition `json:"condition,omitempty"`
	Map RiskPredictorCustomAllOfMap `json:"map"`
}

type _RiskPredictorCustom RiskPredictorCustom

// NewRiskPredictorCustom instantiates a new RiskPredictorCustom object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskPredictorCustom(name string, compactName string, type_ EnumPredictorType, map_ RiskPredictorCustomAllOfMap) *RiskPredictorCustom {
	this := RiskPredictorCustom{}
	this.Name = name
	this.CompactName = compactName
	this.Type = type_
	this.Map = map_
	return &this
}

// NewRiskPredictorCustomWithDefaults instantiates a new RiskPredictorCustom object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskPredictorCustomWithDefaults() *RiskPredictorCustom {
	this := RiskPredictorCustom{}
	return &this
}

func (o RiskPredictorCustom) hasHalLink(linkIndex string) bool {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			if h, ok := v.GetHrefOk(); ok && h != nil && *h != "" {
				return true
			}
		}
	}
	return false
}

func (o RiskPredictorCustom) getHalLink(linkIndex string) LinksHATEOASValue {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return v
		}
	}

	var ret LinksHATEOASValue
	return ret
}

func (o RiskPredictorCustom) getHalLinkOk(linkIndex string) (*LinksHATEOASValue, bool) {
	if l, ok := o.GetLinksOk(); ok && l != nil {
		links := l
		if v, ok := links[linkIndex]; ok {
			return &v, true
		}
	}

	return nil, false
}

func (o RiskPredictorCustom) IsPaginated() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT) || o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskPredictorCustom) HasPaginationSelf() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskPredictorCustom) GetPaginationSelfLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskPredictorCustom) GetPaginationSelfLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_SELF)
}

func (o RiskPredictorCustom) HasPaginationNext() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskPredictorCustom) GetPaginationNextLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskPredictorCustom) GetPaginationNextLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_NEXT)
}

func (o RiskPredictorCustom) HasPaginationPrevious() bool {
	return o.hasHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskPredictorCustom) GetPaginationPreviousLink() LinksHATEOASValue {
	return o.getHalLink(PAGINATION_HAL_LINK_INDEX_PREV)
}

func (o RiskPredictorCustom) GetPaginationPreviousLinkOk() (*LinksHATEOASValue, bool) {
	return o.getHalLinkOk(PAGINATION_HAL_LINK_INDEX_PREV)
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetLinksOk() (map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]LinksHATEOASValue{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *RiskPredictorCustom) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RiskPredictorCustom) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *RiskPredictorCustom) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RiskPredictorCustom) SetName(v string) {
	o.Name = v
}

// GetCompactName returns the CompactName field value
func (o *RiskPredictorCustom) GetCompactName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CompactName
}

// GetCompactNameOk returns a tuple with the CompactName field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetCompactNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompactName, true
}

// SetCompactName sets field value
func (o *RiskPredictorCustom) SetCompactName(v string) {
	o.CompactName = v
}

// GetType returns the Type field value
func (o *RiskPredictorCustom) GetType() EnumPredictorType {
	if o == nil {
		var ret EnumPredictorType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetTypeOk() (*EnumPredictorType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RiskPredictorCustom) SetType(v EnumPredictorType) {
	o.Type = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RiskPredictorCustom) SetDescription(v string) {
	o.Description = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *RiskPredictorCustom) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *RiskPredictorCustom) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetLicensed returns the Licensed field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetLicensed() bool {
	if o == nil || IsNil(o.Licensed) {
		var ret bool
		return ret
	}
	return *o.Licensed
}

// GetLicensedOk returns a tuple with the Licensed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetLicensedOk() (*bool, bool) {
	if o == nil || IsNil(o.Licensed) {
		return nil, false
	}
	return o.Licensed, true
}

// HasLicensed returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasLicensed() bool {
	if o != nil && !IsNil(o.Licensed) {
		return true
	}

	return false
}

// SetLicensed gets a reference to the given bool and assigns it to the Licensed field.
func (o *RiskPredictorCustom) SetLicensed(v bool) {
	o.Licensed = &v
}

// GetDeletable returns the Deletable field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetDeletable() bool {
	if o == nil || IsNil(o.Deletable) {
		var ret bool
		return ret
	}
	return *o.Deletable
}

// GetDeletableOk returns a tuple with the Deletable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetDeletableOk() (*bool, bool) {
	if o == nil || IsNil(o.Deletable) {
		return nil, false
	}
	return o.Deletable, true
}

// HasDeletable returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasDeletable() bool {
	if o != nil && !IsNil(o.Deletable) {
		return true
	}

	return false
}

// SetDeletable gets a reference to the given bool and assigns it to the Deletable field.
func (o *RiskPredictorCustom) SetDeletable(v bool) {
	o.Deletable = &v
}

// GetDefault returns the Default field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetDefault() RiskPredictorCommonDefault {
	if o == nil || IsNil(o.Default) {
		var ret RiskPredictorCommonDefault
		return ret
	}
	return *o.Default
}

// GetDefaultOk returns a tuple with the Default field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetDefaultOk() (*RiskPredictorCommonDefault, bool) {
	if o == nil || IsNil(o.Default) {
		return nil, false
	}
	return o.Default, true
}

// HasDefault returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasDefault() bool {
	if o != nil && !IsNil(o.Default) {
		return true
	}

	return false
}

// SetDefault gets a reference to the given RiskPredictorCommonDefault and assigns it to the Default field.
func (o *RiskPredictorCustom) SetDefault(v RiskPredictorCommonDefault) {
	o.Default = &v
}

// GetCondition returns the Condition field value if set, zero value otherwise.
func (o *RiskPredictorCustom) GetCondition() RiskPredictorCommonCondition {
	if o == nil || IsNil(o.Condition) {
		var ret RiskPredictorCommonCondition
		return ret
	}
	return *o.Condition
}

// GetConditionOk returns a tuple with the Condition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetConditionOk() (*RiskPredictorCommonCondition, bool) {
	if o == nil || IsNil(o.Condition) {
		return nil, false
	}
	return o.Condition, true
}

// HasCondition returns a boolean if a field has been set.
func (o *RiskPredictorCustom) HasCondition() bool {
	if o != nil && !IsNil(o.Condition) {
		return true
	}

	return false
}

// SetCondition gets a reference to the given RiskPredictorCommonCondition and assigns it to the Condition field.
func (o *RiskPredictorCustom) SetCondition(v RiskPredictorCommonCondition) {
	o.Condition = &v
}

// GetMap returns the Map field value
func (o *RiskPredictorCustom) GetMap() RiskPredictorCustomAllOfMap {
	if o == nil {
		var ret RiskPredictorCustomAllOfMap
		return ret
	}

	return o.Map
}

// GetMapOk returns a tuple with the Map field value
// and a boolean to check if the value has been set.
func (o *RiskPredictorCustom) GetMapOk() (*RiskPredictorCustomAllOfMap, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Map, true
}

// SetMap sets field value
func (o *RiskPredictorCustom) SetMap(v RiskPredictorCustomAllOfMap) {
	o.Map = v
}

func (o RiskPredictorCustom) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskPredictorCustom) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
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
	toSerialize["map"] = o.Map
	return toSerialize, nil
}

func (o *RiskPredictorCustom) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"compactName",
		"type",
		"map",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRiskPredictorCustom := _RiskPredictorCustom{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRiskPredictorCustom)

	if err != nil {
		return err
	}

	*o = RiskPredictorCustom(varRiskPredictorCustom)

	return err
}

type NullableRiskPredictorCustom struct {
	value *RiskPredictorCustom
	isSet bool
}

func (v NullableRiskPredictorCustom) Get() *RiskPredictorCustom {
	return v.value
}

func (v *NullableRiskPredictorCustom) Set(val *RiskPredictorCustom) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictorCustom) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictorCustom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictorCustom(val *RiskPredictorCustom) *NullableRiskPredictorCustom {
	return &NullableRiskPredictorCustom{value: val, isSet: true}
}

func (v NullableRiskPredictorCustom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictorCustom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


