# requests
The http Request CRUD Wrapper for Golang.


## Examples
### Get
```go
package main

import (
    "github.com/sdn0303/requests"
    "github.com/labstack/gommon/log"

)

const baseURL = "http://httpbin.org"

func main() {
	
	req := requests.New(baseURL)
	data, err := req.Get("get", map[string]string{"show_env": "1"})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(string(data))
	
}
```

### Post
```go
package main

import (
	"encoding/json"
	"github.com/sdn0303/requests"
	"github.com/labstack/gommon/log"
)

const baseURL = "http://httpbin.org"

func main() {
	
	requests.Header = map[string]string{
		"Content-Type": "application/json",
	}

	body := `{"post": "testing request"}`
	b, err := json.Marshal(body)
	if err != nil {
		log.Fatal(err)
	}
	
	resp, err := requests.Post("post", b)
	if err != nil {
		log.Fatal(err)
	}
	
	log.Printf(string(resp))
}
```