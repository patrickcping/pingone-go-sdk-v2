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

// EntityArrayEmbeddedResourcesInner - struct for EntityArrayEmbeddedResourcesInner
type EntityArrayEmbeddedResourcesInner struct {
	Resource *Resource
	ResourceApplicationResource *ResourceApplicationResource
}

// ResourceAsEntityArrayEmbeddedResourcesInner is a convenience function that returns Resource wrapped in EntityArrayEmbeddedResourcesInner
func ResourceAsEntityArrayEmbeddedResourcesInner(v *Resource) EntityArrayEmbeddedResourcesInner {
	return EntityArrayEmbeddedResourcesInner{
		Resource: v,
	}
}

// ResourceApplicationResourceAsEntityArrayEmbeddedResourcesInner is a convenience function that returns ResourceApplicationResource wrapped in EntityArrayEmbeddedResourcesInner
func ResourceApplicationResourceAsEntityArrayEmbeddedResourcesInner(v *ResourceApplicationResource) EntityArrayEmbeddedResourcesInner {
	return EntityArrayEmbeddedResourcesInner{
		ResourceApplicationResource: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EntityArrayEmbeddedResourcesInner) UnmarshalJSON(data []byte) error {

	var err error
	// try to unmarshal JSON data into Resource
	err = json.Unmarshal(data, &dst.Resource)
	if err == nil {
		jsonResource, _ := json.Marshal(dst.Resource)
		if string(jsonResource) == "{}" { // empty struct
			dst.Resource = nil
		} else {
			if dst.Resource.Type == nil { // we expect an action for this data type
				dst.Resource = nil
			} else {
				return nil // data stored in dst.Resource, return on the first match
			}
		}
	} else {
		dst.Resource = nil
	}

	// try to unmarshal JSON data into ResourceApplicationResource
	err = json.Unmarshal(data, &dst.ResourceApplicationResource)
	if err == nil {
		jsonResourceApplicationResource, _ := json.Marshal(dst.ResourceApplicationResource)
		if string(jsonResourceApplicationResource) == "{}" { // empty struct
			dst.ResourceApplicationResource = nil
		} else {
				return nil // data stored in dst.ResourceApplicationResource, return on the first match
		}
	} else {
		dst.ResourceApplicationResource = nil
	}

	return fmt.Errorf("Data failed to match schemas in oneOf(EntityArrayEmbeddedResourcesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EntityArrayEmbeddedResourcesInner) MarshalJSON() ([]byte, error) {
	if src.Resource != nil {
		return json.Marshal(&src.Resource)
	}

	if src.ResourceApplicationResource != nil {
		return json.Marshal(&src.ResourceApplicationResource)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EntityArrayEmbeddedResourcesInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Resource != nil {
		return obj.Resource
	}

	if obj.ResourceApplicationResource != nil {
		return obj.ResourceApplicationResource
	}

	// all schemas are nil
	return nil
}

type NullableEntityArrayEmbeddedResourcesInner struct {
	value *EntityArrayEmbeddedResourcesInner
	isSet bool
}

func (v NullableEntityArrayEmbeddedResourcesInner) Get() *EntityArrayEmbeddedResourcesInner {
	return v.value
}

func (v *NullableEntityArrayEmbeddedResourcesInner) Set(val *EntityArrayEmbeddedResourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEntityArrayEmbeddedResourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEntityArrayEmbeddedResourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntityArrayEmbeddedResourcesInner(val *EntityArrayEmbeddedResourcesInner) *NullableEntityArrayEmbeddedResourcesInner {
	return &NullableEntityArrayEmbeddedResourcesInner{value: val, isSet: true}
}

func (v NullableEntityArrayEmbeddedResourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntityArrayEmbeddedResourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


