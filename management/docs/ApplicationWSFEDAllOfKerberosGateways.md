# ApplicationWSFEDAllOfKerberosGateways

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The UUID of the LDAP gateway. | 
**Type** | [**EnumApplicationWSFEDKerberosGatewayType**](EnumApplicationWSFEDKerberosGatewayType.md) |  | 
**UserType** | [**ApplicationWSFEDAllOfKerberosUserType**](ApplicationWSFEDAllOfKerberosUserType.md) |  | 

## Methods

### NewApplicationWSFEDAllOfKerberosGateways

`func NewApplicationWSFEDAllOfKerberosGateways(id string, type_ EnumApplicationWSFEDKerberosGatewayType, userType ApplicationWSFEDAllOfKerberosUserType, ) *ApplicationWSFEDAllOfKerberosGateways`

NewApplicationWSFEDAllOfKerberosGateways instantiates a new ApplicationWSFEDAllOfKerberosGateways object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWSFEDAllOfKerberosGatewaysWithDefaults

`func NewApplicationWSFEDAllOfKerberosGatewaysWithDefaults() *ApplicationWSFEDAllOfKerberosGateways`

NewApplicationWSFEDAllOfKerberosGatewaysWithDefaults instantiates a new ApplicationWSFEDAllOfKerberosGateways object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApplicationWSFEDAllOfKerberosGateways) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationWSFEDAllOfKerberosGateways) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationWSFEDAllOfKerberosGateways) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ApplicationWSFEDAllOfKerberosGateways) GetType() EnumApplicationWSFEDKerberosGatewayType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApplicationWSFEDAllOfKerberosGateways) GetTypeOk() (*EnumApplicationWSFEDKerberosGatewayType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApplicationWSFEDAllOfKerberosGateways) SetType(v EnumApplicationWSFEDKerberosGatewayType)`

SetType sets Type field to given value.


### GetUserType

`func (o *ApplicationWSFEDAllOfKerberosGateways) GetUserType() ApplicationWSFEDAllOfKerberosUserType`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *ApplicationWSFEDAllOfKerberosGateways) GetUserTypeOk() (*ApplicationWSFEDAllOfKerberosUserType, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *ApplicationWSFEDAllOfKerberosGateways) SetUserType(v ApplicationWSFEDAllOfKerberosUserType)`

SetUserType sets UserType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


