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

// checks if the PropagationStoreConfigurationServiceNow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PropagationStoreConfigurationServiceNow{}

// PropagationStoreConfigurationServiceNow struct for PropagationStoreConfigurationServiceNow
type PropagationStoreConfigurationServiceNow struct {
	// Password for the administrator.
	AdministratorPassword string `json:"Administrator_Password"`
	// Username for the administrator.
	AdministratorUsername string `json:"Administrator_Username"`
	// Whether or not users are allowed to be created.
	CREATE_USERS *bool `json:"CREATE_USERS,omitempty"`
	// Whether or not users are allowed to be deprovisioned (removed) following action specified in `REMOVE_ACTION`.
	DEPROVISION_USERS *string `json:"DEPROVISION_USERS,omitempty"`
	// Whether or not new users are allowed to be disabled.
	DISABLE_USERS *bool `json:"DISABLE_USERS,omitempty"`
	REMOVE_ACTION *EnumPropagationStoreTypeRemoveActionDisable `json:"REMOVE_ACTION,omitempty"`
	// The URL for the ServiceNow account.
	ServiceNowUrl string `json:"ServiceNow_Url"`
	// Whether or not users are allowed to be updated.
	UPDATE_USERS *bool `json:"UPDATE_USERS,omitempty"`
}

// NewPropagationStoreConfigurationServiceNow instantiates a new PropagationStoreConfigurationServiceNow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropagationStoreConfigurationServiceNow(administratorPassword string, administratorUsername string, serviceNowUrl string) *PropagationStoreConfigurationServiceNow {
	this := PropagationStoreConfigurationServiceNow{}
	this.AdministratorPassword = administratorPassword
	this.AdministratorUsername = administratorUsername
	this.ServiceNowUrl = serviceNowUrl
	return &this
}

// NewPropagationStoreConfigurationServiceNowWithDefaults instantiates a new PropagationStoreConfigurationServiceNow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropagationStoreConfigurationServiceNowWithDefaults() *PropagationStoreConfigurationServiceNow {
	this := PropagationStoreConfigurationServiceNow{}
	return &this
}

// GetAdministratorPassword returns the AdministratorPassword field value
func (o *PropagationStoreConfigurationServiceNow) GetAdministratorPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdministratorPassword
}

// GetAdministratorPasswordOk returns a tuple with the AdministratorPassword field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationServiceNow) GetAdministratorPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdministratorPassword, true
}

// SetAdministratorPassword sets field value
func (o *PropagationStoreConfigurationServiceNow) SetAdministratorPassword(v string) {
	o.AdministratorPassword = v
}

// GetAdministratorUsername returns the AdministratorUsername field value
func (o *PropagationStoreConfigurationServiceNow) GetAdministratorUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdministratorUsername
}

// GetAdministratorUsernameOk returns a tuple with the AdministratorUsername field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationServiceNow) GetAdministratorUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdministratorUsername, true
}

// SetAdministratorUsername sets field value
func (o *PropagationStoreConfigurationServiceNow) SetAdministratorUsername(v string) {
	o.AdministratorUsername = v
}

// GetCREATE_USERS returns the CREATE_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationServiceNow) GetCREATE_USERS() bool {
	if o == nil || IsNil(o.CREATE_USERS) {
		var ret bool
		return ret
	}
	return *o.CREATE_USERS
}

// GetCREATE_USERSOk returns a tuple with the CREATE_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationServiceNow) GetCREATE_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.CREATE_USERS) {
		return nil, false
	}
	return o.CREATE_USERS, true
}

// HasCREATE_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationServiceNow) HasCREATE_USERS() bool {
	if o != nil && !IsNil(o.CREATE_USERS) {
		return true
	}

	return false
}

// SetCREATE_USERS gets a reference to the given bool and assigns it to the CREATE_USERS field.
func (o *PropagationStoreConfigurationServiceNow) SetCREATE_USERS(v bool) {
	o.CREATE_USERS = &v
}

// GetDEPROVISION_USERS returns the DEPROVISION_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationServiceNow) GetDEPROVISION_USERS() string {
	if o == nil || IsNil(o.DEPROVISION_USERS) {
		var ret string
		return ret
	}
	return *o.DEPROVISION_USERS
}

// GetDEPROVISION_USERSOk returns a tuple with the DEPROVISION_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationServiceNow) GetDEPROVISION_USERSOk() (*string, bool) {
	if o == nil || IsNil(o.DEPROVISION_USERS) {
		return nil, false
	}
	return o.DEPROVISION_USERS, true
}

// HasDEPROVISION_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationServiceNow) HasDEPROVISION_USERS() bool {
	if o != nil && !IsNil(o.DEPROVISION_USERS) {
		return true
	}

	return false
}

// SetDEPROVISION_USERS gets a reference to the given string and assigns it to the DEPROVISION_USERS field.
func (o *PropagationStoreConfigurationServiceNow) SetDEPROVISION_USERS(v string) {
	o.DEPROVISION_USERS = &v
}

// GetDISABLE_USERS returns the DISABLE_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationServiceNow) GetDISABLE_USERS() bool {
	if o == nil || IsNil(o.DISABLE_USERS) {
		var ret bool
		return ret
	}
	return *o.DISABLE_USERS
}

// GetDISABLE_USERSOk returns a tuple with the DISABLE_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationServiceNow) GetDISABLE_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.DISABLE_USERS) {
		return nil, false
	}
	return o.DISABLE_USERS, true
}

