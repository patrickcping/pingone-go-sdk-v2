# APIServerOperationAccessControlPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | An application permission ID that defines the access requirements for the operation. The end user must be entitled to the specified application permission to gain access to the operation. This is a required property if &#x60;accessControl.permission&#x60; is set. | 

## Methods

### NewAPIServerOperationAccessControlPermission

`func NewAPIServerOperationAccessControlPermission(id string, ) *APIServerOperationAccessControlPermission`

NewAPIServerOperationAccessControlPermission instantiates a new APIServerOperationAccessControlPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerOperationAccessControlPermissionWithDefaults

`func NewAPIServerOperationAccessControlPermissionWithDefaults() *APIServerOperationAccessControlPermission`

NewAPIServerOperationAccessControlPermissionWithDefaults instantiates a new APIServerOperationAccessControlPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIServerOperationAccessControlPermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIServerOperationAccessControlPermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIServerOperationAccessControlPermission) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


