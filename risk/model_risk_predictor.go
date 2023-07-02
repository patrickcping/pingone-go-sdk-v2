/*
PingOne Platform API - PingOne Risk

The PingOne Platform API covering the PingOne Risk service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package risk

import (
	"encoding/json"
	"fmt"
)

// RiskPredictor - struct for RiskPredictor
type RiskPredictor struct {
	RiskPredictorAnonymousNetwork *RiskPredictorAnonymousNetwork
	RiskPredictorComposite *RiskPredictorComposite
	RiskPredictorCustom *RiskPredictorCustom
	RiskPredictorDevice *RiskPredictorDevice
	RiskPredictorGeovelocity *RiskPredictorGeovelocity
	RiskPredictorIPReputation *RiskPredictorIPReputation
	RiskPredictorUserLocationAnomaly *RiskPredictorUserLocationAnomaly
	RiskPredictorUserRiskBehavior *RiskPredictorUserRiskBehavior
	RiskPredictorVelocity *RiskPredictorVelocity
}

// RiskPredictorAnonymousNetworkAsRiskPredictor is a convenience function that returns RiskPredictorAnonymousNetwork wrapped in RiskPredictor
func RiskPredictorAnonymousNetworkAsRiskPredictor(v *RiskPredictorAnonymousNetwork) RiskPredictor {
	return RiskPredictor{
		RiskPredictorAnonymousNetwork: v,
	}
}

// RiskPredictorCompositeAsRiskPredictor is a convenience function that returns RiskPredictorComposite wrapped in RiskPredictor
func RiskPredictorCompositeAsRiskPredictor(v *RiskPredictorComposite) RiskPredictor {
	return RiskPredictor{
		RiskPredictorComposite: v,
	}
}

// RiskPredictorCustomAsRiskPredictor is a convenience function that returns RiskPredictorCustom wrapped in RiskPredictor
func RiskPredictorCustomAsRiskPredictor(v *RiskPredictorCustom) RiskPredictor {
	return RiskPredictor{
		RiskPredictorCustom: v,
	}
}

// RiskPredictorDeviceAsRiskPredictor is a convenience function that returns RiskPredictorDevice wrapped in RiskPredictor
func RiskPredictorDeviceAsRiskPredictor(v *RiskPredictorDevice) RiskPredictor {
	return RiskPredictor{
		RiskPredictorDevice: v,
	}
}

// RiskPredictorGeovelocityAsRiskPredictor is a convenience function that returns RiskPredictorGeovelocity wrapped in RiskPredictor
func RiskPredictorGeovelocityAsRiskPredictor(v *RiskPredictorGeovelocity) RiskPredictor {
	return RiskPredictor{
		RiskPredictorGeovelocity: v,
	}
}

// RiskPredictorIPReputationAsRiskPredictor is a convenience function that returns RiskPredictorIPReputation wrapped in RiskPredictor
func RiskPredictorIPReputationAsRiskPredictor(v *RiskPredictorIPReputation) RiskPredictor {
	return RiskPredictor{
		RiskPredictorIPReputation: v,
	}
}

// RiskPredictorUserLocationAnomalyAsRiskPredictor is a convenience function that returns RiskPredictorUserLocationAnomaly wrapped in RiskPredictor
func RiskPredictorUserLocationAnomalyAsRiskPredictor(v *RiskPredictorUserLocationAnomaly) RiskPredictor {
	return RiskPredictor{
		RiskPredictorUserLocationAnomaly: v,
	}
}

// RiskPredictorUserRiskBehaviorAsRiskPredictor is a convenience function that returns RiskPredictorUserRiskBehavior wrapped in RiskPredictor
func RiskPredictorUserRiskBehaviorAsRiskPredictor(v *RiskPredictorUserRiskBehavior) RiskPredictor {
	return RiskPredictor{
		RiskPredictorUserRiskBehavior: v,
	}
}

// RiskPredictorVelocityAsRiskPredictor is a convenience function that returns RiskPredictorVelocity wrapped in RiskPredictor
func RiskPredictorVelocityAsRiskPredictor(v *RiskPredictorVelocity) RiskPredictor {
	return RiskPredictor{
		RiskPredictorVelocity: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RiskPredictor) UnmarshalJSON(data []byte) error {

	var common RiskPredictorCommon

	if err := json.Unmarshal(data, &common); err != nil {
		return err
	}

	dst.RiskPredictorAnonymousNetwork = nil
	dst.RiskPredictorComposite = nil
	dst.RiskPredictorCustom = nil
	dst.RiskPredictorGeovelocity = nil
	dst.RiskPredictorIPReputation = nil
	dst.RiskPredictorDevice = nil
	dst.RiskPredictorUserRiskBehavior = nil
	dst.RiskPredictorUserLocationAnomaly = nil
	dst.RiskPredictorVelocity = nil

	objType := common.GetType()

	if !objType.IsValid() {
		return nil
	}

	switch objType {
	case ENUMPREDICTORTYPE_ANONYMOUS_NETWORK:
		if err := json.Unmarshal(data, &dst.RiskPredictorAnonymousNetwork); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_COMPOSITE:
		if err := json.Unmarshal(data, &dst.RiskPredictorComposite); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_MAP:
		if err := json.Unmarshal(data, &dst.RiskPredictorCustom); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_GEO_VELOCITY:
		if err := json.Unmarshal(data, &dst.RiskPredictorGeovelocity); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_IP_REPUTATION:
		if err := json.Unmarshal(data, &dst.RiskPredictorIPReputation); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_DEVICE:
		if err := json.Unmarshal(data, &dst.RiskPredictorDevice); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_USER_RISK_BEHAVIOR:
		if err := json.Unmarshal(data, &dst.RiskPredictorUserRiskBehavior); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_USER_LOCATION_ANOMALY:
		if err := json.Unmarshal(data, &dst.RiskPredictorUserLocationAnomaly); err != nil {
			return err
		}
	case ENUMPREDICTORTYPE_VELOCITY:
		if err := json.Unmarshal(data, &dst.RiskPredictorVelocity); err != nil {
			return err
		}
	default:
		return fmt.Errorf("Data failed to match schemas in oneOf(RiskPredictor)")
	}
	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RiskPredictor) MarshalJSON() ([]byte, error) {
	if src.RiskPredictorAnonymousNetwork != nil {
		return json.Marshal(&src.RiskPredictorAnonymousNetwork)
	}

	if src.RiskPredictorComposite != nil {
		return json.Marshal(&src.RiskPredictorComposite)
	}

	if src.RiskPredictorCustom != nil {
		return json.Marshal(&src.RiskPredictorCustom)
	}

	if src.RiskPredictorDevice != nil {
		return json.Marshal(&src.RiskPredictorDevice)
	}

	if src.RiskPredictorGeovelocity != nil {
		return json.Marshal(&src.RiskPredictorGeovelocity)
	}

	if src.RiskPredictorIPReputation != nil {
		return json.Marshal(&src.RiskPredictorIPReputation)
	}

	if src.RiskPredictorUserLocationAnomaly != nil {
		return json.Marshal(&src.RiskPredictorUserLocationAnomaly)
	}

	if src.RiskPredictorUserRiskBehavior != nil {
		return json.Marshal(&src.RiskPredictorUserRiskBehavior)
	}

	if src.RiskPredictorVelocity != nil {
		return json.Marshal(&src.RiskPredictorVelocity)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RiskPredictor) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.RiskPredictorAnonymousNetwork != nil {
		return obj.RiskPredictorAnonymousNetwork
	}

	if obj.RiskPredictorComposite != nil {
		return obj.RiskPredictorComposite
	}

	if obj.RiskPredictorCustom != nil {
		return obj.RiskPredictorCustom
	}

	if obj.RiskPredictorDevice != nil {
		return obj.RiskPredictorDevice
	}

	if obj.RiskPredictorGeovelocity != nil {
		return obj.RiskPredictorGeovelocity
	}

	if obj.RiskPredictorIPReputation != nil {
		return obj.RiskPredictorIPReputation
	}

	if obj.RiskPredictorUserLocationAnomaly != nil {
		return obj.RiskPredictorUserLocationAnomaly
	}

	if obj.RiskPredictorUserRiskBehavior != nil {
		return obj.RiskPredictorUserRiskBehavior
	}

	if obj.RiskPredictorVelocity != nil {
		return obj.RiskPredictorVelocity
	}

	// all schemas are nil
	return nil
}

type NullableRiskPredictor struct {
	value *RiskPredictor
	isSet bool
}

func (v NullableRiskPredictor) Get() *RiskPredictor {
	return v.value
}

func (v *NullableRiskPredictor) Set(val *RiskPredictor) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskPredictor) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskPredictor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskPredictor(val *RiskPredictor) *NullableRiskPredictor {
	return &NullableRiskPredictor{value: val, isSet: true}
}

func (v NullableRiskPredictor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskPredictor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


