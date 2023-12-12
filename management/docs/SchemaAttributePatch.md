# SchemaAttributePatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | A string that specifies an optional property that specifies the description of the attribute. If provided, it must not be an empty string. Valid characters consists of any Unicode letter, mark (for example, accent or umlaut), numeric character, punctuation character, or space. | [optional] 
**DisplayName** | Pointer to **string** | A string that specifies an optional property that specifies the display name of the attribute such as &#39;T-shirt size’. If provided, it must not be an empty string. Valid characters consist of any Unicode letter, mark (for example, accent or umlaut), numeric character, forward slash, dot, apostrophe, underscore, space, or hyphen. | [optional] 
**Enabled** | Pointer to **bool** | A boolean that specifies whether or not the attribute is enabled. This is a required property for POST and PUT operations; it cannot be omitted or explicitly set to null. Disabled attributes are ignored on create/update and not returned on read. | [optional] 
**Name** | Pointer to **string** | A string that specifies the name of the attribute. The attribute name must be provided during creation, must not be empty and must not exceed 256 characters. It must also be unique within the schema for an environment. It must start with a letter and may be followed by letters, numbers or hyphens. | [optional] 
**Required** | Pointer to **bool** | A boolean that specifies whether or not the attribute is required. Required attributes must be provided a value during create/update. Defaults to false if not provided. | [optional] 
**SchemaType** | Pointer to [**EnumSchemaAttributeSchemaType**](EnumSchemaAttributeSchemaType.md) |  | [optional] 
**SubAttributes** | Pointer to [**[]SchemaAttribute**](SchemaAttribute.md) | The list of sub-attributes of this attribute. Only &#x60;COMPLEX&#x60; attribute types can have sub-attributes, and only one-level of nesting is allowed. The leaf attribute definition must have a type of &#x60;STRING&#x60; or &#x60;JSON&#x60;. A &#x60;COMPLEX&#x60; attribute definition must have at least one child attribute definition. | [optional] 
**Type** | Pointer to [**EnumSchemaAttributeType**](EnumSchemaAttributeType.md) |  | [optional] [default to ENUMSCHEMAATTRIBUTETYPE_STRING]
**Unique** | Pointer to **bool** | A boolean that specifies whether or not the attribute must have a unique value within the environment. This is a required property for POST and PUT operations; it cannot be omitted or explicitly set to null. | [optional] 
**MultiValued** | Pointer to **bool** | A boolean that specifies whether the attribute has multiple values or a single one. This value can only change from false to true, as changing from true to false is not allowed. Maximum number of values stored is 1,000. | [optional] 
**EnumeratedValues** | Pointer to [**[]SchemaAttributeEnumeratedValuesInner**](SchemaAttributeEnumeratedValuesInner.md) |  | [optional] 
**RegexValidation** | Pointer to [**SchemaAttributeRegexValidation**](SchemaAttributeRegexValidation.md) |  | [optional] 

## Methods

### NewSchemaAttributePatch

`func NewSchemaAttributePatch() *SchemaAttributePatch`

NewSchemaAttributePatch instantiates a new SchemaAttributePatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaAttributePatchWithDefaults

`func NewSchemaAttributePatchWithDefaults() *SchemaAttributePatch`

NewSchemaAttributePatchWithDefaults instantiates a new SchemaAttributePatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SchemaAttributePatch) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaAttributePatch) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaAttributePatch) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaAttributePatch) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisplayName

`func (o *SchemaAttributePatch) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *SchemaAttributePatch) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *SchemaAttributePatch) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *SchemaAttributePatch) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnabled

`func (o *SchemaAttributePatch) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SchemaAttributePatch) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SchemaAttributePatch) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SchemaAttributePatch) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetName

`func (o *SchemaAttributePatch) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SchemaAttributePatch) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SchemaAttributePatch) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SchemaAttributePatch) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRequired

`func (o *SchemaAttributePatch) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SchemaAttributePatch) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SchemaAttributePatch) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SchemaAttributePatch) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetSchemaType

`func (o *SchemaAttributePatch) GetSchemaType() EnumSchemaAttributeSchemaType`

GetSchemaType returns the SchemaType field if non-nil, zero value otherwise.

### GetSchemaTypeOk

`func (o *SchemaAttributePatch) GetSchemaTypeOk() (*EnumSchemaAttributeSchemaType, bool)`

