# GatewayLDAPAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindDN** | **string** | A string that specifies the distinguished name information to bind to the LDAP database (for example, uid&#x3D;pingone,dc&#x3D;example,dc&#x3D;com). | 
**BindPassword** | **string** | A string that specifies the bind password for the LDAP database. This is a required property. | 
**ConnectionSecurity** | Pointer to **string** | (Optional) A string that specifies the connection security type. Options are None, TLS, and StartTLS. The default value is None. | [optional] 
**ServersHostAndPort** | Pointer to **[]string** | An array of strings that specifies the LDAP server host name and port number (for example, [\&quot;ds1.example.com:389\&quot;, \&quot;ds2.example.com:389\&quot;]). | [optional] 
**UserTypes** | [**[]GatewayLDAPAllOfUserTypes**](GatewayLDAPAllOfUserTypes.md) | (Optional) An array of the userTypes properties for the users to be provisioned in PingOne. userTypes specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory. | 
**ValidateTlsCertificates** | Pointer to **bool** | (Optional) A boolean that specifies whether or not to trust all SSL certificates (defaults to true). If this value is false, TLS certificates are not validated. When the value is set to true, only certificates that are signed by the default JVM CAs, or the CA certs that the customer has uploaded to the certificate service are trusted. | [optional] 
**Vendor** | **string** | A string that specifies the LDAP vendor. Options are PingDirectory, Microsoft Active Directory, Oracle Directory Server Enterprise Edition, Oracle Unified Directory, CA Directory, OpenDJ Directory, IBM (Tivoli) Security Directory Server, and LDAP v3 compliant Directory Server. | 

## Methods

### NewGatewayLDAPAllOf

`func NewGatewayLDAPAllOf(bindDN string, bindPassword string, userTypes []GatewayLDAPAllOfUserTypes, vendor string, ) *GatewayLDAPAllOf`

NewGatewayLDAPAllOf instantiates a new GatewayLDAPAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayLDAPAllOfWithDefaults

`func NewGatewayLDAPAllOfWithDefaults() *GatewayLDAPAllOf`

NewGatewayLDAPAllOfWithDefaults instantiates a new GatewayLDAPAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindDN

`func (o *GatewayLDAPAllOf) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *GatewayLDAPAllOf) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *GatewayLDAPAllOf) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.


### GetBindPassword

`func (o *GatewayLDAPAllOf) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *GatewayLDAPAllOf) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *GatewayLDAPAllOf) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.


### GetConnectionSecurity

`func (o *GatewayLDAPAllOf) GetConnectionSecurity() string`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *GatewayLDAPAllOf) GetConnectionSecurityOk() (*string, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *GatewayLDAPAllOf) SetConnectionSecurity(v string)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *GatewayLDAPAllOf) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetServersHostAndPort

`func (o *GatewayLDAPAllOf) GetServersHostAndPort() []string`

GetServersHostAndPort returns the ServersHostAndPort field if non-nil, zero value otherwise.

### GetServersHostAndPortOk

`func (o *GatewayLDAPAllOf) GetServersHostAndPortOk() (*[]string, bool)`

GetServersHostAndPortOk returns a tuple with the ServersHostAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersHostAndPort

`func (o *GatewayLDAPAllOf) SetServersHostAndPort(v []string)`

SetServersHostAndPort sets ServersHostAndPort field to given value.

### HasServersHostAndPort

`func (o *GatewayLDAPAllOf) HasServersHostAndPort() bool`

HasServersHostAndPort returns a boolean if a field has been set.

### GetUserTypes

`func (o *GatewayLDAPAllOf) GetUserTypes() []GatewayLDAPAllOfUserTypes`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *GatewayLDAPAllOf) GetUserTypesOk() (*[]GatewayLDAPAllOfUserTypes, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *GatewayLDAPAllOf) SetUserTypes(v []GatewayLDAPAllOfUserTypes)`

SetUserTypes sets UserTypes field to given value.


### GetValidateTlsCertificates

`func (o *GatewayLDAPAllOf) GetValidateTlsCertificates() bool`

GetValidateTlsCertificates returns the ValidateTlsCertificates field if non-nil, zero value otherwise.

### GetValidateTlsCertificatesOk

`func (o *GatewayLDAPAllOf) GetValidateTlsCertificatesOk() (*bool, bool)`

GetValidateTlsCertificatesOk returns a tuple with the ValidateTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateTlsCertificates

`func (o *GatewayLDAPAllOf) SetValidateTlsCertificates(v bool)`

SetValidateTlsCertificates sets ValidateTlsCertificates field to given value.

### HasValidateTlsCertificates

`func (o *GatewayLDAPAllOf) HasValidateTlsCertificates() bool`

HasValidateTlsCertificates returns a boolean if a field has been set.

### GetVendor

`func (o *GatewayLDAPAllOf) GetVendor() string`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GatewayLDAPAllOf) GetVendorOk() (*string, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GatewayLDAPAllOf) SetVendor(v string)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


