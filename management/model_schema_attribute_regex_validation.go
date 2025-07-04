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

// checks if the SchemaAttributeRegexValidation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SchemaAttributeRegexValidation{}

// SchemaAttributeRegexValidation Object representation of the optional regular expression representation of this attribute.
type SchemaAttributeRegexValidation struct {
	// A string that specifies the regular expression to which the attribute must conform.
	Pattern string `json:"pattern"`
	// A string that specifies a developer friendly description of the regular expression requirements.
	Requirements string `json:"requirements"`
	// An array that specifies the list of strings matching the regular expression.
	ValuesPatternShouldMatch []string `json:"valuesPatternShouldMatch,omitempty"`
	// An array that specifies the list of strings not matching the regular expression.
	ValuesPatternShouldNotMatch []string `json:"valuesPatternShouldNotMatch,omitempty"`
}

// NewSchemaAttributeRegexValidation instantiates a new SchemaAttributeRegexValidation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaAttributeRegexValidation(pattern string, requirements string) *SchemaAttributeRegexValidation {
	this := SchemaAttributeRegexValidation{}
	this.Pattern = pattern
	this.Requirements = requirements
	return &this
}

// NewSchemaAttributeRegexValidationWithDefaults instantiates a new SchemaAttributeRegexValidation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaAttributeRegexValidationWithDefaults() *SchemaAttributeRegexValidation {
	this := SchemaAttributeRegexValidation{}
	return &this
}

// GetPattern returns the Pattern field value
func (o *SchemaAttributeRegexValidation) GetPattern() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Pattern
}

// GetPatternOk returns a tuple with the Pattern field value
// and a boolean to check if the value has been set.
func (o *SchemaAttributeRegexValidation) GetPatternOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pattern, true
}

// SetPattern sets field value
func (o *SchemaAttributeRegexValidation) SetPattern(v string) {
	o.Pattern = v
}

// GetRequirements returns the Requirements field value
func (o *SchemaAttributeRegexValidation) GetRequirements() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value
// and a boolean to check if the value has been set.
func (o *SchemaAttributeRegexValidation) GetRequirementsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Requirements, true
}

// SetRequirements sets field value
func (o *SchemaAttributeRegexValidation) SetRequirements(v string) {
	o.Requirements = v
}

// GetValuesPatternShouldMatch returns the ValuesPatternShouldMatch field value if set, zero value otherwise.
func (o *SchemaAttributeRegexValidation) GetValuesPatternShouldMatch() []string {
	if o == nil || IsNil(o.ValuesPatternShouldMatch) {
		var ret []string
		return ret
	}
	return o.ValuesPatternShouldMatch
}

// GetValuesPatternShouldMatchOk returns a tuple with the ValuesPatternShouldMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttributeRegexValidation) GetValuesPatternShouldMatchOk() ([]string, bool) {
	if o == nil || IsNil(o.ValuesPatternShouldMatch) {
		return nil, false
	}
	return o.ValuesPatternShouldMatch, true
}

// HasValuesPatternShouldMatch returns a boolean if a field has been set.
func (o *SchemaAttributeRegexValidation) HasValuesPatternShouldMatch() bool {
	if o != nil && !IsNil(o.ValuesPatternShouldMatch) {
		return true
	}

	return false
}

// SetValuesPatternShouldMatch gets a reference to the given []string and assigns it to the ValuesPatternShouldMatch field.
func (o *SchemaAttributeRegexValidation) SetValuesPatternShouldMatch(v []string) {
	o.ValuesPatternShouldMatch = v
}

// GetValuesPatternShouldNotMatch returns the ValuesPatternShouldNotMatch field value if set, zero value otherwise.
func (o *SchemaAttributeRegexValidation) GetValuesPatternShouldNotMatch() []string {
	if o == nil || IsNil(o.ValuesPatternShouldNotMatch) {
		var ret []string
		return ret
	}
	return o.ValuesPatternShouldNotMatch
}

// GetValuesPatternShouldNotMatchOk returns a tuple with the ValuesPatternShouldNotMatch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttributeRegexValidation) GetValuesPatternShouldNotMatchOk() ([]string, bool) {
	if o == nil || IsNil(o.ValuesPatternShouldNotMatch) {
		return nil, false
	}
	return o.ValuesPatternShouldNotMatch, true
}

// HasValuesPatternShouldNotMatch returns a boolean if a field has been set.
func (o *SchemaAttributeRegexValidation) HasValuesPatternShouldNotMatch() bool {
	if o != nil && !IsNil(o.ValuesPatternShouldNotMatch) {
		return true
	}

	return false
}

// SetValuesPatternShouldNotMatch gets a reference to the given []string and assigns it to the ValuesPatternShouldNotMatch field.
func (o *SchemaAttributeRegexValidation) SetValuesPatternShouldNotMatch(v []string) {
	o.ValuesPatternShouldNotMatch = v
}

func (o SchemaAttributeRegexValidation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SchemaAttributeRegexValidation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["pattern"] = o.Pattern
	toSerialize["requirements"] = o.Requirements
	if !IsNil(o.ValuesPatternShouldMatch) {
		toSerialize["valuesPatternShouldMatch"] = o.ValuesPatternShouldMatch
	}
	if !IsNil(o.ValuesPatternShouldNotMatch) {
		toSerialize["valuesPatternShouldNotMatch"] = o.ValuesPatternShouldNotMatch
	}
	return toSerialize, nil
}

type NullableSchemaAttributeRegexValidation struct {
	value *SchemaAttributeRegexValidation
	isSet bool
}

func (v NullableSchemaAttributeRegexValidation) Get() *SchemaAttributeRegexValidation {
	return v.value
}

func (v *NullableSchemaAttributeRegexValidation) Set(val *SchemaAttributeRegexValidation) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaAttributeRegexValidation) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaAttributeRegexValidation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaAttributeRegexValidation(val *SchemaAttributeRegexValidation) *NullableSchemaAttributeRegexValidation {
	return &NullableSchemaAttributeRegexValidation{value: val, isSet: true}
}

func (v NullableSchemaAttributeRegexValidation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaAttributeRegexValidation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
