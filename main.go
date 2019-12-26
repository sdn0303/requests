package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/cenkalti/backoff"
)

type Requests struct {
	BaseURL    string
	Headers    map[string]string
	HttpClient *http.Client
}

func New(baseURL string) *Requests {
	return &Requests{
		BaseURL:    baseURL,
		Headers:    map[string]string{},
		HttpClient: &http.Client{},
	}
}

func (requests *Requests) doRequest(method, uri string, query map[string]string, data []byte) (resp *http.Response, err error) {

	baseURL, err := url.Parse(requests.BaseURL)
	if err != nil {
		return nil, err
	}

	requestURI, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	endpoint := baseURL.ResolveReference(requestURI).String()
	req, err := http.NewRequest(method, endpoint, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	for k, v := range requests.Headers {
		req.Header.Add(k, v)
	}

	return requests.HttpClient.Do(req)
}

type ResponseData struct {
	Headers    http.Header
	Body       []byte
	Status     string
	StatusCode int
}

func (requests *Requests) Get(uri string, query map[string]string, maxRetry uint64) (*ResponseData, error) {

	var (
		b    []byte
		err  error
		resp *http.Response
	)

	if query == nil {
		query = map[string]string{}
	}

	backOff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetry)
	operation := func() error {
		resp, err = requests.doRequest(http.MethodGet, uri, query, nil)
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

func (requests *Requests) Post(uri string, data []byte, maxRetry uint64) (*ResponseData, error) {

	var (
		b    []byte
		err  error
		resp *http.Response
	)

	backOff := backoff.WithMaxRetries(backoff.NewExponentialBackOff(), maxRetry)
	operation := func() error {
		resp, err = requests.doRequest(http.MethodPost, uri, map[string]string{}, data)
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

func (requests *Requests) Put(uri string, data []byte) (*ResponseData, error) {

	resp, err := requests.doRequest(http.MethodPut, uri, map[string]string{}, data)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
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

func (requests *Requests) Delete(uri string, data []byte) (*ResponseData, error) {

	resp, err := requests.doRequest(http.MethodPut, uri, map[string]string{}, data)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(resp.Body)
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
