/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// SchemaAttribute struct for SchemaAttribute
type SchemaAttribute struct {
	// A string that specifies an optional property that specifies the description of the attribute. If provided, it must not be an empty string. Valid characters consists of any Unicode letter, mark (for example, accent or umlaut), numeric character, punctuation character, or space.
	Description *string `json:"description,omitempty"`
	// A string that specifies an optional property that specifies the display name of the attribute such as 'T-shirt size’. If provided, it must not be an empty string. Valid characters consist of any Unicode letter, mark (for example, accent or umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen.
	DisplayName *string `json:"displayName,omitempty"`
	// A boolean that specifies whether or not the attribute is enabled. This is a required property for POST and PUT operations; it cannot be omitted or explicitly set to null. Disabled attributes are ignored on create/update and not returned on read.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	// A string that specifies the LDAP attribute.
	LdapAttribute *string `json:"ldapAttribute,omitempty"`
	// A string that specifies the name of the attribute. The attribute name must be provided during creation, must not be empty and must not exceed 256 characters. It must also be unique within the schema for an environment. It must start with a letter and may be followed by letters, numbers or hyphens.
	Name string `json:"name"`
	// A boolean that specifies whether or not the attribute is required. Required attributes must be provided a value during create/update. Defaults to false if not provided.
	Required *bool `json:"required,omitempty"`
	Schema *SchemaAttributeSchema `json:"schema,omitempty"`
	SchemaType *EnumSchemaAttributeSchemaType `json:"schemaType,omitempty"`
	Type EnumSchemaAttributeType `json:"type"`
	// A boolean that specifies whether or not the attribute must have a unique value within the environment. This is a required property for POST and PUT operations; it cannot be omitted or explicitly set to null.
	Unique *bool `json:"unique,omitempty"`
	// A boolean that specifies whether the attribute has multiple values or a single one. This value can only change from false to true, as changing from true to false is not allowed. Maximum number of values stored is 1,000.
	MultiValued *bool `json:"multiValued,omitempty"`
}

// NewSchemaAttribute instantiates a new SchemaAttribute object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSchemaAttribute(enabled bool, name string, type_ EnumSchemaAttributeType) *SchemaAttribute {
	this := SchemaAttribute{}
	this.Enabled = enabled
	this.Name = name
	this.Type = type_
	return &this
}

// NewSchemaAttributeWithDefaults instantiates a new SchemaAttribute object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSchemaAttributeWithDefaults() *SchemaAttribute {
	this := SchemaAttribute{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *SchemaAttribute) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *SchemaAttribute) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *SchemaAttribute) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *SchemaAttribute) GetDisplayName() string {
	if o == nil || o.DisplayName == nil {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetDisplayNameOk() (*string, bool) {
	if o == nil || o.DisplayName == nil {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *SchemaAttribute) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *SchemaAttribute) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetEnabled returns the Enabled field value
func (o *SchemaAttribute) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *SchemaAttribute) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *SchemaAttribute) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *SchemaAttribute) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *SchemaAttribute) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SchemaAttribute) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SchemaAttribute) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SchemaAttribute) SetId(v string) {
	o.Id = &v
}

// GetLdapAttribute returns the LdapAttribute field value if set, zero value otherwise.
func (o *SchemaAttribute) GetLdapAttribute() string {
	if o == nil || o.LdapAttribute == nil {
		var ret string
		return ret
	}
	return *o.LdapAttribute
}

// GetLdapAttributeOk returns a tuple with the LdapAttribute field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetLdapAttributeOk() (*string, bool) {
	if o == nil || o.LdapAttribute == nil {
		return nil, false
	}
	return o.LdapAttribute, true
}

// HasLdapAttribute returns a boolean if a field has been set.
func (o *SchemaAttribute) HasLdapAttribute() bool {
	if o != nil && o.LdapAttribute != nil {
		return true
	}

	return false
}

// SetLdapAttribute gets a reference to the given string and assigns it to the LdapAttribute field.
func (o *SchemaAttribute) SetLdapAttribute(v string) {
	o.LdapAttribute = &v
}

