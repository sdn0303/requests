package requests

import (
	"testing"
)

const baseURL = "http://httpbin.org"

func TestRequests_Get(t *testing.T) {
	requests := New(baseURL)
	data, err := requests.Get("get", map[string]string{"show_env": "1"}, 3)
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

	resp, err := requests.Post("post", []byte("testing post"), 3)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}

func TestRequests_Put(t *testing.T) {
	requests := New(baseURL)
	requests.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requests.Put("put", []byte("testing put"), 3)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}

func TestRequests_Patch(t *testing.T) {
	requests := New(baseURL)
	requests.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requests.Patch("patch", []byte("testing patch"), 3)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}

func TestRequests_Delete(t *testing.T) {
	requests := New(baseURL)
	requests.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requests.Delete("delete", nil, 3)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}
