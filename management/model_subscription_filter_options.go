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

// checks if the SubscriptionFilterOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionFilterOptions{}

// SubscriptionFilterOptions struct for SubscriptionFilterOptions
type SubscriptionFilterOptions struct {
	// A non-empty array that specifies the list of action types that should be matched for the subscription. This is a required property.
	IncludedActionTypes []string `json:"includedActionTypes"`
	// An array that specifies the list of applications (by ID) whose events are monitored by the subscription (maximum of 10 IDs in the array). This is an optional property. If a list of applications is not provided, events are monitored for all applications in the environment.
	IncludedApplications []SubscriptionFilterOptionsIncludedApplicationsInner `json:"includedApplications,omitempty"`
	// An array that specifies the list of populations (by ID) whose events are monitored by the subscription (maximum of 10 IDs in the array). This property matches events for users in the specified populations, as opposed to events generated in which the user in one of the populations is the actor. This is an optional property.
	IncludedPopulations []SubscriptionFilterOptionsIncludedApplicationsInner `json:"includedPopulations,omitempty"`
	// An array of tags that events must have to be monitored by the subscription. If tags are not specified, all events are monitored. Currently, the available tags are `adminIdentityEvent`. Identifies the event as the action of an administrator on other administrators.
	IncludedTags []EnumSubscriptionFilterIncludedTags `json:"includedTags,omitempty"`
	// Whether the IP address of an actor should be present in the `source` section of the event.
	IpAddressExposed *bool `json:"ipAddressExposed,omitempty"`
	// Whether the User-Agent HTTP header of an event should be present in the `source` section of the event.
	UserAgentExposed *bool `json:"userAgentExposed,omitempty"`
}

// NewSubscriptionFilterOptions instantiates a new SubscriptionFilterOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionFilterOptions(includedActionTypes []string) *SubscriptionFilterOptions {
	this := SubscriptionFilterOptions{}
	this.IncludedActionTypes = includedActionTypes
	var ipAddressExposed bool = false
	this.IpAddressExposed = &ipAddressExposed
	var userAgentExposed bool = false
	this.UserAgentExposed = &userAgentExposed
	return &this
}

// NewSubscriptionFilterOptionsWithDefaults instantiates a new SubscriptionFilterOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionFilterOptionsWithDefaults() *SubscriptionFilterOptions {
	this := SubscriptionFilterOptions{}
	var ipAddressExposed bool = false
	this.IpAddressExposed = &ipAddressExposed
	var userAgentExposed bool = false
	this.UserAgentExposed = &userAgentExposed
	return &this
}

// GetIncludedActionTypes returns the IncludedActionTypes field value
func (o *SubscriptionFilterOptions) GetIncludedActionTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.IncludedActionTypes
}

// GetIncludedActionTypesOk returns a tuple with the IncludedActionTypes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionFilterOptions) GetIncludedActionTypesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IncludedActionTypes, true
}

// SetIncludedActionTypes sets field value
func (o *SubscriptionFilterOptions) SetIncludedActionTypes(v []string) {
	o.IncludedActionTypes = v
}

// GetIncludedApplications returns the IncludedApplications field value if set, zero value otherwise.
func (o *SubscriptionFilterOptions) GetIncludedApplications() []SubscriptionFilterOptionsIncludedApplicationsInner {
	if o == nil || IsNil(o.IncludedApplications) {
		var ret []SubscriptionFilterOptionsIncludedApplicationsInner
		return ret
	}
	return o.IncludedApplications
}

// GetIncludedApplicationsOk returns a tuple with the IncludedApplications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFilterOptions) GetIncludedApplicationsOk() ([]SubscriptionFilterOptionsIncludedApplicationsInner, bool) {
	if o == nil || IsNil(o.IncludedApplications) {
		return nil, false
	}
	return o.IncludedApplications, true
}

// HasIncludedApplications returns a boolean if a field has been set.
func (o *SubscriptionFilterOptions) HasIncludedApplications() bool {
	if o != nil && !IsNil(o.IncludedApplications) {
		return true
	}

	return false
}

// SetIncludedApplications gets a reference to the given []SubscriptionFilterOptionsIncludedApplicationsInner and assigns it to the IncludedApplications field.
func (o *SubscriptionFilterOptions) SetIncludedApplications(v []SubscriptionFilterOptionsIncludedApplicationsInner) {
	o.IncludedApplications = v
}

// GetIncludedPopulations returns the IncludedPopulations field value if set, zero value otherwise.
func (o *SubscriptionFilterOptions) GetIncludedPopulations() []SubscriptionFilterOptionsIncludedApplicationsInner {
	if o == nil || IsNil(o.IncludedPopulations) {
		var ret []SubscriptionFilterOptionsIncludedApplicationsInner
		return ret
	}
	return o.IncludedPopulations
}

// GetIncludedPopulationsOk returns a tuple with the IncludedPopulations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFilterOptions) GetIncludedPopulationsOk() ([]SubscriptionFilterOptionsIncludedApplicationsInner, bool) {
	if o == nil || IsNil(o.IncludedPopulations) {
		return nil, false
	}
	return o.IncludedPopulations, true
}

