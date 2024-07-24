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

// checks if the AuthorizeEditorDataDefinitionsProcessorDefinitionDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataDefinitionsProcessorDefinitionDTO{}

// AuthorizeEditorDataDefinitionsProcessorDefinitionDTO struct for AuthorizeEditorDataDefinitionsProcessorDefinitionDTO
type AuthorizeEditorDataDefinitionsProcessorDefinitionDTO struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// HAL embedded resources
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	// The resource's unique identifier
	Id *string `json:"id,omitempty"`
	Version *string `json:"version,omitempty"`
	Name string `json:"name"`
	FullName *string `json:"fullName,omitempty"`
	Description *string `json:"description,omitempty"`
	Parent *AuthorizeEditorDataReferenceObjectDTO `json:"parent,omitempty"`
	Processor AuthorizeEditorDataProcessorDTO `json:"processor"`
}

// NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTO instantiates a new AuthorizeEditorDataDefinitionsProcessorDefinitionDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTO(name string, processor AuthorizeEditorDataProcessorDTO) *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO {
	this := AuthorizeEditorDataDefinitionsProcessorDefinitionDTO{}
	this.Name = name
	this.Processor = processor
	return &this
}

// NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTOWithDefaults instantiates a new AuthorizeEditorDataDefinitionsProcessorDefinitionDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataDefinitionsProcessorDefinitionDTOWithDefaults() *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO {
	this := AuthorizeEditorDataDefinitionsProcessorDefinitionDTO{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetId(v string) {
	o.Id = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetVersion(v string) {
	o.Version = &v
}

// GetName returns the Name field value
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetName(v string) {
	o.Name = v
}

// GetFullName returns the FullName field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetFullName() string {
	if o == nil || IsNil(o.FullName) {
		var ret string
		return ret
	}
	return *o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetFullNameOk() (*string, bool) {
	if o == nil || IsNil(o.FullName) {
		return nil, false
	}
	return o.FullName, true
}

// HasFullName returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasFullName() bool {
	if o != nil && !IsNil(o.FullName) {
		return true
	}

	return false
}

// SetFullName gets a reference to the given string and assigns it to the FullName field.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetFullName(v string) {
	o.FullName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetDescription(v string) {
	o.Description = &v
}

// GetParent returns the Parent field value if set, zero value otherwise.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetParent() AuthorizeEditorDataReferenceObjectDTO {
	if o == nil || IsNil(o.Parent) {
		var ret AuthorizeEditorDataReferenceObjectDTO
		return ret
	}
	return *o.Parent
}

// GetParentOk returns a tuple with the Parent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetParentOk() (*AuthorizeEditorDataReferenceObjectDTO, bool) {
	if o == nil || IsNil(o.Parent) {
		return nil, false
	}
	return o.Parent, true
}

// HasParent returns a boolean if a field has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) HasParent() bool {
	if o != nil && !IsNil(o.Parent) {
		return true
	}

	return false
}

// SetParent gets a reference to the given AuthorizeEditorDataReferenceObjectDTO and assigns it to the Parent field.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetParent(v AuthorizeEditorDataReferenceObjectDTO) {
	o.Parent = &v
}

// GetProcessor returns the Processor field value
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetProcessor() AuthorizeEditorDataProcessorDTO {
	if o == nil {
		var ret AuthorizeEditorDataProcessorDTO
		return ret
	}

	return o.Processor
}

// GetProcessorOk returns a tuple with the Processor field value
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) GetProcessorOk() (*AuthorizeEditorDataProcessorDTO, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Processor, true
}

// SetProcessor sets field value
func (o *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) SetProcessor(v AuthorizeEditorDataProcessorDTO) {
	o.Processor = v
}

func (o AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.FullName) {
		toSerialize["fullName"] = o.FullName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Parent) {
		toSerialize["parent"] = o.Parent
	}
	toSerialize["processor"] = o.Processor
	return toSerialize, nil
}

type NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO struct {
	value *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO) Get() *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO) Set(val *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO(val *AuthorizeEditorDataDefinitionsProcessorDefinitionDTO) *NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO {
	return &NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataDefinitionsProcessorDefinitionDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

