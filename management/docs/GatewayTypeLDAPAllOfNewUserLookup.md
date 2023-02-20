# GatewayTypeLDAPAllOfNewUserLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeMappings** | [**[]GatewayTypeLDAPAllOfNewUserLookupAttributeMappings**](GatewayTypeLDAPAllOfNewUserLookupAttributeMappings.md) | A list of objects supplying a mapping of PingOne attributes to external LDAP attributes. One of the entries must be a mapping for \&quot;username&#x60;. This is required for the PingOne user schema. | 
**LdapFilterPattern** | **string** | The LDAP user search filter to use to match users against the entered user identifier at login. For example, (((uid&#x3D;${identifier})(mail&#x3D;${identifier})). Alternatively, this can be a search against the user directory. | 
**Population** | [**GatewayTypeLDAPAllOfNewUserLookupPopulation**](GatewayTypeLDAPAllOfNewUserLookupPopulation.md) |  | 

## Methods

### NewGatewayTypeLDAPAllOfNewUserLookup

`func NewGatewayTypeLDAPAllOfNewUserLookup(attributeMappings []GatewayTypeLDAPAllOfNewUserLookupAttributeMappings, ldapFilterPattern string, population GatewayTypeLDAPAllOfNewUserLookupPopulation, ) *GatewayTypeLDAPAllOfNewUserLookup`

NewGatewayTypeLDAPAllOfNewUserLookup instantiates a new GatewayTypeLDAPAllOfNewUserLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTypeLDAPAllOfNewUserLookupWithDefaults

`func NewGatewayTypeLDAPAllOfNewUserLookupWithDefaults() *GatewayTypeLDAPAllOfNewUserLookup`

NewGatewayTypeLDAPAllOfNewUserLookupWithDefaults instantiates a new GatewayTypeLDAPAllOfNewUserLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeMappings

`func (o *GatewayTypeLDAPAllOfNewUserLookup) GetAttributeMappings() []GatewayTypeLDAPAllOfNewUserLookupAttributeMappings`

GetAttributeMappings returns the AttributeMappings field if non-nil, zero value otherwise.

### GetAttributeMappingsOk

`func (o *GatewayTypeLDAPAllOfNewUserLookup) GetAttributeMappingsOk() (*[]GatewayTypeLDAPAllOfNewUserLookupAttributeMappings, bool)`

GetAttributeMappingsOk returns a tuple with the AttributeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMappings

`func (o *GatewayTypeLDAPAllOfNewUserLookup) SetAttributeMappings(v []GatewayTypeLDAPAllOfNewUserLookupAttributeMappings)`

SetAttributeMappings sets AttributeMappings field to given value.


### GetLdapFilterPattern

`func (o *GatewayTypeLDAPAllOfNewUserLookup) GetLdapFilterPattern() string`

GetLdapFilterPattern returns the LdapFilterPattern field if non-nil, zero value otherwise.

### GetLdapFilterPatternOk

`func (o *GatewayTypeLDAPAllOfNewUserLookup) GetLdapFilterPatternOk() (*string, bool)`

GetLdapFilterPatternOk returns a tuple with the LdapFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapFilterPattern

`func (o *GatewayTypeLDAPAllOfNewUserLookup) SetLdapFilterPattern(v string)`

SetLdapFilterPattern sets LdapFilterPattern field to given value.


### GetPopulation

`func (o *GatewayTypeLDAPAllOfNewUserLookup) GetPopulation() GatewayTypeLDAPAllOfNewUserLookupPopulation`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *GatewayTypeLDAPAllOfNewUserLookup) GetPopulationOk() (*GatewayTypeLDAPAllOfNewUserLookupPopulation, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *GatewayTypeLDAPAllOfNewUserLookup) SetPopulation(v GatewayTypeLDAPAllOfNewUserLookupPopulation)`

SetPopulation sets Population field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


