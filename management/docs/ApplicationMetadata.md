# ApplicationMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**map[string]LinksHATEOASValue**](LinksHATEOASValue.md) |  | [optional] [readonly] 
**ApplicationId** | Pointer to **string** | A string that specifies the ID of the application. | [optional] [readonly] 
**EnvironmentId** | Pointer to **string** | A string that specifies the ID of the environment. | [optional] [readonly] 
**Metadata** | Pointer to **map[string]interface{}** | User-defined custom metadata for the application. The metadata must be a JSON object and can include any custom properties. If metadata is not included in the request body, the existing application metadata will be removed from the application resource. | [optional] 

## Methods

### NewApplicationMetadata

`func NewApplicationMetadata() *ApplicationMetadata`

NewApplicationMetadata instantiates a new ApplicationMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationMetadataWithDefaults

`func NewApplicationMetadataWithDefaults() *ApplicationMetadata`

NewApplicationMetadataWithDefaults instantiates a new ApplicationMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *ApplicationMetadata) GetLinks() map[string]LinksHATEOASValue`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ApplicationMetadata) GetLinksOk() (*map[string]LinksHATEOASValue, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ApplicationMetadata) SetLinks(v map[string]LinksHATEOASValue)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ApplicationMetadata) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetApplicationId

`func (o *ApplicationMetadata) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationMetadata) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationMetadata) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationMetadata) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetEnvironmentId

`func (o *ApplicationMetadata) GetEnvironmentId() string`

GetEnvironmentId returns the EnvironmentId field if non-nil, zero value otherwise.

### GetEnvironmentIdOk

`func (o *ApplicationMetadata) GetEnvironmentIdOk() (*string, bool)`

GetEnvironmentIdOk returns a tuple with the EnvironmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentId

`func (o *ApplicationMetadata) SetEnvironmentId(v string)`

SetEnvironmentId sets EnvironmentId field to given value.

### HasEnvironmentId

`func (o *ApplicationMetadata) HasEnvironmentId() bool`

HasEnvironmentId returns a boolean if a field has been set.

### GetMetadata

`func (o *ApplicationMetadata) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ApplicationMetadata) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ApplicationMetadata) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ApplicationMetadata) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


