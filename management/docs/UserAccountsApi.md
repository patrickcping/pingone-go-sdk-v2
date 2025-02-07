# \UserAccountsApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UserAccount**](UserAccountsApi.md#UserAccount) | **Post** /environments/{environmentID}/users/{userID} | User Account



## UserAccount

> User UserAccount(ctx, environmentID, userID).ContentType(contentType).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).UserAccount(userAccount).Execute()

User Account

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/management"
)

func main() {
    environmentID := "environmentID_example" // string | 
    userID := "userID_example" // string | 
    contentType := openapiclient.EnumUserAccountContentTypeHeader("application/vnd.pingidentity.account.unlock+json") // EnumUserAccountContentTypeHeader |  (optional)
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    userAccount := *openapiclient.NewUserAccount(false, openapiclient.EnumUserStatus("LOCKED")) // UserAccount |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserAccountsApi.UserAccount(context.Background(), environmentID, userID).ContentType(contentType).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).UserAccount(userAccount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAccountsApi.UserAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UserAccount`: User
    fmt.Fprintf(os.Stdout, "Response from `UserAccountsApi.UserAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**userID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUserAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | [**EnumUserAccountContentTypeHeader**](EnumUserAccountContentTypeHeader.md) |  | 
 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **userAccount** | [**UserAccount**](UserAccount.md) |  | 

### Return type

[**User**](User.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

