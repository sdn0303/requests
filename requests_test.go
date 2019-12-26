package requests

import (
	"encoding/json"
	"testing"
)

const baseURL = "http://httpbin.org"

func TestRequests_Get(t *testing.T) {
	requests := New(baseURL)
	data, err := requests.Get("get", map[string]string{"show_env": "1"})
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", data.Headers)
	t.Logf("Body: %s", string(data.Body))
	t.Logf("StatusCode: %v", data.StatusCode)
	t.Logf("Status: %s", data.Status)
}

func TestRequests_Post(t *testing.T) {
	requests := New(baseURL)
	requests.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	body := `{"post": "testing request"}`
	b, err := json.Marshal(body)
	if err != nil {
		t.Error(err)
	}

	resp, err := requests.Post("post", b)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}
