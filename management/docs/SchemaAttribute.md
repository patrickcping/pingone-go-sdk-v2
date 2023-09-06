# SchemaAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Description** | Pointer to **string** | A string that specifies an optional property that specifies the description of the attribute. If provided, it must not be an empty string. Valid characters consists of any Unicode letter, mark (for example, accent or umlaut), numeric character, punctuation character, or space. | [optional] 
**DisplayName** | Pointer to **string** | A string that specifies an optional property that specifies the display name of the attribute such as &#39;T-shirt size’. If provided, it must not be an empty string. Valid characters consist of any Unicode letter, mark (for example, accent or umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. | [optional] 
**Enabled** | **bool** | A boolean that specifies whether or not the attribute is enabled. This is a required property for POST and PUT operations; it cannot be omitted or explicitly set to null. Disabled attributes are ignored on create/update and not returned on read. | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**LdapAttribute** | Pointer to **string** | A string that specifies the LDAP attribute. | [optional] [readonly] 
**Name** | **string** | A string that specifies the name of the attribute. The attribute name must be provided during creation, must not be empty and must not exceed 256 characters. It must also be unique within the schema for an environment. It must start with a letter and may be followed by letters, numbers or hyphens. | 
**Required** | Pointer to **bool** | A boolean that specifies whether or not the attribute is required. Required attributes must be provided a value during create/update. Defaults to false if not provided. | [optional] 
**Schema** | Pointer to [**SchemaAttributeSchema**](SchemaAttributeSchema.md) |  | [optional] 
**SchemaType** | Pointer to [**EnumSchemaAttributeSchemaType**](EnumSchemaAttributeSchemaType.md) |  | [optional] 
**SubAttributes** | Pointer to [**[]SchemaAttribute**](SchemaAttribute.md) | The list of sub-attributes of this attribute. Only &#x60;COMPLEX&#x60; attribute types can have sub-attributes, and only one-level of nesting is allowed. The leaf attribute definition must have a type of &#x60;STRING&#x60; or &#x60;JSON&#x60;. A &#x60;COMPLEX&#x60; attribute definition must have at least one child attribute definition. | [optional] 
**Type** | [**EnumSchemaAttributeType**](EnumSchemaAttributeType.md) |  | [default to ENUMSCHEMAATTRIBUTETYPE_STRING]
**Unique** | Pointer to **bool** | A boolean that specifies whether or not the attribute must have a unique value within the environment. This is a required property for POST and PUT operations; it cannot be omitted or explicitly set to null. | [optional] 
**MultiValued** | Pointer to **bool** | A boolean that specifies whether the attribute has multiple values or a single one. This value can only change from false to true, as changing from true to false is not allowed. Maximum number of values stored is 1,000. | [optional] 
**EnumeratedValues** | Pointer to [**[]SchemaAttributeEnumeratedValuesInner**](SchemaAttributeEnumeratedValuesInner.md) |  | [optional] 
**RegexValidation** | Pointer to [**SchemaAttributeRegexValidation**](SchemaAttributeRegexValidation.md) |  | [optional] 

## Methods

### NewSchemaAttribute

`func NewSchemaAttribute(enabled bool, name string, type_ EnumSchemaAttributeType, ) *SchemaAttribute`

NewSchemaAttribute instantiates a new SchemaAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaAttributeWithDefaults

`func NewSchemaAttributeWithDefaults() *SchemaAttribute`

NewSchemaAttributeWithDefaults instantiates a new SchemaAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SchemaAttribute) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SchemaAttribute) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SchemaAttribute) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SchemaAttribute) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaAttribute) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaAttribute) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaAttribute) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaAttribute) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *SchemaAttribute) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SchemaAttribute) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SchemaAttribute) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SchemaAttribute) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *SchemaAttribute) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SchemaAttribute) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SchemaAttribute) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEnvironment

