/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-07-18
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"encoding/json"
)

// Application struct for Application
type Application struct {
	AccessControl *ApplicationAccessControl `json:"accessControl,omitempty"`
	// A boolean that specifies whether the permissions service should assign default roles to the application. This property is set only on the POST request. The property is ignored when included in a PUT request.
	AssignActorRoles *bool `json:"assignActorRoles,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// A string that specifies the description of the application.
	Description *string `json:"description,omitempty"`
	// A string that specifies the current enabled state of the application. Options are ENABLED or DISABLED.
	Enabled bool `json:"enabled"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Icon *ApplicationIcon `json:"icon,omitempty"`
	// A string that specifies the application ID.
	Id *string `json:"id,omitempty"`
	// A string that specifies the custom login page URL for the application. If you set the loginPageUrl property for applications in an environment that sets a custom domain, the URL should include the top-level domain and at least one additional domain level. Warning To avoid issues with third-party cookies in some browsers, a custom domain must be used, giving your PingOne environment the same parent domain as your authentication application. For more information about custom domains, see Custom domains.
	LoginPageUrl *string `json:"loginPageUrl,omitempty"`
	// A string that specifies the name of the application. This is a required property.
	Name string `json:"name"`
	Protocol EnumApplicationProtocol `json:"protocol"`
	// An array that specifies the list of labels associated with the application. Options are PING_FED_CONNECTION_INTEGRATION.
	Tags []EnumApplicationTags `json:"tags,omitempty"`
	Type EnumApplicationType `json:"type"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
	// A boolean that specifies whether the request query parameter JWT is allowed to be unsigned. If false or null (default), an unsigned request object is not allowed.
	SupportUnsignedRequestObject *bool `json:"supportUnsignedRequestObject,omitempty"`
}

// NewApplication instantiates a new Application object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplication(enabled bool, name string, protocol EnumApplicationProtocol, type_ EnumApplicationType) *Application {
	this := Application{}
	this.Enabled = enabled
	this.Name = name
	this.Protocol = protocol
	this.Type = type_
	return &this
}

// NewApplicationWithDefaults instantiates a new Application object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationWithDefaults() *Application {
	this := Application{}
	return &this
}

// GetAccessControl returns the AccessControl field value if set, zero value otherwise.
func (o *Application) GetAccessControl() ApplicationAccessControl {
	if o == nil || o.AccessControl == nil {
		var ret ApplicationAccessControl
		return ret
	}
	return *o.AccessControl
}

// GetAccessControlOk returns a tuple with the AccessControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAccessControlOk() (*ApplicationAccessControl, bool) {
	if o == nil || o.AccessControl == nil {
		return nil, false
	}
	return o.AccessControl, true
}

// HasAccessControl returns a boolean if a field has been set.
func (o *Application) HasAccessControl() bool {
	if o != nil && o.AccessControl != nil {
		return true
	}

	return false
}

// SetAccessControl gets a reference to the given ApplicationAccessControl and assigns it to the AccessControl field.
func (o *Application) SetAccessControl(v ApplicationAccessControl) {
	o.AccessControl = &v
}

// GetAssignActorRoles returns the AssignActorRoles field value if set, zero value otherwise.
func (o *Application) GetAssignActorRoles() bool {
	if o == nil || o.AssignActorRoles == nil {
		var ret bool
		return ret
	}
	return *o.AssignActorRoles
}

// GetAssignActorRolesOk returns a tuple with the AssignActorRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetAssignActorRolesOk() (*bool, bool) {
	if o == nil || o.AssignActorRoles == nil {
		return nil, false
	}
	return o.AssignActorRoles, true
}

// HasAssignActorRoles returns a boolean if a field has been set.
func (o *Application) HasAssignActorRoles() bool {
	if o != nil && o.AssignActorRoles != nil {
		return true
	}

	return false
}

// SetAssignActorRoles gets a reference to the given bool and assigns it to the AssignActorRoles field.
func (o *Application) SetAssignActorRoles(v bool) {
	o.AssignActorRoles = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Application) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Application) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Application) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Application) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Application) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Application) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value
func (o *Application) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *Application) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *Application) SetEnabled(v bool) {
	o.Enabled = v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *Application) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *Application) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *Application) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetIcon returns the Icon field value if set, zero value otherwise.
func (o *Application) GetIcon() ApplicationIcon {
	if o == nil || o.Icon == nil {
		var ret ApplicationIcon
		return ret
	}
	return *o.Icon
}

// GetIconOk returns a tuple with the Icon field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetIconOk() (*ApplicationIcon, bool) {
	if o == nil || o.Icon == nil {
		return nil, false
	}
	return o.Icon, true
}

// HasIcon returns a boolean if a field has been set.
func (o *Application) HasIcon() bool {
	if o != nil && o.Icon != nil {
		return true
	}

	return false
}

// SetIcon gets a reference to the given ApplicationIcon and assigns it to the Icon field.
func (o *Application) SetIcon(v ApplicationIcon) {
	o.Icon = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Application) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Application) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Application) SetId(v string) {
	o.Id = &v
}

// GetLoginPageUrl returns the LoginPageUrl field value if set, zero value otherwise.
func (o *Application) GetLoginPageUrl() string {
	if o == nil || o.LoginPageUrl == nil {
		var ret string
		return ret
	}
	return *o.LoginPageUrl
}

// GetLoginPageUrlOk returns a tuple with the LoginPageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetLoginPageUrlOk() (*string, bool) {
	if o == nil || o.LoginPageUrl == nil {
		return nil, false
	}
	return o.LoginPageUrl, true
}

// HasLoginPageUrl returns a boolean if a field has been set.
func (o *Application) HasLoginPageUrl() bool {
	if o != nil && o.LoginPageUrl != nil {
		return true
	}

	return false
}

// SetLoginPageUrl gets a reference to the given string and assigns it to the LoginPageUrl field.
func (o *Application) SetLoginPageUrl(v string) {
	o.LoginPageUrl = &v
}

// GetName returns the Name field value
func (o *Application) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Application) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Application) SetName(v string) {
	o.Name = v
}

// GetProtocol returns the Protocol field value
func (o *Application) GetProtocol() EnumApplicationProtocol {
	if o == nil {
		var ret EnumApplicationProtocol
		return ret
	}

	return o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value
// and a boolean to check if the value has been set.
func (o *Application) GetProtocolOk() (*EnumApplicationProtocol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Protocol, true
}

// SetProtocol sets field value
func (o *Application) SetProtocol(v EnumApplicationProtocol) {
	o.Protocol = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Application) GetTags() []EnumApplicationTags {
	if o == nil || o.Tags == nil {
		var ret []EnumApplicationTags
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetTagsOk() ([]EnumApplicationTags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Application) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []EnumApplicationTags and assigns it to the Tags field.
func (o *Application) SetTags(v []EnumApplicationTags) {
	o.Tags = v
}

// GetType returns the Type field value
func (o *Application) GetType() EnumApplicationType {
	if o == nil {
		var ret EnumApplicationType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Application) GetTypeOk() (*EnumApplicationType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Application) SetType(v EnumApplicationType) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Application) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Application) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Application) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

// GetSupportUnsignedRequestObject returns the SupportUnsignedRequestObject field value if set, zero value otherwise.
func (o *Application) GetSupportUnsignedRequestObject() bool {
	if o == nil || o.SupportUnsignedRequestObject == nil {
		var ret bool
		return ret
	}
	return *o.SupportUnsignedRequestObject
}

// GetSupportUnsignedRequestObjectOk returns a tuple with the SupportUnsignedRequestObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Application) GetSupportUnsignedRequestObjectOk() (*bool, bool) {
	if o == nil || o.SupportUnsignedRequestObject == nil {
		return nil, false
	}
	return o.SupportUnsignedRequestObject, true
}

// HasSupportUnsignedRequestObject returns a boolean if a field has been set.
func (o *Application) HasSupportUnsignedRequestObject() bool {
	if o != nil && o.SupportUnsignedRequestObject != nil {
		return true
	}

	return false
}

// SetSupportUnsignedRequestObject gets a reference to the given bool and assigns it to the SupportUnsignedRequestObject field.
func (o *Application) SetSupportUnsignedRequestObject(v bool) {
	o.SupportUnsignedRequestObject = &v
}

func (o Application) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccessControl != nil {
		toSerialize["accessControl"] = o.AccessControl
	}
	if o.AssignActorRoles != nil {
		toSerialize["assignActorRoles"] = o.AssignActorRoles
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Icon != nil {
		toSerialize["icon"] = o.Icon
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LoginPageUrl != nil {
		toSerialize["loginPageUrl"] = o.LoginPageUrl
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["protocol"] = o.Protocol
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.SupportUnsignedRequestObject != nil {
		toSerialize["supportUnsignedRequestObject"] = o.SupportUnsignedRequestObject
	}
	return json.Marshal(toSerialize)
}

type NullableApplication struct {
	value *Application
	isSet bool
}

func (v NullableApplication) Get() *Application {
	return v.value
}

func (v *NullableApplication) Set(val *Application) {
	v.value = val
	v.isSet = true
}

func (v NullableApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplication(val *Application) *NullableApplication {
	return &NullableApplication{value: val, isSet: true}
}

func (v NullableApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


