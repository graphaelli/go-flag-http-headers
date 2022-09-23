# Summary

`go-flag-http-headers` provides a simple way to add a command line flag for specifying HTTP headers.

## Example

```go
package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"

	headerflag "github.com/graphaelli/go-flag-http-headers"
)

func main() {
	hf := headerflag.New()
	flag.Var(hf, "header", "HTTP header, can be specified multiple times")
	flag.Parse()

	req, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
	for h, vs := range hf.Headers() {
		for _, v := range vs {
			req.Header.Add(h, v)
		}
	}
	b, _ := httputil.DumpRequest(req, false)
	fmt.Println(string(b))
}
```
