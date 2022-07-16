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

// Environment struct for Environment
type Environment struct {
	BillOfMaterials *BillOfMaterials `json:"billOfMaterials,omitempty"`
	// The time the resource was created.
	CreatedAt *string `json:"createdAt,omitempty"`
	// A string that specifies the description of the population.
	Description *string `json:"description,omitempty"`
	// A string that specifies the resource’s unique identifier.
	Id *string `json:"id,omitempty"`
	License EnvironmentLicense `json:"license"`
	// A string that specifies the environment name, which must be provided and must be unique within an organization.
	Name string `json:"name"`
	Organization *EnvironmentOrganization `json:"organization,omitempty"`
	// A string that specifies the region in which this environment will be used. The value is set when the environment is created and cannot be updated. Options are NA, EU, and AP.
	Region string `json:"region"`
	// A string that specifies the type of environment to use. Options are PRODUCTION and SANDBOX.
	Type string `json:"type"`
	// The time the resource was last updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// NewEnvironment instantiates a new Environment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironment(license EnvironmentLicense, name string, region string, type_ string) *Environment {
	this := Environment{}
	this.License = license
	this.Name = name
	this.Region = region
	this.Type = type_
	return &this
}

// NewEnvironmentWithDefaults instantiates a new Environment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentWithDefaults() *Environment {
	this := Environment{}
	return &this
}

// GetBillOfMaterials returns the BillOfMaterials field value if set, zero value otherwise.
func (o *Environment) GetBillOfMaterials() BillOfMaterials {
	if o == nil || o.BillOfMaterials == nil {
		var ret BillOfMaterials
		return ret
	}
	return *o.BillOfMaterials
}

// GetBillOfMaterialsOk returns a tuple with the BillOfMaterials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetBillOfMaterialsOk() (*BillOfMaterials, bool) {
	if o == nil || o.BillOfMaterials == nil {
		return nil, false
	}
	return o.BillOfMaterials, true
}

// HasBillOfMaterials returns a boolean if a field has been set.
func (o *Environment) HasBillOfMaterials() bool {
	if o != nil && o.BillOfMaterials != nil {
		return true
	}

	return false
}

// SetBillOfMaterials gets a reference to the given BillOfMaterials and assigns it to the BillOfMaterials field.
func (o *Environment) SetBillOfMaterials(v BillOfMaterials) {
	o.BillOfMaterials = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Environment) GetCreatedAt() string {
	if o == nil || o.CreatedAt == nil {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetCreatedAtOk() (*string, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Environment) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Environment) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Environment) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Environment) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Environment) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Environment) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Environment) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Environment) SetId(v string) {
	o.Id = &v
}

// GetLicense returns the License field value
func (o *Environment) GetLicense() EnvironmentLicense {
	if o == nil {
		var ret EnvironmentLicense
		return ret
	}

	return o.License
}

// GetLicenseOk returns a tuple with the License field value
// and a boolean to check if the value has been set.
func (o *Environment) GetLicenseOk() (*EnvironmentLicense, bool) {
	if o == nil {
		return nil, false
	}
	return &o.License, true
}

// SetLicense sets field value
func (o *Environment) SetLicense(v EnvironmentLicense) {
	o.License = v
}

// GetName returns the Name field value
func (o *Environment) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Environment) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Environment) SetName(v string) {
	o.Name = v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *Environment) GetOrganization() EnvironmentOrganization {
	if o == nil || o.Organization == nil {
		var ret EnvironmentOrganization
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetOrganizationOk() (*EnvironmentOrganization, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *Environment) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given EnvironmentOrganization and assigns it to the Organization field.
func (o *Environment) SetOrganization(v EnvironmentOrganization) {
	o.Organization = &v
}

// GetRegion returns the Region field value
func (o *Environment) GetRegion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Region
}

// GetRegionOk returns a tuple with the Region field value
// and a boolean to check if the value has been set.
func (o *Environment) GetRegionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Region, true
}

// SetRegion sets field value
func (o *Environment) SetRegion(v string) {
	o.Region = v
}

// GetType returns the Type field value
func (o *Environment) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Environment) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Environment) SetType(v string) {
	o.Type = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Environment) GetUpdatedAt() string {
	if o == nil || o.UpdatedAt == nil {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Environment) GetUpdatedAtOk() (*string, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Environment) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *Environment) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o Environment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BillOfMaterials != nil {
		toSerialize["billOfMaterials"] = o.BillOfMaterials
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["license"] = o.License
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["region"] = o.Region
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironment struct {
	value *Environment
	isSet bool
}

func (v NullableEnvironment) Get() *Environment {
	return v.value
}

func (v *NullableEnvironment) Set(val *Environment) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironment) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironment(val *Environment) *NullableEnvironment {
	return &NullableEnvironment{value: val, isSet: true}
}

func (v NullableEnvironment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


