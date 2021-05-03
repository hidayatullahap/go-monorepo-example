package hello

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHello(t *testing.T) {
	expected := "Hello World from TEST"
	actual := GetHello("TEST")

	assert.Equal(t, actual, expected)
}
