# ResourceApplicationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The application resource&#39;s description. | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] 
**Name** | **string** | The application resource name. The name value must be unique. | 
**Parent** | Pointer to [**ResourceApplicationResourceParent**](ResourceApplicationResourceParent.md) |  | [optional] 

## Methods

### NewResourceApplicationResource

`func NewResourceApplicationResource(name string, ) *ResourceApplicationResource`

NewResourceApplicationResource instantiates a new ResourceApplicationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceApplicationResourceWithDefaults

`func NewResourceApplicationResourceWithDefaults() *ResourceApplicationResource`

NewResourceApplicationResourceWithDefaults instantiates a new ResourceApplicationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ResourceApplicationResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceApplicationResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceApplicationResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceApplicationResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ResourceApplicationResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceApplicationResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceApplicationResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceApplicationResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ResourceApplicationResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceApplicationResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceApplicationResource) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *ResourceApplicationResource) GetParent() ResourceApplicationResourceParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ResourceApplicationResource) GetParentOk() (*ResourceApplicationResourceParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ResourceApplicationResource) SetParent(v ResourceApplicationResourceParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ResourceApplicationResource) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


