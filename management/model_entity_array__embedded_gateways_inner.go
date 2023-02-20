/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

package management

import (
	"encoding/json"
	"fmt"
)

// EntityArrayEmbeddedGatewaysInner struct for EntityArrayEmbeddedGatewaysInner
type EntityArrayEmbeddedGatewaysInner struct {
	Gateway           *Gateway
	GatewayTypeLDAP   *GatewayTypeLDAP
	GatewayTypeRADIUS *GatewayTypeRADIUS
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EntityArrayEmbeddedGatewaysInner) UnmarshalJSON(data []byte) error {

	var err error
	// try to unmarshal JSON data into Gateway
	err = json.Unmarshal(data, &dst.Gateway)
	if err == nil {
		jsonGateway, _ := json.Marshal(dst.Gateway)
		if string(jsonGateway) == "{}" { // empty struct
			dst.Gateway = nil
		} else {
			if dst.Gateway.Type != ENUMGATEWAYTYPE_LDAP {
				return nil // data stored in dst.Gateway, return on the first match
			} else {
				dst.Gateway = nil
			}
		}
	} else {
		dst.Gateway = nil
	}

	// try to unmarshal JSON data into GatewayTypeLDAP
	err = json.Unmarshal(data, &dst.GatewayTypeLDAP)
	if err == nil {
		jsonGatewayLDAP, _ := json.Marshal(dst.GatewayTypeLDAP)
		if string(jsonGatewayLDAP) == "{}" { // empty struct
			dst.GatewayTypeLDAP = nil
		} else {
			if dst.GatewayTypeLDAP.Type == ENUMGATEWAYTYPE_LDAP {
				return nil // data stored in dst.GatewayLDAP, return on the first match
			} else {
				dst.GatewayTypeLDAP = nil
			}
		}
	} else {
		dst.GatewayTypeLDAP = nil
	}

	// try to unmarshal JSON data into GatewayTypeRADIUS
	err = json.Unmarshal(data, &dst.GatewayTypeRADIUS)
	if err == nil {
		jsonGatewayRADIUS, _ := json.Marshal(dst.GatewayTypeRADIUS)
		if string(jsonGatewayRADIUS) == "{}" { // empty struct
			dst.GatewayTypeRADIUS = nil
		} else {
			if dst.GatewayTypeRADIUS.Type == ENUMGATEWAYTYPE_RADIUS {
				return nil // data stored in dst.GatewayTypeRADIUS, return on the first match
			} else {
				dst.GatewayTypeRADIUS = nil
			}
		}
	} else {
		dst.GatewayTypeRADIUS = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(EntityArrayEmbeddedGatewaysInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EntityArrayEmbeddedGatewaysInner) MarshalJSON() ([]byte, error) {
	if src.Gateway != nil {
		return json.Marshal(&src.Gateway)
	}

	if src.GatewayTypeLDAP != nil {
		return json.Marshal(&src.GatewayTypeLDAP)
	}

	if src.GatewayTypeRADIUS != nil {
		return json.Marshal(&src.GatewayTypeRADIUS)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEntityArrayEmbeddedGatewaysInner struct {
	value *EntityArrayEmbeddedGatewaysInner
	isSet bool
}

func (v NullableEntityArrayEmbeddedGatewaysInner) Get() *EntityArrayEmbeddedGatewaysInner {
	return v.value
}

func (v *NullableEntityArrayEmbeddedGatewaysInner) Set(val *EntityArrayEmbeddedGatewaysInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbeddedGatewaysInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbeddedGatewaysInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbeddedGatewaysInner(val *EntityArrayEmbeddedGatewaysInner) *NullableEntityArrayEmbeddedGatewaysInner {
	return &NullableEntityArrayEmbeddedGatewaysInner{value: val, isSet: true}
}

func (v NullableEntityArrayEmbeddedGatewaysInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbeddedGatewaysInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
