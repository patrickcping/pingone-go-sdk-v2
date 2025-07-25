/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2023-06-29
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"bytes"
	"context"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

// TranslationsApiService TranslationsApi service
type TranslationsApiService service

type ApiReadTranslationsRequest struct {
	ctx                        context.Context
	ApiService                 *TranslationsApiService
	environmentID              string
	locale                     string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiReadTranslationsRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiReadTranslationsRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiReadTranslationsRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiReadTranslationsRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiReadTranslationsRequest) Execute() EntityArrayPagedIterator {
	return r.ApiService.ReadTranslationsExecute(r)
}

func (r ApiReadTranslationsRequest) ExecuteInitialPage() (*EntityArray, *http.Response, error) {
	return r.ApiService.ReadTranslationsExecuteInitialPage(r)
}

/*
ReadTranslations Read Translations by locale

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param locale
	@return ApiReadTranslationsRequest
*/
func (a *TranslationsApiService) ReadTranslations(ctx context.Context, environmentID string, locale string) ApiReadTranslationsRequest {
	return ApiReadTranslationsRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		locale:        locale,
	}
}

// Execute executes the request
//
//	@return EntityArray
func (a *TranslationsApiService) ReadTranslationsExecute(r ApiReadTranslationsRequest) EntityArrayPagedIterator {
  return a.client.paginationIterator(r.ctx, r.ExecuteInitialPage)
}

func (a *TranslationsApiService) ReadTranslationsExecuteInitialPage(r ApiReadTranslationsRequest) (*EntityArray, *http.Response, error) {
	var (
		err                  error
		response             *http.Response
		localVarReturnValue  *EntityArray
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			return r.ApiService.internalReadTranslationsExecute(r)
		},
		&localVarReturnValue,
	)
	return localVarReturnValue, response, err
}

func (a *TranslationsApiService) internalReadTranslationsExecute(r ApiReadTranslationsRequest) (*EntityArray, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *EntityArray
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TranslationsApiService.ReadTranslations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/translations/{locale}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"locale"+"}", url.PathEscape(parameterValueToString(r.locale, "locale")), -1)

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

type ApiUpdateTranslationsRequest struct {
	ctx                        context.Context
	ApiService                 *TranslationsApiService
	environmentID              string
	locale                     string
	xPingExternalTransactionID *string
	xPingExternalSessionID     *string
	localeTranslation          *[]LocaleTranslation
}

// An ID for telemetry purposes to correlate transactions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiUpdateTranslationsRequest) XPingExternalTransactionID(xPingExternalTransactionID string) ApiUpdateTranslationsRequest {
	r.xPingExternalTransactionID = &xPingExternalTransactionID
	return r
}

// An ID for telemetry purposes to correlate sessions with client systems through PingOne products. This may be a user defined value. If a value isn&#39;t provided on the API request, a unique value will be generated in the API response. See [External transaction and session IDs](https://apidocs.pingidentity.com/pingone/platform/v1/api/#external-transaction-and-session-ids) for more information. Any invalid characters will be converted to underscores. The following characters are allowed: Unicode letters, combining marks, numeric characters, dots, underscores, dashes &#x60;/&#x60;, &#x60;@&#x60;, &#x60;&#x3D;&#x60;, &#x60;#&#x60;, &#x60;+&#x60;
func (r ApiUpdateTranslationsRequest) XPingExternalSessionID(xPingExternalSessionID string) ApiUpdateTranslationsRequest {
	r.xPingExternalSessionID = &xPingExternalSessionID
	return r
}

func (r ApiUpdateTranslationsRequest) LocaleTranslation(localeTranslation []LocaleTranslation) ApiUpdateTranslationsRequest {
	r.localeTranslation = &localeTranslation
	return r
}

func (r ApiUpdateTranslationsRequest) Execute() ([]LocaleTranslation, *http.Response, error) {
	return r.ApiService.UpdateTranslationsExecute(r)
}

/*
UpdateTranslations Update Translations by locale

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param environmentID
	@param locale
	@return ApiUpdateTranslationsRequest
*/
func (a *TranslationsApiService) UpdateTranslations(ctx context.Context, environmentID string, locale string) ApiUpdateTranslationsRequest {
	return ApiUpdateTranslationsRequest{
		ApiService:    a,
		ctx:           ctx,
		environmentID: environmentID,
		locale:        locale,
	}
}

// Execute executes the request
//
//	@return []LocaleTranslation
func (a *TranslationsApiService) UpdateTranslationsExecute(r ApiUpdateTranslationsRequest) ([]LocaleTranslation, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPut
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []LocaleTranslation
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TranslationsApiService.UpdateTranslations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/translations/{locale}"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"locale"+"}", url.PathEscape(parameterValueToString(r.locale, "locale")), -1)

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
	localVarPostBody = r.localeTranslation
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
