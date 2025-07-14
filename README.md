# restc

> 该仓库被废弃⚠️ 请使用: https://github.com/jzero-io/jzero/tree/main/core/restc

rest client for calling apis

## example

```go
package main
import (
	"context"
	
	"github.com/jzero-io/restc"
)

func main() {
	restClient, err := NewClient("http://127.0.0.1:8080", WithRequestMiddleware(func(c Client, request *Request) error {
		request.AddHeader("test", "test")
		return nil
	}))

	result := restClient.Verb("GET").Path("/api/v1/version").Do(context.Background())
	if result.Error() != nil {
		panic(err)
	}
	println(result.Status())
}
```
