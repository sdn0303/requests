# requests
The http Request CRUD Wrapper for Golang.

## Getting Started
```bash
go get github.com/sdn0303/requests
```

## Usage
Sample get
```go
package main

import (
    "fmt"

    "github.com/sdn0303/requests"
)

func main() {
	endpoint := "https://google.com"
	
	req := requests.New(requests.TimeOut(15))
	resp, _ := req.Get(endpoint, map[string]string{})
	
	for k, v := range resp.Headers {
		fmt.Printf("%s: %s", k, v)
	}
	fmt.Println(string(resp.Body))
	fmt.Println(resp.Status)
	fmt.Println(resp.StatusCode)
}
```