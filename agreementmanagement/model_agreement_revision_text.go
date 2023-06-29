/*
PingOne Platform API - Agreement Management

The PingOne Platform API covering the PingOne Agreement Management service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agreementmanagement

import (
	"encoding/json"
)

// checks if the AgreementRevisionText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgreementRevisionText{}

// AgreementRevisionText struct for AgreementRevisionText
type AgreementRevisionText struct {
	ResourcePath *string `json:"resourcePath,omitempty"`
	Data *string `json:"data,omitempty"`
	DataType *string `json:"dataType,omitempty"`
}

// NewAgreementRevisionText instantiates a new AgreementRevisionText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgreementRevisionText() *AgreementRevisionText {
	this := AgreementRevisionText{}
	return &this
}

// NewAgreementRevisionTextWithDefaults instantiates a new AgreementRevisionText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgreementRevisionTextWithDefaults() *AgreementRevisionText {
	this := AgreementRevisionText{}
	return &this
}

// GetResourcePath returns the ResourcePath field value if set, zero value otherwise.
func (o *AgreementRevisionText) GetResourcePath() string {
	if o == nil || IsNil(o.ResourcePath) {
		var ret string
		return ret
	}
	return *o.ResourcePath
}

// GetResourcePathOk returns a tuple with the ResourcePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementRevisionText) GetResourcePathOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePath) {
		return nil, false
	}
	return o.ResourcePath, true
}

// HasResourcePath returns a boolean if a field has been set.
func (o *AgreementRevisionText) HasResourcePath() bool {
	if o != nil && !IsNil(o.ResourcePath) {
		return true
	}

	return false
}

// SetResourcePath gets a reference to the given string and assigns it to the ResourcePath field.
func (o *AgreementRevisionText) SetResourcePath(v string) {
	o.ResourcePath = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AgreementRevisionText) GetData() string {
	if o == nil || IsNil(o.Data) {
		var ret string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementRevisionText) GetDataOk() (*string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AgreementRevisionText) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given string and assigns it to the Data field.
func (o *AgreementRevisionText) SetData(v string) {
	o.Data = &v
}

// GetDataType returns the DataType field value if set, zero value otherwise.
func (o *AgreementRevisionText) GetDataType() string {
	if o == nil || IsNil(o.DataType) {
		var ret string
		return ret
	}
	return *o.DataType
}

// GetDataTypeOk returns a tuple with the DataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgreementRevisionText) GetDataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DataType) {
		return nil, false
	}
	return o.DataType, true
}

// HasDataType returns a boolean if a field has been set.
func (o *AgreementRevisionText) HasDataType() bool {
	if o != nil && !IsNil(o.DataType) {
		return true
	}

	return false
}

// SetDataType gets a reference to the given string and assigns it to the DataType field.
func (o *AgreementRevisionText) SetDataType(v string) {
	o.DataType = &v
}

func (o AgreementRevisionText) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgreementRevisionText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourcePath) {
		toSerialize["resourcePath"] = o.ResourcePath
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.DataType) {
		toSerialize["dataType"] = o.DataType
	}
	return toSerialize, nil
}

type NullableAgreementRevisionText struct {
	value *AgreementRevisionText
	isSet bool
}

func (v NullableAgreementRevisionText) Get() *AgreementRevisionText {
	return v.value
}

func (v *NullableAgreementRevisionText) Set(val *AgreementRevisionText) {
	v.value = val
	v.isSet = true
}

func (v NullableAgreementRevisionText) IsSet() bool {
	return v.isSet
}

func (v *NullableAgreementRevisionText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgreementRevisionText(val *AgreementRevisionText) *NullableAgreementRevisionText {
	return &NullableAgreementRevisionText{value: val, isSet: true}
}

func (v NullableAgreementRevisionText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgreementRevisionText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


