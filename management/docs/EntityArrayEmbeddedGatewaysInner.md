# EntityArrayEmbeddedGatewaysInner

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

### NewEntityArrayEmbeddedGatewaysInner

`func NewEntityArrayEmbeddedGatewaysInner(name string, type_ string, enabled bool, bindDN string, bindPassword string, userTypes []GatewayLDAPAllOfUserTypes, vendor string, ) *EntityArrayEmbeddedGatewaysInner`

NewEntityArrayEmbeddedGatewaysInner instantiates a new EntityArrayEmbeddedGatewaysInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedGatewaysInnerWithDefaults

`func NewEntityArrayEmbeddedGatewaysInnerWithDefaults() *EntityArrayEmbeddedGatewaysInner`

NewEntityArrayEmbeddedGatewaysInnerWithDefaults instantiates a new EntityArrayEmbeddedGatewaysInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityArrayEmbeddedGatewaysInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityArrayEmbeddedGatewaysInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntityArrayEmbeddedGatewaysInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *EntityArrayEmbeddedGatewaysInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EntityArrayEmbeddedGatewaysInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EntityArrayEmbeddedGatewaysInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCredentials

`func (o *EntityArrayEmbeddedGatewaysInner) GetCredentials() []GatewayCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetCredentialsOk() (*[]GatewayCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *EntityArrayEmbeddedGatewaysInner) SetCredentials(v []GatewayCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *EntityArrayEmbeddedGatewaysInner) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *EntityArrayEmbeddedGatewaysInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityArrayEmbeddedGatewaysInner) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntityArrayEmbeddedGatewaysInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityArrayEmbeddedGatewaysInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityArrayEmbeddedGatewaysInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *EntityArrayEmbeddedGatewaysInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityArrayEmbeddedGatewaysInner) SetType(v string)`

SetType sets Type field to given value.


### GetEnabled

`func (o *EntityArrayEmbeddedGatewaysInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EntityArrayEmbeddedGatewaysInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSupportedVersions

`func (o *EntityArrayEmbeddedGatewaysInner) GetSupportedVersions() GatewaySupportedVersions`

GetSupportedVersions returns the SupportedVersions field if non-nil, zero value otherwise.

### GetSupportedVersionsOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetSupportedVersionsOk() (*GatewaySupportedVersions, bool)`

GetSupportedVersionsOk returns a tuple with the SupportedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVersions

`func (o *EntityArrayEmbeddedGatewaysInner) SetSupportedVersions(v GatewaySupportedVersions)`

SetSupportedVersions sets SupportedVersions field to given value.

### HasSupportedVersions

`func (o *EntityArrayEmbeddedGatewaysInner) HasSupportedVersions() bool`

HasSupportedVersions returns a boolean if a field has been set.

### GetBindDN

`func (o *EntityArrayEmbeddedGatewaysInner) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *EntityArrayEmbeddedGatewaysInner) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.


### GetBindPassword

`func (o *EntityArrayEmbeddedGatewaysInner) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *EntityArrayEmbeddedGatewaysInner) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.


### GetConnectionSecurity

`func (o *EntityArrayEmbeddedGatewaysInner) GetConnectionSecurity() string`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetConnectionSecurityOk() (*string, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *EntityArrayEmbeddedGatewaysInner) SetConnectionSecurity(v string)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *EntityArrayEmbeddedGatewaysInner) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetServersHostAndPort

`func (o *EntityArrayEmbeddedGatewaysInner) GetServersHostAndPort() []string`

GetServersHostAndPort returns the ServersHostAndPort field if non-nil, zero value otherwise.

### GetServersHostAndPortOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetServersHostAndPortOk() (*[]string, bool)`

GetServersHostAndPortOk returns a tuple with the ServersHostAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersHostAndPort

`func (o *EntityArrayEmbeddedGatewaysInner) SetServersHostAndPort(v []string)`

SetServersHostAndPort sets ServersHostAndPort field to given value.

### HasServersHostAndPort

`func (o *EntityArrayEmbeddedGatewaysInner) HasServersHostAndPort() bool`

HasServersHostAndPort returns a boolean if a field has been set.

### GetUserTypes

`func (o *EntityArrayEmbeddedGatewaysInner) GetUserTypes() []GatewayLDAPAllOfUserTypes`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetUserTypesOk() (*[]GatewayLDAPAllOfUserTypes, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *EntityArrayEmbeddedGatewaysInner) SetUserTypes(v []GatewayLDAPAllOfUserTypes)`

SetUserTypes sets UserTypes field to given value.


### GetValidateTlsCertificates

`func (o *EntityArrayEmbeddedGatewaysInner) GetValidateTlsCertificates() bool`

GetValidateTlsCertificates returns the ValidateTlsCertificates field if non-nil, zero value otherwise.

### GetValidateTlsCertificatesOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetValidateTlsCertificatesOk() (*bool, bool)`

GetValidateTlsCertificatesOk returns a tuple with the ValidateTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateTlsCertificates

`func (o *EntityArrayEmbeddedGatewaysInner) SetValidateTlsCertificates(v bool)`

SetValidateTlsCertificates sets ValidateTlsCertificates field to given value.

### HasValidateTlsCertificates

`func (o *EntityArrayEmbeddedGatewaysInner) HasValidateTlsCertificates() bool`

HasValidateTlsCertificates returns a boolean if a field has been set.

### GetVendor

`func (o *EntityArrayEmbeddedGatewaysInner) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *EntityArrayEmbeddedGatewaysInner) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *EntityArrayEmbeddedGatewaysInner) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


