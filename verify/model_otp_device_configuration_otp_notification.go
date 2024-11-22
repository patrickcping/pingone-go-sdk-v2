/*
PingOne Platform API - PingOne Verify

The PingOne Platform API covering the PingOne Verify service

API version: 2023-07-20
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package verify

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OTPDeviceConfigurationOtpNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OTPDeviceConfigurationOtpNotification{}

// OTPDeviceConfigurationOtpNotification OTP notification template configuration.
type OTPDeviceConfigurationOtpNotification struct {
	TemplateName string `json:"templateName"`
	VariantName *string `json:"variantName,omitempty"`
}

type _OTPDeviceConfigurationOtpNotification OTPDeviceConfigurationOtpNotification

// NewOTPDeviceConfigurationOtpNotification instantiates a new OTPDeviceConfigurationOtpNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOTPDeviceConfigurationOtpNotification(templateName string) *OTPDeviceConfigurationOtpNotification {
	this := OTPDeviceConfigurationOtpNotification{}
	this.TemplateName = templateName
	return &this
}

// NewOTPDeviceConfigurationOtpNotificationWithDefaults instantiates a new OTPDeviceConfigurationOtpNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOTPDeviceConfigurationOtpNotificationWithDefaults() *OTPDeviceConfigurationOtpNotification {
	this := OTPDeviceConfigurationOtpNotification{}
	return &this
}

// GetTemplateName returns the TemplateName field value
func (o *OTPDeviceConfigurationOtpNotification) GetTemplateName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateName
}

// GetTemplateNameOk returns a tuple with the TemplateName field value
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfigurationOtpNotification) GetTemplateNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateName, true
}

// SetTemplateName sets field value
func (o *OTPDeviceConfigurationOtpNotification) SetTemplateName(v string) {
	o.TemplateName = v
}

// GetVariantName returns the VariantName field value if set, zero value otherwise.
func (o *OTPDeviceConfigurationOtpNotification) GetVariantName() string {
	if o == nil || IsNil(o.VariantName) {
		var ret string
		return ret
	}
	return *o.VariantName
}

// GetVariantNameOk returns a tuple with the VariantName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OTPDeviceConfigurationOtpNotification) GetVariantNameOk() (*string, bool) {
	if o == nil || IsNil(o.VariantName) {
		return nil, false
	}
	return o.VariantName, true
}

// HasVariantName returns a boolean if a field has been set.
func (o *OTPDeviceConfigurationOtpNotification) HasVariantName() bool {
	if o != nil && !IsNil(o.VariantName) {
		return true
	}

	return false
}

// SetVariantName gets a reference to the given string and assigns it to the VariantName field.
func (o *OTPDeviceConfigurationOtpNotification) SetVariantName(v string) {
	o.VariantName = &v
}

func (o OTPDeviceConfigurationOtpNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OTPDeviceConfigurationOtpNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["templateName"] = o.TemplateName
	if !IsNil(o.VariantName) {
		toSerialize["variantName"] = o.VariantName
	}
	return toSerialize, nil
}

func (o *OTPDeviceConfigurationOtpNotification) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"templateName",
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

	varOTPDeviceConfigurationOtpNotification := _OTPDeviceConfigurationOtpNotification{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	// decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOTPDeviceConfigurationOtpNotification)

	if err != nil {
		return err
	}

	*o = OTPDeviceConfigurationOtpNotification(varOTPDeviceConfigurationOtpNotification)

	return err
}

type NullableOTPDeviceConfigurationOtpNotification struct {
	value *OTPDeviceConfigurationOtpNotification
	isSet bool
}

func (v NullableOTPDeviceConfigurationOtpNotification) Get() *OTPDeviceConfigurationOtpNotification {
	return v.value
}

func (v *NullableOTPDeviceConfigurationOtpNotification) Set(val *OTPDeviceConfigurationOtpNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableOTPDeviceConfigurationOtpNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableOTPDeviceConfigurationOtpNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOTPDeviceConfigurationOtpNotification(val *OTPDeviceConfigurationOtpNotification) *NullableOTPDeviceConfigurationOtpNotification {
	return &NullableOTPDeviceConfigurationOtpNotification{value: val, isSet: true}
}

func (v NullableOTPDeviceConfigurationOtpNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOTPDeviceConfigurationOtpNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


