# GatewayTypeLDAPAllOfUserTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPasswordChanges** | Pointer to **bool** | Defaults to false if this property isn&#39;t specified in the request. If false, the user cannot change the password in the remote LDAP directory. In this case, operations for forgotten passwords or resetting of passwords are not available to a user referencing this gateway. | [optional] 
**UpdateUserOnSuccessfulAuthentication** | Pointer to **bool** | If set to true, when users sign on through an LDAP Gateway client, user attributes are updated based on responses from the LDAP server. Defaults to false if this property isn&#39;t specified in the request. | [optional] 
**Id** | Pointer to **string** | The UUID of the user type. This correlates to the password.external.gateway.userType.id User property. | [optional] 
**Name** | **string** | The name of the user type. | 
**NewUserLookup** | Pointer to [**GatewayTypeLDAPAllOfNewUserLookup**](GatewayTypeLDAPAllOfNewUserLookup.md) |  | [optional] 
**OrderedCorrelationAttributes** | **[]string** | A map of key/value entries used to persist the external LDAP directory attributes. | 
**PasswordAuthority** | [**EnumGatewayPasswordAuthority**](EnumGatewayPasswordAuthority.md) |  | 
**SearchBaseDn** | **string** | The LDAP base domain name (DN) for this user type. | 

## Methods

### NewGatewayTypeLDAPAllOfUserTypes

`func NewGatewayTypeLDAPAllOfUserTypes(name string, orderedCorrelationAttributes []string, passwordAuthority EnumGatewayPasswordAuthority, searchBaseDn string, ) *GatewayTypeLDAPAllOfUserTypes`

NewGatewayTypeLDAPAllOfUserTypes instantiates a new GatewayTypeLDAPAllOfUserTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTypeLDAPAllOfUserTypesWithDefaults

`func NewGatewayTypeLDAPAllOfUserTypesWithDefaults() *GatewayTypeLDAPAllOfUserTypes`

NewGatewayTypeLDAPAllOfUserTypesWithDefaults instantiates a new GatewayTypeLDAPAllOfUserTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPasswordChanges

`func (o *GatewayTypeLDAPAllOfUserTypes) GetAllowPasswordChanges() bool`

GetAllowPasswordChanges returns the AllowPasswordChanges field if non-nil, zero value otherwise.

### GetAllowPasswordChangesOk

`func (o *GatewayTypeLDAPAllOfUserTypes) GetAllowPasswordChangesOk() (*bool, bool)`

GetAllowPasswordChangesOk returns a tuple with the AllowPasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPasswordChanges

`func (o *GatewayTypeLDAPAllOfUserTypes) SetAllowPasswordChanges(v bool)`

SetAllowPasswordChanges sets AllowPasswordChanges field to given value.

### HasAllowPasswordChanges

`func (o *GatewayTypeLDAPAllOfUserTypes) HasAllowPasswordChanges() bool`

HasAllowPasswordChanges returns a boolean if a field has been set.

### GetUpdateUserOnSuccessfulAuthentication

`func (o *GatewayTypeLDAPAllOfUserTypes) GetUpdateUserOnSuccessfulAuthentication() bool`

GetUpdateUserOnSuccessfulAuthentication returns the UpdateUserOnSuccessfulAuthentication field if non-nil, zero value otherwise.

### GetUpdateUserOnSuccessfulAuthenticationOk

`func (o *GatewayTypeLDAPAllOfUserTypes) GetUpdateUserOnSuccessfulAuthenticationOk() (*bool, bool)`

GetUpdateUserOnSuccessfulAuthenticationOk returns a tuple with the UpdateUserOnSuccessfulAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateUserOnSuccessfulAuthentication

`func (o *GatewayTypeLDAPAllOfUserTypes) SetUpdateUserOnSuccessfulAuthentication(v bool)`

SetUpdateUserOnSuccessfulAuthentication sets UpdateUserOnSuccessfulAuthentication field to given value.

