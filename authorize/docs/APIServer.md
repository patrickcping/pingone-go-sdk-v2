# APIServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**AuthorizationServer** | [**APIServerAuthorizationServer**](APIServerAuthorizationServer.md) |  | 
**BaseURLs** | **[]string** | An array of string that specifies the possible base URLs that an end-user will use to access the APIs hosted on the customer&#39;s API server. Multiple base URLs may be specified to support cases where the same API may be available from multiple URLs (for example, from a user-friendly domain URL and an internal domain URL). Base URLs must be valid absolute URLs with the https or http scheme. If the path component is non-empty, it must not end in a trailing slash. The path must not contain empty backslash, dot, or double-dot segments. It must not have a query or fragment present, and the host portion of the authority must be a DNS hostname or valid IP (IPv4 or IPv6). The length must be less than or equal to 256 characters. | 
**Id** | Pointer to **string** | A string that specifies the resource&#39;s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies the API server resource name. The name value must be unique among all API servers, and it must be a valid resource name. | 
**Operations** | Pointer to **map[string]interface{}** | A map from the operation name to the operation object. Each key must be valid ObjectName, and each value must be a valid operation. Each key must be unique within the operations object, which means the operation key is unique within an API server. No duplicate operation values are allowed; operations with the same paths and methods members are not allowed. The operations object is limited to 25 keys (25 individual operations). | [optional] 

## Methods

### NewAPIServer

`func NewAPIServer(authorizationServer APIServerAuthorizationServer, baseURLs []string, name string, ) *APIServer`

NewAPIServer instantiates a new APIServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerWithDefaults

`func NewAPIServerWithDefaults() *APIServer`

NewAPIServerWithDefaults instantiates a new APIServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *APIServer) GetLinks() LinksHATEOAS`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *APIServer) GetLinksOk() (*LinksHATEOAS, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *APIServer) SetLinks(v LinksHATEOAS)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *APIServer) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetAuthorizationServer

`func (o *APIServer) GetAuthorizationServer() APIServerAuthorizationServer`

GetAuthorizationServer returns the AuthorizationServer field if non-nil, zero value otherwise.

### GetAuthorizationServerOk

`func (o *APIServer) GetAuthorizationServerOk() (*APIServerAuthorizationServer, bool)`

GetAuthorizationServerOk returns a tuple with the AuthorizationServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationServer

`func (o *APIServer) SetAuthorizationServer(v APIServerAuthorizationServer)`

SetAuthorizationServer sets AuthorizationServer field to given value.


### GetBaseURLs

`func (o *APIServer) GetBaseURLs() []string`

GetBaseURLs returns the BaseURLs field if non-nil, zero value otherwise.

### GetBaseURLsOk

`func (o *APIServer) GetBaseURLsOk() (*[]string, bool)`

GetBaseURLsOk returns a tuple with the BaseURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseURLs

`func (o *APIServer) SetBaseURLs(v []string)`

SetBaseURLs sets BaseURLs field to given value.


### GetId

`func (o *APIServer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *APIServer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *APIServer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *APIServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *APIServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *APIServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *APIServer) SetName(v string)`

SetName sets Name field to given value.


### GetOperations

`func (o *APIServer) GetOperations() map[string]interface{}`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *APIServer) GetOperationsOk() (*map[string]interface{}, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *APIServer) SetOperations(v map[string]interface{})`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *APIServer) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


