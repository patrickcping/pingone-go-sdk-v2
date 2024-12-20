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
	"net/http"
	"net/url"
	"strings"
)


// AuthenticationsPerApplicationApiService AuthenticationsPerApplicationApi service
type AuthenticationsPerApplicationApiService service

type ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest struct {
	ctx context.Context
	ApiService *AuthenticationsPerApplicationApiService
	environmentID string
	limit *int32
	samplePeriod *int32
	samplePeriodCount *int32
	filter *string
}

// Adding a paging value to limit the number of resources displayed per page
func (r ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest) Limit(limit int32) ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest) SamplePeriod(samplePeriod int32) ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	r.samplePeriod = &samplePeriod
	return r
}

func (r ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest) SamplePeriodCount(samplePeriodCount int32) ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	r.samplePeriodCount = &samplePeriodCount
	return r
}

func (r ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest) Filter(filter string) ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.EnvironmentsEnvironmentIDApplicationSignonsGetExecute(r)
}

/*
EnvironmentsEnvironmentIDApplicationSignonsGet READ Authentications Per Application (Partial)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @return ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest
*/
func (a *AuthenticationsPerApplicationApiService) EnvironmentsEnvironmentIDApplicationSignonsGet(ctx context.Context, environmentID string) ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	return ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *AuthenticationsPerApplicationApiService) EnvironmentsEnvironmentIDApplicationSignonsGetExecute(r ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest) (*http.Response, error) {
	var (
		err      error
		response *http.Response
	)
	
	response, err = processResponse(
		func() (any, *http.Response, error) {
			resp, err := r.ApiService.internalEnvironmentsEnvironmentIDApplicationSignonsGetExecute(r)
			return nil, resp, err
		},
		nil,
	)
	return response, err
}

func (a *AuthenticationsPerApplicationApiService) internalEnvironmentsEnvironmentIDApplicationSignonsGetExecute(r ApiEnvironmentsEnvironmentIDApplicationSignonsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationsPerApplicationApiService.EnvironmentsEnvironmentIDApplicationSignonsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/environments/{environmentID}/applicationSignons"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterValueToString(r.environmentID, "environmentID")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "")
	}
	if r.samplePeriod != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "samplePeriod", r.samplePeriod, "")
	}
	if r.samplePeriodCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "samplePeriodCount", r.samplePeriodCount, "")
	}
	if r.filter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "filter", r.filter, "")
	}
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
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	_ = localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
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
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
