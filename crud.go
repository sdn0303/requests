package requests

import "net/http"

// Resource holds the resources needed to send a request
type Resource struct {
	HttpMethod string
	URL        string
	Query      map[string]string
	Data       []byte
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
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodGet,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       nil,
	})
}

// Post
func (requests *Requests) Post(endpoint string, query map[string]string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodPost,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       data,
	})
}

// Put
func (requests *Requests) Put(endpoint string, query map[string]string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodPut,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       data,
	})
}

// Patch
func (requests *Requests) Patch(endpoint string, query map[string]string, data []byte) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodPatch,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       data,
	})
}

// Delete
func (requests *Requests) Delete(endpoint string, query map[string]string) (*ResponseData, error) {
	return requests.handleRequestWithRetry(Resource{
		HttpMethod: http.MethodDelete,
		URL:        endpoint,
		Query:      queryChecker(query),
		Data:       nil,
	})
}