GetSchemaTypeOk returns a tuple with the SchemaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaType

`func (o *SchemaAttributePatch) SetSchemaType(v EnumSchemaAttributeSchemaType)`

SetSchemaType sets SchemaType field to given value.

### HasSchemaType

`func (o *SchemaAttributePatch) HasSchemaType() bool`

HasSchemaType returns a boolean if a field has been set.

### GetSubAttributes

`func (o *SchemaAttributePatch) GetSubAttributes() []SchemaAttribute`

GetSubAttributes returns the SubAttributes field if non-nil, zero value otherwise.

### GetSubAttributesOk

`func (o *SchemaAttributePatch) GetSubAttributesOk() (*[]SchemaAttribute, bool)`

GetSubAttributesOk returns a tuple with the SubAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAttributes

`func (o *SchemaAttributePatch) SetSubAttributes(v []SchemaAttribute)`

SetSubAttributes sets SubAttributes field to given value.

### HasSubAttributes

`func (o *SchemaAttributePatch) HasSubAttributes() bool`

HasSubAttributes returns a boolean if a field has been set.

### GetType

`func (o *SchemaAttributePatch) GetType() EnumSchemaAttributeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SchemaAttributePatch) GetTypeOk() (*EnumSchemaAttributeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SchemaAttributePatch) SetType(v EnumSchemaAttributeType)`

SetType sets Type field to given value.

### HasType

`func (o *SchemaAttributePatch) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnique

`func (o *SchemaAttributePatch) GetUnique() bool`

GetUnique returns the Unique field if non-nil, zero value otherwise.

### GetUniqueOk

`func (o *SchemaAttributePatch) GetUniqueOk() (*bool, bool)`

GetUniqueOk returns a tuple with the Unique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnique

`func (o *SchemaAttributePatch) SetUnique(v bool)`

SetUnique sets Unique field to given value.

### HasUnique

`func (o *SchemaAttributePatch) HasUnique() bool`

HasUnique returns a boolean if a field has been set.

### GetMultiValued

`func (o *SchemaAttributePatch) GetMultiValued() bool`

GetMultiValued returns the MultiValued field if non-nil, zero value otherwise.

### GetMultiValuedOk

`func (o *SchemaAttributePatch) GetMultiValuedOk() (*bool, bool)`

GetMultiValuedOk returns a tuple with the MultiValued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValued

`func (o *SchemaAttributePatch) SetMultiValued(v bool)`

SetMultiValued sets MultiValued field to given value.

### HasMultiValued

`func (o *SchemaAttributePatch) HasMultiValued() bool`

HasMultiValued returns a boolean if a field has been set.

### GetEnumeratedValues

`func (o *SchemaAttributePatch) GetEnumeratedValues() []SchemaAttributeEnumeratedValuesInner`

GetEnumeratedValues returns the EnumeratedValues field if non-nil, zero value otherwise.

### GetEnumeratedValuesOk

`func (o *SchemaAttributePatch) GetEnumeratedValuesOk() (*[]SchemaAttributeEnumeratedValuesInner, bool)`

GetEnumeratedValuesOk returns a tuple with the EnumeratedValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumeratedValues

`func (o *SchemaAttributePatch) SetEnumeratedValues(v []SchemaAttributeEnumeratedValuesInner)`

SetEnumeratedValues sets EnumeratedValues field to given value.

### HasEnumeratedValues

`func (o *SchemaAttributePatch) HasEnumeratedValues() bool`

HasEnumeratedValues returns a boolean if a field has been set.

### GetRegexValidation

`func (o *SchemaAttributePatch) GetRegexValidation() SchemaAttributeRegexValidation`

GetRegexValidation returns the RegexValidation field if non-nil, zero value otherwise.

### GetRegexValidationOk

`func (o *SchemaAttributePatch) GetRegexValidationOk() (*SchemaAttributeRegexValidation, bool)`

GetRegexValidationOk returns a tuple with the RegexValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegexValidation

`func (o *SchemaAttributePatch) SetRegexValidation(v SchemaAttributeRegexValidation)`

SetRegexValidation sets RegexValidation field to given value.

### HasRegexValidation

`func (o *SchemaAttributePatch) HasRegexValidation() bool`

HasRegexValidation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


