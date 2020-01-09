// Copyright (c) 2020 Soichi David Nakahashi
// Released under the MIT license.
// see https://github.com/sdn0303/requests/blob/master/LICENSE
package requests

import "time"

// Option functional options
type Option func(*Requests)

// MaxRetry set retry limit on requests client
func MaxRetry(limit uint64) Option {
	return func(requests *Requests) {
		requests.RetryLimit = limit
	}
}

// TimeOut set time out limit on requests client
func TimeOut(limit int) Option {
	return func(requests *Requests) {
		requests.HttpClient.Timeout = time.Duration(limit) * time.Second
	}
}
