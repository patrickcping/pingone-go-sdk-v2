# UserPasswordExternalGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The UUID of the linked gateway that references the remote directory. | [optional] 
**Type** | Pointer to **string** | An enum indicating one of the supported gateway types. For the supported types, see type in the Gateway base data model. | [optional] 
**UserType** | Pointer to [**UserPasswordExternalGatewayUserType**](UserPasswordExternalGatewayUserType.md) |  | [optional] 
**CorrelationAttributes** | Pointer to **map[string]interface{}** | An object that maps the external LDAP directory attributes to PingOne attributes. We use the correlationAttributes values to read the attributes from the external LDAP directory and map them to the corresponding PingOne attributes. | [optional] 

## Methods

### NewUserPasswordExternalGateway

`func NewUserPasswordExternalGateway() *UserPasswordExternalGateway`

NewUserPasswordExternalGateway instantiates a new UserPasswordExternalGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPasswordExternalGatewayWithDefaults

`func NewUserPasswordExternalGatewayWithDefaults() *UserPasswordExternalGateway`

NewUserPasswordExternalGatewayWithDefaults instantiates a new UserPasswordExternalGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserPasswordExternalGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPasswordExternalGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPasswordExternalGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UserPasswordExternalGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UserPasswordExternalGateway) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UserPasswordExternalGateway) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UserPasswordExternalGateway) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UserPasswordExternalGateway) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserType

`func (o *UserPasswordExternalGateway) GetUserType() UserPasswordExternalGatewayUserType`

GetUserType returns the UserType field if non-nil, zero value otherwise.

### GetUserTypeOk

`func (o *UserPasswordExternalGateway) GetUserTypeOk() (*UserPasswordExternalGatewayUserType, bool)`

GetUserTypeOk returns a tuple with the UserType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserType

`func (o *UserPasswordExternalGateway) SetUserType(v UserPasswordExternalGatewayUserType)`

SetUserType sets UserType field to given value.

### HasUserType

`func (o *UserPasswordExternalGateway) HasUserType() bool`

HasUserType returns a boolean if a field has been set.

### GetCorrelationAttributes

`func (o *UserPasswordExternalGateway) GetCorrelationAttributes() map[string]interface{}`

GetCorrelationAttributes returns the CorrelationAttributes field if non-nil, zero value otherwise.

### GetCorrelationAttributesOk

`func (o *UserPasswordExternalGateway) GetCorrelationAttributesOk() (*map[string]interface{}, bool)`

GetCorrelationAttributesOk returns a tuple with the CorrelationAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationAttributes

`func (o *UserPasswordExternalGateway) SetCorrelationAttributes(v map[string]interface{})`

SetCorrelationAttributes sets CorrelationAttributes field to given value.

### HasCorrelationAttributes

`func (o *UserPasswordExternalGateway) HasCorrelationAttributes() bool`

HasCorrelationAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


