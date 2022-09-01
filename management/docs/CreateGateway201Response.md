# CreateGateway201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the instance ID of the gateway. The gateway instance ID is created by the gateway when it starts up. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Credentials** | Pointer to [**[]GatewayCredential**](GatewayCredential.md) |  | [optional] 
**Name** | **string** | A string that specifies the resource name, which must be provided and must be unique within an environment. Valid characters are any Unicode letter, mark, numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. | 
**Description** | Pointer to **string** | A string that specifies the description of the resource. | [optional] 
**Type** | [**EnumGatewayType**](EnumGatewayType.md) |  | 
**Enabled** | **bool** | A boolean that specifies whether the gateway is enabled. This is a required property. | 
**SupportedVersions** | Pointer to [**GatewaySupportedVersions**](GatewaySupportedVersions.md) |  | [optional] 
**BindDN** | **string** | A string that specifies the distinguished name information to bind to the LDAP database (for example, uid&#x3D;pingone,dc&#x3D;example,dc&#x3D;com). | 
**BindPassword** | **string** | A string that specifies the bind password for the LDAP database. This is a required property. | 
**ConnectionSecurity** | Pointer to [**EnumGatewayLDAPSecurity**](EnumGatewayLDAPSecurity.md) |  | [optional] 
**Kerberos** | Pointer to [**GatewayLDAPAllOfKerberos**](GatewayLDAPAllOfKerberos.md) |  | [optional] 
**ServersHostAndPort** | Pointer to **[]string** | An array of strings that specifies the LDAP server host name and port number (for example, [&#x60;ds1.example.com:389&#x60;, &#x60;ds2.example.com:389&#x60;]). | [optional] 
**UserTypes** | [**[]GatewayLDAPAllOfUserTypes**](GatewayLDAPAllOfUserTypes.md) | An array of the userTypes properties for the users to be provisioned in PingOne. userTypes specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory. | 
**ValidateTlsCertificates** | Pointer to **bool** | A boolean that specifies whether or not to trust all SSL certificates (defaults to true). If this value is false, TLS certificates are not validated. When the value is set to true, only certificates that are signed by the default JVM CAs, or the CA certs that the customer has uploaded to the certificate service are trusted. | [optional] 
**Vendor** | [**EnumGatewayVendor**](EnumGatewayVendor.md) |  | 

## Methods

### NewCreateGateway201Response

`func NewCreateGateway201Response(name string, type_ EnumGatewayType, enabled bool, bindDN string, bindPassword string, userTypes []GatewayLDAPAllOfUserTypes, vendor EnumGatewayVendor, ) *CreateGateway201Response`

NewCreateGateway201Response instantiates a new CreateGateway201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGateway201ResponseWithDefaults

`func NewCreateGateway201ResponseWithDefaults() *CreateGateway201Response`

NewCreateGateway201ResponseWithDefaults instantiates a new CreateGateway201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CreateGateway201Response) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateGateway201Response) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateGateway201Response) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateGateway201Response) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *CreateGateway201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateGateway201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateGateway201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateGateway201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateGateway201Response) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateGateway201Response) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateGateway201Response) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateGateway201Response) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateGateway201Response) GetCredentials() []GatewayCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateGateway201Response) GetCredentialsOk() (*[]GatewayCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateGateway201Response) SetCredentials(v []GatewayCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateGateway201Response) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *CreateGateway201Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGateway201Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGateway201Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateGateway201Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGateway201Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGateway201Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGateway201Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateGateway201Response) GetType() EnumGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateGateway201Response) GetTypeOk() (*EnumGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateGateway201Response) SetType(v EnumGatewayType)`

SetType sets Type field to given value.


### GetEnabled

`func (o *CreateGateway201Response) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateGateway201Response) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateGateway201Response) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSupportedVersions

`func (o *CreateGateway201Response) GetSupportedVersions() GatewaySupportedVersions`

GetSupportedVersions returns the SupportedVersions field if non-nil, zero value otherwise.

### GetSupportedVersionsOk

`func (o *CreateGateway201Response) GetSupportedVersionsOk() (*GatewaySupportedVersions, bool)`

GetSupportedVersionsOk returns a tuple with the SupportedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVersions

`func (o *CreateGateway201Response) SetSupportedVersions(v GatewaySupportedVersions)`

SetSupportedVersions sets SupportedVersions field to given value.

### HasSupportedVersions

`func (o *CreateGateway201Response) HasSupportedVersions() bool`

HasSupportedVersions returns a boolean if a field has been set.

### GetBindDN

`func (o *CreateGateway201Response) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *CreateGateway201Response) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *CreateGateway201Response) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.


### GetBindPassword

`func (o *CreateGateway201Response) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *CreateGateway201Response) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *CreateGateway201Response) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.


### GetConnectionSecurity

`func (o *CreateGateway201Response) GetConnectionSecurity() EnumGatewayLDAPSecurity`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *CreateGateway201Response) GetConnectionSecurityOk() (*EnumGatewayLDAPSecurity, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *CreateGateway201Response) SetConnectionSecurity(v EnumGatewayLDAPSecurity)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *CreateGateway201Response) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetKerberos

`func (o *CreateGateway201Response) GetKerberos() GatewayLDAPAllOfKerberos`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *CreateGateway201Response) GetKerberosOk() (*GatewayLDAPAllOfKerberos, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *CreateGateway201Response) SetKerberos(v GatewayLDAPAllOfKerberos)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *CreateGateway201Response) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetServersHostAndPort

`func (o *CreateGateway201Response) GetServersHostAndPort() []string`

GetServersHostAndPort returns the ServersHostAndPort field if non-nil, zero value otherwise.

### GetServersHostAndPortOk

`func (o *CreateGateway201Response) GetServersHostAndPortOk() (*[]string, bool)`

GetServersHostAndPortOk returns a tuple with the ServersHostAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersHostAndPort

`func (o *CreateGateway201Response) SetServersHostAndPort(v []string)`

SetServersHostAndPort sets ServersHostAndPort field to given value.

### HasServersHostAndPort

`func (o *CreateGateway201Response) HasServersHostAndPort() bool`

HasServersHostAndPort returns a boolean if a field has been set.

### GetUserTypes

`func (o *CreateGateway201Response) GetUserTypes() []GatewayLDAPAllOfUserTypes`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *CreateGateway201Response) GetUserTypesOk() (*[]GatewayLDAPAllOfUserTypes, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *CreateGateway201Response) SetUserTypes(v []GatewayLDAPAllOfUserTypes)`

SetUserTypes sets UserTypes field to given value.


### GetValidateTlsCertificates

`func (o *CreateGateway201Response) GetValidateTlsCertificates() bool`

GetValidateTlsCertificates returns the ValidateTlsCertificates field if non-nil, zero value otherwise.

### GetValidateTlsCertificatesOk

`func (o *CreateGateway201Response) GetValidateTlsCertificatesOk() (*bool, bool)`

GetValidateTlsCertificatesOk returns a tuple with the ValidateTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateTlsCertificates

`func (o *CreateGateway201Response) SetValidateTlsCertificates(v bool)`

SetValidateTlsCertificates sets ValidateTlsCertificates field to given value.

### HasValidateTlsCertificates

`func (o *CreateGateway201Response) HasValidateTlsCertificates() bool`

HasValidateTlsCertificates returns a boolean if a field has been set.

### GetVendor

`func (o *CreateGateway201Response) GetVendor() EnumGatewayVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CreateGateway201Response) GetVendorOk() (*EnumGatewayVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CreateGateway201Response) SetVendor(v EnumGatewayVendor)`

SetVendor sets Vendor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


