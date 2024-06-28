# ApplicationResourceGrant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**Application** | Pointer to [**ApplicationResourceGrantApplication**](ApplicationResourceGrantApplication.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** | The time the resource was created. | [optional] [readonly] 
**Id** | Pointer to **string** | A string that specifies the application resource grant ID. | [optional] [readonly] 
**Resource** | [**ApplicationResourceGrantResource**](ApplicationResourceGrantResource.md) |  | 
**Scopes** | [**[]ApplicationResourceGrantScopesInner**](ApplicationResourceGrantScopesInner.md) |  | 
**UpdatedAt** | Pointer to **time.Time** | The time the resource was last updated. | [optional] [readonly] 

## Methods

### NewApplicationResourceGrant

`func NewApplicationResourceGrant(resource ApplicationResourceGrantResource, scopes []ApplicationResourceGrantScopesInner, ) *ApplicationResourceGrant`

NewApplicationResourceGrant instantiates a new ApplicationResourceGrant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationResourceGrantWithDefaults

`func NewApplicationResourceGrantWithDefaults() *ApplicationResourceGrant`

NewApplicationResourceGrantWithDefaults instantiates a new ApplicationResourceGrant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationResourceGrant) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationResourceGrant) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationResourceGrant) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationResourceGrant) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetApplication

`func (o *ApplicationResourceGrant) GetApplication() ApplicationResourceGrantApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApplicationResourceGrant) GetApplicationOk() (*ApplicationResourceGrantApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApplicationResourceGrant) SetApplication(v ApplicationResourceGrantApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *ApplicationResourceGrant) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApplicationResourceGrant) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApplicationResourceGrant) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApplicationResourceGrant) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApplicationResourceGrant) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *ApplicationResourceGrant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApplicationResourceGrant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApplicationResourceGrant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApplicationResourceGrant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetResource

`func (o *ApplicationResourceGrant) GetResource() ApplicationResourceGrantResource`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ApplicationResourceGrant) GetResourceOk() (*ApplicationResourceGrantResource, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ApplicationResourceGrant) SetResource(v ApplicationResourceGrantResource)`

SetResource sets Resource field to given value.


### GetScopes

`func (o *ApplicationResourceGrant) GetScopes() []ApplicationResourceGrantScopesInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ApplicationResourceGrant) GetScopesOk() (*[]ApplicationResourceGrantScopesInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ApplicationResourceGrant) SetScopes(v []ApplicationResourceGrantScopesInner)`

SetScopes sets Scopes field to given value.


### GetUpdatedAt

`func (o *ApplicationResourceGrant) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ApplicationResourceGrant) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ApplicationResourceGrant) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ApplicationResourceGrant) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


