# SchemaAttributeRegexValidation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pattern** | **string** | A string that specifies the regular expression to which the attribute must conform. | 
**Requirements** | **string** | A string that specifies a developer friendly description of the regular expression requirements. | 
**ValuesPatternShouldMatch** | Pointer to **[]string** | An array that specifies the list of strings matching the regular expression. | [optional] 
**ValuesPatternShouldNotMatch** | Pointer to **[]string** | An array that specifies the list of strings not matching the regular expression. | [optional] 

## Methods

### NewSchemaAttributeRegexValidation

`func NewSchemaAttributeRegexValidation(pattern string, requirements string, ) *SchemaAttributeRegexValidation`

NewSchemaAttributeRegexValidation instantiates a new SchemaAttributeRegexValidation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaAttributeRegexValidationWithDefaults

`func NewSchemaAttributeRegexValidationWithDefaults() *SchemaAttributeRegexValidation`

NewSchemaAttributeRegexValidationWithDefaults instantiates a new SchemaAttributeRegexValidation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPattern

`func (o *SchemaAttributeRegexValidation) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SchemaAttributeRegexValidation) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SchemaAttributeRegexValidation) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetRequirements

`func (o *SchemaAttributeRegexValidation) GetRequirements() string`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *SchemaAttributeRegexValidation) GetRequirementsOk() (*string, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *SchemaAttributeRegexValidation) SetRequirements(v string)`

SetRequirements sets Requirements field to given value.


### GetValuesPatternShouldMatch

`func (o *SchemaAttributeRegexValidation) GetValuesPatternShouldMatch() []string`

GetValuesPatternShouldMatch returns the ValuesPatternShouldMatch field if non-nil, zero value otherwise.

### GetValuesPatternShouldMatchOk

`func (o *SchemaAttributeRegexValidation) GetValuesPatternShouldMatchOk() (*[]string, bool)`

GetValuesPatternShouldMatchOk returns a tuple with the ValuesPatternShouldMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPatternShouldMatch

`func (o *SchemaAttributeRegexValidation) SetValuesPatternShouldMatch(v []string)`

SetValuesPatternShouldMatch sets ValuesPatternShouldMatch field to given value.

### HasValuesPatternShouldMatch

`func (o *SchemaAttributeRegexValidation) HasValuesPatternShouldMatch() bool`

HasValuesPatternShouldMatch returns a boolean if a field has been set.

### GetValuesPatternShouldNotMatch

`func (o *SchemaAttributeRegexValidation) GetValuesPatternShouldNotMatch() []string`

GetValuesPatternShouldNotMatch returns the ValuesPatternShouldNotMatch field if non-nil, zero value otherwise.

### GetValuesPatternShouldNotMatchOk

`func (o *SchemaAttributeRegexValidation) GetValuesPatternShouldNotMatchOk() (*[]string, bool)`

GetValuesPatternShouldNotMatchOk returns a tuple with the ValuesPatternShouldNotMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValuesPatternShouldNotMatch

`func (o *SchemaAttributeRegexValidation) SetValuesPatternShouldNotMatch(v []string)`

SetValuesPatternShouldNotMatch sets ValuesPatternShouldNotMatch field to given value.

### HasValuesPatternShouldNotMatch

`func (o *SchemaAttributeRegexValidation) HasValuesPatternShouldNotMatch() bool`

HasValuesPatternShouldNotMatch returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


