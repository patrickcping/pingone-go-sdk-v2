# APIServerOperationAccessControlScope

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MatchType** | Pointer to [**EnumAPIServerOperationMatchType**](EnumAPIServerOperationMatchType.md) |  | [optional] 
**Scopes** | [**[]APIServerOperationAccessControlScopeScopesInner**](APIServerOperationAccessControlScopeScopesInner.md) | A list of scopes that define the access requirements for the operation. The client must be authorized with &#x60;ANY&#x60; or &#x60;ALL&#x60; of the scopes to be granted access, depending on the &#x60;matchType&#x60; configuration. | 

## Methods

### NewAPIServerOperationAccessControlScope

`func NewAPIServerOperationAccessControlScope(scopes []APIServerOperationAccessControlScopeScopesInner, ) *APIServerOperationAccessControlScope`

NewAPIServerOperationAccessControlScope instantiates a new APIServerOperationAccessControlScope object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAPIServerOperationAccessControlScopeWithDefaults

`func NewAPIServerOperationAccessControlScopeWithDefaults() *APIServerOperationAccessControlScope`

NewAPIServerOperationAccessControlScopeWithDefaults instantiates a new APIServerOperationAccessControlScope object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMatchType

`func (o *APIServerOperationAccessControlScope) GetMatchType() EnumAPIServerOperationMatchType`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *APIServerOperationAccessControlScope) GetMatchTypeOk() (*EnumAPIServerOperationMatchType, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *APIServerOperationAccessControlScope) SetMatchType(v EnumAPIServerOperationMatchType)`

SetMatchType sets MatchType field to given value.

### HasMatchType

`func (o *APIServerOperationAccessControlScope) HasMatchType() bool`

HasMatchType returns a boolean if a field has been set.

### GetScopes

`func (o *APIServerOperationAccessControlScope) GetScopes() []APIServerOperationAccessControlScopeScopesInner`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *APIServerOperationAccessControlScope) GetScopesOk() (*[]APIServerOperationAccessControlScopeScopesInner, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *APIServerOperationAccessControlScope) SetScopes(v []APIServerOperationAccessControlScopeScopesInner)`

SetScopes sets Scopes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


