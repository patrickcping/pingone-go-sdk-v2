# CreateGatewayRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the instance ID of the gateway. The gateway instance ID is created by the gateway when it starts up. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Credentials** | Pointer to [**[]GatewayCredential**](GatewayCredential.md) |  | [optional] [readonly] 
**Name** | **string** | A string that specifies the resource name, which must be provided and must be unique within an environment. Valid characters are any Unicode letter, mark, numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. | 
**Description** | Pointer to **string** | A string that specifies the description of the resource. | [optional] 
**Type** | [**EnumGatewayType**](EnumGatewayType.md) |  | 
**Enabled** | **bool** | A boolean that specifies whether the gateway is enabled. This is a required property. | 
**SupportedVersions** | Pointer to [**[]GatewaySupportedVersionsInner**](GatewaySupportedVersionsInner.md) | An array that lists the LDAP gateway versions associated with this gateway resource. This information is returned on a GET {{apiPath}}/environments/{{environmentID}}/gateways request, and it is used to trigger alerts if the gateway tries to connect with an unsupported version (or a version that is not the latest or recommended version). | [optional] [readonly] 
**CurrentAlerts** | Pointer to **[]map[string]interface{}** |  | [optional] [readonly] 
**BindDN** | **string** | A string that specifies the distinguished name information to bind to the LDAP database (for example, uid&#x3D;pingone,dc&#x3D;example,dc&#x3D;com). | 
**BindPassword** | **string** | A string that specifies the bind password for the LDAP database. This is a required property. | 
**ConnectionSecurity** | Pointer to [**EnumGatewayTypeLDAPSecurity**](EnumGatewayTypeLDAPSecurity.md) |  | [optional] [default to ENUMGATEWAYTYPELDAPSECURITY_NONE]
**Kerberos** | Pointer to [**GatewayTypeLDAPAllOfKerberos**](GatewayTypeLDAPAllOfKerberos.md) |  | [optional] 
**ServersHostAndPort** | **[]string** | An array of strings that specifies the LDAP server host name and port number (for example, [&#x60;ds1.example.com:389&#x60;, &#x60;ds2.example.com:389&#x60;]). | 
**UserTypes** | Pointer to [**[]GatewayTypeLDAPAllOfUserTypes**](GatewayTypeLDAPAllOfUserTypes.md) | An array of the userTypes properties for the users to be provisioned in PingOne. userTypes specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory. | [optional] 
**ValidateTlsCertificates** | Pointer to **bool** | A boolean that specifies whether or not to trust all SSL certificates (defaults to true). If this value is false, TLS certificates are not validated. When the value is set to true, only certificates that are signed by the default JVM CAs, or the CA certs that the customer has uploaded to the certificate service are trusted. | [optional] 
**Vendor** | [**EnumGatewayVendor**](EnumGatewayVendor.md) |  | 
**FollowReferrals** | Pointer to **bool** |  | [optional] [readonly] 
**Davinci** | [**GatewayTypeRADIUSAllOfDavinci**](GatewayTypeRADIUSAllOfDavinci.md) |  | 
**DefaultSharedSecret** | Pointer to **string** | Value to use for the shared secret if the shared secret is not provided for one or more of the RADIUS clients specified. | [optional] 
**NetworkPolicyServer** | Pointer to [**GatewayTypeRADIUSAllOfNetworkPolicyServer**](GatewayTypeRADIUSAllOfNetworkPolicyServer.md) |  | [optional] 
**RadiusClients** | [**[]GatewayTypeRADIUSAllOfRadiusClients**](GatewayTypeRADIUSAllOfRadiusClients.md) | Collection of RADIUS clients. | 

## Methods

### NewCreateGatewayRequest

`func NewCreateGatewayRequest(name string, type_ EnumGatewayType, enabled bool, bindDN string, bindPassword string, serversHostAndPort []string, vendor EnumGatewayVendor, davinci GatewayTypeRADIUSAllOfDavinci, radiusClients []GatewayTypeRADIUSAllOfRadiusClients, ) *CreateGatewayRequest`

NewCreateGatewayRequest instantiates a new CreateGatewayRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateGatewayRequestWithDefaults

`func NewCreateGatewayRequestWithDefaults() *CreateGatewayRequest`

NewCreateGatewayRequestWithDefaults instantiates a new CreateGatewayRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *CreateGatewayRequest) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CreateGatewayRequest) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CreateGatewayRequest) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CreateGatewayRequest) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *CreateGatewayRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateGatewayRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateGatewayRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateGatewayRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *CreateGatewayRequest) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CreateGatewayRequest) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CreateGatewayRequest) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CreateGatewayRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCredentials

