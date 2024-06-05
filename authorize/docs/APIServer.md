# APIServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | Pointer to [**LinksHATEOAS**](LinksHATEOAS.md) |  | [optional] 
**AccessControl** | Pointer to [**APIServerAccessControl**](APIServerAccessControl.md) |  | [optional] 
**AuthorizationServer** | [**APIServerAuthorizationServer**](APIServerAuthorizationServer.md) |  | 
**BaseUrls** | **[]string** | An array of string that specifies the possible base URLs that an end-user will use to access the APIs hosted on the customer&#39;s API server. Multiple base URLs may be specified to support cases where the same API may be available from multiple URLs (for example, from a user-friendly domain URL and an internal domain URL). Base URLs must be valid absolute URLs with the https or http scheme. If the path component is non-empty, it must not end in a trailing slash. The path must not contain empty backslash, dot, or double-dot segments. It must not have a query or fragment present, and the host portion of the authority must be a DNS hostname or valid IP (IPv4 or IPv6). The length must be less than or equal to 256 characters. | 
**Directory** | Pointer to [**APIServerDirectory**](APIServerDirectory.md) |  | [optional] 
**Id** | Pointer to **string** | A string that specifies the resource&#39;s unique identifier. | [optional] [readonly] 
**Name** | **string** | A string that specifies the API server resource name. The name value must be unique among all API servers, and it must be a valid resource name. | 
**Policy** | Pointer to [**APIServerPolicy**](APIServerPolicy.md) |  | [optional] 

## Methods

### NewAPIServer

`func NewAPIServer(authorizationServer APIServerAuthorizationServer, baseUrls []string, name string, ) *APIServer`

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

### GetAccessControl

`func (o *APIServer) GetAccessControl() APIServerAccessControl`

GetAccessControl returns the AccessControl field if non-nil, zero value otherwise.

### GetAccessControlOk

`func (o *APIServer) GetAccessControlOk() (*APIServerAccessControl, bool)`

GetAccessControlOk returns a tuple with the AccessControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessControl

`func (o *APIServer) SetAccessControl(v APIServerAccessControl)`

SetAccessControl sets AccessControl field to given value.

### HasAccessControl

`func (o *APIServer) HasAccessControl() bool`

HasAccessControl returns a boolean if a field has been set.

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


### GetBaseUrls

`func (o *APIServer) GetBaseUrls() []string`

GetBaseUrls returns the BaseUrls field if non-nil, zero value otherwise.

### GetBaseUrlsOk

`func (o *APIServer) GetBaseUrlsOk() (*[]string, bool)`

GetBaseUrlsOk returns a tuple with the BaseUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrls

`func (o *APIServer) SetBaseUrls(v []string)`

SetBaseUrls sets BaseUrls field to given value.


### GetDirectory

`func (o *APIServer) GetDirectory() APIServerDirectory`

GetDirectory returns the Directory field if non-nil, zero value otherwise.

### GetDirectoryOk

`func (o *APIServer) GetDirectoryOk() (*APIServerDirectory, bool)`

GetDirectoryOk returns a tuple with the Directory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirectory

`func (o *APIServer) SetDirectory(v APIServerDirectory)`

SetDirectory sets Directory field to given value.

### HasDirectory

`func (o *APIServer) HasDirectory() bool`

HasDirectory returns a boolean if a field has been set.

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


### GetPolicy

`func (o *APIServer) GetPolicy() APIServerPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *APIServer) GetPolicyOk() (*APIServerPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *APIServer) SetPolicy(v APIServerPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *APIServer) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