`func (o *SchemaAttribute) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SchemaAttribute) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SchemaAttribute) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SchemaAttribute) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SchemaAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SchemaAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SchemaAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SchemaAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLdapAttribute

`func (o *SchemaAttribute) GetLdapAttribute() string`

GetLdapAttribute returns the LdapAttribute field if non-nil, zero value otherwise.

### GetLdapAttributeOk

`func (o *SchemaAttribute) GetLdapAttributeOk() (*string, bool)`

GetLdapAttributeOk returns a tuple with the LdapAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLdapAttribute

`func (o *SchemaAttribute) SetLdapAttribute(v string)`

SetLdapAttribute sets LdapAttribute field to given value.

### HasLdapAttribute

`func (o *SchemaAttribute) HasLdapAttribute() bool`

HasLdapAttribute returns a boolean if a field has been set.

### GetName

`func (o *SchemaAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *SchemaAttribute) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SchemaAttribute) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SchemaAttribute) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SchemaAttribute) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchema

`func (o *SchemaAttribute) GetSchema() SchemaAttributeSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SchemaAttribute) GetSchemaOk() (*SchemaAttributeSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SchemaAttribute) SetSchema(v SchemaAttributeSchema)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *SchemaAttribute) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetSchemaType

`func (o *SchemaAttribute) GetSchemaType() EnumSchemaAttributeSchemaType`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *SchemaAttribute) GetSchemaTypeOk() (*EnumSchemaAttributeSchemaType, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *SchemaAttribute) SetSchemaType(v EnumSchemaAttributeSchemaType)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *SchemaAttribute) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetSubAttributes

`func (o *SchemaAttribute) GetSubAttributes() []SchemaAttribute`

GetSubAttributes returns the SubAttributes field if non-nil, zero value otherwise.

### GetSubAttributesOk

`func (o *SchemaAttribute) GetSubAttributesOk() (*[]SchemaAttribute, bool)`

GetSubAttributesOk returns a tuple with the SubAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAttributes

`func (o *SchemaAttribute) SetSubAttributes(v []SchemaAttribute)`

SetSubAttributes sets SubAttributes field to given value.

### HasSubAttributes

`func (o *SchemaAttribute) HasSubAttributes() bool`

HasSubAttributes returns a boolean if a field has been set.

### GetType

`func (o *SchemaAttribute) GetType() EnumSchemaAttributeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaAttribute) GetTypeOk() (*EnumSchemaAttributeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaAttribute) SetType(v EnumSchemaAttributeType)`

SetType sets Type field to given value.


### GetUnique

`func (o *SchemaAttribute) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *SchemaAttribute) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *SchemaAttribute) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *SchemaAttribute) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetMultiValued

`func (o *SchemaAttribute) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *SchemaAttribute) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *SchemaAttribute) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *SchemaAttribute) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetEnumeratedValues

`func (o *SchemaAttribute) GetEnumeratedValues() []SchemaAttributeEnumeratedValuesInner`

GetEnumeratedValues returns the EnumeratedValues field if non-nil, zero value otherwise.

### GetEnumeratedValuesOk

`func (o *SchemaAttribute) GetEnumeratedValuesOk() (*[]SchemaAttributeEnumeratedValuesInner, bool)`

GetEnumeratedValuesOk returns a tuple with the EnumeratedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumeratedValues

`func (o *SchemaAttribute) SetEnumeratedValues(v []SchemaAttributeEnumeratedValuesInner)`

SetEnumeratedValues sets EnumeratedValues field to given value.

### HasEnumeratedValues

`func (o *SchemaAttribute) HasEnumeratedValues() bool`

HasEnumeratedValues returns a boolean if a field has been set.

### GetRegexValidation

`func (o *SchemaAttribute) GetRegexValidation() SchemaAttributeRegexValidation`

GetRegexValidation returns the RegexValidation field if non-nil, zero value otherwise.

### GetRegexValidationOk

`func (o *SchemaAttribute) GetRegexValidationOk() (*SchemaAttributeRegexValidation, bool)`

GetRegexValidationOk returns a tuple with the RegexValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexValidation

`func (o *SchemaAttribute) SetRegexValidation(v SchemaAttributeRegexValidation)`

SetRegexValidation sets RegexValidation field to given value.

### HasRegexValidation

`func (o *SchemaAttribute) HasRegexValidation() bool`

HasRegexValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


