package requests

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Requests struct {
	BaseURL    string
	Headers     map[string]string
	HttpClient *http.Client
}

func New(baseURL string) *Requests {
	return &Requests{
		BaseURL:    baseURL,
		Headers:     map[string]string{},
		HttpClient: &http.Client{},
	}
}

func (requests *Requests) doRequest(method, uri string, query map[string]string, data []byte) (body []byte, err error) {
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

	resp, err := requests.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return body, nil
}

func (requests *Requests) Get(uri string, query map[string]string) ([]byte, error) {
	if query == nil {
		query = map[string]string{}
	}

	b, err := requests.doRequest(http.MethodGet, uri, query, nil)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func (requests *Requests) Post(uri string, data []byte) ([]byte, error) {
	b, err := requests.doRequest(http.MethodPost, uri, map[string]string{}, data)
	if err != nil {
		return nil, err
	}

	return b, nil
}
