package filestorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFileStorage(t *testing.T) {
	f, err := NewFileStorage("mock")
	assert.Nil(t, err)

	b, err := f.Get("mock.json")
	assert.Nil(t, err)
	assert.NotEmpty(t, b)
}
