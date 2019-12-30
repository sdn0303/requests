package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/cenkalti/backoff"
)

type Requests struct {
	Headers    map[string]string
	HttpClient *http.Client
	RetryLimit uint64
}

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

func (requests *Requests) doRequest(resources Resource) (resp *http.Response, err error) {

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

type ResponseData struct {
	Headers    http.Header
	Body       []byte
	Status     string
	StatusCode int
}

func (requests *Requests) handleRequestWithRetry(resources Resource) (*ResponseData, error) {

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
