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

// checks if the ApplicationResourceGrantResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationResourceGrantResource{}

// ApplicationResourceGrantResource struct for ApplicationResourceGrantResource
type ApplicationResourceGrantResource struct {
	// A string that specifies the ID of the protected resource associated with this grant. This is a required property.
	Id string `json:"id"`
}

// NewApplicationResourceGrantResource instantiates a new ApplicationResourceGrantResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationResourceGrantResource(id string) *ApplicationResourceGrantResource {
	this := ApplicationResourceGrantResource{}
	this.Id = id
	return &this
}

// NewApplicationResourceGrantResourceWithDefaults instantiates a new ApplicationResourceGrantResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationResourceGrantResourceWithDefaults() *ApplicationResourceGrantResource {
	this := ApplicationResourceGrantResource{}
	return &this
}

// GetId returns the Id field value
func (o *ApplicationResourceGrantResource) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApplicationResourceGrantResource) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApplicationResourceGrantResource) SetId(v string) {
	o.Id = v
}

func (o ApplicationResourceGrantResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationResourceGrantResource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableApplicationResourceGrantResource struct {
	value *ApplicationResourceGrantResource
	isSet bool
}

func (v NullableApplicationResourceGrantResource) Get() *ApplicationResourceGrantResource {
	return v.value
}

func (v *NullableApplicationResourceGrantResource) Set(val *ApplicationResourceGrantResource) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationResourceGrantResource) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationResourceGrantResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationResourceGrantResource(val *ApplicationResourceGrantResource) *NullableApplicationResourceGrantResource {
	return &NullableApplicationResourceGrantResource{value: val, isSet: true}
}

func (v NullableApplicationResourceGrantResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationResourceGrantResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


