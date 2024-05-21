# ApplicationResourcePermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | The action associated with this permission. | 
**Description** | Pointer to **string** | The resource&#39;s description. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] 
**Resource** | Pointer to [**ApplicationResourcePermissionResource**](ApplicationResourcePermissionResource.md) |  | [optional] 

## Methods

### NewApplicationResourcePermission

`func NewApplicationResourcePermission(action string, ) *ApplicationResourcePermission`

NewApplicationResourcePermission instantiates a new ApplicationResourcePermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResourcePermissionWithDefaults

`func NewApplicationResourcePermissionWithDefaults() *ApplicationResourcePermission`

NewApplicationResourcePermissionWithDefaults instantiates a new ApplicationResourcePermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *ApplicationResourcePermission) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ApplicationResourcePermission) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ApplicationResourcePermission) SetAction(v string)`

SetAction sets Action field to given value.


### GetDescription

`func (o *ApplicationResourcePermission) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResourcePermission) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResourcePermission) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResourcePermission) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *ApplicationResourcePermission) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApplicationResourcePermission) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApplicationResourcePermission) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApplicationResourcePermission) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *ApplicationResourcePermission) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResourcePermission) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResourcePermission) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationResourcePermission) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResource

`func (o *ApplicationResourcePermission) GetResource() ApplicationResourcePermissionResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ApplicationResourcePermission) GetResourceOk() (*ApplicationResourcePermissionResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ApplicationResourcePermission) SetResource(v ApplicationResourcePermissionResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ApplicationResourcePermission) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


