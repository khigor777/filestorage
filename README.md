# filestorage
PreLoad files in dir and subdir, then just get []byte of filename.
```go
	f, err := NewFileStorage("mock")
	b, err := f.Get("mock.json")
```
