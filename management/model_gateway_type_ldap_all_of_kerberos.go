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

// checks if the GatewayTypeLDAPAllOfKerberos type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GatewayTypeLDAPAllOfKerberos{}

// GatewayTypeLDAPAllOfKerberos Contains the Kerberos authentication settings. Set this to null to disable Kerberos authentication.
type GatewayTypeLDAPAllOfKerberos struct {
	// The password for the Kerberos service account.
	ServiceAccountPassword *string `json:"serviceAccountPassword,omitempty"`
	// The Kerberos service account user principal name (for example, `username@domain.com`).
	ServiceAccountUserPrincipalName string `json:"serviceAccountUserPrincipalName"`
	// The number of minutes for which the previous credentials are persisted.
	MinutesToRetainPreviousCredentials *int32 `json:"minutesToRetainPreviousCredentials,omitempty"`
}

// NewGatewayTypeLDAPAllOfKerberos instantiates a new GatewayTypeLDAPAllOfKerberos object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayTypeLDAPAllOfKerberos(serviceAccountUserPrincipalName string) *GatewayTypeLDAPAllOfKerberos {
	this := GatewayTypeLDAPAllOfKerberos{}
	this.ServiceAccountUserPrincipalName = serviceAccountUserPrincipalName
	return &this
}

// NewGatewayTypeLDAPAllOfKerberosWithDefaults instantiates a new GatewayTypeLDAPAllOfKerberos object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayTypeLDAPAllOfKerberosWithDefaults() *GatewayTypeLDAPAllOfKerberos {
	this := GatewayTypeLDAPAllOfKerberos{}
	return &this
}

// GetServiceAccountPassword returns the ServiceAccountPassword field value if set, zero value otherwise.
func (o *GatewayTypeLDAPAllOfKerberos) GetServiceAccountPassword() string {
	if o == nil || IsNil(o.ServiceAccountPassword) {
		var ret string
		return ret
	}
	return *o.ServiceAccountPassword
}

// GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfKerberos) GetServiceAccountPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceAccountPassword) {
		return nil, false
	}
	return o.ServiceAccountPassword, true
}

// HasServiceAccountPassword returns a boolean if a field has been set.
func (o *GatewayTypeLDAPAllOfKerberos) HasServiceAccountPassword() bool {
	if o != nil && !IsNil(o.ServiceAccountPassword) {
		return true
	}

	return false
}

// SetServiceAccountPassword gets a reference to the given string and assigns it to the ServiceAccountPassword field.
func (o *GatewayTypeLDAPAllOfKerberos) SetServiceAccountPassword(v string) {
	o.ServiceAccountPassword = &v
}

// GetServiceAccountUserPrincipalName returns the ServiceAccountUserPrincipalName field value
func (o *GatewayTypeLDAPAllOfKerberos) GetServiceAccountUserPrincipalName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ServiceAccountUserPrincipalName
}

// GetServiceAccountUserPrincipalNameOk returns a tuple with the ServiceAccountUserPrincipalName field value
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfKerberos) GetServiceAccountUserPrincipalNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ServiceAccountUserPrincipalName, true
}

// SetServiceAccountUserPrincipalName sets field value
func (o *GatewayTypeLDAPAllOfKerberos) SetServiceAccountUserPrincipalName(v string) {
	o.ServiceAccountUserPrincipalName = v
}

// GetMinutesToRetainPreviousCredentials returns the MinutesToRetainPreviousCredentials field value if set, zero value otherwise.
func (o *GatewayTypeLDAPAllOfKerberos) GetMinutesToRetainPreviousCredentials() int32 {
	if o == nil || IsNil(o.MinutesToRetainPreviousCredentials) {
		var ret int32
		return ret
	}
	return *o.MinutesToRetainPreviousCredentials
}

// GetMinutesToRetainPreviousCredentialsOk returns a tuple with the MinutesToRetainPreviousCredentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayTypeLDAPAllOfKerberos) GetMinutesToRetainPreviousCredentialsOk() (*int32, bool) {
	if o == nil || IsNil(o.MinutesToRetainPreviousCredentials) {
		return nil, false
	}
	return o.MinutesToRetainPreviousCredentials, true
}

// HasMinutesToRetainPreviousCredentials returns a boolean if a field has been set.
func (o *GatewayTypeLDAPAllOfKerberos) HasMinutesToRetainPreviousCredentials() bool {
	if o != nil && !IsNil(o.MinutesToRetainPreviousCredentials) {
		return true
	}

	return false
}

// SetMinutesToRetainPreviousCredentials gets a reference to the given int32 and assigns it to the MinutesToRetainPreviousCredentials field.
func (o *GatewayTypeLDAPAllOfKerberos) SetMinutesToRetainPreviousCredentials(v int32) {
	o.MinutesToRetainPreviousCredentials = &v
}

func (o GatewayTypeLDAPAllOfKerberos) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GatewayTypeLDAPAllOfKerberos) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServiceAccountPassword) {
		toSerialize["serviceAccountPassword"] = o.ServiceAccountPassword
	}
	toSerialize["serviceAccountUserPrincipalName"] = o.ServiceAccountUserPrincipalName
	if !IsNil(o.MinutesToRetainPreviousCredentials) {
		toSerialize["minutesToRetainPreviousCredentials"] = o.MinutesToRetainPreviousCredentials
	}
	return toSerialize, nil
}

type NullableGatewayTypeLDAPAllOfKerberos struct {
	value *GatewayTypeLDAPAllOfKerberos
	isSet bool
}

func (v NullableGatewayTypeLDAPAllOfKerberos) Get() *GatewayTypeLDAPAllOfKerberos {
	return v.value
}

func (v *NullableGatewayTypeLDAPAllOfKerberos) Set(val *GatewayTypeLDAPAllOfKerberos) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayTypeLDAPAllOfKerberos) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayTypeLDAPAllOfKerberos) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayTypeLDAPAllOfKerberos(val *GatewayTypeLDAPAllOfKerberos) *NullableGatewayTypeLDAPAllOfKerberos {
	return &NullableGatewayTypeLDAPAllOfKerberos{value: val, isSet: true}
}

func (v NullableGatewayTypeLDAPAllOfKerberos) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayTypeLDAPAllOfKerberos) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
