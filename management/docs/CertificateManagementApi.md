# \CertificateManagementApi

All URIs are relative to *https://api.pingone.eu*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1EnvironmentsEnvironmentIDCertificatesCertIDApplicationsGet**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDCertificatesCertIDApplicationsGet) | **Get** /v1/environments/{environmentID}/certificates/{certID}/applications | GET Certificate Applications
[**V1EnvironmentsEnvironmentIDCertificatesCertIDDelete**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDCertificatesCertIDDelete) | **Delete** /v1/environments/{environmentID}/certificates/{certID} | DELETE Certificate
[**V1EnvironmentsEnvironmentIDCertificatesCertIDGet**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDCertificatesCertIDGet) | **Get** /v1/environments/{environmentID}/certificates/{certID} | GET Certificate
[**V1EnvironmentsEnvironmentIDCertificatesGet**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDCertificatesGet) | **Get** /v1/environments/{environmentID}/certificates | GET Certificates
[**V1EnvironmentsEnvironmentIDCertificatesPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDCertificatesPost) | **Post** /v1/environments/{environmentID}/certificates | CREATE Certificate with PKCS7 or PEM File
[**V1EnvironmentsEnvironmentIDDecryptionsPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDDecryptionsPost) | **Post** /v1/environments/{environmentID}/decryptions | DECRYPT Data
[**V1EnvironmentsEnvironmentIDEncryptionsPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDEncryptionsPost) | **Post** /v1/environments/{environmentID}/encryptions | ENCRYPT Data
[**V1EnvironmentsEnvironmentIDKeysGet**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDKeysGet) | **Get** /v1/environments/{environmentID}/keys | GET Keys
[**V1EnvironmentsEnvironmentIDKeysKeyIDApplicationsGet**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDKeysKeyIDApplicationsGet) | **Get** /v1/environments/{environmentID}/keys/{keyID}/applications | GET Key Applications
[**V1EnvironmentsEnvironmentIDKeysKeyIDCsrGet**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDKeysKeyIDCsrGet) | **Get** /v1/environments/{environmentID}/keys/{keyID}/csr | Export a certificate signing request (CSR)
[**V1EnvironmentsEnvironmentIDKeysKeyIDCsrPut**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDKeysKeyIDCsrPut) | **Put** /v1/environments/{environmentID}/keys/{keyID}/csr | Import Certificate Authority (CA) Response to a CSR
[**V1EnvironmentsEnvironmentIDKeysKeyIDDelete**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDKeysKeyIDDelete) | **Delete** /v1/environments/{environmentID}/keys/{keyID} | DELETE Key
[**V1EnvironmentsEnvironmentIDKeysKeyIDGet**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDKeysKeyIDGet) | **Get** /v1/environments/{environmentID}/keys/{keyID} | EXPORT Public Key (X509 PEM)
[**V1EnvironmentsEnvironmentIDKeysKeyIDPut**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDKeysKeyIDPut) | **Put** /v1/environments/{environmentID}/keys/{keyID} | UPDATE Key
[**V1EnvironmentsEnvironmentIDKeysPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDKeysPost) | **Post** /v1/environments/{environmentID}/keys | CREATE Key with PKCS12 File
[**V1EnvironmentsEnvironmentIDSigningsPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDSigningsPost) | **Post** /v1/environments/{environmentID}/signings | SIGN Data
[**V1EnvironmentsEnvironmentIDVerificationsPost**](CertificateManagementApi.md#V1EnvironmentsEnvironmentIDVerificationsPost) | **Post** /v1/environments/{environmentID}/verifications | VERIFY Signed Data



## V1EnvironmentsEnvironmentIDCertificatesCertIDApplicationsGet

> V1EnvironmentsEnvironmentIDCertificatesCertIDApplicationsGet(ctx, environmentID, certID).Execute()

GET Certificate Applications

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
    certID := "certID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesCertIDApplicationsGet(context.Background(), environmentID, certID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesCertIDApplicationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCertificatesCertIDApplicationsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDCertificatesCertIDDelete

> V1EnvironmentsEnvironmentIDCertificatesCertIDDelete(ctx, environmentID, certID).Execute()

DELETE Certificate

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
    certID := "certID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesCertIDDelete(context.Background(), environmentID, certID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesCertIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCertificatesCertIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDCertificatesCertIDGet

> V1EnvironmentsEnvironmentIDCertificatesCertIDGet(ctx, environmentID, certID).Execute()

GET Certificate

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
    certID := "certID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesCertIDGet(context.Background(), environmentID, certID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesCertIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**certID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCertificatesCertIDGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDCertificatesGet

> V1EnvironmentsEnvironmentIDCertificatesGet(ctx, environmentID).Execute()

GET Certificates

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCertificatesGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDCertificatesPost

> V1EnvironmentsEnvironmentIDCertificatesPost(ctx, environmentID).ContentType(contentType).UsageType(usageType).File(file).Execute()

CREATE Certificate with PKCS7 or PEM File

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
    contentType := "application/x-pkcs7-certificates" // string |  (optional)
    usageType := "usageType_example" // string |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesPost(context.Background(), environmentID).ContentType(contentType).UsageType(usageType).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDCertificatesPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDCertificatesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **usageType** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDDecryptionsPost

> V1EnvironmentsEnvironmentIDDecryptionsPost(ctx, environmentID).Body(body).Execute()

DECRYPT Data

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDDecryptionsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDDecryptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDDecryptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDEncryptionsPost

> V1EnvironmentsEnvironmentIDEncryptionsPost(ctx, environmentID).Body(body).Execute()

ENCRYPT Data

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDEncryptionsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDEncryptionsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDEncryptionsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDKeysGet

> V1EnvironmentsEnvironmentIDKeysGet(ctx, environmentID).Execute()

GET Keys

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysGet(context.Background(), environmentID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDKeysGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDKeysKeyIDApplicationsGet

> V1EnvironmentsEnvironmentIDKeysKeyIDApplicationsGet(ctx, environmentID, keyID).Execute()

GET Key Applications

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
    keyID := "keyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDApplicationsGet(context.Background(), environmentID, keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDApplicationsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDKeysKeyIDApplicationsGetRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDKeysKeyIDCsrGet

> V1EnvironmentsEnvironmentIDKeysKeyIDCsrGet(ctx, environmentID, keyID).ContentType(contentType).Execute()

Export a certificate signing request (CSR)

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
    keyID := "keyID_example" // string | 
    contentType := "application/pkcs10" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDCsrGet(context.Background(), environmentID, keyID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDCsrGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDKeysKeyIDCsrGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDKeysKeyIDCsrPut

> V1EnvironmentsEnvironmentIDKeysKeyIDCsrPut(ctx, environmentID, keyID).ContentType(contentType).Execute()

Import Certificate Authority (CA) Response to a CSR

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
    keyID := "keyID_example" // string | 
    contentType := "application/x-pem-file" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDCsrPut(context.Background(), environmentID, keyID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDCsrPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDKeysKeyIDCsrPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDKeysKeyIDDelete

> V1EnvironmentsEnvironmentIDKeysKeyIDDelete(ctx, environmentID, keyID).Execute()

DELETE Key

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
    keyID := "keyID_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDDelete(context.Background(), environmentID, keyID).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDKeysKeyIDDeleteRequest struct via the builder pattern


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


## V1EnvironmentsEnvironmentIDKeysKeyIDGet

> V1EnvironmentsEnvironmentIDKeysKeyIDGet(ctx, environmentID, keyID).ContentType(contentType).Execute()

EXPORT Public Key (X509 PEM)

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
    keyID := "keyID_example" // string | 
    contentType := "application/x-pkcs7-certificates" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDGet(context.Background(), environmentID, keyID).ContentType(contentType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDKeysKeyIDGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contentType** | **string** |  | 

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


## V1EnvironmentsEnvironmentIDKeysKeyIDPut

> V1EnvironmentsEnvironmentIDKeysKeyIDPut(ctx, environmentID, keyID).Body(body).Execute()

UPDATE Key

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
    keyID := "keyID_example" // string | 
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDPut(context.Background(), environmentID, keyID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysKeyIDPut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 
**keyID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDKeysKeyIDPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDKeysPost

> V1EnvironmentsEnvironmentIDKeysPost(ctx, environmentID).ContentType(contentType).UsageType(usageType).File(file).Execute()

CREATE Key with PKCS12 File

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
    contentType := "application/x-pkcs12-certificates" // string |  (optional)
    usageType := "usageType_example" // string |  (optional)
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysPost(context.Background(), environmentID).ContentType(contentType).UsageType(usageType).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDKeysPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDKeysPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contentType** | **string** |  | 
 **usageType** | **string** |  | 
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDSigningsPost

> V1EnvironmentsEnvironmentIDSigningsPost(ctx, environmentID).Body(body).Execute()

SIGN Data

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDSigningsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDSigningsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDSigningsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## V1EnvironmentsEnvironmentIDVerificationsPost

> V1EnvironmentsEnvironmentIDVerificationsPost(ctx, environmentID).Body(body).Execute()

VERIFY Signed Data

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
    body := map[string]interface{}{ ... } // map[string]interface{} |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CertificateManagementApi.V1EnvironmentsEnvironmentIDVerificationsPost(context.Background(), environmentID).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CertificateManagementApi.V1EnvironmentsEnvironmentIDVerificationsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**environmentID** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiV1EnvironmentsEnvironmentIDVerificationsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 

### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

