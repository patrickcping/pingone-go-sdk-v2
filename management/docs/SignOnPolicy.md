# SignOnPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**CreatedAt** | Pointer to **string** | The date and time the resource was created (format ISO-8061). | [optional] [readonly] 
**Default** | Pointer to **bool** | A boolean that specifies whether this sign-on policy is the environment&#39;s default that is used by applications that do not have application-specific sign-on policy assignments. This property can only be set to true, in which case the isDefault property of all other sign-on policies are set to false. | [optional] [default to false]
**Description** | Pointer to **string** | A string that specifies the description of the sign-on policy. | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the sign-on policy resource’s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies the resource name. The name must be unique within the environment, and can consist of either a string of alphanumeric letters, underscore, hyphen, period &#x60;^[a-zA-Z0-9_. -]+$&#x60; or an absolute URI if the string contains a &#x60;:&#x60; character. | 
**UpdatedAt** | Pointer to **string** | The date and time the resource was last updated (format ISO-8061). | [optional] [readonly] 

## Methods

### NewSignOnPolicy

`func NewSignOnPolicy(name string, ) *SignOnPolicy`

NewSignOnPolicy instantiates a new SignOnPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignOnPolicyWithDefaults

`func NewSignOnPolicyWithDefaults() *SignOnPolicy`

NewSignOnPolicyWithDefaults instantiates a new SignOnPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *SignOnPolicy) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SignOnPolicy) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SignOnPolicy) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SignOnPolicy) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SignOnPolicy) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SignOnPolicy) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SignOnPolicy) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SignOnPolicy) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDefault

`func (o *SignOnPolicy) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SignOnPolicy) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SignOnPolicy) SetDefault(v bool)`

SetDefault sets Default field to given value.

### HasDefault

`func (o *SignOnPolicy) HasDefault() bool`

HasDefault returns a boolean if a field has been set.

### GetDescription

`func (o *SignOnPolicy) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SignOnPolicy) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SignOnPolicy) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SignOnPolicy) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnvironment

`func (o *SignOnPolicy) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *SignOnPolicy) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *SignOnPolicy) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *SignOnPolicy) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetId

`func (o *SignOnPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SignOnPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SignOnPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SignOnPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SignOnPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SignOnPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SignOnPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetUpdatedAt

`func (o *SignOnPolicy) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SignOnPolicy) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SignOnPolicy) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SignOnPolicy) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


