package restc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := NewClient(WithUrl("http://127.0.0.1:8080"))
	if err != nil {
		assert.Nil(t, err)
	}

	_, err = NewClient(WithUrl("https://ragingcd.cloud.jaronnie.com"))
	if err != nil {
		assert.Nil(t, err)
	}

	_, err = NewClient(WithProtocol("http"), WithAddr("127.0.0.1"), WithPort("8080"))
	if err != nil {
		assert.Nil(t, err)
	}
}
