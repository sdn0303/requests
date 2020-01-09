// Copyright (c) 2020 Soichi David Nakahashi
// Released under the MIT license.
// see https://github.com/sdn0303/requests/blob/master/LICENSE
package requests

import (
	"fmt"
	"testing"
)

const baseURL = "http://httpbin.org"

func TestRequests_Get(t *testing.T) {
	requests := New()
	resp, err := requests.Get(fmt.Sprintf("%s/get", baseURL), map[string]string{"show_env": "1"})
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}

func TestRequests_Post(t *testing.T) {
	requests := New(TimeOut(60))
	requests.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requests.Post(fmt.Sprintf("%s/post", baseURL), nil, []byte("testing post"))
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}

func TestRequests_Put(t *testing.T) {
	requests := New(MaxRetry(5))
	requests.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requests.Put(fmt.Sprintf("%s/put", baseURL), nil, []byte("testing put"))
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}

func TestRequests_Patch(t *testing.T) {
	requests := New()
	requests.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requests.Patch(fmt.Sprintf("%s/patch", baseURL), nil, []byte("testing patch"))
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}

func TestRequests_Delete(t *testing.T) {
	requests := New()
	requests.Headers = map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := requests.Delete(fmt.Sprintf("%s/delete", baseURL), nil)
	if err != nil {
		t.Error(err)
	}

	t.Logf("Headers: %v", resp.Headers)
	t.Logf("Body: %s", string(resp.Body))
	t.Logf("StatusCode: %v", resp.StatusCode)
	t.Logf("Status: %s", resp.Status)
}
