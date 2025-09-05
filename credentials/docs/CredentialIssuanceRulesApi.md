# \CredentialIssuanceRulesApi

All URIs are relative to *https://api.pingone.com/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApplyCredentialIssuanceRuleStagedChanges**](CredentialIssuanceRulesApi.md#ApplyCredentialIssuanceRuleStagedChanges) | **Post** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID}/stagedChanges | Apply Credential Issuance Rule Staged Changes
[**CreateCredentialIssuanceRule**](CredentialIssuanceRulesApi.md#CreateCredentialIssuanceRule) | **Post** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules | Create Credential Issuance Rule
[**DeleteCredentialIssuanceRule**](CredentialIssuanceRulesApi.md#DeleteCredentialIssuanceRule) | **Delete** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID} | Delete Credential Issuance Rule
[**ReadAllCredentialIssuanceRules**](CredentialIssuanceRulesApi.md#ReadAllCredentialIssuanceRules) | **Get** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules | Read All Credential Issuance Rules
[**ReadCredentialIssuanceRuleStagedChanges**](CredentialIssuanceRulesApi.md#ReadCredentialIssuanceRuleStagedChanges) | **Get** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID}/stagedChanges | Read Credential Issuance Rule Staged Changes
[**ReadCredentialIssuanceRuleUsageCounts**](CredentialIssuanceRulesApi.md#ReadCredentialIssuanceRuleUsageCounts) | **Get** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID}/usageCounts | Read Credential Issuance Rule Usage Counts
[**ReadCredentialIssuanceRuleUsageDetails**](CredentialIssuanceRulesApi.md#ReadCredentialIssuanceRuleUsageDetails) | **Get** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID}/usageDetails | Read Credential Issuance Rule Usage Details
[**ReadOneCredentialIssuanceRule**](CredentialIssuanceRulesApi.md#ReadOneCredentialIssuanceRule) | **Get** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID} | Read One Credential Issuance Rule
[**UpdateCredentialIssuanceRule**](CredentialIssuanceRulesApi.md#UpdateCredentialIssuanceRule) | **Put** /environments/{environmentID}/credentialTypes/{credentialTypeID}/issuanceRules/{credentialIssuanceRuleID} | Update Credential Issuance Rule



## ApplyCredentialIssuanceRuleStagedChanges

> CredentialIssuanceRuleStagedChange ApplyCredentialIssuanceRuleStagedChanges(ctx, environmentID, credentialTypeID, credentialIssuanceRuleID).ContentType(contentType).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).CredentialIssuanceRuleStagedChange(credentialIssuanceRuleStagedChange).Execute()

Apply Credential Issuance Rule Staged Changes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    contentType := "application/vnd.pingidentity.credentials.applyStagedChanges+json" // string |  (optional)
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    credentialIssuanceRuleStagedChange := *openapiclient.NewCredentialIssuanceRuleStagedChange() // CredentialIssuanceRuleStagedChange |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialIssuanceRulesApi.ApplyCredentialIssuanceRuleStagedChanges(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).ContentType(contentType).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).CredentialIssuanceRuleStagedChange(credentialIssuanceRuleStagedChange).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.ApplyCredentialIssuanceRuleStagedChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplyCredentialIssuanceRuleStagedChanges`: CredentialIssuanceRuleStagedChange
    fmt.Fprintf(os.Stdout, "Response from `CredentialIssuanceRulesApi.ApplyCredentialIssuanceRuleStagedChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApplyCredentialIssuanceRuleStagedChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **contentType** | **string** |  | 
 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **credentialIssuanceRuleStagedChange** | [**CredentialIssuanceRuleStagedChange**](CredentialIssuanceRuleStagedChange.md) |  | 

### Return type

[**CredentialIssuanceRuleStagedChange**](CredentialIssuanceRuleStagedChange.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCredentialIssuanceRule

> CredentialIssuanceRule CreateCredentialIssuanceRule(ctx, environmentID, credentialTypeID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).CredentialIssuanceRule(credentialIssuanceRule).Execute()

Create Credential Issuance Rule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    credentialIssuanceRule := *openapiclient.NewCredentialIssuanceRule(*openapiclient.NewCredentialIssuanceRuleAutomation(openapiclient.EnumCredentialIssuanceRuleAutomationMethod("PERIODIC"), openapiclient.EnumCredentialIssuanceRuleAutomationMethod("PERIODIC"), openapiclient.EnumCredentialIssuanceRuleAutomationMethod("PERIODIC")), openapiclient.EnumCredentialIssuanceRuleStatus("ACTIVE")) // CredentialIssuanceRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialIssuanceRulesApi.CreateCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).CredentialIssuanceRule(credentialIssuanceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.CreateCredentialIssuanceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCredentialIssuanceRule`: CredentialIssuanceRule
    fmt.Fprintf(os.Stdout, "Response from `CredentialIssuanceRulesApi.CreateCredentialIssuanceRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCredentialIssuanceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **credentialIssuanceRule** | [**CredentialIssuanceRule**](CredentialIssuanceRule.md) |  | 

### Return type

[**CredentialIssuanceRule**](CredentialIssuanceRule.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredentialIssuanceRule

> DeleteCredentialIssuanceRule(ctx, environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Delete Credential Issuance Rule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.CredentialIssuanceRulesApi.DeleteCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.DeleteCredentialIssuanceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialIssuanceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

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


## ReadAllCredentialIssuanceRules

> EntityArrayPagedIterator ReadAllCredentialIssuanceRules(ctx, environmentID, credentialTypeID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Read All Credential Issuance Rules

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialIssuanceRulesApi.ReadAllCredentialIssuanceRules(context.Background(), environmentID, credentialTypeID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.ReadAllCredentialIssuanceRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadAllCredentialIssuanceRules`: EntityArrayPagedIterator
    fmt.Fprintf(os.Stdout, "Response from `CredentialIssuanceRulesApi.ReadAllCredentialIssuanceRules`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadAllCredentialIssuanceRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

### Return type

[**EntityArrayPagedIterator**](EntityArrayPagedIterator.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCredentialIssuanceRuleStagedChanges

> CredentialIssuanceRuleStagedChange ReadCredentialIssuanceRuleStagedChanges(ctx, environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Read Credential Issuance Rule Staged Changes

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleStagedChanges(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleStagedChanges``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCredentialIssuanceRuleStagedChanges`: CredentialIssuanceRuleStagedChange
    fmt.Fprintf(os.Stdout, "Response from `CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleStagedChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCredentialIssuanceRuleStagedChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

### Return type

[**CredentialIssuanceRuleStagedChange**](CredentialIssuanceRuleStagedChange.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCredentialIssuanceRuleUsageCounts

> CredentialIssuanceRuleUsageCounts ReadCredentialIssuanceRuleUsageCounts(ctx, environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Read Credential Issuance Rule Usage Counts

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageCounts(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageCounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCredentialIssuanceRuleUsageCounts`: CredentialIssuanceRuleUsageCounts
    fmt.Fprintf(os.Stdout, "Response from `CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageCounts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCredentialIssuanceRuleUsageCountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

### Return type

[**CredentialIssuanceRuleUsageCounts**](CredentialIssuanceRuleUsageCounts.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadCredentialIssuanceRuleUsageDetails

> CredentialIssuanceRuleUsageDetails ReadCredentialIssuanceRuleUsageDetails(ctx, environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Read Credential Issuance Rule Usage Details

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageDetails(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageDetails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadCredentialIssuanceRuleUsageDetails`: CredentialIssuanceRuleUsageDetails
    fmt.Fprintf(os.Stdout, "Response from `CredentialIssuanceRulesApi.ReadCredentialIssuanceRuleUsageDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadCredentialIssuanceRuleUsageDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

### Return type

[**CredentialIssuanceRuleUsageDetails**](CredentialIssuanceRuleUsageDetails.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadOneCredentialIssuanceRule

> CredentialIssuanceRule ReadOneCredentialIssuanceRule(ctx, environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()

Read One Credential Issuance Rule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialIssuanceRulesApi.ReadOneCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.ReadOneCredentialIssuanceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReadOneCredentialIssuanceRule`: CredentialIssuanceRule
    fmt.Fprintf(os.Stdout, "Response from `CredentialIssuanceRulesApi.ReadOneCredentialIssuanceRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReadOneCredentialIssuanceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 

### Return type

[**CredentialIssuanceRule**](CredentialIssuanceRule.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCredentialIssuanceRule

> CredentialIssuanceRule UpdateCredentialIssuanceRule(ctx, environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).CredentialIssuanceRule(credentialIssuanceRule).Execute()

Update Credential Issuance Rule

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/patrickcping/pingone-go-sdk-v2/credentials"
)

func main() {
    environmentID := "environmentID_example" // string | 
    credentialTypeID := "credentialTypeID_example" // string | 
    credentialIssuanceRuleID := "credentialIssuanceRuleID_example" // string | 
    xPingExternalTransactionID := "xPingExternalTransactionID_example" // string | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    xPingExternalSessionID := "xPingExternalSessionID_example" // string | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn't provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes `/`, `@`, `=`, `#`, `+`  (optional)
    credentialIssuanceRule := *openapiclient.NewCredentialIssuanceRule(*openapiclient.NewCredentialIssuanceRuleAutomation(openapiclient.EnumCredentialIssuanceRuleAutomationMethod("PERIODIC"), openapiclient.EnumCredentialIssuanceRuleAutomationMethod("PERIODIC"), openapiclient.EnumCredentialIssuanceRuleAutomationMethod("PERIODIC")), openapiclient.EnumCredentialIssuanceRuleStatus("ACTIVE")) // CredentialIssuanceRule |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CredentialIssuanceRulesApi.UpdateCredentialIssuanceRule(context.Background(), environmentID, credentialTypeID, credentialIssuanceRuleID).XPingExternalTransactionID(xPingExternalTransactionID).XPingExternalSessionID(xPingExternalSessionID).CredentialIssuanceRule(credentialIssuanceRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CredentialIssuanceRulesApi.UpdateCredentialIssuanceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCredentialIssuanceRule`: CredentialIssuanceRule
    fmt.Fprintf(os.Stdout, "Response from `CredentialIssuanceRulesApi.UpdateCredentialIssuanceRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**credentialTypeID** | **string** |  | 
**credentialIssuanceRuleID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCredentialIssuanceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xPingExternalTransactionID** | **string** | An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **xPingExternalSessionID** | **string** | An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;  | 
 **credentialIssuanceRule** | [**CredentialIssuanceRule**](CredentialIssuanceRule.md) |  | 

### Return type

[**CredentialIssuanceRule**](CredentialIssuanceRule.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