// HasDISABLE_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationServiceNow) HasDISABLE_USERS() bool {
	if o != nil && !IsNil(o.DISABLE_USERS) {
		return true
	}

	return false
}

// SetDISABLE_USERS gets a reference to the given bool and assigns it to the DISABLE_USERS field.
func (o *PropagationStoreConfigurationServiceNow) SetDISABLE_USERS(v bool) {
	o.DISABLE_USERS = &v
}

// GetREMOVE_ACTION returns the REMOVE_ACTION field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationServiceNow) GetREMOVE_ACTION() EnumPropagationStoreTypeRemoveActionDisable {
	if o == nil || IsNil(o.REMOVE_ACTION) {
		var ret EnumPropagationStoreTypeRemoveActionDisable
		return ret
	}
	return *o.REMOVE_ACTION
}

// GetREMOVE_ACTIONOk returns a tuple with the REMOVE_ACTION field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationServiceNow) GetREMOVE_ACTIONOk() (*EnumPropagationStoreTypeRemoveActionDisable, bool) {
	if o == nil || IsNil(o.REMOVE_ACTION) {
		return nil, false
	}
	return o.REMOVE_ACTION, true
}

// HasREMOVE_ACTION returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationServiceNow) HasREMOVE_ACTION() bool {
	if o != nil && !IsNil(o.REMOVE_ACTION) {
		return true
	}

	return false
}

// SetREMOVE_ACTION gets a reference to the given EnumPropagationStoreTypeRemoveActionDisable and assigns it to the REMOVE_ACTION field.
func (o *PropagationStoreConfigurationServiceNow) SetREMOVE_ACTION(v EnumPropagationStoreTypeRemoveActionDisable) {
	o.REMOVE_ACTION = &v
}

// GetServiceNowUrl returns the ServiceNowUrl field value
func (o *PropagationStoreConfigurationServiceNow) GetServiceNowUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceNowUrl
}

// GetServiceNowUrlOk returns a tuple with the ServiceNowUrl field value
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationServiceNow) GetServiceNowUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceNowUrl, true
}

// SetServiceNowUrl sets field value
func (o *PropagationStoreConfigurationServiceNow) SetServiceNowUrl(v string) {
	o.ServiceNowUrl = v
}

// GetUPDATE_USERS returns the UPDATE_USERS field value if set, zero value otherwise.
func (o *PropagationStoreConfigurationServiceNow) GetUPDATE_USERS() bool {
	if o == nil || IsNil(o.UPDATE_USERS) {
		var ret bool
		return ret
	}
	return *o.UPDATE_USERS
}

// GetUPDATE_USERSOk returns a tuple with the UPDATE_USERS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropagationStoreConfigurationServiceNow) GetUPDATE_USERSOk() (*bool, bool) {
	if o == nil || IsNil(o.UPDATE_USERS) {
		return nil, false
	}
	return o.UPDATE_USERS, true
}

// HasUPDATE_USERS returns a boolean if a field has been set.
func (o *PropagationStoreConfigurationServiceNow) HasUPDATE_USERS() bool {
	if o != nil && !IsNil(o.UPDATE_USERS) {
		return true
	}

	return false
}

// SetUPDATE_USERS gets a reference to the given bool and assigns it to the UPDATE_USERS field.
func (o *PropagationStoreConfigurationServiceNow) SetUPDATE_USERS(v bool) {
	o.UPDATE_USERS = &v
}

func (o PropagationStoreConfigurationServiceNow) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PropagationStoreConfigurationServiceNow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["Administrator_Password"] = o.AdministratorPassword
	toSerialize["Administrator_Username"] = o.AdministratorUsername
	if !IsNil(o.CREATE_USERS) {
		toSerialize["CREATE_USERS"] = o.CREATE_USERS
	}
	if !IsNil(o.DEPROVISION_USERS) {
		toSerialize["DEPROVISION_USERS"] = o.DEPROVISION_USERS
	}
	if !IsNil(o.DISABLE_USERS) {
		toSerialize["DISABLE_USERS"] = o.DISABLE_USERS
	}
	if !IsNil(o.REMOVE_ACTION) {
		toSerialize["REMOVE_ACTION"] = o.REMOVE_ACTION
	}
	toSerialize["ServiceNow_Url"] = o.ServiceNowUrl
	if !IsNil(o.UPDATE_USERS) {
		toSerialize["UPDATE_USERS"] = o.UPDATE_USERS
	}
	return toSerialize, nil
}

type NullablePropagationStoreConfigurationServiceNow struct {
	value *PropagationStoreConfigurationServiceNow
	isSet bool
}

func (v NullablePropagationStoreConfigurationServiceNow) Get() *PropagationStoreConfigurationServiceNow {
	return v.value
}

func (v *NullablePropagationStoreConfigurationServiceNow) Set(val *PropagationStoreConfigurationServiceNow) {
	v.value = val
	v.isSet = true
}

func (v NullablePropagationStoreConfigurationServiceNow) IsSet() bool {
	return v.isSet
}

func (v *NullablePropagationStoreConfigurationServiceNow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropagationStoreConfigurationServiceNow(val *PropagationStoreConfigurationServiceNow) *NullablePropagationStoreConfigurationServiceNow {
	return &NullablePropagationStoreConfigurationServiceNow{value: val, isSet: true}
}

func (v NullablePropagationStoreConfigurationServiceNow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropagationStoreConfigurationServiceNow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

