# GatewayTypeRADIUS

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
**Davinci** | [**GatewayTypeRADIUSAllOfDavinci**](GatewayTypeRADIUSAllOfDavinci.md) |  | 
**DefaultSharedSecret** | Pointer to **string** | Value to use for the shared secret if the shared secret is not provided for one or more of the RADIUS clients specified. | [optional] 
**NetworkPolicyServer** | Pointer to [**GatewayTypeRADIUSAllOfNetworkPolicyServer**](GatewayTypeRADIUSAllOfNetworkPolicyServer.md) |  | [optional] 
**RadiusClients** | [**[]GatewayTypeRADIUSAllOfRadiusClients**](GatewayTypeRADIUSAllOfRadiusClients.md) | Collection of RADIUS clients. | 

## Methods

### NewGatewayTypeRADIUS

`func NewGatewayTypeRADIUS(name string, type_ EnumGatewayType, enabled bool, davinci GatewayTypeRADIUSAllOfDavinci, radiusClients []GatewayTypeRADIUSAllOfRadiusClients, ) *GatewayTypeRADIUS`

NewGatewayTypeRADIUS instantiates a new GatewayTypeRADIUS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTypeRADIUSWithDefaults

`func NewGatewayTypeRADIUSWithDefaults() *GatewayTypeRADIUS`

NewGatewayTypeRADIUSWithDefaults instantiates a new GatewayTypeRADIUS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *GatewayTypeRADIUS) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GatewayTypeRADIUS) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GatewayTypeRADIUS) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GatewayTypeRADIUS) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetId

`func (o *GatewayTypeRADIUS) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayTypeRADIUS) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayTypeRADIUS) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayTypeRADIUS) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *GatewayTypeRADIUS) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *GatewayTypeRADIUS) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *GatewayTypeRADIUS) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *GatewayTypeRADIUS) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetCredentials

`func (o *GatewayTypeRADIUS) GetCredentials() []GatewayCredential`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *GatewayTypeRADIUS) GetCredentialsOk() (*[]GatewayCredential, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *GatewayTypeRADIUS) SetCredentials(v []GatewayCredential)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *GatewayTypeRADIUS) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetName

`func (o *GatewayTypeRADIUS) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayTypeRADIUS) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayTypeRADIUS) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GatewayTypeRADIUS) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayTypeRADIUS) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayTypeRADIUS) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayTypeRADIUS) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GatewayTypeRADIUS) GetType() EnumGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayTypeRADIUS) GetTypeOk() (*EnumGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayTypeRADIUS) SetType(v EnumGatewayType)`

SetType sets Type field to given value.


### GetEnabled

`func (o *GatewayTypeRADIUS) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GatewayTypeRADIUS) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GatewayTypeRADIUS) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetSupportedVersions

`func (o *GatewayTypeRADIUS) GetSupportedVersions() []GatewaySupportedVersionsInner`

GetSupportedVersions returns the SupportedVersions field if non-nil, zero value otherwise.

### GetSupportedVersionsOk

`func (o *GatewayTypeRADIUS) GetSupportedVersionsOk() (*[]GatewaySupportedVersionsInner, bool)`

GetSupportedVersionsOk returns a tuple with the SupportedVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedVersions

`func (o *GatewayTypeRADIUS) SetSupportedVersions(v []GatewaySupportedVersionsInner)`

SetSupportedVersions sets SupportedVersions field to given value.

### HasSupportedVersions

`func (o *GatewayTypeRADIUS) HasSupportedVersions() bool`

HasSupportedVersions returns a boolean if a field has been set.

### GetCurrentAlerts

`func (o *GatewayTypeRADIUS) GetCurrentAlerts() []map[string]interface{}`

GetCurrentAlerts returns the CurrentAlerts field if non-nil, zero value otherwise.

### GetCurrentAlertsOk

`func (o *GatewayTypeRADIUS) GetCurrentAlertsOk() (*[]map[string]interface{}, bool)`

GetCurrentAlertsOk returns a tuple with the CurrentAlerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAlerts

`func (o *GatewayTypeRADIUS) SetCurrentAlerts(v []map[string]interface{})`

SetCurrentAlerts sets CurrentAlerts field to given value.

### HasCurrentAlerts

`func (o *GatewayTypeRADIUS) HasCurrentAlerts() bool`

HasCurrentAlerts returns a boolean if a field has been set.

### GetDavinci

`func (o *GatewayTypeRADIUS) GetDavinci() GatewayTypeRADIUSAllOfDavinci`

GetDavinci returns the Davinci field if non-nil, zero value otherwise.

### GetDavinciOk

`func (o *GatewayTypeRADIUS) GetDavinciOk() (*GatewayTypeRADIUSAllOfDavinci, bool)`

GetDavinciOk returns a tuple with the Davinci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDavinci

`func (o *GatewayTypeRADIUS) SetDavinci(v GatewayTypeRADIUSAllOfDavinci)`

SetDavinci sets Davinci field to given value.


### GetDefaultSharedSecret

`func (o *GatewayTypeRADIUS) GetDefaultSharedSecret() string`

GetDefaultSharedSecret returns the DefaultSharedSecret field if non-nil, zero value otherwise.

### GetDefaultSharedSecretOk

`func (o *GatewayTypeRADIUS) GetDefaultSharedSecretOk() (*string, bool)`

GetDefaultSharedSecretOk returns a tuple with the DefaultSharedSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSharedSecret

`func (o *GatewayTypeRADIUS) SetDefaultSharedSecret(v string)`

SetDefaultSharedSecret sets DefaultSharedSecret field to given value.

### HasDefaultSharedSecret

`func (o *GatewayTypeRADIUS) HasDefaultSharedSecret() bool`

HasDefaultSharedSecret returns a boolean if a field has been set.

### GetNetworkPolicyServer

`func (o *GatewayTypeRADIUS) GetNetworkPolicyServer() GatewayTypeRADIUSAllOfNetworkPolicyServer`

GetNetworkPolicyServer returns the NetworkPolicyServer field if non-nil, zero value otherwise.

### GetNetworkPolicyServerOk

`func (o *GatewayTypeRADIUS) GetNetworkPolicyServerOk() (*GatewayTypeRADIUSAllOfNetworkPolicyServer, bool)`

GetNetworkPolicyServerOk returns a tuple with the NetworkPolicyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPolicyServer

`func (o *GatewayTypeRADIUS) SetNetworkPolicyServer(v GatewayTypeRADIUSAllOfNetworkPolicyServer)`

SetNetworkPolicyServer sets NetworkPolicyServer field to given value.

### HasNetworkPolicyServer

`func (o *GatewayTypeRADIUS) HasNetworkPolicyServer() bool`

HasNetworkPolicyServer returns a boolean if a field has been set.

### GetRadiusClients

`func (o *GatewayTypeRADIUS) GetRadiusClients() []GatewayTypeRADIUSAllOfRadiusClients`

GetRadiusClients returns the RadiusClients field if non-nil, zero value otherwise.

### GetRadiusClientsOk

`func (o *GatewayTypeRADIUS) GetRadiusClientsOk() (*[]GatewayTypeRADIUSAllOfRadiusClients, bool)`

GetRadiusClientsOk returns a tuple with the RadiusClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRadiusClients

`func (o *GatewayTypeRADIUS) SetRadiusClients(v []GatewayTypeRADIUSAllOfRadiusClients)`

SetRadiusClients sets RadiusClients field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


