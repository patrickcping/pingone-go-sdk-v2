# APIServerDeploymentStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | The deployment status code. - &#x60;POLICIES_CREATE_IN_PROGRESS&#x60; The policy bundle for the API service&#39;s managed policies is being created. - &#x60;DECISION_ENDPOINT_CREATE_IN_PROGRESS&#x60; A decision endpoint is being created for the API service. - &#x60;DECISION_ENDPOINT_UPDATE_IN_PROGRESS&#x60; The API service&#39;s decision endpoint is being updated. - &#x60;DEPLOYMENT_SUCCESSFUL&#x60; The API service&#39;s policies have been successfully deployed. - &#x60;DEPLOYMENT_FAILED&#x60; HAP-MGMT was unable to deploy the API service&#39;s policies. - &#x60;DEPLOYMENT_UNINITIALIZED&#x60; A deployment has not yet been attempted.  | [optional] 
**Error** | Pointer to [**APIServerDeploymentStatusError**](APIServerDeploymentStatusError.md) |  | [optional] 

## Methods

### NewAPIServerDeploymentStatus

`func NewAPIServerDeploymentStatus() *APIServerDeploymentStatus`

NewAPIServerDeploymentStatus instantiates a new APIServerDeploymentStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerDeploymentStatusWithDefaults

`func NewAPIServerDeploymentStatusWithDefaults() *APIServerDeploymentStatus`

NewAPIServerDeploymentStatusWithDefaults instantiates a new APIServerDeploymentStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *APIServerDeploymentStatus) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *APIServerDeploymentStatus) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *APIServerDeploymentStatus) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *APIServerDeploymentStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetError

`func (o *APIServerDeploymentStatus) GetError() APIServerDeploymentStatusError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *APIServerDeploymentStatus) GetErrorOk() (*APIServerDeploymentStatusError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *APIServerDeploymentStatus) SetError(v APIServerDeploymentStatusError)`

SetError sets Error field to given value.

### HasError

`func (o *APIServerDeploymentStatus) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


