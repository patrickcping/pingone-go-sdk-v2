# \GroupMembershipApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToGroup**](GroupMembershipApi.md#AddUserToGroup) | **Post** /v1/environments/{environmentID}/users/{userID}/memberOfGroups | ADD User to Group
[**ReadAllGroupMembershipsForUser**](GroupMembershipApi.md#ReadAllGroupMembershipsForUser) | **Get** /v1/environments/{environmentID}/users/{userID}/memberOfGroups | READ All Group Memberships for User
[**ReadOneGroupMembershipForUser**](GroupMembershipApi.md#ReadOneGroupMembershipForUser) | **Get** /v1/environments/{environmentID}/users/{userID}/memberOfGroups/{groupID} | READ One Group Membership for User
[**RemoveUserFromGroup**](GroupMembershipApi.md#RemoveUserFromGroup) | **Delete** /v1/environments/{environmentID}/users/{userID}/memberOfGroups/{groupID} | REMOVE User from Group



## AddUserToGroup

> Group AddUserToGroup(ctx, environmentID, userID).AddUserToGroupRequest(addUserToGroupRequest).Execute()

ADD User to Group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    addUserToGroupRequest := *openapiclient.NewAddUserToGroupRequest() // AddUserToGroupRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupMembershipApi.AddUserToGroup(context.Background(), environmentID, userID).AddUserToGroupRequest(addUserToGroupRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMembershipApi.AddUserToGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddUserToGroup`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupMembershipApi.AddUserToGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addUserToGroupRequest** | [**AddUserToGroupRequest**](AddUserToGroupRequest.md) |  | 

### Return type

[**Group**](Group.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadAllGroupMembershipsForUser

> EntityArray ReadAllGroupMembershipsForUser(ctx, environmentID, userID).Expand(expand).Limit(limit).Filter(filter).Execute()

READ All Group Memberships for User

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    expand := "group" // string |  (optional)
    limit := int32(100) // int32 |  (optional)
    filter := "type eq "DIRECT"" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupMembershipApi.ReadAllGroupMembershipsForUser(context.Background(), environmentID, userID).Expand(expand).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMembershipApi.ReadAllGroupMembershipsForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllGroupMembershipsForUser`: EntityArray
    fmt.Fprintf(os.Stdout, "Response from `GroupMembershipApi.ReadAllGroupMembershipsForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllGroupMembershipsForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** |  | 
 **limit** | **int32** |  | 
 **filter** | **string** |  | 

### Return type

[**EntityArray**](EntityArray.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneGroupMembershipForUser

> Group ReadOneGroupMembershipForUser(ctx, environmentID, userID, groupID).Expand(expand).Execute()

READ One Group Membership for User

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    groupID := "groupID_example" // string | 
    expand := "group" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupMembershipApi.ReadOneGroupMembershipForUser(context.Background(), environmentID, userID, groupID).Expand(expand).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMembershipApi.ReadOneGroupMembershipForUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneGroupMembershipForUser`: Group
    fmt.Fprintf(os.Stdout, "Response from `GroupMembershipApi.ReadOneGroupMembershipForUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**groupID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneGroupMembershipForUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **expand** | **string** |  | 

### Return type

[**Group**](Group.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveUserFromGroup

> RemoveUserFromGroup(ctx, environmentID, userID, groupID).Execute()

REMOVE User from Group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    groupID := "groupID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GroupMembershipApi.RemoveUserFromGroup(context.Background(), environmentID, userID, groupID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupMembershipApi.RemoveUserFromGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 
**groupID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