// GetName returns the Name field value
func (o *SchemaAttribute) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SchemaAttribute) SetName(v string) {
	o.Name = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *SchemaAttribute) GetRequired() bool {
	if o == nil || o.Required == nil {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetRequiredOk() (*bool, bool) {
	if o == nil || o.Required == nil {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *SchemaAttribute) HasRequired() bool {
	if o != nil && o.Required != nil {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *SchemaAttribute) SetRequired(v bool) {
	o.Required = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *SchemaAttribute) GetSchema() SchemaAttributeSchema {
	if o == nil || o.Schema == nil {
		var ret SchemaAttributeSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetSchemaOk() (*SchemaAttributeSchema, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *SchemaAttribute) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given SchemaAttributeSchema and assigns it to the Schema field.
func (o *SchemaAttribute) SetSchema(v SchemaAttributeSchema) {
	o.Schema = &v
}

// GetSchemaType returns the SchemaType field value if set, zero value otherwise.
func (o *SchemaAttribute) GetSchemaType() EnumSchemaAttributeSchemaType {
	if o == nil || o.SchemaType == nil {
		var ret EnumSchemaAttributeSchemaType
		return ret
	}
	return *o.SchemaType
}

// GetSchemaTypeOk returns a tuple with the SchemaType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetSchemaTypeOk() (*EnumSchemaAttributeSchemaType, bool) {
	if o == nil || o.SchemaType == nil {
		return nil, false
	}
	return o.SchemaType, true
}

// HasSchemaType returns a boolean if a field has been set.
func (o *SchemaAttribute) HasSchemaType() bool {
	if o != nil && o.SchemaType != nil {
		return true
	}

	return false
}

// SetSchemaType gets a reference to the given EnumSchemaAttributeSchemaType and assigns it to the SchemaType field.
func (o *SchemaAttribute) SetSchemaType(v EnumSchemaAttributeSchemaType) {
	o.SchemaType = &v
}

// GetType returns the Type field value
func (o *SchemaAttribute) GetType() EnumSchemaAttributeType {
	if o == nil {
		var ret EnumSchemaAttributeType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetTypeOk() (*EnumSchemaAttributeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SchemaAttribute) SetType(v EnumSchemaAttributeType) {
	o.Type = v
}

// GetUnique returns the Unique field value if set, zero value otherwise.
func (o *SchemaAttribute) GetUnique() bool {
	if o == nil || o.Unique == nil {
		var ret bool
		return ret
	}
	return *o.Unique
}

// GetUniqueOk returns a tuple with the Unique field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetUniqueOk() (*bool, bool) {
	if o == nil || o.Unique == nil {
		return nil, false
	}
	return o.Unique, true
}

// HasUnique returns a boolean if a field has been set.
func (o *SchemaAttribute) HasUnique() bool {
	if o != nil && o.Unique != nil {
		return true
	}

	return false
}

// SetUnique gets a reference to the given bool and assigns it to the Unique field.
func (o *SchemaAttribute) SetUnique(v bool) {
	o.Unique = &v
}

// GetMultiValued returns the MultiValued field value if set, zero value otherwise.
func (o *SchemaAttribute) GetMultiValued() bool {
	if o == nil || o.MultiValued == nil {
		var ret bool
		return ret
	}
	return *o.MultiValued
}

// GetMultiValuedOk returns a tuple with the MultiValued field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SchemaAttribute) GetMultiValuedOk() (*bool, bool) {
	if o == nil || o.MultiValued == nil {
		return nil, false
	}
	return o.MultiValued, true
}

// HasMultiValued returns a boolean if a field has been set.
func (o *SchemaAttribute) HasMultiValued() bool {
	if o != nil && o.MultiValued != nil {
		return true
	}

	return false
}

// SetMultiValued gets a reference to the given bool and assigns it to the MultiValued field.
func (o *SchemaAttribute) SetMultiValued(v bool) {
	o.MultiValued = &v
}

func (o SchemaAttribute) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LdapAttribute != nil {
		toSerialize["ldapAttribute"] = o.LdapAttribute
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Required != nil {
		toSerialize["required"] = o.Required
	}
	if o.Schema != nil {
		toSerialize["schema"] = o.Schema
	}
	if o.SchemaType != nil {
		toSerialize["schemaType"] = o.SchemaType
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Unique != nil {
		toSerialize["unique"] = o.Unique
	}
	if o.MultiValued != nil {
		toSerialize["multiValued"] = o.MultiValued
	}
	return json.Marshal(toSerialize)
}

type NullableSchemaAttribute struct {
	value *SchemaAttribute
	isSet bool
}

func (v NullableSchemaAttribute) Get() *SchemaAttribute {
	return v.value
}

func (v *NullableSchemaAttribute) Set(val *SchemaAttribute) {
	v.value = val
	v.isSet = true
}

func (v NullableSchemaAttribute) IsSet() bool {
	return v.isSet
}

func (v *NullableSchemaAttribute) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSchemaAttribute(val *SchemaAttribute) *NullableSchemaAttribute {
	return &NullableSchemaAttribute{value: val, isSet: true}
}

func (v NullableSchemaAttribute) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSchemaAttribute) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


