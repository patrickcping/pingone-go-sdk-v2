/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"encoding/json"
)

// checks if the CredentialIssuanceRuleNotificationTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CredentialIssuanceRuleNotificationTemplate{}

// CredentialIssuanceRuleNotificationTemplate struct for CredentialIssuanceRuleNotificationTemplate
type CredentialIssuanceRuleNotificationTemplate struct {
	// A string that specifies the ISO 2-character language code used for the notification; for example, en.
	Locale *string `json:"locale,omitempty"`
	// An object of key:value pairs that defines the dynamic variables used by the content variant.
	Variables map[string]interface{} `json:"variables,omitempty"`
	// A string that specifies the unique user-defined name for the content variant that contains the message text used for the notification
	Variant *string `json:"variant,omitempty"`
}

// NewCredentialIssuanceRuleNotificationTemplate instantiates a new CredentialIssuanceRuleNotificationTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCredentialIssuanceRuleNotificationTemplate() *CredentialIssuanceRuleNotificationTemplate {
	this := CredentialIssuanceRuleNotificationTemplate{}
	return &this
}

// NewCredentialIssuanceRuleNotificationTemplateWithDefaults instantiates a new CredentialIssuanceRuleNotificationTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCredentialIssuanceRuleNotificationTemplateWithDefaults() *CredentialIssuanceRuleNotificationTemplate {
	this := CredentialIssuanceRuleNotificationTemplate{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleNotificationTemplate) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleNotificationTemplate) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleNotificationTemplate) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *CredentialIssuanceRuleNotificationTemplate) SetLocale(v string) {
	o.Locale = &v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleNotificationTemplate) GetVariables() map[string]interface{} {
	if o == nil || IsNil(o.Variables) {
		var ret map[string]interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleNotificationTemplate) GetVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Variables) {
		return map[string]interface{}{}, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleNotificationTemplate) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *CredentialIssuanceRuleNotificationTemplate) SetVariables(v map[string]interface{}) {
	o.Variables = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *CredentialIssuanceRuleNotificationTemplate) GetVariant() string {
	if o == nil || IsNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CredentialIssuanceRuleNotificationTemplate) GetVariantOk() (*string, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *CredentialIssuanceRuleNotificationTemplate) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *CredentialIssuanceRuleNotificationTemplate) SetVariant(v string) {
	o.Variant = &v
}

func (o CredentialIssuanceRuleNotificationTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CredentialIssuanceRuleNotificationTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	return toSerialize, nil
}

type NullableCredentialIssuanceRuleNotificationTemplate struct {
	value *CredentialIssuanceRuleNotificationTemplate
	isSet bool
}

func (v NullableCredentialIssuanceRuleNotificationTemplate) Get() *CredentialIssuanceRuleNotificationTemplate {
	return v.value
}

func (v *NullableCredentialIssuanceRuleNotificationTemplate) Set(val *CredentialIssuanceRuleNotificationTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableCredentialIssuanceRuleNotificationTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableCredentialIssuanceRuleNotificationTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCredentialIssuanceRuleNotificationTemplate(val *CredentialIssuanceRuleNotificationTemplate) *NullableCredentialIssuanceRuleNotificationTemplate {
	return &NullableCredentialIssuanceRuleNotificationTemplate{value: val, isSet: true}
}

func (v NullableCredentialIssuanceRuleNotificationTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCredentialIssuanceRuleNotificationTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
