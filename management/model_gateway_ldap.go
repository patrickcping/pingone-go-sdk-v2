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

// GatewayLDAP struct for GatewayLDAP
type GatewayLDAP struct {
	// A string that specifies the instance ID of the gateway. The gateway instance ID is created by the gateway when it starts up.
	Id *string `json:"id,omitempty"`
	Environment *ObjectEnvironment `json:"environment,omitempty"`
	Credentials []GatewayCredential `json:"credentials,omitempty"`
	// A string that specifies the resource name, which must be provided and must be unique within an environment. Valid characters are any Unicode letter, mark, numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen.
	Name string `json:"name"`
	// (Optional) A string that specifies the description of the resource.
	Description *string `json:"description,omitempty"`
	Type EnumGatewayType `json:"type"`
	// A boolean that specifies whether the gateway is enabled. This is a required property.
	Enabled bool `json:"enabled"`
	SupportedVersions *GatewaySupportedVersions `json:"supportedVersions,omitempty"`
	// A string that specifies the distinguished name information to bind to the LDAP database (for example, uid=pingone,dc=example,dc=com).
	BindDN string `json:"bindDN"`
	// A string that specifies the bind password for the LDAP database. This is a required property.
	BindPassword string `json:"bindPassword"`
	ConnectionSecurity *EnumGatewayLDAPSecurity `json:"connectionSecurity,omitempty"`
	// An array of strings that specifies the LDAP server host name and port number (for example, [\"ds1.example.com:389\", \"ds2.example.com:389\"]).
	ServersHostAndPort []string `json:"serversHostAndPort,omitempty"`
	// (Optional) An array of the userTypes properties for the users to be provisioned in PingOne. userTypes specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory.
	UserTypes []GatewayLDAPAllOfUserTypes `json:"userTypes"`
	// (Optional) A boolean that specifies whether or not to trust all SSL certificates (defaults to true). If this value is false, TLS certificates are not validated. When the value is set to true, only certificates that are signed by the default JVM CAs, or the CA certs that the customer has uploaded to the certificate service are trusted.
	ValidateTlsCertificates *bool `json:"validateTlsCertificates,omitempty"`
	// A string that specifies the LDAP vendor. Options are PingDirectory, Microsoft Active Directory, Oracle Directory Server Enterprise Edition, Oracle Unified Directory, CA Directory, OpenDJ Directory, IBM (Tivoli) Security Directory Server, and LDAP v3 compliant Directory Server.
	Vendor string `json:"vendor"`
}

// NewGatewayLDAP instantiates a new GatewayLDAP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGatewayLDAP(name string, type_ EnumGatewayType, enabled bool, bindDN string, bindPassword string, userTypes []GatewayLDAPAllOfUserTypes, vendor string) *GatewayLDAP {
	this := GatewayLDAP{}
	this.Name = name
	this.Type = type_
	this.Enabled = enabled
	this.BindDN = bindDN
	this.BindPassword = bindPassword
	this.UserTypes = userTypes
	this.Vendor = vendor
	return &this
}

// NewGatewayLDAPWithDefaults instantiates a new GatewayLDAP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGatewayLDAPWithDefaults() *GatewayLDAP {
	this := GatewayLDAP{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GatewayLDAP) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GatewayLDAP) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GatewayLDAP) SetId(v string) {
	o.Id = &v
}

// GetEnvironment returns the Environment field value if set, zero value otherwise.
func (o *GatewayLDAP) GetEnvironment() ObjectEnvironment {
	if o == nil || o.Environment == nil {
		var ret ObjectEnvironment
		return ret
	}
	return *o.Environment
}

// GetEnvironmentOk returns a tuple with the Environment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetEnvironmentOk() (*ObjectEnvironment, bool) {
	if o == nil || o.Environment == nil {
		return nil, false
	}
	return o.Environment, true
}

// HasEnvironment returns a boolean if a field has been set.
func (o *GatewayLDAP) HasEnvironment() bool {
	if o != nil && o.Environment != nil {
		return true
	}

	return false
}

// SetEnvironment gets a reference to the given ObjectEnvironment and assigns it to the Environment field.
func (o *GatewayLDAP) SetEnvironment(v ObjectEnvironment) {
	o.Environment = &v
}

// GetCredentials returns the Credentials field value if set, zero value otherwise.
func (o *GatewayLDAP) GetCredentials() []GatewayCredential {
	if o == nil || o.Credentials == nil {
		var ret []GatewayCredential
		return ret
	}
	return o.Credentials
}

