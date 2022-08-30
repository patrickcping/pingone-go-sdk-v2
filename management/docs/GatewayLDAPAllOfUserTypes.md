# GatewayLDAPAllOfUserTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowPasswordChanges** | Pointer to **bool** | Defaults to false if this property isn&#39;t specified in the request. If false, the user cannot change the password in the remote LDAP directory. In this case, operations for forgotten passwords or resetting of passwords are not available to a user referencing this gateway. | [optional] 
**Id** | **string** | The UUID of the user type. This correlates to the password.external.gateway.userType.id User property. | 
**Name** | **string** | The name of the user type. | 
**NewUserLookup** | [**GatewayLDAPAllOfNewUserLookup**](GatewayLDAPAllOfNewUserLookup.md) |  | 

## Methods

### NewGatewayLDAPAllOfUserTypes

`func NewGatewayLDAPAllOfUserTypes(id string, name string, newUserLookup GatewayLDAPAllOfNewUserLookup, ) *GatewayLDAPAllOfUserTypes`

NewGatewayLDAPAllOfUserTypes instantiates a new GatewayLDAPAllOfUserTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayLDAPAllOfUserTypesWithDefaults

`func NewGatewayLDAPAllOfUserTypesWithDefaults() *GatewayLDAPAllOfUserTypes`

NewGatewayLDAPAllOfUserTypesWithDefaults instantiates a new GatewayLDAPAllOfUserTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowPasswordChanges

`func (o *GatewayLDAPAllOfUserTypes) GetAllowPasswordChanges() bool`

GetAllowPasswordChanges returns the AllowPasswordChanges field if non-nil, zero value otherwise.

### GetAllowPasswordChangesOk

`func (o *GatewayLDAPAllOfUserTypes) GetAllowPasswordChangesOk() (*bool, bool)`

GetAllowPasswordChangesOk returns a tuple with the AllowPasswordChanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPasswordChanges

`func (o *GatewayLDAPAllOfUserTypes) SetAllowPasswordChanges(v bool)`

SetAllowPasswordChanges sets AllowPasswordChanges field to given value.

### HasAllowPasswordChanges

`func (o *GatewayLDAPAllOfUserTypes) HasAllowPasswordChanges() bool`

HasAllowPasswordChanges returns a boolean if a field has been set.

### GetId

`func (o *GatewayLDAPAllOfUserTypes) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayLDAPAllOfUserTypes) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayLDAPAllOfUserTypes) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GatewayLDAPAllOfUserTypes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayLDAPAllOfUserTypes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayLDAPAllOfUserTypes) SetName(v string)`

SetName sets Name field to given value.


### GetNewUserLookup

`func (o *GatewayLDAPAllOfUserTypes) GetNewUserLookup() GatewayLDAPAllOfNewUserLookup`

GetNewUserLookup returns the NewUserLookup field if non-nil, zero value otherwise.

### GetNewUserLookupOk

`func (o *GatewayLDAPAllOfUserTypes) GetNewUserLookupOk() (*GatewayLDAPAllOfNewUserLookup, bool)`

GetNewUserLookupOk returns a tuple with the NewUserLookup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewUserLookup

`func (o *GatewayLDAPAllOfUserTypes) SetNewUserLookup(v GatewayLDAPAllOfNewUserLookup)`

SetNewUserLookup sets NewUserLookup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


