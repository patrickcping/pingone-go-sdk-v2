/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"time"
)

// checks if the Subscription type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subscription{}

// Subscription struct for Subscription
type Subscription struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// The time the key resource expires.The date and time at which the subscription resource was created (ISO 8601 format).
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// A boolean that specifies whether a created or updated subscription should be active or suspended. A suspended state (`\"enabled\":false`) accumulates all matched events, but these events are not delivered until the subscription becomes active again (`\"enabled\":true`). For suspended subscriptions, events accumulate for a maximum of two weeks. Events older than two weeks are deleted. Restarted subscriptions receive the saved events (up to two weeks from the restart date). This is a required property.
	Enabled       bool                      `json:"enabled"`
	Environment   *ObjectEnvironment        `json:"environment,omitempty"`
	FilterOptions SubscriptionFilterOptions `json:"filterOptions"`
	Format        EnumSubscriptionFormat    `json:"format"`
	// A string that specifies the user resource’s unique identifier.
	Id           *string                  `json:"id,omitempty"`
	HttpEndpoint SubscriptionHttpEndpoint `json:"httpEndpoint"`
	// A string that specifies the subscription name. This is a required property.
	Name                 string                            `json:"name"`
	TlsClientAuthKeyPair *SubscriptionTlsClientAuthKeyPair `json:"tlsClientAuthKeyPair,omitempty"`
	// The date and time at which the subscription resource was last updated (ISO 8601 format).
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A boolean that specifies whether a certificates should be verified. If this property's value is set to false, then all certificates are trusted. (Setting this property's value to false introduces a security risk.) This is a required property.
	VerifyTlsCertificates bool `json:"verifyTlsCertificates"`
}

// NewSubscription instantiates a new Subscription object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscription(enabled bool, filterOptions SubscriptionFilterOptions, format EnumSubscriptionFormat, httpEndpoint SubscriptionHttpEndpoint, name string, verifyTlsCertificates bool) *Subscription {
	this := Subscription{}
	this.Enabled = enabled
	this.FilterOptions = filterOptions
	this.Format = format
	this.HttpEndpoint = httpEndpoint
	this.Name = name
	this.VerifyTlsCertificates = verifyTlsCertificates
	return &this
}

// NewSubscriptionWithDefaults instantiates a new Subscription object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionWithDefaults() *Subscription {
	this := Subscription{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Subscription) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Subscription) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *Subscription) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Subscription) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Subscription) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *Subscription) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetEnabled returns the Enabled field value
func (o *Subscription) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Subscription) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Subscription) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Subscription) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Subscription) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetFilterOptions returns the FilterOptions field value
func (o *Subscription) GetFilterOptions() SubscriptionFilterOptions {
	if o == nil {
		var ret SubscriptionFilterOptions
		return ret
	}

	return o.FilterOptions
}

// GetFilterOptionsOk returns a tuple with the FilterOptions field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetFilterOptionsOk() (*SubscriptionFilterOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FilterOptions, true
}

// SetFilterOptions sets field value
func (o *Subscription) SetFilterOptions(v SubscriptionFilterOptions) {
	o.FilterOptions = v
}

// GetFormat returns the Format field value
func (o *Subscription) GetFormat() EnumSubscriptionFormat {
	if o == nil {
		var ret EnumSubscriptionFormat
		return ret
	}

	return o.Format
}

// GetFormatOk returns a tuple with the Format field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetFormatOk() (*EnumSubscriptionFormat, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Format, true
}

// SetFormat sets field value
func (o *Subscription) SetFormat(v EnumSubscriptionFormat) {
	o.Format = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Subscription) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Subscription) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Subscription) SetId(v string) {
	o.Id = &v
}

// GetHttpEndpoint returns the HttpEndpoint field value
func (o *Subscription) GetHttpEndpoint() SubscriptionHttpEndpoint {
	if o == nil {
		var ret SubscriptionHttpEndpoint
		return ret
	}

	return o.HttpEndpoint
}

// GetHttpEndpointOk returns a tuple with the HttpEndpoint field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetHttpEndpointOk() (*SubscriptionHttpEndpoint, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpEndpoint, true
}

// SetHttpEndpoint sets field value
func (o *Subscription) SetHttpEndpoint(v SubscriptionHttpEndpoint) {
	o.HttpEndpoint = v
}

// GetName returns the Name field value
func (o *Subscription) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Subscription) SetName(v string) {
	o.Name = v
}

// GetTlsClientAuthKeyPair returns the TlsClientAuthKeyPair field value if set, zero value otherwise.
func (o *Subscription) GetTlsClientAuthKeyPair() SubscriptionTlsClientAuthKeyPair {
	if o == nil || IsNil(o.TlsClientAuthKeyPair) {
		var ret SubscriptionTlsClientAuthKeyPair
		return ret
	}
	return *o.TlsClientAuthKeyPair
}

// GetTlsClientAuthKeyPairOk returns a tuple with the TlsClientAuthKeyPair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetTlsClientAuthKeyPairOk() (*SubscriptionTlsClientAuthKeyPair, bool) {
	if o == nil || IsNil(o.TlsClientAuthKeyPair) {
		return nil, false
	}
	return o.TlsClientAuthKeyPair, true
}

// HasTlsClientAuthKeyPair returns a boolean if a field has been set.
func (o *Subscription) HasTlsClientAuthKeyPair() bool {
	if o != nil && !IsNil(o.TlsClientAuthKeyPair) {
		return true
	}

	return false
}

// SetTlsClientAuthKeyPair gets a reference to the given SubscriptionTlsClientAuthKeyPair and assigns it to the TlsClientAuthKeyPair field.
func (o *Subscription) SetTlsClientAuthKeyPair(v SubscriptionTlsClientAuthKeyPair) {
	o.TlsClientAuthKeyPair = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Subscription) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subscription) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Subscription) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Subscription) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVerifyTlsCertificates returns the VerifyTlsCertificates field value
func (o *Subscription) GetVerifyTlsCertificates() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.VerifyTlsCertificates
}

// GetVerifyTlsCertificatesOk returns a tuple with the VerifyTlsCertificates field value
// and a boolean to check if the value has been set.
func (o *Subscription) GetVerifyTlsCertificatesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyTlsCertificates, true
}

// SetVerifyTlsCertificates sets field value
func (o *Subscription) SetVerifyTlsCertificates(v bool) {
	o.VerifyTlsCertificates = v
}

func (o Subscription) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subscription) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	toSerialize["filterOptions"] = o.FilterOptions
	toSerialize["format"] = o.Format
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["httpEndpoint"] = o.HttpEndpoint
	toSerialize["name"] = o.Name
	if !IsNil(o.TlsClientAuthKeyPair) {
		toSerialize["tlsClientAuthKeyPair"] = o.TlsClientAuthKeyPair
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	toSerialize["verifyTlsCertificates"] = o.VerifyTlsCertificates
	return toSerialize, nil
}

type NullableSubscription struct {
	value *Subscription
	isSet bool
}

func (v NullableSubscription) Get() *Subscription {
	return v.value
}

func (v *NullableSubscription) Set(val *Subscription) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscription(val *Subscription) *NullableSubscription {
	return &NullableSubscription{value: val, isSet: true}
}

func (v NullableSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
