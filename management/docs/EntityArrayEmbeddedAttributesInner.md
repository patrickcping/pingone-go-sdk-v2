# EntityArrayEmbeddedAttributesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | The time the resource was created. | [optional] [readonly] 
**MappingType** | Pointer to [**EnumIdentityProviderAttributeMappingType**](EnumIdentityProviderAttributeMappingType.md) |  | [optional] 
**Name** | **string** | A string that specifies the name of the custom resource attribute to be included in the access token | 
**Required** | **bool** | A boolean that specifies whether or not the attribute is required. Required attributes must be provided a value during create/update. Defaults to false if not provided. | 
**UpdatedAt** | Pointer to **string** | The time the resource was last updated. | [optional] [readonly] 
**Value** | **string** | A string that specifies the value of the custom resource attribute. This value can be a placeholder that references an attribute in the user schema, expressed as “${user.path.to.value}”, or it can be a static string. Placeholders must be valid, enabled attributes in the environment’s user schema. Examples fo valid values are “${user.email}”, “${user.name.family}”, and “myClaimValueString” | 
**Update** | [**EnumIdentityProviderAttributeMappingType**](EnumIdentityProviderAttributeMappingType.md) |  | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**IdentityProvider** | Pointer to [**IdentityProviderAttributeIdentityProvider**](IdentityProviderAttributeIdentityProvider.md) |  | [optional] 
**Description** | Pointer to **string** | A string that specifies an optional property that specifies the description of the attribute. If provided, it must not be an empty string. Valid characters consists of any Unicode letter, mark (for example, accent or umlaut), numeric character, punctuation character, or space. | [optional] 
**DisplayName** | Pointer to **string** | A string that specifies an optional property that specifies the display name of the attribute such as &#39;T-shirt size’. If provided, it must not be an empty string. Valid characters consist of any Unicode letter, mark (for example, accent or umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. | [optional] 
**Enabled** | **bool** | A boolean that specifies whether or not the attribute is enabled. This is a required property for POST and PUT operations; it cannot be omitted or explicitly set to null. Disabled attributes are ignored on create/update and not returned on read. | 
**LdapAttribute** | Pointer to **string** | A string that specifies the LDAP attribute. | [optional] [readonly] 
**Schema** | Pointer to [**SchemaAttributeSchema**](SchemaAttributeSchema.md) |  | [optional] 
**SchemaType** | Pointer to [**EnumSchemaAttributeSchemaType**](EnumSchemaAttributeSchemaType.md) |  | [optional] 
**Type** | [**EnumResourceAttributeType**](EnumResourceAttributeType.md) |  | 
**Unique** | Pointer to **bool** | A boolean that specifies whether or not the attribute must have a unique value within the environment. This is a required property for POST and PUT operations; it cannot be omitted or explicitly set to null. | [optional] 
**MultiValued** | Pointer to **bool** | A boolean that specifies whether the attribute has multiple values or a single one. This value can only change from false to true, as changing from true to false is not allowed. Maximum number of values stored is 1,000. | [optional] 
**Resource** | Pointer to [**IdentityProviderAttributeIdentityProvider**](IdentityProviderAttributeIdentityProvider.md) |  | [optional] 

## Methods

### NewEntityArrayEmbeddedAttributesInner

`func NewEntityArrayEmbeddedAttributesInner(name string, required bool, value string, update EnumIdentityProviderAttributeMappingType, enabled bool, type_ EnumResourceAttributeType, ) *EntityArrayEmbeddedAttributesInner`

NewEntityArrayEmbeddedAttributesInner instantiates a new EntityArrayEmbeddedAttributesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedAttributesInnerWithDefaults

`func NewEntityArrayEmbeddedAttributesInnerWithDefaults() *EntityArrayEmbeddedAttributesInner`

NewEntityArrayEmbeddedAttributesInnerWithDefaults instantiates a new EntityArrayEmbeddedAttributesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityArrayEmbeddedAttributesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityArrayEmbeddedAttributesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityArrayEmbeddedAttributesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntityArrayEmbeddedAttributesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *EntityArrayEmbeddedAttributesInner) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *EntityArrayEmbeddedAttributesInner) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *EntityArrayEmbeddedAttributesInner) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *EntityArrayEmbeddedAttributesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMappingType

`func (o *EntityArrayEmbeddedAttributesInner) GetMappingType() EnumIdentityProviderAttributeMappingType`

GetMappingType returns the MappingType field if non-nil, zero value otherwise.

### GetMappingTypeOk

`func (o *EntityArrayEmbeddedAttributesInner) GetMappingTypeOk() (*EnumIdentityProviderAttributeMappingType, bool)`

GetMappingTypeOk returns a tuple with the MappingType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappingType

`func (o *EntityArrayEmbeddedAttributesInner) SetMappingType(v EnumIdentityProviderAttributeMappingType)`

SetMappingType sets MappingType field to given value.

### HasMappingType

`func (o *EntityArrayEmbeddedAttributesInner) HasMappingType() bool`

HasMappingType returns a boolean if a field has been set.

### GetName

`func (o *EntityArrayEmbeddedAttributesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntityArrayEmbeddedAttributesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntityArrayEmbeddedAttributesInner) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *EntityArrayEmbeddedAttributesInner) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EntityArrayEmbeddedAttributesInner) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EntityArrayEmbeddedAttributesInner) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetUpdatedAt

`func (o *EntityArrayEmbeddedAttributesInner) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *EntityArrayEmbeddedAttributesInner) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *EntityArrayEmbeddedAttributesInner) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *EntityArrayEmbeddedAttributesInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetValue

`func (o *EntityArrayEmbeddedAttributesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntityArrayEmbeddedAttributesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntityArrayEmbeddedAttributesInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetUpdate

`func (o *EntityArrayEmbeddedAttributesInner) GetUpdate() EnumIdentityProviderAttributeMappingType`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *EntityArrayEmbeddedAttributesInner) GetUpdateOk() (*EnumIdentityProviderAttributeMappingType, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *EntityArrayEmbeddedAttributesInner) SetUpdate(v EnumIdentityProviderAttributeMappingType)`

SetUpdate sets Update field to given value.


### GetEnvironment

`func (o *EntityArrayEmbeddedAttributesInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EntityArrayEmbeddedAttributesInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EntityArrayEmbeddedAttributesInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EntityArrayEmbeddedAttributesInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *EntityArrayEmbeddedAttributesInner) GetIdentityProvider() IdentityProviderAttributeIdentityProvider`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *EntityArrayEmbeddedAttributesInner) GetIdentityProviderOk() (*IdentityProviderAttributeIdentityProvider, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *EntityArrayEmbeddedAttributesInner) SetIdentityProvider(v IdentityProviderAttributeIdentityProvider)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *EntityArrayEmbeddedAttributesInner) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetDescription

`func (o *EntityArrayEmbeddedAttributesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityArrayEmbeddedAttributesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityArrayEmbeddedAttributesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityArrayEmbeddedAttributesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *EntityArrayEmbeddedAttributesInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *EntityArrayEmbeddedAttributesInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *EntityArrayEmbeddedAttributesInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *EntityArrayEmbeddedAttributesInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *EntityArrayEmbeddedAttributesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EntityArrayEmbeddedAttributesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EntityArrayEmbeddedAttributesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLdapAttribute

`func (o *EntityArrayEmbeddedAttributesInner) GetLdapAttribute() string`

GetLdapAttribute returns the LdapAttribute field if non-nil, zero value otherwise.

### GetLdapAttributeOk

`func (o *EntityArrayEmbeddedAttributesInner) GetLdapAttributeOk() (*string, bool)`

GetLdapAttributeOk returns a tuple with the LdapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttribute

`func (o *EntityArrayEmbeddedAttributesInner) SetLdapAttribute(v string)`

SetLdapAttribute sets LdapAttribute field to given value.

### HasLdapAttribute

`func (o *EntityArrayEmbeddedAttributesInner) HasLdapAttribute() bool`

HasLdapAttribute returns a boolean if a field has been set.

### GetSchema

`func (o *EntityArrayEmbeddedAttributesInner) GetSchema() SchemaAttributeSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *EntityArrayEmbeddedAttributesInner) GetSchemaOk() (*SchemaAttributeSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *EntityArrayEmbeddedAttributesInner) SetSchema(v SchemaAttributeSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *EntityArrayEmbeddedAttributesInner) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSchemaType

`func (o *EntityArrayEmbeddedAttributesInner) GetSchemaType() EnumSchemaAttributeSchemaType`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *EntityArrayEmbeddedAttributesInner) GetSchemaTypeOk() (*EnumSchemaAttributeSchemaType, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *EntityArrayEmbeddedAttributesInner) SetSchemaType(v EnumSchemaAttributeSchemaType)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *EntityArrayEmbeddedAttributesInner) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetType

`func (o *EntityArrayEmbeddedAttributesInner) GetType() EnumResourceAttributeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EntityArrayEmbeddedAttributesInner) GetTypeOk() (*EnumResourceAttributeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EntityArrayEmbeddedAttributesInner) SetType(v EnumResourceAttributeType)`

SetType sets Type field to given value.


### GetUnique

`func (o *EntityArrayEmbeddedAttributesInner) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *EntityArrayEmbeddedAttributesInner) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *EntityArrayEmbeddedAttributesInner) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *EntityArrayEmbeddedAttributesInner) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetMultiValued

`func (o *EntityArrayEmbeddedAttributesInner) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *EntityArrayEmbeddedAttributesInner) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *EntityArrayEmbeddedAttributesInner) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *EntityArrayEmbeddedAttributesInner) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetResource

`func (o *EntityArrayEmbeddedAttributesInner) GetResource() IdentityProviderAttributeIdentityProvider`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EntityArrayEmbeddedAttributesInner) GetResourceOk() (*IdentityProviderAttributeIdentityProvider, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EntityArrayEmbeddedAttributesInner) SetResource(v IdentityProviderAttributeIdentityProvider)`

SetResource sets Resource field to given value.

### HasResource

`func (o *EntityArrayEmbeddedAttributesInner) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


