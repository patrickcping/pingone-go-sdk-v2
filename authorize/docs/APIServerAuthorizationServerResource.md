# APIServerAuthorizationServerResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A string that specifies the UUID of the custom PingOne resource. This property must identify a PingOne resource with a type property value of CUSTOM. | 

## Methods

### NewAPIServerAuthorizationServerResource

`func NewAPIServerAuthorizationServerResource(id string, ) *APIServerAuthorizationServerResource`

NewAPIServerAuthorizationServerResource instantiates a new APIServerAuthorizationServerResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerAuthorizationServerResourceWithDefaults

`func NewAPIServerAuthorizationServerResourceWithDefaults() *APIServerAuthorizationServerResource`

NewAPIServerAuthorizationServerResourceWithDefaults instantiates a new APIServerAuthorizationServerResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *APIServerAuthorizationServerResource) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIServerAuthorizationServerResource) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIServerAuthorizationServerResource) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


