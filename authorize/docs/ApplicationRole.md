# ApplicationRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the application role. | [optional] 
**Id** | Pointer to **string** | The ID of the application role. | [optional] [readonly] 
**Name** | **string** | The name of the application role. | 

## Methods

### NewApplicationRole

`func NewApplicationRole(name string, ) *ApplicationRole`

NewApplicationRole instantiates a new ApplicationRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationRoleWithDefaults

`func NewApplicationRoleWithDefaults() *ApplicationRole`

NewApplicationRoleWithDefaults instantiates a new ApplicationRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationRole) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationRole) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationRole) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationRole) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationRole) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationRole) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationRole) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationRole) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *ApplicationRole) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationRole) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationRole) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ApplicationRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApplicationRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApplicationRole) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