// GetCredentialsOk returns a tuple with the Credentials field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetCredentialsOk() ([]GatewayCredential, bool) {
	if o == nil || o.Credentials == nil {
		return nil, false
	}
	return o.Credentials, true
}

// HasCredentials returns a boolean if a field has been set.
func (o *GatewayLDAP) HasCredentials() bool {
	if o != nil && o.Credentials != nil {
		return true
	}

	return false
}

// SetCredentials gets a reference to the given []GatewayCredential and assigns it to the Credentials field.
func (o *GatewayLDAP) SetCredentials(v []GatewayCredential) {
	o.Credentials = v
}

// GetName returns the Name field value
func (o *GatewayLDAP) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *GatewayLDAP) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GatewayLDAP) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GatewayLDAP) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GatewayLDAP) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value
func (o *GatewayLDAP) GetType() EnumGatewayType {
	if o == nil {
		var ret EnumGatewayType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetTypeOk() (*EnumGatewayType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GatewayLDAP) SetType(v EnumGatewayType) {
	o.Type = v
}

// GetEnabled returns the Enabled field value
func (o *GatewayLDAP) GetEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Enabled, true
}

// SetEnabled sets field value
func (o *GatewayLDAP) SetEnabled(v bool) {
	o.Enabled = v
}

// GetSupportedVersions returns the SupportedVersions field value if set, zero value otherwise.
func (o *GatewayLDAP) GetSupportedVersions() GatewaySupportedVersions {
	if o == nil || o.SupportedVersions == nil {
		var ret GatewaySupportedVersions
		return ret
	}
	return *o.SupportedVersions
}

// GetSupportedVersionsOk returns a tuple with the SupportedVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetSupportedVersionsOk() (*GatewaySupportedVersions, bool) {
	if o == nil || o.SupportedVersions == nil {
		return nil, false
	}
	return o.SupportedVersions, true
}

// HasSupportedVersions returns a boolean if a field has been set.
func (o *GatewayLDAP) HasSupportedVersions() bool {
	if o != nil && o.SupportedVersions != nil {
		return true
	}

	return false
}

// SetSupportedVersions gets a reference to the given GatewaySupportedVersions and assigns it to the SupportedVersions field.
func (o *GatewayLDAP) SetSupportedVersions(v GatewaySupportedVersions) {
	o.SupportedVersions = &v
}

// GetBindDN returns the BindDN field value
func (o *GatewayLDAP) GetBindDN() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindDN
}

// GetBindDNOk returns a tuple with the BindDN field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetBindDNOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindDN, true
}

// SetBindDN sets field value
func (o *GatewayLDAP) SetBindDN(v string) {
	o.BindDN = v
}

// GetBindPassword returns the BindPassword field value
func (o *GatewayLDAP) GetBindPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BindPassword
}

// GetBindPasswordOk returns a tuple with the BindPassword field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetBindPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BindPassword, true
}

// SetBindPassword sets field value
func (o *GatewayLDAP) SetBindPassword(v string) {
	o.BindPassword = v
}

// GetConnectionSecurity returns the ConnectionSecurity field value if set, zero value otherwise.
func (o *GatewayLDAP) GetConnectionSecurity() EnumGatewayLDAPSecurity {
	if o == nil || o.ConnectionSecurity == nil {
		var ret EnumGatewayLDAPSecurity
		return ret
	}
	return *o.ConnectionSecurity
}

// GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetConnectionSecurityOk() (*EnumGatewayLDAPSecurity, bool) {
	if o == nil || o.ConnectionSecurity == nil {
		return nil, false
	}
	return o.ConnectionSecurity, true
}

// HasConnectionSecurity returns a boolean if a field has been set.
func (o *GatewayLDAP) HasConnectionSecurity() bool {
	if o != nil && o.ConnectionSecurity != nil {
		return true
	}

	return false
}

// SetConnectionSecurity gets a reference to the given EnumGatewayLDAPSecurity and assigns it to the ConnectionSecurity field.
func (o *GatewayLDAP) SetConnectionSecurity(v EnumGatewayLDAPSecurity) {
	o.ConnectionSecurity = &v
}

// GetServersHostAndPort returns the ServersHostAndPort field value if set, zero value otherwise.
func (o *GatewayLDAP) GetServersHostAndPort() []string {
	if o == nil || o.ServersHostAndPort == nil {
		var ret []string
		return ret
	}
	return o.ServersHostAndPort
}

