// Copyright (c) 2020 Soichi David Nakahashi
// Released under the MIT license.
// see https://github.com/sdn0303/requests/blob/master/LICENSE
package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

// Requests struct of http client
type Requests struct {
	Headers    map[string]string
	HttpClient *http.Client
	RetryLimit uint64
}

// ResponseData struct of response data
type ResponseData struct {
	Headers    http.Header
	Body       []byte
	Status     string
	StatusCode int
}

// resource holds the resources needed to send a request
type resource struct {
	HttpMethod string
	URL        string
	Query      map[string]string
	Data       []byte
}

// New initialize http client and set options
func New(options ...Option) *Requests {

	requests := &Requests{
		Headers: map[string]string{},
		HttpClient: &http.Client{
			Timeout: time.Duration(30) * time.Second,
		},
		RetryLimit: 0,
	}

	for _, option := range options {
		option(requests)
	}

	return requests
}

// doRequest sends a request
func (requests *Requests) doRequest(resources *resource) (*http.Response, error) {

	req, err := http.NewRequest(resources.HttpMethod, resources.URL, bytes.NewBuffer(resources.Data))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range resources.Query {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()
	for k, v := range requests.Headers {
		req.Header.Add(k, v)
	}

	return requests.HttpClient.Do(req)
}

// handleRequestWithRetry wraps doRequest function so that retry processing can be performed
func (requests *Requests) handleRequestWithRetry(resources *resource) (*ResponseData, error) {

	var (
		b    []byte
		err  error
		resp *http.Response
	)

	backOff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), requests.RetryLimit)
	operation := func() error {
		resp, err = requests.doRequest(resources)
		return err
	}
	if err := backoff.Retry(operation, backOff); err != nil {
		return nil, err
	}

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return &ResponseData{
		Headers:    resp.Header,
		Body:       b,
		Status:     resp.Status,
		StatusCode: resp.StatusCode,
	}, nil
}

// queryChecker returns map[string]string{} if query is nil
func queryChecker(q map[string]string) map[string]string {
	if q == nil {
		return map[string]string{}
	}
	return q
}

// Get
func (requests *Requests) Get(endpoint string, query map[string]string) (*ResponseData, error) {
	return requests.handleRequestWithRetry(&resource{
		HttpMethod: http.MethodGet,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       nil,
	})
}

// Post
func (requests *Requests) Post(endpoint string, query map[string]string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(&resource{
		HttpMethod: http.MethodPost,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       data,
	})
}

// Put
func (requests *Requests) Put(endpoint string, query map[string]string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(&resource{
		HttpMethod: http.MethodPut,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       data,
	})
}

// Patch
func (requests *Requests) Patch(endpoint string, query map[string]string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(&resource{
		HttpMethod: http.MethodPatch,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       data,
	})
}

// Delete
func (requests *Requests) Delete(endpoint string, query map[string]string) (*ResponseData, error) {
	return requests.handleRequestWithRetry(&resource{
		HttpMethod: http.MethodDelete,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       nil,
	})
}
