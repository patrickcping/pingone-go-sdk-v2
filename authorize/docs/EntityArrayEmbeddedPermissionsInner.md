# EntityArrayEmbeddedPermissionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Action** | **string** | The action associated with this permission. | [readonly] 
**Description** | Pointer to **string** |  | [optional] [readonly] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | **string** | The ID of the application resource permission to associate with this role. | 
**Resource** | Pointer to [**ApplicationRolePermissionResource**](ApplicationRolePermissionResource.md) |  | [optional] 
**Key** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewEntityArrayEmbeddedPermissionsInner

`func NewEntityArrayEmbeddedPermissionsInner(action string, id string, ) *EntityArrayEmbeddedPermissionsInner`

NewEntityArrayEmbeddedPermissionsInner instantiates a new EntityArrayEmbeddedPermissionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityArrayEmbeddedPermissionsInnerWithDefaults

`func NewEntityArrayEmbeddedPermissionsInnerWithDefaults() *EntityArrayEmbeddedPermissionsInner`

NewEntityArrayEmbeddedPermissionsInnerWithDefaults instantiates a new EntityArrayEmbeddedPermissionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EntityArrayEmbeddedPermissionsInner) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntityArrayEmbeddedPermissionsInner) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntityArrayEmbeddedPermissionsInner) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *EntityArrayEmbeddedPermissionsInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAction

`func (o *EntityArrayEmbeddedPermissionsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *EntityArrayEmbeddedPermissionsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *EntityArrayEmbeddedPermissionsInner) SetAction(v string)`

SetAction sets Action field to given value.


### GetDescription

`func (o *EntityArrayEmbeddedPermissionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntityArrayEmbeddedPermissionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntityArrayEmbeddedPermissionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntityArrayEmbeddedPermissionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *EntityArrayEmbeddedPermissionsInner) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *EntityArrayEmbeddedPermissionsInner) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *EntityArrayEmbeddedPermissionsInner) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *EntityArrayEmbeddedPermissionsInner) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *EntityArrayEmbeddedPermissionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityArrayEmbeddedPermissionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityArrayEmbeddedPermissionsInner) SetId(v string)`

SetId sets Id field to given value.


### GetResource

`func (o *EntityArrayEmbeddedPermissionsInner) GetResource() ApplicationRolePermissionResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *EntityArrayEmbeddedPermissionsInner) GetResourceOk() (*ApplicationRolePermissionResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *EntityArrayEmbeddedPermissionsInner) SetResource(v ApplicationRolePermissionResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *EntityArrayEmbeddedPermissionsInner) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetKey

`func (o *EntityArrayEmbeddedPermissionsInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EntityArrayEmbeddedPermissionsInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EntityArrayEmbeddedPermissionsInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *EntityArrayEmbeddedPermissionsInner) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


