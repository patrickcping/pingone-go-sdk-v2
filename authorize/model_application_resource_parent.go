/*
PingOne Platform API - Authorize

The PingOne Platform API covering the PingOne Authorize service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authorize

import (
	"encoding/json"
)

// checks if the ApplicationResourceParent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationResourceParent{}

// ApplicationResourceParent The application resource's parent.
type ApplicationResourceParent struct {
	Type EnumApplicationResourceParentType `json:"type"`
	// The application resource's parent ID.
	Id string `json:"id"`
}

// NewApplicationResourceParent instantiates a new ApplicationResourceParent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResourceParent(type_ EnumApplicationResourceParentType, id string) *ApplicationResourceParent {
	this := ApplicationResourceParent{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewApplicationResourceParentWithDefaults instantiates a new ApplicationResourceParent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResourceParentWithDefaults() *ApplicationResourceParent {
	this := ApplicationResourceParent{}
	return &this
}

// GetType returns the Type field value
func (o *ApplicationResourceParent) GetType() EnumApplicationResourceParentType {
	if o == nil {
		var ret EnumApplicationResourceParentType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ApplicationResourceParent) GetTypeOk() (*EnumApplicationResourceParentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ApplicationResourceParent) SetType(v EnumApplicationResourceParentType) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ApplicationResourceParent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationResourceParent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationResourceParent) SetId(v string) {
	o.Id = v
}

func (o ApplicationResourceParent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationResourceParent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableApplicationResourceParent struct {
	value *ApplicationResourceParent
	isSet bool
}

func (v NullableApplicationResourceParent) Get() *ApplicationResourceParent {
	return v.value
}

func (v *NullableApplicationResourceParent) Set(val *ApplicationResourceParent) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResourceParent) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResourceParent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResourceParent(val *ApplicationResourceParent) *NullableApplicationResourceParent {
	return &NullableApplicationResourceParent{value: val, isSet: true}
}

func (v NullableApplicationResourceParent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResourceParent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