`func (o *CreateGatewayRequest) GetCredentials() []GatewayCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *CreateGatewayRequest) GetCredentialsOk() (*[]GatewayCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *CreateGatewayRequest) SetCredentials(v []GatewayCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *CreateGatewayRequest) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *CreateGatewayRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateGatewayRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateGatewayRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateGatewayRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateGatewayRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateGatewayRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateGatewayRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateGatewayRequest) GetType() EnumGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateGatewayRequest) GetTypeOk() (*EnumGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateGatewayRequest) SetType(v EnumGatewayType)`

SetType sets Type field to given value.


### GetEnabled

`func (o *CreateGatewayRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CreateGatewayRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CreateGatewayRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSupportedVersions

`func (o *CreateGatewayRequest) GetSupportedVersions() []GatewaySupportedVersionsInner`

GetSupportedVersions returns the SupportedVersions field if non-nil, zero value otherwise.

### GetSupportedVersionsOk

`func (o *CreateGatewayRequest) GetSupportedVersionsOk() (*[]GatewaySupportedVersionsInner, bool)`

GetSupportedVersionsOk returns a tuple with the SupportedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVersions

`func (o *CreateGatewayRequest) SetSupportedVersions(v []GatewaySupportedVersionsInner)`

SetSupportedVersions sets SupportedVersions field to given value.

### HasSupportedVersions

`func (o *CreateGatewayRequest) HasSupportedVersions() bool`

HasSupportedVersions returns a boolean if a field has been set.

### GetCurrentAlerts

`func (o *CreateGatewayRequest) GetCurrentAlerts() []map[string]interface{}`

GetCurrentAlerts returns the CurrentAlerts field if non-nil, zero value otherwise.

### GetCurrentAlertsOk

`func (o *CreateGatewayRequest) GetCurrentAlertsOk() (*[]map[string]interface{}, bool)`

GetCurrentAlertsOk returns a tuple with the CurrentAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAlerts

`func (o *CreateGatewayRequest) SetCurrentAlerts(v []map[string]interface{})`

SetCurrentAlerts sets CurrentAlerts field to given value.

### HasCurrentAlerts

`func (o *CreateGatewayRequest) HasCurrentAlerts() bool`

HasCurrentAlerts returns a boolean if a field has been set.

### GetBindDN

`func (o *CreateGatewayRequest) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *CreateGatewayRequest) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *CreateGatewayRequest) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.


### GetBindPassword

