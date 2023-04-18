package risk


import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/exp/slices"
)

type SDKInterfaceFunc func() (interface{}, *http.Response, error)

var (
	maxRetries               = 5
	maximumRetryAfterBackoff = 30
)

func processResponse(f SDKInterfaceFunc) (interface{}, *http.Response, error) {
	obj, response, error := exponentialBackOffRetry(f)
	return obj, response, reformError(error)
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

		backOffTime, isRetryable = testForRetryable(resp, err, backOffTime)

		if (err != nil && isRetryable) || isRetryable {
			log.Printf("Attempt %d failed: %v, backing off by %s.", i+1, err, backOffTime.String())
			time.Sleep(backOffTime)
			continue
		}

		return obj, resp, err
	}

	log.Printf("Request failed after %d attempts", maxRetries)

	return obj, resp, err // output the final error
}

func testForRetryable(r *http.Response, err error, currentBackoff time.Duration) (time.Duration, bool) {

	backoff := currentBackoff

	if r.StatusCode == 501 || r.StatusCode == 503 || r.StatusCode == 429 {
		retryAfter, err := parseRetryAfterHeader(r)
		if err != nil {
			log.Printf("Cannot parse the expected \"Retry-After\" header: %s", err)
			backoff = currentBackoff * 2
		}

		if retryAfter <= time.Duration(maximumRetryAfterBackoff) {
			backoff += time.Duration(maximumRetryAfterBackoff)
		} else {
			backoff += retryAfter
		}
	} else {
		backoff = currentBackoff * 2
	}

	retryAbleCodes := []int{
		429,
		500,
		501,
		502,
		503,
		504,
	}

	if slices.Contains(retryAbleCodes, r.StatusCode) {
		log.Printf("HTTP status code %d detected, available for retry", r.StatusCode)
		return backoff, true
	}

	return backoff, false
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
