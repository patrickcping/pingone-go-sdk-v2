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

// checks if the AuthorizeEditorDataConnectorsConnectorTemplateDTO type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthorizeEditorDataConnectorsConnectorTemplateDTO{}

// AuthorizeEditorDataConnectorsConnectorTemplateDTO struct for AuthorizeEditorDataConnectorsConnectorTemplateDTO
type AuthorizeEditorDataConnectorsConnectorTemplateDTO struct {
	Links *map[string]LinksHATEOASValue `json:"_links,omitempty"`
	// HAL embedded resources
	Embedded map[string]map[string]interface{} `json:"_embedded,omitempty"`
	Code *string `json:"code,omitempty"`
	Channel *string `json:"channel,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	ConnectorCapability []AuthorizeEditorDataConnectorsConnectorCapability `json:"ConnectorCapability,omitempty"`
	Capabilities []AuthorizeEditorDataConnectorsConnectorCapability `json:"capabilities,omitempty"`
}

// NewAuthorizeEditorDataConnectorsConnectorTemplateDTO instantiates a new AuthorizeEditorDataConnectorsConnectorTemplateDTO object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthorizeEditorDataConnectorsConnectorTemplateDTO() *AuthorizeEditorDataConnectorsConnectorTemplateDTO {
	this := AuthorizeEditorDataConnectorsConnectorTemplateDTO{}
	return &this
}

// NewAuthorizeEditorDataConnectorsConnectorTemplateDTOWithDefaults instantiates a new AuthorizeEditorDataConnectorsConnectorTemplateDTO object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthorizeEditorDataConnectorsConnectorTemplateDTOWithDefaults() *AuthorizeEditorDataConnectorsConnectorTemplateDTO {
	this := AuthorizeEditorDataConnectorsConnectorTemplateDTO{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetLinks() map[string]LinksHATEOASValue {
	if o == nil || IsNil(o.Links) {
		var ret map[string]LinksHATEOASValue
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetLinksOk() (*map[string]LinksHATEOASValue, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksHATEOASValue and assigns it to the Links field.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetLinks(v map[string]LinksHATEOASValue) {
	o.Links = &v
}

// GetEmbedded returns the Embedded field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetEmbedded() map[string]map[string]interface{} {
	if o == nil || IsNil(o.Embedded) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.Embedded
}

// GetEmbeddedOk returns a tuple with the Embedded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetEmbeddedOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.Embedded) {
		return map[string]map[string]interface{}{}, false
	}
	return o.Embedded, true
}

// HasEmbedded returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasEmbedded() bool {
	if o != nil && !IsNil(o.Embedded) {
		return true
	}

	return false
}

// SetEmbedded gets a reference to the given map[string]map[string]interface{} and assigns it to the Embedded field.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetEmbedded(v map[string]map[string]interface{}) {
	o.Embedded = v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetCode(v string) {
	o.Code = &v
}

// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetChannel() string {
	if o == nil || IsNil(o.Channel) {
		var ret string
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetChannelOk() (*string, bool) {
	if o == nil || IsNil(o.Channel) {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasChannel() bool {
	if o != nil && !IsNil(o.Channel) {
		return true
	}

	return false
}

// SetChannel gets a reference to the given string and assigns it to the Channel field.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetChannel(v string) {
	o.Channel = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetDescription(v string) {
	o.Description = &v
}

// GetConnectorCapability returns the ConnectorCapability field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetConnectorCapability() []AuthorizeEditorDataConnectorsConnectorCapability {
	if o == nil || IsNil(o.ConnectorCapability) {
		var ret []AuthorizeEditorDataConnectorsConnectorCapability
		return ret
	}
	return o.ConnectorCapability
}

// GetConnectorCapabilityOk returns a tuple with the ConnectorCapability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetConnectorCapabilityOk() ([]AuthorizeEditorDataConnectorsConnectorCapability, bool) {
	if o == nil || IsNil(o.ConnectorCapability) {
		return nil, false
	}
	return o.ConnectorCapability, true
}

// HasConnectorCapability returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasConnectorCapability() bool {
	if o != nil && !IsNil(o.ConnectorCapability) {
		return true
	}

	return false
}

// SetConnectorCapability gets a reference to the given []AuthorizeEditorDataConnectorsConnectorCapability and assigns it to the ConnectorCapability field.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetConnectorCapability(v []AuthorizeEditorDataConnectorsConnectorCapability) {
	o.ConnectorCapability = v
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetCapabilities() []AuthorizeEditorDataConnectorsConnectorCapability {
	if o == nil || IsNil(o.Capabilities) {
		var ret []AuthorizeEditorDataConnectorsConnectorCapability
		return ret
	}
	return o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) GetCapabilitiesOk() ([]AuthorizeEditorDataConnectorsConnectorCapability, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given []AuthorizeEditorDataConnectorsConnectorCapability and assigns it to the Capabilities field.
func (o *AuthorizeEditorDataConnectorsConnectorTemplateDTO) SetCapabilities(v []AuthorizeEditorDataConnectorsConnectorCapability) {
	o.Capabilities = v
}

func (o AuthorizeEditorDataConnectorsConnectorTemplateDTO) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthorizeEditorDataConnectorsConnectorTemplateDTO) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Embedded) {
		toSerialize["_embedded"] = o.Embedded
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Channel) {
		toSerialize["channel"] = o.Channel
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.ConnectorCapability) {
		toSerialize["ConnectorCapability"] = o.ConnectorCapability
	}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	return toSerialize, nil
}

type NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO struct {
	value *AuthorizeEditorDataConnectorsConnectorTemplateDTO
	isSet bool
}

func (v NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO) Get() *AuthorizeEditorDataConnectorsConnectorTemplateDTO {
	return v.value
}

func (v *NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO) Set(val *AuthorizeEditorDataConnectorsConnectorTemplateDTO) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizeEditorDataConnectorsConnectorTemplateDTO(val *AuthorizeEditorDataConnectorsConnectorTemplateDTO) *NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO {
	return &NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO{value: val, isSet: true}
}

func (v NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizeEditorDataConnectorsConnectorTemplateDTO) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

