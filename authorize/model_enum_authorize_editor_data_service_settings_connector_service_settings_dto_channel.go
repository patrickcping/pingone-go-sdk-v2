/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
	"fmt"
)

// EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel the model 'EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel'
type EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel string

// List of EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel
const (
	ENUMAUTHORIZEEDITORDATASERVICESETTINGSCONNECTORSERVICESETTINGSDTOCHANNEL_AUTHORIZE EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel = "AUTHORIZE"
)

// All allowed values of EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel enum
var AllowedEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannelEnumValues = []EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel{
	"AUTHORIZE",
}

func (v *EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel(value)
	for _, existing := range AllowedEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel(fmt.Sprintf("%s", "UNKNOWN"))
	return nil
}

// NewEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannelFromValue returns a pointer to a valid EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannelFromValue(v string) (*EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel, error) {
	ev := EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel: valid values are %v", v, AllowedEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) IsValid() bool {
	for _, existing := range AllowedEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel value
func (v EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) Ptr() *EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel {
	return &v
}

type NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel struct {
	value *EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel
	isSet bool
}

func (v NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) Get() *EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel {
	return v.value
}

func (v *NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) Set(val *EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) {
	v.value = val
	v.isSet = true
}

func (v NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) IsSet() bool {
	return v.isSet
}

func (v *NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel(val *EnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) *NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel {
	return &NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel{value: val, isSet: true}
}

func (v NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnumAuthorizeEditorDataServiceSettingsConnectorServiceSettingsDTOChannel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

