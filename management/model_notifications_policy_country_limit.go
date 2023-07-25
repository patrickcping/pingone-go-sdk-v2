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

// checks if the NotificationsPolicyCountryLimit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsPolicyCountryLimit{}

// NotificationsPolicyCountryLimit You can use the `countryLimit` object to limit the countries where you can send SMS and voice notifications.
type NotificationsPolicyCountryLimit struct {
	Type EnumNotificationsPolicyCountryLimitType `json:"type"`
	// The delivery methods that the defined limitation should be applied to. Content of the array can be `SMS`, `Voice`, or both. If the parameter is not provided, the default is `SMS` and `Voice`.
	DeliveryMethods []EnumNotificationsPolicyCountryLimitDeliveryMethod `json:"deliveryMethods,omitempty"`
	// The countries where the specified methods should be allowed or denied. Use the two-letter country codes from ISO 3166-1.
	Countries []string `json:"countries"`
}

// NewNotificationsPolicyCountryLimit instantiates a new NotificationsPolicyCountryLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsPolicyCountryLimit(type_ EnumNotificationsPolicyCountryLimitType, countries []string) *NotificationsPolicyCountryLimit {
	this := NotificationsPolicyCountryLimit{}
	this.Type = type_
	this.Countries = countries
	return &this
}

// NewNotificationsPolicyCountryLimitWithDefaults instantiates a new NotificationsPolicyCountryLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsPolicyCountryLimitWithDefaults() *NotificationsPolicyCountryLimit {
	this := NotificationsPolicyCountryLimit{}
	return &this
}

// GetType returns the Type field value
func (o *NotificationsPolicyCountryLimit) GetType() EnumNotificationsPolicyCountryLimitType {
	if o == nil {
		var ret EnumNotificationsPolicyCountryLimitType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationsPolicyCountryLimit) GetTypeOk() (*EnumNotificationsPolicyCountryLimitType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationsPolicyCountryLimit) SetType(v EnumNotificationsPolicyCountryLimitType) {
	o.Type = v
}

// GetDeliveryMethods returns the DeliveryMethods field value if set, zero value otherwise.
func (o *NotificationsPolicyCountryLimit) GetDeliveryMethods() []EnumNotificationsPolicyCountryLimitDeliveryMethod {
	if o == nil || IsNil(o.DeliveryMethods) {
		var ret []EnumNotificationsPolicyCountryLimitDeliveryMethod
		return ret
	}
	return o.DeliveryMethods
}

// GetDeliveryMethodsOk returns a tuple with the DeliveryMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsPolicyCountryLimit) GetDeliveryMethodsOk() ([]EnumNotificationsPolicyCountryLimitDeliveryMethod, bool) {
	if o == nil || IsNil(o.DeliveryMethods) {
		return nil, false
	}
	return o.DeliveryMethods, true
}

// HasDeliveryMethods returns a boolean if a field has been set.
func (o *NotificationsPolicyCountryLimit) HasDeliveryMethods() bool {
	if o != nil && !IsNil(o.DeliveryMethods) {
		return true
	}

	return false
}

// SetDeliveryMethods gets a reference to the given []EnumNotificationsPolicyCountryLimitDeliveryMethod and assigns it to the DeliveryMethods field.
func (o *NotificationsPolicyCountryLimit) SetDeliveryMethods(v []EnumNotificationsPolicyCountryLimitDeliveryMethod) {
	o.DeliveryMethods = v
}

// GetCountries returns the Countries field value
func (o *NotificationsPolicyCountryLimit) GetCountries() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Countries
}

// GetCountriesOk returns a tuple with the Countries field value
// and a boolean to check if the value has been set.
func (o *NotificationsPolicyCountryLimit) GetCountriesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Countries, true
}

// SetCountries sets field value
func (o *NotificationsPolicyCountryLimit) SetCountries(v []string) {
	o.Countries = v
}

func (o NotificationsPolicyCountryLimit) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsPolicyCountryLimit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.DeliveryMethods) {
		toSerialize["deliveryMethods"] = o.DeliveryMethods
	}
	toSerialize["countries"] = o.Countries
	return toSerialize, nil
}

type NullableNotificationsPolicyCountryLimit struct {
	value *NotificationsPolicyCountryLimit
	isSet bool
}

func (v NullableNotificationsPolicyCountryLimit) Get() *NotificationsPolicyCountryLimit {
	return v.value
}

func (v *NullableNotificationsPolicyCountryLimit) Set(val *NotificationsPolicyCountryLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsPolicyCountryLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsPolicyCountryLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsPolicyCountryLimit(val *NotificationsPolicyCountryLimit) *NullableNotificationsPolicyCountryLimit {
	return &NullableNotificationsPolicyCountryLimit{value: val, isSet: true}
}

func (v NullableNotificationsPolicyCountryLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsPolicyCountryLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


