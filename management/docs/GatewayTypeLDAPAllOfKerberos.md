# GatewayTypeLDAPAllOfKerberos

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServiceAccountPassword** | Pointer to **string** | The password for the Kerberos service account. | [optional] 
**ServiceAccountUserPrincipalName** | **string** | The Kerberos service account user principal name (for example, &#x60;username@domain.com&#x60;). | 
**MinutesToRetainPreviousCredentials** | Pointer to **int32** | The number of minutes for which the previous credentials are persisted. | [optional] 

## Methods

### NewGatewayTypeLDAPAllOfKerberos

`func NewGatewayTypeLDAPAllOfKerberos(serviceAccountUserPrincipalName string, ) *GatewayTypeLDAPAllOfKerberos`

NewGatewayTypeLDAPAllOfKerberos instantiates a new GatewayTypeLDAPAllOfKerberos object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayTypeLDAPAllOfKerberosWithDefaults

`func NewGatewayTypeLDAPAllOfKerberosWithDefaults() *GatewayTypeLDAPAllOfKerberos`

NewGatewayTypeLDAPAllOfKerberosWithDefaults instantiates a new GatewayTypeLDAPAllOfKerberos object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServiceAccountPassword

`func (o *GatewayTypeLDAPAllOfKerberos) GetServiceAccountPassword() string`

GetServiceAccountPassword returns the ServiceAccountPassword field if non-nil, zero value otherwise.

### GetServiceAccountPasswordOk

`func (o *GatewayTypeLDAPAllOfKerberos) GetServiceAccountPasswordOk() (*string, bool)`

GetServiceAccountPasswordOk returns a tuple with the ServiceAccountPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountPassword

`func (o *GatewayTypeLDAPAllOfKerberos) SetServiceAccountPassword(v string)`

SetServiceAccountPassword sets ServiceAccountPassword field to given value.

### HasServiceAccountPassword

`func (o *GatewayTypeLDAPAllOfKerberos) HasServiceAccountPassword() bool`

HasServiceAccountPassword returns a boolean if a field has been set.

### GetServiceAccountUserPrincipalName

`func (o *GatewayTypeLDAPAllOfKerberos) GetServiceAccountUserPrincipalName() string`

GetServiceAccountUserPrincipalName returns the ServiceAccountUserPrincipalName field if non-nil, zero value otherwise.

### GetServiceAccountUserPrincipalNameOk

`func (o *GatewayTypeLDAPAllOfKerberos) GetServiceAccountUserPrincipalNameOk() (*string, bool)`

GetServiceAccountUserPrincipalNameOk returns a tuple with the ServiceAccountUserPrincipalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountUserPrincipalName

`func (o *GatewayTypeLDAPAllOfKerberos) SetServiceAccountUserPrincipalName(v string)`

SetServiceAccountUserPrincipalName sets ServiceAccountUserPrincipalName field to given value.


### GetMinutesToRetainPreviousCredentials

`func (o *GatewayTypeLDAPAllOfKerberos) GetMinutesToRetainPreviousCredentials() int32`

GetMinutesToRetainPreviousCredentials returns the MinutesToRetainPreviousCredentials field if non-nil, zero value otherwise.

### GetMinutesToRetainPreviousCredentialsOk

`func (o *GatewayTypeLDAPAllOfKerberos) GetMinutesToRetainPreviousCredentialsOk() (*int32, bool)`

GetMinutesToRetainPreviousCredentialsOk returns a tuple with the MinutesToRetainPreviousCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutesToRetainPreviousCredentials

`func (o *GatewayTypeLDAPAllOfKerberos) SetMinutesToRetainPreviousCredentials(v int32)`

SetMinutesToRetainPreviousCredentials sets MinutesToRetainPreviousCredentials field to given value.

### HasMinutesToRetainPreviousCredentials

`func (o *GatewayTypeLDAPAllOfKerberos) HasMinutesToRetainPreviousCredentials() bool`

HasMinutesToRetainPreviousCredentials returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


