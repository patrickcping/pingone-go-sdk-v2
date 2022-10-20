# ResourceScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies the resource scope name. | 
**Description** | Pointer to **string** | A string that specifies the description of the scope. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**SchemaAttributes** | Pointer to **[]string** | An array that specifies the user schema attributes that can be read or updated for the specified PingOne access control scope. The value is an array of schema attribute paths (such as &#x60;username&#x60;, &#x60;name.given&#x60;, &#x60;shirtSize&#x60;) that the scope controls. This property is supported only for the &#x60;p1:read:user&#x60;, &#x60;p1:update:user&#x60; and &#x60;p1:read:user:{suffix}&#x60; and &#x60;p1:update:user:{suffix}&#x60; scopes. No other PingOne platform scopes allow this behavior. Any attributes not listed in the attribute array are excluded from the read or update action. The wildcard path (*) in the array includes all attributes and cannot be used in conjunction with any other user schema attribute paths | [optional] 
**MappedClaims** | Pointer to **[]string** | A list of custom resource attribute IDs. This property applies only for the resource with its type property set to &#x60;OPENID_CONNECT&#x60;. Moreover, this property does not display predefined OpenID Connect (OIDC) mappings, such as the &#x60;email&#x60; claim in the OIDC &#x60;email&#x60; scope or the &#x60;name&#x60; claim in the &#x60;profile&#x60; scope. You can create custom attributes, and these custom attributes can be added to &#x60;mappedClaims&#x60; and will display in the response. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewResourceScope

`func NewResourceScope(name string, ) *ResourceScope`

NewResourceScope instantiates a new ResourceScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceScopeWithDefaults

`func NewResourceScopeWithDefaults() *ResourceScope`

NewResourceScopeWithDefaults instantiates a new ResourceScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceScope) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceScope) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceScope) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceScope) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceScope) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceScope) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceScope) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ResourceScope) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceScope) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceScope) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceScope) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *ResourceScope) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceScope) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceScope) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ResourceScope) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSchemaAttributes

`func (o *ResourceScope) GetSchemaAttributes() []string`

GetSchemaAttributes returns the SchemaAttributes field if non-nil, zero value otherwise.

### GetSchemaAttributesOk

`func (o *ResourceScope) GetSchemaAttributesOk() (*[]string, bool)`

GetSchemaAttributesOk returns a tuple with the SchemaAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaAttributes

`func (o *ResourceScope) SetSchemaAttributes(v []string)`

SetSchemaAttributes sets SchemaAttributes field to given value.

### HasSchemaAttributes

`func (o *ResourceScope) HasSchemaAttributes() bool`

HasSchemaAttributes returns a boolean if a field has been set.

### GetMappedClaims

`func (o *ResourceScope) GetMappedClaims() []string`

GetMappedClaims returns the MappedClaims field if non-nil, zero value otherwise.

### GetMappedClaimsOk

`func (o *ResourceScope) GetMappedClaimsOk() (*[]string, bool)`

GetMappedClaimsOk returns a tuple with the MappedClaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappedClaims

`func (o *ResourceScope) SetMappedClaims(v []string)`

SetMappedClaims sets MappedClaims field to given value.

### HasMappedClaims

`func (o *ResourceScope) HasMappedClaims() bool`

HasMappedClaims returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ResourceScope) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ResourceScope) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ResourceScope) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ResourceScope) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ResourceScope) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ResourceScope) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ResourceScope) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ResourceScope) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


