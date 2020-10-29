package files

import (
	"io"
	"os"
)

var fs fileSystem = osFS{}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

type fileSystem interface {
	Open(name string) (file, error)
}

type osFS struct{}

func (osFS) Open(name string) (file, error) {
	return os.Open(name)
}

//FileRead is responsible for load services
type FileRead struct {
	os fileSystem
}

//Read is responsible for read file and return the bytes
func (fileRead FileRead) Read(path string) ([][]byte, error) {
	handle, err := fileRead.os.Open(path)
	if err != nil {
		return nil, err
	}
	defer handle.Close()

	return nil, nil
}
