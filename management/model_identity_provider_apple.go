/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// checks if the IdentityProviderApple type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdentityProviderApple{}

// IdentityProviderApple struct for IdentityProviderApple
type IdentityProviderApple struct {
	Links map[string]interface{} `json:"_links,omitempty"`
	// The description of the IdP.
	Description *string `json:"description,omitempty"`
	// The current enabled state of the IdP.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Icon *IdentityProviderCommonIcon `json:"icon,omitempty"`
	// The resource ID.
	Id *string `json:"id,omitempty"`
	LoginButtonIcon *IdentityProviderCommonLoginButtonIcon `json:"loginButtonIcon,omitempty"`
	// The name of the IdP.
	Name string `json:"name"`
	Registration *IdentityProviderCommonRegistration `json:"registration,omitempty"`
	Type EnumIdentityProviderExt `json:"type"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// A string that specifies the application ID from Apple. This is the identifier obtained after registering a services ID in the Apple developer portal. This is a required property.
	ClientId string `json:"clientId"`
	// A string that specifies the private key that is used to generate a client secret. This is a required property.
	ClientSecretSigningKey string `json:"clientSecretSigningKey"`
	// A 10-character string that Apple uses to identify an authentication key. This is a required property.
	KeyId string `json:"keyId"`
	// A 10-character string that Apple uses to identify teams. This is a required property.
	TeamId string `json:"teamId"`
}

// NewIdentityProviderApple instantiates a new IdentityProviderApple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityProviderApple(enabled bool, name string, type_ EnumIdentityProviderExt, clientId string, clientSecretSigningKey string, keyId string, teamId string) *IdentityProviderApple {
	this := IdentityProviderApple{}
	this.Enabled = enabled
	this.Name = name
	this.Type = type_
	this.ClientId = clientId
	this.ClientSecretSigningKey = clientSecretSigningKey
	this.KeyId = keyId
	this.TeamId = teamId
	return &this
}

// NewIdentityProviderAppleWithDefaults instantiates a new IdentityProviderApple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityProviderAppleWithDefaults() *IdentityProviderApple {
	this := IdentityProviderApple{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetLinks() map[string]interface{} {
	if o == nil || IsNil(o.Links) {
		var ret map[string]interface{}
		return ret
	}
	return o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetLinksOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Links) {
		return map[string]interface{}{}, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]interface{} and assigns it to the Links field.
func (o *IdentityProviderApple) SetLinks(v map[string]interface{}) {
	o.Links = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *IdentityProviderApple) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *IdentityProviderApple) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *IdentityProviderApple) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetEnvironment() ObjectEnvironment {
	if o == nil || IsNil(o.Environment) {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || IsNil(o.Environment) {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasEnvironment() bool {
	if o != nil && !IsNil(o.Environment) {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *IdentityProviderApple) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetIcon() IdentityProviderCommonIcon {
	if o == nil || IsNil(o.Icon) {
		var ret IdentityProviderCommonIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetIconOk() (*IdentityProviderCommonIcon, bool) {
	if o == nil || IsNil(o.Icon) {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasIcon() bool {
	if o != nil && !IsNil(o.Icon) {
		return true
	}

	return false
}

// SetIcon gets a reference to the given IdentityProviderCommonIcon and assigns it to the Icon field.
func (o *IdentityProviderApple) SetIcon(v IdentityProviderCommonIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdentityProviderApple) SetId(v string) {
	o.Id = &v
}

// GetLoginButtonIcon returns the LoginButtonIcon field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetLoginButtonIcon() IdentityProviderCommonLoginButtonIcon {
	if o == nil || IsNil(o.LoginButtonIcon) {
		var ret IdentityProviderCommonLoginButtonIcon
		return ret
	}
	return *o.LoginButtonIcon
}

// GetLoginButtonIconOk returns a tuple with the LoginButtonIcon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetLoginButtonIconOk() (*IdentityProviderCommonLoginButtonIcon, bool) {
	if o == nil || IsNil(o.LoginButtonIcon) {
		return nil, false
	}
	return o.LoginButtonIcon, true
}

// HasLoginButtonIcon returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasLoginButtonIcon() bool {
	if o != nil && !IsNil(o.LoginButtonIcon) {
		return true
	}

	return false
}

// SetLoginButtonIcon gets a reference to the given IdentityProviderCommonLoginButtonIcon and assigns it to the LoginButtonIcon field.
func (o *IdentityProviderApple) SetLoginButtonIcon(v IdentityProviderCommonLoginButtonIcon) {
	o.LoginButtonIcon = &v
}

// GetName returns the Name field value
func (o *IdentityProviderApple) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *IdentityProviderApple) SetName(v string) {
	o.Name = v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetRegistration() IdentityProviderCommonRegistration {
	if o == nil || IsNil(o.Registration) {
		var ret IdentityProviderCommonRegistration
		return ret
	}
	return *o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetRegistrationOk() (*IdentityProviderCommonRegistration, bool) {
	if o == nil || IsNil(o.Registration) {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasRegistration() bool {
	if o != nil && !IsNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given IdentityProviderCommonRegistration and assigns it to the Registration field.
func (o *IdentityProviderApple) SetRegistration(v IdentityProviderCommonRegistration) {
	o.Registration = &v
}

// GetType returns the Type field value
func (o *IdentityProviderApple) GetType() EnumIdentityProviderExt {
	if o == nil {
		var ret EnumIdentityProviderExt
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetTypeOk() (*EnumIdentityProviderExt, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *IdentityProviderApple) SetType(v EnumIdentityProviderExt) {
	o.Type = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *IdentityProviderApple) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *IdentityProviderApple) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *IdentityProviderApple) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *IdentityProviderApple) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetClientId returns the ClientId field value
func (o *IdentityProviderApple) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *IdentityProviderApple) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecretSigningKey returns the ClientSecretSigningKey field value
func (o *IdentityProviderApple) GetClientSecretSigningKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecretSigningKey
}

// GetClientSecretSigningKeyOk returns a tuple with the ClientSecretSigningKey field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetClientSecretSigningKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecretSigningKey, true
}

// SetClientSecretSigningKey sets field value
func (o *IdentityProviderApple) SetClientSecretSigningKey(v string) {
	o.ClientSecretSigningKey = v
}

// GetKeyId returns the KeyId field value
func (o *IdentityProviderApple) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetKeyIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *IdentityProviderApple) SetKeyId(v string) {
	o.KeyId = v
}

// GetTeamId returns the TeamId field value
func (o *IdentityProviderApple) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *IdentityProviderApple) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *IdentityProviderApple) SetTeamId(v string) {
	o.TeamId = v
}

func (o IdentityProviderApple) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdentityProviderApple) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["enabled"] = o.Enabled
	if !IsNil(o.Environment) {
		toSerialize["environment"] = o.Environment
	}
	if !IsNil(o.Icon) {
		toSerialize["icon"] = o.Icon
	}
	// skip: id is readOnly
	if !IsNil(o.LoginButtonIcon) {
		toSerialize["loginButtonIcon"] = o.LoginButtonIcon
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	toSerialize["type"] = o.Type
	// skip: createdAt is readOnly
	// skip: updatedAt is readOnly
	toSerialize["clientId"] = o.ClientId
	toSerialize["clientSecretSigningKey"] = o.ClientSecretSigningKey
	toSerialize["keyId"] = o.KeyId
	toSerialize["teamId"] = o.TeamId
	return toSerialize, nil
}

type NullableIdentityProviderApple struct {
	value *IdentityProviderApple
	isSet bool
}

func (v NullableIdentityProviderApple) Get() *IdentityProviderApple {
	return v.value
}

func (v *NullableIdentityProviderApple) Set(val *IdentityProviderApple) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityProviderApple) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityProviderApple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityProviderApple(val *IdentityProviderApple) *NullableIdentityProviderApple {
	return &NullableIdentityProviderApple{value: val, isSet: true}
}

func (v NullableIdentityProviderApple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityProviderApple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


