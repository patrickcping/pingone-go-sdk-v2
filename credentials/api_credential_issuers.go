/*
PingOne Platform API - Credentials

The PingOne Platform API covering the PingOne Credentials service

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package credentials

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// CredentialIssuersApiService CredentialIssuersApi service
type CredentialIssuersApiService service

type ApiCreateCredentialIssuerProfileRequest struct {
	ctx                        context.Context
	ApiService                 *CredentialIssuersApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	credentialIssuerProfile    *CredentialIssuerProfile
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiCreateCredentialIssuerProfileRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiCreateCredentialIssuerProfileRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiCreateCredentialIssuerProfileRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiCreateCredentialIssuerProfileRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiCreateCredentialIssuerProfileRequest) CredentialIssuerProfile(credentialIssuerProfile CredentialIssuerProfile) ApiCreateCredentialIssuerProfileRequest {
	r.credentialIssuerProfile = &credentialIssuerProfile
	return r
}

func (r ApiCreateCredentialIssuerProfileRequest) Execute() (*CredentialIssuerProfile, *http.Response, error) {
	return r.ApiService.CreateCredentialIssuerProfileExecute(r)
}

/*
CreateCredentialIssuerProfile Create Credential Issuer Profile

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiCreateCredentialIssuerProfileRequest
*/
func (a *CredentialIssuersApiService) CreateCredentialIssuerProfile(ctx context.Context, environmentID string) ApiCreateCredentialIssuerProfileRequest {
	return ApiCreateCredentialIssuerProfileRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return CredentialIssuerProfile
func (a *CredentialIssuersApiService) CreateCredentialIssuerProfileExecute(r ApiCreateCredentialIssuerProfileRequest) (*CredentialIssuerProfile, *http.Response, error) {
	var (
		err                  error
		response             *http.Response
		localVarReturnValue  *CredentialIssuerProfile
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalCreateCredentialIssuerProfileExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CredentialIssuersApiService) internalCreateCredentialIssuerProfileExecute(r ApiCreateCredentialIssuerProfileRequest) (*CredentialIssuerProfile, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CredentialIssuerProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialIssuersApiService.CreateCredentialIssuerProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/credentialIssuerProfile"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	// body params
	localVarPostBody = r.credentialIssuerProfile
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
	}

	var localVarHTTPResponse *http.Response
	var localVarBody []byte

	for i := range maxRetries {
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if i > 0 {
			slog.Debug("Retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
		}

		localVarHTTPResponse, err = a.client.callAPI(req)
		if err != nil || localVarHTTPResponse == nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}

		localVarBody, err = io.ReadAll(localVarHTTPResponse.Body)
		_ = localVarHTTPResponse.Body.Close()
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}

		if localVarHTTPResponse.StatusCode >= 300 {
			newErr := &GenericOpenAPIError{
				body:  localVarBody,
				error: localVarHTTPResponse.Status,
			}
			if localVarHTTPResponse.StatusCode == 400 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 401 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 403 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 404 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				baseDelay := time.Second
				// check if environment created recently - DOCS-8830
				retryEnvironmentResponse, retryVarHTTPResponse, err := a.client.EnvironmentsApi.ReadOneEnvironment(r.ctx, r.environmentID).Execute()
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				if retryVarHTTPResponse.StatusCode == 200 && retryEnvironmentResponse != nil && retryEnvironmentResponse.CreatedAt != nil {
					// Check if the retryEnvironmentResponse.CreatedAt is within the last 17 seconds
					createdAt, nestedErr := time.Parse(time.RFC3339, *retryEnvironmentResponse.CreatedAt)
					if nestedErr != nil {
						slog.Error("Invalid RFC3339 string", "environment created at", *retryEnvironmentResponse.CreatedAt)
					} else {
						if time.Since(createdAt) < 17*time.Second {
							slog.Debug("The environment was created within the last 17 seconds, retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
							// Retry the request
							delay, nestedErr := calculateExponentialBackoff(i, baseDelay)
							if nestedErr != nil {
								slog.Error("Invalid backoff delay duration", "error", nestedErr, "baseDelay", baseDelay, "retry", false)
							} else {
								time.Sleep(delay)
								continue
							}
						}
					}
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 409 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 429 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 500 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		break
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReadCredentialIssuerProfileRequest struct {
	ctx                        context.Context
	ApiService                 *CredentialIssuersApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiReadCredentialIssuerProfileRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiReadCredentialIssuerProfileRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiReadCredentialIssuerProfileRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiReadCredentialIssuerProfileRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiReadCredentialIssuerProfileRequest) Execute() (*CredentialIssuerProfile, *http.Response, error) {
	return r.ApiService.ReadCredentialIssuerProfileExecute(r)
}

/*
ReadCredentialIssuerProfile Read Credential Issuer Profile

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiReadCredentialIssuerProfileRequest
*/
func (a *CredentialIssuersApiService) ReadCredentialIssuerProfile(ctx context.Context, environmentID string) ApiReadCredentialIssuerProfileRequest {
	return ApiReadCredentialIssuerProfileRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return CredentialIssuerProfile
func (a *CredentialIssuersApiService) ReadCredentialIssuerProfileExecute(r ApiReadCredentialIssuerProfileRequest) (*CredentialIssuerProfile, *http.Response, error) {
	var (
		err                  error
		response             *http.Response
		localVarReturnValue  *CredentialIssuerProfile
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalReadCredentialIssuerProfileExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CredentialIssuersApiService) internalReadCredentialIssuerProfileExecute(r ApiReadCredentialIssuerProfileRequest) (*CredentialIssuerProfile, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CredentialIssuerProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialIssuersApiService.ReadCredentialIssuerProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/credentialIssuerProfile"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
	}

	var localVarHTTPResponse *http.Response
	var localVarBody []byte

	for i := range maxRetries {
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if i > 0 {
			slog.Debug("Retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
		}

		localVarHTTPResponse, err = a.client.callAPI(req)
		if err != nil || localVarHTTPResponse == nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}

		localVarBody, err = io.ReadAll(localVarHTTPResponse.Body)
		_ = localVarHTTPResponse.Body.Close()
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}

		if localVarHTTPResponse.StatusCode >= 300 {
			newErr := &GenericOpenAPIError{
				body:  localVarBody,
				error: localVarHTTPResponse.Status,
			}
			if localVarHTTPResponse.StatusCode == 400 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 401 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 403 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 404 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 409 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 429 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 500 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		break
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateCredentialIssuerProfileRequest struct {
	ctx                        context.Context
	ApiService                 *CredentialIssuersApiService
	environmentID              string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	credentialIssuerProfile    *CredentialIssuerProfile
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiUpdateCredentialIssuerProfileRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiUpdateCredentialIssuerProfileRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiUpdateCredentialIssuerProfileRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiUpdateCredentialIssuerProfileRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiUpdateCredentialIssuerProfileRequest) CredentialIssuerProfile(credentialIssuerProfile CredentialIssuerProfile) ApiUpdateCredentialIssuerProfileRequest {
	r.credentialIssuerProfile = &credentialIssuerProfile
	return r
}

func (r ApiUpdateCredentialIssuerProfileRequest) Execute() (*CredentialIssuerProfile, *http.Response, error) {
	return r.ApiService.UpdateCredentialIssuerProfileExecute(r)
}

/*
UpdateCredentialIssuerProfile Update Credential Issuer Profile

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@return ApiUpdateCredentialIssuerProfileRequest
*/
func (a *CredentialIssuersApiService) UpdateCredentialIssuerProfile(ctx context.Context, environmentID string) ApiUpdateCredentialIssuerProfileRequest {
	return ApiUpdateCredentialIssuerProfileRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
//
//	@return CredentialIssuerProfile
func (a *CredentialIssuersApiService) UpdateCredentialIssuerProfileExecute(r ApiUpdateCredentialIssuerProfileRequest) (*CredentialIssuerProfile, *http.Response, error) {
	var (
		err                  error
		response             *http.Response
		localVarReturnValue  *CredentialIssuerProfile
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalUpdateCredentialIssuerProfileExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *CredentialIssuersApiService) internalUpdateCredentialIssuerProfileExecute(r ApiUpdateCredentialIssuerProfileRequest) (*CredentialIssuerProfile, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CredentialIssuerProfile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CredentialIssuersApiService.UpdateCredentialIssuerProfile")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/credentialIssuerProfile"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPingExternalTransactionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Transaction-ID", r.xPingExternalTransactionID, "")
	}
	if r.xPingExternalSessionID != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-Ping-External-Session-ID", r.xPingExternalSessionID, "")
	}
	// body params
	localVarPostBody = r.credentialIssuerProfile
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	var bodyBytes []byte
	if req.Body != nil {
		bodyBytes, _ = io.ReadAll(req.Body)
	}

	var localVarHTTPResponse *http.Response
	var localVarBody []byte

	for i := range maxRetries {
		if req.Body != nil {
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		if i > 0 {
			slog.Debug("Retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
		}

		localVarHTTPResponse, err = a.client.callAPI(req)
		if err != nil || localVarHTTPResponse == nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}

		localVarBody, err = io.ReadAll(localVarHTTPResponse.Body)
		_ = localVarHTTPResponse.Body.Close()
		localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
		if err != nil {
			return localVarReturnValue, localVarHTTPResponse, err
		}

		if localVarHTTPResponse.StatusCode >= 300 {
			newErr := &GenericOpenAPIError{
				body:  localVarBody,
				error: localVarHTTPResponse.Status,
			}
			if localVarHTTPResponse.StatusCode == 400 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 401 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 403 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 404 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				baseDelay := time.Second
				// check if environment created recently - DOCS-8830
				retryEnvironmentResponse, retryVarHTTPResponse, err := a.client.EnvironmentsApi.ReadOneEnvironment(r.ctx, r.environmentID).Execute()
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				if retryVarHTTPResponse.StatusCode == 200 && retryEnvironmentResponse != nil && retryEnvironmentResponse.CreatedAt != nil {
					// Check if the retryEnvironmentResponse.CreatedAt is within the last 17 seconds
					createdAt, nestedErr := time.Parse(time.RFC3339, *retryEnvironmentResponse.CreatedAt)
					if nestedErr != nil {
						slog.Error("Invalid RFC3339 string", "environment created at", *retryEnvironmentResponse.CreatedAt)
					} else {
						if time.Since(createdAt) < 17*time.Second {
							slog.Debug("The environment was created within the last 17 seconds, retrying request", "attempt", i, "method", localVarHTTPMethod, "path", localVarPath)
							// Retry the request
							delay, nestedErr := calculateExponentialBackoff(i, baseDelay)
							if nestedErr != nil {
								slog.Error("Invalid backoff delay duration", "error", nestedErr, "baseDelay", baseDelay, "retry", false)
							} else {
								time.Sleep(delay)
								continue
							}
						}
					}
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 409 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 429 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			if localVarHTTPResponse.StatusCode == 500 {
				var v P1Error
				err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHTTPResponse, newErr
				}
				newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
				newErr.model = v
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		break
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
