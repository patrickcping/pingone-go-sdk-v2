# ApplicationResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Description** | Pointer to **string** | The application resource&#39;s description. | [optional] 
**Id** | Pointer to **string** | The resource&#39;s unique identifier. | [optional] [readonly] 
**Name** | **string** | The application resource name. The name value must be unique. | 
**Parent** | Pointer to [**ApplicationResourceParent**](ApplicationResourceParent.md) |  | [optional] 

## Methods

### NewApplicationResource

`func NewApplicationResource(name string, ) *ApplicationResource`

NewApplicationResource instantiates a new ApplicationResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResourceWithDefaults

`func NewApplicationResourceWithDefaults() *ApplicationResource`

NewApplicationResourceWithDefaults instantiates a new ApplicationResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationResource) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationResource) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationResource) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationResource) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationResource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationResource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationResource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationResource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApplicationResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResource) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationResource) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationResource) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *ApplicationResource) GetParent() ApplicationResourceParent`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ApplicationResource) GetParentOk() (*ApplicationResourceParent, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ApplicationResource) SetParent(v ApplicationResourceParent)`

SetParent sets Parent field to given value.

### HasParent

`func (o *ApplicationResource) HasParent() bool`

HasParent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


