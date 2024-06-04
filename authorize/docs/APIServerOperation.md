# APIServerOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the API service operation. This is randomly generated when the operation is created. | [optional] [readonly] 
**Name** | **string** | The name of the API service operation. | 
**AccessControl** | Pointer to [**APIServerOperationAccessControl**](APIServerOperationAccessControl.md) |  | [optional] 
**Methods** | Pointer to [**[]EnumAPIServerOperationMethod**](EnumAPIServerOperationMethod.md) | The methods that define the operation. No duplicates are allowed. Each element must be a valid HTTP token, according to [RFC 7230](https://datatracker.ietf.org/doc/html/rfc7230), and cannot exceed 64 characters. An empty array is not valid. To indicate that an operation is defined for every method, the &#x60;methods&#x60; array should be set to null. The &#x60;methods&#x60; array is limited to 10 entries. | [optional] 
**Paths** | [**[]APIServerOperationPathsInner**](APIServerOperationPathsInner.md) | The paths that define the operation. The same literal pattern is not allowed within the same operation (the pattern of a &#x60;paths&#x60; element must be unique as compared to all other patterns in the same &#x60;paths&#x60; array). However, the same literal pattern is allowed in different operations (for example, OperationA, &#x60;/path1&#x60;, OperationB, &#x60;/path1&#x60; is valid). The &#x60;paths&#x60; array is limited to 10 entries. | 
**Policy** | Pointer to [**APIServerPolicy**](APIServerPolicy.md) |  | [optional] 

## Methods

### NewAPIServerOperation

`func NewAPIServerOperation(name string, paths []APIServerOperationPathsInner, ) *APIServerOperation`

NewAPIServerOperation instantiates a new APIServerOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerOperationWithDefaults

`func NewAPIServerOperationWithDefaults() *APIServerOperation`

NewAPIServerOperationWithDefaults instantiates a new APIServerOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIServerOperation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIServerOperation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIServerOperation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIServerOperation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *APIServerOperation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIServerOperation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIServerOperation) SetName(v string)`

SetName sets Name field to given value.


### GetAccessControl

`func (o *APIServerOperation) GetAccessControl() APIServerOperationAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *APIServerOperation) GetAccessControlOk() (*APIServerOperationAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *APIServerOperation) SetAccessControl(v APIServerOperationAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *APIServerOperation) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

### GetMethods

`func (o *APIServerOperation) GetMethods() []EnumAPIServerOperationMethod`

GetMethods returns the Methods field if non-nil, zero value otherwise.

### GetMethodsOk

`func (o *APIServerOperation) GetMethodsOk() (*[]EnumAPIServerOperationMethod, bool)`

GetMethodsOk returns a tuple with the Methods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethods

`func (o *APIServerOperation) SetMethods(v []EnumAPIServerOperationMethod)`

SetMethods sets Methods field to given value.

### HasMethods

`func (o *APIServerOperation) HasMethods() bool`

HasMethods returns a boolean if a field has been set.

### GetPaths

`func (o *APIServerOperation) GetPaths() []APIServerOperationPathsInner`

GetPaths returns the Paths field if non-nil, zero value otherwise.

### GetPathsOk

`func (o *APIServerOperation) GetPathsOk() (*[]APIServerOperationPathsInner, bool)`

GetPathsOk returns a tuple with the Paths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaths

`func (o *APIServerOperation) SetPaths(v []APIServerOperationPathsInner)`

SetPaths sets Paths field to given value.


### GetPolicy

`func (o *APIServerOperation) GetPolicy() APIServerPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *APIServerOperation) GetPolicyOk() (*APIServerPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *APIServerOperation) SetPolicy(v APIServerPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *APIServerOperation) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


