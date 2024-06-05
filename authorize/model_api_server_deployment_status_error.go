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

// checks if the APIServerDeploymentStatusError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIServerDeploymentStatusError{}

// APIServerDeploymentStatusError Error details returned if the last deployment request failed.
type APIServerDeploymentStatusError struct {
	// A unique identifier for the error.
	Id *string `json:"id,omitempty"`
	// A general fault code that identifies the the type of error. See [Error codes](https://apidocs.pingidentity.com/pingone/platform/v1/api/#error-codes).
	Code *string `json:"code,omitempty"`
	// A short human-readable description of the error.
	Message *string `json:"message,omitempty"`
}

// NewAPIServerDeploymentStatusError instantiates a new APIServerDeploymentStatusError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIServerDeploymentStatusError() *APIServerDeploymentStatusError {
	this := APIServerDeploymentStatusError{}
	return &this
}

// NewAPIServerDeploymentStatusErrorWithDefaults instantiates a new APIServerDeploymentStatusError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIServerDeploymentStatusErrorWithDefaults() *APIServerDeploymentStatusError {
	this := APIServerDeploymentStatusError{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *APIServerDeploymentStatusError) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeploymentStatusError) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *APIServerDeploymentStatusError) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *APIServerDeploymentStatusError) SetId(v string) {
	o.Id = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *APIServerDeploymentStatusError) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeploymentStatusError) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *APIServerDeploymentStatusError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *APIServerDeploymentStatusError) SetCode(v string) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *APIServerDeploymentStatusError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIServerDeploymentStatusError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *APIServerDeploymentStatusError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *APIServerDeploymentStatusError) SetMessage(v string) {
	o.Message = &v
}

func (o APIServerDeploymentStatusError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIServerDeploymentStatusError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableAPIServerDeploymentStatusError struct {
	value *APIServerDeploymentStatusError
	isSet bool
}

func (v NullableAPIServerDeploymentStatusError) Get() *APIServerDeploymentStatusError {
	return v.value
}

func (v *NullableAPIServerDeploymentStatusError) Set(val *APIServerDeploymentStatusError) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIServerDeploymentStatusError) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIServerDeploymentStatusError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIServerDeploymentStatusError(val *APIServerDeploymentStatusError) *NullableAPIServerDeploymentStatusError {
	return &NullableAPIServerDeploymentStatusError{value: val, isSet: true}
}

func (v NullableAPIServerDeploymentStatusError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIServerDeploymentStatusError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

