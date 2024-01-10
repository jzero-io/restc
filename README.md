# restc

rest client for calling apis

## example

```go
package main
import (
	"context"
	
	"github.com/jaronnie/restc"
)

func main() {
	restClient, err := restc.New(restc.WithUrl("http://127.0.0.1:8080"))
	if err != nil {
		panic(err)
	}

	result := restClient.Get().SubPath("/api/v1/version").Do(context.Background())
	if result.Error() != nil {
		panic(err)
	}
	println(result.Status())
}
```