### HasUpdateUserOnSuccessfulAuthentication

`func (o *GatewayTypeLDAPAllOfUserTypes) HasUpdateUserOnSuccessfulAuthentication() bool`

HasUpdateUserOnSuccessfulAuthentication returns a boolean if a field has been set.

### GetId

`func (o *GatewayTypeLDAPAllOfUserTypes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayTypeLDAPAllOfUserTypes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayTypeLDAPAllOfUserTypes) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayTypeLDAPAllOfUserTypes) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GatewayTypeLDAPAllOfUserTypes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayTypeLDAPAllOfUserTypes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayTypeLDAPAllOfUserTypes) SetName(v string)`

SetName sets Name field to given value.


### GetNewUserLookup

`func (o *GatewayTypeLDAPAllOfUserTypes) GetNewUserLookup() GatewayTypeLDAPAllOfNewUserLookup`

GetNewUserLookup returns the NewUserLookup field if non-nil, zero value otherwise.

### GetNewUserLookupOk

`func (o *GatewayTypeLDAPAllOfUserTypes) GetNewUserLookupOk() (*GatewayTypeLDAPAllOfNewUserLookup, bool)`

GetNewUserLookupOk returns a tuple with the NewUserLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewUserLookup

`func (o *GatewayTypeLDAPAllOfUserTypes) SetNewUserLookup(v GatewayTypeLDAPAllOfNewUserLookup)`

SetNewUserLookup sets NewUserLookup field to given value.

### HasNewUserLookup

`func (o *GatewayTypeLDAPAllOfUserTypes) HasNewUserLookup() bool`

HasNewUserLookup returns a boolean if a field has been set.

### GetOrderedCorrelationAttributes

`func (o *GatewayTypeLDAPAllOfUserTypes) GetOrderedCorrelationAttributes() []string`

GetOrderedCorrelationAttributes returns the OrderedCorrelationAttributes field if non-nil, zero value otherwise.

### GetOrderedCorrelationAttributesOk

`func (o *GatewayTypeLDAPAllOfUserTypes) GetOrderedCorrelationAttributesOk() (*[]string, bool)`

GetOrderedCorrelationAttributesOk returns a tuple with the OrderedCorrelationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderedCorrelationAttributes

`func (o *GatewayTypeLDAPAllOfUserTypes) SetOrderedCorrelationAttributes(v []string)`

SetOrderedCorrelationAttributes sets OrderedCorrelationAttributes field to given value.


### GetPasswordAuthority

`func (o *GatewayTypeLDAPAllOfUserTypes) GetPasswordAuthority() EnumGatewayPasswordAuthority`

GetPasswordAuthority returns the PasswordAuthority field if non-nil, zero value otherwise.

### GetPasswordAuthorityOk

`func (o *GatewayTypeLDAPAllOfUserTypes) GetPasswordAuthorityOk() (*EnumGatewayPasswordAuthority, bool)`

GetPasswordAuthorityOk returns a tuple with the PasswordAuthority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordAuthority

`func (o *GatewayTypeLDAPAllOfUserTypes) SetPasswordAuthority(v EnumGatewayPasswordAuthority)`

SetPasswordAuthority sets PasswordAuthority field to given value.


### GetSearchBaseDn

`func (o *GatewayTypeLDAPAllOfUserTypes) GetSearchBaseDn() string`

GetSearchBaseDn returns the SearchBaseDn field if non-nil, zero value otherwise.

### GetSearchBaseDnOk

`func (o *GatewayTypeLDAPAllOfUserTypes) GetSearchBaseDnOk() (*string, bool)`

GetSearchBaseDnOk returns a tuple with the SearchBaseDn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchBaseDn

`func (o *GatewayTypeLDAPAllOfUserTypes) SetSearchBaseDn(v string)`

SetSearchBaseDn sets SearchBaseDn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


