# APIServerOperationsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Methods** | Pointer to **[]string** | An array that specifies the methods that define the operation. No duplicates are allowed. Each element must be a valid HTTP token, according to RFC 7230, and cannot exceed 64 characters. An empty array is not valid. To indicate that an operation is defined for every method, the methods array should be set to null. The methods array is limited to 10 entries. | [optional] 
**Paths** | [**[]APIServerOperationsValuePathsInner**](APIServerOperationsValuePathsInner.md) | An array that specifies the paths that define the operation. This is a required property when an operations object is specified. The same literal pattern is not allowed within the same operation (the pattern of a paths element must be unique as compared to all other patterns in the same paths array). However, the same literal pattern is allowed in different operations (for example, OperationA, /path1, OperationB, /path1 is valid). The paths array is limited to 10 entries. | 
**AccessControl** | Pointer to [**APIServerOperationsValueAccessControl**](APIServerOperationsValueAccessControl.md) |  | [optional] 

## Methods

### NewAPIServerOperationsValue

`func NewAPIServerOperationsValue(paths []APIServerOperationsValuePathsInner, ) *APIServerOperationsValue`

NewAPIServerOperationsValue instantiates a new APIServerOperationsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerOperationsValueWithDefaults

`func NewAPIServerOperationsValueWithDefaults() *APIServerOperationsValue`

NewAPIServerOperationsValueWithDefaults instantiates a new APIServerOperationsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMethods

`func (o *APIServerOperationsValue) GetMethods() []string`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *APIServerOperationsValue) GetMethodsOk() (*[]string, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *APIServerOperationsValue) SetMethods(v []string)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *APIServerOperationsValue) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetPaths

`func (o *APIServerOperationsValue) GetPaths() []APIServerOperationsValuePathsInner`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *APIServerOperationsValue) GetPathsOk() (*[]APIServerOperationsValuePathsInner, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *APIServerOperationsValue) SetPaths(v []APIServerOperationsValuePathsInner)`

SetPaths sets Paths field to given value.


### GetAccessControl

`func (o *APIServerOperationsValue) GetAccessControl() APIServerOperationsValueAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *APIServerOperationsValue) GetAccessControlOk() (*APIServerOperationsValueAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *APIServerOperationsValue) SetAccessControl(v APIServerOperationsValueAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *APIServerOperationsValue) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


