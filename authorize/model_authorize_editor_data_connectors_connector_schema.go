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

// checks if the AuthorizeEditorDataConnectorsConnectorSchema type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataConnectorsConnectorSchema{}

// AuthorizeEditorDataConnectorsConnectorSchema struct for AuthorizeEditorDataConnectorsConnectorSchema
type AuthorizeEditorDataConnectorsConnectorSchema struct {
	Version *int32 `json:"version,omitempty"`
	VersionName *string `json:"versionName,omitempty"`
	Input map[string]interface{} `json:"input,omitempty"`
	Output map[string]interface{} `json:"output,omitempty"`
}

// NewAuthorizeEditorDataConnectorsConnectorSchema instantiates a new AuthorizeEditorDataConnectorsConnectorSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataConnectorsConnectorSchema() *AuthorizeEditorDataConnectorsConnectorSchema {
	this := AuthorizeEditorDataConnectorsConnectorSchema{}
	return &this
}

// NewAuthorizeEditorDataConnectorsConnectorSchemaWithDefaults instantiates a new AuthorizeEditorDataConnectorsConnectorSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataConnectorsConnectorSchemaWithDefaults() *AuthorizeEditorDataConnectorsConnectorSchema {
	this := AuthorizeEditorDataConnectorsConnectorSchema{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) GetVersion() int32 {
	if o == nil || IsNil(o.Version) {
		var ret int32
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) GetVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given int32 and assigns it to the Version field.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) SetVersion(v int32) {
	o.Version = &v
}

// GetVersionName returns the VersionName field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) GetVersionName() string {
	if o == nil || IsNil(o.VersionName) {
		var ret string
		return ret
	}
	return *o.VersionName
}

// GetVersionNameOk returns a tuple with the VersionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) GetVersionNameOk() (*string, bool) {
	if o == nil || IsNil(o.VersionName) {
		return nil, false
	}
	return o.VersionName, true
}

// HasVersionName returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) HasVersionName() bool {
	if o != nil && !IsNil(o.VersionName) {
		return true
	}

	return false
}

// SetVersionName gets a reference to the given string and assigns it to the VersionName field.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) SetVersionName(v string) {
	o.VersionName = &v
}

// GetInput returns the Input field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) GetInput() map[string]interface{} {
	if o == nil || IsNil(o.Input) {
		var ret map[string]interface{}
		return ret
	}
	return o.Input
}

// GetInputOk returns a tuple with the Input field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) GetInputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Input) {
		return map[string]interface{}{}, false
	}
	return o.Input, true
}

// HasInput returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) HasInput() bool {
	if o != nil && !IsNil(o.Input) {
		return true
	}

	return false
}

// SetInput gets a reference to the given map[string]interface{} and assigns it to the Input field.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) SetInput(v map[string]interface{}) {
	o.Input = v
}

// GetOutput returns the Output field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) GetOutput() map[string]interface{} {
	if o == nil || IsNil(o.Output) {
		var ret map[string]interface{}
		return ret
	}
	return o.Output
}

// GetOutputOk returns a tuple with the Output field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) GetOutputOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Output) {
		return map[string]interface{}{}, false
	}
	return o.Output, true
}

// HasOutput returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) HasOutput() bool {
	if o != nil && !IsNil(o.Output) {
		return true
	}

	return false
}

// SetOutput gets a reference to the given map[string]interface{} and assigns it to the Output field.
func (o *AuthorizeEditorDataConnectorsConnectorSchema) SetOutput(v map[string]interface{}) {
	o.Output = v
}

func (o AuthorizeEditorDataConnectorsConnectorSchema) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataConnectorsConnectorSchema) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.VersionName) {
		toSerialize["versionName"] = o.VersionName
	}
	if !IsNil(o.Input) {
		toSerialize["input"] = o.Input
	}
	if !IsNil(o.Output) {
		toSerialize["output"] = o.Output
	}
	return toSerialize, nil
}

type NullableAuthorizeEditorDataConnectorsConnectorSchema struct {
	value *AuthorizeEditorDataConnectorsConnectorSchema
	isSet bool
}

func (v NullableAuthorizeEditorDataConnectorsConnectorSchema) Get() *AuthorizeEditorDataConnectorsConnectorSchema {
	return v.value
}

func (v *NullableAuthorizeEditorDataConnectorsConnectorSchema) Set(val *AuthorizeEditorDataConnectorsConnectorSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataConnectorsConnectorSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataConnectorsConnectorSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataConnectorsConnectorSchema(val *AuthorizeEditorDataConnectorsConnectorSchema) *NullableAuthorizeEditorDataConnectorsConnectorSchema {
	return &NullableAuthorizeEditorDataConnectorsConnectorSchema{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataConnectorsConnectorSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataConnectorsConnectorSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


