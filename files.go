package filestorage

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Geter interface {
	Get(name string) ([]byte, error)
}

type FileStorage struct {
	Files map[string]string
}

func NewFileStorage(path string) (*FileStorage, error) {
	fs := &FileStorage{
		Files: make(map[string]string),
	}
	return fs, fs.load(path)
}

func (fs *FileStorage) Get(name string) ([]byte, error) {

	if r, ok := fs.Files[name]; ok {
		return ioutil.ReadFile(r)
	}

	return nil, errors.New("file not found")
}

func (fs *FileStorage) load(path string) error {
	return filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fs.Files[info.Name()] = path
		}
		return nil
	})
}
