# SchemaAttributeEnumeratedValuesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | **string** | A string that specifies the immutable value. Values are case sensitive; two values that differ only by case are not allowed. | 
**Archived** | Pointer to **bool** | A boolean that specifies whether the enumerated value is archived. Archived values cannot be added to a user, but existing archived values are preserved. This allows clients that read the schema to know all possible values of an attribute. | [optional] 
**Description** | Pointer to **string** | A string that specifies the description of the enumerated value. | [optional] 

## Methods

### NewSchemaAttributeEnumeratedValuesInner

`func NewSchemaAttributeEnumeratedValuesInner(value string, ) *SchemaAttributeEnumeratedValuesInner`

NewSchemaAttributeEnumeratedValuesInner instantiates a new SchemaAttributeEnumeratedValuesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaAttributeEnumeratedValuesInnerWithDefaults

`func NewSchemaAttributeEnumeratedValuesInnerWithDefaults() *SchemaAttributeEnumeratedValuesInner`

NewSchemaAttributeEnumeratedValuesInnerWithDefaults instantiates a new SchemaAttributeEnumeratedValuesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *SchemaAttributeEnumeratedValuesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SchemaAttributeEnumeratedValuesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SchemaAttributeEnumeratedValuesInner) SetValue(v string)`

SetValue sets Value field to given value.


### GetArchived

`func (o *SchemaAttributeEnumeratedValuesInner) GetArchived() bool`

GetArchived returns the Archived field if non-nil, zero value otherwise.

### GetArchivedOk

`func (o *SchemaAttributeEnumeratedValuesInner) GetArchivedOk() (*bool, bool)`

GetArchivedOk returns a tuple with the Archived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchived

`func (o *SchemaAttributeEnumeratedValuesInner) SetArchived(v bool)`

SetArchived sets Archived field to given value.

### HasArchived

`func (o *SchemaAttributeEnumeratedValuesInner) HasArchived() bool`

HasArchived returns a boolean if a field has been set.

### GetDescription

`func (o *SchemaAttributeEnumeratedValuesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SchemaAttributeEnumeratedValuesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SchemaAttributeEnumeratedValuesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SchemaAttributeEnumeratedValuesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


