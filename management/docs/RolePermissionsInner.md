# RolePermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Classifier** | Pointer to **string** | A string that specifies the resource for which the permission is applicable. | [optional] [readonly] 
**Description** | Pointer to **string** | A string that specifies the description of what the permission enables for the role. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of a permission associated with this role. | [optional] [readonly] 

## Methods

### NewRolePermissionsInner

`func NewRolePermissionsInner() *RolePermissionsInner`

NewRolePermissionsInner instantiates a new RolePermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRolePermissionsInnerWithDefaults

`func NewRolePermissionsInnerWithDefaults() *RolePermissionsInner`

NewRolePermissionsInnerWithDefaults instantiates a new RolePermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassifier

`func (o *RolePermissionsInner) GetClassifier() string`

GetClassifier returns the Classifier field if non-nil, zero value otherwise.

### GetClassifierOk

`func (o *RolePermissionsInner) GetClassifierOk() (*string, bool)`

GetClassifierOk returns a tuple with the Classifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassifier

`func (o *RolePermissionsInner) SetClassifier(v string)`

SetClassifier sets Classifier field to given value.

### HasClassifier

`func (o *RolePermissionsInner) HasClassifier() bool`

HasClassifier returns a boolean if a field has been set.

### GetDescription

`func (o *RolePermissionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RolePermissionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RolePermissionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RolePermissionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *RolePermissionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RolePermissionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RolePermissionsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RolePermissionsInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


