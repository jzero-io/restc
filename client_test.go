package restc

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	restc, err := NewClient("http://127.0.0.1:8080", WithRequestMiddleware(func(c Client, request *Request) error {
		request.AddHeader("test", "test")
		return nil
	}))
	if err != nil {
		assert.Nil(t, err)
	}
	do, err := restc.Verb("GET").Path("/api/v1/version").Do(context.Background()).RawResponse()
	fmt.Print(do)
}
