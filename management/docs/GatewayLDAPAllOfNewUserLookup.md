# GatewayLDAPAllOfNewUserLookup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AttributeMappings** | [**[]GatewayLDAPAllOfNewUserLookupAttributeMappings**](GatewayLDAPAllOfNewUserLookupAttributeMappings.md) | A list of objects supplying a mapping of PingOne attributes to external LDAP attributes. One of the entries must be a mapping for \&quot;username”. This is required for the PingOne user schema. | 
**LdapFilterPattern** | Pointer to **string** | (Optional) The LDAP user search filter to use to match users against the entered user identifier at login. For example, (((uid&#x3D;${identifier})(mail&#x3D;${identifier})). Alternatively, this can be a search against the user directory. | [optional] 
**Population** | Pointer to [**GatewayLDAPAllOfNewUserLookupPopulation**](GatewayLDAPAllOfNewUserLookupPopulation.md) |  | [optional] 

## Methods

### NewGatewayLDAPAllOfNewUserLookup

`func NewGatewayLDAPAllOfNewUserLookup(attributeMappings []GatewayLDAPAllOfNewUserLookupAttributeMappings, ) *GatewayLDAPAllOfNewUserLookup`

NewGatewayLDAPAllOfNewUserLookup instantiates a new GatewayLDAPAllOfNewUserLookup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayLDAPAllOfNewUserLookupWithDefaults

`func NewGatewayLDAPAllOfNewUserLookupWithDefaults() *GatewayLDAPAllOfNewUserLookup`

NewGatewayLDAPAllOfNewUserLookupWithDefaults instantiates a new GatewayLDAPAllOfNewUserLookup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAttributeMappings

`func (o *GatewayLDAPAllOfNewUserLookup) GetAttributeMappings() []GatewayLDAPAllOfNewUserLookupAttributeMappings`

GetAttributeMappings returns the AttributeMappings field if non-nil, zero value otherwise.

### GetAttributeMappingsOk

`func (o *GatewayLDAPAllOfNewUserLookup) GetAttributeMappingsOk() (*[]GatewayLDAPAllOfNewUserLookupAttributeMappings, bool)`

GetAttributeMappingsOk returns a tuple with the AttributeMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeMappings

`func (o *GatewayLDAPAllOfNewUserLookup) SetAttributeMappings(v []GatewayLDAPAllOfNewUserLookupAttributeMappings)`

SetAttributeMappings sets AttributeMappings field to given value.


### GetLdapFilterPattern

`func (o *GatewayLDAPAllOfNewUserLookup) GetLdapFilterPattern() string`

GetLdapFilterPattern returns the LdapFilterPattern field if non-nil, zero value otherwise.

### GetLdapFilterPatternOk

`func (o *GatewayLDAPAllOfNewUserLookup) GetLdapFilterPatternOk() (*string, bool)`

GetLdapFilterPatternOk returns a tuple with the LdapFilterPattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapFilterPattern

`func (o *GatewayLDAPAllOfNewUserLookup) SetLdapFilterPattern(v string)`

SetLdapFilterPattern sets LdapFilterPattern field to given value.

### HasLdapFilterPattern

`func (o *GatewayLDAPAllOfNewUserLookup) HasLdapFilterPattern() bool`

HasLdapFilterPattern returns a boolean if a field has been set.

### GetPopulation

`func (o *GatewayLDAPAllOfNewUserLookup) GetPopulation() GatewayLDAPAllOfNewUserLookupPopulation`

GetPopulation returns the Population field if non-nil, zero value otherwise.

### GetPopulationOk

`func (o *GatewayLDAPAllOfNewUserLookup) GetPopulationOk() (*GatewayLDAPAllOfNewUserLookupPopulation, bool)`

GetPopulationOk returns a tuple with the Population field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulation

`func (o *GatewayLDAPAllOfNewUserLookup) SetPopulation(v GatewayLDAPAllOfNewUserLookupPopulation)`

SetPopulation sets Population field to given value.

### HasPopulation

`func (o *GatewayLDAPAllOfNewUserLookup) HasPopulation() bool`

HasPopulation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


