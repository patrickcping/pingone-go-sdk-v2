# IdentityProviderAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MappingType** | Pointer to [**EnumIdentityProviderAttributeMappingType**](EnumIdentityProviderAttributeMappingType.md) |  | [optional] 
**Name** | **string** | The user attribute, which is unique per provider. The attribute must not be defined as read only from the user schema or of type COMPLEX based on the user schema. Valid examples username, and name.first. The following attributes may not be used account, id, created, updated, lifecycle, mfaEnabled, and enabled. | 
**Value** | **string** | A placeholder referring to the attribute (or attributes) from the provider. Placeholders must be valid for the attributes returned by the IdP type and use the ${} syntax (for example, username&#x3D;\&quot;${email}\&quot;). For SAML, any placeholder is acceptable, and it is mapped against the attributes available in the SAML assertion after authentication. The ${samlAssertion.subject} placeholder is a special reserved placeholder used to refer to the subject name ID in the SAML assertion response. | 
**Update** | [**EnumIdentityProviderAttributeMappingType**](EnumIdentityProviderAttributeMappingType.md) |  | 
**Id** | Pointer to **string** | The unique identifier for the resource. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**IdentityProvider** | Pointer to [**IdentityProviderAttributeIdentityProvider**](IdentityProviderAttributeIdentityProvider.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewIdentityProviderAttribute

`func NewIdentityProviderAttribute(name string, value string, update EnumIdentityProviderAttributeMappingType, ) *IdentityProviderAttribute`

NewIdentityProviderAttribute instantiates a new IdentityProviderAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdentityProviderAttributeWithDefaults

`func NewIdentityProviderAttributeWithDefaults() *IdentityProviderAttribute`

NewIdentityProviderAttributeWithDefaults instantiates a new IdentityProviderAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMappingType

`func (o *IdentityProviderAttribute) GetMappingType() EnumIdentityProviderAttributeMappingType`

GetMappingType returns the MappingType field if non-nil, zero value otherwise.

### GetMappingTypeOk

`func (o *IdentityProviderAttribute) GetMappingTypeOk() (*EnumIdentityProviderAttributeMappingType, bool)`

GetMappingTypeOk returns a tuple with the MappingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingType

`func (o *IdentityProviderAttribute) SetMappingType(v EnumIdentityProviderAttributeMappingType)`

SetMappingType sets MappingType field to given value.

### HasMappingType

`func (o *IdentityProviderAttribute) HasMappingType() bool`

HasMappingType returns a boolean if a field has been set.

### GetName

`func (o *IdentityProviderAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IdentityProviderAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IdentityProviderAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *IdentityProviderAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *IdentityProviderAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *IdentityProviderAttribute) SetValue(v string)`

SetValue sets Value field to given value.


### GetUpdate

`func (o *IdentityProviderAttribute) GetUpdate() EnumIdentityProviderAttributeMappingType`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *IdentityProviderAttribute) GetUpdateOk() (*EnumIdentityProviderAttributeMappingType, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *IdentityProviderAttribute) SetUpdate(v EnumIdentityProviderAttributeMappingType)`

SetUpdate sets Update field to given value.


### GetId

`func (o *IdentityProviderAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdentityProviderAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdentityProviderAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IdentityProviderAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetEnvironment

`func (o *IdentityProviderAttribute) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *IdentityProviderAttribute) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *IdentityProviderAttribute) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *IdentityProviderAttribute) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *IdentityProviderAttribute) GetIdentityProvider() IdentityProviderAttributeIdentityProvider`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *IdentityProviderAttribute) GetIdentityProviderOk() (*IdentityProviderAttributeIdentityProvider, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *IdentityProviderAttribute) SetIdentityProvider(v IdentityProviderAttributeIdentityProvider)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *IdentityProviderAttribute) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetCreatedAt

`func (o *IdentityProviderAttribute) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *IdentityProviderAttribute) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *IdentityProviderAttribute) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *IdentityProviderAttribute) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *IdentityProviderAttribute) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *IdentityProviderAttribute) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *IdentityProviderAttribute) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *IdentityProviderAttribute) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


