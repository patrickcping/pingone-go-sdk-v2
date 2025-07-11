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

// checks if the AgreementLanguage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgreementLanguage{}

// AgreementLanguage struct for AgreementLanguage
type AgreementLanguage struct {
	Links           *map[string]LinksHATEOASValue     `json:"_links,omitempty"`
	Agreement       *AgreementLanguageAgreement       `json:"agreement,omitempty"`
	CurrentRevision *AgreementLanguageCurrentRevision `json:"currentRevision,omitempty"`
	// A string that is used as the title of the agreement for the language presented to the user. This is a required property.
	DisplayName string `json:"displayName"`
	// A boolean that maps directly with a language being enabled or displayed for the environment within the platform. This is a required property.
	Enabled bool `json:"enabled"`
	// A string that specifies the language ID.
	Id *string `json:"id,omitempty"`
	// A string that specifies the tag for identifying the language resource associated with this agreement consent (for example, en-US). This is a required property. For more information about language tags, see Tags for Identifying Languages.
	Locale         string                           `json:"locale"`
	UserExperience *AgreementLanguageUserExperience `json:"userExperience,omitempty"`
}

// NewAgreementLanguage instantiates a new AgreementLanguage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementLanguage(displayName string, enabled bool, locale string) *AgreementLanguage {
	this := AgreementLanguage{}
	this.DisplayName = displayName
	this.Enabled = enabled
	this.Locale = locale
	return &this
}

// NewAgreementLanguageWithDefaults instantiates a new AgreementLanguage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementLanguageWithDefaults() *AgreementLanguage {
	this := AgreementLanguage{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AgreementLanguage) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguage) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AgreementLanguage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *AgreementLanguage) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetAgreement returns the Agreement field value if set, zero value otherwise.
func (o *AgreementLanguage) GetAgreement() AgreementLanguageAgreement {
	if o == nil || IsNil(o.Agreement) {
		var ret AgreementLanguageAgreement
		return ret
	}
	return *o.Agreement
}

// GetAgreementOk returns a tuple with the Agreement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguage) GetAgreementOk() (*AgreementLanguageAgreement, bool) {
	if o == nil || IsNil(o.Agreement) {
		return nil, false
	}
	return o.Agreement, true
}

// HasAgreement returns a boolean if a field has been set.
func (o *AgreementLanguage) HasAgreement() bool {
	if o != nil && !IsNil(o.Agreement) {
		return true
	}

	return false
}

// SetAgreement gets a reference to the given AgreementLanguageAgreement and assigns it to the Agreement field.
func (o *AgreementLanguage) SetAgreement(v AgreementLanguageAgreement) {
	o.Agreement = &v
}

// GetCurrentRevision returns the CurrentRevision field value if set, zero value otherwise.
func (o *AgreementLanguage) GetCurrentRevision() AgreementLanguageCurrentRevision {
	if o == nil || IsNil(o.CurrentRevision) {
		var ret AgreementLanguageCurrentRevision
		return ret
	}
	return *o.CurrentRevision
}

// GetCurrentRevisionOk returns a tuple with the CurrentRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguage) GetCurrentRevisionOk() (*AgreementLanguageCurrentRevision, bool) {
	if o == nil || IsNil(o.CurrentRevision) {
		return nil, false
	}
	return o.CurrentRevision, true
}

// HasCurrentRevision returns a boolean if a field has been set.
func (o *AgreementLanguage) HasCurrentRevision() bool {
	if o != nil && !IsNil(o.CurrentRevision) {
		return true
	}

	return false
}

// SetCurrentRevision gets a reference to the given AgreementLanguageCurrentRevision and assigns it to the CurrentRevision field.
func (o *AgreementLanguage) SetCurrentRevision(v AgreementLanguageCurrentRevision) {
	o.CurrentRevision = &v
}

// GetDisplayName returns the DisplayName field value
func (o *AgreementLanguage) GetDisplayName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
func (o *AgreementLanguage) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayName, true
}

// SetDisplayName sets field value
func (o *AgreementLanguage) SetDisplayName(v string) {
	o.DisplayName = v
}

// GetEnabled returns the Enabled field value
func (o *AgreementLanguage) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *AgreementLanguage) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *AgreementLanguage) SetEnabled(v bool) {
	o.Enabled = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgreementLanguage) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguage) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgreementLanguage) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgreementLanguage) SetId(v string) {
	o.Id = &v
}

// GetLocale returns the Locale field value
func (o *AgreementLanguage) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *AgreementLanguage) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *AgreementLanguage) SetLocale(v string) {
	o.Locale = v
}

// GetUserExperience returns the UserExperience field value if set, zero value otherwise.
func (o *AgreementLanguage) GetUserExperience() AgreementLanguageUserExperience {
	if o == nil || IsNil(o.UserExperience) {
		var ret AgreementLanguageUserExperience
		return ret
	}
	return *o.UserExperience
}

// GetUserExperienceOk returns a tuple with the UserExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementLanguage) GetUserExperienceOk() (*AgreementLanguageUserExperience, bool) {
	if o == nil || IsNil(o.UserExperience) {
		return nil, false
	}
	return o.UserExperience, true
}

// HasUserExperience returns a boolean if a field has been set.
func (o *AgreementLanguage) HasUserExperience() bool {
	if o != nil && !IsNil(o.UserExperience) {
		return true
	}

	return false
}

// SetUserExperience gets a reference to the given AgreementLanguageUserExperience and assigns it to the UserExperience field.
func (o *AgreementLanguage) SetUserExperience(v AgreementLanguageUserExperience) {
	o.UserExperience = &v
}

func (o AgreementLanguage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgreementLanguage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Agreement) {
		toSerialize["agreement"] = o.Agreement
	}
	if !IsNil(o.CurrentRevision) {
		toSerialize["currentRevision"] = o.CurrentRevision
	}
	toSerialize["displayName"] = o.DisplayName
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["locale"] = o.Locale
	if !IsNil(o.UserExperience) {
		toSerialize["userExperience"] = o.UserExperience
	}
	return toSerialize, nil
}

type NullableAgreementLanguage struct {
	value *AgreementLanguage
	isSet bool
}

func (v NullableAgreementLanguage) Get() *AgreementLanguage {
	return v.value
}

func (v *NullableAgreementLanguage) Set(val *AgreementLanguage) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementLanguage) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementLanguage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementLanguage(val *AgreementLanguage) *NullableAgreementLanguage {
	return &NullableAgreementLanguage{value: val, isSet: true}
}

func (v NullableAgreementLanguage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementLanguage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
