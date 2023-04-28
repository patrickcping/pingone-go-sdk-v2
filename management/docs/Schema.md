# Schema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllowsContainsOperator** | Pointer to **bool** | Indicates whether or not the &#x60;contains&#x60; operator can be used. You can use the &#x60;contains&#x60; operator in a maximum of 5 custom attributes. | [optional] [readonly] 
**Attributes** | Pointer to [**[]SchemaAttribute**](SchemaAttribute.md) |  | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of the schema. | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | Pointer to **string** | A string that specifies the resource name. | [optional] [readonly] 

## Methods

### NewSchema

`func NewSchema() *Schema`

NewSchema instantiates a new Schema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithDefaults

`func NewSchemaWithDefaults() *Schema`

NewSchemaWithDefaults instantiates a new Schema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowsContainsOperator

`func (o *Schema) GetAllowsContainsOperator() bool`

GetAllowsContainsOperator returns the AllowsContainsOperator field if non-nil, zero value otherwise.

### GetAllowsContainsOperatorOk

`func (o *Schema) GetAllowsContainsOperatorOk() (*bool, bool)`

GetAllowsContainsOperatorOk returns a tuple with the AllowsContainsOperator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsContainsOperator

`func (o *Schema) SetAllowsContainsOperator(v bool)`

SetAllowsContainsOperator sets AllowsContainsOperator field to given value.

### HasAllowsContainsOperator

`func (o *Schema) HasAllowsContainsOperator() bool`

HasAllowsContainsOperator returns a boolean if a field has been set.

### GetAttributes

`func (o *Schema) GetAttributes() []SchemaAttribute`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *Schema) GetAttributesOk() (*[]SchemaAttribute, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *Schema) SetAttributes(v []SchemaAttribute)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *Schema) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetDescription

`func (o *Schema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Schema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Schema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Schema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *Schema) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *Schema) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *Schema) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *Schema) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *Schema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schema) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Schema) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Schema) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Schema) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Schema) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Schema) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


