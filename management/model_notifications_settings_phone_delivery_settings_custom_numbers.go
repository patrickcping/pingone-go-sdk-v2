/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
	"fmt"
)

// checks if the NotificationsSettingsPhoneDeliverySettingsCustomNumbers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NotificationsSettingsPhoneDeliverySettingsCustomNumbers{}

// NotificationsSettingsPhoneDeliverySettingsCustomNumbers struct for NotificationsSettingsPhoneDeliverySettingsCustomNumbers
type NotificationsSettingsPhoneDeliverySettingsCustomNumbers struct {
	// The phone number, toll-free number or short code.
	Number string `json:"number"`
	Type EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType `json:"type"`
	// Specifies whether the number is selected by the admin for sending messages.
	Selected *bool `json:"selected,omitempty"`
	// Specifies whether the number is currently available in the provider account.
	Available *bool `json:"available,omitempty"`
	// A collection of the phone delivery service capabilities.
	Capabilities []EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability `json:"capabilities,omitempty"`
	// Specifies the `number`'s supported countries for notification recipients, depending on the phone number type: `SHORT_CODE`: A collection containing a single 2-character ISO country code, for example, `US`, `GB`, `CA`. If the custom provider is of `type=CUSTOM_PROVIDER`, `supportedCountries` must not be empty or null. For other custom provider types, if `supportedCountries` is null (empty is not supported), the specified short code number can only be used to dispatch notifications to United States recipient numbers. `TOLL_FREE`: A collection of valid 2-character country ISO codes, for example, `US`, `GB`, `CA`. If the custom provider is of `type=CUSTOM_PROVIDER`, `supportedCountries` must not be empty or null. For other custom provider types, if `supportedCountries` is null (empty is not supported), the specified toll-free number can only be used to dispatch notifications to United States recipient numbers. `PHONE_NUMBER`: `supportedCountries` can not be specified. If an SMS template has an alphanumeric `sender` ID and also has short code, the `sender` ID will be used for destination countries that support both alphanumeric senders and short codes. For Unites States and Canada that don't support alphanumeric sender IDs, a short code will be used if both an alphanumeric sender and a short code are specified. 
	SupportedCountries []string `json:"supportedCountries,omitempty"`
}

type _NotificationsSettingsPhoneDeliverySettingsCustomNumbers NotificationsSettingsPhoneDeliverySettingsCustomNumbers

// NewNotificationsSettingsPhoneDeliverySettingsCustomNumbers instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNotificationsSettingsPhoneDeliverySettingsCustomNumbers(number string, type_ EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType) *NotificationsSettingsPhoneDeliverySettingsCustomNumbers {
	this := NotificationsSettingsPhoneDeliverySettingsCustomNumbers{}
	this.Number = number
	this.Type = type_
	return &this
}

// NewNotificationsSettingsPhoneDeliverySettingsCustomNumbersWithDefaults instantiates a new NotificationsSettingsPhoneDeliverySettingsCustomNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNotificationsSettingsPhoneDeliverySettingsCustomNumbersWithDefaults() *NotificationsSettingsPhoneDeliverySettingsCustomNumbers {
	this := NotificationsSettingsPhoneDeliverySettingsCustomNumbers{}
	return &this
}

// GetNumber returns the Number field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetNumber(v string) {
	o.Number = v
}

// GetType returns the Type field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetType() EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType {
	if o == nil {
		var ret EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetTypeOk() (*EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetType(v EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersType) {
	o.Type = v
}

// GetSelected returns the Selected field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetSelected() bool {
	if o == nil || IsNil(o.Selected) {
		var ret bool
		return ret
	}
	return *o.Selected
}

// GetSelectedOk returns a tuple with the Selected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetSelectedOk() (*bool, bool) {
	if o == nil || IsNil(o.Selected) {
		return nil, false
	}
	return o.Selected, true
}

// HasSelected returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) HasSelected() bool {
	if o != nil && !IsNil(o.Selected) {
		return true
	}

	return false
}

// SetSelected gets a reference to the given bool and assigns it to the Selected field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetSelected(v bool) {
	o.Selected = &v
}

// GetAvailable returns the Available field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetAvailable() bool {
	if o == nil || IsNil(o.Available) {
		var ret bool
		return ret
	}
	return *o.Available
}

// GetAvailableOk returns a tuple with the Available field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetAvailableOk() (*bool, bool) {
	if o == nil || IsNil(o.Available) {
		return nil, false
	}
	return o.Available, true
}

// HasAvailable returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) HasAvailable() bool {
	if o != nil && !IsNil(o.Available) {
		return true
	}

	return false
}

// SetAvailable gets a reference to the given bool and assigns it to the Available field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetAvailable(v bool) {
	o.Available = &v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetCapabilities() []EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability {
	if o == nil || IsNil(o.Capabilities) {
		var ret []EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetCapabilitiesOk() ([]EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability and assigns it to the Capabilities field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetCapabilities(v []EnumNotificationsSettingsPhoneDeliverySettingsCustomNumbersCapability) {
	o.Capabilities = v
}

// GetSupportedCountries returns the SupportedCountries field value if set, zero value otherwise.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetSupportedCountries() []string {
	if o == nil || IsNil(o.SupportedCountries) {
		var ret []string
		return ret
	}
	return o.SupportedCountries
}

// GetSupportedCountriesOk returns a tuple with the SupportedCountries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) GetSupportedCountriesOk() ([]string, bool) {
	if o == nil || IsNil(o.SupportedCountries) {
		return nil, false
	}
	return o.SupportedCountries, true
}

// HasSupportedCountries returns a boolean if a field has been set.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) HasSupportedCountries() bool {
	if o != nil && !IsNil(o.SupportedCountries) {
		return true
	}

	return false
}

// SetSupportedCountries gets a reference to the given []string and assigns it to the SupportedCountries field.
func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) SetSupportedCountries(v []string) {
	o.SupportedCountries = v
}

func (o NotificationsSettingsPhoneDeliverySettingsCustomNumbers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NotificationsSettingsPhoneDeliverySettingsCustomNumbers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["number"] = o.Number
	toSerialize["type"] = o.Type
	if !IsNil(o.Selected) {
		toSerialize["selected"] = o.Selected
	}
	if !IsNil(o.Available) {
		toSerialize["available"] = o.Available
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.SupportedCountries) {
		toSerialize["supportedCountries"] = o.SupportedCountries
	}
	return toSerialize, nil
}

func (o *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) UnmarshalJSON(bytes []byte) (err error) {
    // This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"number",
		"type",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varNotificationsSettingsPhoneDeliverySettingsCustomNumbers := _NotificationsSettingsPhoneDeliverySettingsCustomNumbers{}

	err = json.Unmarshal(bytes, &varNotificationsSettingsPhoneDeliverySettingsCustomNumbers)

	if err != nil {
		return err
	}

	*o = NotificationsSettingsPhoneDeliverySettingsCustomNumbers(varNotificationsSettingsPhoneDeliverySettingsCustomNumbers)

	return err
}

type NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers struct {
	value *NotificationsSettingsPhoneDeliverySettingsCustomNumbers
	isSet bool
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers) Get() *NotificationsSettingsPhoneDeliverySettingsCustomNumbers {
	return v.value
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers) Set(val *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers(val *NotificationsSettingsPhoneDeliverySettingsCustomNumbers) *NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers {
	return &NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers{value: val, isSet: true}
}

func (v NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationsSettingsPhoneDeliverySettingsCustomNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


