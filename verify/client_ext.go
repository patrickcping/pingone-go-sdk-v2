package verify

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"net/http"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"time"
)

type SDKInterfaceFunc func() (any, *http.Response, error)

var (
	maxRetries               = 10
	maximumRetryAfterBackoff = 30
)

func processResponse(f SDKInterfaceFunc, targetObject any) (*http.Response, error) {

	obj, response, error := exponentialBackOffRetry(f)

	if targetObject != nil {
		v := reflect.ValueOf(targetObject)
		if v.Kind() != reflect.Ptr {
			return nil, fmt.Errorf("Target object must be a pointer.  This is always a problem with the SDK, please raise an issue with the SDK maintainers.")
		}
		if !v.Elem().IsValid() {
			return nil, fmt.Errorf("Target object is not valid.  This is always a problem with the SDK, please raise an issue with the SDK maintainers.")
		}

		if obj != nil {
			v.Elem().Set(reflect.ValueOf(obj))
		}
	}

	return response, reformError(error)
}

func reformError(err error) error {
	return err
}

func exponentialBackOffRetry(f SDKInterfaceFunc) (interface{}, *http.Response, error) {
	var obj interface{}
	var resp *http.Response
	var err error
	backOffTime := time.Second
	var isRetryable bool

	for i := 0; i < maxRetries; i++ {
		obj, resp, err = f()
		retryAttempt := i + 1

		backOffTime, isRetryable = testForRetryable(resp, err, retryAttempt)

		if isRetryable {
			log.Printf("Attempt %d failed: %v, backing off by %s.", retryAttempt, err, backOffTime.String())
			time.Sleep(backOffTime)
			continue
		}

		return obj, resp, err
	}

	log.Printf("Request failed after %d attempts", maxRetries)

	return obj, resp, err // output the final error
}

func testForRetryable(r *http.Response, err error, retryAttempt int) (time.Duration, bool) {

	baseDelay := time.Second
	requestDelayDuration := calculateExponentialBackoff(retryAttempt, baseDelay)

	if r != nil {
		if r.StatusCode == 501 || r.StatusCode == 503 || r.StatusCode == 429 {
			retryAfter, err := parseRetryAfterHeader(r)
			if err != nil {
				log.Printf("Cannot parse the expected \"Retry-After\" header: %s", err)
			}

			if err != nil {
				if retryAfter <= time.Duration(maximumRetryAfterBackoff) {
					requestDelayDuration += time.Duration(maximumRetryAfterBackoff)
				} else {
					requestDelayDuration += retryAfter
				}
			}
		}

		retryAbleCodes := []int{
			408,
			429,
			500,
			501,
			502,
			503,
			504,
		}

		if slices.Contains(retryAbleCodes, r.StatusCode) {
			log.Printf("HTTP status code %d detected, available for retry", r.StatusCode)
			return requestDelayDuration, true
		}
	}

	if err != nil {
		if genericOAError, ok := err.(*GenericOpenAPIError); ok && genericOAError.Model() != nil {
			// We have an application level error

			if modelError, ok := genericOAError.Model().(P1Error); ok {

				// Test for unexpected errors
				if strings.EqualFold(modelError.GetCode(), "UNEXPECTED_ERROR") {
					log.Printf("Unexpected error detected, available for retry")
					return requestDelayDuration, true
				}

				// Test for inconsistent role state
				m, _ := regexp.MatchString(`^Role assignment [a-z0-9\-]* cannot be deleted as it is read only`, modelError.GetMessage())

				if m {
					log.Printf("Inconsistent role assignment, available for retry")
					return requestDelayDuration, true
				}
			}
		}
	}

	return requestDelayDuration, false
}

func parseRetryAfterHeader(resp *http.Response) (time.Duration, error) {
	retryAfterHeader := resp.Header.Get("Retry-After")

	if retryAfterHeader == "" {
		return 0, fmt.Errorf("Retry-After header not found")
	}

	retryAfterSeconds, err := strconv.Atoi(retryAfterHeader)

	if err == nil {
		return time.Duration(retryAfterSeconds) * time.Second, nil
	}

	retryAfterTime, err := http.ParseTime(retryAfterHeader)

	if err != nil {
		return 0, fmt.Errorf("Unable to parse Retry-After header value: %v", err)
	}

	return time.Until(retryAfterTime), nil
}

func calculateExponentialBackoff(attempt int, baseDelay time.Duration) time.Duration {
	jitter := time.Duration(rand.Intn(101)) * time.Millisecond // Add random jitter
	return baseDelay*time.Duration(math.Pow(2, float64(attempt))) + jitter
}
