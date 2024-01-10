package restc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	_, err := New(WithUrl("http://127.0.0.1:8080"))
	if err != nil {
		assert.Nil(t, err)
	}

	_, err = New(WithUrl("https://ragingcd.cloud.jaronnie.com"))
	if err != nil {
		assert.Nil(t, err)
	}

	_, err = New(WithProtocol("http"), WithAddr("127.0.0.1"), WithPort("8080"))
	if err != nil {
		assert.Nil(t, err)
	}
}