`func (o *CreateGatewayRequest) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *CreateGatewayRequest) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *CreateGatewayRequest) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.


### GetConnectionSecurity

`func (o *CreateGatewayRequest) GetConnectionSecurity() EnumGatewayTypeLDAPSecurity`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *CreateGatewayRequest) GetConnectionSecurityOk() (*EnumGatewayTypeLDAPSecurity, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *CreateGatewayRequest) SetConnectionSecurity(v EnumGatewayTypeLDAPSecurity)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *CreateGatewayRequest) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetKerberos

`func (o *CreateGatewayRequest) GetKerberos() GatewayTypeLDAPAllOfKerberos`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *CreateGatewayRequest) GetKerberosOk() (*GatewayTypeLDAPAllOfKerberos, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *CreateGatewayRequest) SetKerberos(v GatewayTypeLDAPAllOfKerberos)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *CreateGatewayRequest) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetServersHostAndPort

`func (o *CreateGatewayRequest) GetServersHostAndPort() []string`

GetServersHostAndPort returns the ServersHostAndPort field if non-nil, zero value otherwise.

### GetServersHostAndPortOk

`func (o *CreateGatewayRequest) GetServersHostAndPortOk() (*[]string, bool)`

GetServersHostAndPortOk returns a tuple with the ServersHostAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersHostAndPort

`func (o *CreateGatewayRequest) SetServersHostAndPort(v []string)`

SetServersHostAndPort sets ServersHostAndPort field to given value.


### GetUserTypes

`func (o *CreateGatewayRequest) GetUserTypes() []GatewayTypeLDAPAllOfUserTypes`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *CreateGatewayRequest) GetUserTypesOk() (*[]GatewayTypeLDAPAllOfUserTypes, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *CreateGatewayRequest) SetUserTypes(v []GatewayTypeLDAPAllOfUserTypes)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *CreateGatewayRequest) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetValidateTlsCertificates

`func (o *CreateGatewayRequest) GetValidateTlsCertificates() bool`

GetValidateTlsCertificates returns the ValidateTlsCertificates field if non-nil, zero value otherwise.

### GetValidateTlsCertificatesOk

`func (o *CreateGatewayRequest) GetValidateTlsCertificatesOk() (*bool, bool)`

GetValidateTlsCertificatesOk returns a tuple with the ValidateTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateTlsCertificates

`func (o *CreateGatewayRequest) SetValidateTlsCertificates(v bool)`

SetValidateTlsCertificates sets ValidateTlsCertificates field to given value.

### HasValidateTlsCertificates

`func (o *CreateGatewayRequest) HasValidateTlsCertificates() bool`

HasValidateTlsCertificates returns a boolean if a field has been set.

### GetVendor

`func (o *CreateGatewayRequest) GetVendor() EnumGatewayVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *CreateGatewayRequest) GetVendorOk() (*EnumGatewayVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *CreateGatewayRequest) SetVendor(v EnumGatewayVendor)`

SetVendor sets Vendor field to given value.


### GetFollowReferrals

`func (o *CreateGatewayRequest) GetFollowReferrals() bool`

GetFollowReferrals returns the FollowReferrals field if non-nil, zero value otherwise.

### GetFollowReferralsOk

`func (o *CreateGatewayRequest) GetFollowReferralsOk() (*bool, bool)`

GetFollowReferralsOk returns a tuple with the FollowReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowReferrals

`func (o *CreateGatewayRequest) SetFollowReferrals(v bool)`

SetFollowReferrals sets FollowReferrals field to given value.

### HasFollowReferrals

`func (o *CreateGatewayRequest) HasFollowReferrals() bool`

HasFollowReferrals returns a boolean if a field has been set.

### GetDavinci

`func (o *CreateGatewayRequest) GetDavinci() GatewayTypeRADIUSAllOfDavinci`

GetDavinci returns the Davinci field if non-nil, zero value otherwise.

### GetDavinciOk

`func (o *CreateGatewayRequest) GetDavinciOk() (*GatewayTypeRADIUSAllOfDavinci, bool)`

GetDavinciOk returns a tuple with the Davinci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDavinci

`func (o *CreateGatewayRequest) SetDavinci(v GatewayTypeRADIUSAllOfDavinci)`

SetDavinci sets Davinci field to given value.


### GetDefaultSharedSecret

`func (o *CreateGatewayRequest) GetDefaultSharedSecret() string`

GetDefaultSharedSecret returns the DefaultSharedSecret field if non-nil, zero value otherwise.

### GetDefaultSharedSecretOk

`func (o *CreateGatewayRequest) GetDefaultSharedSecretOk() (*string, bool)`

GetDefaultSharedSecretOk returns a tuple with the DefaultSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSharedSecret

`func (o *CreateGatewayRequest) SetDefaultSharedSecret(v string)`

SetDefaultSharedSecret sets DefaultSharedSecret field to given value.

### HasDefaultSharedSecret

`func (o *CreateGatewayRequest) HasDefaultSharedSecret() bool`

HasDefaultSharedSecret returns a boolean if a field has been set.

### GetNetworkPolicyServer

`func (o *CreateGatewayRequest) GetNetworkPolicyServer() GatewayTypeRADIUSAllOfNetworkPolicyServer`

GetNetworkPolicyServer returns the NetworkPolicyServer field if non-nil, zero value otherwise.

### GetNetworkPolicyServerOk

`func (o *CreateGatewayRequest) GetNetworkPolicyServerOk() (*GatewayTypeRADIUSAllOfNetworkPolicyServer, bool)`

GetNetworkPolicyServerOk returns a tuple with the NetworkPolicyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicyServer

`func (o *CreateGatewayRequest) SetNetworkPolicyServer(v GatewayTypeRADIUSAllOfNetworkPolicyServer)`

SetNetworkPolicyServer sets NetworkPolicyServer field to given value.

### HasNetworkPolicyServer

`func (o *CreateGatewayRequest) HasNetworkPolicyServer() bool`

HasNetworkPolicyServer returns a boolean if a field has been set.

### GetRadiusClients

`func (o *CreateGatewayRequest) GetRadiusClients() []GatewayTypeRADIUSAllOfRadiusClients`

GetRadiusClients returns the RadiusClients field if non-nil, zero value otherwise.

### GetRadiusClientsOk

`func (o *CreateGatewayRequest) GetRadiusClientsOk() (*[]GatewayTypeRADIUSAllOfRadiusClients, bool)`

GetRadiusClientsOk returns a tuple with the RadiusClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusClients

`func (o *CreateGatewayRequest) SetRadiusClients(v []GatewayTypeRADIUSAllOfRadiusClients)`

SetRadiusClients sets RadiusClients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


