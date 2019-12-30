package requests

import "net/http"

type Resource struct {
	HttpMethod string
	URL        string
	Query      map[string]string
	Data       []byte
}

func queryChecker(q map[string]string) map[string]string {
	if q == nil {
		return map[string]string{}
	}
	return q
}

func (requests *Requests) Get(endpoint string, query map[string]string) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodGet,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       nil,
	})
}

func (requests *Requests) Post(endpoint string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodPost,
		URL:        endpoint,
		Query:      map[string]string{},
		Data:       data,
	})
}

func (requests *Requests) Put(endpoint string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodPut,
		URL:        endpoint,
		Query:      map[string]string{},
		Data:       data,
	})
}

func (requests *Requests) Patch(endpoint string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodPatch,
		URL:        endpoint,
		Query:      map[string]string{},
		Data:       data,
	})
}

func (requests *Requests) Delete(endpoint string, query map[string]string) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodDelete,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       nil,
	})
}
