package requests

import "time"

type Option func(*Requests)

func MaxRetry(limit uint64) Option {
	return func(requests *Requests) {
		requests.RetryLimit = limit
	}
}

func TimeOut(limit int) Option {
	return func(requests *Requests) {
		requests.HttpClient.Timeout = time.Duration(limit) * time.Second
	}
}
