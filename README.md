Requests
====

The http Request Wrapper for Golang.

## Licence

[MIT](https://github.com/sdn0303/requests/blob/master/LICENSE)

## Author

[sdn0303](https://github.com/sdn0303)

## Install

```bash
go get github.com/sdn0303/requests
```

## Usage

Sample GET

```go
package main

import (
    "fmt"

    "github.com/sdn0303/requests"
)

func main() {
	endpoint := "https://google.com"

	// Create New requests client and set timeout 15sec 
	req := requests.New(requests.TimeOut(15))
	resp, _ := req.Get(endpoint, map[string]string{})

	fmt.Println(string(resp.Body))
	fmt.Println(resp.Status)
}
```

## Contribution

1. Fork ([https://github.com/sdn0303/requests/fork](https://github.com/sdn0303/requests/fork))
2. Create a feature branch
3. Commit your changes
4. Rebase your local changes against the master branch
5. Run test suite with the `go test ./...` command and confirm that it passes
6. Run `gofmt -s`
7. Create new Pull Request


