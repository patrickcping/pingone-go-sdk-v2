# APIServerOperations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | A string that specifies the name of the operation. | [optional] 
**Value** | Pointer to [**APIServerOperationsValue**](APIServerOperationsValue.md) |  | [optional] 

## Methods

### NewAPIServerOperations

`func NewAPIServerOperations() *APIServerOperations`

NewAPIServerOperations instantiates a new APIServerOperations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerOperationsWithDefaults

`func NewAPIServerOperationsWithDefaults() *APIServerOperations`

NewAPIServerOperationsWithDefaults instantiates a new APIServerOperations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *APIServerOperations) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *APIServerOperations) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *APIServerOperations) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *APIServerOperations) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *APIServerOperations) GetValue() APIServerOperationsValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *APIServerOperations) GetValueOk() (*APIServerOperationsValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *APIServerOperations) SetValue(v APIServerOperationsValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *APIServerOperations) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


