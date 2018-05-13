package filestorage

import (
	"testing"

)

func TestNewFileStorage(t *testing.T) {
	f, err := NewFileStorage("mock")
	if err != nil{
		t.Error(err)
	}

	b, err := f.Get("mock.json")
	if err != nil{
		t.Error(err)
	}
	if len(b) <=0 {
		t.Error("file is empty")
	}
}