// HasIncludedPopulations returns a boolean if a field has been set.
func (o *SubscriptionFilterOptions) HasIncludedPopulations() bool {
	if o != nil && !IsNil(o.IncludedPopulations) {
		return true
	}

	return false
}

// SetIncludedPopulations gets a reference to the given []SubscriptionFilterOptionsIncludedApplicationsInner and assigns it to the IncludedPopulations field.
func (o *SubscriptionFilterOptions) SetIncludedPopulations(v []SubscriptionFilterOptionsIncludedApplicationsInner) {
	o.IncludedPopulations = v
}

// GetIncludedTags returns the IncludedTags field value if set, zero value otherwise.
func (o *SubscriptionFilterOptions) GetIncludedTags() []EnumSubscriptionFilterIncludedTags {
	if o == nil || IsNil(o.IncludedTags) {
		var ret []EnumSubscriptionFilterIncludedTags
		return ret
	}
	return o.IncludedTags
}

// GetIncludedTagsOk returns a tuple with the IncludedTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFilterOptions) GetIncludedTagsOk() ([]EnumSubscriptionFilterIncludedTags, bool) {
	if o == nil || IsNil(o.IncludedTags) {
		return nil, false
	}
	return o.IncludedTags, true
}

// HasIncludedTags returns a boolean if a field has been set.
func (o *SubscriptionFilterOptions) HasIncludedTags() bool {
	if o != nil && !IsNil(o.IncludedTags) {
		return true
	}

	return false
}

// SetIncludedTags gets a reference to the given []EnumSubscriptionFilterIncludedTags and assigns it to the IncludedTags field.
func (o *SubscriptionFilterOptions) SetIncludedTags(v []EnumSubscriptionFilterIncludedTags) {
	o.IncludedTags = v
}

// GetIpAddressExposed returns the IpAddressExposed field value if set, zero value otherwise.
func (o *SubscriptionFilterOptions) GetIpAddressExposed() bool {
	if o == nil || IsNil(o.IpAddressExposed) {
		var ret bool
		return ret
	}
	return *o.IpAddressExposed
}

// GetIpAddressExposedOk returns a tuple with the IpAddressExposed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFilterOptions) GetIpAddressExposedOk() (*bool, bool) {
	if o == nil || IsNil(o.IpAddressExposed) {
		return nil, false
	}
	return o.IpAddressExposed, true
}

// HasIpAddressExposed returns a boolean if a field has been set.
func (o *SubscriptionFilterOptions) HasIpAddressExposed() bool {
	if o != nil && !IsNil(o.IpAddressExposed) {
		return true
	}

	return false
}

// SetIpAddressExposed gets a reference to the given bool and assigns it to the IpAddressExposed field.
func (o *SubscriptionFilterOptions) SetIpAddressExposed(v bool) {
	o.IpAddressExposed = &v
}

// GetUserAgentExposed returns the UserAgentExposed field value if set, zero value otherwise.
func (o *SubscriptionFilterOptions) GetUserAgentExposed() bool {
	if o == nil || IsNil(o.UserAgentExposed) {
		var ret bool
		return ret
	}
	return *o.UserAgentExposed
}

// GetUserAgentExposedOk returns a tuple with the UserAgentExposed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionFilterOptions) GetUserAgentExposedOk() (*bool, bool) {
	if o == nil || IsNil(o.UserAgentExposed) {
		return nil, false
	}
	return o.UserAgentExposed, true
}

// HasUserAgentExposed returns a boolean if a field has been set.
func (o *SubscriptionFilterOptions) HasUserAgentExposed() bool {
	if o != nil && !IsNil(o.UserAgentExposed) {
		return true
	}

	return false
}

// SetUserAgentExposed gets a reference to the given bool and assigns it to the UserAgentExposed field.
func (o *SubscriptionFilterOptions) SetUserAgentExposed(v bool) {
	o.UserAgentExposed = &v
}

func (o SubscriptionFilterOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionFilterOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["includedActionTypes"] = o.IncludedActionTypes
	if !IsNil(o.IncludedApplications) {
		toSerialize["includedApplications"] = o.IncludedApplications
	}
	if !IsNil(o.IncludedPopulations) {
		toSerialize["includedPopulations"] = o.IncludedPopulations
	}
	if !IsNil(o.IncludedTags) {
		toSerialize["includedTags"] = o.IncludedTags
	}
	if !IsNil(o.IpAddressExposed) {
		toSerialize["ipAddressExposed"] = o.IpAddressExposed
	}
	if !IsNil(o.UserAgentExposed) {
		toSerialize["userAgentExposed"] = o.UserAgentExposed
	}
	return toSerialize, nil
}

type NullableSubscriptionFilterOptions struct {
	value *SubscriptionFilterOptions
	isSet bool
}

func (v NullableSubscriptionFilterOptions) Get() *SubscriptionFilterOptions {
	return v.value
}

func (v *NullableSubscriptionFilterOptions) Set(val *SubscriptionFilterOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionFilterOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionFilterOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionFilterOptions(val *SubscriptionFilterOptions) *NullableSubscriptionFilterOptions {
	return &NullableSubscriptionFilterOptions{value: val, isSet: true}
}

func (v NullableSubscriptionFilterOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionFilterOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


