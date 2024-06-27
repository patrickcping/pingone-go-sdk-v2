# CredentialTypeMultiple

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expression** | **string** | A PingOne Expression Language (PEL) expression evaluated by the P1 Credentials service on issuance.If an array, calculates the array length for the count. Populates the limit to a variable, __ITERATOR__, available to PEL expressions in metadata.fields.attribute. | 

## Methods

### NewCredentialTypeMultiple

`func NewCredentialTypeMultiple(expression string, ) *CredentialTypeMultiple`

NewCredentialTypeMultiple instantiates a new CredentialTypeMultiple object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCredentialTypeMultipleWithDefaults

`func NewCredentialTypeMultipleWithDefaults() *CredentialTypeMultiple`

NewCredentialTypeMultipleWithDefaults instantiates a new CredentialTypeMultiple object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpression

`func (o *CredentialTypeMultiple) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *CredentialTypeMultiple) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *CredentialTypeMultiple) SetExpression(v string)`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


