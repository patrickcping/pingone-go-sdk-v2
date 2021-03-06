# ResourceAttribute

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | A string that specifies the resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies the name of the custom resource attribute to be included in the access token | 
**Type** | Pointer to [**EnumResourceAttributeType**](EnumResourceAttributeType.md) |  | [optional] 
**Value** | **string** | A string that specifies the value of the custom resource attribute. This value can be a placeholder that references an attribute in the user schema, expressed as “${user.path.to.value}”, or it can be a static string. Placeholders must be valid, enabled attributes in the environment’s user schema. Examples fo valid values are “${user.email}”, “${user.name.family}”, and “myClaimValueString” | 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Resource** | Pointer to [**ResourceResource**](ResourceResource.md) |  | [optional] 

## Methods

### NewResourceAttribute

`func NewResourceAttribute(name string, value string, ) *ResourceAttribute`

NewResourceAttribute instantiates a new ResourceAttribute object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAttributeWithDefaults

`func NewResourceAttributeWithDefaults() *ResourceAttribute`

NewResourceAttributeWithDefaults instantiates a new ResourceAttribute object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceAttribute) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceAttribute) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceAttribute) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceAttribute) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceAttribute) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceAttribute) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceAttribute) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *ResourceAttribute) GetType() EnumResourceAttributeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceAttribute) GetTypeOk() (*EnumResourceAttributeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceAttribute) SetType(v EnumResourceAttributeType)`

SetType sets Type field to given value.

### HasType

`func (o *ResourceAttribute) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *ResourceAttribute) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceAttribute) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceAttribute) SetValue(v string)`

SetValue sets Value field to given value.


### GetEnvironment

`func (o *ResourceAttribute) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ResourceAttribute) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ResourceAttribute) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ResourceAttribute) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetResource

`func (o *ResourceAttribute) GetResource() ResourceResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceAttribute) GetResourceOk() (*ResourceResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceAttribute) SetResource(v ResourceResource)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceAttribute) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


