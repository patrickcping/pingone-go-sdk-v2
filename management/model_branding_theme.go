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

// checks if the BrandingTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingTheme{}

// BrandingTheme struct for BrandingTheme
type BrandingTheme struct {
	Configuration BrandingThemeConfiguration `json:"configuration"`
	// Specifies whether this theme is the environment's default branding configuration.
	Default bool `json:"default"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// Specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	Template EnumBrandingThemeTemplate `json:"template"`
}

// NewBrandingTheme instantiates a new BrandingTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingTheme(configuration BrandingThemeConfiguration, default_ bool, template EnumBrandingThemeTemplate) *BrandingTheme {
	this := BrandingTheme{}
	this.Configuration = configuration
	this.Default = default_
	this.Template = template
	return &this
}

// NewBrandingThemeWithDefaults instantiates a new BrandingTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingThemeWithDefaults() *BrandingTheme {
	this := BrandingTheme{}
	return &this
}

// GetConfiguration returns the Configuration field value
func (o *BrandingTheme) GetConfiguration() BrandingThemeConfiguration {
	if o == nil {
		var ret BrandingThemeConfiguration
		return ret
	}

	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value
// and a boolean to check if the value has been set.
func (o *BrandingTheme) GetConfigurationOk() (*BrandingThemeConfiguration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Configuration, true
}

// SetConfiguration sets field value
func (o *BrandingTheme) SetConfiguration(v BrandingThemeConfiguration) {
	o.Configuration = v
}

// GetDefault returns the Default field value
func (o *BrandingTheme) GetDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Default
}

// GetDefaultOk returns a tuple with the Default field value
// and a boolean to check if the value has been set.
func (o *BrandingTheme) GetDefaultOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Default, true
}

// SetDefault sets field value
func (o *BrandingTheme) SetDefault(v bool) {
	o.Default = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *BrandingTheme) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingTheme) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *BrandingTheme) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *BrandingTheme) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrandingTheme) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingTheme) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrandingTheme) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BrandingTheme) SetId(v string) {
	o.Id = &v
}

// GetTemplate returns the Template field value
func (o *BrandingTheme) GetTemplate() EnumBrandingThemeTemplate {
	if o == nil {
		var ret EnumBrandingThemeTemplate
		return ret
	}

	return o.Template
}

// GetTemplateOk returns a tuple with the Template field value
// and a boolean to check if the value has been set.
func (o *BrandingTheme) GetTemplateOk() (*EnumBrandingThemeTemplate, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Template, true
}

// SetTemplate sets field value
func (o *BrandingTheme) SetTemplate(v EnumBrandingThemeTemplate) {
	o.Template = v
}

func (o BrandingTheme) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configuration"] = o.Configuration
	toSerialize["default"] = o.Default
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	// skip: id is readOnly
	toSerialize["template"] = o.Template
	return toSerialize, nil
}

type NullableBrandingTheme struct {
	value *BrandingTheme
	isSet bool
}

func (v NullableBrandingTheme) Get() *BrandingTheme {
	return v.value
}

func (v *NullableBrandingTheme) Set(val *BrandingTheme) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingTheme) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingTheme(val *BrandingTheme) *NullableBrandingTheme {
	return &NullableBrandingTheme{value: val, isSet: true}
}

func (v NullableBrandingTheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