// GetServersHostAndPortOk returns a tuple with the ServersHostAndPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetServersHostAndPortOk() ([]string, bool) {
	if o == nil || o.ServersHostAndPort == nil {
		return nil, false
	}
	return o.ServersHostAndPort, true
}

// HasServersHostAndPort returns a boolean if a field has been set.
func (o *GatewayLDAP) HasServersHostAndPort() bool {
	if o != nil && o.ServersHostAndPort != nil {
		return true
	}

	return false
}

// SetServersHostAndPort gets a reference to the given []string and assigns it to the ServersHostAndPort field.
func (o *GatewayLDAP) SetServersHostAndPort(v []string) {
	o.ServersHostAndPort = v
}

// GetUserTypes returns the UserTypes field value
func (o *GatewayLDAP) GetUserTypes() []GatewayLDAPAllOfUserTypes {
	if o == nil {
		var ret []GatewayLDAPAllOfUserTypes
		return ret
	}

	return o.UserTypes
}

// GetUserTypesOk returns a tuple with the UserTypes field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetUserTypesOk() ([]GatewayLDAPAllOfUserTypes, bool) {
	if o == nil {
		return nil, false
	}
	return o.UserTypes, true
}

// SetUserTypes sets field value
func (o *GatewayLDAP) SetUserTypes(v []GatewayLDAPAllOfUserTypes) {
	o.UserTypes = v
}

// GetValidateTlsCertificates returns the ValidateTlsCertificates field value if set, zero value otherwise.
func (o *GatewayLDAP) GetValidateTlsCertificates() bool {
	if o == nil || o.ValidateTlsCertificates == nil {
		var ret bool
		return ret
	}
	return *o.ValidateTlsCertificates
}

// GetValidateTlsCertificatesOk returns a tuple with the ValidateTlsCertificates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetValidateTlsCertificatesOk() (*bool, bool) {
	if o == nil || o.ValidateTlsCertificates == nil {
		return nil, false
	}
	return o.ValidateTlsCertificates, true
}

// HasValidateTlsCertificates returns a boolean if a field has been set.
func (o *GatewayLDAP) HasValidateTlsCertificates() bool {
	if o != nil && o.ValidateTlsCertificates != nil {
		return true
	}

	return false
}

// SetValidateTlsCertificates gets a reference to the given bool and assigns it to the ValidateTlsCertificates field.
func (o *GatewayLDAP) SetValidateTlsCertificates(v bool) {
	o.ValidateTlsCertificates = &v
}

// GetVendor returns the Vendor field value
func (o *GatewayLDAP) GetVendor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vendor
}

// GetVendorOk returns a tuple with the Vendor field value
// and a boolean to check if the value has been set.
func (o *GatewayLDAP) GetVendorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vendor, true
}

// SetVendor sets field value
func (o *GatewayLDAP) SetVendor(v string) {
	o.Vendor = v
}

func (o GatewayLDAP) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Environment != nil {
		toSerialize["environment"] = o.Environment
	}
	if o.Credentials != nil {
		toSerialize["credentials"] = o.Credentials
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["enabled"] = o.Enabled
	}
	if o.SupportedVersions != nil {
		toSerialize["supportedVersions"] = o.SupportedVersions
	}
	if true {
		toSerialize["bindDN"] = o.BindDN
	}
	if true {
		toSerialize["bindPassword"] = o.BindPassword
	}
	if o.ConnectionSecurity != nil {
		toSerialize["connectionSecurity"] = o.ConnectionSecurity
	}
	if o.ServersHostAndPort != nil {
		toSerialize["serversHostAndPort"] = o.ServersHostAndPort
	}
	if true {
		toSerialize["userTypes"] = o.UserTypes
	}
	if o.ValidateTlsCertificates != nil {
		toSerialize["validateTlsCertificates"] = o.ValidateTlsCertificates
	}
	if true {
		toSerialize["vendor"] = o.Vendor
	}
	return json.Marshal(toSerialize)
}

type NullableGatewayLDAP struct {
	value *GatewayLDAP
	isSet bool
}

func (v NullableGatewayLDAP) Get() *GatewayLDAP {
	return v.value
}

func (v *NullableGatewayLDAP) Set(val *GatewayLDAP) {
	v.value = val
	v.isSet = true
}

func (v NullableGatewayLDAP) IsSet() bool {
	return v.isSet
}

func (v *NullableGatewayLDAP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGatewayLDAP(val *GatewayLDAP) *NullableGatewayLDAP {
	return &NullableGatewayLDAP{value: val, isSet: true}
}

func (v NullableGatewayLDAP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGatewayLDAP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


