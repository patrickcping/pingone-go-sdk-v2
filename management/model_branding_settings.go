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

// checks if the BrandingSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BrandingSettings{}

// BrandingSettings struct for BrandingSettings
type BrandingSettings struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// Specifies the resource’s unique identifier.
	Id          *string            `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The company name associated with the specified environment.
	CompanyName *string               `json:"companyName,omitempty"`
	Logo        *BrandingSettingsLogo `json:"logo,omitempty"`
}

// NewBrandingSettings instantiates a new BrandingSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBrandingSettings() *BrandingSettings {
	this := BrandingSettings{}
	return &this
}

// NewBrandingSettingsWithDefaults instantiates a new BrandingSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBrandingSettingsWithDefaults() *BrandingSettings {
	this := BrandingSettings{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BrandingSettings) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingSettings) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BrandingSettings) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *BrandingSettings) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *BrandingSettings) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingSettings) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *BrandingSettings) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *BrandingSettings) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *BrandingSettings) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingSettings) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *BrandingSettings) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *BrandingSettings) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetCompanyName returns the CompanyName field value if set, zero value otherwise.
func (o *BrandingSettings) GetCompanyName() string {
	if o == nil || IsNil(o.CompanyName) {
		var ret string
		return ret
	}
	return *o.CompanyName
}

// GetCompanyNameOk returns a tuple with the CompanyName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingSettings) GetCompanyNameOk() (*string, bool) {
	if o == nil || IsNil(o.CompanyName) {
		return nil, false
	}
	return o.CompanyName, true
}

// HasCompanyName returns a boolean if a field has been set.
func (o *BrandingSettings) HasCompanyName() bool {
	if o != nil && !IsNil(o.CompanyName) {
		return true
	}

	return false
}

// SetCompanyName gets a reference to the given string and assigns it to the CompanyName field.
func (o *BrandingSettings) SetCompanyName(v string) {
	o.CompanyName = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *BrandingSettings) GetLogo() BrandingSettingsLogo {
	if o == nil || IsNil(o.Logo) {
		var ret BrandingSettingsLogo
		return ret
	}
	return *o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BrandingSettings) GetLogoOk() (*BrandingSettingsLogo, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *BrandingSettings) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given BrandingSettingsLogo and assigns it to the Logo field.
func (o *BrandingSettings) SetLogo(v BrandingSettingsLogo) {
	o.Logo = &v
}

func (o BrandingSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BrandingSettings) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.CompanyName) {
		toSerialize["companyName"] = o.CompanyName
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	return toSerialize, nil
}

type NullableBrandingSettings struct {
	value *BrandingSettings
	isSet bool
}

func (v NullableBrandingSettings) Get() *BrandingSettings {
	return v.value
}

func (v *NullableBrandingSettings) Set(val *BrandingSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandingSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandingSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandingSettings(val *BrandingSettings) *NullableBrandingSettings {
	return &NullableBrandingSettings{value: val, isSet: true}
}

func (v NullableBrandingSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandingSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
