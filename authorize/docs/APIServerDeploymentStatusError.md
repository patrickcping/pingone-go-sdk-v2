# APIServerDeploymentStatusError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A unique identifier for the error. | [optional] 
**Code** | Pointer to **string** | A general fault code that identifies the the type of error. See [Error codes](https://apidocs.pingidentity.com/pingone/platform/v1/api/#error-codes). | [optional] 
**Message** | Pointer to **string** | A short human-readable description of the error. | [optional] 

## Methods

### NewAPIServerDeploymentStatusError

`func NewAPIServerDeploymentStatusError() *APIServerDeploymentStatusError`

NewAPIServerDeploymentStatusError instantiates a new APIServerDeploymentStatusError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerDeploymentStatusErrorWithDefaults

`func NewAPIServerDeploymentStatusErrorWithDefaults() *APIServerDeploymentStatusError`

NewAPIServerDeploymentStatusErrorWithDefaults instantiates a new APIServerDeploymentStatusError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIServerDeploymentStatusError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIServerDeploymentStatusError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIServerDeploymentStatusError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIServerDeploymentStatusError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *APIServerDeploymentStatusError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIServerDeploymentStatusError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIServerDeploymentStatusError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *APIServerDeploymentStatusError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *APIServerDeploymentStatusError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *APIServerDeploymentStatusError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *APIServerDeploymentStatusError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *APIServerDeploymentStatusError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


