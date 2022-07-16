# ApplicationAccessControlRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | A string that specifies the user role required to access the application. Options are ADMIN_USERS_ONLY. A user is an admin user if the user has one or more of the following roles Organization Admin, Environment Admin, Identity Data Admin, or Client Application Developer. | 

## Methods

### NewApplicationAccessControlRole

`func NewApplicationAccessControlRole(type_ string, ) *ApplicationAccessControlRole`

NewApplicationAccessControlRole instantiates a new ApplicationAccessControlRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationAccessControlRoleWithDefaults

`func NewApplicationAccessControlRoleWithDefaults() *ApplicationAccessControlRole`

NewApplicationAccessControlRoleWithDefaults instantiates a new ApplicationAccessControlRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApplicationAccessControlRole) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationAccessControlRole) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationAccessControlRole) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


