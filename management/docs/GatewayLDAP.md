# GatewayLDAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the instance ID of the gateway. The gateway instance ID is created by the gateway when it starts up. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Credentials** | Pointer to [**[]GatewayCredential**](GatewayCredential.md) |  | [optional] 
**Name** | **string** | A string that specifies the resource name, which must be provided and must be unique within an environment. Valid characters are any Unicode letter, mark, numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. | 
**Description** | Pointer to **string** | (Optional) A string that specifies the description of the resource. | [optional] 
**Type** | **string** | A string that specifies the type of gateway resource. Options are LDAP, PING_FEDERATE and PING_INTELLIGENCE. This is a required property. | 
**Enabled** | **bool** | A boolean that specifies whether the gateway is enabled. This is a required property. | 
**SupportedVersions** | Pointer to [**GatewaySupportedVersions**](GatewaySupportedVersions.md) |  | [optional] 
**BindDN** | **string** | A string that specifies the distinguished name information to bind to the LDAP database (for example, uid&#x3D;pingone,dc&#x3D;example,dc&#x3D;com). | 
**BindPassword** | **string** | A string that specifies the bind password for the LDAP database. This is a required property. | 
**ConnectionSecurity** | Pointer to **string** | (Optional) A string that specifies the connection security type. Options are None, TLS, and StartTLS. The default value is None. | [optional] 
**ServersHostAndPort** | Pointer to **[]string** | An array of strings that specifies the LDAP server host name and port number (for example, [\&quot;ds1.example.com:389\&quot;, \&quot;ds2.example.com:389\&quot;]). | [optional] 
**UserTypes** | [**[]GatewayLDAPAllOfUserTypes**](GatewayLDAPAllOfUserTypes.md) | (Optional) An array of the userTypes properties for the users to be provisioned in PingOne. userTypes specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory. | 
**ValidateTlsCertificates** | Pointer to **bool** | (Optional) A boolean that specifies whether or not to trust all SSL certificates (defaults to true). If this value is false, TLS certificates are not validated. When the value is set to true, only certificates that are signed by the default JVM CAs, or the CA certs that the customer has uploaded to the certificate service are trusted. | [optional] 
**Vendor** | **string** | A string that specifies the LDAP vendor. Options are PingDirectory, Microsoft Active Directory, Oracle Directory Server Enterprise Edition, Oracle Unified Directory, CA Directory, OpenDJ Directory, IBM (Tivoli) Security Directory Server, and LDAP v3 compliant Directory Server. | 

## Methods

### NewGatewayLDAP

`func NewGatewayLDAP(name string, type_ string, enabled bool, bindDN string, bindPassword string, userTypes []GatewayLDAPAllOfUserTypes, vendor string, ) *GatewayLDAP`

NewGatewayLDAP instantiates a new GatewayLDAP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayLDAPWithDefaults

`func NewGatewayLDAPWithDefaults() *GatewayLDAP`

NewGatewayLDAPWithDefaults instantiates a new GatewayLDAP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GatewayLDAP) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayLDAP) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayLDAP) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayLDAP) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *GatewayLDAP) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GatewayLDAP) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GatewayLDAP) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GatewayLDAP) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCredentials

`func (o *GatewayLDAP) GetCredentials() []GatewayCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GatewayLDAP) GetCredentialsOk() (*[]GatewayCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GatewayLDAP) SetCredentials(v []GatewayCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GatewayLDAP) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *GatewayLDAP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayLDAP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayLDAP) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GatewayLDAP) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayLDAP) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayLDAP) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayLDAP) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GatewayLDAP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayLDAP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayLDAP) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *GatewayLDAP) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayLDAP) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayLDAP) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSupportedVersions

`func (o *GatewayLDAP) GetSupportedVersions() GatewaySupportedVersions`

GetSupportedVersions returns the SupportedVersions field if non-nil, zero value otherwise.

### GetSupportedVersionsOk

`func (o *GatewayLDAP) GetSupportedVersionsOk() (*GatewaySupportedVersions, bool)`

GetSupportedVersionsOk returns a tuple with the SupportedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVersions

`func (o *GatewayLDAP) SetSupportedVersions(v GatewaySupportedVersions)`

SetSupportedVersions sets SupportedVersions field to given value.

### HasSupportedVersions

`func (o *GatewayLDAP) HasSupportedVersions() bool`

HasSupportedVersions returns a boolean if a field has been set.

### GetBindDN

`func (o *GatewayLDAP) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *GatewayLDAP) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *GatewayLDAP) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.


### GetBindPassword

`func (o *GatewayLDAP) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *GatewayLDAP) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *GatewayLDAP) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.


### GetConnectionSecurity

`func (o *GatewayLDAP) GetConnectionSecurity() string`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *GatewayLDAP) GetConnectionSecurityOk() (*string, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *GatewayLDAP) SetConnectionSecurity(v string)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *GatewayLDAP) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetServersHostAndPort

`func (o *GatewayLDAP) GetServersHostAndPort() []string`

GetServersHostAndPort returns the ServersHostAndPort field if non-nil, zero value otherwise.

### GetServersHostAndPortOk

`func (o *GatewayLDAP) GetServersHostAndPortOk() (*[]string, bool)`

GetServersHostAndPortOk returns a tuple with the ServersHostAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersHostAndPort

`func (o *GatewayLDAP) SetServersHostAndPort(v []string)`

SetServersHostAndPort sets ServersHostAndPort field to given value.

### HasServersHostAndPort

`func (o *GatewayLDAP) HasServersHostAndPort() bool`

HasServersHostAndPort returns a boolean if a field has been set.

### GetUserTypes

`func (o *GatewayLDAP) GetUserTypes() []GatewayLDAPAllOfUserTypes`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *GatewayLDAP) GetUserTypesOk() (*[]GatewayLDAPAllOfUserTypes, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *GatewayLDAP) SetUserTypes(v []GatewayLDAPAllOfUserTypes)`

SetUserTypes sets UserTypes field to given value.


### GetValidateTlsCertificates

`func (o *GatewayLDAP) GetValidateTlsCertificates() bool`

GetValidateTlsCertificates returns the ValidateTlsCertificates field if non-nil, zero value otherwise.

### GetValidateTlsCertificatesOk

`func (o *GatewayLDAP) GetValidateTlsCertificatesOk() (*bool, bool)`

GetValidateTlsCertificatesOk returns a tuple with the ValidateTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateTlsCertificates

`func (o *GatewayLDAP) SetValidateTlsCertificates(v bool)`

SetValidateTlsCertificates sets ValidateTlsCertificates field to given value.

### HasValidateTlsCertificates

`func (o *GatewayLDAP) HasValidateTlsCertificates() bool`

HasValidateTlsCertificates returns a boolean if a field has been set.

### GetVendor

`func (o *GatewayLDAP) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GatewayLDAP) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GatewayLDAP) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


