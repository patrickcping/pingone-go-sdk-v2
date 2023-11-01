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

// PropagationStoreConfiguration - struct for PropagationStoreConfiguration
type PropagationStoreConfiguration struct {
	PropagationStoreConfigurationAquera *PropagationStoreConfigurationAquera
	PropagationStoreConfigurationAzureADSAMLV2 *PropagationStoreConfigurationAzureADSAMLV2
	PropagationStoreConfigurationGithubEMU *PropagationStoreConfigurationGithubEMU
	PropagationStoreConfigurationGoogleApps *PropagationStoreConfigurationGoogleApps
	PropagationStoreConfigurationLDAPGateway *PropagationStoreConfigurationLDAPGateway
	PropagationStoreConfigurationPingOne *PropagationStoreConfigurationPingOne
	PropagationStoreConfigurationSCIM *PropagationStoreConfigurationSCIM
	PropagationStoreConfigurationSalesforce *PropagationStoreConfigurationSalesforce
	PropagationStoreConfigurationSalesforceContacts *PropagationStoreConfigurationSalesforceContacts
	PropagationStoreConfigurationServiceNow *PropagationStoreConfigurationServiceNow
	PropagationStoreConfigurationSlack *PropagationStoreConfigurationSlack
	PropagationStoreConfigurationWorkday *PropagationStoreConfigurationWorkday
	PropagationStoreConfigurationZoom *PropagationStoreConfigurationZoom
}

// PropagationStoreConfigurationAqueraAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationAquera wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationAqueraAsPropagationStoreConfiguration(v *PropagationStoreConfigurationAquera) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationAquera: v,
	}
}

// PropagationStoreConfigurationAzureADSAMLV2AsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationAzureADSAMLV2 wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationAzureADSAMLV2AsPropagationStoreConfiguration(v *PropagationStoreConfigurationAzureADSAMLV2) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationAzureADSAMLV2: v,
	}
}

// PropagationStoreConfigurationGithubEMUAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationGithubEMU wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationGithubEMUAsPropagationStoreConfiguration(v *PropagationStoreConfigurationGithubEMU) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationGithubEMU: v,
	}
}

// PropagationStoreConfigurationGoogleAppsAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationGoogleApps wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationGoogleAppsAsPropagationStoreConfiguration(v *PropagationStoreConfigurationGoogleApps) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationGoogleApps: v,
	}
}

// PropagationStoreConfigurationLDAPGatewayAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationLDAPGateway wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationLDAPGatewayAsPropagationStoreConfiguration(v *PropagationStoreConfigurationLDAPGateway) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationLDAPGateway: v,
	}
}

// PropagationStoreConfigurationPingOneAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationPingOne wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationPingOneAsPropagationStoreConfiguration(v *PropagationStoreConfigurationPingOne) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationPingOne: v,
	}
}

// PropagationStoreConfigurationSCIMAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationSCIM wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationSCIMAsPropagationStoreConfiguration(v *PropagationStoreConfigurationSCIM) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationSCIM: v,
	}
}

// PropagationStoreConfigurationSalesforceAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationSalesforce wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationSalesforceAsPropagationStoreConfiguration(v *PropagationStoreConfigurationSalesforce) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationSalesforce: v,
	}
}

// PropagationStoreConfigurationSalesforceContactsAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationSalesforceContacts wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationSalesforceContactsAsPropagationStoreConfiguration(v *PropagationStoreConfigurationSalesforceContacts) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationSalesforceContacts: v,
	}
}

// PropagationStoreConfigurationServiceNowAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationServiceNow wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationServiceNowAsPropagationStoreConfiguration(v *PropagationStoreConfigurationServiceNow) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationServiceNow: v,
	}
}

// PropagationStoreConfigurationSlackAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationSlack wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationSlackAsPropagationStoreConfiguration(v *PropagationStoreConfigurationSlack) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationSlack: v,
	}
}

// PropagationStoreConfigurationWorkdayAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationWorkday wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationWorkdayAsPropagationStoreConfiguration(v *PropagationStoreConfigurationWorkday) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationWorkday: v,
	}
}

