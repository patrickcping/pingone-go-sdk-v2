# GatewayTypeLDAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to **map[string]interface{}** |  | [optional] 
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
**ConnectionSecurity** | Pointer to [**EnumGatewayTypeLDAPSecurity**](EnumGatewayTypeLDAPSecurity.md) |  | [optional] 
**Kerberos** | Pointer to [**GatewayTypeLDAPAllOfKerberos**](GatewayTypeLDAPAllOfKerberos.md) |  | [optional] 
**ServersHostAndPort** | **[]string** | An array of strings that specifies the LDAP server host name and port number (for example, [&#x60;ds1.example.com:389&#x60;, &#x60;ds2.example.com:389&#x60;]). | 
**UserTypes** | Pointer to [**[]GatewayTypeLDAPAllOfUserTypes**](GatewayTypeLDAPAllOfUserTypes.md) | An array of the userTypes properties for the users to be provisioned in PingOne. userTypes specifies which user properties in PingOne correspond to the user properties in an external LDAP directory. You can use an LDAP browser to view the user properties in the external LDAP directory. | [optional] 
**ValidateTlsCertificates** | Pointer to **bool** | A boolean that specifies whether or not to trust all SSL certificates (defaults to true). If this value is false, TLS certificates are not validated. When the value is set to true, only certificates that are signed by the default JVM CAs, or the CA certs that the customer has uploaded to the certificate service are trusted. | [optional] 
**Vendor** | [**EnumGatewayVendor**](EnumGatewayVendor.md) |  | 
**FollowReferrals** | Pointer to **bool** |  | [optional] [readonly] 

## Methods

### NewGatewayTypeLDAP

`func NewGatewayTypeLDAP(name string, type_ EnumGatewayType, enabled bool, bindDN string, bindPassword string, serversHostAndPort []string, vendor EnumGatewayVendor, ) *GatewayTypeLDAP`

NewGatewayTypeLDAP instantiates a new GatewayTypeLDAP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTypeLDAPWithDefaults

`func NewGatewayTypeLDAPWithDefaults() *GatewayTypeLDAP`

NewGatewayTypeLDAPWithDefaults instantiates a new GatewayTypeLDAP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GatewayTypeLDAP) GetLinks() map[string]interface{}`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GatewayTypeLDAP) GetLinksOk() (*map[string]interface{}, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GatewayTypeLDAP) SetLinks(v map[string]interface{})`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GatewayTypeLDAP) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *GatewayTypeLDAP) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayTypeLDAP) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayTypeLDAP) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayTypeLDAP) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *GatewayTypeLDAP) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GatewayTypeLDAP) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GatewayTypeLDAP) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GatewayTypeLDAP) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCredentials

`func (o *GatewayTypeLDAP) GetCredentials() []GatewayCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GatewayTypeLDAP) GetCredentialsOk() (*[]GatewayCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GatewayTypeLDAP) SetCredentials(v []GatewayCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GatewayTypeLDAP) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *GatewayTypeLDAP) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayTypeLDAP) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayTypeLDAP) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GatewayTypeLDAP) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayTypeLDAP) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayTypeLDAP) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayTypeLDAP) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GatewayTypeLDAP) GetType() EnumGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayTypeLDAP) GetTypeOk() (*EnumGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayTypeLDAP) SetType(v EnumGatewayType)`

SetType sets Type field to given value.


### GetEnabled

