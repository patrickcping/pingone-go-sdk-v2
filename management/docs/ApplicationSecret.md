# ApplicationSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**Environment** | Pointer to [**ObjectEnvironment**](ObjectEnvironment.md) |  | [optional] 
**Secret** | Pointer to **string** | A string that specifies the application secret ID used to authenticate to the authorization server. | [optional] [readonly] 
**Previous** | Pointer to [**ApplicationSecretPrevious**](ApplicationSecretPrevious.md) |  | [optional] 

## Methods

### NewApplicationSecret

`func NewApplicationSecret() *ApplicationSecret`

NewApplicationSecret instantiates a new ApplicationSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSecretWithDefaults

`func NewApplicationSecretWithDefaults() *ApplicationSecret`

NewApplicationSecretWithDefaults instantiates a new ApplicationSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationSecret) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationSecret) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationSecret) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationSecret) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetEnvironment

`func (o *ApplicationSecret) GetEnvironment() ObjectEnvironment`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *ApplicationSecret) GetEnvironmentOk() (*ObjectEnvironment, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *ApplicationSecret) SetEnvironment(v ObjectEnvironment)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *ApplicationSecret) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetSecret

`func (o *ApplicationSecret) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ApplicationSecret) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ApplicationSecret) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ApplicationSecret) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetPrevious

`func (o *ApplicationSecret) GetPrevious() ApplicationSecretPrevious`

GetPrevious returns the Previous field if non-nil, zero value otherwise.

### GetPreviousOk

`func (o *ApplicationSecret) GetPreviousOk() (*ApplicationSecretPrevious, bool)`

GetPreviousOk returns a tuple with the Previous field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrevious

`func (o *ApplicationSecret) SetPrevious(v ApplicationSecretPrevious)`

SetPrevious sets Previous field to given value.

### HasPrevious

`func (o *ApplicationSecret) HasPrevious() bool`

HasPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