// PropagationStoreConfigurationZoomAsPropagationStoreConfiguration is a convenience function that returns PropagationStoreConfigurationZoom wrapped in PropagationStoreConfiguration
func PropagationStoreConfigurationZoomAsPropagationStoreConfiguration(v *PropagationStoreConfigurationZoom) PropagationStoreConfiguration {
	return PropagationStoreConfiguration{
		PropagationStoreConfigurationZoom: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PropagationStoreConfiguration) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into PropagationStoreConfigurationAquera
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationAquera)
	if err == nil {
		jsonPropagationStoreConfigurationAquera, _ := json.Marshal(dst.PropagationStoreConfigurationAquera)
		if string(jsonPropagationStoreConfigurationAquera) == "{}" { // empty struct
			dst.PropagationStoreConfigurationAquera = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationAquera = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationAzureADSAMLV2
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationAzureADSAMLV2)
	if err == nil {
		jsonPropagationStoreConfigurationAzureADSAMLV2, _ := json.Marshal(dst.PropagationStoreConfigurationAzureADSAMLV2)
		if string(jsonPropagationStoreConfigurationAzureADSAMLV2) == "{}" { // empty struct
			dst.PropagationStoreConfigurationAzureADSAMLV2 = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationAzureADSAMLV2 = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationGithubEMU
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationGithubEMU)
	if err == nil {
		jsonPropagationStoreConfigurationGithubEMU, _ := json.Marshal(dst.PropagationStoreConfigurationGithubEMU)
		if string(jsonPropagationStoreConfigurationGithubEMU) == "{}" { // empty struct
			dst.PropagationStoreConfigurationGithubEMU = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationGithubEMU = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationGoogleApps
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationGoogleApps)
	if err == nil {
		jsonPropagationStoreConfigurationGoogleApps, _ := json.Marshal(dst.PropagationStoreConfigurationGoogleApps)
		if string(jsonPropagationStoreConfigurationGoogleApps) == "{}" { // empty struct
			dst.PropagationStoreConfigurationGoogleApps = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationGoogleApps = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationLDAPGateway
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationLDAPGateway)
	if err == nil {
		jsonPropagationStoreConfigurationLDAPGateway, _ := json.Marshal(dst.PropagationStoreConfigurationLDAPGateway)
		if string(jsonPropagationStoreConfigurationLDAPGateway) == "{}" { // empty struct
			dst.PropagationStoreConfigurationLDAPGateway = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationLDAPGateway = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationPingOne
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationPingOne)
	if err == nil {
		jsonPropagationStoreConfigurationPingOne, _ := json.Marshal(dst.PropagationStoreConfigurationPingOne)
		if string(jsonPropagationStoreConfigurationPingOne) == "{}" { // empty struct
			dst.PropagationStoreConfigurationPingOne = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationPingOne = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationSCIM
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationSCIM)
	if err == nil {
		jsonPropagationStoreConfigurationSCIM, _ := json.Marshal(dst.PropagationStoreConfigurationSCIM)
		if string(jsonPropagationStoreConfigurationSCIM) == "{}" { // empty struct
			dst.PropagationStoreConfigurationSCIM = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationSCIM = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationSalesforce
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationSalesforce)
	if err == nil {
		jsonPropagationStoreConfigurationSalesforce, _ := json.Marshal(dst.PropagationStoreConfigurationSalesforce)
		if string(jsonPropagationStoreConfigurationSalesforce) == "{}" { // empty struct
			dst.PropagationStoreConfigurationSalesforce = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationSalesforce = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationSalesforceContacts
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationSalesforceContacts)
	if err == nil {
		jsonPropagationStoreConfigurationSalesforceContacts, _ := json.Marshal(dst.PropagationStoreConfigurationSalesforceContacts)
		if string(jsonPropagationStoreConfigurationSalesforceContacts) == "{}" { // empty struct
			dst.PropagationStoreConfigurationSalesforceContacts = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationSalesforceContacts = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationServiceNow
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationServiceNow)
	if err == nil {
		jsonPropagationStoreConfigurationServiceNow, _ := json.Marshal(dst.PropagationStoreConfigurationServiceNow)
		if string(jsonPropagationStoreConfigurationServiceNow) == "{}" { // empty struct
			dst.PropagationStoreConfigurationServiceNow = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationServiceNow = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationSlack
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationSlack)
	if err == nil {
		jsonPropagationStoreConfigurationSlack, _ := json.Marshal(dst.PropagationStoreConfigurationSlack)
		if string(jsonPropagationStoreConfigurationSlack) == "{}" { // empty struct
			dst.PropagationStoreConfigurationSlack = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationSlack = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationWorkday
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationWorkday)
	if err == nil {
		jsonPropagationStoreConfigurationWorkday, _ := json.Marshal(dst.PropagationStoreConfigurationWorkday)
		if string(jsonPropagationStoreConfigurationWorkday) == "{}" { // empty struct
			dst.PropagationStoreConfigurationWorkday = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationWorkday = nil
	}

	// try to unmarshal data into PropagationStoreConfigurationZoom
	err = newStrictDecoder(data).Decode(&dst.PropagationStoreConfigurationZoom)
	if err == nil {
		jsonPropagationStoreConfigurationZoom, _ := json.Marshal(dst.PropagationStoreConfigurationZoom)
		if string(jsonPropagationStoreConfigurationZoom) == "{}" { // empty struct
			dst.PropagationStoreConfigurationZoom = nil
		} else {
			match++
		}
	} else {
		dst.PropagationStoreConfigurationZoom = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.PropagationStoreConfigurationAquera = nil
		dst.PropagationStoreConfigurationAzureADSAMLV2 = nil
		dst.PropagationStoreConfigurationGithubEMU = nil
		dst.PropagationStoreConfigurationGoogleApps = nil
		dst.PropagationStoreConfigurationLDAPGateway = nil
		dst.PropagationStoreConfigurationPingOne = nil
		dst.PropagationStoreConfigurationSCIM = nil
		dst.PropagationStoreConfigurationSalesforce = nil
		dst.PropagationStoreConfigurationSalesforceContacts = nil
		dst.PropagationStoreConfigurationServiceNow = nil
		dst.PropagationStoreConfigurationSlack = nil
		dst.PropagationStoreConfigurationWorkday = nil
		dst.PropagationStoreConfigurationZoom = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PropagationStoreConfiguration)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PropagationStoreConfiguration)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PropagationStoreConfiguration) MarshalJSON() ([]byte, error) {
	if src.PropagationStoreConfigurationAquera != nil {
		return json.Marshal(&src.PropagationStoreConfigurationAquera)
	}

	if src.PropagationStoreConfigurationAzureADSAMLV2 != nil {
		return json.Marshal(&src.PropagationStoreConfigurationAzureADSAMLV2)
	}

	if src.PropagationStoreConfigurationGithubEMU != nil {
		return json.Marshal(&src.PropagationStoreConfigurationGithubEMU)
	}

	if src.PropagationStoreConfigurationGoogleApps != nil {
		return json.Marshal(&src.PropagationStoreConfigurationGoogleApps)
	}

	if src.PropagationStoreConfigurationLDAPGateway != nil {
		return json.Marshal(&src.PropagationStoreConfigurationLDAPGateway)
	}

	if src.PropagationStoreConfigurationPingOne != nil {
		return json.Marshal(&src.PropagationStoreConfigurationPingOne)
	}

	if src.PropagationStoreConfigurationSCIM != nil {
		return json.Marshal(&src.PropagationStoreConfigurationSCIM)
	}

	if src.PropagationStoreConfigurationSalesforce != nil {
		return json.Marshal(&src.PropagationStoreConfigurationSalesforce)
	}

	if src.PropagationStoreConfigurationSalesforceContacts != nil {
		return json.Marshal(&src.PropagationStoreConfigurationSalesforceContacts)
	}

	if src.PropagationStoreConfigurationServiceNow != nil {
		return json.Marshal(&src.PropagationStoreConfigurationServiceNow)
	}

	if src.PropagationStoreConfigurationSlack != nil {
		return json.Marshal(&src.PropagationStoreConfigurationSlack)
	}

	if src.PropagationStoreConfigurationWorkday != nil {
		return json.Marshal(&src.PropagationStoreConfigurationWorkday)
	}

	if src.PropagationStoreConfigurationZoom != nil {
		return json.Marshal(&src.PropagationStoreConfigurationZoom)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PropagationStoreConfiguration) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.PropagationStoreConfigurationAquera != nil {
		return obj.PropagationStoreConfigurationAquera
	}

	if obj.PropagationStoreConfigurationAzureADSAMLV2 != nil {
		return obj.PropagationStoreConfigurationAzureADSAMLV2
	}

	if obj.PropagationStoreConfigurationGithubEMU != nil {
		return obj.PropagationStoreConfigurationGithubEMU
	}

	if obj.PropagationStoreConfigurationGoogleApps != nil {
		return obj.PropagationStoreConfigurationGoogleApps
	}

	if obj.PropagationStoreConfigurationLDAPGateway != nil {
		return obj.PropagationStoreConfigurationLDAPGateway
	}

	if obj.PropagationStoreConfigurationPingOne != nil {
		return obj.PropagationStoreConfigurationPingOne
	}

	if obj.PropagationStoreConfigurationSCIM != nil {
		return obj.PropagationStoreConfigurationSCIM
	}

	if obj.PropagationStoreConfigurationSalesforce != nil {
		return obj.PropagationStoreConfigurationSalesforce
	}

	if obj.PropagationStoreConfigurationSalesforceContacts != nil {
		return obj.PropagationStoreConfigurationSalesforceContacts
	}

	if obj.PropagationStoreConfigurationServiceNow != nil {
		return obj.PropagationStoreConfigurationServiceNow
	}

	if obj.PropagationStoreConfigurationSlack != nil {
		return obj.PropagationStoreConfigurationSlack
	}

	if obj.PropagationStoreConfigurationWorkday != nil {
		return obj.PropagationStoreConfigurationWorkday
	}

	if obj.PropagationStoreConfigurationZoom != nil {
		return obj.PropagationStoreConfigurationZoom
	}

	// all schemas are nil
	return nil
}

type NullablePropagationStoreConfiguration struct {
	value *PropagationStoreConfiguration
	isSet bool
}

func (v NullablePropagationStoreConfiguration) Get() *PropagationStoreConfiguration {
	return v.value
}

func (v *NullablePropagationStoreConfiguration) Set(val *PropagationStoreConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullablePropagationStoreConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullablePropagationStoreConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropagationStoreConfiguration(val *PropagationStoreConfiguration) *NullablePropagationStoreConfiguration {
	return &NullablePropagationStoreConfiguration{value: val, isSet: true}
}

func (v NullablePropagationStoreConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropagationStoreConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