`func (o *GatewayTypeLDAP) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayTypeLDAP) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayTypeLDAP) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSupportedVersions

`func (o *GatewayTypeLDAP) GetSupportedVersions() []GatewaySupportedVersionsInner`

GetSupportedVersions returns the SupportedVersions field if non-nil, zero value otherwise.

### GetSupportedVersionsOk

`func (o *GatewayTypeLDAP) GetSupportedVersionsOk() (*[]GatewaySupportedVersionsInner, bool)`

GetSupportedVersionsOk returns a tuple with the SupportedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVersions

`func (o *GatewayTypeLDAP) SetSupportedVersions(v []GatewaySupportedVersionsInner)`

SetSupportedVersions sets SupportedVersions field to given value.

### HasSupportedVersions

`func (o *GatewayTypeLDAP) HasSupportedVersions() bool`

HasSupportedVersions returns a boolean if a field has been set.

### GetCurrentAlerts

`func (o *GatewayTypeLDAP) GetCurrentAlerts() []map[string]interface{}`

GetCurrentAlerts returns the CurrentAlerts field if non-nil, zero value otherwise.

### GetCurrentAlertsOk

`func (o *GatewayTypeLDAP) GetCurrentAlertsOk() (*[]map[string]interface{}, bool)`

GetCurrentAlertsOk returns a tuple with the CurrentAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAlerts

`func (o *GatewayTypeLDAP) SetCurrentAlerts(v []map[string]interface{})`

SetCurrentAlerts sets CurrentAlerts field to given value.

### HasCurrentAlerts

`func (o *GatewayTypeLDAP) HasCurrentAlerts() bool`

HasCurrentAlerts returns a boolean if a field has been set.

### GetBindDN

`func (o *GatewayTypeLDAP) GetBindDN() string`

GetBindDN returns the BindDN field if non-nil, zero value otherwise.

### GetBindDNOk

`func (o *GatewayTypeLDAP) GetBindDNOk() (*string, bool)`

GetBindDNOk returns a tuple with the BindDN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindDN

`func (o *GatewayTypeLDAP) SetBindDN(v string)`

SetBindDN sets BindDN field to given value.


### GetBindPassword

`func (o *GatewayTypeLDAP) GetBindPassword() string`

GetBindPassword returns the BindPassword field if non-nil, zero value otherwise.

### GetBindPasswordOk

`func (o *GatewayTypeLDAP) GetBindPasswordOk() (*string, bool)`

GetBindPasswordOk returns a tuple with the BindPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindPassword

`func (o *GatewayTypeLDAP) SetBindPassword(v string)`

SetBindPassword sets BindPassword field to given value.


### GetConnectionSecurity

`func (o *GatewayTypeLDAP) GetConnectionSecurity() EnumGatewayTypeLDAPSecurity`

GetConnectionSecurity returns the ConnectionSecurity field if non-nil, zero value otherwise.

### GetConnectionSecurityOk

`func (o *GatewayTypeLDAP) GetConnectionSecurityOk() (*EnumGatewayTypeLDAPSecurity, bool)`

GetConnectionSecurityOk returns a tuple with the ConnectionSecurity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionSecurity

`func (o *GatewayTypeLDAP) SetConnectionSecurity(v EnumGatewayTypeLDAPSecurity)`

SetConnectionSecurity sets ConnectionSecurity field to given value.

### HasConnectionSecurity

`func (o *GatewayTypeLDAP) HasConnectionSecurity() bool`

HasConnectionSecurity returns a boolean if a field has been set.

### GetKerberos

`func (o *GatewayTypeLDAP) GetKerberos() GatewayTypeLDAPAllOfKerberos`

GetKerberos returns the Kerberos field if non-nil, zero value otherwise.

### GetKerberosOk

`func (o *GatewayTypeLDAP) GetKerberosOk() (*GatewayTypeLDAPAllOfKerberos, bool)`

GetKerberosOk returns a tuple with the Kerberos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKerberos

`func (o *GatewayTypeLDAP) SetKerberos(v GatewayTypeLDAPAllOfKerberos)`

SetKerberos sets Kerberos field to given value.

### HasKerberos

`func (o *GatewayTypeLDAP) HasKerberos() bool`

HasKerberos returns a boolean if a field has been set.

### GetServersHostAndPort

`func (o *GatewayTypeLDAP) GetServersHostAndPort() []string`

GetServersHostAndPort returns the ServersHostAndPort field if non-nil, zero value otherwise.

### GetServersHostAndPortOk

`func (o *GatewayTypeLDAP) GetServersHostAndPortOk() (*[]string, bool)`

GetServersHostAndPortOk returns a tuple with the ServersHostAndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServersHostAndPort

`func (o *GatewayTypeLDAP) SetServersHostAndPort(v []string)`

SetServersHostAndPort sets ServersHostAndPort field to given value.


### GetUserTypes

`func (o *GatewayTypeLDAP) GetUserTypes() []GatewayTypeLDAPAllOfUserTypes`

GetUserTypes returns the UserTypes field if non-nil, zero value otherwise.

### GetUserTypesOk

`func (o *GatewayTypeLDAP) GetUserTypesOk() (*[]GatewayTypeLDAPAllOfUserTypes, bool)`

GetUserTypesOk returns a tuple with the UserTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTypes

`func (o *GatewayTypeLDAP) SetUserTypes(v []GatewayTypeLDAPAllOfUserTypes)`

SetUserTypes sets UserTypes field to given value.

### HasUserTypes

`func (o *GatewayTypeLDAP) HasUserTypes() bool`

HasUserTypes returns a boolean if a field has been set.

### GetValidateTlsCertificates

`func (o *GatewayTypeLDAP) GetValidateTlsCertificates() bool`

GetValidateTlsCertificates returns the ValidateTlsCertificates field if non-nil, zero value otherwise.

### GetValidateTlsCertificatesOk

`func (o *GatewayTypeLDAP) GetValidateTlsCertificatesOk() (*bool, bool)`

GetValidateTlsCertificatesOk returns a tuple with the ValidateTlsCertificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidateTlsCertificates

`func (o *GatewayTypeLDAP) SetValidateTlsCertificates(v bool)`

SetValidateTlsCertificates sets ValidateTlsCertificates field to given value.

### HasValidateTlsCertificates

`func (o *GatewayTypeLDAP) HasValidateTlsCertificates() bool`

HasValidateTlsCertificates returns a boolean if a field has been set.

### GetVendor

`func (o *GatewayTypeLDAP) GetVendor() EnumGatewayVendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *GatewayTypeLDAP) GetVendorOk() (*EnumGatewayVendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *GatewayTypeLDAP) SetVendor(v EnumGatewayVendor)`

SetVendor sets Vendor field to given value.


### GetFollowReferrals

`func (o *GatewayTypeLDAP) GetFollowReferrals() bool`

GetFollowReferrals returns the FollowReferrals field if non-nil, zero value otherwise.

### GetFollowReferralsOk

`func (o *GatewayTypeLDAP) GetFollowReferralsOk() (*bool, bool)`

GetFollowReferralsOk returns a tuple with the FollowReferrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowReferrals

`func (o *GatewayTypeLDAP) SetFollowReferrals(v bool)`

SetFollowReferrals sets FollowReferrals field to given value.

### HasFollowReferrals

`func (o *GatewayTypeLDAP) HasFollowReferrals() bool`

HasFollowReferrals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


